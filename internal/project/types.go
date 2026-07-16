package project

import (
	"log"

	ui "github.com/dizziedbliss/winter/internal/ui/bubbletea"
)

type Project struct {
	Path       string
	Verbose    bool
	Config     string
	ConfigPath string
	Logger     *log.Logger

	UI ui.UI
}

func NewProject(path string, verbose bool, ui ui.UI) *Project {
	return &Project{
		Path:    path,
		Verbose: verbose,
		Logger:  log.Default(),
		UI:      ui,
	}
}
