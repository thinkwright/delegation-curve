# Stage 1: Build SvelteKit frontend
FROM node:22-alpine AS frontend

WORKDIR /app/frontend
COPY frontend/package*.json ./
RUN npm ci
COPY frontend/ ./
RUN npm run build

# Stage 2: Build Go server with embedded frontend
FROM golang:1.25-alpine AS builder

RUN apk add --no-cache git ca-certificates

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY cmd/server/ ./cmd/server/
COPY --from=frontend /app/frontend/build ./cmd/server/static/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s" \
    -o /curve-server ./cmd/server

# Stage 3: Minimal runtime
FROM gcr.io/distroless/static-debian12:nonroot-amd64

WORKDIR /app

COPY --from=builder /curve-server /app/curve-server

ENV PORT=8080

EXPOSE 8080

ENTRYPOINT ["/app/curve-server"]
