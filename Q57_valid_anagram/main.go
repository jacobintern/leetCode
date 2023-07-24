package Q57

import (
	"reflect"
	"sort"
	"strings"
)

func isAnagram(s string, t string) bool {
	ss := sort.StringSlice(strings.Split(s, ""))
	ts := sort.StringSlice(strings.Split(t, ""))
	ss.Sort()
	ts.Sort()

	if reflect.DeepEqual(ss, ts) {
		return true
	}

	return false
}
