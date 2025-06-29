package util

import (
	"fmt"
	"time"
)

func SavePostToDB() {
	time.Sleep(2 * time.Second)
	fmt.Println("Post saved to DB")
}
func SendNotification() {
	time.Sleep(2 * time.Second)
	fmt.Println("Notification sent!")
}
func GenerateThumbnail() {
	time.Sleep(2 * time.Second)
	fmt.Println("Thumbnail generated")
}
func UpdateStats() {
	time.Sleep(2 * time.Second)
	fmt.Println("Stats updated")
}
