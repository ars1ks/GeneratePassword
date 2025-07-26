package main

import (
	"fmt"
	"os"
	"encoding/json"
	"github.com/fatih/color"
)

type message struct{
	Suite string `json:"suite"`
	Password strint `json:"password"`
}

func key(){

}

func encoder(){

}

func decoder(){


}

func main() {
	color.Green(`
                                          ,---,       ,-. 
                                       ,'--.' |   ,--/ /| 
                __  ,-.               /    /  : ,--. :/ | 
              ,' ,'/ /|   .--.--.    :    |.' ' :  : ' /     .--.--. 
   ,--.--.    '  | |' |  /  /    '   '----':  | |  '  /     /  /    ' 
  /       \   |  |   ,' |  :  /./       '   ' ; '  |  :    |  :  /./  
 .--.  .-. |  '  :  /   |  :  ;_        |   | | |  |   \   |  :  ;_  
  \__\/: . .  |  | '     \  \    '.     '   : ; '  : |. \   \  \    '. 
  ," .--.; |  ;  : |      '----.   \    |   | ' |  | ' \ \   '----.   \ 
 /  /  ,.  |  |  , ;     /  /'--'  /    '   : | '  : |--'   /  /'--'  / 
;  :   .'   \  ---'     '--'.     /     ;   |.' ;  |,'     '--'.     /  
|  ,     .-./             '--'---'      '---'   '--'         '--'---'   
 '--'---'`)

color.Green(`
	1.Посмотреть пароли;
	2.Придумать новый пароль;
	3.Сгенирировать Ключ;
	4.Выход`)

	var cout int
	fmt.Scanln(&cout)
	
	switch cout{ 
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	case 4:
		return
	default:
			fmt.Println("Нет такого выбора")
	}
}
