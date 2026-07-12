package deployment

import "log"

type DeployOptions struct {
	Path string;
	Verbose bool;
	Config string;
	ConfigPath string;
	Logger *log.Logger;
}