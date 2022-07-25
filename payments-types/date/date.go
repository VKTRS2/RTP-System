package date

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/jinzhu/now"

	"github.com/domonda/go-types/language"
	"github.com/domonda/go-types/strutil"
)

// Normalize returns str as normalized Date or an error.
// The first given lang argument is used as language hint.
func Normalize(str string, lang ...language.Code) (Date, error) {
	return Date(str).Normalized(lang...)
}

// StringIsDate returns if a string can be parsed as Date.
// The first given lang argument is used as language hint.
func StringIsDate(str string, lang ...language.Code) bool {
	_, err := Normalize(str, lang...)
	return err == nil
}

const (
	// Layout used for the Date type, compatible with time.Time.Format()
	Layout = "2006-01-02"

	Length = 10 // len("2006-01-02")

	// MinLength is the minimum length of a valid date
	MinLength = 6

	// Invalid holds an empty string Date.
	// See NullableDate for where an empty string is a valid value.
	Invalid Date = ""
)

// Date represents a the day of calender date
// Date implements the database/sql.Scanner and database/sql/driver.Valuer interfaces,
// and will treat an empty string or the zero dates "0000-00-00" and "0001-01-01" (see IsZero) as SQL NULL.
type Date string

// Must returns str as normalized Date or panics if str is not a valid Date.
func Must(str string) Date {
	d, err := Date(str).Normalized()
	if err != nil {
		panic(err)
	}
	return d
}

// Of returns a normalized Date for the given year, month, and day.
// The month, day values may be outside
// their usual ranges and will be normalized during the conversion.
// For example, October 32 converts to November 1.
func Of(year int, month time.Month, day int) Date {
	return OfTime(time.Date(year, month, day, 0, 0, 0, 0, time.Local))
}

// OfTime returns the date part of the passed time.Time
// or an empty string if t.IsZero().
// To get the date in a certain time zone,
// pass the time.Time with a location set to the time zone.
func OfTime(t time.Time) Date {
	if t.IsZero() {
		return ""
	}
	return Date(t.Format(Layout))
}

// OfTimePtr returns the date part of the passed time.Time
// or Null (an empty string) if t is nil or t.IsZero().
func OfTimePtr(t *time.Time) NullableDate {
	if t == nil || t.IsZero() {
		return Null
	}
	return NullableDate(OfTime(*t))
}

// OfToday returns the date of today in the local timezone.
func OfToday() Date {
	return OfTime(time.Now())
}

// OfNowInUTC returns the date of the current time in the UTC timezone.
func OfNowInUTC() Date {
	return OfTime(time.Now().UTC())
}

// OfTodayUTC returns the date of today in the timezone of the passed location.
func OfTodayIn(loc *time.Location) Date {
	return OfTime(time.Now().In(loc))
}

// Parse returns the date part from time.Parse(layout, value)
func Parse(layout, value string) (Date, error) {
	t, err := time.Parse(layout, value)
	if t.IsZero() || err != nil {
		return "", err
	}
	return OfTime(t), nil
}

