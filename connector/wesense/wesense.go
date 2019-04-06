package wesense

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/dexidp/dex/connector"
	"github.com/dexidp/dex/pkg/log"
)

var (
	_ connector.PasswordConnector = (*wesenseConnector)(nil)
	_ connector.RefreshConnector  = (*wesenseConnector)(nil)
)

// Open returns an authentication strategy which requires no user interaction.
func (c *Config) Open(id string, logger log.Logger) (connector.Connector, error) {

	u, err := url.Parse(c.URL)
	if err != nil {
		return nil, err
	}
	return &wesenseConnector{u, logger}, nil
}

// LoginURL returns the URL to redirect the user to login with.
func (w *wesenseConnector) Login(ctx context.Context, s connector.Scopes, username, password string) (identity connector.Identity, validPassword bool, err error) {
	requestURL := &url.URL{}
	*requestURL = *w.URL

	q := requestURL.Query()
	q.Set("username", username)
	q.Set("password", password)
	requestURL.RawQuery = q.Encode()

	w.logger.Debugf("Url: %s", requestURL.String())

	resp, err := http.Get(requestURL.String())
	if err != nil {
		return connector.Identity{}, false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return connector.Identity{}, false, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return connector.Identity{}, false, err
	}

	wIdentity, err := NewWeSenseIdentitiy(body)
	if err != nil {
		return connector.Identity{}, false, err
	}
	w.logger.Debugf("Identity: %v", wIdentity)

	return wIdentity.ConnectorIdentity(), true, nil
}

func (w *wesenseConnector) Refresh(_ context.Context, _ connector.Scopes, identity connector.Identity) (connector.Identity, error) {

	requestURL := &url.URL{}
	*requestURL = *w.URL

	q := requestURL.Query()
	q.Set("token", string(identity.ConnectorData))
	requestURL.RawQuery = q.Encode()

	w.logger.Debugf("Url: %s", requestURL.String())

	resp, err := http.Get(requestURL.String())
	if err != nil {
		return connector.Identity{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return connector.Identity{}, fmt.Errorf("User %s is unauthorized", identity.Username)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return connector.Identity{}, err
	}

	wIdentity, err := NewWeSenseIdentitiy(body)
	if err != nil {
		return connector.Identity{}, err
	}
	w.logger.Debugf("Identity: %v", wIdentity)

	return wIdentity.ConnectorIdentity(), nil
}

func (w *wesenseConnector) Prompt() string {
	return "Username"
}
