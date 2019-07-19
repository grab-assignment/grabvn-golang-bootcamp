package main

import ("fmt"
        "bufio"
	"os"
	"strings"
	"strconv"
)


func eval(str string){
	s := strings.Split(str, " ")
	if len(s) == 3 {
		a, err := strconv.ParseFloat(s[0], 32)
		err_flag := false
		if err != nil {
			fmt.Println("Math error!")
			err_flag = true
		}
		b, err := strconv.ParseFloat(s[2], 32)
		if err != nil {
			fmt.Println("Math error!")
			err_flag = true
		}
		if err_flag == false {
			switch s[1]{
			case "+" :
				fmt.Println(str, " = ", a + b)
			case "-" :
				fmt.Println(str, " = ", a - b)
			case "*" :
				fmt.Println(str, " = ", a * b)
			case "/" :
				if b == 0{
					fmt.Println("Math error!")
				} else {
					fmt.Println(str, " = ", a / b)
				}
			default:
				fmt.Print("Operator error! (+, -, *, /)")
			}
		}
	} else {
		fmt.Println("Math error!")
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(">")
	for scanner.Scan(){
		text := scanner.Text()
		if text != "" {
			eval(text)
			}
		fmt.Print(">")
	}
}

