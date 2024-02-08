package school

import (
	"fmt"
)

// Define the Grade and School types here.
type Grade struct {
	level int
	names []string
}

type School struct {
	grades map[int]Grade
}

func New() *School {
	return &School{
		grades: make(map[int]Grade),
	}
}

func list(e []Grade) (s string) {
	for _, l := range e {
		s += fmt.Sprintln(l)
	}
	return s
}

func (s *School) Add(student string, grade int) {
	_, ok := s.grades[grade]
	if ok {
		s.grades[grade] = Grade{level: grade, names: append(s.grades[grade].names, student)}
	} else {
		newGrade := Grade{level: grade, names: []string{student}}
		s.grades[grade] = newGrade
	}
}

func (s *School) Grade(level int) []string {
	return s.grades[level].names
}

func (s *School) Enrollment() []Grade {
	result := []Grade{}
	for _, v := range s.grades {
		result = append(result, v)
	}
	for k1, v := range result {
		for k2, v2 := range result {
			if v.level < v2.level {
				result[k1], result[k2] = result[k2], result[k1]
			}
		}
	}
	return result
}

func TestSchool() {
	s := New()
	s.Add("Aimee", 2)
	s.Add("Qn", 2)
	s.Add("Quizra", 3)
	fmt.Println(s)
	getGrades := s.Grade(4)
	fmt.Println(getGrades)
	s.Add("Jennifer", 4)
	s.Add("Kareem", 6)
	s.Add("Christopher", 4)
	s.Add("Kyle", 3)
	fmt.Println(list(s.Enrollment()))
}
