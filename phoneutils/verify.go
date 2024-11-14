package phoneutils

import (
	"errors"
	"log"

	openapi "github.com/twilio/twilio-go/rest/verify/v2"
)

//sends an OTP code to the phone number passed as an argument
func SendOtp(phone_number string) error {
	params := &openapi.CreateVerificationParams{}
	params.SetTo(phone_number)
	params.SetChannel("sms")

	resp, err := TwilioInstance.Client.VerifyV2.CreateVerification(TwilioInstance.VerifyServiceSid, params)

	if err != nil {
		log.Println(err.Error())
		return errors.New("error sending verification code to: " + phone_number + " please verify format")
	} else {
		log.Printf("Sent verification '%s'\n", *resp.Sid)
	}
	return nil
}

//verifies the validity of an OTP code with regard to the phone number. Takes both the OTP and the phone number as arguments
func CheckOtp(phone_number string, code string) error {

	params := &openapi.CreateVerificationCheckParams{}
	params.SetTo(phone_number)
	params.SetCode(code)

	resp, err := TwilioInstance.Client.VerifyV2.CreateVerificationCheck(TwilioInstance.VerifyServiceSid, params)

	if err != nil {
		log.Println(err.Error())
		return errors.New("error vefifying code")
	} else if *resp.Status == "approved" {
		log.Println("Entered Correct Code!")
	} else {
		log.Println("Entered Incorrect Code!")
		return errors.New("entered incorrect code")
	}
	return nil
}
