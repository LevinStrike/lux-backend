package v1

import (
	"encoding/json"
	"io"

	"github.com/LevinStrike/lux-backend/internal/core/apperror"
)

func JsonDecode(buf io.Reader, body any) error {
	if buf == nil {
		return apperror.ErrUnexpectedNillValues
	}
	if err := json.NewDecoder(buf).Decode(body); err != nil {
		return apperror.NewBadRequestError(err)
	}
	return nil
}
