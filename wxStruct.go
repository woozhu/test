package wxXunHuan

type taiji interface{
    xchange() interface
}
type YinYang struct{
    id int
    name string
    cname string
}
func (y *YinYang) String() string{
   str:=fmt.Sprintf("id=[%d] cname=[%s]",y.id,y.cname)
   return str
}

func newYinYang(id int,name string,cname string) YinYang{
    return YinYang{id=id,name=name,cname=cname}
}

func (y *YinYang) xchange() YinYang{
    if y.id == 0{
    return newYinYang(1,"Yang","阳")
    }else{
    return newYinYang（0,"Yin","阴"）
    }
}