// PeriodRange returns the dates [from, until] for a period
// defined in one the following formats:
// period of a ISO 8601 week of a year: YYYY-Wnn
// period of a month of a year: YYYY-MM,
// period of a quarter of a year: YYYY-Qn,
// period of a half year: YYYY-Hn,
// period of full year: YYYY.
// The returned from Date is the first day of the month of the period,
// the returned until Date is the last day of the month of the period.
// Exmaples:
// Period of June 2018: PeriodRange("2018-06") == Date("2018-06-01"), Date("2018-06-30"), nil
// Period of Q3 2018: PeriodRange("2018-Q3") == Date("2018-07-01"), Date("2018-09-30"), nil
// Period of second half of 2018: PeriodRange("2018-H2") == Date("2018-07-01"), Date("2018-12-31"), nil
// Period of year 2018: PeriodRange("2018") == Date("2018-07-01"), Date("2018-12-31"), nil
// Period of week 1 2019: PeriodRange("2019-W01") == Date("2018-12-31"), Date("2019-01-06"), nil
func PeriodRange(period string) (from, until Date, err error) {
	if len(period) != 4 && len(period) != 7 && len(period) != 8 {
		return "", "", fmt.Errorf("invalid period format length: %q", period)
	}

	if len(period) == 4 {
		year, err := strconv.Atoi(period)
		if err != nil || year <= 0 {
			return "", "", fmt.Errorf("invalid period format: %q", period)
		}
		from = Date(period + "-01-01")
		until = Date(period + "-12-31")
		return from, until, nil
	}

	if period[4] != '-' {
		return "", "", fmt.Errorf("invalid period format, expected '-' after year: %q", period)
	}

	year, err := strconv.Atoi(period[:4])
	if err != nil {
		return "", "", fmt.Errorf("invalid period format, can't parse year: %q", period)
	}

	switch period[5] {
	case 'W', 'w':
		week, err := strconv.Atoi(period[6:])
		if err != nil || week < 1 || week > 53 {
			return "", "", fmt.Errorf("invalid period format, can't parse week: %q", period)
		}
		from, until = YearWeekRange(year, week)
		return from, until, nil

	case 'Q', 'q':
		quarter, err := strconv.Atoi(period[6:])
		if err != nil || quarter < 1 || quarter > 4 {
			return "", "", fmt.Errorf("invalid period format, can't parse quarter: %q", period)
		}
		from = Of(year, time.Month(quarter-1)*3+1, 1)
		until = Of(year, time.Month(quarter)*3+1, 0) // 0th day is the last day of the previous month
		return from, until, nil

	case 'H', 'h':
		half, err := strconv.Atoi(period[6:])
		if err != nil || half < 1 || half > 2 {
			return "", "", fmt.Errorf("invalid period format, can't parse half-year: %q", period)
		}
		from = Of(year, time.Month(half-1)*6+1, 1)
		until = Of(year, time.Month(half)*6+1, 0) // 0th day is the last day of the previous month
		return from, until, nil
	}

	month, err := strconv.Atoi(period[5:])
	if err != nil || month < 1 || month > 12 {
		return "", "", fmt.Errorf("invalid period format, can't parse month: %q", period)
	}

	from = Of(year, time.Month(month), 1)
	until = Of(year, time.Month(month)+1, 0) // 0th day is the last day of the previous month
	return from, until, nil
}

// YearRange returns the date range from
// first of January to 31st of December of a year.
func YearRange(year int) (from, until Date) {
	yyyy := fmt.Sprintf("%04d", year)
	return Date(yyyy + "-01-01"), Date(yyyy + "-12-31")
}

// YearWeekMonday returns the date of Monday of an ISO 8601 week.
func YearWeekMonday(year, week int) (monday Date) {
	// January 1st of the year
	t := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	// Go back to Monday of week
	t = t.AddDate(0, 0, int(time.Monday-t.Weekday()))

	// Add week days
	t = t.AddDate(0, 0, (week-1)*7)

	return OfTime(t)
}

// YearWeekRange returns the dates of Monday and Sunday of an ISO 8601 week.
func YearWeekRange(year, week int) (monday, sunday Date) {
	monday = YearWeekMonday(year, week)
	sunday = monday.AddDays(6)
	return monday, sunday
}

func FromUntilFromYearAndMonths(year, months string) (fromDate, untilDate Date, err error) {
	if year == "" {
		return "", "", nil
	}
	if months == "" {
		months = "1-12"
	}
	parts := strings.Split(months, "-")
	fromMonth := parts[0]
	untilMonth := parts[0]
	if len(parts) == 2 {
		untilMonth = parts[1]
	}

	yearInt, err := strconv.Atoi(year)
	if err != nil {
		return "", "", err
	}
	fromMonthInt, err := strconv.Atoi(fromMonth)
	if err != nil {
		return "", "", err
	}
	untilMonthInt, err := strconv.Atoi(untilMonth)
	if err != nil {
		return "", "", err
	}
	fromDate = Of(yearInt, time.Month(fromMonthInt), 1)
	untilDate = Of(yearInt, time.Month(untilMonthInt+1), 0) // 0th day is the last day of the previous month

	return fromDate, untilDate, nil
}

// ScanString tries to parse and assign the passed
// source string as value of the implementing type.
// It returns an error if source could not be parsed.
// If the source string could be parsed, but was not
// in the expected normalized format, then false is
// returned for sourceWasNormalized and nil for err.
// ScanString implements the strfmt.Scannable interface.
func (date *Date) ScanString(source string) (sourceWasNormalized bool, err error) {
	newDate, err := Date(source).Normalized()
	if err != nil {
		return false, err
	}
	*date = newDate
	return newDate == Date(source), nil
}

