package handlers

import (
	"net/http" // Пакет для работы с HTTP
	"strconv"  // Пакет для конвертации строк в другие типы данных

	"api-server/storage" // Импортирование модуля для работы с базой данных

	"github.com/gin-gonic/gin" // Веб-фреймворк Gin
)

// Обработчик для получения всех заметок
func GetNotes(c *gin.Context) {
	notes := storage.GetAllNotes() // Получение всех заметок из хранилища
	c.JSON(http.StatusOK, notes)   // Возвращение заметок в формате JSON с кодом 200 (OK)
}

// Обработчик для получения заметки по ID
func GetNoteByID(c *gin.Context) {
	// Конвертация параметра ID из строки в целое число
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// Возвращение ошибки 400 (Bad Request), если ID некорректен
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note ID"})
		return
	}
	// Получение заметки по ID из хранилища
	note := storage.GetNoteByID(id)
	if note == nil {
		// Возвращение ошибки 404 (Not Found), если заметка не найдена
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return
	}
	// Возвращение найденной заметки в формате JSON с кодом 200 (OK)
	c.JSON(http.StatusOK, note)
}

// Обработчик для создания новой заметки
func CreateNote(c *gin.Context) {
	// Структура для хранения входных данных
	var input struct {
		Title   string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
	}
	// Привязка входных данных в формате JSON к структуре input
	if err := c.ShouldBindJSON(&input); err != nil {
		// Возвращение ошибки 400 (Bad Request), если входные данные некорректны
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Создание новой заметки в хранилище
	note := storage.CreateNote(input.Title, input.Content)
	// Возвращение созданной заметки в формате JSON с кодом 201 (Created)
	c.JSON(http.StatusCreated, note)
}

// Обработчик для обновления существующей заметки по ID
func UpdateNoteByID(c *gin.Context) {
	// Конвертация параметра ID из строки в целое число
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// Возвращение ошибки 400 (Bad Request), если ID некорректен
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note ID"})
		return
	}
	// Структура для хранения входных данных
	var input struct {
		Title   string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
	}
	// Привязка входных данных в формате JSON к структуре input
	if err := c.ShouldBindJSON(&input); err != nil {
		// Возвращение ошибки 400 (Bad Request), если входные данные некорректны
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Обновление заметки в хранилище
	note := storage.UpdateNote(id, input.Title, input.Content)
	if note == nil {
		// Возвращение ошибки 404 (Not Found), если заметка не найдена
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return
	}
	// Возвращение обновленной заметки в формате JSON с кодом 200 (OK)
	c.JSON(http.StatusOK, note)
}

// Обработчик для удаления заметки по ID
func DeleteNoteByID(c *gin.Context) {
	// Конвертация параметра ID из строки в целое число
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// Возвращение ошибки 400 (Bad Request), если ID некорректен
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note ID"})
		return
	}
	// Удаление заметки из хранилища
	if success := storage.DeleteNoteByID(id); !success {
		// Возвращение ошибки 404 (Not Found), если заметка не найдена
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return
	}
	// Возвращение кода 204 (No Content) при успешном удалении
	c.Status(http.StatusNoContent)
}
