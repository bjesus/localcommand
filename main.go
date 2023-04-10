package main

import (
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"strings"

	"github.com/adrg/xdg"
	"github.com/pelletier/go-toml"
)

type Config struct {
	AllowedCommands []string `toml:"allowed_commands"`
	Terminal        []string `toml:"terminal" default:[]`
	Interface       struct {
		Confirmation  string `toml:"confirmation"`
		Output        string `toml:"output"`
		DisplayOutput bool   `toml:"display_output" default:"true"`
	} `toml:"interface"`
}

func main() {
	configFilePath, _ := xdg.ConfigFile("localcommand/config.toml")
	config, err := loadConfig(configFilePath)
	if err != nil {
		fmt.Printf("Error loading config file from %s: %v\n", configFilePath, err)
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		fmt.Println("localcommand: missing command to run\nUsage: localcommand \"cmd://url-encoded-command\"")
		os.Exit(1)
	}

	command, err := url.QueryUnescape(os.Args[1][6:])

	if err != nil {
		fmt.Printf("Error decoding command: %v\n", err)
		os.Exit(1)
	}

	if !isAllowedCommand(command, config) {
		if !confirmCommand(command, config) {
			fmt.Println("Command not confirmed. Exiting.")
			os.Exit(1)
		}
	}

	output, err := runCommand(command, config)
	if err != nil {
		fmt.Printf("Error running command: %v\n", err)
		os.Exit(1)
	}

	if config.Interface.DisplayOutput {
		if strings.Contains(config.Interface.Output, "%s") {
			outputCommand := fmt.Sprintf(config.Interface.Output, output)
			err = exec.Command("sh", "-c", outputCommand).Run()
		} else {
			infoCmd := exec.Command("sh", "-c", config.Interface.Output)
			infoCmd.Stdin = strings.NewReader(output)
			err = infoCmd.Run()
		}
		if err != nil {
			fmt.Printf("Error showing info: %v\n", err)
			os.Exit(1)
		}
	}
}

func loadConfig(filename string) (Config, error) {
	var config Config
	configFile, err := os.Open(filename)
	if err != nil {
		return config, err
	}
	defer configFile.Close()

	decoder := toml.NewDecoder(configFile)
	err = decoder.Decode(&config)
	return config, err
}

func isAllowedCommand(command string, config Config) bool {
	args := strings.Split(command, " ")

	for _, allowedCommand := range config.AllowedCommands {
		if args[0] == allowedCommand {
			return true
		}
	}
	return false
}

func confirmCommand(command string, config Config) bool {
	questionCommand := fmt.Sprintf(config.Interface.Confirmation, command)
	_, err := exec.Command("sh", "-c", questionCommand).Output()
	if err != nil {
		return false
	}
	return true
}

func runCommand(command string, config Config) (string, error) {
	args := make([]string, 0)
	if len(config.Terminal) > 0 {
		args = append(args, config.Terminal...)
		args = append(args, "--")
	}
	parts := strings.Fields(command)
	args = append(args, parts...)
	cmd := exec.Command(args[0], args[1:]...)
	output, err := cmd.Output()
	return string(output), err
}
