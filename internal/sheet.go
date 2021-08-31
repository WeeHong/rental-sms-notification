package sheet

import (
	"context"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type TenantInfo struct {
	Contact string
	Amount  string
}

func extractValue(resp *sheets.ValueRange, err error) []TenantInfo {
	var info []TenantInfo
	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
	} else {
		for _, row := range resp.Values {
			amount := row[1].(string)
			if len(amount) <= 0 {
				log.Fatalf("Unable to display amount correctly")
			}

			if err != nil {
				log.Fatalf("Failed to parse amount")
			}

			if len(row[0].(string)) > 0 {
				info = append(info, TenantInfo{
					row[0].(string),
					row[2].(string),
				})
			}
		}
	}
	return info
}

func TenantList() []TenantInfo {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx := context.Background()

	dirname, err := os.Getwd()
	if err != nil {
		log.Fatalf("Unable to locate the directory name: %v", err)
	}

	dirname, err = os.Getwd()

	if err != nil {
		panic(err)
	}

	filepath := path.Join(dirname, "credentials.json")

	client, err := sheets.NewService(ctx, option.WithCredentialsFile(filepath))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	readRange := "Current!A2:C5"
	resp, err := client.Spreadsheets.Values.Get(os.Getenv("SHEET_ID"), readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	return extractValue(resp, err)
}
