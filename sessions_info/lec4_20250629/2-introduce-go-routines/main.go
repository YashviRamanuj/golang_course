// We want to speed up publishing by doing tasks (notifications, thumbnail, stats) at the same time.

package main

import (
	"fmt"
	util "lec4/util"
	"time"
)

func publishPostV2() {
	go util.SendNotification()
	go util.GenerateThumbnail()
	go util.UpdateStats()
	util.SavePostToDB()
	fmt.Println("Post published!")
}

func main() {
	starttime := time.Now()
	publishPostV2()
	fmt.Println("Time taken:", time.Since(starttime))
}

// another analogy
// Imagine you’re a chef (main function).
// Instead of making toast, coffee, and eggs one by one, you shout instructions (“go make toast!”) to your helpers.
// Everyone works at once, breakfast is ready faster.

// TODO more theory in technical terms
