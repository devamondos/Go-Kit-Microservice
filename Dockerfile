# golang:1.12.7-alpine3.10
FROM golang@sha256:0eb3e3a320d5711006ac3b8afa0bf5b100e16966bf61351dfa8f7c99a9238260 AS BUILDER
WORKDIR /app/

# install git for fetching dependencies
RUN apk add --quiet --no-cache git

# creates low priviledge user to run app
RUN adduser -D -g '' RUNNER

# copy dependencies list to container - helps with cache
ADD go.* ./

# install dependencies and verify
RUN go mod download
RUN go mod verify

# copy app to container
ADD . .

# compile code to /bin/app
RUN CGO_ENABLED=0 go build -o ./app .

# rhel-atomic:7.6-305
FROM registry.access.redhat.com/rhel-atomic@sha256:5f85cb9db7fbd40888a1ace32a9733543b4bddb0cb42632bd6ac11877909a31c

WORKDIR /

# expose port 8081
EXPOSE 8081

# uses low priviledge user created in BUILDER
COPY --from=builder /etc/passwd /etc/passwd
USER RUNNER


# copy in binary from app
COPY --from=BUILDER /app/app /

CMD ["/app"]