func (date *Date) ScanStringWithLang(source string, lang language.Code) (sourceWasNormalized bool, monthMustBeFirst bool, err error) {
	newDate, monthMustBeFirst, err := normalizeAndCheckDate(source, lang)
	if err != nil {
		return false, false, err
	}
	*date = newDate
	return newDate == Date(source), monthMustBeFirst, nil
}

// String returns the normalized date if possible,
// else it will be returned unchanged as string.
// String implements the fmt.Stringer interface.
func (date Date) String() string {
	norm, err := date.Normalized()
	if err != nil {
		return string(date)
	}
	return string(norm)
}

// WithinIncl returns if date is within and inclusive from and until.
func (date Date) WithinIncl(from, until Date) bool {
	t := date.MidnightUTC()
	tFrom := from.MidnightUTC()
	tUntil := until.MidnightUTC()
	return (t == tFrom || t.After(tFrom)) && (t == tUntil || t.Before(tUntil))
}

// BetweenExcl returns if date is between and exlusive after and until.
func (date Date) BetweenExcl(after, before Date) bool {
	t := date.MidnightUTC()
	return t.After(after.MidnightUTC()) && t.Before(before.MidnightUTC())
}

// Nullable returns the date as NullableDate
func (date Date) Nullable() NullableDate {
	if date.IsZero() {
		return Null
	}
	return NullableDate(date)
}

// Validate returns an error if the date is not in a valid, normalizeable format.
func (date Date) Validate() error {
	_, err := date.Normalized()
	return err
}

// Valid returns if the format of the date is correct, see Format
func (date Date) Valid() bool {
	return date.Validate() == nil
}

func (date Date) ValidAndNormalized() bool {
	_, err := time.Parse(Layout, string(date))
	return err == nil
}

func (date Date) Time(hour, minute, second int, location *time.Location) time.Time {
	if date.IsZero() {
		return time.Time{}
	}
	year, month, day := date.YearMonthDay()
	return time.Date(year, month, day, hour, minute, second, 0, location)
}

func (date Date) TimeLocal(hour, minute, second int) time.Time {
	return date.Time(hour, minute, second, time.Local)
}

func (date Date) TimeUTC(hour, minute, second int) time.Time {
	return date.Time(hour, minute, second, time.UTC)
}

// MidnightUTC returns the midnight (00:00) time.Time of the date in UTC,
// or a zero time.Time value if the date is not valid.
func (date Date) MidnightUTC() time.Time {
	if date.IsZero() {
		return time.Time{}
	}
	t, err := time.Parse(Layout, string(date))
	if err != nil {
		return time.Time{}
	}
	return t
}

// Midnight returns the midnight (00:00) time.Time of the date
// in the local time zone,
// or a zero time.Time value if the date is not valid.
func (date Date) Midnight() time.Time {
	return date.MidnightInLocation(time.Local)
}

// MidnightInLocation returns the midnight (00:00) time.Time of the date
// in the given location,
// or a zero time.Time value if the date is not valid.
func (date Date) MidnightInLocation(loc *time.Location) time.Time {
	if date.IsZero() {
		return time.Time{}
	}
	t, err := time.ParseInLocation(Layout, string(date), loc)
	if err != nil {
		return time.Time{}
	}
	return t
}

// Format returns date.MidnightUTC().Format(layout),
// or an empty string if date or layout are an empty string.
func (date Date) Format(layout string) string {
	if date == "" || layout == "" {
		return ""
	}
	if layout == Layout {
		return string(date)
	}
	return date.MidnightUTC().Format(layout)
}

func (date Date) NormalizedOrUnchanged(lang ...language.Code) Date {
	normalized, err := date.Normalized(lang...)
	if err != nil {
		return date
	}
	return normalized
}

// func (date Date) NormalizedOrInvalid(lang ...language.Code) Date {
// 	normalized, err := date.Normalized(lang...)
// 	if err != nil {
// 		return Invalid
// 	}
// 	return normalized
// }

func (date Date) NormalizedOrNull(lang ...language.Code) NullableDate {
	normalized, err := date.Normalized(lang...)
	if err != nil {
		return Null
	}
	return NullableDate(normalized)
}

