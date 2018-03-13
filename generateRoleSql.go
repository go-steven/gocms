package main

import (
	"github.com/go-steven/cms/src/controllers"
	controller_base "github.com/go-steven/cms/src/controllers/base"

	"fmt"
	"reflect"
)

func main1() {
	baseController := new(controller_base.BaseController)
	typ := reflect.TypeOf(baseController)
	baseMethod := make(map[string]bool, typ.NumMethod())
	for i := 0; i < typ.NumMethod(); i++ {
		baseMethod[typ.Method(i).Name] = true
	}

	controller := new(controllers.MainController)
	controllerType := reflect.TypeOf(controller)

	for i := 0; i < controllerType.NumMethod(); i++ {
		methodName := controllerType.Method(i).Name
		if _, ok := baseMethod[methodName]; !ok {
			fmt.Println(methodName)
		}
	}

}
