package models

// Определение структуры заметки
type Note struct {
	ID      int    `json:"id" gorm:"primaryKey;autoIncrement"` // Уникальный идентификатор заметки, автоматически инкрементируемый
	Title   string `json:"title"`                              // Заголовок заметки
	Content string `json:"content"`                            // Содержание заметки
}
