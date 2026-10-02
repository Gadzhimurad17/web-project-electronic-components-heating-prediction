package main

import (
	"log"
	"os"

	"heating-prediction/internal/infrastructure/db"
	gormrepo "heating-prediction/internal/infrastructure/repo/gorm"
	"heating-prediction/internal/infrastructure/storage"
	"heating-prediction/internal/interfaces/http/handlers"
	"heating-prediction/internal/interfaces/http/router"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// 1. Загружаем .env (если файл есть рядом с бинарником / в корне проекта)
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  .env не найден, используем переменные окружения")
	}

	// 2. Подключение к PostgreSQL через GORM
	gormDB, err := db.NewPostgres()
	if err != nil {
		log.Fatalf("❌ Не удалось подключиться к PostgreSQL: %v", err)
	}
	log.Println("✅ PostgreSQL подключён")

	// 3. Клиент MinIO
	minioClient, err := storage.NewMinIO()
	if err != nil {
		log.Fatalf("❌ Не удалось подключиться к MinIO: %v", err)
	}
	log.Printf("✅ MinIO подключён, bucket=%s", os.Getenv("MINIO_BUCKET"))

	// 4. Репозитории (реализации интерфейсов repo.*)
	componentRepo := gormrepo.NewComponentRepo(gormDB)
	likeRepo := gormrepo.NewLikeRepo(gormDB)
	userRepo := gormrepo.NewUserRepo(gormDB)

	// 5. Хендлер с зависимостями (интерфейсы + шаблоны)
	h := handlers.NewHandler(componentRepo, likeRepo, userRepo, minioClient)
	log.Println("✅ Шаблоны загружены из ./templates/*.html")

	// 6. Gin
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	router.SetupRouter(r, h)

	// 7. Запуск
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("🚀 Сервер запущен: http://localhost:%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("❌ Ошибка запуска сервера: %v", err)
	}
}
