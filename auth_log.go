package godummylogs

import "time"

const (
	AUTH_LOG_INVALID_LINE_1 = "%s %s %s[%d] Invalid user %s from %s port %d\n"
	AUTH_LOG_INVALID_LINE_2 = "%s %s %s[%d] Connection closed by invalid user %s from %s port %d [preauth]\n"

	AUTH_LOG_VALID_LINE_1 = "%s %s %s[%d] Accepted password for %s from %s port %d ssh2\n"
	AUTH_LOG_VALID_LINE_2 = "%s %s %s[%d] pam_unix(sshd:session): session opened for user %s[uid=%d] by (uid=0)\n"
)

// auth_log is a sample struct for auth.log in linux
type auth_log struct {
	// Log time stamp: Nov 10 01:44:15
	Date time.Time `json:"date" xml:"date"`

	// Linux Hostname, will use comma-delimited hostnames in godummylogs.conf
	ServerName string `json:"host" xml:"host"`

	// Process Name
	// to simplified it, currently we will use sshd[random process id]
	Process string `json:"process" xml:"process"`

	// User name
	Username string `json:"username" xml:"username"`

	// IP address
	IP string `json:"ip" xml:"ip"`

	//

}
