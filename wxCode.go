package wxXunHuan

//"fmt"

var (
	//[子-亥]月的五行计分系数{木，火，土，金，水}
	monthWXk = [12][5]int{
		{1414, 500, 707, 1000, 2000},  //【子】月的五行分数调整系数
		{898, 821, 1512, 1348, 1041},  //【丑】月的五行分数调整系数
		{1571, 1548, 924, 716, 862},   //【寅】月的五行分数调整系数
		{2000, 1414, 500, 707, 1000},  //【卯】月的五行分数调整系数
		{1166, 1074, 1421, 1161, 800}, //【辰】月的五行分数调整系数
		{861, 1571, 1548, 924, 716},   //【巳】月的五行分数调整系数
		{912, 1700, 1590, 774, 645},   //【午】月的五行分数调整系数
		{924, 1340, 1674, 1069, 612},  //【未】月的五行分数调整系数
		{795, 674, 1012, 1641, 1498},  //【申】月的五行分数调整系数
		{500, 707, 1000, 2000, 1414},  //【酉】月的五行分数调整系数
		{674, 1012, 1641, 1498, 795},  //【戌】月的五行分数调整系数
		{1590, 774, 645, 912, 1700}}   //【亥】月的五行分数调整系数

	//天干中的五行分数，地址及地址藏干中的五行[木，火，土，金，水]分数
	tgcode = [10][5]int{
		{10, 0, 0, 0, 0}, //天干【甲】的五行分数
		{10, 0, 0, 0, 0}, //天干【乙】的五行分数
		{0, 10, 0, 0, 0}, //天干【丙】的五行分数
		{0, 10, 0, 0, 0}, //天干【丁】的五行分数
		{0, 0, 10, 0, 0}, //天干【戊】的五行分数
		{0, 0, 10, 0, 0}, //天干【己】的五行分数
		{0, 0, 0, 10, 0}, //天干【庚】的五行分数
		{0, 0, 0, 10, 0}, //天干【辛】的五行分数
		{0, 0, 0, 0, 10}, //天干【壬】的五行分数
		{0, 0, 0, 0, 10}} //天干【癸】的五行分数
	//十二地支【子-亥】在十二个月份【子-亥】中的藏干分数
	dzcgcode = [12][12][5]int{
		{{0, 0, 0, 0, 10}, {0, 0, 0, 0, 10}, {0, 0, 0, 0, 10}, {0, 0, 0, 0, 10}, {0, 0, 0, 0, 10}, {0, 0, 0, 0, 10}, {0, 0, 0, 0, 10}, {0, 0, 0, 0, 10}, {0, 0, 0, 0, 10}, {0, 0, 0, 0, 10}, {0, 0, 0, 0, 10}, {0, 0, 0, 0, 10}}, //地支【子】在12个月【子-亥】中的藏干分数
		{{0, 0, 7, 0, 3}, {0, 0, 7, 1, 2}, {0, 0, 6, 2, 2}, {0, 0, 6, 2, 2}, {0, 0, 6, 2, 2}, {0, 0, 6, 2, 2}, {0, 0, 6, 2, 2}, {0, 0, 6, 2, 2}, {0, 0, 4, 3, 3}, {0, 0, 4, 3, 3}, {0, 0, 6, 2, 2}, {0, 0, 7, 0, 3}},             //地支【丑】在12个月【子-亥】中的藏干分数
		{{7, 2, 1, 0, 0}, {7, 2, 1, 0, 0}, {7, 2, 1, 0, 0}, {7, 2, 1, 0, 0}, {7, 2, 1, 0, 0}, {7, 2, 1, 0, 0}, {7, 2, 1, 0, 0}, {7, 2, 1, 0, 0}, {7, 2, 1, 0, 0}, {7, 2, 1, 0, 0}, {7, 2, 1, 0, 0}, {7, 2, 1, 0, 0}},             //地支【寅】在12个月【子-亥】中的藏干分数
		{{10, 0, 0, 0, 0}, {10, 0, 0, 0, 0}, {10, 0, 0, 0, 0}, {10, 0, 0, 0, 0}, {10, 0, 0, 0, 0}, {10, 0, 0, 0, 0}, {10, 0, 0, 0, 0}, {10, 0, 0, 0, 0}, {10, 0, 0, 0, 0}, {10, 0, 0, 0, 0}, {10, 0, 0, 0, 0}, {10, 0, 0, 0, 0}}, //地支【卯】在12个月【子-亥】中的藏干分数
		{{4, 0, 0, 0, 6}, {2, 0, 6, 0, 2}, {2, 0, 7, 0, 1}, {2, 0, 7, 0, 1}, {2, 0, 7, 0, 1}, {2, 0, 7, 0, 1}, {2, 0, 7, 0, 1}, {2, 0, 7, 0, 1}, {2, 0, 5, 0, 3}, {2, 0, 5, 0, 3}, {2, 0, 7, 0, 1}, {4, 0, 0, 0, 6}},             //地支【辰】在12个月【子-亥】中的藏干分数
		{{0, 7, 2, 0, 1}, {0, 7, 2, 0, 1}, {0, 7, 2, 0, 1}, {0, 7, 2, 0, 1}, {0, 7, 2, 0, 1}, {0, 7, 2, 0, 1}, {0, 7, 2, 0, 1}, {0, 7, 2, 0, 1}, {0, 7, 2, 0, 1}, {0, 7, 2, 0, 1}, {0, 7, 2, 0, 1}, {0, 7, 2, 0, 1}},             //地支【巳】在12个月【子-亥】中的藏干分数
		{{0, 7, 3, 0, 0}, {0, 7, 3, 0, 0}, {0, 7, 3, 0, 0}, {0, 7, 3, 0, 0}, {0, 7, 3, 0, 0}, {0, 7, 3, 0, 0}, {0, 7, 3, 0, 0}, {0, 7, 3, 0, 0}, {0, 7, 3, 0, 0}, {0, 7, 3, 0, 0}, {0, 7, 3, 0, 0}, {0, 7, 3, 0, 0}},             //地支【午】在12个月【子-亥】中的藏干分数
		{{2, 3, 5, 0, 0}, {2, 3, 5, 0, 0}, {2, 3, 5, 0, 0}, {2, 3, 5, 0, 0}, {2, 3, 5, 0, 0}, {0, 6, 4, 0, 0}, {0, 6, 4, 0, 0}, {0, 6, 4, 0, 0}, {2, 3, 5, 0, 0}, {2, 3, 5, 0, 0}, {0, 5, 5, 0, 0}, {2, 3, 5, 0, 0}},             //地支【未】在12个月【子-亥】中的藏干分数
		{{0, 0, 1, 7, 2}, {0, 0, 1, 7, 2}, {0, 0, 1, 7, 2}, {0, 0, 1, 7, 2}, {0, 0, 1, 7, 2}, {0, 0, 1, 7, 2}, {0, 0, 1, 7, 2}, {0, 0, 1, 7, 2}, {0, 0, 1, 7, 2}, {0, 0, 1, 7, 2}, {0, 0, 1, 7, 2}, {0, 0, 1, 7, 2}},             //地支【申】在12个月【子-亥】中的藏干分数
		{{0, 0, 0, 10, 0}, {0, 0, 0, 10, 0}, {0, 0, 0, 10, 0}, {0, 0, 0, 10, 0}, {0, 0, 0, 10, 0}, {0, 0, 0, 10, 0}, {0, 0, 0, 10, 0}, {0, 0, 0, 10, 0}, {0, 0, 0, 10, 0}, {0, 0, 0, 10, 0}, {0, 0, 0, 10, 0}, {0, 0, 0, 10, 0}}, //地支【酉】在12个月【子-亥】中的藏干分数
		{{0, 1, 7, 2, 0}, {0, 1, 7, 2, 0}, {0, 3, 5, 2, 0}, {0, 3, 5, 2, 0}, {0, 2, 5, 3, 0}, {0, 6, 4, 0, 0}, {0, 6, 4, 0, 0}, {0, 6, 4, 0, 0}, {0, 2, 5, 3, 0}, {0, 2, 5, 3, 0}, {0, 5, 5, 0, 0}, {0, 1, 7, 2, 0}},             //地支【戌】在12个月【子-亥】中的藏干分数
		{{3, 0, 0, 0, 7}, {3, 0, 0, 0, 7}, {3, 0, 0, 0, 7}, {3, 0, 0, 0, 7}, {3, 0, 0, 0, 7}, {3, 0, 0, 0, 7}, {3, 0, 0, 0, 7}, {3, 0, 0, 0, 7}, {3, 0, 0, 0, 7}, {3, 0, 0, 0, 7}, {3, 0, 0, 0, 7}, {3, 0, 0, 0, 7}}}             //地支【亥】在12个月【子-亥】中的藏干分数
)

