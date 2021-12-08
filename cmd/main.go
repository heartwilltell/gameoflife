package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"time"

	"gameoflife/app/game"
	"gameoflife/app/game/console"
)

// Variables which are related to Version command.
// Should be specified by '-ldflags' during the build phase.
// Example:
// GOOS=linux GOARCH=amd64 go build -ldflags="-X main.Version=$VERSION \
//  -X main.Branch=$BRANCH \
//  -X main.Commit=$COMMIT \
//  -X main.Environment=$ENV" \
//  -o receiver
var (
	// Version is the current version of application.
	Version = "unknown"
	// Branch is the branch this binary built from.
	Branch = "unknown"
	// Commit is the commit this binary built from.
	Commit = "unknown"
	// Environment represents the environment if the application.
	Environment = "local"
	// BuildTime is the time this binary built.
	BuildTime = time.Now().Format(time.RFC822)
)

const (
	runCmd     = "run"
	versionCmd = "version"
)

func main() {
	flag.Parse()

	if len(os.Args[1:]) < 1 {
		fmt.Printf("game: '%s' or '%s' command expected\n", versionCmd, runCmd)
		os.Exit(1)
	}

	switch os.Args[1] {
	case versionCmd:
		if err := versionCommand(os.Args[2:]); err != nil {
			fmt.Printf("game: '%s' command: %v\n", versionCmd, err)
			os.Exit(1)
		}
	case runCmd:
		if err := runCommand(os.Args[2:]); err != nil {
			fmt.Printf("game: '%s' command: %v\n", runCmd, err)
			os.Exit(1)
		}
	default:
		fmt.Printf("game: on of the following command expected: '%v'\n", []string{runCmd, versionCmd})
		os.Exit(1)
	}
}

func runCommand(args []string) error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	cmd := flag.NewFlagSet(runCmd, flag.ExitOnError)
	if err := cmd.Parse(args); err != nil {
		return err
	}

	s := console.NewConsoleSimulator(os.Stdout, console.WithTimeout(500*time.Millisecond))
	if err := s.Simulate(ctx, game.Glider); err != nil && !errors.Is(err, context.Canceled) {
		return fmt.Errorf("simulation failed: %w", err)
	}

	return nil
}

func versionCommand(args []string) error {
	cmd := flag.NewFlagSet(versionCmd, flag.ExitOnError)
	if err := cmd.Parse(args); err != nil {
		return err
	}

	fmt.Printf("Version: %s\n", Version)
	fmt.Printf("Environment: %s\n", Environment)
	fmt.Printf("Built from: %s [%s]\n", Branch, Commit)
	fmt.Printf("Built on: %s\n", BuildTime)
	fmt.Printf("Built time: %v\n", time.Now().UTC())

	return nil
}
