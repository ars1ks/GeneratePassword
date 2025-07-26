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


 ____             _   __     __    _ _      _   
/ ___| _   _  ___| | _\ \   / /_ _| | | ___| |_ 
\___ \| | | |/ __| |/ /\ \ / / _` | | |/ _ \ __|
 ___) | |_| | (__|   <  \ V / (_| | | |  __/ |_ 
|____/ \__,_|\___|_|\_\  \_/ \__,_|_|_|\___|\__|


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
