package userdb

import (
	"fmt"

	"github.com/vim-diesel/new-service/business/core/user"
	"github.com/vim-diesel/new-service/business/data/order"
)

var orderByFields = map[string]string{
	user.OrderByID:      "user_id",
	user.OrderByName:    "full_name",
	user.OrderByEmail:   "email",
	user.OrderByEnabled: "enabled",
}

func orderByClause(orderBy order.By) (string, error) {
	by, exists := orderByFields[orderBy.Field]
	if !exists {
		return "", fmt.Errorf("field %q does not exist", orderBy.Field)
	}

	return " ORDER BY " + by + " " + orderBy.Direction, nil
}
