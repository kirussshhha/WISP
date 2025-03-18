package domain

type TimeDifferenceRequest struct {
	FromTime string `json:"from_time" binding:"required"`
	ToTime   string `json:"to_time" binding:"required"`
}
