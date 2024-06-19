package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var Rom = map[string]int{"M": 1000, "D": 500, "C": 100,
	"L": 50, "X": 10, "V": 5, "I": 1}

//							||||||
//----------------Объявление римских чисел----------------------

func plus(a, b int) int {
	return a + b
}
func minus(a, b int) int {
	return a - b
}
func mul(a, b int) int {
	return a * b
}
func di(a, b int) int {
	return a / b
}

//							  ||||||
//-----------------------Работа со знаками------------------------------

func findZnak(line string) (string, error) {
	switch {
	case strings.Contains(line, "+"):
		return "+", nil
	case strings.Contains(line, "-"):
		return "-", nil
	case strings.Contains(line, "*"):
		return "*", nil
	case strings.Contains(line, "/"):
		return "/", nil
	default:
		return "", fmt.Errorf("Введите знак!")
	}
}

//                                        ||||||||
//----------------------------Распознование знака операции------------------

func calc(a, b int, op string) (num int, err error) {
	switch op {
	case "+":
		num = plus(a, b)
	case "-":
		num = minus(a, b)
	case "*":
		num = mul(a, b)
	case "/":
		num = di(a, b)
	default:
		err = fmt.Errorf("%s Ничего нет", op)

	}
	return
}

//	|||||||||||
//
// ------------------------------------ВЫЧИСЛЕНИЯ ИДУТ!!!!!------------------------------------

func isRoma(num string) bool {
	if _, err := Rom[strings.Split(num, "")[0]]; err {
		return true
	}
	return false
}
func rtInt(num string) int {
	val := 0
	nnn := len(num)
	for i := 0; i < nnn; i++ {
		if i != nnn-1 && Rom[string(num[i])] < Rom[string(num[i+1])] {
			val += Rom[string(num[i+1])] - Rom[string(num[i])]
			i++
			continue
		}
		val += Rom[string(num[i])]
	}
	return val
}
func intTR(num int) string {
	var roman string = ""
	var nums = []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	var romans = []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	var ind = len(romans) - 1
	for num > 0 {
		for nums[ind] <= num {
			roman += romans[ind]
			num -= nums[ind]
		}
		ind -= 1
	}
	return roman
}

func getNumsAndType(line string, op string) (x, y int, rom bool, err error) {
	nums := strings.Split(line, op)
	if len(nums) > 2 {
		return x, y, rom, fmt.Errorf("Много оппов")
	}
	firstRomT := isRoma(nums[0])
	secondRomT := isRoma(nums[1])
	if firstRomT != secondRomT {
		return x, y, rom, fmt.Errorf("Др. формат")
	}
	if firstRomT && secondRomT {
		rom = true
		x = rtInt(nums[0])
		y = rtInt(nums[1])
	} else {
		x, err = strconv.Atoi(nums[0])
		if err != nil {
			return
		}
		y, err = strconv.Atoi(nums[1])
		if err != nil {
			return
		}
	}
	if x < 1 || x > 10 || y < 0 || y > 10 {
		return x, y, rom, fmt.Errorf("Введите от 1 до 10")
	}
	return x, y, rom, nil
}
func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Для выхода введите !exit\nВведите пример: ")
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		line = strings.ReplaceAll(line, " ", "")

		if line == "!exit" {
			fmt.Println("exiting..")
			return
		}

		operator, err := findZnak(line)
		if err != nil {
			panic(err)
		}

		a, b, isRom, err := getNumsAndType(line, operator)
		if err != nil {
			panic(err)
		}

		result, err := calc(a, b, operator)
		if err != nil {
			panic(err)
		}

		if isRom {
			if result <= 0 {
				panic("roman numbers can't less 0")
			}

			first := intTR(a)
			second := intTR(b)
			res := intTR(result)

			fmt.Println(first, operator, second, "=", res)
		} else {
			fmt.Println(a, operator, b, "=", result)
		}
	}
}
