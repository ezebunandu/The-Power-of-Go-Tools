package battery

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
)

type Status struct {
	ChargePercent int
}

func ParsePmSetOutput(output string) (Status, error) {
	re := regexp.MustCompile(`(\d+)%`)
	matches := re.FindStringSubmatch(output)

	if len(matches) < 2 {
		return Status{}, fmt.Errorf("failed to parse pmset output: %q", output)
	}
	charge, err := strconv.Atoi(matches[1])
	if err != nil {
		return Status{}, fmt.Errorf("failed to parse charge percentage: %q", err)
	}
	return Status{ChargePercent: charge}, nil
}

func GetPmSetOutput() (string, error) {
	data, err := exec.Command("/usr/bin/pmset", "-g", "ps").CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(data), nil
}
