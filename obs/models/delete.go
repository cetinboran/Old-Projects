package models

import (
	"log"
	"obs/db"
	"strconv"

	cla "github.com/cetinboran/goarg/CLA"
)

type Delete struct {
	StudentId int
	GradeId   int
}

func InitDelete() *Delete {
	return &Delete{}
}

func (d *Delete) TakeInputs(args []cla.Input) {
	for _, arg := range args {
		if arg.Argument == "sid" {
			value, err := strconv.Atoi(arg.Value)
			if err != nil {
				log.Fatal("Enter a integer value.")
			}

			d.StudentId = value
		}

		if arg.Argument == "gid" {
			value, err := strconv.Atoi(arg.Value)
			if err != nil {
				log.Fatal("Enter a integer value.")
			}

			d.GradeId = value
		}
	}

	if d.GradeId != 0 && d.StudentId != 0 {
		log.Fatal("You can do one operation at a time")
	}
}

func (d *Delete) Start() {
	if d.StudentId != 0 {
		db.StudentT.Delete("id", d.StudentId)
	}

	if d.GradeId != 0 {
		db.GradeT.Delete("id", d.GradeId)
	}
}
