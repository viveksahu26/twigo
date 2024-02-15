package send

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/twilio/twilio-go"

	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

func SendMsg(ctx context.Context, source, destination, body string, args []string) error {
	accountSid := ""
	authToken := ""

	params := &twilioApi.CreateMessageParams{}
	// params.SetTo("+")
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
	accountSid := ""
	authToken := ""

	params := &twilioApi.CreateMessageParams{}
	// params.SetTo("+")
	// params.SetFrom("+19519774919")
	// params.SetBody("Hello from Go again !!")

	params.SetTo("whatsapp:" + source)
	params.SetFrom("whatsapp:" + destination)
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

func SendVoiceMsg(ctx context.Context, source, destination, body string, args []string) error {
	accountSid := ""
	authToken := ""

	params := &twilioApi.CreateCallParams{}
	// params.SetTo("+")
	// params.SetFrom("+19519774919")
	// params.SetBody("Hello from Go again !!")

	params.SetTo(source)
	params.SetFrom(destination)
	params.SetUrl("http://demo.twilio.com/docs/voice.xml")

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	resp, err := client.Api.CreateCall(params)
	if err != nil {
		return fmt.Errorf("Error sending SMS message: " + err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
	}
	return nil
}
