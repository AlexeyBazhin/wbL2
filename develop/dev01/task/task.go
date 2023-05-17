package task

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func GetNTPTime(host string) (time.Time, error) {
	response, err := ntp.Query(host)
	if err != nil {
		_, err := fmt.Fprintln(os.Stderr, err)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(1)
	}
	return time.Now().Add(response.ClockOffset), nil
}
