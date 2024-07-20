FROM golang:1.22-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download && go mod verify

COPY *.go ./

RUN go build -o /app/maradona

# Runtime Stage
FROM alpine:3.15

COPY --from=build /app/main /bin/main

EXPOSE 8000

CMD [ "/bin/main" ]