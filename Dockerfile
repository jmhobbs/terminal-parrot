FROM golang:alpine
WORKDIR /project
RUN apk update && apk add --no-cache git
COPY go.* ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o parrot .

FROM scratch
COPY --from=0 /project/parrot /parrot
COPY animations/ /etc/terminal-parrot
ENTRYPOINT ["/parrot"]

