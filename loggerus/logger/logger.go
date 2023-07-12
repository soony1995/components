package logger

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

var (
	LogToFile *logrus.Logger = logrus.New() // 대문자로 하는 이유는 다른 컴포넌트에서 사용할 수 도 있기 때문이다.
	StdToFile *logrus.Logger = logrus.New()

	LoggerFile *os.File
)

func InitLogrus() {
	f, err := os.OpenFile("./detail.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Failed to create logfile" + "log.txt")
		panic(err)
	}

	LogToFile = &logrus.Logger{
		// Log into f file handler and on os.LogToFileout

		// io.Writer는 인터페이스이고, os.file은 io.Writer를 구현한 객체이기 때문에 사용이 가능하다.
		// Out:   io.MultiWriter(f, os.LogToFileout),

		Out:   f,
		Level: logrus.DebugLevel,
		Formatter: &logrus.TextFormatter{
			DisableColors: false,
			ForceColors:   true,
		},
		ReportCaller: true,
	}

	LoggerFile = f

	LogToFile.Trace("Trace message")
	LogToFile.Info("Info message")
	LogToFile.Warn("Warn message")
	LogToFile.Error("Error message")
}

func StdLogrus() {
	StdToFile = &logrus.Logger{
		// Log into f file handler and on os.LogToFileout

		Out:   os.Stdout,
		Level: logrus.DebugLevel,
		Formatter: &logrus.TextFormatter{
			DisableColors: false,
			ForceColors:   true,
		},
		ReportCaller: true,
	}
}

type User struct {
	Information string `json:"information"`
	Port        uint16 `json:"port"`
	Status      bool   `json:"status"`
}

func NewUser() *User {
	return &User{
		Information: "sooncheol",
		Port:        3535,
	}
}

// Logrus logging Standard Output Print (Info level)
func (u *User) Print(args ...interface{}) {
	args = append(append([]interface{}{}, u.Information), args...)
	// LogToFile.Print(args...)
	LogToFile.Print(args...)
}
func (u *User) Printf(format string, args ...interface{}) {
	// LogToFile.Printf(format, args...)
	LogToFile.Printf(u.Information+" "+format, args...)
}
func (u *User) Println(args ...interface{}) {
	args = append(append([]interface{}{}, u.Information), args...)
	// LogToFile.Println(args...)
	LogToFile.Println(args...)
}

// Logrus logging Panic
func (u *User) Panic(args ...interface{}) {
	args = append(append([]interface{}{}, u.Information), args...)
	LogToFile.Panic(args...)
	StdToFile.Panic(args...)
}
func (u *User) Panicf(format string, args ...interface{}) {
	LogToFile.Panicf(u.Information+" "+format, args...)
	StdToFile.Panicf(u.Information+" "+format, args...)
}
func (u *User) Panicln(args ...interface{}) {
	args = append(append([]interface{}{}, u.Information), args...)
	LogToFile.Panicln(args...)
	StdToFile.Panicln(args...)
}

// Logrus logging Fatal
func (u *User) Fatal(args ...interface{}) {
	args = append(append([]interface{}{}, u.Information), args...)
	LogToFile.Fatal(args...)
	StdToFile.Fatal(args...)
}
func (u *User) Fatalf(format string, args ...interface{}) {
	LogToFile.Fatalf(u.Information+" "+format, args...)
	StdToFile.Fatalf(u.Information+" "+format, args...)
}
func (u *User) Fatalln(args ...interface{}) {
	args = append(append([]interface{}{}, u.Information), args...)
	LogToFile.Fatalln(args...)
	StdToFile.Fatalln(args...)
}

// Logrus logging Error
func (u *User) Error(args ...interface{}) {
	args = append(append([]interface{}{}, u.Information), args...)
	LogToFile.Error(args...)
	StdToFile.Error(args...)
}
func (u *User) Errorf(format string, args ...interface{}) {
	LogToFile.Errorf(u.Information+" "+format, args...)
	StdToFile.Errorf(u.Information+" "+format, args...)
}
func (u *User) Errorln(args ...interface{}) {
	args = append(append([]interface{}{}, u.Information), args...)
	LogToFile.Errorln(args...)
	StdToFile.Errorln(args...)
}

// Logrus logging Warn
func (u *User) Warn(args ...interface{}) {
	args = append(append([]interface{}{}, u.Information), args...)
	LogToFile.Warn(args...)
	StdToFile.Warn(args...)
}
func (u *User) Warnf(format string, args ...interface{}) {
	LogToFile.Warnf(u.Information+" "+format, args...)
	StdToFile.Warnf(u.Information+" "+format, args...)
}
func (u *User) Warnln(args ...interface{}) {
	args = append(append([]interface{}{}, u.Information), args...)
	LogToFile.Warnln(args...)
	StdToFile.Warnln(args...)
}

