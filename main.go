package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Что считаем?")
		input_raw, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("С такими не работаю. Прощай...")
			os.Exit(1)
		}

		input_split := strings.Split(input_raw, " ")
		if len(input_split) < 3 {
			fmt.Println("Вводи в формате 10 + 10. C Пробелами!")
			os.Exit(1)
		}

		first := strings.TrimSpace(input_split[0])
		oper := strings.TrimSpace(input_split[1])
		secon := strings.TrimSpace(input_split[2])

		first_num := RomanToInt(first)
		secon_num := RomanToInt(secon)
		roma := false

		if first_num != -1 && secon_num != -1 {
			roma = true
		} else {
			first_num = conv(first)
			secon_num = conv(secon)
		}

		var res int
		switch oper {
		case "+":
			res = first_num + secon_num
		case "-":
			res = first_num - secon_num
		case "*":
			res = first_num * secon_num
		case "/":
			res = first_num / secon_num
		default:
			fmt.Println("Неверно, выбирай операции на твой вкус +,-,*,/")
			os.Exit(1)
		}

		if roma {
			if res <= 0 {
				fmt.Println("Рома ушел в отрицание...")
				os.Exit(1)
			} else {
				IntToRomain(res)
			}
		} else {
			fmt.Println(res)
		}
	}
}

// При ошибке возвращается -1
func RomanToInt(roman string) int {
	conversions := []struct {
		index int // Вычитать или прибавлять
		count int // Сколько всего
	}{
		{0, 0},
		{0, 0},
		{0, 0},
	}

	var ret int

	for pos, char := range roman {
		switch string(char) {
		case "I":
			conversions[0].index = pos
			conversions[0].count += 1
		case "V":
			conversions[1].index = pos
			conversions[1].count += 1
			ret += 5
		case "X":
			conversions[2].index = pos
			conversions[2].count += 1
			ret += 10
		default:
			return -1
		}
	}

	// Проверка на чепуху
	if (conversions[1].count+conversions[2].count > 1) || conversions[0].count > 3 {
		fmt.Println("Не обманывай себя и меня...")
		return -1
	}

	if conversions[0].index >= conversions[1].index+conversions[2].index {
		ret += conversions[0].count
	} else {
		ret -= conversions[0].count
	}

	return ret
}

// function from https://www.geeksforgeeks.org/python-program-to-convert-integer-to-roman/
func IntToRomain(number int) {
	num := []int{1, 4, 5, 9, 10, 40, 50, 90,
		100, 400, 500, 900, 1000}
	sym := []string{"I", "IV", "V", "IX", "X", "XL",
		"L", "XC", "C", "CD", "D", "CM", "M"}

	for i := 12; number > 0; i -= 1 {
		div := number / num[i]
		number %= num[i]

		for ; div > 0; div -= 1 {
			fmt.Printf("%s", sym[i])
		}
	}
	fmt.Printf("\n")
}

func conv(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil || num > 10 || num <= 0 {
		fmt.Println("Давай числа от 1 до 10)")
		os.Exit(1)
	}
	return num
}
