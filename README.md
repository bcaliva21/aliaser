# Aliaser
Simple go program to alias files in a zsh shell

## Version Info
- Make sure go is >= version 1.20
- [Install go here](https://go.dev/doc/install)
- This program assumes you are use a zsh shell

### Getting Started
- pull down repo
- build the binary `go build aliaser.go`
- move the binary into your local bin `sudo mv aliaser /usr/local/bin/aliaser`
- restart terminal or run `source ~/.zshrc` for good measure

#### How to Use
- `-h, --help`  displays usage and options
- `-p, --path`  sets the path for your alias
- `-a, --alias` sets the alias for the specified path
- example `aliaser -p /Users/user_name/Desktop -a desk`
- restart terminal or run `source ~/.zshrc` for alias to take effect

