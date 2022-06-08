// Client code
package templatemethod

import "testing"

func Test_Template_Method(t *testing.T) {
// 	otp := otp{}
	
// 	smsOTP := &sms{ otp: otp }
// 	smsOTP.genAndSendOTP(smsOTP, 4)

// 	emailOTP := &email{ otp: otp }
// 	emailOTP.genAndSendOTP(emailOTP, 4)

	smsOTP := &sms{ uid: "sms" }
	emailOTP := &email{ uid: "email"}

	otp := otp{}

	otp.iOtp = smsOTP
	otp.genAndSendOTP()

	if !otp.checkOtpUid(smsOTP.uid) {
		t.Error("wrong smsOTP")
	}

	otp.iOtp = emailOTP
	otp.genAndSendOTP()

	if !otp.checkOtpUid(emailOTP.uid) {
		t.Error("wrong emailOTP")
	}
}