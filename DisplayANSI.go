package main

import (
	"fmt"
	_ "strconv"
)

/*
https://en.wikipedia.org/wiki/ANSI_escape_code 
ESC = 27 = \x1b = \033
0: reset
1~10: style
30~37, 90~97: foreground
40~47, 100~107: background
"\x1b[%d, %d, %dm"
ERROR := "\x1b[33m %s\x1b[0m"
*/


func main() {
	styles := []int{1, 2, 4, 7, 9}
	//foreground+background
	for _, style := range styles {
		for light:=0;light<2;light++{
			for j:=0;j < 7;j++{
				for light2:=0;light2<2;light2++{
					for k:=0;k < 7;k++{
						format := fmt.Sprintf(
							"\x1b[%d;%d;%dm", style, 30+j+light*60, 40+k+light2*60)
							format += "%s"
							format += "\x1b[0m"//%d;%d;%d\n"
							fmt.Printf(format, "show!!!")
						}
					}
					fmt.Println("\n")

				}

			}
		}
		//only foreground
		for _, style := range styles {
			for light:=0;light<2;light++{
				for j:=0;j < 7;j++{
					format := fmt.Sprintf(
						"\x1b[%d;%dm", style, 30+j+light*60)
						format += "%s"
						format += "\x1b[0m"//%d;%d;%d\n"
						fmt.Printf(format, "show!!!")
					}

				}
				fmt.Println("\n")
			}
		}
