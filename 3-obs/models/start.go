package models

import (
	cla "github.com/cetinboran/goarg/CLA"
)

func AddStart(args []cla.Input) {
	add := InitAdd()
	add.TakeInputs(args)
	add.Start()
}

func DeleteStart(args []cla.Input) {
	delete := InitDelete()
	delete.TakeInputs(args)
	delete.Start()
}

func GetStart(args []cla.Input) {
	get := InitGet()
	get.TakeInputs(args)
	get.Start()

}
