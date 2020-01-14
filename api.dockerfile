# For api, consumer, producer

#BASE IMAGE SETS UP GOLANG ENV AND REFLEX
FROM golang:1.13.6-alpine3.10 as builder

# Install the Certificate-Authority certificates for the app to be able to make
# calls to HTTPS endpoints.
# Git is required for fetching the dependencies.
RUN apk add --no-cache ca-certificates

# Set the environment variables for the go command:
# * CGO_ENABLED=0 to build a statically-linked executable
ENV CGO_ENABLED=0 GOOS=linux

# Set the working directory outside $GOPATH to enable the support for modules.
WORKDIR /src

#COPY CODE INTO WORKSPACE
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

#BUILD APP
RUN go build -a -installsuffix nocgo -o /app *.go

#MINIMAL CONTAINER COPIES WHAT IS NEEDED
FROM alpine

WORKDIR /work
# Import the Certificate-Authority certificates for enabling HTTPS.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Import the compiled executable from the first stage.
COPY --from=builder /app ./

# Import zoneinfo for read timezone in app
COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /usr/local/go/lib/time/zoneinfo.zip

# set host time zone 
ENV TZ=Asia/Bangkok
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# Create minimal nsswitch.conf file to prioritize the usage of /etc/hosts over DNS queries.
# This resolves the conflict between:
# * fluxd using netgo for static compilation. netgo reads nsswitch.conf to mimic glibc,
#   defaulting to prioritize DNS queries over /etc/hosts if nsswitch.conf is missing:
#   https://github.com/golang/go/issues/22846
# * Alpine not including a nsswitch.conf file. Since Alpine doesn't use glibc
#   (it uses musl), maintainers argue that the need of nsswitch.conf is a Go bug:
#   https://github.com/gliderlabs/docker-alpine/issues/367#issuecomment-354316460
RUN [ ! -e /etc/nsswitch.conf ] && echo 'hosts: files dns' > /etc/nsswitch.conf

ENTRYPOINT ["./app"]

EXPOSE 8080

LABEL   org.opencontainers.image.authors="jiraphon.sa <jiraphon.sa@tnis.com>" \
        org.opencontainers.image.created="${BUILD_DATE}" \
        org.opencontainers.image.revision="${VCS_REF}" \
        org.opencontainers.image.vendor="T.N. Incorporation Ltd."