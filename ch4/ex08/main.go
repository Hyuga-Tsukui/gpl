package main

import (
	"errors"
	"fmt"
)

var ErrMyError = errors.New("my error")

func main() {
	statckerr := dynamicError()
	dynamicerr := dynamicError()

	if statckerr == dynamicerr {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}
}

func example1(x int) error {
	if x > 0 {
		return errors.New("x is lather 0")
	}
	return nil
}

func staticError() error {
	return ErrMyError
}

func dynamicError() error {
	return errors.New("my error")
}

func example2() error {
	err := example1(1)
	if err != nil {
		return fmt.Errorf("wrap error %w", err)
	}
	return nil
}
