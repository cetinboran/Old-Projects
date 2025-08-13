package db

import (
	"os"

	"github.com/cetinboran/gojson/gojson"
)

// Tables
var (
	StudentT *gojson.Table
	GradeT   *gojson.Table
	LessonT  *gojson.Table
)

func CreateDatabase() {
	var exists bool
	if _, err := os.Stat("OBS/Lesson.json"); err == nil {
		exists = true
	}

	Database := gojson.CreateDatabase("OBS", "./")

	Students := gojson.CreateTable("Student")
	Students.AddProperty("id", "int", "PK")
	Students.AddProperty("name", "string", "")
	Students.AddProperty("surname", "string", "")

	Grades := gojson.CreateTable("Grade")
	Grades.AddProperty("id", "int", "PK")
	Grades.AddProperty("student_id", "int", "") // FK
	Grades.AddProperty("lesson_id", "int", "")  // FK
	Grades.AddProperty("grade", "int", "")

	Lessons := gojson.CreateTable("Lesson")
	Lessons.AddProperty("id", "int", "PK")
	Lessons.AddProperty("name", "string", "")

	Database.AddTable(&Students)
	Database.AddTable(&Grades)
	Database.AddTable(&Lessons)

	// Table ismine dikkat et. Gojson da bunun kontrolünü yapmamışsın.
	StudentT = Database.Tables["Student"]
	GradeT = Database.Tables["Grade"]
	LessonT = Database.Tables["Lesson"]

	Database.CreateFiles()

	// Dummy Data
	if !exists {
		LessonT.Save(gojson.DataInit([]string{"name"}, []interface{}{"Cyber Security"}, LessonT))
		LessonT.Save(gojson.DataInit([]string{"name"}, []interface{}{"Mathematics"}, LessonT))
		LessonT.Save(gojson.DataInit([]string{"name"}, []interface{}{"Cryptology"}, LessonT))
		LessonT.Save(gojson.DataInit([]string{"name"}, []interface{}{"Data Structures"}, LessonT))
	}

}
