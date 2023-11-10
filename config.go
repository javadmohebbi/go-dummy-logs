package godummylogs

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
)

type Usernames struct {
	User     string
	UserID   string
	Fullname string
	Email    string
}

type Config struct {
	Hostnames []string
	GoodUsers []Usernames
}

// Read and parse confing.conf and parse the values
func ParseConfig(cnf string) Config {
	c := Config{}
	var hosts []string
	var goodpeople []Usernames

	f, err := os.Open(cnf)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		_t := scanner.Text()
		// fmt.Println(_t)

		re_comment, _ := regexp.Compile(`^(.*?#)|^\s*$`)
		if re_comment.MatchString(_t) {
			continue
		}

		re := regexp.MustCompile(`(\w+)\s*=\s*(.*)`)
		matches := re.FindStringSubmatch(_t)
		if len(matches) > 2 {
			// fmt.Printf("%s = %s\n", matches[1], matches[2])
			switch matches[1] {
			case "hostnames":
				// hosts = append(hosts, strings.Split(matches[2], ",")...)
				for _, h := range strings.Split(matches[2], ",") {
					hosts = append(hosts, strings.TrimSpace(h))
				}
			case "goodusers":
				for _, gu := range strings.Split(matches[2], ",") {
					u_id := strings.Split(gu, ";")
					if len(u_id) == 4 {
						goodpeople = append(goodpeople, Usernames{
							User:     strings.TrimSpace(u_id[0]),
							UserID:   strings.TrimSpace(u_id[1]),
							Fullname: strings.TrimSpace(u_id[2]),
							Email:    strings.TrimSpace(u_id[3]),
						})
					} else {
						log.Printf("Invalid user for the value of '%v'. this value will be ignored", u_id)
					}
				}
			}
		}

	}

	c.GoodUsers = goodpeople
	c.Hostnames = hosts

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return c
}
