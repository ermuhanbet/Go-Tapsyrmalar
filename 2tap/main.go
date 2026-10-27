package main

import (
	"2tap/student"
)

func main() {
	s1 := student.Student{"Ali", 18, "IS-24A"}
	s2 := student.Student{"Nurasyl", 19, "IS-24B"}
	s3 := student.Student{"Ayan", 29, "IS-24A"}
	s1.ShowInfo()
	s2.ShowInfo()
	s3.ShowInfo()
}
