package version

import (
	"runtime"
)

// New function...
func Version() {
	println(runtime.Version())
}

func CPUValue() {
	println(runtime.NumCPU())
}
