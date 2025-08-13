package handler

import (
	"obs/models"

	cla "github.com/cetinboran/goarg/CLA"
)

func Handle(args []cla.Input) {
	modeName := args[0].ModeName

	switch modeName {
	case "student":
		models.AddStart(args)
	case "grade":
		models.AddStart(args)
		break
	case "get":
		models.GetStart(args)
		break
	case "delete":
		models.DeleteStart(args)
		break
	}
}
