# ===== BUILD STAGE =====
FROM golang:1.25-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./ 
RUN go mod download
RUN apk add gcc musl-dev

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /language-learning-app cmd/api/main.go

RUN strip /language-learning-app

# ===== RUN STAGE =====
FROM scratch

WORKDIR /app

COPY --from=build /language-learning-app /language-learning-app
COPY --from=build /app/migrations/*.sql /migrations/
COPY --from=build /app/resources ./resources
COPY config.development.yaml ./config.development.yaml

EXPOSE 8080

ENTRYPOINT ["/language-learning-app"]
