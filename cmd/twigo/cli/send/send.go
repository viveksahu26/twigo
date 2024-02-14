package send

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/twilio/twilio-go"

	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

func SendMsg(ctx context.Context, source, destination, body string, args []string) error {
	accountSid := "ACfed4d0cbfb8da0889f49c406ae40c24a"
	authToken := "12df78bfb3eb3d7d0500288a16807859"

	params := &twilioApi.CreateMessageParams{}
	// params.SetTo("+919721938076")
	// params.SetFrom("+19519774919")
	// params.SetBody("Hello from Go again !!")

	params.SetTo(source)
	params.SetFrom(destination)
	params.SetBody(body)

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		return fmt.Errorf("Error sending SMS message: " + err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
	}
	return nil
}

func SendWhatsUP(ctx context.Context, source, destination, body string, args []string) error {
	return nil
}

func SendVoiceMsg(ctx context.Context, source, destination, body string, args []string) error {
	return nil
}
