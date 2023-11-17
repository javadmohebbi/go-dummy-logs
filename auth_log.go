package godummylogs

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/go-faker/faker/v4"
)

const (
	AUTH_LOG_INVALID_LINE_1 = "%s %s %s[%d] Invalid user %s from %s port %d\n"
	AUTH_LOG_INVALID_LINE_2 = "%s %s %s[%d] Connection closed by invalid user %s from %s port %d [preauth]\n"

	AUTH_LOG_VALID_LINE_1 = "%s %s %s[%d] Accepted password for %s from %s port %d ssh2\n"
	AUTH_LOG_VALID_LINE_2 = "%s %s %s[%d] pam_unix(sshd:session): session opened for user %s[uid=%d] by (uid=0)\n"
)

// AuthLog is a sample struct for auth.log in linux
type ConfAuthLog struct {
	File                      string `json:"file"`
	KnownUsersSuccessInterval string `json:"knownUsers_success_interval"`
	KnownUsersFaieldInterval  string `json:"knownUsers_failed_interval"`
	OtherFailed               string `json:"other_failed"`
}

type AuthLog struct {
	mu            sync.Mutex
	f             *os.File
	conf          Config
	path          string
	known_success time.Duration
	known_failure time.Duration
	other_failure time.Duration
}

func NewAuthLog(c Config) *AuthLog {
	return &AuthLog{
		conf:          c,
		path:          filepath.Join(c.ContextPath, c.Auth.File),
		known_success: parseDuration(c.Auth.KnownUsersSuccessInterval, "15m"),
		known_failure: parseDuration(c.Auth.KnownUsersFaieldInterval, "1h"),
		other_failure: parseDuration(c.Auth.OtherFailed, "10m"),
	}
}

func (a AuthLog) Generate() {
	// log.Println(a.path)
	f, err := os.OpenFile(a.path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	a.f = f
	if err != nil {
		panic(err)
	}
	defer f.Close()

}

func (a AuthLog) generate_known_success() {
	ticker := time.NewTicker(a.known_success)
	defer func() { ticker.Stop() }()

	for {
		select {
		case <-ticker.C:
			_host := ""
			if len(a.conf.Hostnames) <= 0 {
				_host = "www1"
			} else {
				_host = a.conf.Hostnames[(rand.Intn(len(a.conf.Hostnames) - 1))]
			}

			_users := Users{}
			if len(a.conf.Known.Users) <= 0 {
				_users = Users{
					Username: "root",
					UserID:   1,
					FullName: "root",
					Email:    "root@example.local",
				}
			} else {
				_users = a.conf.Known.Users[(rand.Intn(len(a.conf.Known.Users) - 1))]
			}

			_pid := rand.Intn(100000-10000) + 10000
			_s_port := rand.Intn(60000-50000) + 50000

			line := fmt.Sprintf("%s %s %s[%d] Accepted password for %s from %s port %d ssh2\n",
				time.Now().Format("Jan 02 15:04:05"), _host, "sshd", _pid, _users.Username,
				faker.IPv4(), _s_port,
			)

			fmt.Printf("%s", line)

			a.mu.Lock()
			_, err := a.f.WriteString(line)
			if err != nil {
				log.Println("ERR:", err)
			}
			a.mu.Unlock()

		}
	}
}
