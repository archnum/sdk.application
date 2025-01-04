/*
####### sdk.application (c) 2024 Archivage Numérique ############################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package waitend

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/archnum/sdk.application/container"
)

const (
	Name = "waitend"
)

type (
	implComponent struct {
		*container.Component
	}
)

func New(c container.Container) *implComponent {
	return &implComponent{
		Component: container.NewComponent(Name, c),
	}
}

//////////////////////
/// Implementation ///
//////////////////////

func (impl *implComponent) Start() error {
	signals := []os.Signal{syscall.SIGABRT, syscall.SIGQUIT, syscall.SIGTERM}

	if impl.C().App().Environment() != "" {
		signals = append(signals, os.Interrupt)
	}

	ctx, stop := signal.NotifyContext(context.Background(), signals...)

	defer stop()
	<-ctx.Done()

	return nil
}

/*
####### END ############################################################################################################
*/
