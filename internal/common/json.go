package common

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
)

func ReadJSON(c echo.Context, obj any) error {
	reqBody := c.Request().Body
	dec := json.NewDecoder(reqBody)
	dec.DisallowUnknownFields()
	return dec.Decode(obj)
}
