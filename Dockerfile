# FROM golang:alpine
#
# WORKDIR /app
#
# COPY . .
#
# RUN go build -o main .
#
# EXPOSE 8080
#
# CMD ["./main"]


FROM golang:alpine3.19 as builder

ENV APP_HOME /app

WORKDIR "$APP_HOME"
COPY src/ .

RUN go mod download
RUN go mod verify
RUN go build -o main

FROM golang:alpine3.19

ENV APP_HOME /app
RUN mkdir -p "$APP_HOME"
WORKDIR "$APP_HOME"

COPY src/ app/
COPY --from=builder "$APP_HOME"/main $APP_HOME

EXPOSE 8080
CMD ["./main"]
