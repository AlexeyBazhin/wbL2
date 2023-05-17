package task

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func GetNTPTime(host string) (time.Time, error) {
	// получаем текущее время с NTP сервера с дополнительными метаданными
	response, err := ntp.Query(host)
	if err != nil {
		if _, err := fmt.Fprintln(os.Stderr, err); err != nil {
			log.Fatal(err)
		}
		return time.Time{}, err
	}
	// Предполагаемое смещение часов локальной системы относительно часов сервера прибавляем ко времени локальной системы
	return time.Now().Add(response.ClockOffset), nil
}
