package main

import (
	"context"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type GithubClient struct {
	client *github.Client
}

func NewGithubClientWithToken(token string) *GithubClient {
	if token == "" {
		return &GithubClient{github.NewClient(nil)}
	}

	return &GithubClient{
		github.NewClient(
			oauth2.NewClient(
				context.Background(),
				oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token}),
			),
		),
	}
}

type Notification struct {
	ID  string `json:"id"`
	URL string `json:"urls"`
}

func (c *GithubClient) getNotifications(ctx context.Context) ([]*Notification, error) {
	req, err := c.client.NewRequest("GET", "notifications", 0)
	if err != nil {
		return nil, err
	}

	var notifs []*Notification
	_, err = c.client.Do(ctx, req, &notifs)
	if err != nil {
		return nil, err
	}

	return notifs, err
}
