package wxXunHuan

import (
	"errors"
	"math"
	"strings"
	"time"
)

// Represents the number of elements in a given period
const (
	secondsPerMinute  = 60
	minutesPerHour    = 60
	hoursPerDay       = 24
	daysPerWeek       = 7
	monthsPerQuarter  = 3
	monthsPerYear     = 12
	yearsPerCenturies = 100
	yearsPerDecade    = 10
	weeksPerLongYear  = 53
	daysInLeapYear    = 366
	daysInNormalYear  = 365
	secondsInWeek     = 691200
	secondsInMonth    = 2678400
)

// Represents the different string formats for dates
const (
	DefaultFormat       = "2006-01-02 15:04:05"
	DefaultSiZhuFormat  = "2006-01-02 15:04"
	DateFormat          = "2006-01-02"
	FormattedDateFormat = "Jan 2, 2006"
	TimeFormat          = "15:04:05"
	HourMinuteFormat    = "15:04"
	HourFormat          = "15"
	DayDateTimeFormat   = "Mon, Aug 2, 2006 3:04 PM"
	CookieFormat        = "Monday, 02-Jan-2006 15:04:05 MST"
	RFC822Format        = "Mon, 02 Jan 06 15:04:05 -0700"
	RFC1036Format       = "Mon, 02 Jan 06 15:04:05 -0700"
	RFC2822Format       = "Mon, 02 Jan 2006 15:04:05 -0700"
	RFC3339Format       = "2006-01-02T15:04:05-07:00"
	RSSFormat           = "Mon, 02 Jan 2006 15:04:05 -0700"
)

type WXTime struct {
	time.Time
	weekStartsAt time.Weekday
	weekEndsAt   time.Weekday
	weekendDays  []time.Weekday
	stringFormat string
	sunrise      time.Time
	sunset       time.Time
	latitude     float64
	longitude    float64
}

var defaultlatitudelongitude = []float64{32.7601261133, 111.9238493474}

// Used for testing purposes
var (
	isTimeFrozen      bool
	currentFrozenTime time.Time
)

// NewWXTime returns a pointer to a new WXTime instance
func NewWXTime(t time.Time) *WXTime {
	wds := []time.Weekday{
		time.Saturday,
		time.Sunday,
	}
	rise, set := SunriseSunset(defaultlatitudelongitude[0], defaultlatitudelongitude[1], t.Year(), t.Month(), t.Day())
	return &WXTime{
		Time:         t,
		weekStartsAt: time.Monday,
		weekEndsAt:   time.Sunday,
		weekendDays:  wds,
		stringFormat: DefaultSiZhuFormat,
		latitude:     defaultlatitudelongitude[0],
		longitude:    defaultlatitudelongitude[1],
		sunrise:      rise,
		sunset:       set,
	}
}

// Freeze allows time to be frozen to facilitate testing
func Freeze(time time.Time) {
	currentFrozenTime = time
	isTimeFrozen = true
}

// UnFreeze returns time to normal operation
func UnFreeze() {
	isTimeFrozen = false
}

// IsTimeFrozen allows checking if time has been frozen
func IsTimeFrozen() bool {
	return isTimeFrozen
}

// After will be behave like time.After unless time has been frozen
// If time is frozen it will add the expected delay and immediately send the frozen time on the returned channel
func After(d time.Duration) <-chan time.Time {
	if isTimeFrozen {
		currentFrozenTime = currentFrozenTime.Add(d)
		c := make(chan time.Time, 1)
		c <- currentFrozenTime
		return c
	}

	return time.After(d)
}

// Tick will be behave like time.Tick unless time has been frozen
// If time is frozen it will tick normally but the date will be based on the frozen date
func Tick(d time.Duration) <-chan time.Time {
	if isTimeFrozen {
		c := make(chan time.Time, 1)
		go func() {
			for {
				currentFrozenTime = currentFrozenTime.Add(d)
				c <- currentFrozenTime
			}
		}()
		return c
	}

	return time.Tick(d)
}

// Sleep will be behave like time.Sleep unless time has been frozen
// If time is frozen it will add the expected sleep delay and return immediately
func Sleep(d time.Duration) {
	if isTimeFrozen && d > 0 {
		currentFrozenTime = currentFrozenTime.Add(d)

		return
	}

	time.Sleep(d)
}

// create returns a new WXTime pointe. It is a helper function to create new dates
func create(y int, mon time.Month, d, h, m, s, ns int, l *time.Location) *WXTime {
	return NewWXTime(time.Date(y, mon, d, h, m, s, ns, l))
}

// Create returns a new pointer to WXTime instance from a specific date and time.
// If the location is invalid, it returns an error instead.
func Create(y int, mon time.Month, d, h, m, s, ns int, location string) (*WXTime, error) {
	l, err := time.LoadLocation(location)
	if err != nil {
		return nil, err
	}
	return create(y, mon, d, h, m, s, ns, l), nil
}

// CreateFromDate returns a new pointer to a WXTime instance from just a date.
// The time portion is set to now.
// If the location is invalid, it returns an error instead.
func CreateFromDate(y int, mon time.Month, d int, location string) (*WXTime, error) {
	h, m, s := Now().Clock()
	ns := Now().Nanosecond()

	return Create(y, mon, d, h, m, s, ns, location)
}

// CreateFromTime returns a new pointer to a WXTime instance from just a date.
// The time portion is set to now.
// If the locations is invalid, it returns an error instead.
func CreateFromTime(h, m, s, ns int, location string) (*WXTime, error) {
	y, mon, d := Now().Date()

	return Create(y, mon, d, h, m, s, ns, location)
}

// CreateFromFormat returns a new pointer to a WXTime instance from a specific format.
// If the location is invalid, it returns an error instead.
func CreateFromFormat(layout, value string, location string) (*WXTime, error) {
	l, err := time.LoadLocation(location)
	if err != nil {
		return nil, err
	}
	t, err := time.ParseInLocation(layout, value, l)
	if err != nil {
		return nil, err
	}

	return NewWXTime(t), nil
}

// CreateFromTimestamp returns a new pointer to a WXTime instance from a timestamp.
// If the location is invalid, it returns an error instead.
func CreateFromTimestamp(timestamp int64, location string) (*WXTime, error) {
	l, err := time.LoadLocation(location)
	if err != nil {
		return nil, err
	}
	t := NewWXTime(Now().In(l))
	t.SetTimestamp(timestamp)

	return t, nil
}