// NormalizedEqual returns if two dates are equal in normalized form.
func (date Date) NormalizedEqual(other Date) bool {
	a, _ := date.Normalized()
	b, _ := other.Normalized()
	return a == b
}

// After returns if the date is after the passed other one.
func (date Date) After(other Date) bool {
	return date.MidnightUTC().After(other.MidnightUTC())
}

// EqualOrAfter returns if the date is equal or after the passed other one.
func (date Date) EqualOrAfter(other Date) bool {
	return date.NormalizedEqual(other) || date.After(other)
}

// Before returns if the date is before the passed other one.
func (date Date) Before(other Date) bool {
	return date.MidnightUTC().Before(other.MidnightUTC())
}

// EqualOrBefore returns if the date is equal or before the passed other one.
func (date Date) EqualOrBefore(other Date) bool {
	return date.NormalizedEqual(other) || date.Before(other)
}

// AfterTime returns if midnight of the date in location of the passed
// time is after the time.
func (date Date) AfterTime(other time.Time) bool {
	return date.MidnightInLocation(other.Location()).After(other)
}

// BeforeTime returns if midnight of the date in location of the passed
// time is before the time.
func (date Date) BeforeTime(other time.Time) bool {
	return date.MidnightInLocation(other.Location()).Before(other)
}

func (date Date) AddDate(years int, months int, days int) Date {
	return OfTime(date.MidnightUTC().AddDate(years, months, days))
}

func (date Date) BeginningOfWeek() Date {
	n := (now.Now{Time: date.MidnightUTC()})
	return OfTime(n.BeginningOfWeek())
}

func (date Date) BeginningOfMonth() Date {
	n := (now.Now{Time: date.MidnightUTC()})
	return OfTime(n.BeginningOfMonth())
}

func (date Date) BeginningOfQuarter() Date {
	n := (now.Now{Time: date.MidnightUTC()})
	return OfTime(n.BeginningOfQuarter())
}

func (date Date) BeginningOfYear() Date {
	n := (now.Now{Time: date.MidnightUTC()})
	return OfTime(n.BeginningOfYear())
}

func (date Date) EndOfWeek() Date {
	n := (now.Now{Time: date.MidnightUTC()})
	return OfTime(n.EndOfWeek())
}

func (date Date) EndOfMonth() Date {
	n := (now.Now{Time: date.MidnightUTC()})
	return OfTime(n.EndOfMonth())
}

func (date Date) EndOfQuarter() Date {
	n := (now.Now{Time: date.MidnightUTC()})
	return OfTime(n.EndOfQuarter())
}

func (date Date) EndOfYear() Date {
	n := (now.Now{Time: date.MidnightUTC()})
	return OfTime(n.EndOfYear())
}

func (date Date) LastMonday() Date {
	n := (now.Now{Time: date.MidnightUTC()})
	return OfTime(n.Monday())
}

func (date Date) NextSunday() Date {
	n := (now.Now{Time: date.MidnightUTC()})
	return OfTime(n.Sunday())
}

// YearMonthDay returns the year, month, day components of the Date.
// Zero values will be returned when the date is not valid.
func (date Date) YearMonthDay() (year int, month time.Month, day int) {
	if len(date) != Length {
		return 0, 0, 0
	}
	year, _ = strconv.Atoi(string(date)[:4])
	monthInt, _ := strconv.Atoi(string(date)[5:7])
	day, _ = strconv.Atoi(string(date)[8:])
	return year, time.Month(monthInt), day
}

// ISOWeek returns the ISO 8601 year and week number in which the date occurs.
// Week ranges from 1 to 53. Jan 01 to Jan 03 of year n might belong to
// week 52 or 53 of year n-1, and Dec 29 to Dec 31 might belong to week 1
// of year n+1.
func (date Date) ISOWeek() (year, week int) {
	t := date.MidnightUTC()
	if t.IsZero() {
		return 0, 0
	}
	return t.ISOWeek()
}

// Scan implements the database/sql.Scanner interface.
func (date *Date) Scan(value interface{}) (err error) {
	switch x := value.(type) {
	case string:
		d := Date(x)
		if !d.IsZero() {
			d, err = d.Normalized()
			if err != nil {
				return err
			}
		}
		*date = d
		return nil

	case time.Time:
		*date = OfTime(x)
		return nil

	case nil:
		*date = ""
		return nil
	}

	return fmt.Errorf("can't scan value '%#v' of type %T as data.Date", value, value)
}

