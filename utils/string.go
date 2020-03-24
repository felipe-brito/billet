package utils

import (
	"strconv"
)

type StringBuilder struct {
	newString string
}

type StringBuilderInterface interface {
	Append(string) StringBuilderInterface
	AppendInt(int) StringBuilderInterface
	Insert(int, string) StringBuilderInterface
	InsertInt(int, int) StringBuilderInterface
	Builder() string
}

func NewStringBuilder() StringBuilderInterface {
	return &StringBuilder{}
}

// Append
func (s *StringBuilder) Append(value string) StringBuilderInterface {
	s.newString += value
	return s
}

// AppendInt
func (s *StringBuilder) AppendInt(value int) StringBuilderInterface {
	s.newString += strconv.Itoa(value)
	return s
}

//Insert
func (s *StringBuilder) Insert(index int, value string) StringBuilderInterface {
	s.newString = s.newString[0:index] + value + s.newString[index:]
	return s
}

//Insert
func (s *StringBuilder) InsertInt(index int, value int) StringBuilderInterface {
	s.newString = s.newString[0:index] + strconv.Itoa(value) + s.newString[index:]
	return s
}

// Builder
func (s *StringBuilder) Builder() string {
	return s.newString
}