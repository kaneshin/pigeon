package credentials

import (
	"fmt"
	"net/url"
)

// A StaticProvider is a set of credentials which are set pragmatically,
// and will never expire.
type StaticProvider struct {
	Value
}

// NewStaticCredentials returns a pointer to a new Credentials object
// wrapping a static credentials value provider.
func NewStaticCredentials(projectID, privateKeyID, privateKey, clientEmail, clientID string) *Credentials {
	u, err := url.ParseQuery(clientEmail)
	if err != nil {
		return nil
	}
	return NewCredentials(&StaticProvider{Value: Value{
		Type:                    "service_account",
		ProjectID:               projectID,
		PrivateKeyID:            privateKeyID,
		PrivateKey:              privateKey,
		ClientEmail:             clientEmail,
		ClientID:                clientID,
		AuthURI:                 "https://accounts.google.com/o/oauth2/auth",
		TokenURI:                "https://accounts.google.com/o/oauth2/token",
		AuthProviderX509CertURL: "https://www.googleapis.com/oauth2/v1/certs",
		ClientX509CertURL:       fmt.Sprintf("https://www.googleapis.com/robot/v1/metadata/x509/%s", u.Encode()),
	}})

}

// Retrieve returns the credentials or error if the credentials are invalid.
func (s *StaticProvider) Retrieve() (Value, error) {
	if s.ProjectID == "" ||
		s.PrivateKeyID == "" ||
		s.PrivateKey == "" ||
		s.ClientEmail == "" ||
		s.ClientID == "" {
		return Value{}, fmt.Errorf("static credentials are empty")
	}
	return s.Value, nil
}
