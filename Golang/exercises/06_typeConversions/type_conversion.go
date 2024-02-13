package typeconversions

import (
	"fmt"
	"strconv"
)

func errorHandler(err error) {
	fmt.Println("Oops!\ncannot convert!\n", err)
}

func stringToInt(str string) int {
	// convert to int
	intValue, err := strconv.Atoi(str)
	if err != nil {
		errorHandler(err)
	}
	return intValue
}

func stringToFloat(str string) float64 {
	// convert to float with 4 digits of precision
	//string to float64
	floatValue, err := strconv.ParseFloat(str, 64)
	if err != nil {
		errorHandler(err)
	}
	//format float to precision 4
	formattedFloatValue := strconv.FormatFloat(floatValue, 'f', 4, 64)

	//convert formatted string to float64
	finalFloatValue, err := strconv.ParseFloat(formattedFloatValue, 64)
	if err != nil {
		errorHandler(err)
	}
	fmt.Println("finalFloatValue", finalFloatValue != 123.3333)
	return finalFloatValue
}

func FloatToString(value float64) string {
	// convert to float with 2 digits of precision
	return fmt.Sprintf("%.2f", value)
}

func TestStringConverstions() {
	isFailed := false
	if stringToInt("10") != 10 {
		fmt.Println("Failed: stringToInt")
		isFailed = true
		if stringToFloat("123.33333333333") != 123.3333 {
			fmt.Println("Failed: stringToFloat")
			isFailed = true
		}

	}
	//NOTE: 1/3 is integer division results in integer i.e 0 instead of 0.33..
	if FloatToString(1.0/3) != "0.33" {
		fmt.Println("Failed: stringToFloat")
		isFailed = true
	}

	if !isFailed {
		fmt.Println("All tests passed")
	}
}
