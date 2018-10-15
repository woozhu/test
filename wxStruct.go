package wxXunHuan

import "fmt"

type taiji interface {
	xchange() interface{}
}
type YinYang struct {
	id    int
	name  string
	cname string
}

func (y *YinYang) String() string {
	str := fmt.Sprintf("id=[%d] cname=[%s]", y.id, y.cname)
	return str
}

//构造函数
func newYinYang(id int, name string, cname string) YinYang {
	return YinYang{id: id, name: name, cname: cname}
}

//集成taiji的xchange接口
func (y *YinYang) xchange() YinYang {
	if y.id == 0 {
		return newYinYang(1, "Yang", "阳")
	} else {
		return newYinYang(0, "Yin", "阴")
	}
}

var (
	yin     = newYinYang(0, "Yin", "阴")
	yang    = newYinYang(1, "Yang", "阳")
	yinyang = []YinYang{yin, yang}
)

type SiXiang struct {
	id    int
	name  string
	cname string
}

func (sx *SiXiang) String() string {
	str := fmt.Sprintf("id=[%d] cname=[%s]", sx.id, sx.cname)
	return str
}

//构造函数
func newSiXiang(id int, name string, cname string) SiXiang {
	return SiXiang{id: id, name: name, cname: cname}
}

//集成taiji的xchange接口
func (sx *SiXiang) xchange() SiXiang {
	switch sx.id {
	case 0:
		{
			return newSiXiang(1, "ShaoYang", "少阳")
		}
	case 1:
		{
			return newSiXiang(2, "LaoYang", "老阳")
		}
	case 2:
		{
			return newSiXiang(3, "ShaoYin", "少阴")
		}
	case 3:
		{
			return newSiXiang(0, "LaoYin", "老阴")
		}
	}
	return SiXiang{}
}

var (
	laoyin   = newSiXiang(0, "LaoYin", "老阴")
	shaoyang = newSiXiang(1, "ShaoYang", "少阳")
	laoyang  = newSiXiang(2, "LaoYang", "老阳")
	shaoyin  = newSiXiang(3, "ShaoYin", "少阴")
	sixiang  = []SiXiang{laoyin, shaoyang, laoyang, shaoyin, laoyin}
)

type WuXing struct {
	id    int
	name  string
	cname string
}

func (wx *WuXing) String() string {
	str := fmt.Sprintf("id=[%d] cname=[%s]", wx.id, wx.cname)
	return str
}

//构造函数
func newWuXing(id int, name string, cname string) WuXing {
	return WuXing{id: id, name: name, cname: cname}
}

//集成taiji的xchange接口
func (wx *WuXing) xchange() WuXing {
	switch wx.id {
	case 0:
		{
			return newWuXing(1, "Huo", "火")
		}
	case 1:
		{
			return newWuXing(2, "Tu", "土")
		}
	case 2:
		{
			return newWuXing(3, "Jin", "金")
		}
	case 3:
		{
			return newWuXing(4, "Shui", "水")
		}
	case 4:
		{
			return newWuXing(0, "Mu", "木")
		}
	}
	return WuXing{}
}

var (
	mu     = newWuXing(0, "Mu", "木")
	huo    = newWuXing(1, "Huo", "火")
	tu     = newWuXing(2, "Tu", "土")
	jin    = newWuXing(3, "Jin", "金")
	shui   = newWuXing(4, "Shui", "水")
	wuxing = []WuXing{mu, huo, tu, jin, shui}
)

type TianGan struct {
	id    int
	name  string
	cname string
	wx    WuXing
	yy    YinYang
}

func (tg *TianGan) String() string {
	str := fmt.Sprintf("id=[%d] cname=[%s]", tg.id, tg.cname)
	return str
}

//构造函数
func newTianGan(id int, name string, cname string) TianGan {
	tgwx := WuXing{}
	tgyy := YinYang{}
	switch id {
	case 0:
		{
			tgwx = mu
			tgyy = yang
		}
	case 1:
		{
			tgwx = mu
			tgyy = yin
		}
	case 2:
		{
			tgwx = huo
			tgyy = yang
		}
	case 3:
		{
			tgwx = huo
			tgyy = yin
		}
	case 4:
		{
			tgwx = tu
			tgyy = yang
		}
	case 5:
		{
			tgwx = tu
			tgyy = yin
		}
	case 6:
		{
			tgwx = jin
			tgyy = yang
		}
	case 7:
		{
			tgwx = jin
			tgyy = yin
		}
	case 8:
		{
			tgwx = shui
			tgyy = yang
		}
	case 9:
		{
			tgwx = shui
			tgyy = yin
		}
	}
	return TianGan{id: id, name: name, cname: cname, wx: tgwx, yy: tgyy}
}

