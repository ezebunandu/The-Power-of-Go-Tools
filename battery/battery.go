package battery

import (
	"fmt"
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
