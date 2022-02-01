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
		fmt.Println(" %{F#FC4D44}☠%{F-} ")
	} else if len(notifs) == 0 {
		fmt.Println(" %{F#4E9A06}%{F-} ")
	} else {
		fmt.Println(" %{F#FC4D44}%{F-} ")
	}
}
