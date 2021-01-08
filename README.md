# Arg

A command-line arguments parser

# Installation

Use the alias "arg". To use docopt in your Go code:

```go
import "github.com/AkvicorEdwards/arg"
```

To install docopt in your $GOPATH:

```shell script
go get github.com/AkvicorEdwards/arg"
```

# Usage

```go
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

// Add a Command
err = AddCommand([]string{"build"}, 2, "build to fi", 
        "build a file to fi", "", "[ori filename] [target filename]", 
        func(str []string) error {
	fmt.Println("build", str[1:])
	return nil
}, func(err error) error {
	fmt.Println("Handled build err:", err)
	return err
})

// Add a Option
err = AddOption([]string{"build", "-type"}, 1, 10, 
        "This is a type for test", "test type", "", "", 
        func(str []string) error {
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

// Generate Help
RootCommand.GenerateHelp()

// The specified "help" parameters
AddHelpCommandArg("help")

// Parse
err = Parse()
```

After `RootCommand.GenerateHelp()`

```
Arg

    Arg is a project for go,
to parse arguments and execute command
This project is free

Usage:
        Arg [arguments...]        Arg <command> [arguments]

The commands are:

        build build a file to fi

Use "Arg help <command>" for more information about a command.
```

# API

## Package

```go
AddCommand(arg []string, size int, describe, describeBrief, help, 
	usage string, executor FuncExecutor, errExecutor FuncErrorHandler)
```

```go
AddOption(arg []string, size, priority int, describe, describeBrief, help,
	usage string, executor FuncExecutor, errExecutor FuncErrorHandler) 
```

```go
Add(isCmd bool, arg []string, size, priority int, describe, 
	describeBrief, help, usage string, executor FuncExecutor, 
	errExecutor FuncErrorHandler) 
```

```go
Parse()
```
```go
AddHelpCommandArg("help")
```

## Command

```go
PrintHelp()
```

```go
GenerateHelp()
```
