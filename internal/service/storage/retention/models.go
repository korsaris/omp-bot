package retention

/*var allRetention = map[int]Retention{
	{0, {retentionID: 0}},
	{1, {retentionID: 1}},
	{2, {retentionID: 2}},
	{3, {retentionID: 3}},
	{4, {retentionID: 4}},
}*/
var allRetention = []Retention {}


type Retention struct {
	retentionID uint64
}
