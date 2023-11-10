package godummylogs

import "time"

const (
	AUTH_LOG_INVALID_LINE_1 = "%s %s %s[%d] Invalid user %s from %s port %d\n"
	AUTH_LOG_INVALID_LINE_2 = "%s %s %s[%d] Connection closed by invalid user %s from %s port %d [preauth]\n"

	AUTH_LOG_VALID_LINE_1 = "%s %s %s[%d] Accepted password for %s from %s port %d ssh2\n"
	AUTH_LOG_VALID_LINE_2 = "%s %s %s[%d] pam_unix(sshd:session): session opened for user %s[uid=%d] by (uid=0)\n"
)

// AuthLog is a sample struct for auth.log in linux
type AuthLogConf struct {
	File                      string        `json:"file"`
	KnownUsersSuccessInterval time.Duration `json:"knownUsers_success_interval"`
	KnownUsersFaieldInterval  time.Duration `json:"knownUsers_failed_interval"`
	OtherFailed               time.Duration `json:"other_failed"`
}
