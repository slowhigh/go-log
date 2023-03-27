// Concrete implementation
package templatemethod

import (
	"log"
	"math/rand"
	"strconv"
)

type email struct {
	op otp
	uid string
}

func (e *email) genRandomOTP() string {
	randomOTP := e.uid + strconv.Itoa(rand.Intn(8999) + 1000)
	log.Printf("EMAIL: generating random otp %s\n", randomOTP)
	return randomOTP
}

func (e *email) saveOTPCache(otp string) {
	log.Printf("EMAIL: saving otp: %s to cache\n", otp)
}

func (e *email) getMessage(otp string) string {
	return "EMAIL OTP for login is" + otp
}

func (e *email) sendNotification(message string) error {
	log.Printf("EMAIL: sending email: %s\n", message)
	return nil
}

func (e *email) publishMetric() {
	log.Printf("EMAIL: publishing metrics\n")
}