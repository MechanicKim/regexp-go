package backreference

import (
	"regexp"
	"strconv"
)

func getRepeatCharGroup(input string, chars string, count int) []string {
	pattern := "(" + regexp.QuoteMeta(chars) + "){" + strconv.Itoa(count) + ",}"
	re := regexp.MustCompile(pattern)
	return re.FindAllString(input, -1)
}

func getSameLengthCharGroup(input string, pattern string) []string {
	re := regexp.MustCompile(pattern)
	return re.FindAllString(input, -1)
}

// FindNTimesRepeated : find N times repeated string
func FindNTimesRepeated(input string, pattern string, count int) map[string][]string {
	result := make(map[string][]string)
	mGroup := getSameLengthCharGroup(input, pattern)
	buf := make(map[string]bool)
	for _, v := range mGroup {
		if _, ok := buf[v]; ok {
			continue
		} else {
			buf[v] = true
		}

		mGroup2 := getRepeatCharGroup(input, v, count)
		if len(mGroup2) > 0 {
			result[v] = mGroup2
		}
	}

	return result
}

// FindRepeated : find repeated string
func FindRepeated(input string, pattern string) map[string][]string {
	return FindNTimesRepeated(input, pattern, 2)
}
