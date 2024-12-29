/*
####### sdk.application (c) 2024 Archivage Numérique ############################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package container

import (
	"github.com/archnum/sdk.base/failure"
	"github.com/archnum/sdk.base/kv"
)

type (
	Component struct {
		c     Container
		name  string
		value any
	}
)

func NewComponent(name string, c Container) *Component {
	return &Component{
		c:    c,
		name: name,
	}
}

func (cpt *Component) SetValue(value any) {
	cpt.value = value
}

/////////////////
/// Interface ///
/////////////////

func (cpt *Component) C() Container {
	return cpt.c
}

func (cpt *Component) Name() string {
	return cpt.name
}

func (cpt *Component) Value() any {
	if cpt.value == nil {
		panic(failure.New("zero value component", kv.String("name", cpt.Name()))) //!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	}

	return cpt.value
}

func (cpt *Component) Build() error { return nil }
func (cpt *Component) Start() error { return nil }
func (cpt *Component) Stop() error  { return nil }
func (cpt *Component) Close() error { return nil }

/*
####### END ############################################################################################################
*/
