package wxXunHuan

type ZiWeiGong struct{
    id int
    name string
    cname string
    tg TianGan
    dz Dizhi
    daXian []int
    xiaoXian []int
    zwStars []ZiWeiXing
    
}
type ZiWeiStars struct{
    id int
    name string
    cnmae string
    ziweiGong ZiWeiGong
    wangShuan []string
}

type ZiWei struct{
    sizhu SiZhu
    ziweiList ZiWeiList
}
