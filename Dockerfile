FROM golang:1.16-alpine AS builder 
WORKDIR /echo
COPY go.mod .
COPY *.go .
RUN go build -o /echo/server
FROM alpine 
COPY --from=builder /echo/server /bin/echo
EXPOSE 8080
CMD ["/bin/echo", "-addr", ":8080"]

