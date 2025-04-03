// made by recanman
package registration

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type RadiusAddUserRequest struct {
	User RadiusUser `json:"user"`
}

func registerUser(w http.ResponseWriter, user AddUserRequest) {
	radiusUser := RadiusUser{
		Enabled:            "1",
		Username:           user.Username,
		Password:           user.Password,
		PasswordEncryption: "Cleartext-Password",
		SimUse:             "5",
	}
	radiusReq := RadiusAddUserRequest{User: radiusUser}

	startTime := time.Now()
	res, err := client.Post("/api/freeradius/user/addUser/", radiusReq)
	elapsedTime := time.Since(startTime)

	log.Printf("API request to add user took %s", elapsedTime)

	if err != nil {
		log.Printf("Failed to send request to OPNsense server: %s", err.Error())
		http.Error(w, "Failed to process registration", http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Printf("Failed to register OPNsense user: %s", res.Status)
		http.Error(w, "Failed to process registration.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "User successfully registered")

}
