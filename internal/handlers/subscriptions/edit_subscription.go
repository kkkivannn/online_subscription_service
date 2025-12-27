package subscriptions

import (
	"net/http"
	"online_subscription_service/internal/domain/models"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// editSubscription — HTTP-обработчик для редактирования существующей подписки.
// - Получает ID подписки из URL-параметра.
// - Валидирует и парсит UUID.
// - Привязывает тело запроса к структуре EditSubRequest.
// - Вызывает сервисный слой для обновления подписки.
// - Возвращает статус 200 при успешном обновлении или соответствующую ошибку.
func (h *Handlers) editSubscription(c echo.Context) error {
	r := new(models.EditSubRequest)

	param := c.Param("id")

	if param == "" {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": "param ID is required",
		})
	}

	uuid, err := uuid.Parse(param)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"error": err.Error(),
		})
	}

	if err := c.Bind(&r); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
	}

	ctx := c.Request().Context()

	err = h.subsService.EditSubscription(ctx, uuid, *r.ToSubsUpdateDTO())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"error": err.Error(),
		})
	}

	return c.NoContent(http.StatusOK)
}
