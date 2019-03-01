/**
 */
package log

import (
	"fmt"
	"strconv"
)

func Debug(format string, v ...interface{}) {
	datetime := getTimeStr()
	file, line, function := getCallerInfo(0);
	text := fmt.Sprintf(format, v...)
	fmt.Println(datetime + ", " + file + ", Line " + strconv.Itoa(line) + ", " + function + "() - " + text)
	return
}
