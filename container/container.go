/*
####### sdk.application (c) 2024 Archivage Numérique ############################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package container

import (
	"log"

	"github.com/archnum/sdk.application/component"
	"github.com/archnum/sdk.base/application"
	"github.com/archnum/sdk.base/failure"
	"github.com/archnum/sdk.base/kv"
	"github.com/archnum/sdk.base/tracer"
	"github.com/archnum/sdk.base/util"
)

type (
	Container interface {
		App() *application.Application
		ComponentValue(name string) (any, error)
	}

	implContainer struct {
		app *application.Application
		cs  []component.Component
		msc map[string]component.Component
	}
)

func New(app *application.Application) *implContainer {
	return &implContainer{
		app: app,
		msc: make(map[string]component.Component),
	}
}

func (impl *implContainer) App() *application.Application {
	return impl.app
}

func (impl *implContainer) ComponentValue(name string) (any, error) {
	if cpt, ok := impl.msc[name]; ok {
		return cpt.Value(), nil
	}

	return nil,
		failure.New("this component doesn't exist", kv.String("name", name)) ///////////////////////////////////////////
}

func (impl *implContainer) AddComponents(cs ...component.Component) {
	impl.cs = cs
}

func safeRun(fn func() error) (err error) {
	defer func() {
		if data := recover(); data != nil {
			err = failure.New( /////////////////////////////////////////////////////////////////////////////////////////
				"Recovered error",
				kv.Any("data", data),
				kv.String("stack", util.Stack(5)),
			)
		}
	}()

	err = fn()
	return
}

func (impl *implContainer) Run(cs ...string) error {
	var buildClose int

	defer func() {
		for i := buildClose - 1; i >= 0; i-- {
			cpt := impl.cs[i]

			tracer.Log("Component", kv.String(cpt.Name(), "Close")) //::::::::::::::::::::::::::::::::::::::::::::::::::

			if err := safeRun(cpt.Close); err != nil {
				log.Print( //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
					failure.WithMessage(
						err,
						"an error occurred",
						kv.String("component", cpt.Name()),
						kv.String("func", "Close"),
					).Error(),
				)
			}
		}
	}()

	for _, cpt := range impl.cs {
		tracer.Log("Component", kv.String(cpt.Name(), "Build")) //::::::::::::::::::::::::::::::::::::::::::::::::::::::

		if err := safeRun(cpt.Build); err != nil {
			if err == failure.NoError {
				return nil
			}

			return failure.WithMessage( ////////////////////////////////////////////////////////////////////////////////
				err,
				"an error occurred",
				kv.String("component", cpt.Name()),
				kv.String("func", "Build"),
			)
		}

		buildClose++

		impl.msc[cpt.Name()] = cpt
	}

	var startStop int

	defer func() {
		for i := startStop - 1; i >= 0; i-- {
			cpt := impl.msc[cs[i]]

			tracer.Log("Component", kv.String(cpt.Name(), "Stop")) //:::::::::::::::::::::::::::::::::::::::::::::::::::

			if err := safeRun(cpt.Stop); err != nil {
				log.Print( //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
					failure.WithMessage(
						err,
						"an error occurred",
						kv.String("component", cpt.Name()),
						kv.String("func", "Stop"),
					).Error(),
				)
			}
		}
	}()

	for _, name := range cs {
		cpt, ok := impl.msc[name]
		if !ok {
			return failure.New("this component doesn't exist", kv.String("name", name)) ////////////////////////////////
		}

		tracer.Log("Component", kv.String(cpt.Name(), "Start")) //::::::::::::::::::::::::::::::::::::::::::::::::::::::

		if err := safeRun(cpt.Start); err != nil {
			if err == failure.NoError {
				return nil
			}

			return failure.WithMessage( ////////////////////////////////////////////////////////////////////////////////
				err,
				"an error occurred",
				kv.String("component", cpt.Name()),
				kv.String("func", "Start"),
			)
		}

		startStop++
	}

	return nil
}

/*
####### END ############################################################################################################
*/
