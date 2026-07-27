package command

import "github.com/AmirAbaris/mygit/cmd/internal/repository"

func Init() error {
	return  repository.Create()

}
