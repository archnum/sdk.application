/*
####### sdk.application (c) 2024 Archivage Numérique ############################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package logger

import (
	"os"
	"time"

	"github.com/archnum/sdk.base/failure"
	"github.com/archnum/sdk.base/kv"
	"github.com/archnum/sdk.base/logger"
	"github.com/archnum/sdk.base/logger/handler"
	"github.com/archnum/sdk.base/logger/level"

	"github.com/archnum/sdk.application/container"
)

type (
	implComponent struct {
		*container.Component
		logger *logger.Logger
	}
)

func New(c container.Container) *implComponent {
	return &implComponent{
		Component: container.NewComponent("logger", c),
	}
}

//////////////////////
/// Implementation ///
//////////////////////

func (impl *implComponent) Build() error {
	c := impl.C()
	app := c.App()
	env := app.Environment()
	cfg := config(c)

	impl.logger = logger.New(app.ID(), app.Name())
	impl.logger.Register()

	if cfg.Level != "" {
		impl.logger.SetLevel(level.StringToLevel(cfg.Level))
	} else if env == "" {
		impl.logger.SetLevel(level.Info)
	} else {
		impl.logger.SetLevel(level.Trace)
	}

	switch cfg.Output {
	case "":
		if env == "" {
			impl.logger.AddHandler(handler.Journal)
		} else {
			impl.logger.AddHandler(handler.Console)
		}
	case "console":
		impl.logger.AddHandler(handler.Console)
	default:
		return failure.New("unknown logger output", kv.String("name", cfg.Output)) /////////////////////////////////////
	}

	impl.logger.Info( //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
		"APP",
		kv.String("id", string(app.ID())),
		kv.String("name", app.Name()),
		kv.String("version", app.Version()),
		kv.Time("builtAt", app.BuiltAt()),
		kv.Int("pid", os.Getpid()),
	)

	impl.SetValue(impl.logger)

	return nil
}

func (impl *implComponent) Close() error {
	impl.logger.Info( //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
		"END",
		kv.String("uptime", time.Since(impl.C().App().StartedAt()).Round(time.Second).String()),
	)

	return nil
}

/*
####### END ############################################################################################################
*/
