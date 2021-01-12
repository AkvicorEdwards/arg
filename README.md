# Arg

A command-line arguments parser

# Installation

Use the alias "arg". To use Arg in your Go code:

```go
import "github.com/AkvicorEdwards/arg"
```

To install Arg in your $GOPATH:

```shell script
go get "github.com/AkvicorEdwards/arg"
```

# Usage

## Template Code

```go
package main

import (
	"errors"
	"fmt"
	"github.com/AkvicorEdwards/arg"
)

var Err1 = errors.New("err1")
var Err2 = errors.New("err2")

func main() {
	// Set Root Command
	arg.RootCommand.Name = "Arg"
	arg.RootCommand.Size = -1
	arg.RootCommand.Describe = `Arg is a project for go,
to parse arguments and execute command
This project is free`
	arg.RootCommand.DescribeBrief = "Arg is a arguments parser"
	arg.RootCommand.Usage = "[arguments...]"
	arg.RootCommand.Executor = func(str []string) error {
		fmt.Println("RootCommand.Executor is run", str)
		return nil
	}
	arg.RootCommand.ErrorHandler = func(err error) error {
		fmt.Println("RootCommand.Handler Error:", err)
		return nil
	}

	// Add a Command
	vBuildType := ""
	vBuildDel := false
	err := arg.AddCommand([]string{"build"}, 2, "build to fi",
		"build a file to fi", "", "[ori filename] [target filename]",
		func(str []string) error {
			fmt.Println("build", vBuildType, str[1:])
			if vBuildDel {
				fmt.Println("build delete file", str[1:])
			}
			return nil
		}, func(err error) error {
			fmt.Println("Handled build err:", err)
			return err
		})

	// Add a Option
	err = arg.AddOption([]string{"build", "-type"}, 1, 10,
		"This is a type for test", "test type", "", "[type]",
		func(str []string) error {
			fmt.Println("build type", str[1])
			vBuildType = str[1]
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

	// Add a Option
	err = arg.AddOption([]string{"build", "-del"}, 0, 10,
		"delete origin file after build", "del origin file", "User design help", "",
		func(str []string) error {
			vBuildDel = true
			return nil
		}, nil)
	if err != nil {
		fmt.Println("2:", err)
	}

	arg.RootCommand.GenerateHelp()
	arg.AddHelpCommandArg("help")
	err = arg.Parse()
	if err != nil {
		fmt.Println("Parse Error:", err)
	}
	fmt.Println("Finished")
}
```

## Test

```shell script
go build
```

### RootCommand

```
$ ./aflag Akvicor
RootCommand.Executor is run [./aflag Akvicor]
Finished
```

### build

```
$ ./aflag build            
Handled build err: wrong number of arg
Parse Error: wrong number of arg
Finished
```

### build with args

```
$ ./aflag build file1 file2               
build  [file1 file2]
Finished
```

### build with -type

```
$ ./aflag build file1 file2 -type tgz     
build type tgz
build tgz [file1 file2]
Finished
```

### build with -type and -del

```
$ ./aflag build file1 file2 -type tgz -del
build type tgz
build tgz [file1 file2]
build delete file [file1 file2]
Finished
```

### build with err1

```
$ ./aflag build -type err1 file1 file2
build type err1
Handle build.type err err1
build fixed [file1 file2]
Finished
```

### build with err2

```
$ ./aflag build -type err2 file1 file2
build type err2
Handle build.type err err2
Parse Error: err2
Finished
```

### help

```
$ ./aflag help

Arg

    Arg is a project for go,
to parse arguments and execute command
This project is free

Usage:
        Arg [arguments...]        Arg <command> [arguments]

The commands are:

        build build a file to fi

Use "Arg help <command>" for more information about a command.

Parse Error: help
Finished

```

## OptionCombination

if an arguments is not

- Command
- Option
- "Help"

1. Each letter of this parameter will be prefixed and then
checked whether it is a parameter.
2. Option.Size must be equal to 0
3. if all the newly formed parameters are legal, execute them

`EnableOptionCombination()` equal `OptionCombination='-'`

if `OptionCombination = ' '`, set `prefix = ''`

**Example:**

`OptionCombination=' '` `arguments = "abcd"` = `a` `b` `c` `d`

`OptionCombination='-'` `arguments = "-abcd"` = `-a` `-b` `-c` `-d`

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
