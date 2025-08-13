package models

import (
	"log"
	"obs/db"
	"strconv"

	cla "github.com/cetinboran/goarg/CLA"
	"github.com/cetinboran/gojson/gojson"
)

type Add struct {
	LessonId  int
	StudentId int
	Grade     int
	Name      string
	Surname   string
	Mode      string
}

func InitAdd() *Add {
	return &Add{}
}

func (a *Add) takeStudentInputs(args []cla.Input) {
	for _, arg := range args {
		if arg.Argument == "name" || arg.Argument == "n" {
			a.Name = arg.Value
		}

		if arg.Argument == "surname" || arg.Argument == "sn" {
			a.Surname = arg.Value
		}
	}

	if a.Name == "" || a.Surname == "" {
		log.Fatal("Name or surname cannot be empty string")
	}
}

func (a *Add) takeGradeInputs(args []cla.Input) {
	// For error check
	a.Grade = -1

	for _, arg := range args {
		if arg.Argument == "sid" {
			value, err := strconv.Atoi(arg.Value)
			if err != nil {
				log.Fatal("Enter a integer value.")
			}
			a.StudentId = value
		}

		if arg.Argument == "lid" {
			value, err := strconv.Atoi(arg.Value)
			if err != nil {
				log.Fatal("Enter a integer value.")
			}
			a.LessonId = value
		}

		if arg.Argument == "g" || arg.Argument == "grade" {
			value, err := strconv.Atoi(arg.Value)
			if err != nil {
				log.Fatal("Enter a integer value.")
			}
			a.Grade = value
		}
	}

	if a.StudentId == 0 || a.Grade < 0 || a.LessonId == 0 {
		log.Fatal("Stuent Id, Lesson Id or Grade cannot be empty string")
	}
}

func (a *Add) TakeInputs(args []cla.Input) {
	modeName := args[0].ModeName

	a.Mode = modeName
	if modeName == "student" {
		a.takeStudentInputs(args)
	} else if modeName == "grade" {
		a.takeGradeInputs(args)
	}
}

func (a *Add) Start() {
	if a.Mode == "student" {
		studentData := gojson.DataInit([]string{"name", "surname"}, []interface{}{a.Name, a.Surname}, db.StudentT)
		db.StudentT.Save(studentData)
	} else if a.Mode == "grade" {
		student := db.StudentT.Find("id", a.StudentId)
		if student == nil {
			log.Fatal("There is no such student!")
		}

		lesson := db.LessonT.Find("id", a.LessonId)
		if lesson == nil {
			log.Fatal("There is no such lesson!")
		}

		gradeData := gojson.DataInit([]string{"student_id", "lesson_id", "grade"}, []interface{}{a.StudentId, a.LessonId, a.Grade}, db.GradeT)
		db.GradeT.Save(gradeData)
	}

}
