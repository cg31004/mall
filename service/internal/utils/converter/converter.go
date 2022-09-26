package converter

import (
	"strconv"
	"time"

	"golang.org/x/xerrors"
)

func ConvertIntToIntRef(int int) *int {
	return &int
}

func ConvertInt64ToStr(int64 int64) string {
	return strconv.FormatInt(int64, 10)
}

func ConvertStrToStrRef(str string) *string {
	if str == "" {
		return nil
	}
	return &str
}

func ConvertStrRefToStr(strRef *string) string {
	if strRef == nil {
		return ""
	}

	return *strRef
}

func ConvertStrToInt(intStr string) (int, error) {
	targetInt64, err := strconv.ParseInt(intStr, 10, 0)
	if err != nil {
		return 0, xerrors.Errorf("無法將字串 %s 轉換成 int: %w", intStr, err)
	}
	return int(targetInt64), nil
}

func ConvertStrToIntRef(intStr string) (*int, error) {
	if intStr == "" {
		return nil, nil
	}

	targetInt64, err := strconv.ParseInt(intStr, 10, 0)
	if err != nil {
		return nil, xerrors.Errorf("無法將字串 %s 轉換成 *int: %w", intStr, err)
	}

	targetInt := int(targetInt64)

	return &targetInt, nil
}

func ConvertStrToInt64(intStr string) (int64, error) {
	targetInt64, err := strconv.ParseInt(intStr, 10, 64)
	if err != nil {
		return 0, xerrors.Errorf("無法將字串 %s 轉換成 int64: %w", intStr, err)
	}
	return targetInt64, nil
}

func ConvertStrToTime(timeStr string) (time.Time, error) {
	targetTime, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		return time.Time{}, xerrors.Errorf("無法將字串 %s 轉換成 time.Time: %w", timeStr, err)
	}

	return targetTime, nil
}

func ConvertStrToTimeRef(timeStr string) (*time.Time, error) {
	if timeStr == "" {
		return nil, nil
	}

	targetTime, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		return nil, xerrors.Errorf("無法將字串 %s 轉換成 *time.Time: %w", timeStr, err)
	}

	return &targetTime, nil
}

func ConvertStrToBool(boolStr string) bool {
	return boolStr == "1"
}

func CovertIntSliceToStringSlice(intSlice []int) []string {
	stringSlice := make([]string, 0, len(intSlice))
	for _, i := range intSlice {
		s := strconv.Itoa(i)

		stringSlice = append(stringSlice, s)
	}
	return stringSlice
}
