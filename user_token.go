package graylog

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"time"
)

// UserTokenService is an interface for interfacing with the User Tokens endpoints of the Graylog API
type UserTokenService interface {
	List(context.Context, string) ([]Token, *Response, error)
	Create(context.Context, string, string) (Token, *Response, error)
	Delete(context.Context, string, Token) (*Response, error)
	// Get(context.Context, string) (*User, *Response, error)
}

// UserTokenServiceOp handles communication with the User Token related methods of the API.
type UserTokenServiceOp struct {
	client *Client
}

// Token returned by API
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

	return root.Tokens, resp, err
}

// Performs a create request given a path.
func (s *UserTokenServiceOp) create(ctx context.Context, path string) (Token, *Response, error) {
	root := new(Token)
	req, err := s.client.NewRequest(ctx, http.MethodPost, path, "{}")
	if err != nil {
		return *root, nil, err
	}

	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return *root, resp, err
	}

	return *root, resp, err
}

// Performs a delete request given a path.
func (s *UserTokenServiceOp) delete(ctx context.Context, path string) (*Response, error) {
	buf := &bytes.Buffer{}
	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, buf)
	if err != nil {
		return resp, err
	}

	return resp, err
}

// List all user tokens.
func (s *UserTokenServiceOp) List(ctx context.Context, user string) ([]Token, *Response, error) {
	path := fmt.Sprintf("%s/%s/tokens", usersBasePath, user)

	return s.list(ctx, path)
}

// Create a new user token.
func (s *UserTokenServiceOp) Create(ctx context.Context, user, tokenName string) (Token, *Response, error) {
	path := fmt.Sprintf("%s/%s/tokens/%s", usersBasePath, user, tokenName)

	return s.create(ctx, path)
}

// Delete a user token.
func (s *UserTokenServiceOp) Delete(ctx context.Context, user string, token Token) (*Response, error) {
	path := fmt.Sprintf("%s/%s/tokens/%s", usersBasePath, user, token.Token)

	return s.delete(ctx, path)
}
