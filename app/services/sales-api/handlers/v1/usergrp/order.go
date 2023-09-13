package usergrp

import (
	"errors"
	"net/http"

	"github.com/vim-diesel/new-service/business/core/user"
	"github.com/vim-diesel/new-service/business/cview/user/summary"

	// "github.com/vim-diesel/new-service/business/cview/user/summary"
	"github.com/vim-diesel/new-service/business/data/order"
	"github.com/vim-diesel/new-service/foundation/validate"
)

var orderByFields = map[string]struct{}{
	user.OrderByID:      {},
	user.OrderByName:    {},
	user.OrderByEmail:   {},
	user.OrderByEnabled: {},
}

func parseOrder(r *http.Request) (order.By, error) {
	orderBy, err := order.Parse(r, user.DefaultOrderBy)
	if err != nil {
		return order.By{}, err
	}

	if _, exists := orderByFields[orderBy.Field]; !exists {
		return order.By{}, validate.NewFieldsError(orderBy.Field, errors.New("order field does not exist"))
	}

	return orderBy, nil
}

// =============================================================================

var orderBySummaryFields = map[string]struct{}{
	summary.OrderByUserID:   {},
	summary.OrderByUserName: {},
}

func parseSummaryOrder(r *http.Request) (order.By, error) {
	orderBy, err := order.Parse(r, user.DefaultOrderBy)
	if err != nil {
		return order.By{}, err
	}

	if _, exists := orderBySummaryFields[orderBy.Field]; !exists {
		return order.By{}, validate.NewFieldsError(orderBy.Field, errors.New("order field does not exist"))
	}

	return orderBy, nil
}
