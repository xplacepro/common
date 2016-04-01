package common

import (
	"fmt"
)

func Dd(from, to string) (string, error) {
	return RunCommand("dd", []string{fmt.Sprintf("if=%s", from), fmt.Sprintf("of=%s", to)})
}
