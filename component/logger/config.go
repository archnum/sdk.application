/*
####### sdk.application (c) 2024 Archivage Numérique ############################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package logger

import "github.com/archnum/sdk.application/container"

type (
	Config struct {
		Level  string `dm:"level"`
		Output string `dm:"output"`
	}

	configProvider interface {
		ConfigLogger() *Config
	}
)

func config(c container.Container) *Config {
	return container.Value[configProvider](c, "config").ConfigLogger()
}

/*
####### END ############################################################################################################
*/
