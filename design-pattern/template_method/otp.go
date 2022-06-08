// Template method
package templatemethod

import "strings"

type iOtp interface {
	genRandomOTP() string
	saveOTPCache(string)
	getMessage(string) string
	sendNotification(string) error
	publishMetric()
}

// type otp struct {}

// func (o *otp) genAndSendOTP(iOtp iOtp) error {
// 	otp := iOtp.genRandomOTP()
// 	iOtp.saveOTPCache(otp)
// 	message := iOtp.getMessage(otp)
// 	err := iOtp.sendNotification(message)
// 	if err != nil {
// 		return err
// 	}
// 	iOtp.publishMetric()
// 	return nil
// }

type otp struct {
	iOtp iOtp
	pw string
}

func (o *otp) genAndSendOTP() error {
	otpString := o.iOtp.genRandomOTP()
	o.iOtp.saveOTPCache(otpString)
	message := o.iOtp.getMessage(otpString)
	o.pw = otpString
	err := o.iOtp.sendNotification(message)
	if err != nil {
		return err
	}

	o.iOtp.publishMetric()
	return nil
}

func (o *otp) checkOtpUid(uid string) bool {
	return strings.Contains(o.pw, uid)
}