// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package di

import (
	"github.com/int128/kubectl-external-forward/pkg/cmd"
	"github.com/int128/kubectl-external-forward/pkg/externalforwarder"
	"github.com/int128/kubectl-external-forward/pkg/portforwarder"
)

// Injectors from di.go:

func NewCmd() cmd.Interface {
	portForwarder := &portforwarder.PortForwarder{}
	externalForwarder := &externalforwarder.ExternalForwarder{
		PortForwarder: portForwarder,
	}
	cmdCmd := &cmd.Cmd{
		ExternalForwarder: externalForwarder,
	}
	return cmdCmd
}
