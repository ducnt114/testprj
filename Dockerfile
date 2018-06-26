FROM golang:1.9.2 as builder
COPY . /go/src/github.com/ducnt114/testprj
WORKDIR /go/src/testprj
RUN cd cmd/testprj && go get && env GOOS=linux GOARCH=amd64 go build
RUN mkdir /app && cd cmd/testprj && cp testprj /app && cp -r conf /app

FROM alpine:3.7
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app /app
# expose port for api
EXPOSE 8000
ENTRYPOINT ["./testprj"]