// made by recanman
package registration

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/MoneroKon/wifireg/internal/config"
	"github.com/MoneroKon/wifireg/internal/opnsense"
)

var (
	client    *opnsense.Client
	userMutex sync.Mutex
)

func Init() {
	var err error

	skipVerify := config.OpnsenseCert() == ""
	clientConfig := &opnsense.ClientConfig{
		Host:       config.OpnsenseHost(),
		SkipVerify: skipVerify,
		CertPath:   config.OpnsenseCert(),
		Key:        config.OpnsenseKey(),
		Secret:     config.OpnsenseSecret(),
	}

	if client, err = opnsense.New(*clientConfig); err != nil {
		log.Fatalf("Could not initialize OPNsense client: %s", err.Error())
	}
}

func HandleAddUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	var user AddUserRequest
	var err error
	if err = json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if !OPNSENSE_USERNAME_REGEX.MatchString(user.Username) {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	} else if !OPNSENSE_PASSWORD_REGEX.MatchString(user.Password) {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// FreeRADIUS does not care about duplicate users so we have to check ourselves
	exists, err := checkUsers(user.Username)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Failed to process registration", http.StatusInternalServerError)
		return
	} else if exists {
		// We don't really care about user enumeration here, do we?
		http.Error(w, "Failed to process registration", http.StatusConflict)
		return
	}

	registerUser(w, user)
}
