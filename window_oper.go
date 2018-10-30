package shenwu

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"strconv"
	"strings"
	"unicode"
	"unsafe"

	"github.com/go-vgo/robotgo"
)

type Rect robotgo.MRect

type Point struct {
	X int
	Y int
}

type Game struct {
	HGame int32
	Rec   Rect
}

type connectInfo struct {
	Method func()
	Dst    *mapInfo
}

type mapInfo struct {
	W    int
	H    int
	R    float32
	Cons map[string]connectInfo
	Fly  func()
}

type TaskState int

const (
	_TaskState = iota
	PATROL
	BUY
	FINDING
	PET
)

type baoTuTaskInfo struct {
	N     int
	S     TaskState
	Scene string
	Pos   Point
	Spos  string
	Goods string
}

const (
	game_name       = "神武3 - 平步青云"
	search_npc      = "sucai\\map_search.bmp"
	task_goal       = "sucai\\renwu_mubiao.bmp"
	task_reply      = "sucai\\renwu_huifu.bmp"
	task_strategy   = "sucai\\renwu_gonglue.bmp"
	task_baotu1     = "sucai\\renwu_baotu0.bmp"
	task_baotu2     = "sucai\\renwu_baotu1.bmp"
	get_baotu       = "sucai\\baotu.bmp"
	close_button    = "sucai\\shutdown_button.bmp"
	task_path       = "sucai\\renwu.bmp"
	baotu_patrol    = "sucai\\baotu_xunluo.bmp"
	check_battle    = "sucai\\check_battle.bmp"
	check_battle0   = "sucai\\battle_button0.bmp"
	check_battle1   = "sucai\\battle_button1.bmp"
	close_auto      = "sucai\\close_auto.bmp"
	friends_        = "sucai\\haoyou.bmp"
	friends_round0  = "sucai\\haoyou_zhouwei0.bmp"
	friends_round1  = "sucai\\haoyou_zhouwei1.bmp"
	round_npc0      = "sucai\\zhouwei_npc0.bmp"
	round_npc1      = "sucai\\zhouwei_npc1.bmp"
	fight_dadangjia = "sucai\\dadangjia_zhandou.bmp"
	task_pet        = "sucai\\aichong.bmp"
	trading_pet     = "sucai\\jiaoyi_pet.bmp"
	trading_goods   = "sucai\\jiaoyi_goods.bmp"
	pet_name        = "sucai\\chongwu.bmp"
	pet_need0       = "sucai\\pet_xu0.bmp"
	pet_need1       = "sucai\\pet_xu1.bmp"
	co_bmp_path     = "game\\co.bmp"
	task_detail     = "game\\detail"
	round_npc       = "game\\round.bmp"
	spet_name       = "game\\pet_name.bmp"
	wait_micsec     = 500
	task_width      = 219
	retry_times     = 2
	fly0            = "f6"
	fly1            = "f7"
	fly2            = "f8"
)

var cordinate_rec Rect

var ThisGame Game

func (t Game) GetEnemyPos(i int) Point {
	x0 := 300
	y0 := 240
	xi := i
	yi := -i
	return Point{
		x0 + xi*75,
		y0 + yi*37,
	}
}

func MoveToNone() {
	checkActive()
	robotgo.MoveMouseSmooth(ThisGame.Rec.Top_x, ThisGame.Rec.Top_y)
}

//得到坐标rect
func (g Game) GetCoRect() Rect {
	return Rect{
		g.Rec.Top_x + cordinate_rec.Top_x,
		g.Rec.Top_y + cordinate_rec.Top_y,
		cordinate_rec.Width,
		cordinate_rec.Height,
	}
}

func init() {
	tmp := robotgo.FindWindow(game_name)
	ThisGame.HGame = *((*int32)(unsafe.Pointer(&tmp)))
	ThisGame.Rec = Rect(robotgo.GetClientRect(int32(ThisGame.HGame)))
	cordinate_rec = Rect{
		124,
		2,
		74,
		18,
	}
}

func Test() robotgo.MRect {
	rec := robotgo.GetClientRect(ThisGame.HGame)
	return rec
}

