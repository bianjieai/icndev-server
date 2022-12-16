package main

import (
	"github.com/bianjieai/icndev-server/cmd"
	"runtime"
)

// @title SupTechServer Swagger API
// @version visit /version
// @description sup-tech-server api document
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() * 8)
	cmd.Execute()
}
