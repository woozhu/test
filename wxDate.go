package wxXunHuan
//从1600年开始，到2400年结束。
//立春时间

import (
	"strconv"
	//"time"
)

//天干列表
var (
	dTGinfo = []string{`甲`, `乙`, `丙`, `丁`, `戊`, `己`, `庚`, `辛`, `壬`, `癸`}

	//地支列表
	dDZinfo = []string{`子`, `丑`, `寅`, `卯`, `辰`, `巳`, `午`, `未`, `申`, `酉`, `戌`, `亥`}
	//
	lunarInfo = []int{
		0x04bd8, 0x04ae0, 0x0a570, 0x054d5, 0x0d260, 0x0d950, 0x16554, 0x056a0, 0x09ad0, 0x055d2, //1900-1909
		0x04ae0, 0x0a5b6, 0x0a4d0, 0x0d250, 0x1d255, 0x0b540, 0x0d6a0, 0x0ada2, 0x095b0, 0x14977, //1910-1919
		0x04970, 0x0a4b0, 0x0b4b5, 0x06a50, 0x06d40, 0x1ab54, 0x02b60, 0x09570, 0x052f2, 0x04970, //1920-1929
		0x06566, 0x0d4a0, 0x0ea50, 0x06e95, 0x05ad0, 0x02b60, 0x186e3, 0x092e0, 0x1c8d7, 0x0c950, //1930-1939
		0x0d4a0, 0x1d8a6, 0x0b550, 0x056a0, 0x1a5b4, 0x025d0, 0x092d0, 0x0d2b2, 0x0a950, 0x0b557, //1940-1949
		0x06ca0, 0x0b550, 0x15355, 0x04da0, 0x0a5b0, 0x14573, 0x052b0, 0x0a9a8, 0x0e950, 0x06aa0, //1950-1959
		0x0aea6, 0x0ab50, 0x04b60, 0x0aae4, 0x0a570, 0x05260, 0x0f263, 0x0d950, 0x05b57, 0x056a0, //1960-1969
		0x096d0, 0x04dd5, 0x04ad0, 0x0a4d0, 0x0d4d4, 0x0d250, 0x0d558, 0x0b540, 0x0b6a0, 0x195a6, //1970-1979
		0x095b0, 0x049b0, 0x0a974, 0x0a4b0, 0x0b27a, 0x06a50, 0x06d40, 0x0af46, 0x0ab60, 0x09570, //1980-1989
		0x04af5, 0x04970, 0x064b0, 0x074a3, 0x0ea50, 0x06b58, 0x055c0, 0x0ab60, 0x096d5, 0x092e0, //1990-1999
		0x0c960, 0x0d954, 0x0d4a0, 0x0da50, 0x07552, 0x056a0, 0x0abb7, 0x025d0, 0x092d0, 0x0cab5, //2000-2009
		0x0a950, 0x0b4a0, 0x0baa4, 0x0ad50, 0x055d9, 0x04ba0, 0x0a5b0, 0x15176, 0x052b0, 0x0a930, //2010-2019
		0x07954, 0x06aa0, 0x0ad50, 0x05b52, 0x04b60, 0x0a6e6, 0x0a4e0, 0x0d260, 0x0ea65, 0x0d530, //2020-2029
		0x05aa0, 0x076a3, 0x096d0, 0x04bd7, 0x04ad0, 0x0a4d0, 0x1d0b6, 0x0d250, 0x0d520, 0x0dd45, //2030-2039
		0x0b5a0, 0x056d0, 0x055b2, 0x049b0, 0x0a577, 0x0a4b0, 0x0aa50, 0x1b255, 0x06d20, 0x0ada0, //2040-2049
		0x14b63, 0x09370, 0x049f8, 0x04970, 0x064b0, 0x168a6, 0x0ea50, 0x06b20, 0x1a6c4, 0x0aae0, //2050-2059
		0x0a2e0, 0x0d2e3, 0x0c960, 0x0d557, 0x0d4a0, 0x0da50, 0x05d55, 0x056a0, 0x0a6d0, 0x055d4, //2060-2069
		0x052d0, 0x0a9b8, 0x0a950, 0x0b4a0, 0x0b6a6, 0x0ad50, 0x055a0, 0x0aba4, 0x0a5b0, 0x052b0, //2070-2079
		0x0b273, 0x06930, 0x07337, 0x06aa0, 0x0ad50, 0x14b55, 0x04b60, 0x0a570, 0x054e4, 0x0d160, //2080-2089
		0x0e968, 0x0d520, 0x0daa0, 0x16aa6, 0x056d0, 0x04ae0, 0x0a9d4, 0x0a2d0, 0x0d150, 0x0f252, //2090-2099
		0x0d520}
	solarMonth = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	//
	stemBranchTable = []string{
		`甲子`, `乙丑`, `丙寅`, `丁卯`, `戊辰`, `己巳`, `庚午`, `辛未`, `壬申`, `癸酉`, `甲戌`, `乙亥`, //甲或己日
		`丙子`, `丁丑`, `戊寅`, `己卯`, `庚辰`, `辛巳`, `壬午`, `癸未`, `甲申`, `乙酉`, `丙戌`, `丁亥`, //乙或庚日
		`戊子`, `己丑`, `庚寅`, `辛卯`, `壬辰`, `癸巳`, `甲午`, `乙未`, `丙申`, `丁酉`, `戊戌`, `己亥`, //丙或辛日
		`庚子`, `辛丑`, `壬寅`, `癸卯`, `甲辰`, `乙巳`, `丙午`, `丁未`, `戊申`, `己酉`, `庚戌`, `辛亥`, //丁或壬日
		`壬子`, `癸丑`, `甲寅`, `乙卯`, `丙辰`, `丁巳`, `戊午`, `己未`, `庚申`, `辛酉`, `壬戌`, `癸亥`, //戊或癸日
	}

	//
	yearNumber = []int{
		0x9, 0xE, 0x13, 0x18, 0x1E, 0x23, 0x28, 0x2D, 0x33, 0x38, //1900-1909
		0x1, 0x6, 0xC, 0x11, 0x16, 0x1B, 0x21, 0x26, 0x2B, 0x30, //1910-1919
		0x36, 0x3B, 0x4, 0x9, 0xF, 0x14, 0x19, 0x1E, 0x24, 0x36, //1920-1929
		0x2E, 0x33, 0x39, 0x2, 0x7, 0xC, 0x12, 0x17, 0x1C, 0x2E, //1930-1939
		0x27, 0x2C, 0x31, 0x36, 0x0, 0x5, 0xA, 0xF, 0x15, 0x1A, //1940-1949
		0x1F, 0x24, 0x2A, 0x2F, 0x34, 0x39, 0x3, 0x8, 0xD, 0x12, //1950-1959
		0x18, 0x1D, 0x22, 0x27, 0x2D, 0x32, 0x37, 0x0, 0x6, 0xB, //1960-1969
		0x10, 0x15, 0x1B, 0x20, 0x25, 0x2A, 0x30, 0x35, 0x3A, 0x3, //1970-1979
		0x9, 0xE, 0x13, 0x18, 0x1E, 0x23, 0x28, 0x2D, 0x33, 0x38, //1980-1989
		0x1, 0x6, 0xC, 0x11, 0x16, 0x1B, 0x21, 0x26, 0x2B, 0x30, //1990-1999
		0x36, 0x3B, 0x4, 0x9, 0xF, 0x14, 0x19, 0x1E, 0x24, 0x29, //2000-2009
		0x2E, 0x33, 0x39, 0x2, 0x7, 0xC, 0x12, 0x17, 0x1C, 0x21, //2010-2019
		0x27, 0x2C, 0x31, 0x36, 0x0, 0x5, 0xA, 0xF, 0x15, 0x1A, //2020-2029
		0x1F, 0x24, 0x2A, 0x2F, 0x34, 0x39, 0x3, 0x8, 0xD, 0x12, //2030-2039
		0x18, 0x1D, 0x22, 0x27, 0x2D, 0x32, 0x37, 0x0, 0x6, 0xB, //2040-2049
		0x10, 0x15, 0x1B, 0x20, 0x25, 0x2A, 0x30, 0x35, 0x3A, 0x3, //2050-2059
		0x9, 0xE, 0x13, 0x18, 0x1E, 0x23, 0x28, 0x2D, 0x33, 0x38, //2060-2069
		0x1, 0x6, 0xC, 0x11, 0x16, 0x1B, 0x21, 0x26, 0x2B, 0x30, //2070-2079
		0x36, 0x3B, 0x4, 0x9, 0xF, 0x14, 0x19, 0x1E, 0x24, 0x29, //2080-2089
		0x2E, 0x33, 0x39, 0x2, 0x7, 0xC, 0x12, 0x17, 0x1C, 0x21, //2090-2099
	}

	//
	monthNumber = []int{
		0x6,  //next year 1
		0x25, //next year 2
		0x0,  //3
		0x1F, //4
		0x1,  //5
		0x20, //6
		0x2,  //7
		0x21, //8
		0x4,  //9
		0x22, //10
		0x5,  //11
		0x23, //12
	}

	sTermInfo = []string{
		`9778397bd097c36b0b6fc9274c91aa`, `97b6b97bd19801ec9210c965cc920e`, `97bcf97c3598082c95f8c965cc920f`, `97bd0b06bdb0722c965ce1cfcc920f`, `b027097bd097c36b0b6fc9274c91aa`,
		`97b6b97bd19801ec9210c965cc920e`, `97bcf97c359801ec95f8c965cc920f`, `97bd0b06bdb0722c965ce1cfcc920f`, `b027097bd097c36b0b6fc9274c91aa`, `97b6b97bd19801ec9210c965cc920e`,
		`97bcf97c359801ec95f8c965cc920f`, `97bd0b06bdb0722c965ce1cfcc920f`, `b027097bd097c36b0b6fc9274c91aa`, `9778397bd19801ec9210c965cc920e`, `97b6b97bd19801ec95f8c965cc920f`,
		`97bd09801d98082c95f8e1cfcc920f`, `97bd097bd097c36b0b6fc9210c8dc2`, `9778397bd197c36c9210c9274c91aa`, `97b6b97bd19801ec95f8c965cc920e`, `97bd09801d98082c95f8e1cfcc920f`,
		`97bd097bd097c36b0b6fc9210c8dc2`, `9778397bd097c36c9210c9274c91aa`, `97b6b97bd19801ec95f8c965cc920e`, `97bcf97c3598082c95f8e1cfcc920f`, `97bd097bd097c36b0b6fc9210c8dc2`,
		`9778397bd097c36c9210c9274c91aa`, `97b6b97bd19801ec9210c965cc920e`, `97bcf97c3598082c95f8c965cc920f`, `97bd097bd097c35b0b6fc920fb0722`, `9778397bd097c36b0b6fc9274c91aa`,
		`97b6b97bd19801ec9210c965cc920e`, `97bcf97c3598082c95f8c965cc920f`, `97bd097bd097c35b0b6fc920fb0722`, `9778397bd097c36b0b6fc9274c91aa`, `97b6b97bd19801ec9210c965cc920e`,
		`97bcf97c359801ec95f8c965cc920f`, `97bd097bd097c35b0b6fc920fb0722`, `9778397bd097c36b0b6fc9274c91aa`, `97b6b97bd19801ec9210c965cc920e`, `97bcf97c359801ec95f8c965cc920f`,
		`97bd097bd097c35b0b6fc920fb0722`, `9778397bd097c36b0b6fc9274c91aa`, `97b6b97bd19801ec9210c965cc920e`, `97bcf97c359801ec95f8c965cc920f`, `97bd097bd07f595b0b6fc920fb0722`,
		`9778397bd097c36b0b6fc9210c8dc2`, `9778397bd19801ec9210c9274c920e`, `97b6b97bd19801ec95f8c965cc920f`, `97bd07f5307f595b0b0bc920fb0722`, `7f0e397bd097c36b0b6fc9210c8dc2`,
		`9778397bd097c36c9210c9274c920e`, `97b6b97bd19801ec95f8c965cc920f`, `97bd07f5307f595b0b0bc920fb0722`, `7f0e397bd097c36b0b6fc9210c8dc2`, `9778397bd097c36c9210c9274c91aa`,
		`97b6b97bd19801ec9210c965cc920e`, `97bd07f1487f595b0b0bc920fb0722`, `7f0e397bd097c36b0b6fc9210c8dc2`, `9778397bd097c36b0b6fc9274c91aa`, `97b6b97bd19801ec9210c965cc920e`,
		`97bcf7f1487f595b0b0bb0b6fb0722`, `7f0e397bd097c35b0b6fc920fb0722`, `9778397bd097c36b0b6fc9274c91aa`, `97b6b97bd19801ec9210c965cc920e`, `97bcf7f1487f595b0b0bb0b6fb0722`,
		`7f0e397bd097c35b0b6fc920fb0722`, `9778397bd097c36b0b6fc9274c91aa`, `97b6b97bd19801ec9210c965cc920e`, `97bcf7f1487f531b0b0bb0b6fb0722`, `7f0e397bd097c35b0b6fc920fb0722`,
		`9778397bd097c36b0b6fc9274c91aa`, `97b6b97bd19801ec9210c965cc920e`, `97bcf7f1487f531b0b0bb0b6fb0722`, `7f0e397bd07f595b0b6fc920fb0722`, `9778397bd097c36b0b6fc9274c91aa`,
		`97b6b97bd19801ec9210c9274c920e`, `97bcf7f0e47f531b0b0bb0b6fb0722`, `7f0e397bd07f595b0b0bc920fb0722`, `9778397bd097c36b0b6fc9210c91aa`, `97b6b97bd197c36c9210c9274c920e`,
		`97bcf7f0e47f531b0b0bb0b6fb0722`, `7f0e397bd07f595b0b0bc920fb0722`, `9778397bd097c36b0b6fc9210c8dc2`, `9778397bd097c36c9210c9274c920e`, `97b6b7f0e47f531b0723b0b6fb0722`,
		`7f0e37f5307f595b0b0bc920fb0722`, `7f0e397bd097c36b0b6fc9210c8dc2`, `9778397bd097c36b0b70c9274c91aa`, `97b6b7f0e47f531b0723b0b6fb0721`, `7f0e37f1487f595b0b0bb0b6fb0722`,
		`7f0e397bd097c35b0b6fc9210c8dc2`, `9778397bd097c36b0b6fc9274c91aa`, `97b6b7f0e47f531b0723b0b6fb0721`, `7f0e27f1487f595b0b0bb0b6fb0722`, `7f0e397bd097c35b0b6fc920fb0722`,
		`9778397bd097c36b0b6fc9274c91aa`, `97b6b7f0e47f531b0723b0b6fb0721`, `7f0e27f1487f531b0b0bb0b6fb0722`, `7f0e397bd097c35b0b6fc920fb0722`, `9778397bd097c36b0b6fc9274c91aa`,
		`97b6b7f0e47f531b0723b0b6fb0721`, `7f0e27f1487f531b0b0bb0b6fb0722`, `7f0e397bd097c35b0b6fc920fb0722`, `9778397bd097c36b0b6fc9274c91aa`, `97b6b7f0e47f531b0723b0b6fb0721`,
		`7f0e27f1487f531b0b0bb0b6fb0722`, `7f0e397bd07f595b0b0bc920fb0722`, `9778397bd097c36b0b6fc9274c91aa`, `97b6b7f0e47f531b0723b0787b0721`, `7f0e27f0e47f531b0b0bb0b6fb0722`,
		`7f0e397bd07f595b0b0bc920fb0722`, `9778397bd097c36b0b6fc9210c91aa`, `97b6b7f0e47f149b0723b0787b0721`, `7f0e27f0e47f531b0723b0b6fb0722`, `7f0e397bd07f595b0b0bc920fb0722`,
		`9778397bd097c36b0b6fc9210c8dc2`, `977837f0e37f149b0723b0787b0721`, `7f07e7f0e47f531b0723b0b6fb0722`, `7f0e37f5307f595b0b0bc920fb0722`, `7f0e397bd097c35b0b6fc9210c8dc2`,
		`977837f0e37f14998082b0787b0721`, `7f07e7f0e47f531b0723b0b6fb0721`, `7f0e37f1487f595b0b0bb0b6fb0722`, `7f0e397bd097c35b0b6fc9210c8dc2`, `977837f0e37f14998082b0787b06bd`,
		`7f07e7f0e47f531b0723b0b6fb0721`, `7f0e27f1487f531b0b0bb0b6fb0722`, `7f0e397bd097c35b0b6fc920fb0722`, `977837f0e37f14998082b0787b06bd`, `7f07e7f0e47f531b0723b0b6fb0721`,
		`7f0e27f1487f531b0b0bb0b6fb0722`, `7f0e397bd097c35b0b6fc920fb0722`, `977837f0e37f14998082b0787b06bd`, `7f07e7f0e47f531b0723b0b6fb0721`, `7f0e27f1487f531b0b0bb0b6fb0722`,
		`7f0e397bd07f595b0b0bc920fb0722`, `977837f0e37f14998082b0787b06bd`, `7f07e7f0e47f531b0723b0b6fb0721`, `7f0e27f1487f531b0b0bb0b6fb0722`, `7f0e397bd07f595b0b0bc920fb0722`,
		`977837f0e37f14998082b0787b06bd`, `7f07e7f0e47f149b0723b0787b0721`, `7f0e27f0e47f531b0b0bb0b6fb0722`, `7f0e397bd07f595b0b0bc920fb0722`, `977837f0e37f14998082b0723b06bd`,
		`7f07e7f0e37f149b0723b0787b0721`, `7f0e27f0e47f531b0723b0b6fb0722`, `7f0e397bd07f595b0b0bc920fb0722`, `977837f0e37f14898082b0723b02d5`, `7ec967f0e37f14998082b0787b0721`,
		`7f07e7f0e47f531b0723b0b6fb0722`, `7f0e37f1487f595b0b0bb0b6fb0722`, `7f0e37f0e37f14898082b0723b02d5`, `7ec967f0e37f14998082b0787b0721`, `7f07e7f0e47f531b0723b0b6fb0722`,
		`7f0e37f1487f531b0b0bb0b6fb0722`, `7f0e37f0e37f14898082b0723b02d5`, `7ec967f0e37f14998082b0787b06bd`, `7f07e7f0e47f531b0723b0b6fb0721`, `7f0e37f1487f531b0b0bb0b6fb0722`,
		`7f0e37f0e37f14898082b072297c35`, `7ec967f0e37f14998082b0787b06bd`, `7f07e7f0e47f531b0723b0b6fb0721`, `7f0e27f1487f531b0b0bb0b6fb0722`, `7f0e37f0e37f14898082b072297c35`,
		`7ec967f0e37f14998082b0787b06bd`, `7f07e7f0e47f531b0723b0b6fb0721`, `7f0e27f1487f531b0b0bb0b6fb0722`, `7f0e37f0e366aa89801eb072297c35`, `7ec967f0e37f14998082b0787b06bd`,
		`7f07e7f0e47f149b0723b0787b0721`, `7f0e27f1487f531b0b0bb0b6fb0722`, `7f0e37f0e366aa89801eb072297c35`, `7ec967f0e37f14998082b0723b06bd`, `7f07e7f0e47f149b0723b0787b0721`,
		`7f0e27f0e47f531b0723b0b6fb0722`, `7f0e37f0e366aa89801eb072297c35`, `7ec967f0e37f14998082b0723b06bd`, `7f07e7f0e37f14998083b0787b0721`, `7f0e27f0e47f531b0723b0b6fb0722`,
		`7f0e37f0e366aa89801eb072297c35`, `7ec967f0e37f14898082b0723b02d5`, `7f07e7f0e37f14998082b0787b0721`, `7f07e7f0e47f531b0723b0b6fb0722`, `7f0e36665b66aa89801e9808297c35`,
		`665f67f0e37f14898082b0723b02d5`, `7ec967f0e37f14998082b0787b0721`, `7f07e7f0e47f531b0723b0b6fb0722`, `7f0e36665b66a449801e9808297c35`, `665f67f0e37f14898082b0723b02d5`,
		`7ec967f0e37f14998082b0787b06bd`, `7f07e7f0e47f531b0723b0b6fb0721`, `7f0e36665b66a449801e9808297c35`, `665f67f0e37f14898082b072297c35`, `7ec967f0e37f14998082b0787b06bd`,
		`7f07e7f0e47f531b0723b0b6fb0721`, `7f0e26665b66a449801e9808297c35`, `665f67f0e37f1489801eb072297c35`, `7ec967f0e37f14998082b0787b06bd`, `7f07e7f0e47f531b0723b0b6fb0721`,
		`7f0e27f1487f531b0b0bb0b6fb0722`, //2100
	}

	number = []string{`一`, `二`, `三`, `四`, `五`, `六`, `七`, `八`, `九`, `十`, `十一`, `十二`}
	ten    = []string{`初`, `十`, `廿`, `卅`}
	//二十四节气 wxdate
	solarTerm = []string{"小寒", "大寒", "立春", "雨水", "惊蛰", "春分", "清明", "谷雨", "立夏", "小满", "芒种", "夏至", "小暑", "大暑", "立秋", "处暑", "白露", "秋分", "寒露", "霜降", "立冬", "小雪", "大雪", "冬至"}
	//月历月份
	chineseNumber = []string{`正`, `二`, `三`, `四`, `五`, `六`, `七`, `八`, `九`, `十`, `十一`, `腊`}

	//公历每个月份的天数
	monthDay = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	//12属相
	Animals = []string{"鼠", "牛", "虎", "兔", "龙", "蛇", "马", "羊", "猴", "鸡", "狗", "猪"}
	//
	nStr1 = []string{"日", "一", "二", "三", "四", "五", "六", "七", "八", "九", "十"}
	//
	nStr2 = []string{"初", "十", "廿", "卅"}
	//
	nStr3 = []string{"正", "二", "三", "四", "五", "六", "七", "八", "九", "十", "十一", "腊"}
	//

	//12星座
	constellation = []string{
		`魔羯`, `水瓶`, `双鱼`, `白羊`, `金牛`, `双子`, `巨蟹`, `狮子`, `处女`, `天秤`, `天蝎`, `射手`,
	}
)

