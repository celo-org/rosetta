package fileutils

import "os"

func FileExists(filepath string) bool {
	_, err := os.Stat(filepath)
	return err == nil || !os.IsNotExist(err)
}

func IsDirectory(dirpath string) (bool, error) {
	stat, err := os.Stat(dirpath)
	if err != nil {
		return false, err
	}
	return stat.IsDir(), nil

}

func TouchFile(filepath string) error {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}
