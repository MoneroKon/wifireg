// made by recanman
package registration

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type RadiusSearchUsersRequest struct {
	Current      uint              `json:"current"`
	RowCount     int               `json:"rowCount"`
	SearchPhrase string            `json:"searchPhrase"`
	Sort         map[string]string `json:"sort"`
}

type RadiusSearchUsersResponse struct {
	Users []RadiusUser `json:"rows"`
}

func checkUsers(username string) (bool, error) {
	userMutex.Lock()
	defer userMutex.Unlock()

	radiusReq := RadiusSearchUsersRequest{
		Current:  1,
		RowCount: -1,
	}

	startTime := time.Now()
	res, err := client.Post("/api/freeradius/user/searchUser", radiusReq)
	elapsedTime := time.Since(startTime)

	log.Printf("API request to search users took %s", elapsedTime)

	if err != nil {
		return false, fmt.Errorf("failed to send request to opnsense server")
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return false, fmt.Errorf("opnsense response was not OK")
	}

	var users RadiusSearchUsersResponse
	if err = json.NewDecoder(res.Body).Decode(&users); err != nil {
		return false, fmt.Errorf("failed to decode user list: %v", err)
	}

	return userExists(username, users.Users)
}

func userExists(username string, users []RadiusUser) (bool, error) {
	for _, existingUser := range users {
		if existingUser.Username == username {
			return true, fmt.Errorf("user with same username already exists")
		}
	}
	return false, nil
}
