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
//构造函数
func newYinYang(id int,name string,cname string) YinYang{
    return YinYang{id=id,name=name,cname=cname}
}
//集成taiji的xchange接口
func (y *YinYang) xchange() YinYang{
    if y.id == 0{
        return newYinYang(1,"Yang","阳")
    }else{
        return newYinYang(0,"Yin","阴")
    }
}                     

type SiXiang struct{
    id int
    name string
    cname string
}

type WuXing struct{
    id int
    name string
    cname string
}

type TianGan struct{
    id int
    name string
    cname string
    wx WuXing
    yy YinYang
}

type DiZhi struct{
    id int
    name string
    cname string
    wx WuXing
    yy YinYang
    canggan []TianGan
}
