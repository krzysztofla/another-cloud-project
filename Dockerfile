# --------------------------------------- build stage --------------------------------------- #

# setting up image for build stage
FROM golang:1.18-bullseye as builder

WORKDIR $GOPATH/src/demo-golang/app/

# Copy files needed for build
COPY ./api .

RUN go mod download
RUN go mod verify

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /main .
# RUN CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o /main .

# --------------------------------------- mounting stage --------------------------------------- #

# Start from golang base image
FROM gcr.io/distroless/static-debian11

COPY --from=builder /main .

# Run the executable
CMD ["./main"]