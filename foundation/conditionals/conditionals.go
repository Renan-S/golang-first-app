package main

import "fmt"

func main() {
	//IF ELSE CONDITIONAL
	number := 10

	if number >= 15 {
		fmt.Println("Is bigger or equal than 15")
	} else {
		fmt.Println("Is lower than 15")
	}

	// anotherNum exists only inside the if scope
	if anotherNum := number; anotherNum == 10 {
		body := fmt.Sprintf("anotherNum is %d", anotherNum)
		fmt.Println(body)
	}

	//SWITCH CASE
	fmt.Println("-------------------")
	fmt.Println(weekDay(0))
	fmt.Println(weekDay(6))
	fmt.Println(weekDay(7))
	fmt.Println(caseAsIfWeekDay(2))
}

func weekDay(number int) string {
	var returnValue string

	switch number {
	case 0:
		returnValue = "Sunday"
	case 1:
		returnValue = "Monday"
	case 2:
		returnValue = "Tuesday"
	case 3:
		returnValue = "Wednesday"
	case 4:
		returnValue = "Thursday"
	case 5:
		returnValue = "Friday"
	case 6:
		returnValue = "Saturday"
		fallthrough //Will run the next validation - Useful for code repetition
	default:
		returnValue = "Invalid number"
	}

	return returnValue
}

func caseAsIfWeekDay(number int) string { //Useful if we wanna use more than one variable or param
	switch {
	case number == 0:
		return "Sunday"
	case number == 1:
		return "Monday"
	case number == 2:
		return "Tuesday"
	case number == 3:
		return "Wednesday"
	case number == 4:
		return "Thursday"
	case number == 5:
		return "Friday"
	case number == 6:
		return "Saturday"
	default:
		return "Invalid number"
	}
}
