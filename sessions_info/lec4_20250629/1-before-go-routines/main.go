// world before go routines

// blog example:
// Publishing a blog post sequentially

package main

import (
	"fmt"
	util "lec4/util"
	"time"
)

func publishPost() {
	util.SavePostToDB()
	util.SendNotification()
	util.GenerateThumbnail()
	util.UpdateStats()
	fmt.Println("Post published!")
}

func main() {
	starttime := time.Now()
	publishPost()
	fmt.Println("Time taken:", time.Since(starttime))
}

// If sendNotification() is slow (e.g., sending emails), the whole process is blocked.
// Users have to wait even if the other tasks could be done independently.
