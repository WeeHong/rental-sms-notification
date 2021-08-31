# Rental SMS Notification

This is a SMS notification service which remind the tenant to pay for the rental
<br />
<br />
This service develop in `Golang` and utilise the `Google Sheet API` to read the contact number from the spreadsheet and send the message via `Twilio Messaging API`.
<br />
<br />
This service has been deployed to `AWS Lambda` and utilise `AWS CloudWatch` to trigger the function every last day of the month.

## Notes

The Google Sheet API requires `credentials.json` and `token.json` to works.
<br />
<br />

**The steps to generate the credentials might vary based on the present UI**

Based on [Using OAuth 2.0 for Web Server Applications](https://developers.google.com/identity/protocols/oauth2/web-server)

_Certain portion of the instrument wrote on this page are outdated. Hence, I create the step below._

### Step 1: Create an

1. Go to the [Credentials Page](https://console.developers.google.com/apis/credentials)
1. Click **Crete credentials** > **Service account**
1. Fill in the form and click **Create and Continue** button

### Step 2: Generate a JSON API key

1. Go to the [Credentials Page](https://console.developers.google.com/apis/credentials)
1. Click the service account you created earlier
1. Click the **Keys** tab
1. Click the **Add Key** button and select **Create new key**
1. Select **JSON** for the Key type and click **Create** button
1. Download the json file and place it on your project
