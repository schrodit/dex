package wesense

import (
	"encoding/json"
	"fmt"

	"github.com/dexidp/dex/connector"
)

// NewWeSenseIdentitiy creates a new wesense identity out of a request
func NewWeSenseIdentitiy(data []byte) (*Identity, error) {
	fmt.Print(string(data))
	var status IStatus
	err := json.Unmarshal(data, &status)
	if err != nil {
		return nil, err
	}

	if status.Code != OkStatus {
		return nil, fmt.Errorf("Status with code %s: %s", status.Code, status.Data)
	}

	var wIdentity Identity
	err = json.Unmarshal(status.Data, &wIdentity)
	if err != nil {
		return nil, err
	}

	return &wIdentity, nil
}

// ConnectorIdentity converts a WeSenseIdentity into a connector Identity
func (id *Identity) ConnectorIdentity() connector.Identity {
	return connector.Identity{
		UserID:        id.ID,
		Username:      id.ID,
		Email:         id.Email,
		EmailVerified: true,
		Groups:        id.Groups,
		ConnectorData: []byte(id.Token),
	}
}
