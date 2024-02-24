package main

import (
	"flag"
    "fmt"
    "io/ioutil"
    "os"
    "strings"
)

var (
	path   string
	alias  string
	help   bool
)

func init() {
	flag.StringVar(&path, "p", "", "path to assign alias to")
	flag.StringVar(&path, "path", "", "path to assign alias to")
	flag.StringVar(&alias, "a", "", "alias to assign to path")
	flag.StringVar(&alias, "alias", "", "alias to assign to path")
	flag.BoolVar(&help, "h", false, "show help message and exit")
	flag.BoolVar(&help, "help", false, "show help message and exit")
}

func main() {
	flag.Parse()

	if help {
		fmt.Println("Usage: aliaser [options]")
        fmt.Println("Options:")
        fmt.Println("  -p, --path <path>      Path to assign alias to")
        fmt.Println("  -a, --alias <alias>  Alias to assign to path")
        fmt.Println("  -h, --help           Show help message and exit")
        os.Exit(0)
	}

	if path == "" || alias == "" {
		fmt.Println("Error: both -p/--path and -a/--alias flags are required")
		os.Exit(1)
	}

    // Append config to .zshrc
    zshrcPath := fmt.Sprintf("%s/.zshrc", os.Getenv("HOME"))
    contents, err := ioutil.ReadFile(zshrcPath)
    if err != nil {
        fmt.Println("Error reading .zshrc:", err)
        os.Exit(1)
    }

    newContents := fmt.Sprintf(`alias %s=%s`, alias, path)
    if !strings.Contains(string(contents), newContents) {
        contents = append(contents, []byte(newContents+"\n")...)
        err = ioutil.WriteFile(zshrcPath, contents, 0644)
        if err != nil {
            fmt.Println("Error writing to .zshrc:", err)
            os.Exit(1)
        }
    }

    fmt.Println("Alias created! You can now use", alias, "to access", path)
}

