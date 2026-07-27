package command

import "github.com/AmirAbaris/mygit/cmd/internal/repository"

func Init() {
	repository.Create()
}
