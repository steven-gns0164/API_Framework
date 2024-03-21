package models

type Students struct {
	NIM  int    `gorm:"primaryKey" json:"nim"`
	Nama string `gorm:"type:varchar(255)" json:"nama"`
}

type StudentsResponse struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    []Students `json:"data"`
}
