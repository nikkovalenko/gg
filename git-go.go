package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	catFileCmd := flag.NewFlagSet("cat-file", flag.ExitOnError)
	checkoutCmd := flag.NewFlagSet("checkout", flag.ExitOnError)
	commitCmd := flag.NewFlagSet("commit", flag.ExitOnError)
	hashObjectCmd := flag.NewFlagSet("hash-object", flag.ExitOnError)
	initCmd := flag.NewFlagSet("init", flag.ExitOnError)
	logCmd := flag.NewFlagSet("log", flag.ExitOnError)
	lsTreeCmd := flag.NewFlagSet("ls-tree", flag.ExitOnError)
	mergeCmd := flag.NewFlagSet("merge", flag.ExitOnError)
	rebaseCmd := flag.NewFlagSet("rebase", flag.ExitOnError)
	revParseCmd := flag.NewFlagSet("rev-parse", flag.ExitOnError)
	rmCmd := flag.NewFlagSet("rm", flag.ExitOnError)
	showRefCmd := flag.NewFlagSet("show-ref", flag.ExitOnError)
	tagCmd := flag.NewFlagSet("tag", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("expected a subcommand")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		fmt.Printf("%s has been invoked\n", os.Args[1])
	case "cat-file":
		catFileCmd.Parse(os.Args[2:])
		fmt.Printf("%s has been invoked\n", os.Args[1])
	case "checkout":
		checkoutCmd.Parse(os.Args[2:])
		fmt.Printf("%s has been invoked\n", os.Args[1])
	case "commit":
		commitCmd.Parse(os.Args[2:])
		fmt.Printf("%s has been invoked\n", os.Args[1])
	case "hash-object":
		hashObjectCmd.Parse(os.Args[2:])
		fmt.Printf("%s has been invoked\n", os.Args[1])
	case "init":
		initCmd.Parse(os.Args[2:])
		fmt.Printf("%s has been invoked\n", os.Args[1])
	case "log":
		logCmd.Parse(os.Args[2:])
		fmt.Printf("%s has been invoked\n", os.Args[1])
	case "ls-tree":
		lsTreeCmd.Parse(os.Args[2:])
		fmt.Printf("%s has been invoked\n", os.Args[1])
	case "merge":
		mergeCmd.Parse(os.Args[2:])
		fmt.Printf("%s has been invoked\n", os.Args[1])
	case "rebase":
		rebaseCmd.Parse(os.Args[2:])
		fmt.Printf("%s has been invoked\n", os.Args[1])
	case "rev-parse":
		revParseCmd.Parse(os.Args[2:])
		fmt.Printf("%s has been invoked\n", os.Args[1])
	case "rm":
		rmCmd.Parse(os.Args[2:])
		fmt.Printf("%s has been invoked\n", os.Args[1])
	case "show-ref":
		showRefCmd.Parse(os.Args[2:])
		fmt.Printf("%s has been invoked\n", os.Args[1])
	case "tag":
		tagCmd.Parse(os.Args[2:])
		fmt.Printf("%s has been invoked\n", os.Args[1])
	default:
		fmt.Printf("%s subcommand is not recognized\n", os.Args[1])
		os.Exit(1)
	}
}
