package common

import (
	"io/ioutil"
	"os"
	"strings"
)

func ReplaceInFile(fPath, old, repl string, perm os.FileMode) error {
	contents, err := ioutil.ReadFile(fPath)
	if err != nil {
		return err
	}
	newContents := strings.Replace(string(contents), old, repl, -1)

	if err := ioutil.WriteFile(fPath, []byte(newContents), perm); err != nil {
		return err
	}
	return nil
}
