package arg

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func Example() {
	var err error
	var Err1 = errors.New("error 1") // handled
	var Err2 = errors.New("error 2") // unhandled
	os.Args = []string{"Arg", "Akvicer", "-type", "err2"}

	AddHelpCommandArg("help")

	// Set Root Command
	RootCommand.Name = "Arg"
	RootCommand.Size = -1
	RootCommand.Describe = `Arg is a project for go,
to parse arguments and execute command
This project is free`
	RootCommand.DescribeBrief = "Arg is a arguments parser"
	RootCommand.Usage = "[arguments...]"
	RootCommand.Executor = func(str []string) error {
		fmt.Println("RootCommand.Executor is run", str)
		return nil
	}
	RootCommand.ErrorHandler = func(err error) error {
		fmt.Println("Handle Error:", err)
		return nil
	}

	err = AddOption([]string{"-type"}, 1, 10, "This is a type for test",
		"test type", "", "", func(str []string) error {
			fmt.Println("Enter type", str[1])
			if str[1] == "err1" {
				return Err1
			}
			if str[1] == "err2" {
				return Err2
			}
			return nil
		}, func(err error) error {
			fmt.Println("Handle err", err)
			if err == Err1 {
				return nil
			}
			if err == Err2 {
				return err
			}
			return nil
		})
	if err != nil {
		fmt.Println("1:", err)
	}

	RootCommand.GenerateHelp()

	err = Parse()
	if err != nil {
		log.Println("2:", err)
	}

	fmt.Println("Finished")

	// Output:
	// Enter type err2
	// Handle err error 2
	// Finished
}

func ExampleParse() {
	var err error
	var Err1 = errors.New("error 1") // handled
	var Err2 = errors.New("error 2") // unhandled
	os.Args = []string{"Arg", "build", "-type", "tgz", "fileA", "fileB"}

	AddHelpCommandArg("help")

	// Set Root Command
	RootCommand.Name = "Arg"
	RootCommand.Size = -1
	RootCommand.Describe = `Arg is a project for go,
to parse arguments and execute command
This project is free`
	RootCommand.DescribeBrief = "Arg is a arguments parser"
	RootCommand.Usage = "[arguments...]"
	RootCommand.Executor = func(str []string) error {
		fmt.Println("RootCommand.Executor is run", str)
		return nil
	}
	RootCommand.ErrorHandler = func(err error) error {
		fmt.Println("Handle Error:", err)
		return nil
	}

	err = AddCommand([]string{"build"}, 2, "build to fi", "build a file to fi", "",
		"[ori filename] [target filename]", func(str []string) error {
			fmt.Println("build", str[1:])
			return nil
		}, func(err error) error {
			fmt.Println("Handled build err:", err)
			return err
		})

	err = AddOption([]string{"build", "-type"}, 1, 10, "This is a type for test",
		"test type", "", "", func(str []string) error {
			fmt.Println("build type", str[1])
			if str[1] == "err1" {
				return Err1
			}
			if str[1] == "err2" {
				return Err2
			}
			return nil
		}, func(err error) error {
			fmt.Println("Handle build.type err", err)
			if err == Err1 {
				return nil
			}
			if err == Err2 {
				return err
			}
			return nil
		})
	if err != nil {
		fmt.Println("1:", err)
	}

	RootCommand.GenerateHelp()

	err = Parse()
	if err != nil {
		log.Println("2:", err)
	}

	fmt.Println("Finished")

	// Output:
	// build type tgz
	// build [fileA fileB]
	// Finished
}

func ExampleAddCommand() {
	err := AddCommand([]string{"version"}, 1, "check version", "check programme version",
		"", "[version]", func(str []string) error {
			if Version == str[1] {
				return nil
			}
			return errors.New("check version failure")
		}, func(err error) error {
			log.Println("Handled Command Err:", err)
			return nil
		})
	if err != nil {
		panic(err)
	}
}

func ExampleAddOption() {
	err := AddCommand([]string{"-phone"}, -1, "specify phone number",
		"specify user phone number", "", "[phone number]", func(str []string) error {
			for k, v := range str {
				if len(v) != 11 {
					return errors.New(fmt.Sprintf("check phone number on [%d]: [%s]", k, v))
				}
			}
			return nil
		}, func(err error) error {
			log.Println("Handled Option Err:", err)
			return err
		})
	if err != nil {
		panic(err)
	}
}

func ExampleAdd() {
	// Add command
	_ = Add(true, []string{"phone"}, -1, 0 /* invalid parameter */, "specify phone number",
		"specify user phone number", "", "[phone number]", func(str []string) error {
			for k, v := range str {
				if len(v) != 11 {
					return errors.New(fmt.Sprintf("check phone number on [%d]: [%s]", k, v))
				}
			}
			return nil
		}, func(err error) error {
			log.Println("Handled Option Err:", err)
			return err
		})
	// Add option
	_ = Add(false, []string{"-phone"}, -1, 100, "specify phone number",
		"specify user phone number", "", "[phone number]", func(str []string) error {
			for k, v := range str {
				if len(v) != 11 {
					return errors.New(fmt.Sprintf("check phone number on [%d]: [%s]", k, v))
				}
			}
			return nil
		}, func(err error) error {
			log.Println("Handled Option Err:", err)
			return err
		})
}

