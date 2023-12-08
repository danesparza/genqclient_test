// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package main

import (
	"context"

	"github.com/Khan/genqlient/graphql"
)

// __getUserInput is used internally by genqlient
type __getUserInput struct {
	Login string `json:"login"`
}

// GetLogin returns __getUserInput.Login, and is useful for accessing the field via an interface.
func (v *__getUserInput) GetLogin() string { return v.Login }

// getUserResponse is returned by getUser on success.
type getUserResponse struct {
	// Lookup a user by login.
	User getUserUser `json:"user"`
}

// GetUser returns getUserResponse.User, and is useful for accessing the field via an interface.
func (v *getUserResponse) GetUser() getUserUser { return v.User }

// getUserUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A user is an individual's account on GitHub that owns repositories and can make new content.
type getUserUser struct {
	// The user's public profile name.
	Name string `json:"name"`
	// The user's public profile bio.
	Bio string `json:"bio"`
}

// GetName returns getUserUser.Name, and is useful for accessing the field via an interface.
func (v *getUserUser) GetName() string { return v.Name }

// GetBio returns getUserUser.Bio, and is useful for accessing the field via an interface.
func (v *getUserUser) GetBio() string { return v.Bio }

// The query or mutation executed by getUser.
const getUser_Operation = `
query getUser ($login: String!) {
	user(login: $login) {
		name
		bio
	}
}
`

func getUser(
	ctx context.Context,
	client graphql.Client,
	login string,
) (*getUserResponse, error) {
	req := &graphql.Request{
		OpName: "getUser",
		Query:  getUser_Operation,
		Variables: &__getUserInput{
			Login: login,
		},
	}
	var err error

	var data getUserResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}