//激活窗口 游戏关闭或者重新打开返回false其他返回true
func checkActive() bool {
	ac := robotgo.GetActive()
	if *((*int32)(unsafe.Pointer(&ac.HWnd))) == ThisGame.HGame && ThisGame.HGame != 0 {
		rec := robotgo.GetClientRect(int32(ThisGame.HGame))
		if Rect(rec) != ThisGame.Rec {
			ThisGame.Rec = Rect(rec)
		}
		return true
	} else {
		hwnd := robotgo.FindWindow(game_name)
		i := *((*int32)(unsafe.Pointer(&hwnd)))
		if i != 0 {
			if i != ThisGame.HGame {
				return false
			} else {
				robotgo.SwitchToThisWidow(i)
				robotgo.Sleep(1)
				ThisGame.Rec = Rect(robotgo.GetClientRect(int32(ThisGame.HGame)))
				return true
			}
		}
	}
	return false
}

//打开任务
func OpenOrCloseTask() bool {
	if checkActive() == false {
		return false
	}
	robotgo.MicroSleep(wait_micsec)
	robotgo.KeyToggle("alt", "down")
	robotgo.MicroSleep(wait_micsec)
	robotgo.KeyTap("q")
	robotgo.MicroSleep(wait_micsec)
	robotgo.KeyToggle("alt", "up")
	return true
}

//打开背包
func OpneOrCloseBag() bool {
	if checkActive() == false {
		return false
	}
	robotgo.MicroSleep(wait_micsec)
	robotgo.KeyToggle("alt", "down")
	robotgo.MicroSleep(wait_micsec)
	robotgo.KeyTap("e")
	robotgo.MicroSleep(wait_micsec)
	robotgo.KeyToggle("alt", "up")
	return true
}

func OpenOrCloseFriends() {
	checkActive()
	robotgo.MicroSleep(wait_micsec)
	robotgo.KeyToggle("alt", "down")
	robotgo.MicroSleep(wait_micsec)
	robotgo.KeyTap("f")
	robotgo.MicroSleep(wait_micsec)
	robotgo.KeyToggle("alt", "up")
}

func CheckAndOpenFriends() {
	checkActive()
	x, _ := SearchImg(friends_, ThisGame.Rec)
	if x == -1 {
		OpenOrCloseFriends()
	}
}

func CheckAndCloseFriends() {
	checkActive()
	x, _ := SearchImg(friends_, ThisGame.Rec)
	if x != -1 {
		OpenOrCloseFriends()
	}
}

//刷新挂机回合
func FlushAuto() bool {
	if checkActive() == false {
		return false
	}
	robotgo.MicroSleep(wait_micsec)
	robotgo.KeyToggle("control", "down")
	robotgo.MicroSleep(wait_micsec)
	robotgo.KeyTap("a")
	robotgo.MicroSleep(wait_micsec)
	robotgo.KeyToggle("control", "up")
	return true
}

//打开地图
func OpenOrCloseMap() bool {
	if checkActive() == false {
		return false
	}
	robotgo.MicroSleep(wait_micsec)
	robotgo.KeyTap("tab")
	robotgo.MicroSleep(wait_micsec)
	return false
}

//搜索指定图片
//世界坐标
func SearchBitmap(img image.Image, rec Rect) (int, int) {
	checkActive()
	robotgo.MicroSleep(1000)
	MoveToNone()
	screen := robotgo.CaptureScreen(rec.Top_x, rec.Top_y, rec.Width, rec.Height)
	npc := robotgo.Img2CBitmap(img)
	npc_x, npc_y := robotgo.FindBitmap(npc, screen, 0.2)
	if npc_x == -1 {
		return npc_x, npc_y
	} else {
		return rec.Top_x + npc_x, rec.Top_y + npc_y
	}
}

func SearchImg(str string, rec Rect) (int, int) {
	checkActive()
	img, _, _ := robotgo.DecodeImg(str)
	var npc_x, npc_y int
	for i := 0; i < retry_times; i++ {
		screen := robotgo.CaptureScreen(rec.Top_x, rec.Top_y, rec.Width, rec.Height)
		npc := robotgo.Img2CBitmap(img)
		npc_x, npc_y = robotgo.FindBitmap(npc, screen, 0.2)
		if npc_x == -1 {
			MoveToNone()
		} else {
			break
		}
	}
	if npc_x == -1 {
		return npc_x, npc_y
	} else {
		return rec.Top_x + npc_x, rec.Top_y + npc_y
	}
}

//点击搜索到的图片
func ClickOnImg(str string, r ...Rect) bool {
	img_tmp, _, _ := robotgo.DecodeImg(str)
	npc_w := img_tmp.Bounds().Max.X - img_tmp.Bounds().Min.X
	npc_h := img_tmp.Bounds().Max.Y - img_tmp.Bounds().Min.Y
	checkActive()
	var rec Rect
	if len(r) == 0 {
		rec = ThisGame.Rec
	} else {
		rec = r[0]
	}
	npc_x, npc_y := SearchBitmap(img_tmp, rec)
	if npc_x == -1 {
		return false
	}
	robotgo.MoveMouseSmooth(npc_x+npc_w/2, npc_y+npc_h/2)
	robotgo.Click()
	return true
}

