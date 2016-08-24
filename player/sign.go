package player

import "fmt"
import "strings"

type XYZ struct {
	X int
	Y int
	Z int
}

type Player struct {
	Name        string
	House_site  XYZ
	First_sign  XYZ
	Second_sign XYZ
	Third_sign  XYZ
	Name_sign   XYZ
	First_text  string
	Second_text string
	Third_text  string
}

type strlang struct {
	str string
	eng bool //1:eng 0:other
}

func (p Player) Sign_reset() {
	x, y, z := p.First_sign.X, p.First_sign.Y, p.First_sign.Z
	for a, b := 0, 0; a < 5 && b < 3; {
		fmt.Printf("setblock %d %d %d wall_sign 3 replace {id:\"Sign\",Text1:\"\"}\n", x+a, y-b, z)
		//time.Sleep(100*time.Millisecond)
		if a+1 == 5 {
			a = 0
			b++
		} else {
			a++
		}
	}

	x, y, z = p.Second_sign.X, p.Second_sign.Y, p.Second_sign.Z
	for a, b := 0, 0; a < 5 && b < 3; {
		fmt.Printf("setblock %d %d %d wall_sign 4 replace {id:\"Sign\",Text1:\"\"}\n", x, y-b, z+a)
		//time.Sleep(100*time.Millisecond)
		if a+1 == 5 {
			a = 0
			b++
		} else {
			a++
		}
	}

	x, y, z = p.Third_sign.X, p.Third_sign.Y, p.Third_sign.Z
	for a, b := 0, 0; a < 5 && b < 3; {
		fmt.Printf("setblock %d %d %d wall_sign 2 replace {id:\"Sign\",Text1:\"\"}\n", x-a, y-b, z)
		//time.Sleep(100*time.Millisecond)
		if a+1 == 5 {
			a = 0
			b++
		} else {
			a++
		}
	}
}

func (p Player)Sign_show(){
  p.sign_show_detail(p.First_sign.X, p.First_sign.Y, p.First_sign.Z,1,-1,0,3,p.First_text)
  p.sign_show_detail(p.Second_sign.X, p.Second_sign.Y, p.Second_sign.Z,0,-1,1,4,p.Second_text)
  p.sign_show_detail(p.Third_sign.X, p.Third_sign.Y, p.Third_sign.Z,-1,-1,0,2,p.Third_text)
}

func (p Player) sign_show_detail(x,y,z,xc,yc,zc,dir int,s string) {
	ss := strings.Split(s, " ")
	var final [500]strlang
	var final_length int = 0
	for a := 0; a < len(ss); a++ {
		//判斷中英及切割
		tmpr := []rune(ss[a])
		var lang bool
		for start, now := 0, 0; now < len(tmpr); now++ {
			if now == 0 {
				if tmpr[now] < 127 {
					lang = true
				} else {
					lang = false
				}
				if len(tmpr) == 1 {
					final[final_length].str = ss[a]
					final[final_length].eng = lang
					final_length++
					break
				}
			} else {
				if lang {
					if tmpr[now] < 127 {
						if len(tmpr) == now+1 {
							final[final_length].str = string(tmpr[start:])
							final[final_length].eng = lang
							for true {
								if len(final[final_length].str) > 23 {
									final[final_length+1].str = final[final_length].str[23:]
									final[final_length].str = final[final_length].str[0:23]
									final[final_length+1].eng = lang
									final_length++
								} else {
									break
								}
							}
							final_length++
							break
						} else {
							continue
						}
					} else {
						final[final_length].str = string(tmpr[start:now])
						final[final_length].eng = lang
						for true {
							if len(final[final_length].str) > 23 {
								final[final_length+1].str = final[final_length].str[23:]
								final[final_length].str = final[final_length].str[0:23]
								final[final_length+1].eng = lang
								final_length++
							} else {
								break
							}
						}
						final_length++
						start = now
						lang = false
					}
				} else {
					if tmpr[now] >= 127 {
						if len(tmpr) == now+1 {
							final[final_length].str = string(tmpr[start:])
							final[final_length].eng = lang
							for true {
								if len(final[final_length].str) > 21 {
									final[final_length+1].str = final[final_length].str[21:]
									final[final_length].str = final[final_length].str[0:21]
									final[final_length+1].eng = lang
									final_length++
								} else {
									break
								}
							}
							final_length++
							break
						} else {
							continue
						}
					} else {
						final[final_length].str = string(tmpr[start:now])
						final[final_length].eng = lang
						for true {
							if len(final[final_length].str) > 21 {
								final[final_length+1].str = final[final_length].str[21:]
								final[final_length].str = final[final_length].str[0:21]
								final[final_length+1].eng = lang
								final_length++
							} else {
								break
							}
						}
						final_length++
						start = now
						lang = true
					}
				}
			}
		}
	}
	/*fmt.Printf("say %d\n", final_length)
	for a := 0; a < final_length; a++ {
		fmt.Printf("say %s\n", final[a].str)
	}*/

	var sc [100]string
	var b int = 0
	var pre string = ""
	for a := 0; a < final_length; a++ {
		if len(sc[b])+len(pre)+len(final[a].str) <= 23 {
			sc[b] = sc[b] + pre + final[a].str
			if final[a].eng {
				pre = " "
			} else {
				pre = ""
			}
		} else {
			pre = ""
			b += 1
			a -= 1
		}
	}
	for xl, yl, a := 0, 0, 0; xl < 5 && yl < 3 && a <= b; a += 4 {
		fmt.Printf("setblock %d %d %d wall_sign %d replace {id:\"Sign\" ,Text1:\"%s\",Text2:\"%s\",Text3:\"%s\",Text4:\"%s\"}\n", x+xc*xl, y+yc*yl, z+zc*xl,dir, sc[a], sc[a+1], sc[a+2], sc[a+3])
		if xl+1 == 5 {
			xl = 0
			yl++
		} else {
			xl++
		}
	}
}
