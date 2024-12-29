/*
####### sdk.application (c) 2024 Archivage Numérique ############################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package config

import "github.com/archnum/sdk.application/container"

type (
	implComponent struct {
		*container.Component
	}
)

func New(c container.Container, value any) *implComponent {
	impl := &implComponent{
		Component: container.NewComponent("config", c),
	}

	impl.SetValue(value)

	return impl
}

/*
####### END ############################################################################################################
*/
