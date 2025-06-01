package utils

import "fmt"

func HandleErr(err error) {
	if err != nil {
		// TODO: Alert Discord instead of panicing.
		panic(fmt.Errorf("%s", err))
	}
}
