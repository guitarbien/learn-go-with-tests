package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "GO"
const countdownStart = 3

type Sleeper interface {
	Sleep()
}

// for test
type spySleeper struct {
	Calls int
}

func (s *spySleeper) Sleep()  {
	s.Calls++
}

// for production
type ConfigurableSleeper struct {
	duration time.Duration
}

func (o ConfigurableSleeper) Sleep()  {
	time.Sleep(o.duration)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
	}
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
	}

	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second}
	Countdown(os.Stdout, sleeper)
}
