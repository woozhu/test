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
func (sx *SiXiang) String() string{
   str:=fmt.Sprintf("id=[%d] cname=[%s]",sx.id,sx.cname)
   return str
}
//构造函数
func newSiXiang(id int,name string,cname string) SiXiang{
    return SiXiang{id=id,name=name,cname=cname}
}
//集成taiji的xchange接口
func (sx *SiXiang) xchange() SiXiang{
    switch sx.id{
        case 0{
            return newSiXiang(1,"ShaoYang","少阳")
        }
        case 1{
            return newSiXiang(2,"LaoYang","老阳")
        }
        case 2{
            return newSiXiang(3,"ShaoYin","少阴")
        }
        case 3{
            return newSiXiang(0,"LaoYin","老阴")
        }
        default return nil
    }
}

type WuXing struct{
    id int
    name string
    cname string
}
func (wx *WuXing) String() string{
   str:=fmt.Sprintf("id=[%d] cname=[%s]",wx.id,wx.cname)
   return str
}
//构造函数
func newWuXing(id int,name string,cname string) WuXing{
    return WuXing{id=id,name=name,cname=cname}
}
//集成taiji的xchange接口
func (wx *WuXing) xchange() WuXing{
    switch wx.id{
        case 0{
            return newSiXiang(1,"ShaoYang","少阳")
        }
        case 1{
            return newSiXiang(2,"LaoYang","老阳")
        }
        case 2{
            return newSiXiang(3,"ShaoYin","少阴")
        }
        case 3{
            return newSiXiang(0,"LaoYin","老阴")
        }
        default return nil
    }
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
