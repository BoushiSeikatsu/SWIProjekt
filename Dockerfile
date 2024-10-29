FROM golang:1.23-alpine AS build

WORKDIR /app

COPY backend .

ENV CGO_ENABLED=1

RUN apk add --no-cache gcc musl-dev libc-dev
RUN go mod download
RUN go build -o app .

FROM alpine:latest AS runtime

RUN apk add --no-cache libc6-compat

WORKDIR /app
COPY --from=build /app/app app
COPY --from=build /app/config.yaml config.yaml

EXPOSE 8083

ENTRYPOINT ["/app/app"]
