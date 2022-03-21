package utility

import (
	"strconv"
	"strings"
)

type ActionValue struct {
	Action string
	Val    int
}

type Struct5_1 struct {
	X1, Y1, X2, Y2 int
}

type Struct13_1 struct {
	X, Y int
}

type Struct13_1XY struct {
	XY string
	val int
}

func GetActionValue(input []string) []ActionValue {
	var actVal ActionValue
	var actValArr []ActionValue
	for _, l := range input {
		splittedStr := strings.Split(l, " ")
		actVal.Action = splittedStr[0]
		intVal, err := strconv.Atoi(splittedStr[1])
		if err == nil {
			actVal.Val = intVal
		} else {
			actVal.Val = 0
		}
		actValArr = append(actValArr, actVal)
	}
	return actValArr
}