// CreateFromTimestampUTC returns a new pointer to a WXTime instance from an UTC timestamp.
// If the location is invalid, it returns an error instead.
func CreateFromTimestampUTC(timestamp int64) (*WXTime, error) {
	return CreateFromTimestamp(timestamp, "UTC")
}

// Parse returns a pointer to a new WXTime instance from a string
// If the location is invalid, it returns an error instead.
func Parse(layout, value, location string) (*WXTime, error) {
	l, err := time.LoadLocation(location)
	if err != nil {
		return nil, err
	}
	t, err := time.ParseInLocation(layout, value, l)
	if err != nil {
		return nil, err
	}

	return NewWXTime(t), nil
}

// Today returns a pointer to a new WXTime instance for today
// If the location is invalid, it returns an error instead.
func Today(location string) (*WXTime, error) {
	l, err := time.LoadLocation(location)
	if err != nil {
		return nil, err
	}

	return NewWXTime(Now().In(l)), err
}

// Tomorrow returns a pointer to a new WXTime instance for tomorrow
// If the location is invalid, it returns an error instead.
func Tomorrow(location string) (*WXTime, error) {
	c, err := Today(location)
	if err != nil {
		return nil, err
	}

	return c.AddDay(), nil
}

// Yesterday returns a pointer to a new WXTime instance for yesterday
// If the location is invalid, it returns an error instead.
func Yesterday(location string) (*WXTime, error) {
	c, err := Today(location)
	if err != nil {
		return nil, err
	}

	return c.SubDay(), nil
}

// unixTimeInSeconds represents the number of seconds between Year 1 and 1970
const unixTimeInSeconds = 62135596801

const maxNSecs = 999999999

// MaxValue returns a pointer to a new WXTime instance for greatest supported date
func MaxValue() *WXTime {
	return NewWXTime(time.Unix(math.MaxInt64-unixTimeInSeconds, maxNSecs))
}

// MinValue returns a pointer to a new WXTime instance for lowest supported date
func MinValue() *WXTime {
	return NewWXTime(time.Unix(math.MinInt64+unixTimeInSeconds, 0))
}

// Now returns a new WXTime instance for right now in current localtime
func Now() *WXTime {
	if isTimeFrozen {
		return NewWXTime(currentFrozenTime)
	}

	return NewWXTime(time.Now())
}

// NowInLocation returns a new WXTime instance for right now in given location.
// The location is in IANA Time Zone database, such as "America/New_York".
func NowInLocation(loc string) (*WXTime, error) {
	l, err := time.LoadLocation(loc)
	if err != nil {
		return nil, err
	}
	return nowIn(l), nil
}

func nowIn(loc *time.Location) *WXTime {
	return NewWXTime(Now().In(loc))
}

// Copy returns a new copy of the current WXTime instance
func (c *WXTime) Copy() *WXTime {
	return create(c.Year(), c.Month(), c.Day(), c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), c.Location())
}

// WeekStartsAt get the starting day of the week
func (c *WXTime) WeekStartsAt() time.Weekday {
	return c.weekStartsAt
}

// WeekEndsAt gets the ending day of the week
func (c *WXTime) WeekEndsAt() time.Weekday {
	return c.weekEndsAt
}

// WeekendDays gets the weekend days of the week
func (c *WXTime) WeekendDays() []time.Weekday {
	return c.weekendDays
}

// Quarter gets the current quarter
func (c *WXTime) Quarter() int {
	month := c.Month()
	switch {
	case month < 4:
		return 1
	case month >= 4 && month < 7:
		return 2
	case month >= 7 && month < 10:
		return 3
	}
	return 4
}

// Age gets the age from the current instance time to now
func (c *WXTime) Age() int {
	return int(c.DiffInYears(Now(), true))
}

// DaysInMonth returns the number of days in the month
func (c *WXTime) DaysInMonth() int {
	return c.EndOfMonth().Day()
}

// DaysInYear returns the number of days in the year
func (c *WXTime) DaysInYear() int {
	if c.IsLeapYear() {
		return daysInLeapYear
	}

	return daysInNormalYear
}

// WeekOfMonth returns the week of the month
func (c *WXTime) WeekOfMonth() int {
	w := math.Ceil(float64(c.Day() / daysPerWeek))
	return int(w + 1)
}

// WeekOfYear returns the week of the current year.
// This is an alias for time.ISOWeek
func (c *WXTime) WeekOfYear() (int, int) {
	return c.ISOWeek()
}

// TimeZone gets the current timezone
func (c *WXTime) TimeZone() string {
	return c.Location().String()
}

// Timestamp gets the current time since January 1, 1970 UTC
func (c *WXTime) Timestamp() int64 {
	return c.Unix()
}

// String gets the current date using the previously set format
func (c *WXTime) String() string {
	return c.Format(c.stringFormat)
}

// AddYears adds a year to the current time.
// Positive values travel forward while negative values travel into the past
func (c *WXTime) AddYears(y int) *WXTime {
	return NewWXTime(c.AddDate(y, 0, 0))
}

// AddYear adds a year to the current time
func (c *WXTime) AddYear() *WXTime {
	return c.AddYears(1)
}

// AddQuarters adds quarters to the current time.
// Positive values travel forward while negative values travel into the past
func (c *WXTime) AddQuarters(q int) *WXTime {
	return NewWXTime(c.AddDate(0, monthsPerQuarter*q, 0))
}

// AddQuarter adds a quarter to the current time
func (c *WXTime) AddQuarter() *WXTime {
	return c.AddQuarters(1)
}

// AddCenturies adds centuries to the time.
// Positive values travels forward while negative values travels into the past
func (c *WXTime) AddCenturies(cent int) *WXTime {
	return NewWXTime(c.AddDate(yearsPerCenturies*cent, 0, 0))
}

// AddCentury adds a century to the current time
func (c *WXTime) AddCentury() *WXTime {
	return c.AddCenturies(1)
}

// AddMonths adds months to the current time.
// Positive value travels forward while negative values travels into the past
func (c *WXTime) AddMonths(m int) *WXTime {
	return NewWXTime(c.AddDate(0, m, 0))
}

