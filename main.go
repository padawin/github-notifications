package main

import (
	"context"
	"fmt"
	"os"
)

func main() {
	c := NewGithubClientWithToken(os.Getenv("GITHUB_NOTIFICATIONS_TOKEN"))
	notifs, err := c.getNotifications(context.Background())
	if err != nil {
		fmt.Print("%{B#FC4D44}%{F#2E3436} ☠ %{B-}")
	} else if len(notifs) > 0 {
		fmt.Printf("%%{B#FC4D44}%%{F#2E3436}  %d %%{F-}%%{B-}", len(notifs))
	}
}
