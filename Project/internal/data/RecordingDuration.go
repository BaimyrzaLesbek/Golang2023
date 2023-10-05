package data

import (
	"fmt"
	"strconv"
)

type RecordingDuration int64

func (r RecordingDuration) MarshalJSON() ([]byte, error) {

	jsonValue := fmt.Sprintf("%d mins", r)

	quotedJSONValue := strconv.Quote(jsonValue)

	return []byte(quotedJSONValue), nil
}
