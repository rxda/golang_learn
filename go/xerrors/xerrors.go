package main

import (
	"fmt"
	"golang.org/x/xerrors"
	"os"
	"time"
)

// MyError is an error implementation that includes a time and message.
type MyError struct {
	When time.Time
	What string
}

func (e MyError) Error() string {
	return fmt.Sprintf("%v: %v", e.When, e.What)
}

func oops() error {
	return MyError{
		time.Date(1989, 3, 15, 22, 30, 0, 0, time.UTC),
		"the file system has gone away",
	}
}

func main() {
	_, err := os.Open("non-existing")
	if err != nil {
		var pathError *os.PathError
		if xerrors.As(err, &pathError) {
			fmt.Println("Failed at path:", pathError.Path)
		}
	}
}





