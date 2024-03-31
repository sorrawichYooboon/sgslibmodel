package personinfo

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score int
}

func NewStudent(name string, age int, score int) *Student {
	return &Student{
		Name:  name,
		Age:   age,
		Score: score,
	}
}

func (s *Student) GetStudentInfo() string {
	return fmt.Sprintf("Name: %s, Age: %d, Score: %d, Grade: %s", s.Name, s.Age, s.Score, s.GetStudentGrade())
}

func (s *Student) GetStudentGrade() string {
	if s.Score >= 90 {
		return "A"
	} else if s.Score >= 80 {
		return "B"
	} else if s.Score >= 70 {
		return "C"
	} else if s.Score >= 60 {
		return "D"
	} else {
		return "F"
	}
}
