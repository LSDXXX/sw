package shenwu

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/go-vgo/robotgo"
)

//宝图任务
func DoBaotu(n int) bool {
	for i := 0; i < n; i++ {
		for {
			ClickCloseButton()
			CheckAndOpenTask()
			if !checkClickFist(task_baotu1, task_baotu2) {
				GetBaotuTask()
				CheckAndOpenTask()
				if !checkClickFist(task_baotu1, task_baotu2) {
					return false
				}
			}
			k, _ := GetTaskDetail("goal", task_detail)
			var str string
			for i := 0; i < k; i++ {
				tmp := GetText(task_detail + strconv.Itoa(i) + ".bmp")
				str += tmp
			}
			info := parseBaotuDetail(str)
			if info.N == 1 {
				s := map_infos[info.Scene]
				s.Fly()
				if info.S == PATROL {
					Patrol()
				} else if info.S == PET {
					CloseAuto()
					if len(info.Spos) == 0 {
						MoveToPos(info.Scene, info.Pos)
						ClickOnRound("")
					} else {
						CheckAllPossiblePos(info.Scene, info.Spos, "群")
					}
					x, y := SearchImg(task_pet, ThisGame.Rec)
					if x != -1 {
						min := ThisGame.Rec.Width*ThisGame.Rec.Width + ThisGame.Rec.Height*ThisGame.Rec.Height
						var min_pos Point
						for i := -2; i <= 2; i++ {
							p := ThisGame.GetEnemyPos(i)
							m := (p.X-x)*(p.X-x) + (p.Y-y)*(p.Y-y)
							if min > m {
								min = m
								min_pos = p
							}
						}
						robotgo.MoveMouseSmooth(ThisGame.Rec.Top_x+min_pos.X,
							ThisGame.Rec.Top_y+min_pos.Y)
						robotgo.Click()
						robotgo.MicroSleep(wait_micsec)
						robotgo.Click()
						for {
							if !CheckBattle() {
								break
							}
						}
					}
				}
			} else if info.N == 2 || info.N == 3 {
				s := map_infos[info.Scene]
				if s != nil {
					s.Fly()
					var name string
					if info.N == 2 {
						name = "宝"
					} else {
						name = "大"
					}
					if len(info.Spos) == 0 {
						MoveToPos(info.Scene, info.Pos)
						ClickOnRound(name)
					} else {
						CheckAllPossiblePos(info.Scene, info.Spos, name)
					}
				}
				if info.N == 3 {
					ClickOnImg(fight_dadangjia)
					OpenAuto()
					robotgo.Sleep(1)
					for {
						if !CheckBattle() {
							break
						}
					}
					break
				}
			}
		}
	}
	return true
}

func parseBaotuDetail(str string) baoTuTaskInfo {
	var out baoTuTaskInfo
	cur := strings.Index(str, "第")
	if cur == -1 {
		if unicode.IsDigit(rune(str[0])) {
			out.N = int(str[0]) - 48
		} else {
			out.N = 1
		}
	} else {
		if unicode.IsDigit(rune(str[3])) {
			out.N = int(str[3]) - 48
		} else {
			out.N = 1
		}
	}
	var sub string
	if out.N == 1 {
		CheckAndOpenTask()
		checkClickFist(task_baotu1, task_baotu2)
		x, _ := SearchImg(baotu_patrol, ThisGame.Rec)
		CheckAndCloseTask()
		if x != -1 {
			pos1 := strings.Index(str, "在")
			pos2 := strings.Index(str, "巡")
			out.S = PATROL
			out.Scene = str[pos1+3 : pos2]
		} else {
			pos1 := strings.Index(str, "从")
			sub = str[pos1+3:]
			pos2 := strings.IndexFunc(sub, func(r rune) bool {
				return r == '[' || r == ']' || r == '(' || r == '{' || r == '}' || r == ')' || unicode.IsDigit(r)
			})
			out.S = PET
			out.Scene = sub[:pos2]
			sub = sub[pos2:]
			sub = strings.TrimFunc(sub, func(r rune) bool {
				return !unicode.IsDigit(r)
			})
		}
	} else if out.N == 2 {
		out.S = FINDING
		pos1 := strings.Index(str, "在")
		sub = str[pos1+3:]
		var pos2 int
		for i, v := range sub {
			fmt.Println(string(v))
			if rune(v) == rune('[') || rune(v) == rune('{') || rune(v) == rune('(') || unicode.IsDigit(v) {
				pos2 = i
				break
			}
		}
		out.Scene = sub[:pos2]
		sub = sub[pos2:]
		sub = strings.TrimFunc(sub, func(r rune) bool {
			return !unicode.IsDigit(r)
		})
	} else if out.N == 3 {
		pos := strings.Index(str, "去")
		sub = str[pos+3:]
		pos = strings.IndexFunc(sub, func(r rune) bool {
			return (r == '{' || r == '(' || r == ')' || r == '}' || unicode.IsDigit(r))
		})
		out.Scene = sub[:pos]
		sub = sub[pos:]
		sub = strings.TrimFunc(sub, func(r rune) bool {
			return !unicode.IsDigit(r)
		})
	}
	if len(sub) > 0 {
		pos1 := strings.IndexFunc(sub, func(r rune) bool {
			return !unicode.IsDigit(r)
		})
		if pos1 != -1 {
			x, _ := strconv.Atoi(sub[:pos1])
			sub = sub[pos1:]
			sub = strings.TrimFunc(sub, func(r rune) bool {
				return !unicode.IsDigit(r)
			})
			y, _ := strconv.Atoi(sub)
			out.Pos = Point{
				x,
				y,
			}
		} else {
			out.Spos = sub
		}
	}
	return out
}
