package exec

import (
	"time"
	"errors"
)

var ErrGoTimeout = errors.New("GoTimeoutFunc")

func GoFunc(f func() error) chan error {
	ch := make(chan error)
	go func() {
		ch <- f()
	}()
	return ch
}

func GoTimeout(f func() error, timeout time.Duration) (err error) {
	done := make(chan bool)
	go func() {
		err = f()
		done <- true
	}()
	select {
	case <-time.After(timeout):
		return ErrGoTimeout
	case <-done:
		return
	}
}

func GoTimeoutFunc(timeout time.Duration, f func() error) chan error {
	ch := make(chan error)
	go func() {
		var err error
		select {
		case err = <-GoFunc(f):
			ch <- err
		case <-time.After(timeout):
			ch <- ErrGoTimeout
		}
	}()
	return ch
}

