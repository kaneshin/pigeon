package credentials

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// A ApplicationCredentialsProvider retrieves credentials from the current user's home
// directory, and keeps track if those credentials are expired.
type ApplicationCredentialsProvider struct {
	// Path to the shared credentials file.
	//
	// If empty will look for "GOOGLE_APPLICATION_CREDENTIALS" env variable. If the
	// env value is empty will default to current user's home directory.
	Filename string

	// retrieved states if the credentials have been successfully retrieved.
	retrieved bool
}

// NewApplicationCredentials returns a pointer to a new Credentials object
// wrapping the file provider.
func NewApplicationCredentials(filename string) *Credentials {
	return NewCredentials(&ApplicationCredentialsProvider{
		Filename:  filename,
		retrieved: false,
	})
}

// Retrieve reads and extracts the shared credentials from the current
// users home directory.
func (p *ApplicationCredentialsProvider) Retrieve() (Value, error) {
	p.retrieved = false

	filename, err := p.filename()
	if err != nil {
		return Value{}, err
	}

	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return Value{}, err
	}
	creds := Value{}
	if err := json.Unmarshal(b, &creds); err != nil {
		return Value{}, err
	}

	p.retrieved = true
	return creds, nil
}

// filename returns the filename to use to read google application credentials.
// Will return an error if the user's home directory path cannot be found.
func (p *ApplicationCredentialsProvider) filename() (string, error) {
	if p.Filename == "" {
		if p.Filename = os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"); p.Filename == "" {
			return "", fmt.Errorf("Unable to read GOOGLE_APPLICATION_CREDENTIALS")
		}
	}
	return p.Filename, nil
}
