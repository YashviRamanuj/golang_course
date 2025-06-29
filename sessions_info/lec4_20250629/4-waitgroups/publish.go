package main

import (
	"fmt"
	"sync"
	"time"
)

// Simulate thumbnail generation
func generateThumbnail(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Generating thumbnail...")
	time.Sleep(time.Second)
	fmt.Println("Thumbnail generated.")
}

// Simulate sending notifications to multiple subscribers
func sendNotifications(subscribers []string, wg *sync.WaitGroup) {
	defer wg.Done()

	var notifWG sync.WaitGroup
	notifWG.Add(len(subscribers))

	for _, user := range subscribers {
		go func(user string) {
			defer notifWG.Done()
			fmt.Printf("Sending notification to %s...\n", user)
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("Notification sent to %s\n", user)
		}(user)
	}

	notifWG.Wait()
	fmt.Println("All notifications sent.")
}

// Simulate analytics update
func updateAnalytics(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Updating analytics...")
	time.Sleep(700 * time.Millisecond)
	fmt.Println("Analytics updated.")
}

func afterPublish() {
	var wg sync.WaitGroup

	subscribers := []string{"alice@site.com", "bob@site.com", "carol@site.com"}

	wg.Add(3)
	go generateThumbnail(&wg)
	go sendNotifications(subscribers, &wg)
	go updateAnalytics(&wg)

	wg.Wait()
	fmt.Println("Blog post published: all tasks done!")
}
