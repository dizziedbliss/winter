package bubbletea

import "fmt"
import tea "charm.land/bubbletea/v2"

type ErrMsg error
type StatusMsg string
type SuccessMsg string

type UI interface {
	Status(string)
	Progress(int)
	Error(error)
	Success(string)
}

type BubbleTeaUI struct {
	program *tea.Program
}

func (b *BubbleTeaUI) Status(msg string) {
	b.program.Send(StatusMsg(msg))
}

func (b *BubbleTeaUI) Progress(percent int) {
	fmt.Printf("Progress: %d%%\n", percent)
}

func (b *BubbleTeaUI) Error(err error) {
	b.program.Send(ErrMsg(err))
}

func (b *BubbleTeaUI) Success(msg string) {
	b.program.Send(SuccessMsg(msg))
}

func New() *BubbleTeaUI {

	model := initModel()

	return &BubbleTeaUI{
		program: tea.NewProgram(model),
	}
}

func (b *BubbleTeaUI) Run(task func() error) error {

	go func() {
		_ = task()
	}()

	_, err := b.program.Run()

	return err
}
