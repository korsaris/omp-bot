package retention

import "fmt"

var allRetention = []Retention{
	{RetentionID: 0},
	{RetentionID: 1},
	{RetentionID: 2},
	{RetentionID: 3},
	{RetentionID: 4},
}

type Retention struct {
	RetentionID uint64
}

func NewRetention(retentionID uint64) *Retention {
	return &Retention{
		RetentionID: retentionID,
	}
}

func (r *Retention) String() string {
	return fmt.Sprintf("{ RetentionID: %v }", r.RetentionID)
}
