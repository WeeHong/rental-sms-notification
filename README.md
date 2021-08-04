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

### Step 1: Activate Sheet API

1. Login to the [Google Cloud Console](https://console.cloud.google.com/)
1. Click on the **Navigation Menu** on the top left of the screen
1. Choose **APIs & Services** > **Library**
1. Enter _Google Sheets API_ on the search bar
1. Click **Enabled** button

### Step 2: Create credentials

1. Login to the [Google Cloud Console](https://console.cloud.google.com/)
1. Click on the **Navigation Menu** on the top left corner of the screen
1. Choose **APIs & Services** > **Credentials**
1. Click on **Create Credentials** > **OAuth client ID**
1. Select _Desktop app_ option under the **Application Type** (Server-side)
1. Rename the **Name** if necessary
1. Click **Create** button
1. Download the JSON file after the pop out

### Step 3: Add test users

1. Login to the [Google Cloud Console](https://console.cloud.google.com/)
1. Click on the **Navigation Menu** on the top left corner of the screen
1. Choose **APIs & Services** > **OAuth consent screen**
1. Click on **Add users** button under the **Test users** section
1. Add the email address you would like to asssociate with, click **Save**

### Step 4: Generate token

This step is ambiguous.

1. Once you run your API to trigger the Google Sheet API, you will notice there is a chunk of text similar below.

```
Go to the following link in your browser then type the authorization code:
https://accounts.google.com/o/oauth2/auth?access_type=offline&client_id=<client_id>&redirect_uri=<redirect_uri>&response_type=code&scope=https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fspreadsheets.readonly&state=state-token
```

1. Paste the link in your browser or
   - `Ctrl + Left click` for Windows user
   - `Cmd + Left click` for macOS user
1. You will see a screen like this.
   ![Step-1.png](https://postimg.cc/LgYt2Kyp)
1. Click **Continue**
1. Check the **Checkbox** and Click **Continue**
   ![Step-2.png](https://postimg.cc/LnPtWSYk)
1. Copy the code and save it as `token.json` file and place it on the root directory
   ![Step-3.png](https://postimg.cc/z36wpvBX)

You should not be getting 403 Permission Denied if you follow the step.
Otherwise, back to Step 3 again.
![Permission-Denied-Error.png](https://postimg.cc/XXpcwdnz)
