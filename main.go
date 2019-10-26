package lesson2package

import (
	"fmt"
	"log"
)

// UserInput implements user data input
func UserInput(minimum, maximum int64)  int64 {
	p := log.Println // the alias for log.Print in order to simplify the code
	input := 0
	p("Please, enter a number from", minimum, "to", maximum, ":")
	for {
		fmt.Scan(&input)
		if input >= minimum && input <= maximum {
			p("The data input is done")
			break
		} else {
			p("The data input has been wrong. Please, repeat:")
		}
	}
	return input
}

// IfEven recognises even and odd numbers
func IfEven() {
	p := log.Println // the alias for log.Print in order to simplify the code
	input = 0
	var intRange int64 = 9223372036854775807
	p("Enter a number:")
	input := UserInput(intRange*-1, intRange)
	if input%2 == 0 {
		p("The entered number is even!")
	} else {
		p("The entered number is odd!")
	}
}

// IfDevisionHasRemainder finds devision remainder
func IfDevisionHasRemainder() {
	p := log.Println // the alias for log.Print in order to simplify the code
	input = 0
	var intRange int64 = 9223372036854775807
	p("Enter a number:")
	input := UserInput(intRange*-1, intRange)
	if input%3 == 0 {
		p(input, "/ 3 has no remainder!")
	} else {
		p(input, "/ 3 has a remainder!")
	}
}

// Fibonacci genereates 100 Fibonacci numbers
func Fibonacci(counter byte, value1, value2 complex128) (returnCounter byte) {
	if counter == 101 {
		return counter
	}
	log.Printf("%3d - %21.0f\n", counter, real(value1))
	return Fibonacci(counter+1, value2, value1+value2)
}

// Sieve generates 100 prime numbers
func Sieve(limit int) {
	numberArray := make(map[int]bool, limit)

	// 0 и 1 не простые числа, это мы сразу знаем
	numberArray[0] = true
	numberArray[1] = true

	// заполняем массив
	for start := 2; start <= limit; start++ {
		numberArray[start] = false
	}

	for key := 2; key*key <= limit; key++ {

		// если key ещё не вычеркнуто
		if numberArray[key] == false {

			// то вычеркнем кратные key
			for toDelete := key * key; toDelete <= limit; toDelete += key {
				numberArray[toDelete] = true
			}
		}
	}

	// выводим список из 100 чисел
	log.Println("The list of the prime numbers is:")
	for newValue := 2; newValue <= limit; newValue++ {
		if numberArray[newValue] == false {
			log.Println(newValue)
		}
	}
}
