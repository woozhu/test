package wxXunHuan

import (
	"strconv"
	"time"
	//"time"
)

type LunarObj struct {
	lunarYear             int
	lunarMonth            int
	lunarDay              int
	solarYear             int
	solarMonth            int
	solarDay              int
	weekNumber            int
	lastTermDay           int //上一个节气日期
	nextTermDay           int //下一个节气日期
	animal                string
	lunarMonthChineseName string
	lunarDayChineseName   string
	ganzhiYear            string
	ganzhiMonth           string
	ganzhiDay             string
	term                  string
	weekChineseName       string //星期几
	lastTerm              string //上一个节气
	nextTerm              string //下一个节气
	xingzuo               string //星座
	xingxiu               string //星宿
	jijie                 string //季节
	jieri                 string //节日
	pengzubaiji           string //彭祖百忌
	jinriyi               string
	jinriji               string
	sunrise               string //日出时间
	noon                  string //正午时间
	sunset                string //日落时间

	isToday bool
	isLeap  bool
	isTerm  bool
}

var (
	LunarYear   int
	LunarMonth  int
	LunarDay    int
	SolarYear   int
	SolarMonth  int
	SolarDay    int
	WeekNumber  int
	LastTermDay int //上一个节气日期
	NextTermDay int //下一个节气日期

	Animal                string
	LunarMonthChineseName string
	LunarDayChineseName   string
	GanzhiYear            string
	GanzhiMonth           string
	GanzhiDay             string
	Term                  string
	WeekChineseName       string //星期几
	LastTerm              string //上一个节气
	NextTerm              string //下一个节气
	Xingzuo               string
	Xingxiu               string
	Jijie                 string //季节
	Jieri                 string //节日
	Pengzubaiji           string //彭祖百忌
	Jinriyi               string
	Jinriji               string

	Sunrise string //日出时间
	Noon    string //正午时间
	Sunset  string //日落时间

	IsToday bool
	IsLeap  bool
	IsTerm  bool
)

//solarterm id [0-23]
type SolarTerms struct {
	id    int
	name  string
	cname string
	time  time.Time
}

func lYearDays(y int) int {
	var (
		i   = 0
		sum = 348
	)
	for i = 0x8000; i > 0x8; i >>= 1 {
		if lunarInfo[y-1900]&i != 0 {
			sum += 1
		}
	}
	return sum + leapDays(y)
}

//获得某年的闰月是几月
func leapMonth(y int) int {
	return lunarInfo[y-1900] & 0xf
}
func leapDays(y int) int {
	if leapMonth(y) != 0 {
		if lunarInfo[y-1900]&0x10000 != 0 {
			return 30
		} else {
			return 29
		}
	}
	return 0
}
func monthDays(y int, m int) int {
	if m > 12 || m < 1 {
		return -1
	} //月份参数从1至12，参数错误返回-1
	if lunarInfo[y-1900]&(0x10000>>uint32(m)) != 0 {
		return 30
	} else {
		return 29
	}
}

//计算公历的某年某月的天数
func solarDays(y int, m int) int {
	if m > 12 || m < 1 {
		return -1
	} //若参数错误 返回-1
	var ms = m - 1
	if ms == 1 { //2月份的闰平规律测算后确认返回28或29
		if (y%4 == 0) && (y%100 != 0) || (y%400 == 0) {
			return 29
		} else {
			return 28
		}
	} else {
		return (solarMonth[ms])
	}
}

func toGanZhi(offset int) string {
	return dTGcnameList[offset%10] + dDZcnameList[offset%12]
}

//截取字符串 start 起点下标 length 需要截取的长度
func Substr(str string, start int, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}

//GetLunarInfo 取得月历信息
func GetLunarInfo(y int) int {
	y = y - 1900
	if y < 0 || y > len(lunarInfo) {
		return 0
	}
	return lunarInfo[y]
}

//传入公历(!)y年获得该年第n个节气的公历日期
//@param y公历年(1900-2100)；n二十四节气中的第几个节气(1~24)；从n=1(小寒)算起
//@return day Number
//@eg:var _qm = getTerm(1987,7) ;//_qm=5;意即1987年4月5日清明
func getTerm(y, n int) int {
	y = y - 1900
	if y < 0 || y > len(sTermInfo) {
		return -1
	}
	if n < 1 || n > 24 {
		return -1
	}
	i := (n - 1) / 4 * 5
	n = (n - 1) % 4
	idx, _ := strconv.ParseInt(sTermInfo[y][i:i+5], 16, 64)
	a := strconv.FormatInt(idx, 10)
	day := []string{a[0:1], a[1:3], a[3:4], a[4:6]}
	i, _ = strconv.Atoi(day[n])
	return i
}

func toChinaMonth(m int) string {
	if m > 12 || m < 1 {
		return "?月"
	}
	var s = nStr3[m-1]
	s += "月" //加上月字
	return s
}

