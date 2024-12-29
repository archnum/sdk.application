/*
####### sdk.application (c) 2024 Archivage Numérique ############################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package container

import (
	"fmt"

	"github.com/archnum/sdk.base/crypto"
	"github.com/archnum/sdk.base/failure"
	"github.com/archnum/sdk.base/kv"
	"github.com/archnum/sdk.base/logger"
)

func Value[T any](c Container, name string) T {
	tmp, err := c.ComponentValue(name)
	if err != nil {
		panic(err) //!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	}

	value, ok := tmp.(T)
	if !ok {
		var unused T

		err := failure.New( ////////////////////////////////////////////////////////////////////////////////////////////
			"unexpected value type for this component",
			kv.String("name", name),
			kv.String("got", fmt.Sprintf("%T", tmp)),
			kv.String("want", fmt.Sprintf("%T", unused)),
		)

		panic(err) //!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	}

	return value
}

func Crypto(c Container) crypto.Crypto {
	return Value[crypto.Crypto](c, "crypto")
}

func Logger(c Container) *logger.Logger {
	return Value[*logger.Logger](c, "logger")
}

/*
####### END ############################################################################################################
*/