// Logrus logging Info
func (u *User) Info(args ...interface{}) {
	args = append(append([]interface{}{}, u.Information), args...)
	LogToFile.Info(args...)
	StdToFile.Info(args...)
}
func (u *User) Infof(format string, args ...interface{}) {
	LogToFile.Infof(u.Information+" "+format, args...)
	StdToFile.Infof(u.Information+" "+format, args...)
}
func (u *User) Infoln(args ...interface{}) {
	args = append(append([]interface{}{}, u.Information), args...)
	LogToFile.Infoln(args...)
	StdToFile.Infoln(args...)
}

// Logrus logging Debug
func (u *User) Debug(args ...interface{}) {
	args = append(append([]interface{}{}, u.Information), args...)
	LogToFile.Debug(args...)
	StdToFile.Debug(args...)
}
func (u *User) Debugf(format string, args ...interface{}) {
	LogToFile.Debugf(u.Information+" "+format, args...)
	StdToFile.Debugf(u.Information+" "+format, args...)
}
func (u *User) Debugln(args ...interface{}) {
	args = append(append([]interface{}{}, u.Information), args...)
	LogToFile.Debugln(args...)
	StdToFile.Debugln(args...)
}

// Logrus logging Trace
func (u *User) Trace(args ...interface{}) {
	args = append(append([]interface{}{}, u.Information), args...)
	LogToFile.Trace(args...)
	StdToFile.Trace(args...)
}
func (u *User) Tracef(format string, args ...interface{}) {
	LogToFile.Tracef(u.Information+" "+format, args...)
	StdToFile.Tracef(u.Information+" "+format, args...)
}
func (u *User) Traceln(args ...interface{}) {
	args = append(append([]interface{}{}, u.Information), args...)
	LogToFile.Traceln(args...)
	StdToFile.Traceln(args...)
}

/***
* @ Standard log print out
 */
// Logrus logging Standard Output Print (Info level)
func Print(args ...interface{}) {
	// LogToFile.Print(args...)
	LogToFile.Print(args...)
}
func Printf(format string, args ...interface{}) {
	// LogToFile.Printf(format, args...)
	LogToFile.Printf(format, args...)
}
func Println(args ...interface{}) {
	// LogToFile.Println(args...)
	LogToFile.Println(args...)
}

// Logrus logging Panic
func Panic(args ...interface{}) {
	LogToFile.Panic(args...)
	StdToFile.Panic(args...)
}
func Panicf(format string, args ...interface{}) {
	LogToFile.Panicf(format, args...)
	StdToFile.Panicf(format, args...)
}
func Panicln(args ...interface{}) {
	LogToFile.Panicln(args...)
	StdToFile.Panicln(args...)
}

// Logrus logging Fatal
func Fatal(args ...interface{}) {
	LogToFile.Fatal(args...)
	StdToFile.Fatal(args...)
}
func Fatalf(format string, args ...interface{}) {
	LogToFile.Fatalf(format, args...)
	StdToFile.Fatalf(format, args...)
}
func Fatalln(args ...interface{}) {
	LogToFile.Fatalln(args...)
	StdToFile.Fatalln(args...)
}

// Logrus logging Error
func Error(args ...interface{}) {
	LogToFile.Error(args...)
	StdToFile.Error(args...)
}
func Errorf(format string, args ...interface{}) {
	LogToFile.Errorf(format, args...)
	StdToFile.Errorf(format, args...)
}
func Errorln(args ...interface{}) {
	LogToFile.Errorln(args...)
	StdToFile.Errorln(args...)
}

// Logrus logging Warn
func Warn(args ...interface{}) {
	LogToFile.Warn(args...)
	StdToFile.Warn(args...)
}
func Warnf(format string, args ...interface{}) {
	LogToFile.Warnf(format, args...)
	StdToFile.Warnf(format, args...)
}
func Warnln(args ...interface{}) {
	LogToFile.Warnln(args...)
	StdToFile.Warnln(args...)
}

// Logrus logging Info
func Info(args ...interface{}) {
	LogToFile.Info(args...)
	StdToFile.Info(args...)
}
func Infof(format string, args ...interface{}) {
	LogToFile.Infof(format, args...)
	StdToFile.Infof(format, args...)
}
func Infoln(args ...interface{}) {
	LogToFile.Infoln(args...)
	StdToFile.Infoln(args...)
}

// Logrus logging Debug
func Debug(args ...interface{}) {
	LogToFile.Debug(args...)
	StdToFile.Debug(args...)
}
func Debugf(format string, args ...interface{}) {
	LogToFile.Debugf(format, args...)
	StdToFile.Debugf(format, args...)
}
func Debugln(args ...interface{}) {
	LogToFile.Debugln(args...)
	StdToFile.Debugln(args...)
}

// Logrus logging Trace
func Trace(args ...interface{}) {
	LogToFile.Trace(args...)
	StdToFile.Trace(args...)
}
func Tracef(format string, args ...interface{}) {
	LogToFile.Tracef(format, args...)
	StdToFile.Tracef(format, args...)
}
func Traceln(args ...interface{}) {
	LogToFile.Traceln(args...)
	StdToFile.Traceln(args...)
}
