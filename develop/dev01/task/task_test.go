package task

import "testing"

func TestGetNTPTimeFromInvalidServer(t *testing.T) {
	ntpTime, err := GetNTPTime("some invalid server")
	if err != nil {
		t.Logf("Test passed: Error getting NTP time from server: %s", err)
		return
	}
	t.Errorf("Test failed: Time from server: %v", ntpTime)
}

func TestGetNTPTimeFromValidServer(t *testing.T) {
	ntpTime, err := GetNTPTime("0.ru.pool.ntp.org")
	if err != nil {
		t.Errorf("Test failed: Error getting NTP time from server: %s", err)
	}
	t.Logf("Test passed: time from server: %v", ntpTime)
}
