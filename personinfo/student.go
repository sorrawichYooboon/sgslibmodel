package personinfo

type Student struct {
	Name  string
	Age   int
	Score int
}

func (s *Student) NewStudent(name string, age int, score int) {
	s.Name = name
	s.Age = age
	s.Score = score
}

func (s *Student) GetStudentInfo() string {
	return s.Name + " is " + string(s.Age) + " years old and has a score of " + string(s.Score)
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