var (
	tgJia   = newTianGan(0, "Jia", "甲")
	tgYi    = newTianGan(1, "Yi", "乙")
	tgBing  = newTianGan(2, "Bing", "丙")
	tgDing  = newTianGan(3, "Ding", "丁")
	tgWu    = newTianGan(4, "Wu", "戊")
	tgJi    = newTianGan(5, "Ji", "己")
	tgGeng  = newTianGan(6, "Geng", "庚")
	tgXin   = newTianGan(7, "Xin", "辛")
	tgRen   = newTianGan(8, "Ren", "壬")
	tgGui   = newTianGan(9, "Gui", "癸")
	tiangan = []TianGan{tgJia, tgYi, tgBing, tgDing, tgWu, tgJi, tgGeng, tgXin, tgRen, tgGui}
)

type DiZhi struct {
	id      int
	name    string
	cname   string
	wx      WuXing
	yy      YinYang
	canggan []TianGan
}

func (dz *DiZhi) String() string {
	str := fmt.Sprintf("id=[%d] cname=[%s]", dz.id, dz.cname)
	return str
}
func (dz *DiZhi) GetShengXiao() string {
	switch dz.id {
	case 0:
		{
			return "鼠"
		}
	case 1:
		{
			return "牛"
		}
	case 2:
		{
			return "虎"
		}
	case 3:
		{
			return "兔"
		}
	case 4:
		{
			return "龙"
		}
	case 5:
		{
			return "蛇"
		}
	case 6:
		{
			return "马"
		}
	case 7:
		{
			return "羊"
		}
	case 8:
		{
			return "猴"
		}
	case 9:
		{
			return "鸡"
		}
	case 10:
		{
			return "狗"
		}
	case 11:
		{
			return "猪"
		}
	}
	return ""
}
func newDiZhi(id int, name string, cname string) DiZhi {
	dzwx := WuXing{}
	dzyy := YinYang{}
	dzcg := []TianGan{}
	switch id {
	case 0:
		{
			dzwx = shui
			dzyy = yang
			dzcg = []TianGan{tgGui}
		}
	case 1:
		{
			dzwx = tu
			dzyy = yin
			dzcg = []TianGan{tgJi, tgGui, tgXin}
		}
	case 2:
		{
			dzwx = mu
			dzyy = yang
			dzcg = []TianGan{tgJia, tgBing, tgWu}
		}
	case 3:
		{
			dzwx = mu
			dzyy = yin
			dzcg = []TianGan{tgYi}
		}
	case 4:
		{
			dzwx = tu
			dzyy = yang
			dzcg = []TianGan{tgWu, tgJi, tgGui}
		}
	case 5:
		{
			dzwx = huo
			dzyy = yin
			dzcg = []TianGan{tgBing, tgWu, tgGeng}
		}
	case 6:
		{
			dzwx = huo
			dzyy = yang
			dzcg = []TianGan{tgDing, tgJi}
		}
	case 7:
		{
			dzwx = tu
			dzyy = yin
			dzcg = []TianGan{tgJi, tgDing, tgYi}
		}
	case 8:
		{
			dzwx = jin
			dzyy = yang
			dzcg = []TianGan{tgGeng, tgRen, tgWu}
		}
	case 9:
		{
			dzwx = jin
			dzyy = yin
			dzcg = []TianGan{tgXin}
		}
	case 10:
		{
			dzwx = tu
			dzyy = yang
			dzcg = []TianGan{tgWu, tgXin, tgDing}
		}
	case 11:
		{
			dzwx = shui
			dzyy = yin
			dzcg = []TianGan{tgRen, tgJia}
		}
	}
	return DiZhi{id: id, name: name, cname: cname, wx: dzwx, yy: dzyy, canggan: dzcg}
}

var (
	dzZi   = newDiZhi(0, "Zi", "子")
	dzChou = newDiZhi(1, "Chou", "丑")
	dzYin  = newDiZhi(2, "Yin", "寅")
	dzMao  = newDiZhi(3, "Mao", "卯")
	dzChen = newDiZhi(4, "Chen", "辰")
	dzSi   = newDiZhi(5, "Si", "巳")
	dzWu   = newDiZhi(6, "Wu", "午")
	dzWei  = newDiZhi(7, "Wei", "未")
	dzShen = newDiZhi(8, "Shen", "申")
	dzYou  = newDiZhi(9, "You", "酉")
	dzXu   = newDiZhi(10, "Xu", "戌")
	dzHai  = newDiZhi(11, "Hai", "亥")
	dizhi  = []DiZhi{dzZi, dzChou, dzYin, dzMao, dzChen, dzSi, dzWu, dzWei, dzShen, dzYou, dzXu, dzHai}
)

//干支纪年法，六十甲子
type GanZhi struct {
	id int
	tg TianGan
	dz DiZhi
}

//干支的id是从0到59的一个循环。
func newGanZhi(id int) GanZhi {
	var (
		gz = GanZhi{}
	)
	gz.id = id
	for _, t := range tiangan {
		if t.id == id%10 {
			gz.tg = t
		}
	}
	for _, d := range dizhi {
		if d.id == id%12 {
			gz.dz = d
		}
	}
	return gz
}
func (gz *GanZhi) String() string {
	str := fmt.Sprintf("id=[%d] cname=[%s%s]", gz.id, gz.tg.cname, gz.dz.cname)
	return str
}

