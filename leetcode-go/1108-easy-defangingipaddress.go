package main

func defangIPaddr(address string) string {
	defanged := ""

	for _, val := range address {
		if string(val) == "." {
			defanged += "[.]"
		} else {
			defanged += string(val)
		}
	}

	return defanged
}
