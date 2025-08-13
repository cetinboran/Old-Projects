package core

import (
	inputhandler "github.com/cetinboran/ssher/input"
)

type IBruteForce interface {
	Start() error
}

type BruteForce struct {
	SSH         ISSH
	UserOptions []inputhandler.Arg
}

func NewBruteForce(server, port string, userOptions []inputhandler.Arg) IBruteForce {
	return &BruteForce{
		SSH:         NewSSH(server, port),
		UserOptions: userOptions,
	}
}

func (b *BruteForce) Start() error {
	return nil
}
