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

// AddSubRequest — структура запроса на создание подписки через HTTP.
// Содержит минимальные данные: название услуги, цену и ID пользователя.
type AddSubRequest struct {
	Name   string    `json:"service_name"`
	Price  int       `json:"price"`
	UserID uuid.UUID `json:"user_id"`
}

type EditSubRequest struct {
	Name      *string    `json:"service_name"`
	Price     *int       `json:"price"`
	UserID    *uuid.UUID `json:"user_id"`
	StartDate *time.Time `json:"start_date"`
	EndDate   *time.Time `json:"end_date"`
}

func (s *EditSubRequest) ToSubsUpdateDTO() *SubsUpdateDTO {
	return &SubsUpdateDTO{
		Name:      s.Name,
		Price:     s.Price,
		UserID:    s.UserID,
		StartDate: s.StartDate,
		EndDate:   s.EndDate,
	}
}

// monthStart — возвращает начало месяца для указанного времени.
// Используется для установки даты начала подписки.
func monthStart(t time.Time) time.Time {
	return time.Date(
		t.Year(),
		t.Month(),
		1,
		0, 0, 0, 0,
		t.Location(),
	)
}

// ToSubsDTO — конвертирует SubRequest в DTO для хранения в сервисном слое.
// Устанавливает дату начала подписки на первый день текущего месяца.
func (s *AddSubRequest) ToSubsDTO() *SubsDTO {
	return &SubsDTO{
		Name:      s.Name,
		Price:     s.Price,
		UserID:    s.UserID,
		StartDate: monthStart(time.Now()),
		EndDate:   nil,
	}
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

// ToSubs — конвертирует DTO обратно в модель для работы с базой данных.
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

type SubsUpdateDTO struct {
	Name      *string    `json:"service_name"`
	Price     *int       `json:"price"`
	UserID    *uuid.UUID `json:"user_id"`
	StartDate *time.Time `json:"start_date"`
	EndDate   *time.Time `json:"end_date"`
}
