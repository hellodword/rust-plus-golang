//go:build darwin || freebsd || linux
package main

// NOTE: There should be NO space between the comments and the `import "C"` line.
// The -ldl is sometimes necessary to fix linker errors about `dlsym`.

/*
#cgo LDFLAGS: ./lib/libhello.a -ldl
#include "./lib/hello.h"
#include <stdlib.h>
*/
import "C"

func main() {
	C.hello()
}
