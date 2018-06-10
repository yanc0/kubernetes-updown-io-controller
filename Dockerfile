FROM debian:latest

COPY kubernetes-updown-io-controller /

ENV API_KEY "xxxxxxxxxx"

CMD ["/kubernetes-updown-io-controller","-apikey=$API_KEY"]