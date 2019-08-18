package utils

import (
	"crypto/sha1"
	"database/sql"
	"encoding/base64"
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

func Hash(salt string, data string) string {
	hasher := sha1.New()
	hasher.Write([]byte(salt + data))
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}