var (
	gzJiaZi    = newGanZhi(0)
	gzYiChou   = newGanZhi(1)
	gzBingYin  = newGanZhi(2)
	gzDingMao  = newGanZhi(3)
	gzWuChen   = newGanZhi(4)
	gzJiSi     = newGanZhi(5)
	gzGengWu   = newGanZhi(6)
	gzXinWei   = newGanZhi(7)
	gzRenShen  = newGanZhi(8)
	gzGuiYou   = newGanZhi(9)
	gzJiaXu    = newGanZhi(10)
	gzYiHai    = newGanZhi(11)
	gzBingZi   = newGanZhi(12)
	gzDingChou = newGanZhi(13)
	gzWuYin    = newGanZhi(14)
	gzJiMao    = newGanZhi(15)
	gzGengChen = newGanZhi(16)
	gzXinSi    = newGanZhi(17)
	gzRenWu    = newGanZhi(18)
	gzGuiWei   = newGanZhi(19)
	gzJiaShen  = newGanZhi(20)
	gzYiYou    = newGanZhi(21)
	gzBingXu   = newGanZhi(22)
	gzDingHai  = newGanZhi(23)
	gzWuZi     = newGanZhi(24)
	gzJiChou   = newGanZhi(25)
	gzGengYin  = newGanZhi(26)
	gzXinMao   = newGanZhi(27)
	gzRenChen  = newGanZhi(28)
	gzGuiSi    = newGanZhi(29)
	gzJiaWu    = newGanZhi(30)
	gzYiWei    = newGanZhi(31)
	gzBingShen = newGanZhi(32)
	gzDingYou  = newGanZhi(33)
	gzWuXu     = newGanZhi(34)
	gzJiHai    = newGanZhi(35)
	gzGengZi   = newGanZhi(36)
	gzXinChou  = newGanZhi(37)
	gzRenYin   = newGanZhi(38)
	gzGuiMao   = newGanZhi(39)
	gzJiaChen  = newGanZhi(40)
	gzYiSi     = newGanZhi(41)
	gzBingWu   = newGanZhi(42)
	gzDingWei  = newGanZhi(43)
	gzWuShen   = newGanZhi(44)
	gzJiYou    = newGanZhi(45)
	gzGengXu   = newGanZhi(46)
	gzXinHai   = newGanZhi(47)
	gzRenZi    = newGanZhi(48)
	gzGuiChou  = newGanZhi(49)
	gzJiaYin   = newGanZhi(50)
	gzYiMao    = newGanZhi(51)
	gzBingChen = newGanZhi(52)
	gzDingSi   = newGanZhi(53)
	gzWuWu     = newGanZhi(54)
	gzJiWei    = newGanZhi(55)
	gzGengShen = newGanZhi(56)
	gzXinYou   = newGanZhi(57)
	gzRenXu    = newGanZhi(58)
	gzGuiHai   = newGanZhi(59)
	ganzhi     = []GanZhi{
		gzJiaZi, gzYiChou, gzBingYin, gzDingMao, gzWuChen, gzJiSi, gzGengWu, gzXinWei, gzRenShen, gzGuiYou,
		gzJiaXu, gzYiHai, gzBingZi, gzDingChou, gzWuYin, gzJiMao, gzGengChen, gzXinSi, gzRenWu, gzGuiWei,
		gzJiaShen, gzYiYou, gzBingXu, gzDingHai, gzWuZi, gzJiChou, gzGengYin, gzXinMao, gzRenChen, gzGuiSi,
		gzJiaWu, gzYiWei, gzBingShen, gzDingYou, gzWuXu, gzJiHai, gzGengZi, gzXinChou, gzRenYin, gzGuiMao,
		gzJiaChen, gzYiSi, gzBingWu, gzDingWei, gzWuShen, gzJiYou, gzGengXu, gzXinHai, gzRenZi, gzGuiChou,
		gzJiaYin, gzYiMao, gzBingChen, gzDingSi, gzWuWu, gzJiWei, gzGengShen, gzXinYou, gzRenXu, gzGuiHai}
)

type NianZhu struct {
	id    int
	name  string
	cname string
	tg    TianGan
	dz    DiZhi
}
type YueZhu struct {
	id    int
	name  string
	cname string
	tg    TianGan
	dz    DiZhi
}

type RiZhu struct {
	id    int
	name  string
	cname string
	tg    TianGan
	dz    DiZhi
}

type ShiZhu struct {
	id    int
	name  string
	cname string
	tg    TianGan
	dz    DiZhi
}

type SiZhu struct {
	nz NianZhu
	yz YueZhu
	rz RiZhu
	sz ShiZhu
}
