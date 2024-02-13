package erratum

import (
	"errors"
	"fmt"
)

// Please review the README for this exercise carefully before implementation.

// Little helper to let us customize behaviour of the resource on a per-test
// basis.
type mockResource struct {
	close  func() error
	frob   func(string)
	defrob func(string)
}

const hello = "hello"

func (mr mockResource) Close() error      { return mr.close() }
func (mr mockResource) Frob(input string) { mr.frob(input) }
func (mr mockResource) Defrob(tag string) { mr.defrob(tag) }

// Use should not return an error on the "happy" path.
func TestNoErrors() {
	var frobInput string
	var closeCalled bool
	mr := mockResource{
		close: func() error { closeCalled = true; return nil },
		frob:  func(input string) { frobInput = input },
	}
	opener := func() (Resource, error) { return mr, nil }
	err := Use(opener, hello)
	if err != nil {
		fmt.Println("Unexpected error from Use: %v", err)
	}
	if frobInput != hello {
		fmt.Println("Wrong string passed to Frob: got %v, expected %v", frobInput, hello)
	}
	if !closeCalled {
		fmt.Println("Close was not called")
	}
}

// Use should keep trying if a transient error is returned on open.
func TestKeepTryOpenOnTransient() {
	var frobInput string
	mr := mockResource{
		close: func() error { return nil },
		frob:  func(input string) { frobInput = input },
	}
	nthCall := 0
	opener := func() (Resource, error) {
		if nthCall < 3 {
			nthCall++
			return mockResource{}, TransientError{errors.New("some error")}
		}
		return mr, nil
	}
	err := Use(opener, hello)
	if err != nil {
		fmt.Println("Unexpected error from Use: %v", err)
	}
	if frobInput != hello {
		fmt.Println("Wrong string passed to Frob: got %v, expected %v", frobInput, hello)
	}
	fmt.Println("error is", err)
}

// Use should fail if a non-transient error is returned on open.
func TestFailOpenOnNonTransient() {
	nthCall := 0
	opener := func() (Resource, error) {
		if nthCall < 3 {
			nthCall++
			return mockResource{}, TransientError{errors.New("some error")}
		}
		return nil, errors.New("too awesome")
	}
	err := Use(opener, hello)
	if err == nil {
		fmt.Println("Unexpected lack of error from Use")
	}
	if err.Error() != "too awesome" {
		fmt.Println("Invalid error returned from Use")
	}
	fmt.Println("error is", err)
}

// Use should call Defrob and Close on FrobError panic from Frob
// and return the error.
func TestCallDefrobAndCloseOnFrobError() {
	tag := "moo"
	var closeCalled bool
	var defrobTag string
	mr := mockResource{
		close: func() error { closeCalled = true; return nil },
		frob:  func(input string) { panic(FrobError{tag, errors.New("meh")}) },
		defrob: func(tag string) {
			if closeCalled {
				fmt.Println("Close was called before Defrob")
			}
			defrobTag = tag
		},
	}
	opener := func() (Resource, error) { return mr, nil }
	err := Use(opener, hello)
	if err == nil {
		fmt.Println("Unexpected lack of error from Use")
	}
	if err.Error() != "meh" {
		fmt.Println("Invalid error returned from Use")
	}
	if defrobTag != tag {
		fmt.Println("Wrong string passed to Defrob: got %v, expected %v", defrobTag, tag)
	}
	if !closeCalled {
		fmt.Println("Close was not called")
	}
	fmt.Println("error is", err)
}

// Use should call Close but not Defrob on non-FrobError panic from Frob
// and return the error.
func TestCallCloseOnNonFrobError() {
	var closeCalled bool
	var defrobCalled bool
	mr := mockResource{
		close:  func() error { closeCalled = true; return nil },
		frob:   func(input string) { panic(errors.New("meh")) },
		defrob: func(tag string) { defrobCalled = true },
	}
	opener := func() (Resource, error) { return mr, nil }
	err := Use(opener, hello)
	if err == nil {
		fmt.Println("Unexpected lack of error from Use")
	}
	if err.Error() != "meh" {
		fmt.Println("Invalid error returned from Use")
	}
	if defrobCalled {
		fmt.Println("Defrob was called")
	}
	if !closeCalled {
		fmt.Println("Close was not called")
	}
	fmt.Println("error is", err)
}
