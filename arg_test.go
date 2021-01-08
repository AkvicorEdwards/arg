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
