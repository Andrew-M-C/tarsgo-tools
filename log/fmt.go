/**
 */
package log

import (
	"runtime"
	"strings"
	"path/filepath"
	"time"
	"fmt"
	"strconv"
)

func getCallerInfo(invoke_level int) (fileName string, line int, funcName string) {
	funcName = "FILE"
	line = -1
	fileName = "FUNC"

	if invoke_level <= 0 {
		invoke_level = 2
	} else {
		invoke_level += 1
	}

	pc, file_name, line, ok := runtime.Caller(invoke_level)
	if ok {
		fileName = filepath.Base(file_name)
		func_name := runtime.FuncForPC(pc).Name()
		func_name = filepath.Ext(func_name)
		funcName = strings.TrimPrefix(func_name, ".")
	}
	// fmt.Println(reflect.TypeOf(pc), reflect.ValueOf(pc))
	return
}

func getTimeStr() string {
	return time.Now().Local().Format("2006-01-02 15:04:05.000000")
}

func Gen(format string, v ...interface{}) string {
	datetime := getTimeStr()
	file, line, function := getCallerInfo(0);
	text := fmt.Sprintf(format, v...)
	return datetime + ", " + file + ", Line " + strconv.Itoa(line) + ", " + function + "() - " + text
}
