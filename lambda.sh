#!/bin/sh
ENVFILE=.env
if [[ ! -f "$ENVFILE" ]]; then
    echo "$ENVFILE does not exists."
    exit
fi

go build main.go \
  && zip function.zip main .env credentials.json \
  && aws lambda update-function-code \
      --function-name go_rental_sms \
      --zip-file fileb://function.zip \
  && rm -rf function.zip main

echo "DONE"