//地图上搜索npc
func SearchNpc(str string) bool {
	OpenOrCloseMap()
	robotgo.MicroSleep(wait_micsec)
	if !ClickOnImg(search_npc) {
		return false
	}
	robotgo.MicroSleep(wait_micsec)
	robotgo.TypeStr(str)
	robotgo.MicroSleep(wait_micsec)
	robotgo.KeyTap("enter")
	robotgo.MicroSleep(wait_micsec)
	robotgo.KeyTap("backspace")
	OpenOrCloseMap()
	return true
}

//获取游戏截图
func GetGameScreen(save_path string, r ...Rect) {
	checkActive()
	robotgo.MicroSleep(1000)
	var rec Rect
	if len(r) == 0 {
		rec = ThisGame.Rec
	} else {
		rec = r[0]
	}
	bmp_ref := robotgo.CaptureScreen(rec.Top_x, rec.Top_y, rec.Width, rec.Height)
	defer robotgo.FreeBitmap(bmp_ref)
	robotgo.SaveBitmap(bmp_ref, save_path)
}

//移动到游戏中坐标位置
func MoveToPos(map_name string, pos Point) {
	CheckAndOpenMap()
	robotgo.MicroSleep(wait_micsec)
	if info, ok := map_infos[map_name]; ok {
		x0, y0 := GetPosition()
		rec := &ThisGame.Rec
		center := Point{
			rec.Top_x + rec.Width/2,
			rec.Top_y + rec.Height/2,
		}
		map_ori := Point{
			center.X - info.W/2,
			center.Y + info.H/2,
		}
		abs_pos := Point{
			map_ori.X + int(float32(pos.X)*info.R),
			map_ori.Y - int(float32(pos.Y)*info.R),
		}
		robotgo.MoveMouseSmooth(abs_pos.X, abs_pos.Y)
		robotgo.MicroSleep(wait_micsec)
		robotgo.Click()
		for {
			robotgo.Sleep(3)
			x1, y1 := GetPosition()
			if x1 == x0 && y1 == y0 {
				break
			} else {
				x0, y0 = x1, y1
			}
		}
	}
	CheckAndCloseMap()
}

//得到坐标图
func GetPosition() (int, int) {
	checkActive()
	rec := ThisGame.GetCoRect()
	robotgo.MicroSleep(wait_micsec)
	bmp_ref := robotgo.CaptureScreen(rec.Top_x, rec.Top_y, rec.Width, rec.Height)
	robotgo.SaveBitmap(bmp_ref, co_bmp_path)
	defer robotgo.FreeBitmap(bmp_ref)
	/*str, err := robotgo.GetText(co_bmp_path, "co")
	if err != nil {
		return -1, -1
	}*/
	str := GetText(co_bmp_path)
	fmt.Println(str)
	strings.ToUpper(str)
	parse_int := func(str string, pos int) int {
		beg := pos
		for ; ; beg++ {
			if unicode.IsNumber(rune(str[beg])) {
				break
			}
		}
		end := beg
		for ; ; end++ {
			if end >= len(str) || !unicode.IsNumber(rune(str[end])) {
				break
			}
		}
		x, err := strconv.Atoi(str[beg:end])
		if err != nil {
			return -1
		}
		return x
	}
	pos := strings.Index(str, "X")
	if pos == -1 {
		return -1, -1
	}
	x := parse_int(str, pos)
	pos = strings.Index(str, "Y")
	if pos == -1 {
		return -1, -1
	}
	y := parse_int(str, pos)
	return x, y
}

