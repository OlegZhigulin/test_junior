package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func GetDaysLeft(ctx echo.Context) error {
	targetDate := ctx.QueryParam("date")
	if targetDate == "" {
		return ctx.String(http.StatusBadRequest, "Параметр date обязателен в формате YYYY-MM-DD")
	}

	target, _ := time.Parse("2006-01-02", targetDate)
	currentTimeUntil := time.Until(target)
	daysLeft := int(currentTimeUntil.Hours() / 24)
	hoursLeft := int(currentTimeUntil.Hours()) % 24
	return ctx.String(http.StatusOK,
		fmt.Sprintf("До %s осталось %d дней и %d часов",
			targetDate, daysLeft, hoursLeft))
}
