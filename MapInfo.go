package shenwu

import "github.com/go-vgo/robotgo"

var map_infos map[string]*mapInfo

func init() {
	/*map_infos = map[string]*mapInfo{
		"长安城": &mapInfo{
			660,
			390,
			1.2,
			make(map[string]connectInfo),
		},
		"女儿国": &mapInfo{
			520,
			390,
			3.25,
			make(map[string]connectInfo),
		},
		"青河镇": &mapInfo{
			520,
			390,
			3.25,
			make(map[string]connectInfo),
		},
		"傲来国": &mapInfo{
			560,
			420,
			3.5,
			make(map[string]connectInfo),
		},
		"临县镇": &mapInfo{
			600,
			360,
			2.5,
			make(map[string]connectInfo),
		},
		"乌斯藏": &mapInfo{
			600,
			300,
			2.5,
			make(map[string]connectInfo),
		},
		"大唐境外": &mapInfo{
			600,
			336,
			1.95,
			make(map[string]connectInfo),
		},
		"大唐国境": &mapInfo{
			500,
			450,
			2.5,
			make(map[string]connectInfo),
		},
		"长安城外": &mapInfo{
			664,
			372,
			4.176,
			make(map[string]connectInfo),
		},
		"青河镇外": &mapInfo{
			560,
			422,
			4.667,
			make(map[string]connectInfo),
		},
	}*/
	map_infos = make(map[string]*mapInfo)
	//长安城
	changan := new(mapInfo)
	changan.W = 660
	changan.H = 390
	changan.R = 1.2
	changan.Cons = make(map[string]connectInfo)
	changan.Fly = func() {
		PressKey(fly0)
		PressKey("1")
	}
	changan.Cons["大唐国境"] = connectInfo{
		Method: func() {
			MoveToPos("长安城", Point{30, 307})
		},
		Dst: map_infos["大唐国境"],
	}
	changan.Cons["长安城外"] = connectInfo{
		Method: func() {
			MoveToPos("长安城", Point{543, 4})
		},
		Dst: map_infos["长安城外"],
	}
	map_infos["长安城"] = changan

	//长安城外
	chengwai := new(mapInfo)
	chengwai.W = 664
	chengwai.H = 372
	chengwai.R = 4.176
	chengwai.Cons = make(map[string]connectInfo)
	chengwai.Fly = func() {
		PressKey(fly1)
		robotgo.MicroSleep(wait_micsec)
		PressKey("8")
		robotgo.MicroSleep(wait_micsec)
		MoveToPos("长安城", Point{543, 4})
	}
	chengwai.Cons["长安城"] = connectInfo{
		Method: func() {
			MoveToPos("长安城外", Point{3, 83})
		},
		Dst: map_infos["长安城"],
	}
	chengwai.Cons["青河镇"] = connectInfo{
		Method: func() {
			MoveToPos("长安城外", Point{156, 87})
		},
		Dst: map_infos["青河镇"],
	}
	map_infos["长安城外"] = chengwai

	//青河镇
	qinghe := new(mapInfo)
	qinghe.W = 520
	qinghe.H = 390
	qinghe.R = 3.25
	qinghe.Cons = make(map[string]connectInfo)
	qinghe.Fly = func() {
		PressKey(fly0)
		robotgo.MicroSleep(wait_micsec)
		PressKey("3")
	}
	qinghe.Cons["青河镇外"] = connectInfo{
		Method: func() {
			MoveToPos("青河镇", Point{154, 117})
		},
		Dst: map_infos["青河镇外"],
	}
	map_infos["青河镇"] = qinghe

	//青河镇外
	zhenwai := new(mapInfo)
	zhenwai.W = 560
	zhenwai.H = 422
	zhenwai.R = 4.667
	zhenwai.Cons = make(map[string]connectInfo)
	zhenwai.Fly = func() {
		PressKey(fly1)
		robotgo.MicroSleep(wait_micsec)
		PressKey("9")
		robotgo.MicroSleep(wait_micsec)
		MoveToPos("青河镇", Point{154, 117})
	}
	zhenwai.Cons["青河镇"] = connectInfo{
		Method: func() {
			MoveToPos("青河镇外", Point{3, 85})
		},
		Dst: map_infos["青河镇"],
	}
	map_infos["青河镇外"] = zhenwai

	//傲来国
	aolai := new(mapInfo)
	aolai.W = 560
	aolai.H = 420
	aolai.R = 3.5
	aolai.Cons = make(map[string]connectInfo)
	aolai.Fly = func() {
		PressKey(fly0)
		robotgo.MicroSleep(wait_micsec)
		PressKey("2")
	}
	map_infos["傲来国"] = aolai

	//临仙镇
	linxian := new(mapInfo)
	linxian.W = 600
	linxian.H = 360
	linxian.R = 3
	linxian.Cons = make(map[string]connectInfo)
	linxian.Fly = func() {
		PressKey(fly0)
		robotgo.MicroSleep(wait_micsec)
		PressKey("4")
	}

	map_infos["临仙镇"] = linxian

	//女儿国
	nver := new(mapInfo)
	nver.W = 520
	nver.H = 390
	nver.R = 3.25
	nver.Fly = func() {
		PressKey(fly0)
		robotgo.MicroSleep(wait_micsec)
		PressKey("5")
	}

	map_infos["女儿国"] = nver

	//乌斯藏
	wusi := new(mapInfo)
	wusi.W = 600
	wusi.H = 300
	wusi.R = 2.5
	wusi.Cons = make(map[string]connectInfo)
	wusi.Fly = func() {
		PressKey(fly2)
		robotgo.MicroSleep(wait_micsec)
		PressKey("4")
		robotgo.MicroSleep(wait_micsec)
		MoveToPos("临仙镇", Point{184, 6})
	}
	map_infos["乌斯藏"] = wusi

	//大唐境外
	jingwai := new(mapInfo)
	jingwai.W = 600
	jingwai.H = 336
	jingwai.R = 1.95
	jingwai.Cons = make(map[string]connectInfo)
	jingwai.Fly = func() {
		PressKey(fly2)
		robotgo.MicroSleep(wait_micsec)
		PressKey("9")
		MoveToPos("女儿国", Point{159, 8})
	}
	map_infos["大唐境外"] = jingwai

	//大唐国境
	guojing := new(mapInfo)
	guojing.W = 500
	guojing.H = 450
	guojing.R = 2.5
	guojing.Cons = make(map[string]connectInfo)
	guojing.Fly = func() {
		PressKey(fly1)
		robotgo.MicroSleep(wait_micsec)
		PressKey("6")
		MoveToPos("长安城", Point{528, 319})
	}
	map_infos["大唐国境"] = guojing

	//金銮殿
	jinluan := new(mapInfo)
	jinluan.Fly = func() {
		PressKey(fly1)
		robotgo.MicroSleep(wait_micsec)
		PressKey("1")
		MoveToPos("长安城", Point{32, 265})
	}
	map_infos["金銮殿"] = jinluan

	//万兽岭
	wanshou := new(mapInfo)
	wanshou.Fly = func() {
		m := map_infos["乌斯藏"]
		m.Fly()
		MoveToPos("乌斯藏", Point{3, 5})
	}
	map_infos["万兽岭"] = wanshou

	//魔王山
	mowang := new(mapInfo)
	mowang.Fly = func() {
		m := map_infos["乌斯藏"]
		m.Fly()
		MoveToPos("乌斯藏", Point{14, 108})
	}
	map_infos["魔王山"] = mowang
}

func GetMap() map[string]*mapInfo {
	return map_infos
}
