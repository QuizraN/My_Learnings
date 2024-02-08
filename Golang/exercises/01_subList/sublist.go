package sublist

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
			firstValueOfL2 := l2[0]
			for key1, value1 := range l1 {
				isSuperList := true
				if value1 == firstValueOfL2 && key1+lengthOfL2 <= lengthOfL1 {
					for key2, value2 := range l2 {
						if value2 != l1[key1+key2] {
							isSuperList = false
							break
						}
					}
					if isSuperList {
						return RelationSuperlist
					}
				}
			}
		} else {
			firstValueOfL1 := l1[0]
			for key1, value1 := range l2 {
				isSubList := true
				if value1 == firstValueOfL1 && key1+lengthOfL1 <= lengthOfL2 {
					for key2, value2 := range l1 {
						if value2 != l2[key1+key2] {
							isSubList = false
							break
						}
					}
					if isSubList {
						return RelationSublist
					}
				}
			}
		}
	}
	return RelationUnequal
}
