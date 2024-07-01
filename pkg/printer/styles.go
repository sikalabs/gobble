package printer

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"strconv"
)

// Define styles for the printer
var (
	// PlayStyle provides a style for plays
	PlayStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#FFA500", Dark: "#FFA500"}).
			Bold(true)

	// TaskStyle provides a style for tasks
	TaskStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#FFD700", Dark: "#FFD700"}).
			Bold(true).
			MarginLeft(2)

	// HostStyle provides a style for hosts
	HostStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#32CD32", Dark: "#32CD32"}).
			Bold(true).
			MarginLeft(4)

	// TextStyle provides a style for text
	TextStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#6A7FFF", Dark: "#6A7FFF"}).
			Bold(true).
			MarginLeft(1)

	// IndexStyle provides a style for index numbers
	IndexStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#4AB3FF", Dark: "#4AB3FF"}).
			Bold(true).
			MarginLeft(1)

	ProtocolStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#FFA500", Dark: "#FFA500"}).
			SetString("protocol:").
			MarginLeft(1)

	// DeprecatedStyle provides a style for plays
	DeprecatedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#FF0000", Dark: "#FF0000"}).
			Bold(true).SetString("DEPRECATED:")

	// BlockStyle provides a style for plays
	BlockStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#FF0000", Dark: "#FF0000"}).
			Bold(true).SetString("DEPRECATED:")
)

func formatHost(host, protocol string, hostIndex int64, totalHosts int) string {
	count := fmt.Sprintf("(%s)", strconv.FormatInt(hostIndex, 10)+"/"+strconv.Itoa(totalHosts))
	return HostStyle.Render("host:") + TextStyle.Render(host) + ProtocolStyle.Render(protocol) + IndexStyle.Render(count) + "\n"
}

func formatTask(task string, taskIndex int64, totalTasks int) string {
	count := fmt.Sprintf("(%s)", strconv.FormatInt(taskIndex, 10)+"/"+strconv.Itoa(totalTasks))
	return TaskStyle.Render("task:") + TextStyle.Render(task) + IndexStyle.Render(count) + "\n"
}

func formatPlay(play string, playIndex int64, totalPlays int) string {
	count := fmt.Sprintf("(%s)", strconv.FormatInt(playIndex, 10)+"/"+strconv.Itoa(totalPlays))
	return PlayStyle.Render("play:") + TextStyle.Render(play) + IndexStyle.Render(count) + "\n"
}

func formatDeprecated(text string) string {
	return DeprecatedStyle.Render(text) + "\n"
}
