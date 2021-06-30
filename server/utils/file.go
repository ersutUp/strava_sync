package utils

import "os"

func CreateFolder(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(path, os.ModePerm)
			if  err != nil {
				return err
			}
		}
	}
	return nil
}
