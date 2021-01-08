package arg

import (
	"fmt"
	"sort"
)

type FuncErrorHandler func(error) error
type FuncExecutor func([]string) error

// Work Queue item
type work struct {
	// Priority for work
	Priority int
	// Work Executor
	Executor FuncExecutor
	// Error Handler, If Executor return err
	ErrorHandler FuncErrorHandler
	// Arguments for Executor
	Args []string
}

// Work Queue
type workQueue []work

// Sort Work Queue, Order by Priority DESC
func (q *workQueue) sort() {
	sort.Slice((*q)[:], func(i, j int) bool {
		return (*q)[i].Priority > (*q)[j].Priority
	})
}

// Add a work to Work Queue
func (q *workQueue) add(priority int, executor FuncExecutor, errExecutor FuncErrorHandler, args []string) {
	*q = append(*q, work{
		Priority:     priority,
		Executor:     executor,
		ErrorHandler: errExecutor,
		Args:         args,
	})
}

// Exec All work, Start execution from the first item in the Work Queue
func (q *workQueue) exec() (err error) {
	for _, v := range *q {
		err = v.Executor(v.Args)
		if err != nil {
			return v.ErrorHandler(err)
		}
	}
	return nil
}

// Command
type Command struct {
	Name          string
	Father        string
	Describe      string
	DescribeBrief string
	Help          string
	Usage         string

	Options map[string]*Option

	Commands map[string]*Command

	// Number of parameters required
	// if Size=-1 All parameters that follow belong to this command
	Size         int
	Executor     FuncExecutor
	ErrorHandler FuncErrorHandler
}

// Print Help
func (c *Command) PrintHelp() {
	fmt.Print(c.Help)
}

// Generate Help
func (c *Command) GenerateHelp() {
	if len(c.Help) != 0 {
		return
	}
	if c.Commands != nil {
		for _, v := range c.Commands {
			v.GenerateHelp()
		}
	}
	if c.Options != nil {
		for _, v := range c.Options {
			v.GenerateHelp()
		}
	}
	fullName := func() (name string) {
		name = c.Father
		if len(name) == 0 {
			name = c.Name
		} else {
			name += " " + c.Name
		}
		return
	}()
	describe := fmt.Sprintf(TplDescribe, fullName, c.Describe)
	cmd := false
	opt := false
	usageHead := func() string {
		hCommand := ""
		hOptionUp := ""
		hOptionDown := ""
		if c.Executor != nil {
			if c.Options != nil && len(c.Options) != 0 {
				hOptionDown = fmt.Sprintf(TplCommandUsageOption, fullName)
				opt = true
			} else {
				hOptionUp = fmt.Sprintf(TplCommandUsageSelf, fullName, c.Usage)
			}
		}

		if c.Commands != nil && len(c.Commands) != 0 {
			hCommand = fmt.Sprintf(TplCommandUsageCommand, fullName)
			cmd = true
		}
		return fmt.Sprintf(TplUsage, hOptionUp+hCommand+hOptionDown)
	}()

	commands := ""
	if cmd {
		cmdLine := ""
		lMax := 0
		for _, v := range c.Commands {
			if lMax < len(v.Name) {
				lMax = len(v.Name)
			}
		}
		for _, v := range c.Commands {
			line := fmt.Sprintf(fmt.Sprintf(HTplLineCommand, lMax), v.Name, v.DescribeBrief)
			cmdLine += line
		}
		commands = fmt.Sprintf(HTplCommandList, cmdLine, fullName)
	}

	options := ""
	if opt {
		optLine := ""
		lMax := 0
		for _, v := range c.Options {
			if lMax < len(v.Name) {
				lMax = len(v.Name)
			}
		}
		for _, v := range c.Options {
			line := fmt.Sprintf(fmt.Sprintf(HTplLineOption, lMax, lMax), v.Name, v.Usage, " ", v.DescribeBrief)
			optLine += line
		}
		options = fmt.Sprintf(HTplOptionList, optLine, fullName)
	}
	c.Help = fmt.Sprintf(TplHelp, describe, usageHead, commands, options)
}

// Create a new Command
func NewCommand(name, father string) *Command {
	return &Command{
		Name:          name,
		Describe:      "",
		DescribeBrief: "",
		Father:        father,
		Help:          "",
		Usage:         "",
		Options:       nil,
		Commands:      nil,
		Size:          0,
		Executor:      nil,
		ErrorHandler:  nil,
	}
}

// Create a new Command
func NewCommandFull(name, father, describe, describeBrief, help, usage string, size int,
	executor FuncExecutor, errExecutor FuncErrorHandler) *Command {
	return &Command{
		Name:          name,
		Describe:      describe,
		DescribeBrief: describeBrief,
		Father:        father,
		Help:          help,
		Usage:         usage,
		Options:       nil,
		Commands:      nil,
		Size:          size,
		Executor:      executor,
		ErrorHandler:  errExecutor,
	}
}

// Create a new Command map
func NewCommands() map[string]*Command {
	return make(map[string]*Command)
}

// Option
type Option struct {
	Name          string
	Father        string
	Size          int
	Priority      int
	Describe      string
	DescribeBrief string
	Help          string
	Usage         string
	Executor      FuncExecutor
	ErrorExecutor FuncErrorHandler
}

// Print Help
func (o *Option) PrintHelp() {
	fmt.Print(o.Help)
}

// Generate Help
func (o *Option) GenerateHelp() {
	if len(o.Help) != 0 {
		return
	}
	fullName := func() (name string) {
		name = o.Father
		if len(name) == 0 {
			name = o.Name
		} else {
			name += " " + o.Name
		}
		return
	}()
	//describe := fmt.Sprintf(TplDescribe, fullName, o.Describe)
	o.Help = fmt.Sprintf(HTplOptionUsage, fullName+" "+o.Usage, o.Describe)
}

// Create a new Option
func NewOption(name, father string) *Option {
	return &Option{
		Name:          name,
		Father:        father,
		Size:          0,
		Priority:      1000,
		Describe:      "",
		DescribeBrief: "",
		Help:          "",
		Usage:         "",
		Executor:      nil,
		ErrorExecutor: nil,
	}
}

// Create a new Option
func NewOptionFull(name, father string, size, priority int, describe, describeBrief, help, usage string,
	executor FuncExecutor, errExecutor FuncErrorHandler) *Option {
	return &Option{
		Name:          name,
		Father:        father,
		Size:          size,
		Priority:      priority,
		Describe:      describe,
		DescribeBrief: describeBrief,
		Help:          help,
		Usage:         usage,
		Executor:      executor,
		ErrorExecutor: errExecutor,
	}
}

// Create a new Option map
func NewOptions() map[string]*Option {
	return make(map[string]*Option)
}
