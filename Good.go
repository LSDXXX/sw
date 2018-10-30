package shenwu

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

type GoodsType int

const (
	_GoodsType = iota
	PETS
	YIJ
	ERJ
	SANJ
	FURNITURE
	PENGREN
	GUDONG
	TINGYUAN
)

var goods_info map[string]GoodsType

func init() {
	goods_info = map[string]GoodsType{
		"蛇胆":  ERJ,
		"七星草": ERJ,
		"松节":  ERJ,
		"紫冥果": ERJ,
		"血玛瑙": ERJ,
		"叶上珠": ERJ,
		"也白头": ERJ,
		"绿莲散": ERJ,
		"长寿果": ERJ,
		"山长青": ERJ,
		"灵芝":  ERJ,
		"赤石脂": ERJ,
		"紫花兰": ERJ,
		"药神花": ERJ,
		"龙尾骨": ERJ,
		"贝母花": ERJ,
		"吕宋果": ERJ,
		"奇异果": ERJ,
		"麻黄":  ERJ,
		"月月花": ERJ,

		"还灵散":   SANJ,
		"济生丸":   SANJ,
		"玉蟾丸":   SANJ,
		"灵心丸":   SANJ,
		"驱魔丹":   SANJ,
		"血还丹":   SANJ,
		"聚元丹":   SANJ,
		"九转还魂丹": SANJ,
		"金疮药":   SANJ,
		"乾坤丹":   SANJ,
		"逍遥散":   SANJ,

		"女儿红":  PENGREN,
		"踏雪燕窝": PENGREN,
		"珍珠丸子": PENGREN,
		"香辣蟹":  PENGREN,
		"长寿面":  PENGREN,
		"佛跳墙":  PENGREN,
		"叫花鸡":  PENGREN,
		"百味酒":  PENGREN,
		"醉生梦死": PENGREN,
		"玉蓉糕":  PENGREN,
		"脆皮乳猪": PENGREN,
		"八宝粥":  PENGREN,
		"珍露酒":  PENGREN,
	}
}

func BuyPets(name string) bool {
	//OpenTradingPet()
	checkActive()
	x, y := SearchImg(pet_name, ThisGame.Rec)
	if x == -1 {
		return false
	}
	rec := Rect{
		x,
		y + 10,
		160,
		252,
	}
	fmt.Println(SearchImg(pet_need0, rec))
	for i := 0; i < 10; i++ {
		if checkClickFist(pet_need0, pet_need1, rec) {
			break
		}
		robotgo.MoveMouseSmooth(x+180, y+90)
		robotgo.Sleep(1)
		robotgo.ScrollMouse(9, "down")
		robotgo.Sleep(1)
	}
	return true
}

func BuyGoods(t GoodsType, name string) {
}
