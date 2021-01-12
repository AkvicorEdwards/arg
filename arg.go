package arg

import (
	"fmt"
	"os"
)

// Default Root Command, init by os.Args[0]
var RootCommand = NewCommand(os.Args[0], "")

// The specified "help" parameters
var HelpCommandArgs = make(map[string]bool)

// Work Queue
var queue = make(workQueue, 0)

var OptionCombination int32 = 0

// The Command to be executed
var command = RootCommand

// The Arguments of the Command
var commandArgs = []string{command.Name}

func AddHelpCommandArg(h string) {
	HelpCommandArgs[h] = true
}

// AddCommand
//
// add a command to RootCommand
//
// arg is the path for command, like "go mod download" is []string{"mod", "download"}
//	size  is the number of arguments
func AddCommand(arg []string, size int, describe, describeBrief, help, usage string,
	executor FuncExecutor, errExecutor FuncErrorHandler) error {
	return Add(true, arg, size, 0, describe, describeBrief, help, usage, executor, errExecutor)
}

// AddOption
//
// add a option to RootCommand
//
//	arg  is the path for option, like "go mod -version" is []string{"mod", "-version"}
//	size  is the number of arguments
//	priority  is the execution priority
func AddOption(arg []string, size, priority int, describe, describeBrief, help, usage string,
	executor FuncExecutor, errExecutor FuncErrorHandler) error {
	return Add(false, arg, size, priority, describe, describeBrief, help, usage, executor, errExecutor)
}

// Add
//
// add a Command or Option to RootCommand
//
// You can use this function through AddCommand and AddOption
func Add(isCmd bool, arg []string, size, priority int, describe, describeBrief, help, usage string,
	executor FuncExecutor, errExecutor FuncErrorHandler) error {
	var args = RootCommand
	argLength := len(arg) - 1
	father := RootCommand.Name

	for k, v := range arg {
		if k == argLength {
			if isCmd {
				if args.Commands == nil {
					args.Commands = NewCommands()
				}
				args.Commands[v] = NewCommandFull(v, father, describe, describeBrief, help,
					usage, size, executor, errExecutor)
			} else {
				if args.Options == nil {
					args.Options = NewOptions()
				}
				args.Options[v] = NewOptionFull(v, father, size, priority, describe, describeBrief,
					help, usage, executor, errExecutor)
			}
			return nil
		}
		if args.Commands != nil {
			if c, ok := args.Commands[v]; ok {
				args = c
				father += " " + v
				continue
			}
		}
		return ErrWrongArgPath
	}
	return nil
}

func EnableOptionCombination() {
	OptionCombination = '-'
}

// Parse
//
// Parse os.Args use RootCommand
func Parse() (err error) {
	err = parse(command, os.Args[1:])
	if err != nil {
		return err
	}
	if command.Size != -1 && len(commandArgs) != command.Size+1 {
		if command.ErrorHandler == nil {
			fmt.Printf(TplNeedMoreArguments, "command", command.Name, command.Size)
			return ErrNeedMoreArguments
		} else if err = command.ErrorHandler(ErrNeedMoreArguments); err != nil {
			return err
		}
		return nil
	}
	queue.sort()
	err = queue.exec()
	if err != nil {
		return err
	}
	if command.Executor == nil {
		return nil
	}
	err = command.Executor(commandArgs)
	if err != nil {
		if command.ErrorHandler != nil {
			err = command.ErrorHandler(err)
		}
	}
	return err
}

// Parse "args" use "cmd"
func parse(cmd *Command, args []string) error {
	if cmd == nil || len(args) == 0 {
		return nil
	}
	if cmd.Commands != nil {
		if c, ok := cmd.Commands[args[0]]; ok {
			// clear queue
			queue = make(workQueue, 0)
			// reset command
			command = c
			// reset command args
			commandArgs = []string{c.Name}

			return parse(c, args[1:])
		}
	}

	if cmd.Options != nil {
		opt, ok := cmd.Options[args[0]]
		if ok {
			if opt.Size == -1 {
				queue.add(opt.Priority, opt.Executor, opt.ErrorExecutor, args[:])
				return nil
			}
			if len(args) < 1+opt.Size {
				if opt.ErrorExecutor == nil {
					fmt.Printf(TplNeedMoreArguments, "option", opt.Father+" "+opt.Name, opt.Size)
					return ErrNeedMoreArguments
				} else if err := opt.ErrorExecutor(ErrNeedMoreArguments); err != nil {
					return err
				}
				queue.add(opt.Priority, opt.Executor, opt.ErrorExecutor, args[:])
				return nil
			}
			queue.add(opt.Priority, opt.Executor, opt.ErrorExecutor, args[:1+opt.Size])
			return parse(cmd, args[1+opt.Size:])
		}
	}

	if h, ok := HelpCommandArgs[args[0]]; ok {
		if h && len(args) >= 2 {
			if cmd.Commands != nil {
				if c, ok := cmd.Commands[args[1]]; ok {
					c.PrintHelp()
					return ErrHelp
				}
			}
			if cmd.Options != nil {
				if c, ok := cmd.Options[args[1]]; ok {
					c.PrintHelp()
					return ErrHelp
				}
			}
		}
		cmd.PrintHelp()
		return ErrHelp
	}

	if OptionCombination != 0 {
		format := ""
		if OptionCombination == ' ' {
			format = "%c"
		} else {
			format = fmt.Sprintf("%c%%c", OptionCombination)
		}
		checked := true
		for k, v := range args[0] {
			if k == 0 && v == OptionCombination {
				continue
			}
			o, ok := cmd.Options[fmt.Sprintf(format, v)]
			if !ok || o.Size != 0 {
				checked = false
				break
			}
		}
		if checked {
			op := ""
			for k, v := range args[0] {
				if k == 0 && v == OptionCombination {
					continue
				}
				op = fmt.Sprintf(format, v)
				o, _ := cmd.Options[op]
				queue.add(o.Priority, o.Executor, o.ErrorExecutor, []string{op})
			}
			return parse(cmd, args[1:])
		}
	}

	commandArgs = append(commandArgs, args[0])
	return parse(cmd, args[1:])
}
