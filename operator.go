package songo

func VerifyQueryOperator(operator string) bool {
	switch operator {
	case "$and", "$or":
	case "$eq", "$ne":
	case "$lt", "$lte", "$gt", "$gte":
	case "$like", "$in", "$nin", "$bt", "$nbt":
	default:
		return false
	}
	return true
}

func IsQueryOperatorGroup(operator string) bool {
	switch operator {
	case "$and", "$or":
	default:
		return false
	}
	return true
}

func IsQueryOperatorV(operator string) bool {
	switch operator {
	case "$eq", "$ne":
	default:
		return false
	}
	return true
}

func IsQueryOperatorKV(operator string) bool {
	switch operator {
	case "$lt", "$lte", "$gt", "$gte":
	case "$like", "$in", "$nin", "$bt", "$nbt":
	default:
		return false
	}
	return true
}
