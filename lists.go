package lists

import ()

type StringSlice []string

func (list StringSlice) Subtract(excludeList StringSlice) StringSlice {
	newList := make(StringSlice, 0, len(list))
	for _, v := range list {
		if !(excludeList.Includes(v)) {
			newList = append(newList, v)
		}
	}
	return newList
}

func (list StringSlice) Add(otherList StringSlice) StringSlice {
	n := len(list) + len(otherList)
	newList := make(StringSlice, n, n)
	for _, v := range list {
		if !(newList.Includes(v)) {
			newList = append(newList, v)
		}
	}
	for _, v := range otherList {
		if !(newList.Includes(v)) {
			newList = append(newList, v)
		}
	}
	return newList
}

func (list StringSlice) Includes(arg string) bool {
	result := false
	for _, v := range list {
		if v == arg {
			result = true
			break
		}
	}
	return result
}

func (list StringSlice) Equals(otherList StringSlice) bool {
	l1 := list.Subtract(otherList)
	l2 := otherList.Subtract(list)
	l3 := l1.Add(l2)
	count := 0
	for _, v := range l3 {
		if v != "" {
			count += 1
		}
	}
	return (count == 0)
}
