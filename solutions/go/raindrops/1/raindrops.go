package raindrops

import (
	"strconv"
)

func Convert(number int) string {
	var resultString string

    if number%3 == 0 || number%5 == 0 || number%7 == 0 {
            if number%3 == 0 {
        resultString = resultString + "Pling"
    }

    if number%5 == 0 {
        resultString = resultString + "Plang"
    }

    if number%7 == 0 {
        resultString = resultString + "Plong"
    }
    } else {
        resultString =  strconv.Itoa(number) 
    }

    return resultString
}
