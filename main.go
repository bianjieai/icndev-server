package main

import (
	"github.com/bianjieai/icndev-server/cmd"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() * 8)
	cmd.Execute()
}