// AddMonth adds a month to the current time
func (c *WXTime) AddMonth() *WXTime {
	return c.AddMonths(1)
}

// AddSeconds adds seconds to the current time.
// Positive values travels forward while negative values travels into the past.
func (c *WXTime) AddSeconds(s int) *WXTime {
	d := time.Duration(s) * time.Second
	return NewWXTime(c.Add(d))
}

// AddSecond adds a second to the time
func (c *WXTime) AddSecond() *WXTime {
	return c.AddSeconds(1)
}

// AddDays adds a day to the current time.
// Positive value travels forward while negative value travels into the past
func (c *WXTime) AddDays(d int) *WXTime {
	return NewWXTime(c.AddDate(0, 0, d))
}

// AddDay adds a day to the current time
func (c *WXTime) AddDay() *WXTime {
	return c.AddDays(1)
}

// AddWeekdays adds a weekday to the current time.
// Positive value travels forward while negative value travels into the past
func (c *WXTime) AddWeekdays(wd int) *WXTime {
	d := 1
	if wd < 0 {
		wd, d = -wd, -d
	}
	t := c.Copy()
	for wd > 0 {
		t = t.AddDays(d)
		if t.IsWeekday() {
			wd--
		}
	}

	return t
}

// AddWeekday adds a weekday to the current time
func (c *WXTime) AddWeekday() *WXTime {
	return c.AddWeekdays(1)
}

// AddWeeks adds a week to the current time.
// Positive value travels forward while negative value travels into the past.
func (c *WXTime) AddWeeks(w int) *WXTime {
	return NewWXTime(c.AddDate(0, 0, daysPerWeek*w))
}

// AddWeek adds a week to the current time
func (c *WXTime) AddWeek() *WXTime {
	return c.AddWeeks(1)
}

// AddHours adds an hour to the current time.
// Positive value travels forward while negative value travels into the past
func (c *WXTime) AddHours(h int) *WXTime {
	d := time.Duration(h) * time.Hour

	return NewWXTime(c.Add(d))
}

// AddHour adds an hour to the current time
func (c *WXTime) AddHour() *WXTime {
	return c.AddHours(1)
}

// AddMonthsNoOverflow adds a month to the current time, not overflowing in case the
// destination month has less days than the current one.
// Positive value travels forward while negative value travels into the past.
func (c *WXTime) AddMonthsNoOverflow(m int) *WXTime {
	addedDate := NewWXTime(c.AddDate(0, m, 0))
	if c.Day() != addedDate.Day() {
		return addedDate.PreviousMonthLastDay()
	}

	return addedDate
}

// PreviousMonthLastDay returns the last day of the previous month
func (c *WXTime) PreviousMonthLastDay() *WXTime {
	return NewWXTime(c.AddDate(0, 0, -c.Day()))
}

// AddMonthNoOverflow adds a month with no overflow to the current time
func (c *WXTime) AddMonthNoOverflow() *WXTime {
	return c.AddMonthsNoOverflow(1)
}

// AddMinutes adds minutes to the current time.
// Positive value travels forward while negative value travels into the past.
func (c *WXTime) AddMinutes(m int) *WXTime {
	d := time.Duration(m) * time.Minute

	return NewWXTime(c.Add(d))
}

// AddMinute adds a minute to the current time
func (c *WXTime) AddMinute() *WXTime {
	return c.AddMinutes(1)
}

// SubYear removes a year from the current time
func (c *WXTime) SubYear() *WXTime {
	return c.SubYears(1)
}

// SubYears removes years from current time
func (c *WXTime) SubYears(y int) *WXTime {
	return c.AddYears(-1 * y)
}

// SubQuarter removes a quarter from the current time
func (c *WXTime) SubQuarter() *WXTime {
	return c.SubQuarters(1)
}

// SubQuarters removes quarters from current time
func (c *WXTime) SubQuarters(q int) *WXTime {
	return c.AddQuarters(-q)
}

// SubCentury removes a century from the current time
func (c *WXTime) SubCentury() *WXTime {
	return c.SubCenturies(1)
}

// SubCenturies removes centuries from the current time
func (c *WXTime) SubCenturies(cent int) *WXTime {
	return c.AddCenturies(-cent)
}

// SubMonth removes a month from the current time
func (c *WXTime) SubMonth() *WXTime {
	return c.SubMonths(1)
}

// SubMonths removes months from the current time
func (c *WXTime) SubMonths(m int) *WXTime {
	return c.AddMonths(-m)
}

// SubMonthNoOverflow remove a month with no overflow from the current time
func (c *WXTime) SubMonthNoOverflow() *WXTime {
	return c.SubMonthsNoOverflow(1)
}

// SubMonthsNoOverflow removes months with no overflow from the current time
func (c *WXTime) SubMonthsNoOverflow(m int) *WXTime {
	return c.AddMonthsNoOverflow(-m)
}

// SubDay removes a day from the current instance
func (c *WXTime) SubDay() *WXTime {
	return c.SubDays(1)
}

// SubDays removes days from the current time
func (c *WXTime) SubDays(d int) *WXTime {
	return c.AddDays(-d)
}

// SubWeekday removes a weekday from the current time
func (c *WXTime) SubWeekday() *WXTime {
	return c.SubWeekdays(1)
}

// SubWeekdays removes a weekday from the current time
func (c *WXTime) SubWeekdays(wd int) *WXTime {
	return c.AddWeekdays(-wd)
}

// SubWeek removes a week from the current time
func (c *WXTime) SubWeek() *WXTime {
	return c.SubWeeks(1)
}

// SubWeeks removes weeks to the current time
func (c *WXTime) SubWeeks(w int) *WXTime {
	return c.AddWeeks(-w)
}

// SubHour removes an hour from the current time
func (c *WXTime) SubHour() *WXTime {
	return c.SubHours(1)
}

// SubHours removes hours from the current time
func (c *WXTime) SubHours(h int) *WXTime {
	return c.AddHours(-h)
}

// SubMinute removes a minute from the current time
func (c *WXTime) SubMinute() *WXTime {
	return c.SubMinutes(1)
}

// SubMinutes removes minutes from the current time
func (c *WXTime) SubMinutes(m int) *WXTime {
	return c.AddMinutes(-m)
}

