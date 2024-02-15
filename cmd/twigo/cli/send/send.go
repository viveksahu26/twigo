package send

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/twilio/twilio-go"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

func SendMsg(ctx context.Context, from, to, body string, args []string) error {
	accountSid := os.Getenv("TWILLIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")

	params := &twilioApi.CreateMessageParams{}

	params.SetTo(to)
	params.SetFrom(from)
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
	accountSid := os.Getenv("TWILLIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")

	params := &twilioApi.CreateMessageParams{}

	params.SetTo("whatsapp:" + source)
	params.SetFrom("whatsapp:" + destination)
	params.SetBody(body)

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		return fmt.Errorf("Error sending WhatsApps message: " + err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
	}
	return nil
}

func SendVoiceMsg(ctx context.Context, source, destination, body string, args []string) error {
	accountSid := os.Getenv("TWILLIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")

	params := &twilioApi.CreateCallParams{}

	params.SetTo(source)
	params.SetFrom(destination)
	// params.SetUrl("http://github.com/viveksahu26/twigo/blob/main/voice.xml")
	params.SetUrl("http://demo.twilio.com/docs/voice.xml")

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	resp, err := client.Api.CreateCall(params)
	if err != nil {
		return fmt.Errorf("Error sending Voice Note: " + err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
	}
	return nil
}

func SendEmailMsg(ctx context.Context, sendUserName, from, recvUserName, to, subject, htmlContent, plainTextContent string, args []string) error {
	fromEmail := mail.NewEmail(sendUserName, from)
	toEmail := mail.NewEmail(recvUserName, to)
	message := mail.NewSingleEmail(fromEmail, subject, toEmail, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
	return nil
}
