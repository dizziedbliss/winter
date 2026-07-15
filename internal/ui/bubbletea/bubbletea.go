package ui

import "fmt"

type UI interface {
	Status(string)
	Progress(int)
	Error(error)
	Success(string)
}

type BubbleTeaUI struct{}

func (b *BubbleTeaUI) Status(msg string) {
	fmt.Println(msg)
}

func (b *BubbleTeaUI) Progress(percent int) {
	fmt.Printf("Progress: %d%%\n", percent)
}

func (b *BubbleTeaUI) Error(err error) {
	fmt.Println("Error:", err)
}

func (b *BubbleTeaUI) Success(msg string) {
	fmt.Println("Success:", msg)
}
