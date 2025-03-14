package main

import (
	"api-server/handlers"
	"api-server/storage"
	"log"

	"github.com/gin-gonic/gin"
)

// Пакет для логирования
// Импортирование модуля с обработчиками запросов
// Импортирование модуля для работы с базой данных
func main() {

	// Инициализация базы данных
	if err := storage.InitDatabase(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err) // Логирование ошибки и завершение программы, если инициализация базы данных не удалась
	}

	// Создание нового роутера Gin с настройками по умолчанию
	router := gin.Default()

	// Определение маршрутов и привязка их к соответствующим обработчикам
	router.GET("/notes", handlers.GetNotes)        // Маршрут для получения всех заметок
	router.GET("/notes/:id", handlers.GetNoteByID) // Маршрут для получения заметки по ID

	router.POST("/notes", handlers.CreateNote) // Маршрут для создания новой заметки

	router.PUT("/notes/:id", handlers.UpdateNoteByID)    // Маршрут для обновления заметки по ID
	router.DELETE("/notes/:id", handlers.DeleteNoteByID) // Маршрут для удаления заметки по ID

	// Запуск веб-сервера на порту 8080
	router.Run(":8080")
}

// var apiKey = "QJ1WDJHZ2OOMUWM4"

// func getSymbols(c *gin.Context) {
// 	symbol_list := GetSymbols(apiKey)

// 	c.JSON(200, gin.H{
// 		"symbols": symbol_list,
// 	})
// }

// func pong(c *gin.Context) {
// 	c.JSON(200, gin.H{
// 		"messsage": "pong",
// 	})
// }

// func postUser(c *gin.Context) {
// 	user := struct {
// 		Id   string `json:"id"`
// 		Name string `json:"name"`
// 	}{}
// 	c.BindJSON(&user) // Reading data in 'post' request

// 	user.Name = "Петя"
// 	c.JSON(200, gin.H{
// 		"user": user,
// 	})
// }

// func getUser(c *gin.Context) {
// 	id := c.Query("id") // Reading "id" in a 'get' request
// 	//name := c.Query("name")

// 	user := struct {
// 		Id   string `json:"id"`
// 		Name string `json:"name"`
// 	}{id, "Vasya"}

// 	c.JSON(200, gin.H{
// 		"user": user,
// 	})
//}
