# Rental SMS Notification

This is a SMS notification service which remind the tenant to pay for the rental
<br />
This service develop in Golang and utilise the Google Sheet API to read the contact number from the spreadsheet and send the message via Twilio Messaging API.
<br />
This service has been deployed to AWS Lambda and utilise AWS CloudWatch to trigger the function every last day of the month.
