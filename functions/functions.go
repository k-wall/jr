package functions

import (
	"math"
	"math/rand"
	"strconv"
	"strings"
	"text/template"
)

func FunctionsMap() template.FuncMap {
	return template.FuncMap(fmap)
}

var fmap = map[string]interface{}{

	// text utilities
	"trim":    strings.TrimSpace,
	"upper":   strings.ToUpper,
	"lower":   strings.ToLower,
	"title":   strings.Title,
	"substr":  func(start, length int, s string) string { return s[start:length] },
	"first":   func(s string) string { return s[:1] },
	"join":    func(sep string, ss []string) string { return strings.Join(ss, sep) },
	"repeat":  func(count int, str string) string { return strings.Repeat(str, count) },
	"trimall": func(c, s string) string { return strings.Trim(s, c) },
	"atoi":    func(a string) int { i, _ := strconv.Atoi(a); return i },
	"split":   func(sep, s string) []string { return strings.Split(s, sep) },
	"markov":  func(prefixLen, numWords int, baseText string) string { return Nonsense(prefixLen, numWords, baseText) },
	"lorem":   func(size int) string { return Lorem(size) },

	//math utilities
	"add":      func(a, b int) int { return a + b },
	"sub":      func(a, b int) int { return a - b },
	"div":      func(a, b int) int { return a / b },
	"mod":      func(a, b int) int { return a % b },
	"mul":      func(a, b int) int { return a * b },
	"max":      math.Max,
	"min":      math.Min,
	"integer":  func(min, max int) int { return rand.Intn(min-max+1) + max },
	"floating": func(min, max float32) float32 { return min + rand.Float32()*(max-min) },
	"random":   func(s []string) string { return s[rand.Intn(len(s))] },
	"randoms":  func(s string) string { a := strings.Split(s, "|"); return a[rand.Intn(len(a))] },

	//networking and time utilities
	"ip":                 func(s string) string { return ip(s) },
	"ip_known_protocols": ipKnownProtocols,
	"ip_known_ports":     ipKnownPorts,
	"unix_time_stamp":    func(days int) int64 { return unixTimeStamp(days) },

	//people
	"name":       name,
	"surname":    surname,
	"middlename": middlename,
	"address":    address,
	"city":       city,
	"state":      state,

	//generic
	"uuid":    uniqueId,
	"bool":    randomBool,
	"yesorno": yesOrNo,
}