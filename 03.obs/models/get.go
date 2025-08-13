package models

import (
	"fmt"
	"log"
	"obs/db"
	"os"
	"strconv"

	cla "github.com/cetinboran/goarg/CLA"
	"github.com/olekukonko/tablewriter"
)

type Get struct {
	StudentId   int
	ListStudent bool
	ListLessons bool
	ListGrades  bool
}

func InitGet() *Get {
	return &Get{}
}

func (g *Get) TakeInputs(args []cla.Input) {
	for _, arg := range args {
		if arg.Argument == "sL" {
			g.ListStudent = true
		}

		if arg.Argument == "lL" {
			g.ListLessons = true
		}

		if arg.Argument == "gL" {
			g.ListGrades = true
		}

		if arg.Argument == "id" {
			value, err := strconv.Atoi(arg.Value)
			if err != nil {
				log.Fatal("Enter a integer value.")
			}
			g.StudentId = value
		}
	}
}

func (g *Get) studentList() {
	Students := db.StudentT.Get()

	studentTable := tablewriter.NewWriter(os.Stdout)
	studentTable.SetHeader([]string{"Id", "Name", "Surname"})

	for _, student := range Students {
		studentTable.Append([]string{
			fmt.Sprintf("%.0f", student["id"]),
			student["name"].(string),
			student["surname"].(string),
		})
	}

	studentTable.Render()
}

func (g *Get) LessonList() {
	Lessons := db.LessonT.Get()

	lessonTable := tablewriter.NewWriter(os.Stdout)
	lessonTable.SetHeader([]string{"Id", "Name"})

	for _, lesson := range Lessons {
		lessonTable.Append([]string{
			fmt.Sprintf("%.0f", lesson["id"]),
			lesson["name"].(string),
		})
	}

	lessonTable.Render()
}

func (g *Get) GradeList() {
	student := db.StudentT.Find("id", g.StudentId)
	if student == nil {
		log.Fatal("There is no such student")
	}

	Grades := db.GradeT.Get()

	gradeTable := tablewriter.NewWriter(os.Stdout)
	gradeTable.SetHeader([]string{"Grade Id", "Student Name", "Student Surname", "Lesson Name", "Grade"})

	for _, grade := range Grades {
		// PK ya göre aradığım için dizi şeklinde gelse bile 1 adet gelicek
		// Bu yüzden direkt [0] olarak erişebilirim.
		lessonData := db.LessonT.Find("id", grade["lesson_id"])

		gradeTable.Append([]string{
			fmt.Sprintf("%.0f", grade["id"]),
			student[0]["name"].(string),
			student[0]["surname"].(string),
			lessonData[0]["name"].(string),
			fmt.Sprintf("%.0f", grade["grade"]),
		})
	}

	gradeTable.Render()
}

func (g *Get) Start() {
	if g.ListStudent {
		g.studentList()
	}

	if g.ListLessons {
		g.LessonList()
	}

	if g.ListGrades {
		g.GradeList()
	}
}
