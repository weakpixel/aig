package main

import (
	"aig/pkg/module"
	"fmt"
)

func main() {

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
