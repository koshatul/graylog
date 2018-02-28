package graylog

import (
	"context"
	"fmt"
	"net/http"
)

const usersBasePath = "users"

// UserService is an interface for interfacing with the Users endpoints of the Graylog API
type UserService interface {
	List(context.Context) ([]User, *Response, error)
	Get(context.Context, string) (User, *Response, error)
}

// UserServiceOp handles communication with the User related methods of the API.
type UserServiceOp struct {
	client *Client
}

// User object returned from API
type User struct {
	ID          string   `json:"id"`
	Username    string   `json:"username"`
	Email       string   `json:"email"`
	FullName    string   `json:"full_name"`
	Permissions []string `json:"permissions"`
	Preferences struct {
		UpdateUnfocussed  bool `json:"updateUnfocussed"`
		EnableSmartSearch bool `json:"enableSmartSearch"`
	} `json:"preferences"`
	Timezone         string      `json:"timezone"`
	SessionTimeoutMs int         `json:"session_timeout_ms"`
	External         bool        `json:"external"`
	Startpage        interface{} `json:"startpage"`
	Roles            []string    `json:"roles"`
	ReadOnly         bool        `json:"read_only"`
	SessionActive    bool        `json:"session_active"`
	LastActivity     string      `json:"last_activity"`
	ClientAddress    string      `json:"client_address"`
}

type userListRoot struct {
	Users []User `json:"users"`
}

// Performs a list request given a path.
func (s *UserServiceOp) list(ctx context.Context, path string) ([]User, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(userListRoot)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root.Users, resp, err
}

// Performs a list request given a path.
func (s *UserServiceOp) get(ctx context.Context, path string) (User, *Response, error) {
	user := new(User)
	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return *user, nil, err
	}

	resp, err := s.client.Do(ctx, req, user)
	if err != nil {
		return *user, resp, err
	}

	return *user, resp, err
}

// List all users.
func (s *UserServiceOp) List(ctx context.Context) ([]User, *Response, error) {
	path := usersBasePath

	return s.list(ctx, path)
}

// Get user.
func (s *UserServiceOp) Get(ctx context.Context, user string) (User, *Response, error) {
	path := fmt.Sprintf("%s/%s", usersBasePath, user)

	return s.get(ctx, path)
}
