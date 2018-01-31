package graylog

import "context"

const usersBasePath = "users"

// UserService is an interface for interfacing with the Users endpoints of the Graylog API
type UserService interface {
	// List(context.Context, *ListOptions) ([]User, *Response, error)
	// Get(context.Context, string) (*User, *Response, error)
	// Create(context.Context, *UserCreateRequest) (*User, *Response, error)
	// Delete(context.Context, string) (*Response, error)
}

// UserServiceOp handles communication with the User related methods of the API.
type UserServiceOp struct {
	client *Client
}

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

// List all Users.
func (s *UserServiceOp) List(ctx context.Context, opt *ListOptions) ([]User, *Response, error) {
	// path := usersBasePath
	// path, err := addOptions(path, opt)
	// if err != nil {
	// 	return nil, nil, err
	// }

	// return s.list(ctx, path)
	return []User{}, nil, nil
}