// Value implements the driver database/sql/driver.Valuer interface.
func (date Date) Value() (driver.Value, error) {
	if date.IsZero() {
		return nil, nil
	}
	normalized, err := date.Normalized()
	if err != nil {
		return nil, err
	}
	return string(normalized), nil
}

// IsZero returns true when the date is any of ["", "0000-00-00", "0001-01-01"]
// "0001-01-01" is treated as zero because it's the zero value of time.Time.
// "0000-00-00" may be the zero value of other date implementations.
func (date Date) IsZero() bool {
	return date == "" || date == "0000-00-00" || date == "0001-01-01"
}

func (date Date) IsToday() bool {
	return date == OfToday()
}

func (date Date) IsTodayInUTC() bool {
	return date == OfNowInUTC()
}

func (date Date) AfterToday() bool {
	return date.After(OfToday())
}

func (date Date) AfterTodayInUTC() bool {
	return date.After(OfNowInUTC())
}

func (date Date) BeforeToday() bool {
	return date.Before(OfToday())
}

func (date Date) BeforeTodayInUTC() bool {
	return date.Before(OfNowInUTC())
}

func (date Date) Add(years, months, days int) Date {
	return OfTime(date.Midnight().AddDate(years, months, days))
}

func (date Date) AddYears(years int) Date {
	return OfTime(date.Midnight().AddDate(years, 0, 0))
}

func (date Date) AddMonths(months int) Date {
	return OfTime(date.Midnight().AddDate(0, months, 0))
}

func (date Date) AddDays(days int) Date {
	return OfTime(date.Midnight().AddDate(0, 0, days))
}

func isDateSeparatorRune(r rune) bool {
	return unicode.IsSpace(r) || r == '.' || r == '/' || r == '-'
}

func isDateTrimRune(r rune) bool {
	return unicode.IsSpace(r) || (!unicode.IsLetter(r) && !unicode.IsDigit(r))
}

// Normalized returns the date in normalized form,
// or an error if the format can't be detected.
// The first given lang argument is used as language hint.
func (date Date) Normalized(lang ...language.Code) (Date, error) {
	normalized, _, err := normalizeAndCheckDate(string(date), getLangHint(lang))
	return normalized, err
}

func normalizeAndCheckDate(str string, langHint language.Code) (Date, bool, error) {
	normalized, monthMustBeFirst, err := normalizeDate(str, langHint)
	if err != nil {
		return "", monthMustBeFirst, err
	}
	_, err = time.Parse(Layout, normalized)
	if err != nil {
		return "", monthMustBeFirst, err
	}
	return Date(normalized), monthMustBeFirst, nil
}