// SubSecond removes a second from the current time
func (c *WXTime) SubSecond() *WXTime {
	return c.SubSeconds(1)
}

// SubSeconds removes seconds from the current time
func (c *WXTime) SubSeconds(s int) *WXTime {
	return c.AddSeconds(-s)
}

// SetYear sets the year of the current time
func (c *WXTime) SetYear(y int) {
	c.Time = time.Date(y, c.Month(), c.Day(), c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), c.Location())
}

// SetMonth sets the month of the current time
func (c *WXTime) SetMonth(m time.Month) {
	c.Time = time.Date(c.Year(), m, c.Day(), c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), c.Location())
}

// SetDay sets the day of the current time
func (c *WXTime) SetDay(d int) {
	c.Time = time.Date(c.Year(), c.Month(), d, c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), c.Location())
}

// SetHour sets the hour of the current time
func (c *WXTime) SetHour(h int) {
	c.Time = time.Date(c.Year(), c.Month(), c.Day(), h, c.Minute(), c.Second(), c.Nanosecond(), c.Location())
}

// SetMinute sets the minute of the current time
func (c *WXTime) SetMinute(m int) {
	c.Time = time.Date(c.Year(), c.Month(), c.Day(), c.Hour(), m, c.Second(), c.Nanosecond(), c.Location())
}

// SetSecond sets the second of the current time
func (c *WXTime) SetSecond(s int) {
	c.Time = time.Date(c.Year(), c.Month(), c.Day(), c.Hour(), c.Minute(), s, c.Nanosecond(), c.Location())
}

// SetDate sets only the date of the current time
func (c *WXTime) SetDate(y int, m time.Month, d int) {
	c.Time = time.Date(y, m, d, c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), c.Location())
}

// SetDateTime sets the date and the time
func (c *WXTime) SetDateTime(y int, mon time.Month, d, h, m, s int) {
	c.Time = time.Date(y, mon, d, h, m, s, c.Nanosecond(), c.Location())
}

// SetTimeFromTimeString receives a string and sets the current time
// It accepts the following formats: "hh:mm:ss", "hh:mm" and "hh"
func (c *WXTime) SetTimeFromTimeString(timeString string) error {
	layouts := []string{
		TimeFormat,
		HourMinuteFormat,
		HourFormat,
	}

	var t time.Time
	var err error
	for i, layout := range layouts {
		t, err = time.Parse(layout, timeString)
		if err == nil {
			h, m, s := t.Clock()
			switch i {
			case 1:
				s = c.Second()
			case 2:
				m, s = c.Minute(), c.Second()
			}
			c.SetHour(h)
			c.SetMinute(m)
			c.SetSecond(s)
			return nil
		}
	}

	return errors.New("only supports hh:mm:ss, hh:mm and hh formats")
}

// SetWeekEndsAt sets the last day of week
func (c *WXTime) SetWeekEndsAt(wd time.Weekday) {
	c.weekEndsAt = wd
}

// SetWeekStartsAt sets the first day of week
func (c *WXTime) SetWeekStartsAt(wd time.Weekday) {
	c.weekStartsAt = wd
}

// SetWeekendDays sets the weekend days
func (c *WXTime) SetWeekendDays(wds []time.Weekday) {
	c.weekendDays = wds
}

// SetTimestamp sets the current time given a timestamp
func (c *WXTime) SetTimestamp(sec int64) {
	t := time.Unix(sec, 0)
	c.Time = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), c.Location())
}

// SetTimeZone sets the location from a string
// If the location is invalid, it returns an error instead.
func (c *WXTime) SetTimeZone(name string) error {
	loc, err := time.LoadLocation(name)
	if err != nil {
		return err
	}
	c.Time = time.Date(c.Year(), c.Month(), c.Day(), c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), loc)

	return nil
}

// ResetStringFormat changes the format to the DefaultFormat
func (c *WXTime) ResetStringFormat() {
	c.stringFormat = DefaultFormat
}

// SetStringFormat formats the current time with the set format string
func (c *WXTime) SetStringFormat(format string) {
	c.stringFormat = format
}

// DateString return the current time in Y-m-d format
func (c *WXTime) DateString() string {
	return c.Format(DateFormat)
}

// FormattedDateString returns the current time as a readable date
func (c *WXTime) FormattedDateString() string {
	return c.Format(FormattedDateFormat)
}

// TimeString returns the current time in hh:mm:ss format
func (c *WXTime) TimeString() string {
	return c.Format(TimeFormat)
}

// DateTimeString returns the current time in Y-m-d hh:mm:ss format
func (c *WXTime) DateTimeString() string {
	return c.Format(DefaultFormat)
}

// DayDateTimeString returns the current time with a day, date and time format
func (c *WXTime) DayDateTimeString() string {
	return c.Format(DayDateTimeFormat)
}

// AtomString formats the current time to a Atom date format
func (c *WXTime) AtomString() string {
	return c.Format(RFC3339Format)
}

// CookieString formats the current time to a Cookie date format
func (c *WXTime) CookieString() string {
	return c.Format(CookieFormat)
}

// ISO8601String returns the current time in ISO8601 format
func (c *WXTime) ISO8601String() string {
	return c.Format(RFC3339Format)
}

// RFC822String returns the current time in RFC 822 format
func (c *WXTime) RFC822String() string {
	return c.Format(RFC822Format)
}

// RFC850String returns the current time in RFC 850 format
func (c *WXTime) RFC850String() string {
	return c.Format(time.RFC850)
}

// RFC1036String returns the current time in RFC 1036 format
func (c *WXTime) RFC1036String() string {
	return c.Format(RFC1036Format)
}

// RFC1123String returns the current time in RFC 1123 format
func (c *WXTime) RFC1123String() string {
	return c.Format(time.RFC1123Z)
}

// RFC2822String returns the current time in RFC 2822 format
func (c *WXTime) RFC2822String() string {
	return c.Format(RFC2822Format)
}

// RFC3339String returns the current time in RFC 3339 format
func (c *WXTime) RFC3339String() string {
	return c.Format(RFC3339Format)
}

// RSSString returns the current time for RSS format
func (c *WXTime) RSSString() string {
	return c.Format(RSSFormat)
}

// W3CString returns the current time for WWW Consortium format
func (c *WXTime) W3CString() string {
	return c.Format(RFC3339Format)
}

