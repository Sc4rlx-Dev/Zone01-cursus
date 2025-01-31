package main

import ("strconv")

func binToInt(bin string) string { val, err := strconv.ParseInt(bin, 2, 64)
	if err != nil { return bin }
	return strconv.FormatInt(val, 10)
}
