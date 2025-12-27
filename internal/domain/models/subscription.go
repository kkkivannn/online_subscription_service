package models

import (
	"time"

	"github.com/google/uuid"
)

// Subs — модель подписки для хранения в базе данных.
// Содержит ID, название услуги, цену, ID пользователя, дату начала и необязательную дату окончания.
type Subs struct {
	ID        uuid.UUID  `json:"id"`
	Name      string     `json:"service_name"`
	Price     int        `json:"price"`
	UserID    uuid.UUID  `json:"user_id"`
	StartDate time.Time  `json:"start_date"`
	EndDate   *time.Time `json:"end_date"`
}

// SubsDTO — Data Transfer Object для подписки.
// Используется для передачи данных между слоями приложения.
type SubsDTO struct {
	ID        uuid.UUID  `json:"id"`
	Name      string     `json:"service_name"`
	Price     int        `json:"price"`
	UserID    uuid.UUID  `json:"user_id"`
	StartDate time.Time  `json:"start_date"`
	EndDate   *time.Time `json:"end_date"`
}

// SubsUpdateDTO — DTO для обновления подписки. Все поля опциональны.
type SubsUpdateDTO struct {
	Name      *string    `json:"service_name"`
	Price     *int       `json:"price"`
	UserID    *uuid.UUID `json:"user_id"`
	StartDate *time.Time `json:"start_date"`
	EndDate   *time.Time `json:"end_date"`
}

// AddSubRequest — структура запроса на создание подписки через HTTP.
type AddSubRequest struct {
	Name   string    `json:"service_name"`
	Price  int       `json:"price"`
	UserID uuid.UUID `json:"user_id"`
}

// EditSubRequest — структура запроса на редактирование подписки через HTTP.
type EditSubRequest struct {
	Name      *string    `json:"service_name"`
	Price     *int       `json:"price"`
	UserID    *uuid.UUID `json:"user_id"`
	StartDate *time.Time `json:"start_date"`
	EndDate   *time.Time `json:"end_date"`
}

// PricePeriodRequest — структура запроса для получения цены за период.
type PricePeriodRequest struct {
	From   time.Time `json:"from" query:"from"`
	To     time.Time `json:"to" query:"to"`
	UserID uuid.UUID `json:"user_id" query:"user_id"`
	Name   string    `json:"service_name" query:"service_name"`
}

// Методы конвертации

// ToSubsDTO — конвертирует AddSubRequest в DTO для хранения в сервисном слое.
// Устанавливает дату начала подписки на текущий момент.
func (s *AddSubRequest) ToSubsDTO() *SubsDTO {
	return &SubsDTO{
		Name:      s.Name,
		Price:     s.Price,
		UserID:    s.UserID,
		StartDate: time.Now(),
		EndDate:   nil,
	}
}

// ToSubsUpdateDTO — конвертирует EditSubRequest в DTO для обновления.
func (s *EditSubRequest) ToSubsUpdateDTO() *SubsUpdateDTO {
	return &SubsUpdateDTO{
		Name:      s.Name,
		Price:     s.Price,
		UserID:    s.UserID,
		StartDate: s.StartDate,
		EndDate:   s.EndDate,
	}
}

// ToSubs — конвертирует SubsDTO обратно в модель для работы с базой данных.
func (s *SubsDTO) ToSubs() Subs {
	return Subs{
		ID:        s.ID,
		Name:      s.Name,
		Price:     s.Price,
		UserID:    s.UserID,
		StartDate: s.StartDate,
		EndDate:   s.EndDate,
	}
}
