FROM debian:stable-slim

RUN apt-get update && apt-get install -y ca-certificates

ADD notely /usr/bin/notely

# Use the absolute path to the binary
CMD ["/usr/bin/notely"]