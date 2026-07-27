package repository

import (
	"os"
)

func Exists() bool {
	_, err := os.Stat(".mygit")
	return err == nil
}
