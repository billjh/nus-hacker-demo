# build stage
FROM golang:1.15-alpine AS build

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /src
COPY main.go .
RUN go build -o /app .

# final stage
FROM scratch

WORKDIR /app
COPY --from=build /app .
ENTRYPOINT [ "/app/app" ]