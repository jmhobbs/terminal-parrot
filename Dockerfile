FROM golang:alpine3.7
WORKDIR /project
COPY *.go ./
RUN apk update && apk add --no-cache git
RUN go get github.com/nsf/termbox-go
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o parrot parrot.go data.go draw.go

FROM scratch
COPY --from=0 /project/parrot /parrot
ENTRYPOINT ["/parrot"]

