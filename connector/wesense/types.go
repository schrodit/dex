package wesense

import (
	"encoding/json"
	"net/url"

	"github.com/dexidp/dex/pkg/log"
)

// Status is the type definition for IStatus
type Status string

// IStatus types
const (
	OkStatus    Status = "OK"
	ErrorStatus Status = "ERROR"
)

// Config holds the configuration parameters for a connector which returns an
// identity with the HTTP header X-Remote-User as verified email.
type Config struct {
	URL string `json:"url"`
}

type wesenseConnector struct {
	URL *url.URL

	logger log.Logger
}

// IStatus represents the WeSense status
type IStatus struct {
	Code Status          `json:"code"`
	Data json.RawMessage `json:"data"`
}

// Identity represents the Identity that comes from the wesene account
type Identity struct {
	ID      string   `json:"_id"`
	Name    string   `json:"name"`
	Surname string   `json:"surename"`
	Email   string   `json:"email"`
	Token   string   `json:"token"`
	Groups  []string `json:"groups"`
}
