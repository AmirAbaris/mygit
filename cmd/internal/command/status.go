package command

import "github.com/AmirAbaris/mygit/cmd/internal/repository"

func Status() string {
	repoExists := repository.Exists()

	if repoExists {
		return "mygit exists"
	}

	return "no mygit here"
}
