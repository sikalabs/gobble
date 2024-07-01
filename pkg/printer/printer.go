package printer

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Printer provides a concurrency-safe way to print structured output for plays, tasks, and hosts.
type Printer struct {
	mu         sync.Mutex // Mutex to protect prints that should not interleave
	playMu     sync.Mutex // Mutex for play-related prints
	taskMu     sync.Mutex // Mutex for task-related prints
	hostMu     sync.Mutex // Mutex for host-related prints
	print      sync.Mutex // Mutex for general prints
	deprecated sync.Mutex // Mutex for deprecated warnings
	playCount  int64      // Counter for the number of plays
	taskCount  int64      // Counter for the number of tasks
	hostCount  int64      // Counter for the number of hosts
	totalPlays int        // Total number of plays
	totalTasks int        // Total number of tasks in the current play
	totalHosts int        // Total number of hosts in the current task
	verbosity  int        // Verbosity level for the printer

}

// GlobalPrinter Global printer instance
var GlobalPrinter *Printer

// InitPrinter initializes the global printer with the given verbosity level.
func InitPrinter(verbosity int) {
	GlobalPrinter = &Printer{
		verbosity: verbosity,
	}
}

// SetPlayLength sets the total number of plays.
func (p *Printer) SetPlayLength(length int) {
	p.totalPlays = length
}

// SetTaskLength sets the total number of tasks for the current play.
func (p *Printer) SetTaskLength(length int) {
	p.totalTasks = length
	p.taskCount = 0
}

// SetHostLength sets the total number of hosts for the current task.
func (p *Printer) SetHostLength(length int) {
	p.totalHosts = length
	p.hostCount = 0
}

// PrintPlay prints play information in a structured and thread-safe manner.
func (p *Printer) PrintPlay(play string) {
	if p.verbosity > 0 {
		playIndex := atomic.AddInt64(&p.playCount, 1) // Increment play count
		p.playMu.Lock()
		fmt.Printf(formatPlay(play, playIndex, p.totalPlays))
		p.playMu.Unlock()
	}
}

// PrintTask prints task information in a structured and thread-safe manner.
func (p *Printer) PrintTask(task string) {
	if p.verbosity > 0 {
		p.taskMu.Lock()
		taskIndex := atomic.AddInt64(&p.taskCount, 1) // Increment task count
		fmt.Printf(formatTask(task, taskIndex, p.totalTasks))
		p.taskMu.Unlock()
	}
}

// PrintHost safely prints host information in a structured and thread-safe manner.
func (p *Printer) PrintHost(host, protocol string) {
	if p.verbosity > 0 {
		p.hostMu.Lock()
		hostIndex := atomic.AddInt64(&p.hostCount, 1) // Increment host count
		fmt.Printf(formatHost(host, protocol, hostIndex, p.totalHosts))
		p.hostMu.Unlock()
	}
}

// Print prints any data with optional format, checking the verbosity level.
func (p *Printer) Print(format string, text ...interface{}) {
	if p.verbosity > 0 {
		p.print.Lock()
		defer p.print.Unlock()
		if format == "" {
			// Print without formatting if format string is empty
			fmt.Println(text...)
		} else {
			// Print with formatting
			fmt.Printf(format, text...)
		}
	}
}

func (p *Printer) PrintDeprecated(text string) {
	if p.verbosity > 0 {
		p.deprecated.Lock()
		fmt.Printf(formatDeprecated(text))
		p.deprecated.Unlock()

	}
}
