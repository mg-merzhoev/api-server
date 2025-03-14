package storage

import (
	"api-server/models" // Импортирование пакета с определением моделей данных
)

// Функция для получения всех заметок из базы данных
func GetAllNotes() []models.Note {
	var notes []models.Note
	// Использование GORM для выполнения SQL-запроса SELECT и заполнения среза notes
	db.Find(&notes)
	return notes // Возвращение всех найденных заметок
}

// Функция для получения заметки по ID
func GetNoteByID(id int) *models.Note {
	var note models.Note
	// Использование GORM для выполнения SQL-запроса SELECT с условием WHERE id = id заметки
	if result := db.First(&note, id); result.Error != nil {
		return nil // Возвращение nil, если заметка с указанным ID не найдена
	}
	return &note // Возвращение найденной заметки
}

// Функция для создания новой заметки
func CreateNote(title, content string) models.Note {
	note := models.Note{
		Title:   title,
		Content: content,
	}
	// Использование GORM для выполнения SQL-запроса INSERT и сохранения новой заметки в базе данных
	db.Create(&note)
	return note // Возвращение созданной заметки
}

// Функция для обновления существующей заметки по ID
func UpdateNote(id int, title, content string) *models.Note {
	var note models.Note
	// Использование GORM для выполнения SQL-запроса SELECT с условием WHERE id = id заметки
	if result := db.First(&note, id); result.Error != nil {
		return nil // Возвращение nil, если заметка с указанным ID не найдена
	}
	note.Title = title
	note.Content = content
	// Использование GORM для выполнения SQL-запроса UPDATE и сохранения обновленной заметки в базе данных
	db.Save(&note)
	return &note // Возвращение обновленной заметки
}

// Функция для удаления заметки по ID
func DeleteNoteByID(id int) bool {
	// Использование GORM для выполнения SQL-запроса DELETE с условием WHERE id = id заметки
	if result := db.Delete(&models.Note{}, id); result.Error != nil {
		return false // Возвращение false, если удаление не удалось
	}
	return true // Возвращение true при успешном удалении заметки
}
