package utils

import (
	"strconv"
	"time"
)


func ConvertEpochToHuman(convertEpoch string) string {
	epoch, err := strconv.ParseInt(convertEpoch, 10, 64)
	if err != nil {
		return "Erro ao converter epoch"
	}
	tempo := time.Unix(epoch, 0)
	return tempo.Format("02-Jan-2006 15:04:05")
}