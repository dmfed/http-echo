# base image
FROM golang:1.16-alpine
# creating working directory
WORKDIR /echo
# copying go.mod to WORKDIR
COPY go.mod .
# could use RUN go mod download
# copying source code
COPY *.go .
# building 
RUN go build -o /echo/server
# opening container's port to the outside
EXPOSE 8080
# running the app
CMD ["/echo/server", "-ip", "0.0.0.0"]

