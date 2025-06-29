// problem: unfinished go routines
package main

import (
	"fmt"
	util "lec4/util"
)

func notify() {
	go util.SendNotification()
	fmt.Println("Done!") // main exits, program ends, sendNotification might NOT HAVE BEEN finished
}

func main() {
	notify()
}
