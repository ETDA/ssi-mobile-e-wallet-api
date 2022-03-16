package helpers

import (
	"fmt"
	"gitlab.finema.co/finema/etda/mobile-app-api/emsgs"
	"strconv"
	"strings"
)

func BirthdateTransform(date string) (string, error) {
	splited := strings.Split(date, " ")
	if len(splited) != 3 {
		return "", emsgs.InvalidDateFormat
	}
	months := map[string]string{
		"ม.ค.":  "01",
		"ก.พ.":  "02",
		"มี.ค.": "03",
		"เม.ย.": "04",
		"พ.ค.":  "05",
		"มิ.ย.": "06",
		"ก.ค.":  "07",
		"ส.ค.":  "08",
		"ก.ย.":  "09",
		"ต.ค.":  "10",
		"พ.ย.":  "11",
		"ธ.ค.":  "12",
	}
	month, ok := months[splited[1]]
	if !ok {
		return "", emsgs.InvalidDateFormat
	}
	day, err := strconv.Atoi(splited[0])
	if err != nil {
		return "", emsgs.InvalidDateFormat
	}
	formatted := splited[2] + month + fmt.Sprintf("%02d", day)
	return formatted, nil
}
