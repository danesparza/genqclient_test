package main

import (
	"context"
	"github.com/Khan/genqlient/graphql"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
)

type authedTransport struct {
	key     string
	wrapped http.RoundTripper
}

func (t *authedTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "bearer "+t.key)
	return t.wrapped.RoundTrip(req)
}

func main() {
	log.Info().Msg("Starting up...")

	key := os.Getenv("GITHUB_TOKEN")
	if key == "" {
		log.Error().Msg("must set GITHUB_TOKEN=<github token>")
		return
	}

	//	Create the http client with the token
	httpClient := http.Client{
		Transport: &authedTransport{
			key:     key,
			wrapped: http.DefaultTransport,
		},
	}

	//	Create the context and the graphQL client
	ctx := context.Background()
	graphqlClient := graphql.NewClient("https://api.github.com/graphql", &httpClient)

	//	Call the function:
	resp, err := getUser(ctx, graphqlClient, "danesparza")
	if err != nil {
		log.Err(err).Msg("problem calling getUser")
		return
	}

	log.Info().Str("Name", resp.User.Name).Str("bio", resp.User.Bio).Msg("Success!")
}

//go:generate go run github.com/Khan/genqlient genqlient.yaml
