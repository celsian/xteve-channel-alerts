package utils

import (
	"fmt"
)

func PanicOnErr(err error) {
	if err != nil {
		e := fmt.Errorf("%v", err)
		// TODO: Alert Discord instead of panicing.
		panic(e)
	}
}
