package common

func Rsync(from, to string) (string, error) {
	return RunCommand("rsync", []string{"-avv", from, to})
}
