package main

import (
	"encoding/json"
	"fmt"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

// func main() {
// 	http.HandleFunc("/sms", handleSMS)
// 	fmt.Println("Server starting on port 8080...")
// 	http.ListenAndServe(":8080", nil)
// } response, _ := json.Marshal(*resp)
// fmt.Println("Response: " + string(response))

func main() {
	accountSid := "ACfed4d0cbfb8da0889f49c406ae40c24a"
	authToken := "12df78bfb3eb3d7d0500288a16807859"
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := &twilioApi.CreateMessageParams{}
	params.SetTo("+919721938076")
	params.SetFrom("+19519774919")
	params.SetBody("Hello from Go !!")

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println("Error sending SMS message: " + err.Error())
		return
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
	}

	// fmt.Fprintf(w, "Message sent successfully")
}
