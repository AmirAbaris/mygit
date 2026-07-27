package repository

import (
	"errors"
	"os"
)

func Create() error {
	_, err := os.Stat(".mygit")

	if err == nil {
		return errors.New("repository already initialized")
	}

	if !os.IsNotExist(err) {
		return err
	}

	dirs := []string{
		".mygit",
		".mygit/objects",
		".mygit/refs",
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	files := []string{
		".mygit/HEAD",
	}

	for _, file := range files {
		if err := os.WriteFile(file, []byte{}, 0644); err != nil {
			return err
		}
	}

	return nil
}
