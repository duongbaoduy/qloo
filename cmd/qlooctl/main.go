package main

import (
	"fmt"
	"os"

	"github.com/solo-io/qloo/pkg/qlooctl"
)

func main() {
	if err := qlooctl.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}