package wxXunHuan

import (
	"fmt"
	//	"time"
)

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
func newYinYang(id int) YinYang {
	return YinYang{id: id, name: dYYnameList[id], cname: dYYcnameList[id]}
}

//集成taiji的xchange接口
func (y *YinYang) xchange() YinYang {
	return yinyang[(y.id+1)%2]
}

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
func newSiXiang(id int) SiXiang {
	return SiXiang{id: id, name: dSXnameList[id], cname: dSXcnameList[id]}
}

//集成taiji的xchange接口
func (sx *SiXiang) xchange() SiXiang {
	return sixiang[(sx.id+1)%4]
}

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
	return wuxing[(wx.id+1)%5]
}

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
func newTianGan(id int) TianGan {
	tgwx := wuxing[((id-id%2)/2)%5]
	tgyy := yinyang[(id+1)%2]
	return TianGan{id: id, name: dTGnameList[id], cname: dTGcnameList[id], wx: tgwx, yy: tgyy}
}

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
	return Animals[dz.id]
}
func newDiZhi(id int) DiZhi {
	dzwx := WuXing{}
	dzyy := yinyang[(id+1)%2]
	dzcg := []TianGan{}
	switch id {
	case 0:
		{
			dzwx = shui
			dzcg = []TianGan{tgGui}
		}
	case 1:
		{
			dzwx = tu
			dzcg = []TianGan{tgJi, tgGui, tgXin}
		}
	case 2:
		{
			dzwx = mu
			dzcg = []TianGan{tgJia, tgBing, tgWu}
		}
	case 3:
		{
			dzwx = mu
			dzcg = []TianGan{tgYi}
		}
	case 4:
		{
			dzwx = tu
			dzcg = []TianGan{tgWu, tgJi, tgGui}
		}
	case 5:
		{
			dzwx = huo
			dzcg = []TianGan{tgBing, tgWu, tgGeng}
		}
	case 6:
		{
			dzwx = huo
			dzcg = []TianGan{tgDing, tgJi}
		}
	case 7:
		{
			dzwx = tu
			dzcg = []TianGan{tgJi, tgDing, tgYi}
		}
	case 8:
		{
			dzwx = jin
			dzcg = []TianGan{tgGeng, tgRen, tgWu}
		}
	case 9:
		{
			dzwx = jin
			dzcg = []TianGan{tgXin}
		}
	case 10:
		{
			dzwx = tu
			dzcg = []TianGan{tgWu, tgXin, tgDing}
		}
	case 11:
		{
			dzwx = shui
			dzcg = []TianGan{tgRen, tgJia}
		}
	}
	return DiZhi{id: id, name: dDZnameList[id], cname: dDZcnameList[id], wx: dzwx, yy: dzyy, canggan: dzcg}
}

//干支纪年法，六十甲子
type GanZhi struct {
	id int
	//cname string
	tg TianGan
	dz DiZhi
}

//干支的id是从0到59的一个循环。
func newGanZhi(id int) GanZhi {
	return GanZhi{id: id, tg: tiangan[id%10], dz: dizhi[id%12]}
}
func (gz *GanZhi) String() string {
	str := fmt.Sprintf("id=[%d] cname=[%s%s]", gz.id, gz.tg.cname, gz.dz.cname)
	return str
}
