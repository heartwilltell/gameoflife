package console

import (
	"context"
	"fmt"
	"io"
	"time"

	"gameoflife/app/game"
)

const (
	// generationTimeout represents default timeout
	// for Simulator struct.
	generationTimeout = 1 * time.Second
)

// NewConsoleSimulator takes io.Writer and returns
// a pointer to a new instance of Simulator.
func NewConsoleSimulator(w io.Writer, options ...Option) *Simulator {
	s := Simulator{
		w:  w,
		to: generationTimeout,
	}

	for _, option := range options {
		option(&s)
	}

	return &s
}

// Option changes default values for Simulator.
type Option func(*Simulator)

// WithTimeout changes default value of Simulator timeout field.
func WithTimeout(to time.Duration) Option { return func(s *Simulator) { s.to = to } }

// Simulator implements Simulator interface by
// writings Generation to the given io.Writer.
type Simulator struct {
	w  io.Writer
	to time.Duration
}

func (s *Simulator) Simulate(ctx context.Context, gen game.Generation) error {
	currentGen := gen

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			for _, i := range currentGen {
				fmt.Fprintf(s.w, "%v\n", i)
			}

			fmt.Println("===================================================")

			currentGen = currentGen.NextGeneration()
			time.Sleep(s.to)
		}
	}
}
