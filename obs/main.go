package main

import (
	"obs/db"
	"obs/handler"

	cla "github.com/cetinboran/goarg/CLA"
)

func main() {
	db.CreateDatabase()

	Setup := cla.Init()
	Setup.SetUsage("Student Manager", "You can manage students", []string{})

	// ADD MODE START
	Add := cla.ModInit()
	Add.SetUsage("Add Mode", "You can add student and grades", []string{})

	AddStudent := cla.ModInit()
	AddStudent.SetUsage("Add Student", "You can add student to the db.", []string{})

	AddStudent.AddOptionTitle("Options")
	AddStudent.AddOption("-n,--name", false, "Student Name")
	AddStudent.AddOption("-sn,--surname", false, "Student Surname")

	AddStudent.SetExamples([]string{"go run main.go -n Boran -sn Mesüm"})

	AddGrade := cla.ModInit()
	AddGrade.SetUsage("Add Grade", "You can add grade to the db.", []string{})

	AddGrade.AddOptionTitle("Options")
	AddGrade.AddOption("-lid", false, "Lesson Id")
	AddGrade.AddOption("-sid", false, "Student Id")
	AddGrade.AddOption("-g, --grade", false, "Student Grade")

	AddGrade.SetExamples([]string{"go run main.go -sid 1 -lid -3 -g -75"})

	Add.AddMode("grade", &AddGrade)
	Add.AddMode("student", &AddStudent)

	AddStudent.AutomaticUsage()
	AddGrade.AutomaticUsage()

	// ADD MODE END

	Delete := cla.ModInit()
	Delete.SetUsage("Delete Mode", "You can delete student or grade", []string{})

	Delete.AddOptionTitle("Options")
	Delete.AddOption("-sid", false, "Delete Student By Id")
	Delete.AddOption("-gid", false, "Delete Grade By Id")

	Delete.SetExamples([]string{"go run main.go delete -sid 1", "go run main.go delete -gid 1"})

	Get := cla.ModInit()
	Get.SetUsage("Get Mode", "You can list everyting", []string{})

	Get.AddOptionTitle("Options")
	Get.AddOption("-id", false, "Student Id For Listing Grades")
	Get.AddOption("-sL", true, "List Students")
	Get.AddOption("-lL", true, "List Lessons")
	Get.AddOption("-gL", true, "List Grades By Id")

	Get.SetExamples([]string{"go run main.go -sL", "go run main.go -lL", "go run main.go -id 1 -gL", "Or at the same time"})

	Setup.AddMode("delete", &Delete)
	Setup.AddMode("add", &Add)
	Setup.AddMode("get", &Get)

	Delete.AutomaticUsage()
	Add.AutomaticUsage()
	Get.AutomaticUsage()
	Setup.AutomaticUsage()

	args, _ := Setup.Start()
	handler.Handle(args)
}
