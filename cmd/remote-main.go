package main

import (
	"fmt"
	"log"

	"github.com/k0sproject/rig"
	"github.com/weakpixel/aig/pkg/ansible"
	"github.com/weakpixel/aig/pkg/module"
)

func main() {

	r, err := ansible.NewRemote(&rig.SSH{
		Address: "orko.run",
		// User:    "root",
		// KeyPath: "/Users/dkuffner/.ssh/orko-run",
		// PasswordCallback: func() (string, error) {
		// 	fmt.Println("Enter password:")
		// 	pass, err := terminal.ReadPassword(int(syscall.Stdin))
		// 	return string(pass), err
		// },
	})

	if err != nil {
		log.Fatal(err)
	}

	// c := module.NewCommand()
	// c.Params.Cmd = "find"
	// c.Params.Argv = []string{"/tmp"}
	// f := module.NewFind()
	// f.Params.Paths = []string{"/var/lib/"}
	// f.Params.
	// c := module.NewApt()
	// c.Params.Name = []string{"jq"}
	// c.Params.State = "present"

	// // c.Params.Service
	// res, err := ansible.ExecuteRemote2(r, c)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(res)
	// fmt.Println("--------")

	// c = module.NewApt()
	// c.Params.Name = []string{"jq"}
	// c.Params.State = "absent"

	// // c.Params.Service
	// res, err = ansible.ExecuteRemote2(r, c)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(res)
	// fmt.Println("--------")
	g := module.NewGit()
	g.Params.Clone = true
	g.Params.Dest = "/tmp/clone"
	g.Params.Repo = "https://github.com/k0sproject/rig.git"

	res, err := ansible.ExecuteRemote(r, g)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
	fmt.Println("--------")

	s := module.NewSetup()

	res, err = ansible.ExecuteRemote(r, s)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
	fmt.Println("--------")

	// fmt.Println(c.Result.)
	// fmt.Println(c.Result.Msg)

	// for _, f := range f.Result.Files {
	// 	fmt.Println(f)
	// }

}
