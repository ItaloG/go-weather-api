FROM golang:latest as builder
WORKDIR /app
COPY . .
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o server cmd/server.go
# CMD [ "./server" ]

# FROM scratch
# COPY --from=builder /app/server .

CMD [ "tail", "-f", "/dev/null" ]