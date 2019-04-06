package wesense

import (
	"net/url"

	"github.com/dexidp/dex/pkg/log"
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

type WeSenseIdentity struct {
	ID      string `json:"_id"`
	Name    string `json:"name"`
	Surname string `json:"surename"`
	Email   string `json:"email"`
	Token   string `json:"token"`
}