func normalizeDate(str string, langHint language.Code) (string, bool, error) {
	trimmed := strings.ToLower(strings.TrimFunc(str, isDateTrimRune))

	if len(trimmed) < MinLength {
		return "", false, fmt.Errorf("too short for a date: %q", str)
	}

	if len(trimmed) > 10 && trimmed[10] == 't' {
		// Use date part of this date-time format: "2006-01-02T15:04:05"
		trimmed = trimmed[:10]
	}

	langHint, _ = langHint.Normalized()

	parts := strings.FieldsFunc(trimmed, isDateSeparatorRune)
	if len(parts) == 4 {
		i := strutil.IndexInStrings("of", parts)
		if i > 0 && i <= 2 {
			// remove the word "of" within date
			parts = append(parts[:i], parts[i+1:]...)
		}
	}
	if len(parts) != 3 {
		return "", false, fmt.Errorf("date must have 3 parts: %q", str)
	}
	dayHint := -1
	totalLen := 0
	for i := range parts {
		l := len(parts[i])
		totalLen += l
		if l == 1 {
			parts[i] = "0" + parts[i]
		} else if parts[i] == "1st" {
			parts[i] = "01"
			dayHint = i
		} else if parts[i] == "2nd" {
			parts[i] = "02"
			dayHint = i
		} else if parts[i] == "3rd" {
			parts[i] = "03"
			dayHint = i
		} else if strings.HasSuffix(parts[i], "th") {
			parts[i] = strings.TrimSuffix(parts[i], "th")
			if len(parts[i]) == 1 {
				parts[i] = "0" + parts[i]
			}
			dayHint = i
		}
	}
	if totalLen < 5 {
		return "", false, fmt.Errorf("date is too short: %q", str)
	}

	len0 := len(parts[0])
	len1 := len(parts[1])
	len2 := len(parts[2])
	val0, _ := strconv.Atoi(parts[0])
	val1, _ := strconv.Atoi(parts[1])
	val2, _ := strconv.Atoi(parts[2])
	month0, _ := monthNameMap[parts[0]]
	month1, _ := monthNameMap[parts[1]]
	month2, _ := monthNameMap[parts[2]]

	// fmt.Println(len0, len1, len2)
	// fmt.Println(val0, val1, val2)

	expandVal2ToFullYear := func() {
		if len2 != 2 {
			panic("len2")
		}
		if val2 < 45 {
			parts[2] = "20" + parts[2]
			val2 = 2000 + val2
		} else {
			parts[2] = "19" + parts[2]
			val2 = 1900 + val2
		}
		len2 = 4
	}

	switch {
	case month0 != 0:
		if len2 == 2 {
			expandVal2ToFullYear()
		}
		if !validDay(val1) || !validYear(val2) {
			return "", false, fmt.Errorf("invalid date: %q", str)
		}
		// m DD YYYY
		return fmt.Sprintf("%s-%02d-%s", parts[2], month0, parts[1]), false, nil

	case len0 == 2 && month1 != 0 && len2 == 2:
		if (!validDay(val0) && validDay(val2)) || dayHint == 2 {
			// YY m DD
			parts[0], parts[2] = parts[2], parts[0]
			val0, val2 = val2, val0
			// DD m YY
		}
		expandVal2ToFullYear()
		fallthrough

	case len0 == 2 && month1 != 0 && len2 == 4:
		if !validDay(val0) || !validYear(val2) {
			return "", false, fmt.Errorf("invalid date: %q", str)
		}
		// DD m YYYY
		return fmt.Sprintf("%s-%02d-%s", parts[2], month1, parts[0]), false, nil

	case len0 == 4 && month1 != 0 && len2 == 2:
		if !validYear(val0) || !validDay(val2) {
			return "", false, fmt.Errorf("invalid date: %q", str)
		}
		// YYYY m DD
		return fmt.Sprintf("%s-%02d-%s", parts[0], month1, parts[2]), false, nil

	case month2 != 0:
		if len0 == 2 {
			// YY DD m
			if val0 < 45 {
				parts[0] = "20" + parts[0]
				val0 = 2000 + val0
			} else {
				parts[0] = "19" + parts[0]
				val0 = 1900 + val0
			}
			len0 = 4
		}
		// YYYY DD m
		return fmt.Sprintf("%s-%02d-%s", parts[0], month2, parts[1]), false, nil

	case len0 == 4 && len1 == 2 && len2 == 2:
		if !validYear(val0) || !validMonth(val1) || !validDay(val2) {
			return "", false, fmt.Errorf("invalid date: %q", str)
		}
		return strings.Join(parts, "-"), false, nil

	case len0 == 2 && len1 == 2 && len2 == 2:
		expandVal2ToFullYear()
		fallthrough

	case len0 == 2 && len1 == 2 && len2 == 4:
		monthMustBeFirst := validMonth(val0) && !validMonth(val1)
		if (!validMonth(val1) && validMonth(val0)) || dayHint == 1 || langHint == "en" {
			// MM DD YYYY
			parts[0], parts[1] = parts[1], parts[0]
			val0, val1 = val1, val0
			// DD MM YYYY
		}
		if !validDay(val0) || !validMonth(val1) || !validYear(val2) {
			return "", false, fmt.Errorf("invalid date: %q", str)
		}
		// DD MM YYYY
		parts[0], parts[2] = parts[2], parts[0]
		// YYYY MM DD
		return strings.Join(parts, "-"), monthMustBeFirst, nil
	}

	return "", false, fmt.Errorf("invalid date: %q", str)
}

func validYear(year int) bool {
	return year > 0
}

func validMonth(month int) bool {
	return month >= 1 && month <= 12
}

func validDay(day int) bool {
	return day >= 1 && day <= 31
}

func getLangHint(lang []language.Code) language.Code {
	if len(lang) == 0 || len(lang[0]) < 2 {
		return ""
	}
	return lang[0][:2]
}
