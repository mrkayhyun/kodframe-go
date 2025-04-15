# 1단계: 빌드 스테이지
FROM golang:1.24 AS builder

WORKDIR /app

# go.mod 및 go.sum을 복사하고 의존성 설치
COPY go.mod go.sum ./
RUN go mod download

# 전체 소스 복사 및 빌드
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd

# 2단계: 실행 스테이지 (경량 이미지)
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/app .

# 포트 설정 (예: 8080)
EXPOSE 8080

CMD ["./app"]