func toChinaDay(d int) string {
	if d < 0 || d > 31 {
		return "?日"
	}
	var s string
	switch d {
	case 10:
		s = `初十`
	case 20:
		s = `二十`
	case 30:
		s = `三十`
	default:
		s = nStr2[d/10] + nStr1[d%10]
	}
	return s + "日"
}

//年份转生肖[!仅能大致转换] => 精确划分生肖分界线是“立春”
func getAnimal(y int) string {
	return Animals[(y-4)%12]
}

func getAstro(m int, d int) string {
	arr := []int{20, 19, 21, 21, 21, 22, 23, 23, 23, 23, 22, 22}
	idx := d < arr[m-1]
	index := m * 2
	if idx {
		index = m*2 - 2
	}
	return constellation[index] + "座"
}

//主函数，返回年月日对应的LunarObj数据格式。
func solar2lunar(year int, month int, day int) *LunarObj {
	var obj = &LunarObj{}
	if year < 1900 || year > 2100 {
		return obj
	}
	if year == 1900 && month == 1 && day < 31 {
		return obj
	}

	var (
		i    = 0
		temp = 0
		leap = 0

		offset = deltaDaysWith19000131(year, month, day)
	)

	for i = 1900; i < 2101 && offset > 0; i++ {
		temp = lYearDays(i)
		offset -= temp
	}

	if offset < 0 {
		offset += temp
		i--
	}

	var (
		lunarYear = i

		isLeap = false
	)
	leap = leapMonth(i)
	for i = 1; i < 13 && offset > 0; i++ {
		if leap > 0 && i == (leap+1) && isLeap == false {
			i--
			isLeap = true
			temp = leapDays(lunarYear)
		} else {
			temp = monthDays(lunarYear, i)
		}

		if isLeap == true && i == (leap+1) {
			isLeap = false
		}

		offset -= temp
	}

	if offset == 0 && leap > 0 && i == leap+1 {
		if isLeap {
			isLeap = false
		} else {
			isLeap = true
			i--
		}
	}

	if offset < 0 {
		offset += temp
		i--
	}

	var (
		lunarMonth = i
		lunarDay   = offset + 1

		sm    = month - 1
		term3 = getTerm(lunarYear, 3)
		gzY   = toGanZhi(lunarYear - 4)
	)
	//term3为立春日期
	if sm < 2 && day < term3 {
		gzY = toGanZhi(lunarYear - 5)
	} else {
		gzY = toGanZhi(lunarYear - 4)
	}
	//firstNode为当月的上一个节气，secondNode为第二个节气。
	var (
		firstNode  = getTerm(year, month*2-1)
		secondNode = getTerm(year, month*2)

		gzM = toGanZhi((year-1900)*12 + month + 11)
	)
	if day >= firstNode {
		gzM = toGanZhi((year-1900)*12 + month + 12)
	}

	var (
		isTerm = false
		Term   = ""
	)
	if firstNode == day {
		isTerm = true
		Term = solarTerm[month*2-2]
	}

	if secondNode == day {
		isTerm = true
		Term = solarTerm[month*2-1]
	}

	var (
		dayCyclical = (deltaDaysWith19000131(year, sm+1, 1) + 30) + 10
		gzD         = toGanZhi(dayCyclical + day - 1)
	)

	obj.lunarYear = lunarYear
	obj.lunarMonth = lunarMonth
	obj.lunarDay = lunarDay

	obj.animal = getAnimal(lunarYear)
	if isLeap == true {
		obj.lunarMonthChineseName = "闰" + toChinaMonth(lunarMonth)
	} else {
		obj.lunarMonthChineseName = toChinaMonth(lunarMonth)
	}
	obj.lunarDayChineseName = toChinaDay(lunarDay)

	obj.solarYear = year
	obj.solarMonth = month
	obj.solarDay = day

	obj.ganzhiYear = gzY
	obj.ganzhiMonth = gzM
	obj.ganzhiDay = gzD

	obj.isLeap = isLeap
	obj.term = Term
	obj.isTerm = isTerm

	return obj
}
func lunar2solar(year int, month int, day int, isLeapMonth bool) *LunarObj {
	var (
		obj       = &LunarObj{}
		leapmonth = leapMonth(year)
	)

	if isLeapMonth && leapmonth != month {
		return obj
	}
	if (year == 2100 && month == 12 && day > 1) || (year == 1900 && month == 1 && day < 31) {
		return obj
	}
	var dd = monthDays(year, month)
	var ddd int
	if isLeapMonth == true {
		ddd = leapDays(year)
	} else {
		ddd = dd
	}
	if year < 1900 || year > 2100 || day > ddd {
		return obj
	}

	var offset = 0

	for i := 1900; i < year; i++ {
		offset += lYearDays(i)
	}

	var (
		leap  = 0
		isAdd = false
	)
	for i := 1; i < month; i++ {
		leap = leapMonth(year)
		if !isAdd {
			if leap <= i && leap > 0 {
				offset += leapDays(year)
				isAdd = true
			}
		}

		offset += monthDays(year, i)
	}

	if isLeapMonth == true {
		offset += dd
	}

	var (
		y = 1900
		m = 1
		d = 30
	)
	offset = offset + day

	for true {
		yy := y
		var t int
		if (yy%4 == 0 && yy%100 != 0) || yy%400 == 0 {
			t = 366
		} else {
			t = 365
		}

		if offset >= t {
			offset -= t
			y++
		} else {
			for i := 0; i < 12; i++ {
				var tmp = solarMonth[i]
				if t == 366 && i == 1 {
					tmp++
				}

				if offset > tmp {
					offset -= tmp
				} else {
					m = i + 1
					d += offset

					a := i
					for d > tmp {
						d -= tmp
						m += 1

						if m > 12 {
							y++
							m = 1
						}

						tmp = solarMonth[(a+1)%12]
						if t == 366 && a+1 == 1 {
							tmp++
						}
						a++
					}

					break
				}
			}
			break
		}
	}

	return solar2lunar(y, m, d)
}