type wxCode struct {
	code []int
}

//s string "甲子"
func newCode() *wxCode {
	wc := new(wxCode)
	for i := 0; i < 5; {
		wc.code[i] = 0
		i++
	}
	return wc
}

func (wc *wxCode) caculateRiZhu(ygz GanZhi, rgz GanZhi) {
	for i, ls := range tgcode {
		if rgz.tg.id == i {
			for j, tgl := range ls {
				wc.code[j] = tgl
			}
		}
	}
	for i, ls := range dzcgcode {
		if rgz.dz.id == i {
			for j, dzl := range ls {
				if ygz.dz.id == j {
					for a, b := range dzl {
						wc.code[a] = b
					}
				}
			}
		}
	}
}
func (wc *wxCode) caculateShiZhu(ygz GanZhi, sgz GanZhi) {
	wc.caculateRiZhu(ygz, sgz)
}
func (wc *wxCode) caculateYueZhu(ygz GanZhi) {
	for i, ls := range tgcode {
		if ygz.tg.id == i {
			for j, tgl := range ls {
				wc.code[j] = tgl
			}
		}
	}
	for i, ls := range dzcgcode {
		if ygz.dz.id == i {
			for j, dzl := range ls {
				if ygz.dz.id == j {
					for a, b := range dzl {
						wc.code[a] = b
					}
				}
			}
		}
	}
}
func (wc *wxCode) caculateNianZhu(ygz GanZhi, ngz GanZhi) {
	wc.caculateRiZhu(ygz, ngz)
}
func (wc *wxCode) caculateSiZhu(sz wxSiZhu) {
	wc.caculateRiZhu(sz.yuezhu, sz.rizhu)
	wc.caculateShiZhu(sz.yuezhu, sz.shizhu)
	wc.caculateYueZhu(sz.yuezhu)
	wc.caculateNianZhu(sz.yuezhu, sz.nianzhu)
}
