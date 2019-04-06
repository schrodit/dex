package wesense

import (
	"encoding/json"

	"github.com/dexidp/dex/connector"
)

// NewWeSenseIdentitiy creates a new wesense identity out of a request
func NewWeSenseIdentitiy(data []byte) (*WeSenseIdentity, error) {
	var wIdentity WeSenseIdentity
	err := json.Unmarshal(data, &wIdentity)
	if err != nil {
		return nil, err
	}

	return &wIdentity, nil
}

// ConnectorIdentity converts a WeSenseIdentity into a connector Identity
func (id *WeSenseIdentity) ConnectorIdentity() connector.Identity {
	return connector.Identity{
		UserID:        id.ID,
		Username:      id.ID,
		Email:         id.Email,
		EmailVerified: true,
		ConnectorData: []byte(id.Token),
	}
}
