/*
####### sdk.application (c) 2024 Archivage Numérique ############################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package crypto

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/archnum/sdk.base/crypto"
	"github.com/archnum/sdk.base/failure"
	"github.com/archnum/sdk.base/util"

	"github.com/archnum/sdk.application/container"
)

const (
	_name       = "crypto"
	_rootFolder = "/var/lib"
)

type (
	implComponent struct {
		*container.Component
	}
)

func New(c container.Container) *implComponent {
	return &implComponent{
		Component: container.NewComponent(_name, c),
	}
}

func (impl *implComponent) readFile() (string, error) {
	app := impl.C().App()
	ecosystem := app.Ecosystem()
	name := util.If(ecosystem == "", app.Name(), ecosystem)
	filename := filepath.Join(_rootFolder, fmt.Sprintf(".%s", name))

	if !util.FileExist(filename) {
		return "", failure.NotFound
	}

	key, err := os.ReadFile(filename)
	if err != nil {
		return "",
			failure.WithMessage(err, "failed to retrieve encryption key") //////////////////////////////////////////////
	}

	return string(key), nil
}

func Value(c container.Container) crypto.Crypto {
	return container.Value[crypto.Crypto](c, _name)
}

//////////////////////
/// Implementation ///
//////////////////////

func (impl *implComponent) Build() error {
	crypto := crypto.New()

	key, err := impl.readFile()
	if err != nil {
		if err != failure.NotFound {
			return err
		}

		if impl.C().App().Environment() == "" {
			return failure.New("in production environment the encryption key must be configured") //////////////////////
		}
	} else if err := crypto.SetKey(key); err != nil {
		return err
	}

	impl.SetValue(crypto)

	return nil
}

/*
####### END ############################################################################################################
*/
