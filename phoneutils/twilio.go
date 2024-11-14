package phoneutils

import (
	"log"
	"os"

	"github.com/twilio/twilio-go"
)

var TwilioInstance *Twilio

type Twilio struct{
	Client *twilio.RestClient
	VerifyServiceSid string
}

func NewTwilio() *Twilio {
	TWILIO_ACCOUNT_SID := os.Getenv("TWILIO_ACCOUNT_SID")
	TWILIO_AUTH_TOKEN  := os.Getenv("TWILIO_AUTH_TOKEN")
	VERIFY_SERVICE_SID := os.Getenv("TWILIO_VERIFY_SERVICE_SID")
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
	   Username: TWILIO_ACCOUNT_SID,
	   Password: TWILIO_AUTH_TOKEN,
	})
	

	return &Twilio{
		Client: client,
		VerifyServiceSid: VERIFY_SERVICE_SID,
	}
}

// Initializes the TwilioInstance global variable
func InitializeTwilio(){
	TwilioInstance = NewTwilio()
	log.Println("Initialize Twilio at VerifyServiceSid: " + TwilioInstance.VerifyServiceSid)
}