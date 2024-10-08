package convertthai

import (
	"fmt"
	"strings"
	"time"
)

var fullMonths = map[string]string{
	"January":   "มกราคม",
	"February":  "กุมภาพันธ์",
	"March":     "มีนาคม",
	"April":     "เมษายน",
	"May":       "พฤษภาคม",
	"June":      "มิถุนายน",
	"July":      "กรกฎาคม",
	"August":    "สิงหาคม",
	"September": "กันยายน",
	"October":   "ตุลาคม",
	"November":  "พฤศจิกายน",
	"December":  "ธันวาคม",
}

var shortMonths = map[string]string{
	"January":   "ม.ค.",
	"February":  "ก.พ.",
	"March":     "มี.ค.",
	"April":     "เม.ย.",
	"May":       "พ.ค.",
	"June":      "มิ.ย.",
	"July":      "ก.ค.",
	"August":    "ส.ค.",
	"September": "ก.ย.",
	"October":   "ต.ค.",
	"November":  "พ.ย.",
	"December":  "ธ.ค.",
}

var fullDays = map[string]string{
	"Sunday":    "อาทิตย์",
	"Monday":    "จันทร์",
	"Tuesday":   "อังคาร",
	"Wednesday": "พุธ",
	"Thursday":  "พฤหัสบดี",
	"Friday":    "ศุกร์",
	"Saturday":  "เสาร์",
}

var shortDays = map[string]string{
	"Sunday": "อา", "Monday": "จ", "Tuesday": "อ", "Wednesday": "พ", "Thursday": "พฤ", "Friday": "ศ", "Saturday": "ส",
}

func addZero(numberStr string) string {
	if len(numberStr) == 3 {
		return "๐" + numberStr
	}

	return numberStr
}

func getSeparator(format string) string {
	if strings.Contains(format, "/") {
		return "/"
	} else if strings.Contains(format, "-") {
		return "-"
	} else if strings.Contains(format, " ") {
		return " "
	}
	return ""
}

func DateFormat(date interface{}, format string) string {
	var newDate time.Time
	var err error

	switch v := date.(type) {
	case time.Time:
		newDate = v
	case string:
		layout := "2006-01-02"
		newDate, err = time.Parse(layout, v)
		if err != nil {
			fmt.Println("Error parsing date:", err)
			return ""
		}
	default:
		fmt.Println("Unsupported date type")
		return ""
	}

	day := newDate.Day()
	weekday := newDate.Weekday().String()
	month := newDate.Month().String()
	monthInt := int(newDate.Month())
	year := newDate.Year()
	hour := newDate.Hour()
	minute := newDate.Minute()
	second := newDate.Second()

	if format == "" {
		format = "dd/mm/yyyy"
	}

	formatDate := map[string]string{
		"dddd": fullDays[weekday],
		"ddd":  shortDays[weekday],
		"dd":   addZero(NumberFormat(float64(day), 0, false)),
		"d":    NumberFormat(float64(day), 0, false),
		"mmmm": fullMonths[month],
		"mmm":  shortMonths[month],
		"mm":   addZero(NumberFormat(float64(monthInt), 0, false)),
		"m":    NumberFormat(float64(monthInt), 0, false),
		"yyyy": NumberFormat(float64(year), 0, false),
		"HH":   addZero(NumberFormat(float64(hour), 0, false)),
		"H":    NumberFormat(float64(hour), 0, false),
		"MM":   addZero(NumberFormat(float64(minute), 0, false)),
		"M":    NumberFormat(float64(minute), 0, false),
		"SS":   addZero(NumberFormat(float64(second), 0, false)),
		"S":    NumberFormat(float64(second), 0, false),
	}

	separator := getSeparator(format)
	formatArr := strings.Split(format, separator)

	var result string
	for i, key := range formatArr {
		if val, exists := formatDate[key]; exists {
			result += val
		}
		if i < len(formatArr)-1 {
			result += separator
		}
	}

	return result
}
