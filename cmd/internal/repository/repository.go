package repository

import "os"

func Create() error {
	// directories
	dirs := []string{
		".git",
		".git/objects",
		".git/refs",
	}

	for _, dir := range dirs {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}

	// files
	files := []string{
		".git/HEAD",
	}

	for _, file := range files {
		err := os.WriteFile(file, []byte{}, 0644)
		if err != nil {
			return err
		}
	}

	return nil
}
