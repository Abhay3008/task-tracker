FROM golang:1.24-alpine as builder

RUN mkdir /app
WORKDIR /app
COPY . /app
RUN CGO_ENABLED=0 go build -o task-tracker .

RUN chmod +x task-tracker


FROM alpine:latest
RUN mkdir /app

COPY --from=builder app/task-tracker /app/task-tracker

ENTRYPOINT ["/app/task-tracker"]