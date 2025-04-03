// made by recanman
package registration

import "regexp"

type AddUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Reference: https://github.com/opnsense/plugins/blob/master/net/freeradius/src/opnsense/mvc/app/models/OPNsense/Freeradius/User.xml
var OPNSENSE_USERNAME_REGEX = regexp.MustCompile(`^([0-9a-zA-Z@._\-\/:]){1,128}$`)
var OPNSENSE_PASSWORD_REGEX = regexp.MustCompile(`^([0-9a-zA-Z._\-\!\$\%\/\(\)\+\#\=\{\}:&]){1,128}$`)

type RadiusUser struct {
	Enabled            string `json:"enabled"`
	Username           string `json:"username"`
	Password           string `json:"password"`
	PasswordEncryption string `json:"passwordencryption"`
	SimUse             string `json:"simuse"`
}