// IsWeekday determines if the current time is a weekday
func (c *WXTime) IsWeekday() bool {
	return !c.IsWeekend()
}

// IsWeekend determines if the current time is a weekend day
func (c *WXTime) IsWeekend() bool {
	d := c.Weekday()
	for _, wd := range c.WeekendDays() {
		if d == wd {
			return true
		}
	}

	return false
}

// IsYesterday determines if the current time is yesterday
func (c *WXTime) IsYesterday() bool {
	n := Now().SubDay()

	return c.IsSameDay(n)
}

// IsToday determines if the current time is today
func (c *WXTime) IsToday() bool {
	return c.IsSameDay(Now())
}

// IsTomorrow determines if the current time is tomorrow
func (c *WXTime) IsTomorrow() bool {
	n := Now().AddDay()

	return c.IsSameDay(n)
}

// IsFuture determines if the current time is in the future, ie. greater (after) than now
func (c *WXTime) IsFuture() bool {
	return c.After(Now().Time)
}

// IsPast determines if the current time is in the past, ie. less (before) than now
func (c *WXTime) IsPast() bool {
	return c.Before(Now().Time)
}

// IsLeapYear determines if current current time is a leap year
func (c *WXTime) IsLeapYear() bool {
	y := c.Year()
	if (y%4 == 0 && y%100 != 0) || y%400 == 0 {
		return true
	}

	return false
}

// IsLongYear determines if the instance is a long year
func (c *WXTime) IsLongYear() bool {
	carb := create(c.Year(), time.December, 31, 0, 0, 0, 0, c.Location())
	_, w := carb.WeekOfYear()

	return w == weeksPerLongYear
}

// IsSameAs compares the formatted values of the two dates.
// If passed date is nil, compares against today
func (c *WXTime) IsSameAs(format string, carb *WXTime) bool {
	if carb == nil {
		return c.Format(DefaultFormat) == Now().Format(DefaultFormat)
	}

	return c.Format(DefaultFormat) == carb.Format(DefaultFormat)
}

// IsCurrentYear determines if the current time is in the current year
func (c *WXTime) IsCurrentYear() bool {
	return c.Year() == Now().Year()
}

// IsSameYear checks if the passed in date is in the same year as the current time year.
// If passed date is nil, compares against today
func (c *WXTime) IsSameYear(carb *WXTime) bool {
	if carb == nil {
		return c.Year() == nowIn(c.Location()).Year()
	}

	return c.Year() == carb.Year()
}

// IsCurrentMonth determines if the current time is in the current month
func (c *WXTime) IsCurrentMonth() bool {
	return c.Month() == Now().Month()
}

// IsSameMonth checks if the passed in date is in the same month as the current month
// If passed date is nil, compares against today
func (c *WXTime) IsSameMonth(carb *WXTime, sameYear bool) bool {
	m := nowIn(c.Location()).Month()
	if carb != nil {
		m = carb.Month()
	}
	if sameYear {
		return c.IsSameYear(carb) && c.Month() == m
	}

	return c.Month() == m
}

// IsSameDay checks if the passed in date is the same day as the current day.
// If passed date is nil, compares against today
func (c *WXTime) IsSameDay(carb *WXTime) bool {
	n := nowIn(c.Location())
	if carb != nil {
		n = carb
	}

	return c.Year() == n.Year() && c.Month() == n.Month() && c.Day() == n.Day()
}

// IsSunday checks if this day is a Sunday.
func (c *WXTime) IsSunday() bool {
	return c.Weekday() == time.Sunday
}

// IsMonday checks if this day is a Monday.
func (c *WXTime) IsMonday() bool {
	return c.Weekday() == time.Monday
}

// IsTuesday checks if this day is a Tuesday.
func (c *WXTime) IsTuesday() bool {
	return c.Weekday() == time.Tuesday
}

// IsWednesday checks if this day is a Wednesday.
func (c *WXTime) IsWednesday() bool {
	return c.Weekday() == time.Wednesday
}

// IsThursday checks if this day is a Thursday.
func (c *WXTime) IsThursday() bool {
	return c.Weekday() == time.Thursday
}

// IsFriday checks if this day is a Friday.
func (c *WXTime) IsFriday() bool {
	return c.Weekday() == time.Friday
}

// IsSaturday checks if this day is a Saturday.
func (c *WXTime) IsSaturday() bool {
	return c.Weekday() == time.Saturday
}

// IsLastWeek returns true is the date is within last week
func (c *WXTime) IsLastWeek() bool {
	secondsInWeek := float64(secondsInWeek)
	difference := Now().Sub(c.Time)
	if difference.Seconds() > 0 && difference.Seconds() < secondsInWeek {
		return true
	}

	return false
}

// IsLastMonth returns true is the date is within last month
func (c *WXTime) IsLastMonth() bool {
	now := Now()

	monthDifference := now.Month() - c.Month()

	if absValue(true, int64(monthDifference)) != 1 {
		return false
	}

	if now.UnixNano() > c.UnixNano() && monthDifference == 1 {
		return true
	}

	return false
}

// Eq determines if the current WXTime is equal to another
func (c *WXTime) Eq(carb *WXTime) bool {
	return c.Equal(carb.Time)
}

// EqualTo determines if the current WXTime is equal to another
func (c *WXTime) EqualTo(carb *WXTime) bool {
	return c.Eq(carb)
}

// Ne determines if the current WXTime is not equal to another
func (c *WXTime) Ne(carb *WXTime) bool {
	return !c.Eq(carb)
}

// NotEqualTo determines if the current WXTime is not equal to another
func (c *WXTime) NotEqualTo(carb *WXTime) bool {
	return c.Ne(carb)
}

// Gt determines if the current WXTime is greater (after) than another
func (c *WXTime) Gt(carb *WXTime) bool {
	return c.After(carb.Time)
}

// GreaterThan determines if the current WXTime is greater (after) than another
func (c *WXTime) GreaterThan(carb *WXTime) bool {
	return c.Gt(carb)
}

// Gte determines if the instance is greater (after) than or equal to another
func (c *WXTime) Gte(carb *WXTime) bool {
	return c.Gt(carb) || c.Eq(carb)
}

// GreaterThanOrEqualTo determines if the instance is greater (after) than or equal to another
func (c *WXTime) GreaterThanOrEqualTo(carb *WXTime) bool {
	return c.Gte(carb) || c.Eq(carb)
}

