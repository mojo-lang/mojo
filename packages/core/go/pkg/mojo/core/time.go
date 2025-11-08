package core

import "strings"

// time format
//
//	Year:
//		YYYY: 	"2006"
//		YY  :	"06"
//	Month:
//		MM	:	"01"
//		MS	:   "_1"  // the underscore in "_1" represent space
//		M	:	"1"
// 		Mon	:	"Jan"
//		Month:	"January"
//	Day of the week:
//		Wk	:	"Mon"
//		Week:	"Monday"
//	Day of the month:
//		D	:	"2"
//		DS	:	"_2"  // the underscore in "_2" represent space
//		DD	:	"02"
//	Hour:
//		hh	:	"15"
//		h	:	"3"
//		hh PM:	"03" (PM or AM)
//	Minute:
//		m	:	"4"
//		mm	:	"04"
//	Second:
//		s	:	"5"
//		ss	:	"05"
//		s.milli:	"5.999"
//		s.micro:	"5.999999"
//		s.nano:		"5.999999999"
// 		ss.milli:	"05.999"
//		ss.micro:	"05.999999"
//		ss.nano:	"05.999999999"
//	AM/PM mark:"PM"
//
// Numeric time zone offsets format as follows:
//
// -hh			"-07"
// -hhmm 		"-0700"
// -hh:mm 		"-07:00"
// -hhmmss		"-070000"
// -hh:mm:ss 	"-07:00:00"
//
// Zhh			"Z07"
// Zhhmm		"Z0700"
// Zhh:mm		"Z07:00"
// Zhhmmss		"Z070000"
// Zhh:mm:ss	"Z07:00:00"
//

const (
	RFC3339     = "YYYY-MM-DDThh:mm:ssZhh:mm"      // "2006-01-02T15:04:05Z07:00"
	RFC3339Nano = "YYYY-MM-DDThh:mm:ss.nanoZhh:mm" // "2006-01-02T15:04:05.999999999Z07:00"
	RFC850      = "Week, DD-Mon-YY hh:mm:ss MST"   // "Monday, 02-Jan-06 15:04:05 MST"
	Kitchen     = "h:mmPM"                         // "3:04PM"
)

func TimeFormat2Layout(format string) string {
	layout := &strings.Builder{}
	year, month, day, hour, minute, second := false, false, false, false, false, false
	for i := 0; i < len(format); {
		switch c := int(format[i]); c {
		case 'Y':
			if !year && len(format) >= i+2 && format[i:i+2] == "YY" {
				if len(format) >= i+4 && format[i:i+4] == "YYYY" {
					layout.WriteString("2006")
					i += 4
				} else {
					layout.WriteString("06")
					i += 2
				}
				year = true
			} else {
				layout.WriteByte(format[i])
				i += 1
			}
		case 'M':
			if !month {
				if len(format) >= i+3 && format[i:i+3] == "Mon" {
					if len(format) >= i+5 && format[i:i+5] == "Month" {
						layout.WriteString("January")
						i += 5
					} else {
						layout.WriteString("Jan")
						i += 3
					}
				} else if len(format) >= i+2 && format[i:i+2] == "MM" {
					layout.WriteString("01")
					i += 2
				} else if len(format) >= i+2 && format[i:i+2] == "MS" {
					layout.WriteString("_1")
					i += 2
				} else {
					layout.WriteString("1")
					i += 1
				}
				month = true
			} else {
				layout.WriteByte(format[i])
				i += 1
			}
		case 'W':
			if len(format) >= i+2 && format[i:i+2] == "Wk" {
				layout.WriteString("Mon")
				i += 2
			} else if len(format) >= i+4 && format[i:i+4] == "Week" {
				layout.WriteString("Monday")
				i += 4
			} else {
				layout.WriteByte(format[i])
				i += 1
			}
		case 'D':
			if !day {
				if len(format) >= i+2 && format[i:i+2] == "DD" {
					layout.WriteString("02")
					i += 2
				} else if len(format) >= i+2 && format[i:i+2] == "DS" {
					layout.WriteString("_2")
					i += 2
				} else {
					layout.WriteString("2")
					i += 1
				}
				day = true
			} else {
				layout.WriteByte(format[i])
				i += 1
			}
		case 'P':
			if len(format) >= i+2 && format[i:i+2] == "PM" {
				layout.WriteString(format[i : i+2])
				i += 2
			}
		case 'h':
			if !hour {
				pm := strings.Contains(format, "PM")
				if len(format) >= i+2 && format[i:i+2] == "hh" {
					if pm {
						layout.WriteString("03")
					} else {
						layout.WriteString("15")
					}
					i += 2
				} else {
					layout.WriteString("3")
					i += 1
				}
				hour = true
			} else {
				layout.WriteByte(format[i])
				i += 1
			}
		case 'm':
			if !minute {
				if len(format) >= i+2 && format[i:i+2] == "mm" {
					layout.WriteString("04")
					i += 2
				} else {
					layout.WriteString("4")
					i += 1
				}
				minute = true
			} else {
				layout.WriteByte(format[i])
				i += 1
			}
		case 's':
			if !second {
				if len(format) >= i+2 && (format[i:i+2] == "ss" || format[i:i+2] == "s.") {
					if format[i:i+2] == "ss" {
						if len(format) > i+7 && format[i:i+7] == "ss.nano" {
							layout.WriteString("05.999999999")
							i += 7
						} else if len(format) > i+8 && format[i:i+8] == "ss.milli" {
							layout.WriteString("05.999")
							i += 8
						} else if len(format) > i+8 && format[i:i+8] == "ss.micro" {
							layout.WriteString("05.999999")
							i += 8
						} else {
							layout.WriteString("05")
							i += 2
						}
					} else {
						if len(format) > i+6 && format[i:i+6] == "s.nano" {
							layout.WriteString("5.999999999")
							i += 7
						} else if len(format) > i+7 && format[i:i+7] == "s.milli" {
							layout.WriteString("5.999")
							i += 8
						} else if len(format) > i+7 && format[i:i+7] == "s.micro" {
							layout.WriteString("5.999999")
							i += 8
						} else {
							layout.WriteString("5")
							i += 1
						}
					}
				} else {
					layout.WriteString("5")
					i += 1
				}
				second = true
			} else {
				layout.WriteByte(format[i])
				i += 1
			}
		case '-', 'Z':
			if len(format) >= i+3 && format[i+1:i+3] == "hh" {
				if len(format) >= i+5 && format[i+1:i+5] == "hhmm" {
					if len(format) >= i+7 && format[i+1:i+7] == "hhmmss" {
						layout.WriteByte(format[i])
						layout.WriteString("070000")
						i += 7
					} else {
						layout.WriteByte(format[i])
						layout.WriteString("0700")
						i += 5
					}
				} else if len(format) >= i+6 && format[i+1:i+6] == "hh:mm" {
					if len(format) >= i+9 && format[i+1:i+9] == "hh:mm:ss" {
						layout.WriteByte(format[i])
						layout.WriteString("07:00:00")
						i += 9
					} else {
						layout.WriteByte(format[i])
						layout.WriteString("07:00")
						i += 6
					}
				} else {
					layout.WriteByte(format[i])
					layout.WriteString("07")
					i += 3
				}
			} else {
				layout.WriteByte(format[i])
				i += 1
			}
		default:
			layout.WriteByte(format[i])
			i += 1
		}
	}
	return layout.String()
}
