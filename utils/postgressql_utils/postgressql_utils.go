package postgressql_utils

import (
	"strings"

	"github.com/jademnp/go-store-user-api/utils/errors"
	"github.com/lib/pq"
)

const (
	errorNoRow  = "no row"
	errorUnique = "unique_violation"
)

func ParseError(err error) *errors.RestErr {
	pqErr, ok := err.(*pq.Error)
	if !ok {
		if strings.Contains(err.Error(), errorNoRow) {
			return errors.NotFoundError("user not found")
		}
		return errors.InternalServerError("error parsing database response")
	}
	switch pqErr.Code.Name() {
	case errorUnique:
		return errors.BadRequestError("duplicate value found")
	}
	return errors.InternalServerError("error parsing database request")
}