// Lt determines if the instance is less (before) than another
func (c *WXTime) Lt(carb *WXTime) bool {
	return c.Before(carb.Time)
}

// LessThan determines if the instance is less (before) than another
func (c *WXTime) LessThan(carb *WXTime) bool {
	return c.Lt(carb)
}

// Lte determines if the instance is less (before) or equal to another
func (c *WXTime) Lte(carb *WXTime) bool {
	return c.Lt(carb) || c.Eq(carb)
}

// LessThanOrEqualTo determines if the instance is less (before) or equal to another
func (c *WXTime) LessThanOrEqualTo(carb *WXTime) bool {
	return c.Lte(carb)
}

// Between determines if the current instance is between two others
// eq Indicates if a > and < comparison should be used or <= or >=
func (c *WXTime) Between(a, b *WXTime, eq bool) bool {
	if a.Gt(b) {
		a, b = swap(a, b)
	}
	if eq {
		return c.Gte(a) && c.Lte(b)
	}

	return c.Gt(a) && c.Lt(b)
}

// Closest returns the closest date from the current time
func (c *WXTime) Closest(a, b *WXTime) *WXTime {
	if c.DiffInSeconds(a, true) < c.DiffInSeconds(b, true) {
		return a
	}

	return b
}

// Farthest returns the farthest date from the current time
func (c *WXTime) Farthest(a, b *WXTime) *WXTime {
	if c.DiffInSeconds(a, true) > c.DiffInSeconds(b, true) {
		return a
	}

	return b
}

// Min returns the minimum instance between a given instance and the current instance
func (c *WXTime) Min(carb *WXTime) *WXTime {
	if carb == nil {
		carb = nowIn(c.Location())
	}
	if c.Lt(carb) {
		return c
	}

	return carb
}

// Minimum returns the minimum instance between a given instance and the current instance
func (c *WXTime) Minimum(carb *WXTime) *WXTime {
	return c.Min(carb)
}

// Max returns the maximum instance between a given instance and the current instance
func (c *WXTime) Max(carb *WXTime) *WXTime {
	if carb == nil {
		carb = nowIn(c.Location())
	}

	if c.Gt(carb) {
		return c
	}

	return carb
}

// Maximum returns the maximum instance between a given instance and the current instance
func (c *WXTime) Maximum(carb *WXTime) *WXTime {
	return c.Max(carb)
}

// DiffInYears returns the difference in years
func (c *WXTime) DiffInYears(carb *WXTime, abs bool) int64 {
	if carb == nil {
		carb = nowIn(c.Location())
	}

	if c.Year() == carb.Year() {
		return 0
	}

	start := NewWXTime(c.Time)
	end := NewWXTime(carb.Time)
	if end.UnixNano() < start.UnixNano() {
		aux := start
		start = end
		end = aux
	}

	yearsAmmount := int64(end.Year()-start.Year()) - 1

	start.SetYear(end.Year())

	if start.UnixNano() <= end.UnixNano() {
		yearsAmmount++
	}

	return absValue(abs, yearsAmmount)
}

// DiffInMonths returns the difference in months
func (c *WXTime) DiffInMonths(carb *WXTime, abs bool) int64 {
	if carb == nil {
		carb = nowIn(c.Location())
	}

	if c.Month() == carb.Month() && c.Year() == carb.Year() {
		return 0
	}

	if c.Month() != carb.Month() && c.Year() == carb.Year() {
		diffInMonths := int64(carb.In(time.UTC).Month() - c.In(time.UTC).Month())
		remainingTime := int(carb.DiffInHours(c, true))

		if remainingTime < c.DaysInMonth()*hoursPerDay {
			return 0
		}

		return absValue(abs, diffInMonths)
	}

	m := monthsPerYear - c.In(time.UTC).Month() + carb.In(time.UTC).Month() - 1
	if c.Year() < carb.Year() && c.hasRemainingHours(carb) {
		m = m + 1
	}

	if c.Year() > carb.Year() {
		m = monthsPerYear - carb.In(time.UTC).Month() + c.In(time.UTC).Month() - 1

		if carb.hasRemainingHours(c) {
			m = m + 1
		}
	}

	diffYr := c.Year() - carb.Year()
	if math.Abs(float64(diffYr)) > 1 {
		dateWithoutMonths := c.AddMonths(int(m))
		diff := dateWithoutMonths.DiffInYears(carb, abs)*monthsPerYear + int64(m)

		return absValue(abs, diff)
	}

	diff := int64(m)

	if c.GreaterThan(carb) {
		diff = -diff
	}

	return absValue(abs, diff)
}

func (c *WXTime) hasRemainingHours(carb *WXTime) bool {
	totalHr := int64(c.DaysInMonth() * hoursPerDay)
	cHr := c.StartOfMonth().DiffInHours(c, false)
	remainHr := totalHr - cHr
	spentInHr := carb.StartOfMonth().DiffInHours(carb, false)

	return remainHr+spentInHr >= totalHr
}

// DiffDurationInString returns the duration difference in string format
func (c *WXTime) DiffDurationInString(carb *WXTime) string {
	if carb == nil {
		carb = nowIn(c.Location())
	}

	return strings.Replace(carb.Sub(c.Time).String(), "-", "", 1)
}

// DiffInWeeks returns the difference in weeks
func (c *WXTime) DiffInWeeks(carb *WXTime, abs bool) int64 {
	if carb == nil {
		carb = nowIn(c.Location())
	}
	return c.DiffInDays(carb, abs) / daysPerWeek
}

// DiffInDays returns the difference in days
func (c *WXTime) DiffInDays(carb *WXTime, abs bool) int64 {
	if carb == nil {
		carb = nowIn(c.Location())
	}
	return c.DiffInHours(carb, abs) / hoursPerDay
}

// DiffInNights returns the difference in nights
func (c *WXTime) DiffInNights(carb *WXTime, abs bool) int64 {
	if carb == nil {
		carb = nowIn(c.Location())
	}
	return c.DiffInDays(carb, abs)
}

// Filter represents a predicate used for filtering diffs
type Filter func(*WXTime) bool

// dayDuration reprensets a day in time.Duration format
const dayDuration = time.Hour * hoursPerDay

