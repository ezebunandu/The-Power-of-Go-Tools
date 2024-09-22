package greet_test

import (
	"bytes"
	"errors"
	"testing"
	"testing/iotest"

	"github.com/ezebunandu/greet"
)

func TestGreetUser_PromptsUserForANameAndRendersGreeting(t * testing.T){
    t.Parallel()
    input := bytes.NewBufferString("Greg")
    output := new(bytes.Buffer)
    greet.GreetUser(input, output)
    got := output.String()
    want := "What is your name?\nHello, Greg.\n"
    if want != got {
        t.Fatalf("wanted %q but got %q", want, got)
    }
}

func TestGreetUser_PrintsHelloYouOnReadError(t *testing.T){
    t.Parallel()
    input := iotest.ErrReader(errors.New("bad reader"))
    output := new(bytes.Buffer)
    greet.GreetUser(input, output)
    got := output.String()
    want := "What is your name?\nHello, you.\n"
    if want != got {
        t.Fatalf("wanted %q but got %q", want, got)
    }
}