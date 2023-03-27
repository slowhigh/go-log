// Concrete implementation
package templatemethod

import (
	"log"
	"math/rand"
	"strconv"
)

type sms struct {
	otp otp
	uid string
}

func (s *sms) genRandomOTP() string {
	randomOTP := s.uid + strconv.Itoa(rand.Intn(8999) + 1000)
	log.Printf("SMS: generating random otp %s\n", randomOTP)
	return randomOTP
}

func (s *sms) saveOTPCache(otp string) {
	log.Printf("SMS: saving otp: %s to cache\n", otp)
}

func(s *sms) getMessage(otp string) string {
	return "SMS OTP for login is " + otp
}

func (s *sms) sendNotification(message string) error {
	log.Printf("SMS: sending sms: %s\n", message)
	return nil
}

func (s *sms) publishMetric() {
	log.Printf("SMS: publishing metrics\n")
}