// DiffInDaysFiltered returns the difference in days using a filter
func (c *WXTime) DiffInDaysFiltered(f Filter, carb *WXTime, abs bool) int64 {
	return c.DiffFiltered(dayDuration, f, carb, abs)
}

// DiffInHoursFiltered returns the difference in hours using a filter
func (c *WXTime) DiffInHoursFiltered(f Filter, carb *WXTime, abs bool) int64 {
	return c.DiffFiltered(time.Hour, f, carb, abs)
}

// DiffInWeekdays returns the difference in weekdays
func (c *WXTime) DiffInWeekdays(carb *WXTime, abs bool) int64 {
	f := func(t *WXTime) bool {
		return t.IsWeekday()
	}

	return c.DiffFiltered(dayDuration, f, carb, abs)
}

// DiffInWeekendDays returns the difference in weekend days using a filter
func (c *WXTime) DiffInWeekendDays(carb *WXTime, abs bool) int64 {
	f := func(t *WXTime) bool {
		return t.IsWeekend()
	}

	return c.DiffFiltered(dayDuration, f, carb, abs)
}

// DiffFiltered returns the difference by the given duration using a filter
func (c *WXTime) DiffFiltered(duration time.Duration, f Filter, carb *WXTime, abs bool) int64 {
	if carb == nil {
		carb = nowIn(c.Location())
	}
	if c.IsSameDay(carb) {
		return 0
	}

	inverse := false
	var counter int64
	s := int64(duration.Seconds())
	start, end := c.Copy(), carb.Copy()
	if start.Gt(end) {
		start, end = swap(start, end)
		inverse = true
	}
	for start.DiffInSeconds(end, true)/s > 0 {
		if f(end) {
			counter++
		}
		end = NewWXTime(end.Add(-duration))
	}
	if inverse {
		counter = -counter
	}

	return absValue(abs, counter)
}

// DiffInHours returns the difference in hours
func (c *WXTime) DiffInHours(d *WXTime, abs bool) int64 {
	return c.DiffInMinutes(d, abs) / minutesPerHour
}

// DiffInMinutes returns the difference in minutes
func (c *WXTime) DiffInMinutes(d *WXTime, abs bool) int64 {
	return c.DiffInSeconds(d, abs) / secondsPerMinute
}

// DiffInSeconds returns the difference in seconds
func (c *WXTime) DiffInSeconds(carb *WXTime, abs bool) int64 {
	if carb == nil {
		carb = nowIn(c.Location())
	}
	diff := carb.Timestamp() - c.Timestamp()

	return absValue(abs, diff)
}

// SecondsSinceMidnight returns the number of seconds since midnight.
func (c *WXTime) SecondsSinceMidnight() int {
	startOfDay := c.StartOfDay()

	return int(c.DiffInSeconds(startOfDay, true))
}

// SecondsUntilEndOfDay returns the number of seconds until 23:59:59.
func (c *WXTime) SecondsUntilEndOfDay() int {
	dayEnd := c.EndOfDay()

	return int(c.DiffInSeconds(dayEnd, true))
}

// absValue returns the abs value if needed
func absValue(needsAbs bool, value int64) int64 {
	if needsAbs && value < 0 {
		return -value
	}

	return value
}

func swap(a, b *WXTime) (*WXTime, *WXTime) {
	return b, a
}

// StartOfDay returns the time at 00:00:00 of the same day
func (c *WXTime) StartOfDay() *WXTime {
	return create(c.Year(), c.Month(), c.Day(), 0, 0, 0, 0, c.Location())
}

// EndOfDay returns the time at 23:59:59 of the same day
func (c *WXTime) EndOfDay() *WXTime {
	return create(c.Year(), c.Month(), c.Day(), 23, 59, 59, maxNSecs, c.Location())
}

// StartOfMonth returns the date on the first day of the month and the time to 00:00:00
func (c *WXTime) StartOfMonth() *WXTime {
	return create(c.Year(), c.Month(), 1, 0, 0, 0, 0, c.Location())
}

// EndOfMonth returns the date at the end of the month and time at 23:59:59
func (c *WXTime) EndOfMonth() *WXTime {
	return create(c.Year(), c.Month()+1, 0, 23, 59, 59, maxNSecs, c.Location())
}

// StartOfQuarter returns the date at the first day of the quarter and time at 00:00:00
func (c *WXTime) StartOfQuarter() *WXTime {
	month := time.Month((c.Quarter()-1)*monthsPerQuarter + 1)

	return create(c.Year(), time.Month(month), 1, 0, 0, 0, 0, c.Location())
}

// EndOfQuarter returns the date at end of the quarter and time at 23:59:59
func (c *WXTime) EndOfQuarter() *WXTime {
	return c.StartOfQuarter().AddMonths(monthsPerQuarter - 1).EndOfMonth()
}

// StartOfYear returns the date at the first day of the year and the time at 00:00:00
func (c *WXTime) StartOfYear() *WXTime {
	return create(c.Year(), time.January, 1, 0, 0, 0, 0, c.Location())
}

// EndOfYear returns the date at end of the year and time to 23:59:59
func (c *WXTime) EndOfYear() *WXTime {
	return create(c.Year(), time.December, 31, 23, 59, 59, maxNSecs, c.Location())
}

// StartOfDecade returns the date at the first day of the decade and time at 00:00:00
func (c *WXTime) StartOfDecade() *WXTime {
	year := c.Year() - c.Year()%yearsPerDecade

	return create(year, time.January, 1, 0, 0, 0, 0, c.Location())
}

// EndOfDecade returns the date at the end of the decade and time at 23:59:59
func (c *WXTime) EndOfDecade() *WXTime {
	year := c.Year() - c.Year()%yearsPerDecade + yearsPerDecade - 1

	return create(year, time.December, 31, 23, 59, 59, maxNSecs, c.Location())
}

// StartOfCentury returns the date of the first day of the century at 00:00:00
func (c *WXTime) StartOfCentury() *WXTime {
	year := c.Year() - c.Year()%yearsPerCenturies

	return create(year, time.January, 1, 0, 0, 0, 0, c.Location())
}

// EndOfCentury returns the date of the end of the century at 23:59:59
func (c *WXTime) EndOfCentury() *WXTime {
	year := c.Year() - 1 - c.Year()%yearsPerCenturies + yearsPerCenturies

	return create(year, time.December, 31, 23, 59, 59, maxNSecs, c.Location())
}

