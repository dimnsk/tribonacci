FROM golang:1.8

COPY src /app/
WORKDIR /app

ENTRYPOINT ["/app/bin/run.sh"]