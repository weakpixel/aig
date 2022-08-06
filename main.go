package main

import (
	"fmt"

	"github.com/weakpixel/aig/pkg/ansible"
	"github.com/weakpixel/aig/pkg/module"
)

func main() {

	s, _ := module.GetSpec()
	for _, m := range s.Modules {
		fmt.Println(m.ModuleName, "---", m.ShortDescription)
	}
	cmd := module.NewFile()
	cmd.Params.Path = "/tmp/myfile"
	cmd.Params.State = "absent"
	_, err := ansible.ExecuteLocal(cmd)
	fmt.Println(cmd.Result.Raw)
	if err != nil {
		panic(err)
	}
	cmd2 := module.NewFind()
	cmd2.Params.Recurse = true
	cmd2.Params.Paths = []string{
		"/tmp/",
	}
	_, err = ansible.ExecuteLocal(cmd2)
	if err != nil {
		fmt.Println(cmd2.Result.Raw)
		panic(err)
	}
	fmt.Println("Matched: ", cmd2.Result.Matched)
	fmt.Println("Skiped:", cmd2.Result.SkippedPaths)
	if cmd2.Result.Matched > 0 {
		fmt.Println(cmd2.Result.Files[0])
	}
	m := module.ModuleByName("setup")
	_, err = ansible.ExecuteLocal(m)
	if err != nil {
		panic(err)
	}
	fmt.Println("Failed: ", m.GetCommonResult().Failed, "Cause: ", m.GetCommonResult().Msg)
}
