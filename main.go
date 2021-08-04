package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
	sheet "weehong.me/rental-sms-notification/internal"
)

func handleRequest() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	tenantList := sheet.TenantList()

	from := os.Getenv("TWILIO_MESSAGING_SID")
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")

	client := twilio.NewRestClientWithParams(twilio.RestClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := &openapi.CreateMessageParams{}

	for _, tenant := range tenantList {
		params.SetTo("+" + tenant.Contact)
		params.SetFrom(from)
		params.SetBody("Dear #26-66 tenant,\nPlease transfer your rental and utilise today.\n\nTotal amount: " + tenant.Amount)

		resp, err := client.ApiV2010.CreateMessage(params)
		if err != nil {
			fmt.Println(err.Error())
			err = nil
		} else {
			response, _ := json.Marshal(*resp)
			fmt.Println("Response: " + string(response))
		}
	}
}

func main() {
	lambda.Start(handleRequest)
}
