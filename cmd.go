package common

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"unicode"
)

type RunError struct {
	err string
}

func (err RunError) Error() string {
	return err.err
}

func printCommand(cmd *exec.Cmd) {
	log.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
}

func printError(err error, message string) {
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("==> Error: %s, %s\n", err.Error(), message))
	}
}

func printOutput(outs string) {
	if len(outs) > 0 {
		log.Printf("==> Output: %s\n", outs)
	}
}

func RunCommand(cmd_name string, args []string) (string, error) {
	// Create an *exec.Cmd
	cmd := exec.Command(cmd_name, args...)

	printCommand(cmd)

	// Stdout buffer
	cmdOutput := &bytes.Buffer{}
	cmdErr := &bytes.Buffer{}
	// Attach buffer to command
	cmd.Stdout = cmdOutput
	cmd.Stderr = cmdErr

	err := cmd.Run() // will wait for command to return
	printError(err, string(cmdErr.Bytes()))
	result := string(cmdOutput.Bytes())
	printOutput(result)
	if err == nil {
		return result, nil
	}
	return result, RunError{strings.Trim(string(cmdErr.Bytes()), "\n")}

	// http://play.golang.org/p/_6xke11GMp
}

func ParseValues(data string, eq_sign rune, comment_sign byte) map[string]string {
	splitter := func(c rune) bool {
		return unicode.IsSpace(c) || c == eq_sign
	}

	values := make(map[string]string)
	lines := strings.Split(data, "\n")
	for _, line := range lines {
		splitted := strings.FieldsFunc(line, splitter)
		if len(splitted) >= 2 && len(splitted[0]) > 0 && splitted[0][0] != comment_sign {
			values[splitted[0]] = splitted[1]
		}
	}

	return values
}
