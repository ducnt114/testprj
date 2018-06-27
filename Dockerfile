FROM golang:1.9.2 as builder
RUN mkdir -p /go/src/github.com/ducnt114
COPY . /go/src/github.com/ducnt114/testprj
WORKDIR /go/src/github.com/ducnt114/testprj
RUN make build
RUN mkdir /app && cp testprj /app && cp -r conf /app && cp run.sh /app

FROM alpine:3.7
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app /app
RUN chmod +x run.sh
# expose port for api
EXPOSE 8000
ENTRYPOINT ["./run.sh", "start"]