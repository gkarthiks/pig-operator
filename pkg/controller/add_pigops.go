package controller

import (
	"github.com/gkarthiks/pig-operator/pig-operator/pkg/controller/pigops"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, pigops.Add)
}