func ExampleEnableOptionCombination() {
	// Example '-'
	os.Args = []string{"fi", "-pdwa", "Akvicor"}
	queue = make(workQueue, 0)
	RootCommand = NewCommand("fi", "")
	commandArgs = []string{"fi"}
	command = RootCommand
	RootCommand.Name = "fi"
	RootCommand.Size = -1
	RootCommand.Executor = func(str []string) error {
		fmt.Println("Root Command", str)
		return nil
	}
	_ = Add(false, []string{"-p"}, 0, 100, "",
	"", "", "", func(str []string) error {
			fmt.Println("Enter -p:", str[0])
			return nil
		}, nil)
	_ = Add(false, []string{"-d"}, 0, 100, "",
	"", "", "", func(str []string) error {
			fmt.Println("Enter -d:", str[0])
			return nil
		}, nil)
	_ = Add(false, []string{"-w"}, 0, 100, "",
	"", "", "", func(str []string) error {
			fmt.Println("Enter -w:", str[0])
			return nil
		}, nil)
	_ = Add(false, []string{"-a"}, 0, 100, "",
	"", "", "", func(str []string) error {
			fmt.Println("Enter -a:", str[0])
			return nil
		}, nil)
	EnableOptionCombination()
	err := Parse()
	if err != nil {
		fmt.Println(err)
	}

	// Example ' '
	queue = make(workQueue, 0)
	commandArgs = []string{"fi"}
	os.Args = []string{"fi", "pdwa", "Akvicor"}
	_ = Add(false, []string{"p"}, 0, 100, "",
		"", "", "", func(str []string) error {
			fmt.Println("Enter p:", str[0])
			return nil
		}, nil)
	_ = Add(false, []string{"d"}, 0, 100, "",
		"", "", "", func(str []string) error {
			fmt.Println("Enter d:", str[0])
			return nil
		}, nil)
	_ = Add(false, []string{"w"}, 0, 100, "",
		"", "", "", func(str []string) error {
			fmt.Println("Enter w:", str[0])
			return nil
		}, nil)
	_ = Add(false, []string{"a"}, 0, 100, "",
		"", "", "", func(str []string) error {
			fmt.Println("Enter a:", str[0])
			return nil
		}, nil)
	OptionCombination = ' '
	err = Parse()
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Enter -p: -p
	// Enter -d: -d
	// Enter -w: -w
	// Enter -a: -a
	// Root Command [fi Akvicor]
	// Enter p: p
	// Enter d: d
	// Enter w: w
	// Enter a: a
	// Root Command [fi Akvicor]
}

func ExampleAddHelpCommandArg() {
	os.Args = []string{"fi", "help"}
	queue = make(workQueue, 0)
	RootCommand = NewCommand("fi", "")
	commandArgs = []string{"fi"}
	command = RootCommand
	RootCommand.Name = "fi"
	RootCommand.Describe = "this is describe"

	RootCommand.Executor = func(str []string) error {
		fmt.Println("Root Command", str)
		return nil
	}

	err := AddCommand([]string{"version"}, 1, "check version", "check programme version",
		"", "[version]", func(str []string) error {
			fmt.Println("version", str)
			return nil
		}, nil)
	if err != nil {
		panic(err)
	}
	err = AddOption([]string{"-phone"}, -1, 100, "specify phone number",
		"specify user phone number", "", "[phone number]", func(str []string) error {
			fmt.Println("Phone number:", str)
			return nil
		}, nil)
	if err != nil {
		panic(err)
	}
	err = AddCommand([]string{"version", "v1"}, 0, "display version 1", "dis v1", "",
		"u1", func(str []string) error {
			fmt.Println("version 1")
			return nil
		}, nil)
	err = AddCommand([]string{"version", "v2"}, 0, "display version 2", "dis v2", "",
		"", func(str []string) error {
			fmt.Println("version 2")
			return nil
		}, nil)
	AddHelpCommandArg("help")
	RootCommand.GenerateHelp()

	err = Parse()
	if err != nil {
		if err != ErrHelp {
			fmt.Println(err)
		}
	}

	fmt.Println("======================================")

	os.Args = []string{"fi", "help", "version"}
	err = Parse()
	if err != nil {
		if err != ErrHelp {
			fmt.Println(err)
		}
	}

	fmt.Println("======================================")

	os.Args = []string{"fi", "version", "help", "v1"}
	err = Parse()
	if err != nil {
		if err != ErrHelp {
			fmt.Println(err)
		}
	}

	//fi
	//
	//    this is describe
	//
	//Usage:
	//
	//        fi <command> [arguments]
	//        fi <option>  [arguments]
	//
	//The commands are:
	//
	//        version  check programme version
	//
	//Use "fi help <command>" for more information about a command.
	//
	//The options are:
	//
	//        -phone  [phone number]
	//                  specify user phone number
	//
	//Use "fi help <option>" for more information about a option.
	//
	//======================================
	//
	//fi version
	//
	//    check version
	//
	//Usage:
	//
	//        fi version [version]        fi version <command> [arguments]
	//
	//The commands are:
	//
	//        v1  dis v1
	//        v2  dis v2
	//
	//Use "fi version help <command>" for more information about a command.
	//
	//======================================
	//
	//fi version v1
	//
	//    display version 1
	//
	//Usage:
	//
	//        fi version v1 u1

}
