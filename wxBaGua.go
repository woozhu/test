package wxBaGua

import “strconv”

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