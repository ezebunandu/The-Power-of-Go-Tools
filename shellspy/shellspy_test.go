package shellspy_test

import (
	"testing"

	"github.com/ezebunandu/shellspy"
	"github.com/google/go-cmp/cmp"
)

func TestCmdFromString_CreatesExpectedCmd(t *testing.T) {
	t.Parallel()
	cmd, err := shellspy.CmdFromString("ls -l")
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"ls", "-l"}
	got := cmd.Args
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestCmdFromString_ErrorsOnEmptyInput(t *testing.T) {
	t.Parallel()
	_, err := shellspy.CmdFromString("")
	if err == nil {
		t.Fatal("want error on emtpy input, got nil")
	}
}
