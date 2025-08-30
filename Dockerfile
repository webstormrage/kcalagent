# ---------- 1) Frontend build (Node v22.18.0) ----------
FROM node:22.18.0 AS frontend
WORKDIR /app/frontend

# deps
COPY frontend/package.json frontend/package-lock.json* ./
RUN npm ci --no-audit --no-fund

# build
COPY frontend ./
RUN npm run build

# Если сборка положила файлы в build (CRA) — переименуем в dist
RUN if [ -d build ] && [ ! -d dist ]; then mv build dist; fi


# ---------- 2) Backend build (Go 1.23.4) ----------
FROM golang:1.23.4 AS gobuilder
WORKDIR /src

# кеш модулей
COPY go.mod go.sum ./
RUN go mod download

# код приложения
COPY . .

# бинарь (linux/amd64, без CGO)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/app ./cmd/


# ---------- 3) Runtime (distroless) ----------
FROM gcr.io/distroless/static:nonroot
WORKDIR /app

# бинарь и статик (ожидается ./frontend/dist)
COPY --from=gobuilder /bin/app /app/app
COPY --from=frontend /app/frontend/dist /app/frontend/dist

EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/app/app"]