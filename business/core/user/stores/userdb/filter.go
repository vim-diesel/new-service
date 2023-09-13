package userdb

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/vim-diesel/new-service/business/core/user"
)

func (s *Store) applyFilter(filter user.QueryFilter, data map[string]interface{}, buf *bytes.Buffer) {
	var wc []string

	if filter.ID != nil {
		data["user_id"] = *filter.ID
		wc = append(wc, "user_id = :user_id")
	}

	if filter.FullName != nil {
		data["full_name"] = fmt.Sprintf("%%%s%%", *filter.FullName)
		wc = append(wc, "name LIKE :full_name")
	}

	if filter.Email != nil {
		data["email"] = (*filter.Email).String()
		wc = append(wc, "email = :email")
	}

	if filter.StartCreatedDate != nil {
		data["start_date_created"] = *filter.StartCreatedDate
		wc = append(wc, "date_created >= :start_date_created")
	}

	if filter.EndCreatedDate != nil {
		data["end_date_created"] = *filter.EndCreatedDate
		wc = append(wc, "date_created <= :end_date_created")
	}

	if len(wc) > 0 {
		buf.WriteString(" WHERE ")
		buf.WriteString(strings.Join(wc, " AND "))
	}
}