func deltaDaysWith19000131(year int, month int, day int) int {
	var s = 0
	if year == 1900 {
		for m := 2; m < month; m++ {
			s += solarDays(year, m)
		}
		if month != 1 {
			s += day
		}
		return s
	} else {
		var s = 365 - 31
		for i := 1901; i < year; i++ {
			if (i%4 == 0 && i%100 != 0) || i%400 == 0 {
				s += 366
			} else {
				s += 365
			}
		}

		if month != 1 {
			for i := 1; i < month; i++ {
				s += solarDays(year, i)
			}
		}
		s += day
		return s
	}
}

//GetStemBranch 取得干支
func GetStemBranch(y int) string {
	return dTGcnameList[y%10] + dDZcnameList[y%12]
}

//StemBranchHour 获取时柱
//　	子 　　丑 　　寅 　　卯 　　辰 　　己
//　　　23-01：01-03：03-05 :05-07：07-09：09-11
//　　　午 　　未 　　申 　　酉 　　戊 　　亥
//　　　11-13：13-15：15-17：17-19：19-21：21-23
//`甲子`, `乙丑`, `丙寅`, `丁卯`, `戊辰`, `己巳`, `庚午`, `辛未`, `壬申`, `癸酉`, `甲戌`, `乙亥`, //甲或己日
//`丙子`, `丁丑`, `戊寅`, `己卯`, `庚辰`, `辛巳`, `壬午`, `癸未`, `甲申`, `乙酉`, `丙戌`, `丁亥`, //乙或庚日
//`戊子`, `己丑`, `庚寅`, `辛卯`, `壬辰`, `癸巳`, `甲午`, `乙未`, `丙申`, `丁酉`, `戊戌`, `己亥`, //丙或辛日
//`庚子`, `辛丑`, `壬寅`, `癸卯`, `甲辰`, `乙巳`, `丙午`, `丁未`, `戊申`, `己酉`, `庚戌`, `辛亥`, //丁或壬日
//`壬子`, `癸丑`, `甲寅`, `乙卯`, `丙辰`, `丁巳`, `戊午`, `己未`, `庚申`, `辛酉`, `壬戌`, `癸亥`, //戊或癸日
func StemBranchHour(y, m, d, h int) string {
	i := stemBranchIndex(y, m, d) % 5 * 12
	h = h / 2 % 12
	return stemBranchTable[i+h]
}

//
func stemBranchIndex(y, m, d int) int {
	y = y - 1900
	if y < 0 || y > len(yearNumber) {
		return 0
	}
	if m < 3 {
		y--
	}
	m = (m - 1) % 12
	return (yearNumber[y] + monthNumber[m] + d - 1) % 60
}

// StemBranchDay 获取日柱
func StemBranchDay(y, m, d int) string {
	return stemBranchTable[stemBranchIndex(y, m, d)]
}

//StemBranchMonth 获取月柱
func StemBranchMonth(y, m, d int) string {
	//月柱 1900年1月小寒以前为 丙子月(60进制12)
	fir := getTerm(y, m*2-1) //返回当月「节」为几日开始
	//sec := GetTermInfo(y, m*2)   //返回当月「节」为几日开始

	//依据12节气修正干支月
	var sb = GetStemBranch((y-1900)*12 + m + 11)
	if d >= fir {
		sb = GetStemBranch((y-1900)*12 + m + 12)
	}
	return sb
}

//StemBranchYear 获取年柱
func StemBranchYear(y int) string {
	num := y - 4
	return dTGcnameList[num%10] + dDZcnameList[num%12]
}
