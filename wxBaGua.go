package wxBaGua

import “strconv”

//日辰空亡!!!
func GetRiChenXunKongDiZhiIndexList(rc []int) []int{
    var r=[]int{}
    var index int
    var tg=rc[0]
    var dz=rc[1]
    k1:= (dz+(9-tg)%10+1)%12
    k2:= (k1+1)%12
    r=append(r,k1)
    r=append(r,k2)
    return r
}
//月破!!!
func GetYuePoDiZhiIndex(yc []int) int{
    return (yc[1]+6)%12
}
//付神oooo
func GetFuShenIndexList(sxl []int) []int{
    
}
//飞神oooo
func GetFeiShenIndexList(sxl[]int) []int{
    
}

//六亲之爻序!!!
func GetYaoIndexByGuaLiuQinIndex(lqindex int,lqindexlist []int)int{
    var r int
    for i,j:=range lqindexlist{
        if j==lqindex{r=i}
    }
    return r
}

//五行之爻序!!!
func GetYaoIndexByWuXingIndex(wxindex int,wxindexlist []int) int{
    var r int
    for i,j:=range wxindexlist{
        if j== wxindex {r=i}
    }
    return r
}

//六神之爻序!!!
func GetYaoIndexByLiuShenIndex(lsindex int,lsindexlist []int) int{
    var r int
    for i,j:=range lsindexlist{
        if j==lsindex{r=i}
    }
    return r
}

//选项卡模式

type LiuYaoOptions struct {
    optionStr1 string
    optionStr2 string
    optionInt1 int
    optionInt2 int
}

var defaultLiuYaoOptions = LiuYaoOptions{
    optionStr1: "defaultStr1",
    optionStr2: "defaultStr2",
    optionInt1: 1,
    optionInt2: 2,
}

func LiuYao(requiredStr string, opts ...LiuYaoOption) {
    options := defaultLiuYaoOptions
    for _, o := range opts {
        o(&options)
    }

    fmt.Println(requiredStr, options.optionStr1, options.optionStr2, options.optionInt1, options.optionInt2)
}

type LiuYaoOption func(options *LiuYaoOptions)

func WithOptionStr1(str1 string) LiuYaoOption {
    return func(options *LiuYaoOptions) {
        options.optionStr1 = str1
    }
}

func WithOptionInt1(int1 int) LiuYaoOption {
    return func(options *LiuYaoOptions) {
        options.optionInt1 = int1
    }
}

func WithOptionStr2AndInt2(str2 string, int2 int) LiuYaoOption {
    return func(options *LiuYao。Options) {
        options.optionStr2 = str2
        options.optionInt2 = int2
    }
}


LiuYao("requiredStr")
LiuYao("requiredStr", WithOptionStr1("mystr1"))
LiuYao("requiredStr", WithOptionStr2AndInt2("mystr2", 22), WithOptionInt1(11))

//g关于数序的一些转化函数。

var ShuXuList=[]int{}

