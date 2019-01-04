package wxBaGua

import “strconv”

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

//列表转化为字符串
func List2String(sxl []int) string{
    var s string
    for _,i:= range(sxl){
        s += strconv.ItoA(i)
    }
    return s
}
//字符串转化为列表
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
//对数序处理
func GetGongGua(sxl []int) []int {
    s:= List2Int(sxl)
    var r = []int{}
    for i,j:=range(64Gua){
        for _,k:=range(j){
            if s==k{r=Int2List(i)}
        }
    }
    return r
}
//对数序处理
func GetBenGua(sxl []int) []int {
    var r= []int{}
    for _,i:=range(sxl){
        r = append(r,i%2)
    }
    return r
}
//对数序处理
func GetBianGua(sxl []int) []int {
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