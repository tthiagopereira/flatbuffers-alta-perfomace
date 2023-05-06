package main

import (
	"fmt"

	flatbuffers "github.com/google/flatbuffers/go"

	"teste/Teste"
)

func main() {
	builder := flatbuffers.NewBuilder(1024)

	coursesName := builder.CreateString("Flatbuffers teste")
	courseDescription := builder.CreateString("curso full go")

	Teste.CourseStart(builder)
	Teste.CourseAddId(builder, 1)
	Teste.CourseAddName(builder, coursesName)
	Teste.CourseAddDescription(builder, courseDescription)
	course := Teste.CourseEnd(builder)
	builder.Finish(course)

	buf := builder.FinishedBytes()
	c := Teste.GetRootAsCourse(buf, 0)

	fmt.Println(c.Id())
	fmt.Println(c.Name())
	fmt.Println(c.Description())
}