//列表转化为字符串!!!
func List2String(sxl []int) string{
    var s string
    for _,i:= range(sxl){
        s += strconv.ItoA(i)
    }
    return s
}
//字符串转化为列表!!!
func String2List(s string) []int{
    var l=[]int{}
    for _,i:=range(s){
         l= append(l,strconv.AtoI(i))
    }
    return l
}
//列表转化为数字
func List2Int(sxl []int) int{
    s:=List2String(sxl)
    return int(strconv.AtoI(s))
}
//数字转化为列表
func Int2List(i int) []int {
    var l=[]int{}
    for _,ii:=range(strconv.ItoA(i)){
        l = append(l,strconv.AtoI(ii))
    }
    return l
}
//根据卦序数得到宫卦的五行序数！！！
func GetGongGuaWuXingIndex(sxl []int)int{
    s:=List2String(GetBenGuaYinYangIndexList(sxl))
    var r int
    var rr int
    for i,j:=range(GuaIndexStringList){
        for _,k:= range(j){
            if s==k{r=i}
        }
    }
    switch{
        case r==0||r==7:{rr=3}
        case r==1||r==5:{rr=0}
        case r==2:{rr=4}
        case r==3||r==4:{rr=2}
        case r==6:{rr=1}
    }
    return rr
}
//根据卦序数得到宫卦阴阳序数！！！
func GetGongGuaYinYangIndexList(sxl []int) []int {
    s:= List2String(GetBenGuaYinYangIndexList(sxl))
    var rs string
    for i,j:=range(GuaIndexStringList){
        for _,k:=range(j){
            if s==k{rs=j[0]}
        }
    }
    return String2List(rs)
}
//八宫卦的上下干支序数列表
var(
    QianShangGZIndexList =[]int{}
    QianXiaGzIndexList =[]int{}
    XunShangGZIndexList=[]int{}
    XunXiaGZIndexList=[]int{}
    KanShangGZIndexList=[]int{}
    KanXiaGZIndexList=[]int{}
    GenShangGZIndexList=[]int{}
    GenXiaGZIndexList=[]int{}
    KunShangGZIndexList=[]int{}
    KunXiaGZIndexList=[]int{}
    ZhenShangGZIndexList=[]int{}
    ZhenXiaGZIndexList=[]int{}
    LiShangGZIndexList=[]int{}
    LiXiaGZIndexList=[]int{}
    DuiShang
)

//乾宫卦序字符串列表。
//初爻在前，六爻在末位。!!!
var QianGongisl=[]string{
    “111111”,
    “011111”,
    “001111”,
    “000111”,
    “000011”,
    “000001”,
    “000101”,
    “111101”,
}
var XunGongisl=[]string{
    “011011”,
    “111011”,
    “101011”,
    “100011”,
    “100111”,
    “100101”,
    “100001”,
    “011001”,
}
var KanGongisl=[]string{
    “010010”,
    “110010”,
    “100010”,
    “101010”,
    “101110”,
    “101100”,
    “101000”,
    “010000”
}
var GenGongisl=[]string{
    “001001”,
    “101001”,
    “111001”,
    “110001”,
    “110101”,
    “110111”,
    “110011”,
    “001011”,
}
var KunGongisl=[]string{
    “000000”,
    “100000”,
    “110000”,
    “111000”,
    “111100”,
    “111110”,
    “111010”,
    “000010”,
}
var ZhenGongisl=[]string{
    “100100”,
    “000100”,
    “010100”,
    “011100”,
    “011000”,
    “011010”,
    “011110”,
    “100110”,
}
var LiGongisl=[]string{
    “101101”,
    “001101”,
    “011101”,
    “010101”,
    “010001”,
    “010011”,
    “010111”,
    “101111”,
}
var DuiGongisl=[]string{
    “110110”,
    “010110”,
    “000110”,
    “001110”,
    “001010”,
    “001000”,
    “001100”,
    “110100”,
}
var GuaIndexStringList=[][]string{
    QianGongisl,
    XunGongisl,
    KanGongisl,
    GenGongisl,
    KunGongisl,
    ZhenGongisl,
    LiGongisl,
    DuiGongisl,
}
//根据卦序数得到本卦阴阳卦序！！！
func GetBenGuaYinYangIndexList(sxl []int) []int {
    var r= []int{}
    for _,i:=range(sxl){
        r = append(r,i%2)
    }
    return r
}
//根据卦数序得到变卦阴阳卦序！！！
func GetBianGuaYinYangIndexList(sxl []int) []int {
    var r= []int{}
    for _,i:= range(sxl){
        i=i%4
        switch {
        case i%4==0||i%4==1:{r=append(r,1)}
        case i%4==3||i%4==2:{r=append(r,0)}
        }
    }
    return r
}
//根据卦序数得到动爻的卦序列表!!!
func GetGuaDongYaoIndexList(sxl []int){
    var r=[]int{}
    for i,j:=range sxl{
        if j%4==0||j%4==3{
            r=append(r,i)
        }
    }
    return r
}
