package pipeline_test

import (
	"bytes"
	"errors"
	"io"
	"testing"

	"github.com/ezebunandu/pipeline"
	"github.com/google/go-cmp/cmp"
)

func TestStdoutPrintsMessageToOutput(t *testing.T) {
	t.Parallel()
	want := "Hello, world\n"
	p := pipeline.FromString(want)
	buf := new(bytes.Buffer)
	p.Output = buf
	p.Stdout()
	if p.Error != nil {
		t.Fatal(p.Error)
	}
	got := buf.String()
	if !cmp.Equal(want, got) {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestStdoutPrintsNothingOnError(t *testing.T) {
	t.Parallel()
	p := pipeline.FromString("Hello, world\n")
	p.Error = errors.New("silly billy")
	buf := new(bytes.Buffer)
	p.Output = buf
	p.Stdout()
	got := buf.String()
	if got != "" {
		t.Errorf("want no output from Stdout after an error, got %q", got)
	}
}

func TestFromFile_ReadsAllDataFromFile(t *testing.T) {
	t.Parallel()
	want := []byte("Hello, world\n")
	p := pipeline.FromFile("testing/hello.txt")
	if p.Error != nil {
		t.Fatal(p.Error)
	}
	got, err := io.ReadAll(p.Reader)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestFromFile_SetsErrorGivenNonexistentFile(t *testing.T) {
	t.Parallel()
	p := pipeline.FromFile("doesnt-exist.txt")
	if p.Error == nil {
		t.Fatal("want error opening non-existent file, got nil")
	}
}

func TestStringReturnsPipeContent(t *testing.T) {
	t.Parallel()
	want := "Hello, world\n"
	p := pipeline.FromString(want)
	got, err := p.String()
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestStringReturnsErrorWhenPipeErrorSet(t *testing.T) {
	t.Parallel()
	p := pipeline.FromString("Hello, world\n")
	p.Error = errors.New("oh no")
	_, err := p.String()
	if err == nil {
		t.Fatal("want error when pipe error set but got nil")
	}
}
