FROM golang:1-alpine3.20 AS build

WORKDIR /app/

RUN apk add build-base

COPY . .

RUN go mod tidy
RUN CGO_ENABLED=1 go build -o ./stwhh-mensa main.go



FROM alpine:3.20 AS final

WORKDIR /app/

COPY --from=build /app/stwhh-mensa /app/stwhh-mensa

ENTRYPOINT [ "./stwhh-mensa" ]
