package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "GO"
const countdownStart = 3
const sleep = "sleep"
const write = "write"

type Sleeper interface {
	Sleep()
}

// for test
type countdownOperationsSpy struct {
	Calls []string
}

func (s *countdownOperationsSpy) Sleep()  {
	s.Calls = append(s.Calls, sleep)
}

func (s *countdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
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
