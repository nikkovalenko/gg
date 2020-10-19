package main

import (
	"bufio"
	"compress/zlib"
	"flag"
	"fmt"
	"gopkg.in/ini.v1"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
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
		log.Fatal("expected a subcommand")
	}

	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		fmt.Printf("%s has been invoked\n", os.Args[1])
	case "cat-file":
		catFileCmd.Parse(os.Args[2:])
		positionalArgs := catFileCmd.Args()
		if len(positionalArgs) != 2 {
			log.Fatal("wrong number of arguments")
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
			log.Fatal("too many arguments")
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
		log.Fatalf("%s subcommand is not recognized\n", os.Args[1])
	}
}

func createRepository(path string) {
	repo := new(GitRepository)
	repo.init(path)

	_, err := os.Stat(repo.GitDir)
	if !os.IsNotExist(err) {
		log.Fatal(".git directory already exists")
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

func findRepository(path string) GitRepository {
	_, err := os.Stat(fmt.Sprintf("%s/.git", path))
	if os.IsNotExist(err) {
		parent := filepath.Dir(path)
		if parent == path {
			log.Fatal("not a git repository")
		} else {
			return findRepository(parent)
		}
	}

	repository := new(GitRepository)
	repository.init(path)

	return *repository
}

func catFile(objectType string, object string) {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	repository := findRepository(wd)

	obj := readObject(repository, object)
	fmt.Print(obj)
}

func readObject(repository GitRepository, sha string) string {
	path := fmt.Sprintf("%s/objects/%s/%s", repository.GitDir, sha[0:2], sha[2:])

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	reader, err := zlib.NewReader(file)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	b, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}

	raw := string(b)
	x := strings.Index(raw, " ")
	y := strings.Index(raw, "\x00")
	size, err := strconv.Atoi(raw[x+1 : y])
	if err != nil {
		log.Fatal(err)
	}

	if size != len(raw)-y-1 {
		log.Fatal("incorrect object size")
	}

	return raw[y+1:]
}

type GitRepository struct {
	WorkTree string
	GitDir   string
}

func (r *GitRepository) init(path string) {
	r.WorkTree = path
	r.GitDir = filepath.Join(path, ".git")
}
