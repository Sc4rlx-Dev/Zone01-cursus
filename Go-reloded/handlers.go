package main

import ("regexp";"strings")

func handlers(final_res string) string{

	rgx := regexp.MustCompile(`(\s+)([.,!?:;]+)`)
	final_res = rgx.ReplaceAllString(final_res, "${2}")
	rgx = regexp.MustCompile(`([.,!?:;]+)`)
	final_res = rgx.ReplaceAllString(final_res, "${1} ")
	rgx = regexp.MustCompile(`'\s*(.*?)\s*'`)
	final_res = rgx.ReplaceAllString(final_res, "'$1'")
	wordsfinal := strings.Fields(final_res)
	final_res = strings.Join(wordsfinal, " ")

return final_res
}