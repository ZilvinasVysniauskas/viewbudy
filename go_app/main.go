package main

/*
#cgo CFLAGS: -x objective-c -fmodules -fobjc-arc
#cgo LDFLAGS: -L${SRCDIR}/bin -lplayer
#include <stdlib.h>

// Add the typedef here so Go knows about it
typedef void (*CLogCallback)(const char*);

#include "libplayer.h"
*/
import "C"
import (
	"fmt"
	"time"
)

//export StartGoLogic
func StartGoLogic() {
	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("Simulating incoming request: pause")
		C.pauseVideo()
	}()
}

func main() {}
