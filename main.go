package main

import (
	"aig/pkg/module"
	"fmt"
)

func main() {

	// cmd := module.NewFind()
	// cmd.Options.Paths = []string{
	// 	"/tmp",
	// }
	// cmd.Options.Recurse = true
	cmd := module.NewFile()
	cmd.Options.Path = "/tmp/myfile"
	cmd.Options.State = "absent"
	err := cmd.Run()
	if err != nil {
		fmt.Println(cmd.Result.Raw)

		panic(err)
	}
	// for _, f := range cmd.Result.Files {
	// 	fmt.Println(f["path"])
	// }
}
