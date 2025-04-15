# kodframe-go

gin framework 를 이용한 rest api 서버

## Run
```bash
# Run the server
go run main.go
# Run the server with hot reload
air
```

## Docker 
```bash
# Build the image
docker build -t kodframe-go .

# Run the container
docker run -p 8080:8080 kodframe-go
```

## Swagger
```bash
# 파일 변경시 
swag init
```

## Setting list
- [ ] github action ci 설정
- [ ] gin framework 폴더 구조잡기
- [ ] 배포설정하기

## Reference
- https://github.com/gin-gonic/gin
- https://github.com/swaggo/gin-swagger
