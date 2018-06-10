FROM debian:latest

RUN apt update && apt-get install -y ca-certificates

COPY kubernetes-updown-io-controller /

ENV API_KEY "xxxxxxxxxx"

CMD ["/kubernetes-updown-io-controller","-apikey=$API_KEY"]