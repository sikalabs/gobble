package exec_utils

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/sikalabs/gobble/pkg/libtask"
	"github.com/sikalabs/gobble/pkg/utils/random_utils"
)

func RawExec(password string, cmd string, args ...string) error {
	c := exec.Command(cmd, args...)
	if password != "" {
		c.Stdin = bytes.NewBufferString(password + "\n")
	}
	err := c.Run()
	return err
}

func RawExecStdout(cmd string, args ...string) error {
	c := exec.Command(cmd, args...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	err := c.Run()
	return err
}

func Exec(taskInput libtask.TaskInput, cmd string, args ...string) error {
	if taskInput.Dry {
		dryArgs := []string{}
		for _, arg := range args {
			if strings.Contains(arg, " ") {
				dryArgs = append(dryArgs, "'"+arg+"'")
			} else {
				dryArgs = append(dryArgs, arg)
			}
		}
		if taskInput.SudoPassword != "" {
			fmt.Printf("echo %s | ", taskInput.SudoPassword)
		}
		fmt.Println(strings.Join(append([]string{cmd}, dryArgs...), " "))
		return nil
	}

	err := RawExec(taskInput.SudoPassword, cmd, args...)
	return err
}

func SSH(taskInput libtask.TaskInput, cmdArray ...string) error {
	args := append([]string{taskInput.SSHTarget}, cmdArray...)
	if taskInput.Sudo {
		args = append([]string{taskInput.SSHTarget, "sudo"}, cmdArray...)
	} else if taskInput.SudoPassword != "" {
		args = append([]string{taskInput.SSHTarget, "sudo", "-S"}, cmdArray...)
	}
	if taskInput.NoStrictHostKeyChecking {
		args = append([]string{"-o", "StrictHostKeyChecking=no"}, args...)
	}
	if taskInput.SSHPort != 0 {
		args = append([]string{"-p", fmt.Sprintf("%d", taskInput.SSHPort)}, args...)
	}
	for _, option := range taskInput.SSHOptions {
		args = append([]string{"-o", option}, args...)
	}
	if taskInput.SSHPassword != "" {
		return Exec(taskInput, "sshpass", append([]string{"-p", taskInput.SSHPassword, "ssh"}, args...)...)
	} else {
		return Exec(taskInput, "ssh", args...)
	}
}

func rawSCP(taskInput libtask.TaskInput, localPath string, remotePath string) error {
	args := []string{localPath, taskInput.SSHTarget + ":" + remotePath}
	if taskInput.NoStrictHostKeyChecking {
		args = append([]string{"-o", "StrictHostKeyChecking=no"}, args...)
	}
	if taskInput.SSHPort != 0 {
		args = append([]string{"-P", fmt.Sprintf("%d", taskInput.SSHPort)}, args...)
	}
	if taskInput.SSHPassword != "" {
		return Exec(taskInput, "sshpass", append([]string{"-p", taskInput.SSHPassword, "scp"}, args...)...)
	} else {
		return Exec(taskInput, "scp", args...)
	}
}

func rawSCPRemoteToLocal(taskInput libtask.TaskInput, remotePath string, localPath string) error {
	args := []string{taskInput.SSHTarget + ":" + remotePath, localPath}
	if taskInput.NoStrictHostKeyChecking {
		args = append([]string{"-o", "StrictHostKeyChecking=no"}, args...)
	}
	if taskInput.SSHPort != 0 {
		args = append([]string{"-P", fmt.Sprintf("%d", taskInput.SSHPort)}, args...)
	}
	if taskInput.SSHPassword != "" {
		return Exec(taskInput, "sshpass", append([]string{"-p", taskInput.SSHPassword, "scp"}, args...)...)
	} else {
		return Exec(taskInput, "scp", args...)
	}
}

func SCP(taskInput libtask.TaskInput, localSrc string, remoteDst string) error {
	var err error
	tmpPath := "/tmp/" + random_utils.RandomString(32)
	err = rawSCP(taskInput, localSrc, tmpPath)
	if err != nil {
		return err
	}
	err = SSH(
		taskInput,
		"mv", tmpPath, remoteDst,
	)
	return err
}

func SCPRemoteToLocal(taskInput libtask.TaskInput, remoteSrc string, localDst string) error {
	var err error
	tmpPath := "/tmp/" + random_utils.RandomString(32)
	err = SSH(
		taskInput,
		"cp", remoteSrc, tmpPath,
	)
	if err != nil {
		return err
	}
	err = rawSCPRemoteToLocal(taskInput, tmpPath, localDst)
	return err
}
