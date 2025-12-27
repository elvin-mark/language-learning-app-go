# ===== BUILD STAGE =====
FROM golang:1.25-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./ 
RUN go mod download
RUN apk add gcc musl-dev

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -o /migrate cmd/migrate/main.go
RUN CGO_ENABLED=1 GOOS=linux go build -o /language-learning-app cmd/api/main.go

RUN strip /migrate
RUN strip /language-learning-app

# ===== RUN STAGE =====
FROM alpine:latest

WORKDIR /app

COPY --from=build /language-learning-app /language-learning-app
COPY --from=build /migrate /migrate
COPY --from=build /app/migrations/*.sql /migrations/
COPY --from=build /app/resources ./resources
COPY config.development.yaml ./config.development.yaml

EXPOSE 8080

RUN echo "\
    /migrate && /language-learning-app \
    " > /app/entrypoint.sh

ENTRYPOINT ["sh", "/app/entrypoint.sh"]
