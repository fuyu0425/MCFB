package main

import (
    "fmt"
    //"html/template"
    //"math"
    "./house"
    "log"
    "net/http"
    "strings"
    //"time"
    //"github.com/FrozenKP/go-mcpi-api"
)

type xyz struct{
  x int
  y int
  z int
}

type player struct{
  name string
  house_site xyz
  first_sign xyz
  second_sign xyz
  third_sign xyz
  name_sign xyz
  first_text string
  second_text string
  third_text string
}

var build =xyz{-246,3,-167} //13 -1 -20
var userlist=make(map[string]player)

func (p player)sign_reset(){
  x,y,z:=p.first_sign.x , p.first_sign.y , p.first_sign.z
  for a,b:=0,0 ; a<5 && b<3 ; {
    fmt.Printf("setblock %d %d %d wall_sign 3 replace {id:\"Sign\",Text1:\"\"}\n",x+a,y-b,z)
    //time.Sleep(100*time.Millisecond)
    if a+1==5 {
      a=0
      b++
    }else{
      a++
    }
  }

  x,y,z=p.second_sign.x , p.second_sign.y , p.second_sign.z
  for a,b:=0,0 ; a<5 && b<3 ; {
    fmt.Printf("setblock %d %d %d wall_sign 4 replace {id:\"Sign\",Text1:\"\"}\n",x,y-b,z+a)
    //time.Sleep(100*time.Millisecond)
    if a+1==5 {
      a=0
      b++
    }else{
      a++
    }
  }

  x,y,z=p.third_sign.x , p.third_sign.y , p.third_sign.z
  for a,b:=0,0 ; a<5 && b<3 ; {
    fmt.Printf("setblock %d %d %d wall_sign 2 replace {id:\"Sign\",Text1:\"\"}\n",x-a,y-b,z)
    //time.Sleep(100*time.Millisecond)
    if a+1==5 {
      a=0
      b++
    }else{
      a++
    }
  }
}

func (p player)sign_show(){
    fmt.Println("say in")
  x,y,z:=p.first_sign.x , p.first_sign.y , p.first_sign.z
  ssf:=strings.Split(p.first_text," ")
  var scf [100]string
  var b int=0
  for a:=0;a<len(ssf);a++{
    if len(scf[b])+1+len(ssf[a])<=13{
        scf[b]=scf[b]+" "+ssf[a]
    }else{
        b+=1
        a-=1
    }
  }
  for xl,yl,a:=0,0,0 ; xl<5 && yl<3 && a<=b; a+=4{
    fmt.Printf("setblock %d %d %d wall_sign 3 replace {id:\"Sign\" ,Text1:\"%s\",Text2:\"%s\",Text3:\"%s\",Text4:\"%s\"}\n",x+xl,y-yl,z,scf[a],scf[a+1],scf[a+2],scf[a+3])
    //time.Sleep(100*time.Millisecond)
    if xl+1==5 {
      xl=0
      yl++
    }else{
      xl++
    }
  }

  x,y,z=p.second_sign.x , p.second_sign.y , p.second_sign.z
  sss:=strings.Split(p.second_text," ")
  var scs [100]string
  b=0
  for a:=0;a<len(sss);a++{
    if len(scs[b])+1+len(sss[a])<=13{
        scs[b]=scs[b]+" "+sss[a]
    }else{
        b+=1
        a-=1
    }
  }
  for xl,yl,a:=0,0,0 ; xl<5 && yl<3 && a<=b; a+=4{
    fmt.Printf("setblock %d %d %d wall_sign 4 replace {id:\"Sign\" ,Text1:\"%s\",Text2:\"%s\",Text3:\"%s\",Text4:\"%s\"}\n",x,y-yl,z+xl,scs[a],scs[a+1],scs[a+2],scs[a+3])
    //time.Sleep(100*time.Millisecond)
    if xl+1==5 {
      xl=0
      yl++
    }else{
      xl++
    }
  }

  x,y,z=p.third_sign.x , p.third_sign.y , p.third_sign.z
  sst:=strings.Split(p.third_text," ")
  var sct [100]string
  b=0
  for a:=0;a<len(sst);a++{
    if len(sct[b])+1+len(sst[a])<=13{
        sct[b]=sct[b]+" "+sst[a]
    }else{
        b+=1
        a-=1
    }
  }
  for xl,yl,a:=0,0,0 ; xl<5 && yl<3 && a<=b; a+=4{
    fmt.Printf("setblock %d %d %d wall_sign 2 replace {id:\"Sign\" ,Text1:\"%s\",Text2:\"%s\",Text3:\"%s\",Text4:\"%s\"}\n",x-xl,y-yl,z,sct[a],sct[a+1],sct[a+2],sct[a+3])
    //time.Sleep(100*time.Millisecond)
    if xl+1==5 {
      xl=0
      yl++
    }else{
      xl++
    }
  }
fmt.Println("say done")
}

