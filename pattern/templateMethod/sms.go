package templateMethod

import "fmt"

type SMS struct {
	OTP
}

func (s *SMS) GenRandomOTP(len int) string {
	randomOTP := "1234"
	fmt.Printf("SMS: generating random otp %s\n", randomOTP)
	return randomOTP
}

func (s *SMS) SaveOTPCache(otp string) {
	fmt.Printf("SMS: saving otp: %s to cache\n", otp)
}

func (s *SMS) GetMessage(otp string) string {
	return "SMS OTP for login is " + otp
}

func (s *SMS) SendNotification(message string) error {
	fmt.Printf("SMS: sending SMS: %s\n", message)
	return nil
}

func (s *SMS) PublishMetric() {
	fmt.Printf("SMS: publishing metrics\n")
}
