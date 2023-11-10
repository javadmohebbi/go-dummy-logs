package godummylogs

import "time"

func parseDuration(s, _defult string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		d, _ = time.ParseDuration(_defult)
	}
	return d
}
