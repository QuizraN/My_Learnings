package sublist

func checkIfSublist(list, sublist []int) bool {
	for i := 0; i <= len(list)-len(sublist); i++ {
		matches := true
		for j := 0; j < len(sublist); j++ {
			if list[i+j] != sublist[j] {
				matches = false
				break
			}
		}
		if matches {
			return true
		}
	}
	return false
}

func Sublist(l1, l2 []int) Relation {
	lengthOfL1, lengthOfL2 := len(l1), len(l2)
	if lengthOfL1 == 0 && lengthOfL2 == 0 {
		return RelationEqual
	} else if lengthOfL1 == 0 {
		return RelationSublist
	} else if lengthOfL2 == 0 {
		return RelationSuperlist
	} else {
		if lengthOfL1 == lengthOfL2 {
			for key, value := range l1 {
				if value != l2[key] {
					return RelationUnequal
				}
			}
			return RelationEqual
		} else if lengthOfL1 > lengthOfL2 {
			if checkIfSublist(l1, l2) {
				return RelationSuperlist
			}
		} else {
			if checkIfSublist(l2, l1) {
				return RelationSublist
			}
		}
	}
	return RelationUnequal
}
