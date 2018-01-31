package graylog

import (
	"context"
	"net/http"
	"time"
)

type UserTokenService interface {
	List(context.Context, *ListOptions) ([]Token, *Response, error)
	// Create(context.Context, *UserCreateTokenRequest) (*Response, error)
	// Get(context.Context, string) (*User, *Response, error)
	// Delete(context.Context, string) (*Response, error)
}

// UserTokenServiceOp handles communication with the User Token related methods of the API.
type UserTokenServiceOp struct {
	client *Client
}

type Token struct {
	Name       string    `json:"name"`
	Token      string    `json:"token"`
	LastAccess time.Time `json:"last_access"`
}

type tokensRoot struct {
	Tokens []Token `json:"tokens"`
}

// Performs a list request given a path.
func (s *UserTokenServiceOp) list(ctx context.Context, path string) ([]Token, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(tokensRoot)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}
	// if l := root.Links; l != nil {
	// 	resp.Links = l
	// }

	return root.Tokens, resp, err
}

// List all user tokens.
func (s *UserTokenServiceOp) List(ctx context.Context, opt *ListOptions) ([]Token, *Response, error) {
	path := usersBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}

	return s.list(ctx, path)
}
