package main

import (
	"fmt"
	"testing"
)

type MockReadWriter struct{}

func (cli MockReadWriter) Read() string {
	return "testing read"
}

func (cli MockReadWriter) Write(output string) {
	fmt.Println(output)
}

func TestStudent(t *testing.T) {
	Student(MockReadWriter{})
}