package student

import "fmt"

type Student struct {
	Name  string
	Age   int
	Group string
}

func (s Student) ShowInfo() {
	fmt.Println(s.Name, s.Age, s.Group)
}
