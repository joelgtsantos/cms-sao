//go:generate go run gen/main.go

package main

import (
	_ "github.com/jossemargt/cms-sao/design"

	"github.com/goadesign/goa/design"
	"github.com/goadesign/goa/goagen/codegen"
	"github.com/goadesign/goa/goagen/gen_app"
	"github.com/goadesign/goa/goagen/gen_client"
	"github.com/goadesign/goa/goagen/gen_controller"
	"github.com/goadesign/goa/goagen/gen_main"
	"github.com/goadesign/goa/goagen/gen_swagger"
)

func main() {
	codegen.ParseDSL()
	codegen.Run(
		genmain.NewGenerator(
			genmain.API(design.Design),
		),
		genapp.NewGenerator(
			genapp.API(design.Design),
			genapp.OutDir("app"),
			genapp.Target("app"),
			genapp.NoTest(false),
		),
		gencontroller.NewGenerator(
			gencontroller.API(design.Design),
		),
		genclient.NewGenerator(
			genclient.API(design.Design),
			genclient.NoTool(true),
		),
		genswagger.NewGenerator(
			genswagger.API(design.Design),
		),
	)
}
