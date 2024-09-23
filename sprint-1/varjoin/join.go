//go:build !solution

package varjoin

import (
	"strings"
)

func Join(sep string, args ...string) string {
	ans := strings.Builder{}

	for index := range args {

		ans.WriteString(args[index])
		if index != len(args)-1 {
			ans.WriteString(sep)
		}
	}

	return ans.String()
}
