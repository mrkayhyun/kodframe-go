package server

import (
	"github.com/kodframe-go/api"
	"github.com/kodframe-go/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func Start() {
	// 릴리즈 모드 설정
	//gin.SetMode(gin.ReleaseMode)

	// Gin의 기본 라우터 생성
	r := gin.Default()

	// Swagger 문서 설정
	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	ginSwagger.WrapHandler(swaggerfiles.Handler,
		ginSwagger.URL("http://localhost:8080/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1))

	// 라우터 설정
	api.SetupRoutes(r)

	// 서버 실행
	r.Run(":8080") // 기본적으로 localhost:8080에서 실행
}
