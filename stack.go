package gerrors

import (
	"fmt"
	"io"
	"runtime"
	"strconv"
)

//frame 栈帧
type frame uintptr

// stackTrace 调用的栈信息
type stackTrace []frame

//stack 调用的栈集合地址
type stack []uintptr

// pc 返回程序计数器
func (f frame) pc() uintptr { return uintptr(f) - 1 }

// file 返回栈帧执行的文件的名称
func (f frame) file() string {
	fn := runtime.FuncForPC(f.pc())
	if fn == nil {
		return "unknown"
	}
	file, _ := fn.FileLine(f.pc())
	return file
}

// line 返回栈帧执行的行数
func (f frame) line() int {
	fn := runtime.FuncForPC(f.pc())
	if fn == nil {
		return 0
	}
	_, line := fn.FileLine(f.pc())
	return line
}

// name 返回栈帧执行的方法名称
func (f frame) name() string {
	fn := runtime.FuncForPC(f.pc())
	if fn == nil {
		return "unknown"
	}
	return fn.Name()
}

// Format 重新格式化方法
func (f frame) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v', 's':
		_, _ = io.WriteString(s, fmt.Sprintf("%s:%d", f.file(), f.line()))
	case 'd':
		_, _ = io.WriteString(s, strconv.Itoa(f.line()))
	}
}

//stackTrace 获取执行的栈帧集合
func (s *stack) stackTrace() stackTrace {
	f := make([]frame, len(*s))
	for i := 0; i < len(f); i++ {
		f[i] = frame((*s)[i])
	}
	return f
}

//callers 创建栈，默认深度为8
func callers() *stack {
	const depth = 8
	var pcs [depth]uintptr
	n := runtime.Callers(3, pcs[:])
	var st stack = pcs[0:n]
	return &st
}
