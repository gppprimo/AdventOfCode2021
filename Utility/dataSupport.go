package utility

import (
	"strconv"
	"strings"
)

type ActionValue struct {
	Action string
	Val    int
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
