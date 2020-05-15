FROM golang as builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.* .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o client main.go 

# final stage
FROM scratch
COPY --from=builder /app/client /app/
EXPOSE 8080
ENTRYPOINT ["/app/client"]