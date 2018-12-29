package wxBaGua

import “strconv”

var ShuXuList=[]int{}
func List2String(sxl []int) string{
    var s string
    for _,i:= range(sxl){
        s += strconv.ItoA(i)
    }
    return s
}

func String2List(s string) []int{
    var l=[]int{}
    for _,i:=range(s){
         l= append(l,strconv.AtoI(i))
    }
    return l
}

func List2Int(sxl []int) int{
    s:=List2String(sxl)
    return int(strconv.AtoI(s))
}

func Int2List(i int) []int {
    var l=[]int{}
    for _,ii:=range(strconv.ItoA(i)){
        l = append(l,strconv.AtoI(ii))
    }
    return l
}