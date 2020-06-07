FROM golang:1.14-alpine

RUN apk add git

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

ENV SLACK_VERIFICATION_TOKEN ""
ENV DEVOPS_ONDUTY_NAMESPACE ""
ENV DEVOPS_ONDUTY_CONFIGMAP ""

EXPOSE 8080
CMD ["app"]