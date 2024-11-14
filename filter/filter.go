package filter

import "strings"

var FuncMap map[string]func(*string) bool

// strategies
func init() {
	FuncMap = make(map[string]func(*string) bool)
	FuncMap["email"] = FilterByEmail
	FuncMap["tel"] = FilterByTel
	FuncMap["cv"] = FilterByCV
	FuncMap["name"] = FilterByName
	FuncMap["resume"] = FilterByResume
	FuncMap["telphone"] = FilterByTelephone
}

const TRUE = true
const FALSE = false

func FilterByEmail(body *string) bool {
	if strings.Contains(*body, "email") {
		return TRUE
	}
	return FALSE
}

func FilterByTel(body *string) bool {
	if strings.Contains(*body, "tel") {
		return TRUE
	}
	return FALSE
}

func FilterByName(body *string) bool {
	if strings.Contains(*body, "name") {
		return TRUE
	}
	return FALSE
}

func FilterByCV(body *string) bool {
	if strings.Contains(*body, "cv") {
		return TRUE
	}
	return FALSE
}

func FilterByResume(body *string) bool {
	if strings.Contains(*body, "resume") {
		return TRUE
	}
	return FALSE
}

func FilterByTelephone(body *string) bool {
	if strings.Contains(*body, "+(86)") {
		return TRUE
	}
	return FALSE
}
