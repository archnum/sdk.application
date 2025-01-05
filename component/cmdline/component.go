/*
####### sdk.application (c) 2024 Archivage Numérique ############################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package cmdline

import (
	"os"

	"github.com/archnum/sdk.base/cmdline"

	"github.com/archnum/sdk.application/container"
)

type (
	Builder func() (*cmdline.CmdLine, error)

	implComponent struct {
		*container.Component
		builder Builder
	}
)

func New(c container.Container, builder Builder) *implComponent {
	return &implComponent{
		Component: container.NewComponent("cmdline", c),
		builder:   builder,
	}
}

/////////////////
/// Component ///
/////////////////

func (impl *implComponent) Build() error {
	cmdline, err := impl.builder()
	if err != nil {
		return err
	}

	return cmdline.Run(os.Args[1:])
}

/*
####### END ############################################################################################################
*/