//得到任务目标
func GetTaskDetail(mode, path string) (int, bool) {
	CheckAndOpenTask()
	robotgo.MicroSleep(wait_micsec)
	var path1, path2 string
	if mode == "goal" {
		path1 = task_goal
		path2 = task_reply
	} else if mode == "reply" {
		path1 = task_reply
		path2 = task_strategy
	} else {
		return 0, false
	}
	img1, _, _ := robotgo.DecodeImg(path1)
	img2, _, _ := robotgo.DecodeImg(path2)
	x1, y1 := SearchBitmap(img1, ThisGame.Rec)
	x2, y2 := SearchBitmap(img2, ThisGame.Rec)
	fmt.Println(x1, y1)
	fmt.Println(x2, y2)
	if x1 == -1 {
		OpenOrCloseTask()
		x1, y1 = SearchBitmap(img1, ThisGame.Rec)
	}
	if x2 == -1 {
		img2, _, _ := robotgo.DecodeImg(task_strategy)
		x2, y2 = SearchBitmap(img2, ThisGame.Rec)
	}
	if (x1 == -1) || (x2 == -1) {
		return 0, false
	}
	rec := Rect{
		x1,
		y1 + img1.Bounds().Max.Y,
		task_width,
		y2 - y1 - img1.Bounds().Max.Y,
	}
	n := rec.Height / 18
	for i := 0; i < n; i++ {
		img := robotgo.CaptureScreen(rec.Top_x, rec.Top_y+18*i, rec.Width, 18)
		img_tmp := image.NewGray(image.Rectangle{
			image.Point{0, 0},
			image.Point{rec.Width, 18},
		})
		for i := 0; i < rec.Height; i++ {
			for j := 0; j < rec.Width; j++ {
				c := robotgo.GetColor(img, j, i)
				r := ((int(c) >> 16) & 0xFF)
				g := ((int(c) >> 8) & 0xFF)
				b := (int(c) & 0xFF)
				delt_r := math.Abs(float64(r - 92))
				delt_g := math.Abs(float64(g - 137))
				delt_b := math.Abs(float64((b - 142)))
				sum := math.Sqrt(float64(delt_r*delt_r + delt_g*delt_g + delt_b*delt_b))
				if sum >= 75 {
					img_tmp.SetGray(j, i, color.Gray{
						0,
					})
				} else {
					img_tmp.SetGray(j, i, color.Gray{255})
				}
			}
		}
		out := robotgo.Img2CBitmap(img_tmp)
		robotgo.SaveBitmap(out, path+strconv.Itoa(i)+".bmp")
		robotgo.FreeBitmap(img)
	}
	CheckAndCloseTask()
	return n, true
}

//按下对应快捷键
func PressKey(key string) {
	checkActive()
	robotgo.MicroSleep(wait_micsec)
	robotgo.KeyTap(key)
}

//检查并开启地图
func CheckAndOpenMap() {
	checkActive()
	robotgo.MicroSleep(wait_micsec)
	x, _ := SearchImg(search_npc, ThisGame.Rec)
	if x == -1 {
		OpenOrCloseMap()
	}
}

//检查并关闭地图
func CheckAndCloseMap() {
	checkActive()
	robotgo.MicroSleep(wait_micsec)
	x, _ := SearchImg(search_npc, ThisGame.Rec)
	if x != -1 {
		OpenOrCloseMap()
	}
}

//检查并开启任务
func CheckAndOpenTask() {
	checkActive()
	robotgo.MicroSleep(wait_micsec)
	x, _ := SearchImg(task_path, ThisGame.Rec)
	if x == -1 {
		OpenOrCloseTask()
	}
}

//检查并关闭任务
func CheckAndCloseTask() {
	checkActive()
	robotgo.MicroSleep(wait_micsec)
	x, _ := SearchImg(task_path, ThisGame.Rec)
	if x != -1 {
		OpenOrCloseTask()
	}
}

//识别字体
func GetText(path string) string {
	rep, _ := postToTencent(path)
	fmt.Println(rep)
	item_tmp := &rep.Data.Items
	var out string
	for _, v := range *item_tmp {
		out += v.Itemstring
	}
	fmt.Println(out)
	return out
}

func GetTextMoreLine(path string) []string {
	rep, _ := postToTencent(path)
	out := make([]string, 0)
	fmt.Println(rep)
	item_tmp := &rep.Data.Items
	for _, v := range *item_tmp {
		out = append(out, v.Itemstring)
	}
	return out
}

//等待人物停止运动
func waitForStop() {
	x0, y0 := GetPosition()
	for {
		robotgo.MicroSleep(100)
		x1, y1 := GetPosition()
		if x0 == x1 && y0 == y1 {
			break
		} else {
			x0 = x1
			y0 = y1
		}
	}
}

//把所有带有关闭按钮的窗口关闭
func ClickCloseButton() {
	for {
		if !ClickOnImg(close_button) {
			break
		}
		robotgo.Sleep(1)
	}
}