// StartOfWeek returns the date of the first day of week at 00:00:00
func (c *WXTime) StartOfWeek() *WXTime {
	return c.Previous(c.WeekStartsAt())
}

// EndOfWeek returns the date of the last day of the week at 23:59:59
func (c *WXTime) EndOfWeek() *WXTime {
	return c.Next(c.WeekEndsAt()).EndOfDay()
}

// Next changes the time to the next occurrence of a given day of the week
func (c *WXTime) Next(wd time.Weekday) *WXTime {
	c = c.AddDay()
	for c.Weekday() != wd {
		c = c.AddDay()
	}

	return c.StartOfDay()
}

// NextWeekday goes forward to the next weekday
func (c *WXTime) NextWeekday() *WXTime {
	return c.AddWeekday()
}

// PreviousWeekday goes back to the previous weekday
func (c *WXTime) PreviousWeekday() *WXTime {
	return c.SubWeekday()
}

// NextWeekendDay goes forward to the next weekend day
func (c *WXTime) NextWeekendDay() *WXTime {
	c = c.AddDay()
	for !c.IsWeekend() {
		c = c.AddDay()
	}

	return c
}

// PreviousWeekendDay goes back to the previous weekend day
func (c *WXTime) PreviousWeekendDay() *WXTime {
	c = c.SubDay()
	for !c.IsWeekend() {
		c = c.SubDay()
	}

	return c
}

// Previous changes the time to the previous occurrence of a given day of the week
func (c *WXTime) Previous(wd time.Weekday) *WXTime {
	c = c.SubDay()
	for c.Weekday() != wd {
		c = c.SubDay()
	}

	return c.StartOfDay()
}

// FirstOfMonth returns the first occurrence of a given day of the week in the current month
func (c *WXTime) FirstOfMonth(wd time.Weekday) *WXTime {
	d := c.StartOfMonth()
	if d.Weekday() != wd {
		return d.Next(wd)
	}

	return d
}

// LastOfMonth returns the last occurrence of a given day of the week in the current month
func (c *WXTime) LastOfMonth(wd time.Weekday) *WXTime {
	d := c.EndOfMonth()
	if d.Weekday() != wd {
		return d.Previous(wd)
	}

	return d.StartOfDay()
}

// LastDayOfMonth returns a new WXTime instance with the last day of current month
func (c *WXTime) LastDayOfMonth() *WXTime {
	return NewWXTime(time.Date(c.Year(), c.Month(), c.DaysInMonth(), 0, 0, 0, 0, time.UTC))
}

// FirstDayOfMonth returns a new WXTime instance with the first day of current month
func (c *WXTime) FirstDayOfMonth() *WXTime {
	return NewWXTime(time.Date(c.Year(), c.Month(), 1, 0, 0, 0, 0, time.UTC))
}

// NthOfMonth returns the given occurrence of a given day of the week in the current month
// If the calculated occurrence is outside the scope of current month, no modifications are made
func (c *WXTime) NthOfMonth(nth int, wd time.Weekday) *WXTime {
	cp := c.Copy().StartOfMonth()
	i := 0
	if cp.Weekday() == wd {
		i++
	}
	for i < nth {
		cp = cp.Next(wd)
		i++
	}
	if cp.Gt(c.EndOfMonth()) {
		return c
	}

	return cp
}

// FirstOfQuarter returns the first occurrence of a given day of the week in the current quarter
func (c *WXTime) FirstOfQuarter(wd time.Weekday) *WXTime {
	d := c.StartOfQuarter()
	if d.Weekday() != wd {
		return d.Next(wd)
	}

	return d
}

// LastOfQuarter returns the last occurrence of a given day of the week in the current quarter
func (c *WXTime) LastOfQuarter(wd time.Weekday) *WXTime {
	d := c.EndOfQuarter()
	if d.Weekday() != wd {
		return d.Previous(wd)
	}

	return d.StartOfDay()
}

// NthOfQuarter returns the given occurrence of a given day of the week in the current quarter
// If the calculated occurrence is outside the scope of current quarter, no modifications are made
func (c *WXTime) NthOfQuarter(nth int, wd time.Weekday) *WXTime {
	cp := c.Copy().StartOfQuarter()
	i := 0
	if cp.Weekday() == wd {
		i++
	}
	for i < nth {
		cp = cp.Next(wd)
		i++
	}
	if cp.Gt(c.EndOfQuarter()) {
		return c
	}

	return cp
}

// FirstOfYear returns the first occurrence of a given day of the week in the current year
func (c *WXTime) FirstOfYear(wd time.Weekday) *WXTime {
	d := c.StartOfYear()
	if d.Weekday() != wd {
		return d.Next(wd)
	}

	return d
}

// LastOfYear returns the last occurrence of a given day of the week in the current year
func (c *WXTime) LastOfYear(wd time.Weekday) *WXTime {
	d := c.EndOfYear()
	if d.Weekday() != wd {
		return d.Previous(wd)
	}

	return d.StartOfDay()
}

// NthOfYear returns the given occurrence of a given day of the week in the current year
// If the calculated occurrence is outside the scope of current year, no modifications are made
func (c *WXTime) NthOfYear(nth int, wd time.Weekday) *WXTime {
	cp := c.Copy().StartOfYear()
	i := 0
	if cp.Weekday() == wd {
		i++
	}
	for i < nth {
		cp = cp.Next(wd)
		i++
	}
	if cp.Gt(c.EndOfYear()) {
		return c
	}

	return cp
}

// Average returns the average between a given WXTime date and the current date
func (c *WXTime) Average(carb *WXTime) *WXTime {
	if carb == nil {
		carb = nowIn(c.Location())
	}
	if c.Eq(carb) {
		return c.Copy()
	}
	average := int(c.DiffInSeconds(carb, false) / 2)

	return c.AddSeconds(average)
}

//sizhu
type wxSiZhu struct {
	WXTime
	name             string
	sex              string
	nianzhu          GanZhi
	yuezhu           GanZhi
	rizhu            GanZhi
	shizhu           GanZhi
	SolarTermStartAt WXTime //上个节气时间
	SolarTermEndsAt  WXTime //下个节气时间
}
