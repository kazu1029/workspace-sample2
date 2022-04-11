package main

import (
	"github.com/kazu1029/workspace-sample1/modules/module1"
	"github.com/kazu1029/workspace-sample1/modules/module2"
)

func main() {
	module2Service := module2.NewModule2Service()
	module2Service.Print("from workspace2")

	module1Service := module1.NewModule1Service()
	module1Service.Print("from workspace2")
}
