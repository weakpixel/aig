package main

import (
	"fmt"

	"github.com/weakpixel/aig/pkg/module"
)

func main() {

	s, _ := module.GetSpec()
	for _, m := range s.Modules {
		fmt.Println(m.ModuleName, "---", m.ShortDescription)
	}
	// cmd := module.NewFind()
	// cmd.Params.Paths = []string{
	// 	"/tmp",
	// }
	// cmd.Params.Recurse = true
	cmd := module.NewFile()
	cmd.Params.Path = "/tmp/myfile"
	cmd.Params.State = "absent"
	err := cmd.Run()
	if err != nil {
		fmt.Println(cmd.Result.Raw)
		panic(err)
	}
	// for _, f := range cmd.Result.Files {
	// 	fmt.Println(f["path"])
	// }
}
