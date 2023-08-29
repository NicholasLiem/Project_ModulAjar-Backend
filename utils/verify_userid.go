package utils

import (
	"errors"
	"strconv"
)

func VerifyUserId(UserID string) (uint64, error) {
	userID, err := strconv.ParseUint(UserID, 10, 64)
	if err != nil {
		return 0, errors.New("cannot parse id")
	}
	return userID, nil
}