type LunarObj struct {
	lunarYear   int
	lunarMonth  int
	lunarDay    int
	solarYear   int
	solarMonth  int
	solarDay    int
	weekNumber  int
	lastTermDay int //上一个节气日期
	nextTermDay int //下一个节气日期

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
	xingzuo               string
	xingxiu               string
	jijie                 string //季节
	jieri                 string //节日
	pengzubaiji           string //彭祖百忌
	jinriyi               string
	jinriji               string

	sunrise string //日出时间
	noon    string //正午时间
	sunset  string //日落时间

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
	return dTGinfo[offset%10] + dDZinfo[offset%12]
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

//传入公历年月日以及传入的月份是否闰月获得详细的公历、农历object信息
//func lunar2solar(y int, m int, d int, isLeapMonth bool) *LunarObj{
//    if(y<1900 || y>2100) {return -1}//年份限定、上限
//	if(y==1900&&m==1&&d<31) {return -1}//下限
//	var objDate time.Time
//	if(!y) { //未传参  获得当天
//		objDate = time.Now()
//		//new Date()
//	}else {
//		objDate = time.Date(y,strconv.ParseInt(m)-1,d)
//		//new Date(y,parseInt(m)-1,d)
//	}
//	var i, leap=0, temp=0;
//	//修正ymd参数
//	var y = objDate.getFullYear(),m = objDate.getMonth()+1,d = objDate.getDate();
//	var offset   = (Date.UTC(objDate.getFullYear(),objDate.getMonth(),objDate.getDate()) - Date.UTC(1900,0,31))/86400000;
//	for(i=1900; i<2101 && offset>0; i++) { temp=lYearDays(i); offset-=temp; }
//	if(offset<0) { offset+=temp; i--; }

//	//是否今天
//	var isTodayObj = new Date(),isToday=false;
//	if(isTodayObj.getFullYear()==y && isTodayObj.getMonth()+1==m && isTodayObj.getDate()==d) {
//		isToday = true;
//	}
//	//星期几
//	var nWeek = objDate.getDay(),cWeek = nStr1[nWeek];
//	if(nWeek==0) {nWeek =7;}//数字表示周几顺应天朝周一开始的惯例
//		//农历年
//	var year = i;

//	var leap = leapMonth(i); //闰哪个月
//	var isLeap = false;

//	//效验闰月
//	for(i=1; i<13 && offset>0; i++) {
//		//闰月
//		if(leap>0 && i==(leap+1) && isLeap==false){
//			--i;
//			isLeap = true; temp = leapDays(year); //计算农历闰月天数
//		}else{
//			temp = monthDays(year, i);//计算农历普通月天数
//		}
//		//解除闰月
//		if(isLeap==true && i==(leap+1)) {
//		    isLeap = false
//		}
//		offset -= temp;
//	}

//	if(offset==0 && leap>0 && i==leap+1)
//	if(isLeap){
//		isLeap = false;
//	}else{
//		isLeap = true; --i;
//	}
//	if(offset<0){ offset += temp; --i; }
//		//农历月
//	var month 	= i;
//		//农历日
//	var day 		= offset + 1;

//		//天干地支处理
//	var sm 		= 	m-1;
//	var term3	=	getTerm(year,3); //该农历年立春日期
//	var gzY 		= 	toGanZhi(year-4);//普通按年份计算，下方尚需按立春节气来修正

//		//依据立春日进行修正gzY
//	if(sm<2 && d<term3) {
//		gzY = toGanZhi(year-5);
//	}else {
//		gzY = toGanZhi(year-4);
//	}

//		//月柱 1900年1月小寒以前为 丙子月(60进制12)
//	var firstNode 	= getTerm(y,(m*2-1));//返回当月「节」为几日开始
//	var secondNode = getTerm(y,(m*2));//返回当月「节」为几日开始

//	//依据12节气修正干支月
//	var gzM 	= 	toGanZhi((y-1900)*12+m+11);
//	if(d>=firstNode) {
//		gzM 	= 	toGanZhi((y-1900)*12+m+12);
//	}

//		//传入的日期的节气与否
//	var isTerm = false;
//	var Term = null;
//	if(firstNode==d) {
//		isTerm 	= true;
//		Term 	= solarTerm[m*2-2];
//	}
//	if(secondNode==d) {
//		isTerm 	= true;
//		Term 	= solarTerm[m*2-1];
//	}
//		//日柱 当月一日与 1900/1/1 相差天数
//	var dayCyclical = Date.UTC(y,sm,1,0,0,0,0)/86400000+25567+10;
//	var gzD = toGanZhi(dayCyclical+d-1);

//	return {'lYear':year,'lMonth':month,'lDay':day,'Animal':calendar.getAnimal(year),'IMonthCn':(isLeap?"\u95f0":'')+calendar.toChinaMonth(month),'IDayCn':calendar.toChinaDay(day),'cYear':y,'cMonth':m,'cDay':d,'gzYear':gzY,'gzMonth':gzM,'gzDay':gzD,'isToday':isToday,'isLeap':isLeap,'nWeek':nWeek,'ncWeek':"\u661f\u671f"+cWeek,'isTerm':isTerm,'Term':Term}
//}
//
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
	return dTGinfo[y%10] + dDZinfo[y%12]
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
	return dTGinfo[num%10] + dDZinfo[num%12]
}
