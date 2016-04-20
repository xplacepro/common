package common

func Mkfs_ext4(path string) (string, error) {
	return RunCommand("mkfs.ext4", []string{path})
}
