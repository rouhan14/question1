package main

import "fmt"

type Stack struct {
	s []string
}

func (s *Stack) IsEmpty() bool {
	return len(s.s) == 0
}

func (s *Stack) Length() int {
	return len(s.s)
}

func (s *Stack) Print() {
	for _, value := range s.s {
		fmt.Print(value, " ")
	}
	fmt.Println()
}

func (s *Stack) Push(value string) {
	s.s = append(s.s, value)
}

func (s *Stack) Pop() string {
	if s.IsEmpty() {
		return ""
	}
	res := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]
	return res
}

func (s *Stack) Top() string {
	if s.IsEmpty() {
		return ""
	}
	return s.s[len(s.s)-1]
}
