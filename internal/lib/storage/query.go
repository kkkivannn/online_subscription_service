package storage

import (
	"fmt"
	"online_subscription_service/internal/domain/models"
	"reflect"
	"strings"

	"github.com/google/uuid"
)

func BuildUpdateQuery(serviceID uuid.UUID, sub models.SubsUpdateDTO) (string, []any) {
	fields := map[string]any{
		"name":       sub.Name,
		"price":      sub.Price,
		"user_id":    sub.UserID,
		"start_date": sub.StartDate,
		"end_date":   sub.EndDate,
	}

	set := make([]string, 0)
	args := make([]any, 0)
	idx := 1
	for column, value := range fields {
		if reflect.ValueOf(value).IsNil() {
			continue
		}

		set = append(set, fmt.Sprintf("%s = $%d", column, idx))
		args = append(args, value)
		idx++
	}

	if len(set) == 0 {
		return "", nil
	}

	query := fmt.Sprintf("update services set %s where id=$%d", strings.Join(set, ", "), idx)

	args = append(args, serviceID)

	return query, args
}
