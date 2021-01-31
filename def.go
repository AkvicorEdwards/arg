package arg

var Version = "Arg 1.0.0"

var TplNeedMoreArguments = "The %s [%s] requires %d arguments to execute\n"

var TplCommandUsageSelf = "        %s %s"
var TplCommandUsageCommand = "        %s <command> [arguments]\n"
var TplCommandUsageOption = "        %s <option>  [arguments]\n"

// TplHelp =====================================================
/*

Arg

  Arg is a flag parser

Usage:

		arg [filename]
		arg <command> [arguments]
		arg <option>  [arguments]

The commands are

		build    build a programme
		encrypt  encrypt a file

Usage "arg help <command>" for more information about a command

The options are

		-o  [filename]
			  Specify out file
		-e  [password]
			  The password for file

Usage "arg help <option>" for more information about a option

===========================================================================

TplDescribe Begin ---

TplDescribe End ---

TplUsage Begin ---

	# TplCommandUsageSelf Begin ---
	# TplCommandUsageSelf End ---
	# TplCommandUsageCommand Begin ---
	# TplCommandUsageCommand End ---
	# TplCommandUsageOption Begin ---
	# TplCommandUsageOption End ---

TplUsage End ---

HTplCommandList Begin ---

	# HTplLineCommand Begin ---
	# HTplLineCommand End ---
	# HTplLineCommand Begin ---
	# HTplLineCommand End ---

HTplCommandList End ---

HTplOptionList Begin ---

	# HTplLineOption Begin ---
	# HTplLineOption End ---
	# HTplLineOption Begin ---
	# HTplLineOption End ---

HTplOptionList End ---
*/
var TplHelp = `
%s

%s%s%s
`

// TplDescribe =====================================================
/*
Arg
	Arg is a arguments parser
*/
//var TplDescribe = "%s\n\n    %s"
var TplDescribeUp = "%s"
var TplDescribeDown = "\n\n    %s"

// TplUsage =====================================================
/*
Usage:
		arg <command> [arguments]
*/
var TplUsage = `Usage:

%s`

// HTplCommandList ======================================================
/*
The commands are:

        build      build a programme
        encrypt    encrypt a file

Use "%s help <command>" for more information about a command.
*/
var HTplCommandList = `
The commands are:

%s
Use "%s %s <command>" for more information about a command.
`

// HTplOptionList ======================================================
/*
The options are:

        -o [filename]
                Specify out file
        -e [password]
                Encrypt File

Use "%s help <option>" for more information about a command.
*/
var HTplOptionList = `
The options are:

%s
Use "%s %s <option>" for more information about a option.
`

// HTplOptionUsage ======================================================
/*
Usage: Arg -build [-o file.tgz]
#Usage: father name usage

build file to file.tgz
#describe
*/
var HTplOptionUsage = `
Usage: %s

%s
`

// HTplLineCommand ======================================================
/*
		%%-%ds  %%s
		%-8s    %s
		build   build to tgz
#       name    describeBrief
*/
var HTplLineCommand = `        %%-%ds  %%s
`

// HTplLineOption ======================================================
/*
	%%-%ds  %%s
		%%s
-------------------
	%-8s    %s
		%s
-------------------
	-o      [filename]
		Specify out file
--------------------
    name    usage
        describeBrief
*/
var HTplLineOption = `        %%-%ds  %%s
        %%%ds    %%s
`
