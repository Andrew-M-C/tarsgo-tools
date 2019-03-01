/**
 */
package log

import (
	"fmt"
)

func Debug(format string, v ...interface{}) string {
	datetime := getTimeStr()
	file, line, function := getCallerInfo(0);
	text := fmt.Sprintf(format, v...)
	fmt.Println(datetime + ", " + file + ", Line " + strconv.Itoa(line) + ", " + function + "() - " + text)
}