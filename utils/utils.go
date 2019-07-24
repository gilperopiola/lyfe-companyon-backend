package utils

import (
	"database/sql"
	"strconv"
)

func ToString(i int) string {
	return strconv.Itoa(i)
}

func ToInt(s string) int {
	val, _ := strconv.Atoi(s)
	return val
}

func BoolToString(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

func GetEntryID(result sql.Result) int {
	id, _ := result.LastInsertId()
	return int(id)
}
