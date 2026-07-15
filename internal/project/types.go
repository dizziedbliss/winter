package project

import (
	"log"

	"github.com/dizziedbliss/winter/internal/ui/bubbletea"
)

type Project struct {
	Path       string
	Verbose    bool
	Config     string
	ConfigPath string
	Logger     *log.Logger

	DeployUI ui.UI
}
