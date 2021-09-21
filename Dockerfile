FROM golang:alpine

COPY . .

ENV LOGGLY_TOKEN=fdc31234-adb1-41b1-85ad-9d8ffe1544fa


CMD ["./assignment4"]