func messagein(res http.ResponseWriter, req *http.Request){
  if req.Method == "POST"{
      req.ParseForm()
        //fmt.Println(req)
        //判斷是否存在
        AAA:=strings.Split(req.Form["text"][0],"$$$")
        fmt.Println(req.Form["name"][0],AAA)
        _,exist:=userlist[req.Form["name"][0]]
        fmt.Println(exist)
        //u:=userlist[req.Form["name"][0]]
        var t player
      if exist {//存在->替換訊息
        var times int
        if AAA[0]==userlist[req.Form["name"][0]].first_text {
            return
        }else if AAA[1]==userlist[req.Form["name"][0]].first_text{
            times=1
        }else if AAA[2]==userlist[req.Form["name"][0]].first_text{
            times=2
        }else{times=3}

        t=userlist[req.Form["name"][0]]
        switch times {
            case 1:
                t.third_text=userlist[req.Form["name"][0]].second_text
                t.second_text=userlist[req.Form["name"][0]].first_text
                t.first_text=AAA[0]
            case 2:
                t.third_text=userlist[req.Form["name"][0]].first_text
                t.second_text=AAA[1]
                t.first_text=AAA[0]
            case 3:
                t.third_text=AAA[2]
                t.second_text=AAA[1]
                t.first_text=AAA[0]
        }
        /*t=userlist[req.Form["name"][0]]
        t.third_text=userlist[req.Form["name"][0]].second_text
        t.second_text=userlist[req.Form["name"][0]].first_text
        t.first_text=AAA[0]*/

        userlist[req.Form["name"][0]]=t
      }else{//不存在->初始化
        /*fmt.Println("in")
        var original_xyz=xyz{0,0,0}
        var original_player=player{"",original_xyz,original_xyz,original_xyz,original_xyz,original_xyz,"","",""}
        userlist[req.Form["name"][0]]=original_player*/

        t.name=req.Form["name"][0]

        t.house_site.x=build.x
        t.house_site.y=build.y
        t.house_site.z=build.z

        t.first_sign.x=build.x+2
        t.first_sign.y=build.y+4
        t.first_sign.z=build.z+1

        t.second_sign.x=build.x+7
        t.second_sign.y=build.y+4
        t.second_sign.z=build.z+2

        t.third_sign.x=build.x+6
        t.third_sign.y=build.y+4
        t.third_sign.z=build.z+7

        t.name_sign.x=build.x-1
        t.name_sign.y=build.y+3
        t.name_sign.z=build.z+4

        t.first_text=AAA[0]
        t.second_text=AAA[1]
        t.third_text=AAA[2]
        //建造小屋
        build.z+=10
        hxc:=t.house_site.x+349
        hyc:=t.house_site.y-4
        hzc:=t.house_site.z-695
        house.Create(hxc,hyc,hzc)
        //門牌
        hxc=t.name_sign.x
        hyc=t.name_sign.y
        hzc=t.name_sign.z
        fmt.Printf("setblock %d %d %d wall_sign 4 replace {id:\"Sign\" ,Text2:\"%s\"}\n",hxc,hyc,hzc,req.Form["name"][0])
        userlist[req.Form["name"][0]]=t
        fmt.Println("done")
        }
        fmt.Println(t)
      if userlist[req.Form["name"][0]].first_text != "" {
        userlist[req.Form["name"][0]].sign_reset()
        userlist[req.Form["name"][0]].sign_show()
      }
  }
}

func main() {
    fmt.Println(build)
    http.HandleFunc("/play", messagein)
    err := http.ListenAndServe(":9090", nil) //设置监听的端口
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
