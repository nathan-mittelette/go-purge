package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:                 "go-purge",
		Usage:                "Clean your system or directory",
		Suggest:              true,
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name:    "global",
				Aliases: []string{"g"},
				Usage:   "Clean global system",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "force",
						Aliases: []string{"f"},
						Usage:   "Force cleaning without confirmation",
					},
				},
				Action: cleanGlobal,
			},
			{
				Name:    "directory",
				Aliases: []string{"d"},
				Usage:   "Clean current sub-directories",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "force",
						Aliases: []string{"f"},
						Usage:   "Force cleaning without confirmation",
					},
				},
				Action: cleanDirectory,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		color.Red("Error:", err)
	}
}

func confirmAction(message string, force bool) bool {
	if force {
		return true
	}

	color.Magenta("%s [y/N]: ", message)

	var response string
	_, err := fmt.Scanln(&response)

	if err != nil {
		return false
	}

	return strings.ToLower(response) == "y"
}

func confirmCommand(force bool, name string, arg ...string) error {
	var command = name + " " + strings.Join(arg, " ")
	var message = "Do you want to execute this command?"
	if confirmAction(message+" ("+command+")", force) {
		return execCommand(name, arg...)
	}
	return nil
}

func execCommand(name string, arg ...string) error {
	s := spinner.New(spinner.CharSets[35], 100*time.Millisecond)
	s.Suffix = "     " + name + " " + strings.Join(arg, " ") + "    "
	s.Start()

	defer s.Stop()

	cmd := exec.Command(name, arg...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

func removeAll(path string) error {
	s := spinner.New(spinner.CharSets[35], 100*time.Millisecond)
	s.Suffix = "     Removing " + path + "    "
	s.Start()

	defer s.Stop()

	return os.RemoveAll(path)
}

func cleanGlobal(c *cli.Context) error {
	force := c.Bool("force")

	var err error

	// Podman check
	if commandExists("podman") {
		color.Cyan("Cleaning Podman...")

		// Check if Podman machine is already running
		out, err := exec.Command("podman", "machine", "ls", "--format", "{{.Running}}").Output()
		if err != nil {
			color.Red("Error checking Podman machine status")
			return err
		}
		if strings.TrimSpace(string(out)) != "true" {
			err = exec.Command("podman", "machine", "start").Run()
			if err != nil {
				color.Red("Error starting Podman machine")
				return err
			}
		} else {
			color.Cyan("Podman machine already running")
		}

		err = confirmCommand(force, "podman", "system", "prune", "--all", "--volumes", "-f")

		if err != nil {
			color.Red("Error cleaning Podman")
			return err
		}

		color.Green("Podman cleaned!")
	} else {
		color.Yellow("podman not found")
	}

	// Docker check
	if commandExists("docker") {
		color.Cyan("Cleaning Docker...")

		err = confirmCommand(force, "docker", "system", "prune", "--all", "--volumes", "-f")

		if err != nil {
			color.Red("Error cleaning Docker")
			return err
		}

		color.Green("Docker cleaned!")
	} else {
		color.Yellow("docker not found")
	}

	// Maven check
	homeDir, _ := os.UserHomeDir()
	m2Path := filepath.Join(homeDir, ".m2", "repository")
	if _, err := os.Stat(m2Path); err == nil {
		color.Cyan("Cleaning Maven repository...")

		if confirmAction("Do you want to delete the Maven repository?", force) {
			err = removeAll(m2Path)

			if err != nil {
				color.Red("Error cleaning Maven repository")
				return err
			}

			color.Green("Maven repository cleaned!")
		}
	} else {
		color.Yellow("Maven repository not found")
	}

	// Gradle check
	gradlePath := filepath.Join(homeDir, ".gradle", "caches")
	if _, err := os.Stat(gradlePath); err == nil {
		color.Cyan("Cleaning Gradle cache...")

		if confirmAction("Do you want to delete the Gradle cache?", force) {
			err = removeAll(gradlePath)

			if err != nil {
				color.Red("Error cleaning Gradle cache")
				return err
			}

			color.Green("Gradle cache cleaned!")
		}
	} else {
		color.Yellow("Gradle cache not found")
	}

	// Go check
	if commandExists("go") {
		color.Cyan("Cleaning Go cache...")
		err = confirmCommand(force, "go", "clean", "-cache", "-modcache", "-testcache")

		if err != nil {
			color.Red("Error cleaning Go cache")
			return err
		}

		color.Green("Go cache cleaned!")
	} else {
		color.Yellow("go not found")
	}

	// Brew check
	if commandExists("brew") {
		color.Cyan("Cleaning Brew...")

		err = confirmCommand(force, "brew", "cleanup")

		if err != nil {
			color.Red("Error cleaning Brew")
			return err
		}

		color.Green("Brew cleaned!")
	} else {
		color.Yellow("brew not found")
	}

	// NPM check
	if commandExists("npm") {
		color.Cyan("Cleaning NPM cache...")

		err = confirmCommand(force, "npm", "cache", "clean", "--force")

		if err != nil {
			color.Red("Error cleaning NPM cache")
			return err
		}

		color.Green("NPM cache cleaned!")
	} else {
		color.Yellow("npm not found")
	}

	// Yarn check
	if commandExists("yarn") {
		color.Cyan("Cleaning Yarn cache...")

		err = confirmCommand(force, "yarn", "cache", "clean")

		if err != nil {
			color.Red("Error cleaning Yarn cache")
			return err
		}

		color.Green("Yarn cache cleaned!")
	} else {
		color.Yellow("yarn not found")
	}

	// PNPM check
	if commandExists("pnpm") {
		color.Cyan("Cleaning PNPM cache...")

		err = confirmCommand(force, "pnpm", "store", "prune")

		if err != nil {
			color.Red("Error cleaning PNPM cache")
			return err
		}

		color.Green("PNPM cache cleaned!")
	} else {
		color.Yellow("pnpm not found")
	}

	color.Green("Global cleaning completed!")
	return nil
}

func cleanDirectory(c *cli.Context) error {
	force := c.Bool("force")
	dirsToClean := []string{".terraform", "node_modules", "target"}
	var err error

	for _, dir := range dirsToClean {

		if !force {
			color.Cyan(fmt.Sprintf("Looking for %v directories...", dir))

			err = execCommand("find", ".", "-type", "d", "-name", dir)

			if err != nil {
				color.Red("Error finding directories to clean")
				return err
			}
		}

		if confirmAction("Do you want to delete these directories?", force) {
			color.Cyan(fmt.Sprintf("Cleaning %v directories ...", dir))

			err = execCommand("find", ".", "-type", "d", "-name", dir, "-exec", "rm", "-rf", "{}", "+")

			if err != nil {
				color.Red("Error cleaning directories")
				return err
			}

			color.Green(fmt.Sprintf("Directories %v cleaned!", dir))
		}
	}

	color.Green("Directory cleaning completed!")
	return nil
}
