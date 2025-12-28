package subscriptions

import (
	"net/http"
	"online_subscription_service/internal/domain/models"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type EditSubscriptionResponse struct {
	Status string `json:"status"`
}

// editSubscription — HTTP-обработчик для редактирования существующей подписки.
//
// - Получает ID подписки из URL-параметра.
// - Валидирует и парсит UUID.
// - Привязывает тело запроса к структуре EditSubRequest.
// - Вызывает сервисный слой для обновления подписки.
// - Возвращает статус 200 при успешном обновлении или соответствующую ошибку.
//
// EditSubscription godoc
// @Summary Edit subscription
// @Description Edit subscription in service
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param        id path string true "Subscription ID" format(uuid)
// @Param        request body models.EditSubRequest true "Subscription data"
// @Success      200 {object} EditSubscriptionResponse
// @Failure      400 {object} models.ErrorResponse
// @Failure      500 {object} models.ErrorResponse
// @Router       /subscriptions/{id} [patch]
func (h *Handlers) editSubscription(c echo.Context) error {
	r := new(models.EditSubRequest)

	param := c.Param("id")

	if param == "" {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "param ID is required"})
	}

	uuid, err := uuid.Parse(param)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
	}

	if err := c.Bind(&r); err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
	}

	ctx := c.Request().Context()

	err = h.subsService.EditSubscription(ctx, uuid, *r.ToSubsUpdateDTO())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, EditSubscriptionResponse{Status: "Ok"})
}
