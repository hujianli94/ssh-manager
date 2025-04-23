FROM golang:1.16-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# 获取当前 Git Commit 作为版本号
ARG GIT_COMMIT
ENV VERSION=${GIT_COMMIT}

# 如果没有 Git Commit，则使用 "dev" 作为默认版本
RUN if [ -z "${VERSION}" ]; then \
    VERSION="dev"; \
fi

RUN go build -ldflags "-X main.Version=${VERSION}" -o ssh-manager main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/ssh-manager /app/ssh-manager

ENTRYPOINT ["/app/ssh-manager"]