package main

import (
	"fmt"
	//"html/template"
	//"math"
	"./player"
	"log"
	"net/http"
	"strings"
	//"time"
	//"github.com/FrozenKP/go-mcpi-api"
)

var build = player.XYZ{-246, 3, -167} //13 -1 -20
var userlist = make(map[string]player.Player)

func messagein(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		req.ParseForm()
		//fmt.Println(req)
		//判斷是否存在
		AAA := strings.Split(req.Form["text"][0], "$$$")
		fmt.Println(req.Form["name"][0], AAA)
		_, exist := userlist[req.Form["name"][0]]
		fmt.Println(exist)
		//u:=userlist[req.Form["name"][0]]
		var t player.Player
		if exist { //存在->替換訊息
			var times int
			if AAA[0] == userlist[req.Form["name"][0]].First_text {
				return
			} else if AAA[1] == userlist[req.Form["name"][0]].First_text {
				times = 1
			} else if AAA[2] == userlist[req.Form["name"][0]].First_text {
				times = 2
			} else {
				times = 3
			}

			t = userlist[req.Form["name"][0]]
			switch times {
			case 1:
				t.Third_text = userlist[req.Form["name"][0]].Second_text
				t.Second_text = userlist[req.Form["name"][0]].First_text
				t.First_text = AAA[0]
			case 2:
				t.Third_text = userlist[req.Form["name"][0]].First_text
				t.Second_text = AAA[1]
				t.First_text = AAA[0]
			case 3:
				t.Third_text = AAA[2]
				t.Second_text = AAA[1]
				t.First_text = AAA[0]
			}
			userlist[req.Form["name"][0]] = t

		} else { //不存在->初始化

			t.Name = req.Form["name"][0]

			t.House_site.X = build.X
			t.House_site.Y = build.Y
			t.House_site.Z = build.Z

			t.First_sign.X = build.X + 2
			t.First_sign.Y = build.Y + 4
			t.First_sign.Z = build.Z + 1

			t.Second_sign.X = build.X + 7
			t.Second_sign.Y = build.Y + 4
			t.Second_sign.Z = build.Z + 2

			t.Third_sign.X = build.X + 6
			t.Third_sign.Y = build.Y + 4
			t.Third_sign.Z = build.Z + 7

			t.Name_sign.X = build.X - 1
			t.Name_sign.Y = build.Y + 3
			t.Name_sign.Z = build.Z + 4

			t.First_text = AAA[0]
			t.Second_text = AAA[1]
			t.Third_text = AAA[2]
			//建造小屋
			build.Z += 10
			hxc := t.House_site.X + 349
			hyc := t.House_site.Y - 4
			hzc := t.House_site.Z - 695
			player.CreateHouse(hxc, hyc, hzc)
			//門牌
			hxc = t.Name_sign.X
			hyc = t.Name_sign.Y
			hzc = t.Name_sign.Z
			fmt.Printf("setblock %d %d %d wall_sign 4 replace {id:\"Sign\" ,Text2:\"%s\"}\n", hxc, hyc, hzc, req.Form["name"][0])
			userlist[req.Form["name"][0]] = t
			fmt.Println("done")
		}
		//fmt.Println(t)
		if userlist[req.Form["name"][0]].First_text != "" {
			userlist[req.Form["name"][0]].Sign_reset()
			userlist[req.Form["name"][0]].Sign_show()
		}
	}
}

func main() {
	http.HandleFunc("/play", messagein)
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