//领取宝图任务
func GetBaotuTask() {
	PressKey(fly1)
	robotgo.MicroSleep(wait_micsec)
	PressKey("5")
	robotgo.MicroSleep(1000)
	SearchNpc("江湖密探")
	waitForStop()
	ClickOnImg(get_baotu)
	robotgo.MicroSleep(wait_micsec)
	ClickCloseButton()
}

func checkClickFist(img1, img2 string, r ...Rect) bool {
	checkActive()
	var rec Rect
	if len(r) == 0 {
		rec = ThisGame.Rec
	} else {
		rec = r[0]
	}
	x, _ := SearchImg(img1, rec)
	if x == -1 {
		x, _ = SearchImg(img2, rec)
		if x == -1 {
			return false
		}
	} else {
		ClickOnImg(img1, rec)
	}
	return true
}

//开启自动战斗
func OpenAuto() {
	robotgo.KeyToggle("control", "down")
	robotgo.MicroSleep(wait_micsec)
	robotgo.KeyTap("a")
	robotgo.MicroSleep(wait_micsec)
	robotgo.KeyToggle("control", "up")
}

func CloseAuto() {
	ClickOnImg(close_auto)
}

//等待检查是否进入战斗
func CheckBattle() bool {
	count := 0
	x, _ := SearchImg(check_battle0, ThisGame.Rec)
	for {
		if x == -1 {
			x, _ = SearchImg(check_battle1, ThisGame.Rec)
			if x == -1 {
				if count == 1 {
					return false
				} else {
					count = 1
				}
			} else {
				if count == 2 {
					return true
				} else {
					count = 2
				}
			}
		} else {
			if count == 2 {
				return true
			} else {
				count = 2
			}
		}
	}
}

//巡逻（遇暗怪)
func Patrol() {
	CheckAndOpenMap()
	robotgo.MicroSleep(wait_micsec)
	robotgo.MoveMouseSmooth(ThisGame.Rec.Top_x+ThisGame.Rec.Width/2,
		ThisGame.Rec.Top_y+ThisGame.Rec.Height/2)
	robotgo.Click()
	bx := 1
	by := 1
	l := 150
	for {
		waitForStop()
		if CheckBattle() {
			break
		} else {
			checkActive()
			rec := &ThisGame.Rec
			center := Point{
				rec.Top_x + rec.Width/2,
				rec.Top_y + rec.Height/2,
			}
			robotgo.MoveMouseSmooth(center.X+bx*l,
				center.Y+by*l)
			bx = -1 * bx
			by = -1 * by
		}
	}

	OpenAuto()

	for {
		robotgo.Sleep(2)
		if !CheckBattle() {
			break
		}
	}
	ClickCloseButton()
}

//点击周围
func ClickOnRound(name string) bool {
	CheckAndOpenFriends()
	robotgo.MicroSleep(wait_micsec)
	checkClickFist(friends_round0, friends_round1)
	checkClickFist(round_npc0, round_npc1)
	x, y := SearchImg(round_npc1, ThisGame.Rec)
	defer CheckAndCloseFriends()
	y += 22
	GetGameScreen(round_npc, Rect{
		x,
		y,
		100,
		20,
	})
	if len(name) != 0 {
		str := GetText(round_npc)
		if strings.Index(str, name) == -1 {
			return false
		}
	}
	robotgo.MoveMouseSmooth(ThisGame.Rec.Top_x+x, ThisGame.Rec.Top_y+y+8)
	robotgo.Click()
	waitForStop()
	return true
}

//当识别的坐标不准确的时候查找可能的坐标
func CheckAllPossiblePos(scene, spos, name string) {
	p := Point{
		map_infos[scene].W,
		map_infos[scene].H,
	}
	for i, _ := range spos {
		if i == 0 {
			continue
		}
		x, _ := strconv.Atoi(spos[:i])
		y, _ := strconv.Atoi(spos[i:])
		if float32(x)*map_infos[scene].R > float32(p.X) || map_infos[scene].R*float32(y) > float32(p.Y) {
			continue
		}
		MoveToPos(scene, Point{x, y})
		if ClickOnRound(name) {
			return
		}
	}
}

func OpenTradingPet() {
	PressKey("f7")
	robotgo.MicroSleep(wait_micsec)
	PressKey("2")
	robotgo.Sleep(1)
	SearchNpc("交易中心")
	waitForStop()
	ClickOnImg(trading_pet)
}

func OpenTradingGoods() {
	PressKey("f7")
	robotgo.MicroSleep(wait_micsec)
	PressKey("2")
	robotgo.Sleep(1)
	SearchNpc("交易中心")
	waitForStop()
	ClickOnImg(trading_goods)
}
