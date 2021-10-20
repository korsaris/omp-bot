package retention

/*var allRetention = map[int]Retention{
	{0, {RetentionID: 0}},
	{1, {RetentionID: 1}},
	{2, {RetentionID: 2}},
	{3, {RetentionID: 3}},
	{4, {RetentionID: 4}},
}*/
var allRetention = []Retention{}

type Retention struct {
	RetentionID uint64
}
