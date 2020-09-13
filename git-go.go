package main

import (
	"bufio"
	"flag"
	"fmt"
	"gopkg.in/ini.v1"
	"log"
	"os"
	"path/filepath"
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
		positionalArgs := catFileCmd.Args()
		if len(positionalArgs) != 2 {
			fmt.Println("wrong number of arguments")
			os.Exit(1)
		} else {
			catFile(positionalArgs[0], positionalArgs[1])
		}
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
		positionalArgs := initCmd.Args()
		if len(positionalArgs) == 0 {
			wd, err := os.Getwd()
			if err != nil {
				log.Fatal(err)
			}
			createRepository(wd)
		} else if len(positionalArgs) == 1 {
			createRepository(positionalArgs[0])
		} else {
			fmt.Println("too many arguments")
			os.Exit(1)
		}
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

func createRepository(path string) {
	repo := new(GitRepository)
	repo.init(path)

	_, err := os.Stat(repo.GitDir)
	if !os.IsNotExist(err) {
		log.Println(".git directory already exists")
	} else {
		objectsDir := fmt.Sprintf("%s/objects", repo.GitDir)
		tagsDir := fmt.Sprintf("%s/refs/tags", repo.GitDir)
		headsDir := fmt.Sprintf("%s/refs/heads", repo.GitDir)
		descriptionFile := fmt.Sprintf("%s/description", repo.GitDir)
		headFile := fmt.Sprintf("%s/HEAD", repo.GitDir)
		configFile := fmt.Sprintf("%s/config", repo.GitDir)

		err = os.MkdirAll(objectsDir, 0755)
		if err != nil {
			log.Fatal(err)
		}

		err = os.MkdirAll(tagsDir, 0755)
		if err != nil {
			log.Fatal(err)
		}

		err = os.MkdirAll(headsDir, 0755)
		if err != nil {
			log.Fatal(err)
		}

		file, err := os.Create(descriptionFile)
		if err != nil {
			log.Fatal(err)
		}

		writer := bufio.NewWriter(file)
		_, err = writer.WriteString("Unnamed repository; edit this file 'description' to name the repository.\n")
		if err != nil {
			log.Fatal(err)
		}
		writer.Flush()

		file, err = os.Create(headFile)
		if err != nil {
			log.Fatal(err)
		}

		writer = bufio.NewWriter(file)
		_, err = writer.WriteString("ref: refs/heads/main\n")
		if err != nil {
			log.Fatal(err)
		}
		writer.Flush()

		cfg := ini.Empty(ini.LoadOptions{})

		ini.PrettyFormat = false
		ini.PrettyEqual = true
		coreSection, err := cfg.NewSection("core")
		if err != nil {
			log.Fatal(err)
		}

		_, err = coreSection.NewKey("repositoryformatverion", "0")
		if err != nil {
			log.Fatal(err)

		}
		_, err = coreSection.NewKey("filemode", "true")
		if err != nil {
			log.Fatal(err)

		}
		_, err = coreSection.NewKey("bare", "false")
		if err != nil {
			log.Fatal(err)

		}
		_, err = coreSection.NewKey("logallrefupdates", "true")
		if err != nil {
			log.Fatal(err)

		}
		_, err = coreSection.NewKey("ignorecase", "true")
		if err != nil {
			log.Fatal(err)
		}

		_, err = coreSection.NewKey("precomposeunicode", "true")
		if err != nil {
			log.Fatal(err)
		}

		cfg.SaveToIndent(configFile, "    ")
	}
}

func catFile(objectType string, object string) {
	fmt.Println(objectType)
	fmt.Println(object)
}

type GitRepository struct {
	WorkTree string
	GitDir   string
}

func (r *GitRepository) init(path string) {
	r.WorkTree = path
	r.GitDir = filepath.Join(path, ".git")
}
