package songo

func VerifyQueryOperator(operator string) bool {
	switch operator {
	case "$and", "$or":
	case "$eq", "$ne":
	case "$lt", "$lte", "$gt", "$gte":
	case "$like", "$in", "$bt", "$nbt":
	default:
		return false
	}
	return true
}
