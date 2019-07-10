package sesslog
import(
	"github.com/TarsCloud/TarsGo/tars"
	"github.com/TarsCloud/TarsGo/tars/util/rogger"
	"golang.org/x/sync/syncmap"
	"sync"
	"strings"
	"runtime"
	"path/filepath"
	"container/list"
	"strconv"
	"fmt"
	"time"
)

type loglevel int

const (
	_DEBUG loglevel = iota
	_INFO
	_WARN
	_ERROR
)


var (
	_loggerLocks	syncmap.Map		// stores *sync.Mutex
	_shouldDebug	bool = false
	_shouldInfo		bool = true
	_shouldWarn		bool = true
	_shouldError	bool = true
)


func init() {
	cfg := tars.GetServerConfig()
	switch strings.ToUpper(cfg.LogLevel) {
	case "DEBUG":
		_shouldDebug = true
	case "INFO":
		// nothing to change
	case "WARN", "WARNING":
		_shouldInfo = false
	case "ERROR":
		_shouldInfo = false
		_shouldWarn = false
	case "NONE", "NOTHING":
		_shouldInfo = false
		_shouldWarn = false
		_shouldError = false
	}

}


type logitem struct {
	level	loglevel
	log		string
}


type SessLogger struct {
	name		string
	logs		*list.List
	writeLock	*sync.Mutex		// lock when calling write function
	autoWrite	bool
	logger		*rogger.Logger
}


func New(name string) *SessLogger {
	ret := new(SessLogger)
	ret.name = name
	ret.logs = list.New()
	ret.logger = tars.GetLogger(name)
	lock_value, _ := _loggerLocks.LoadOrStore(name, &sync.Mutex{})
	ret.writeLock = lock_value.(*sync.Mutex)
	return ret
}


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


func sesslogFinalizer(l *SessLogger) {
	l.Close()
	return
}


// object functions
func (l *SessLogger) AutoClose() *SessLogger {
	if false == l.autoWrite {
		l.autoWrite = true
		runtime.SetFinalizer(l, sesslogFinalizer)
	}
	return l
}


func write(lock *sync.Mutex, logger *rogger.Logger, logs *list.List) {
	lock.Lock()
	defer lock.Unlock()

	for e := logs.Front(); e != nil; e = e.Next() {
		item := e.Value.(*logitem)
		switch item.level {
		case _DEBUG:
			logger.Debug(item.log)
		case _INFO:
			logger.Info(item.log)
		case _WARN:
			logger.Warn(item.log)
		case _ERROR:
			logger.Error(item.log)
		default:
			// do nothing
		}
	}

	if logs.Len() > 0 {
		logger.Debug("")
	}
	return
}


func (l *SessLogger) Close() {
	logs := l.logs
	l.logs = list.New()
	go write(l.writeLock, l.logger, logs)
	return
}


func (l *SessLogger) Debugf(format string, v ...interface{}) {
	if false == _shouldDebug {
		return
	}
	datetime := getTimeStr()
	file, line, function := getCallerInfo(0);
	text := fmt.Sprintf(format, v...)
	log := datetime + ", " + file + ", Line " + strconv.Itoa(line) + ", " + function + "() - " + text

	l.logs.PushBack(&logitem{
		level: _DEBUG,
		log: log,
	})
	return
}


func (l *SessLogger) Infof(format string, v ...interface{}) {
	if false == _shouldInfo {
		return
	}
	datetime := getTimeStr()
	file, line, function := getCallerInfo(0);
	text := fmt.Sprintf(format, v...)
	log := datetime + ", " + file + ", Line " + strconv.Itoa(line) + ", " + function + "() - " + text

	l.logs.PushBack(&logitem{
		level: _INFO,
		log: log,
	})
	return
}


func (l *SessLogger) Warnf(format string, v ...interface{}) {
	if false == _shouldWarn {
		return
	}
	datetime := getTimeStr()
	file, line, function := getCallerInfo(0);
	text := fmt.Sprintf(format, v...)
	log := datetime + ", " + file + ", Line " + strconv.Itoa(line) + ", " + function + "() - " + text

	l.logs.PushBack(&logitem{
		level: _WARN,
		log: log,
	})
	return
}


func (l *SessLogger) Errorf(format string, v ...interface{}) {
	if false == _shouldError {
		return
	}
	datetime := getTimeStr()
	file, line, function := getCallerInfo(0);
	text := fmt.Sprintf(format, v...)
	log := datetime + ", " + file + ", Line " + strconv.Itoa(line) + ", " + function + "() - " + text

	l.logs.PushBack(&logitem{
		level: _ERROR,
		log: log,
	})
	return
}


func (l *SessLogger) Debug(format string, v ...interface{}) {
	if false == _shouldDebug {
		return
	}
	datetime := getTimeStr()
	file, line, function := getCallerInfo(0);
	text := fmt.Sprintf(format, v...)
	log := datetime + ", " + file + ", Line " + strconv.Itoa(line) + ", " + function + "() - " + text

	l.logs.PushBack(&logitem{
		level: _DEBUG,
		log: log,
	})
	return
}


func (l *SessLogger) Info(format string, v ...interface{}) {
	if false == _shouldInfo {
		return
	}
	datetime := getTimeStr()
	file, line, function := getCallerInfo(0);
	text := fmt.Sprintf(format, v...)
	log := datetime + ", " + file + ", Line " + strconv.Itoa(line) + ", " + function + "() - " + text

	l.logs.PushBack(&logitem{
		level: _INFO,
		log: log,
	})
	return
}


func (l *SessLogger) Warn(format string, v ...interface{}) {
	if false == _shouldWarn {
		return
	}
	datetime := getTimeStr()
	file, line, function := getCallerInfo(0);
	text := fmt.Sprintf(format, v...)
	log := datetime + ", " + file + ", Line " + strconv.Itoa(line) + ", " + function + "() - " + text

	l.logs.PushBack(&logitem{
		level: _WARN,
		log: log,
	})
	return
}


func (l *SessLogger) Error(format string, v ...interface{}) {
	if false == _shouldError {
		return
	}
	datetime := getTimeStr()
	file, line, function := getCallerInfo(0);
	text := fmt.Sprintf(format, v...)
	log := datetime + ", " + file + ", Line " + strconv.Itoa(line) + ", " + function + "() - " + text

	l.logs.PushBack(&logitem{
		level: _ERROR,
		log: log,
	})
	return
}
