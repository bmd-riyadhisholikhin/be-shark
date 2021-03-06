// Code generated by candi v1.6.4.

package main

import (
	"fmt"
	"runtime/debug"
	_ "github.com/lib/pq"

	"pkg.agungdp.dev/candi/codebase/app"
	"pkg.agungdp.dev/candi/config"

	service "be-shark/internal"
)

const serviceName = "be-shark"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("\x1b[31;1mFailed to start %s service: %v\x1b[0m\n", serviceName, r)
			fmt.Printf("Stack trace: \n%s\n", debug.Stack())
		}
	}()

	cfg := config.Init(serviceName)
	defer cfg.Exit()

	app.New(service.NewService(cfg)).Run()
}
