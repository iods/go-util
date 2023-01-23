package debug

import (
	"io"
	"math/rand"
	"os"
	"sync"
	"time"
)

var (
	output     io.Writer = os.Stdout
	outputOnce sync.Once

	// fill and update with new random one
	random = rand.New(rand.NewSource(time.Now().Unix()))
)

// SetOutput sets the output for the debug util package. Can be set once.
func SetOutput(w io.Writer) {
	outputOnce.Do(func() {
		output = w
	})
}

// DebugType declares the labeling for the debugging type.
type DebugType string

// DebugValue declares the level for DebugType label.
type DebugValue int32

// DebugLevels declares the type and level for debugging.
type DebugLevels map[DebugType]DebugValue

// level sets a global default for debugging levels.
var level DebugValue = value_info

const (
	value_info  DebugValue = 1
	value_debug DebugValue = 2
	value_warn  DebugValue = 3
	value_error DebugValue = 4
	value_none  DebugValue = 5
)

const (
	Debug_Info  DebugType = "info"
	Debug_Debug DebugType = "debug"
	Debug_Warn  DebugType = "warn"
	Debug_Error DebugType = "error"
	Debug_None  DebugType = "none"
)

// caller is the default for what level will be called.
const (
	caller = 2
)

const (
	defaultOutput      = "[%s][%s] %s\n"
	outputWithFileline = "[%s][%s] %s in file %s on line %d\n"
)

// levels map of the DebugLevels.
var levels = DebugLevels{
	Debug_Info:  value_info,
	Debug_Debug: value_debug,
	Debug_Warn:  value_warn,
	Debug_Error: value_error,
}

// Allowed returns true or false if the log level is greater than or equal to value_info.
func (d DebugLevels) Allowed(t DebugType) bool {
	if l, ok := d[t]; ok {
		return l >= level
	}
	return false
}

// Set will set the level var with the provided type.
func Set(t DebugType) {
	level = levels.Get(t)
}

// Get returns the DebugValue for the provided DebugType.
func (d DebugLevels) Get(t DebugType) DebugValue {
	if v, ok := d[t]; ok {
		return v
	}
	return value_none
}

func echo(format string, t DebugType, args ...interface{}) {
	if !levels.Allowed(t) {
		return
	}
	switch t {
	case Debug_Warn:
		if lineSwitch == ShowFileline {

		}
	default:
		if lineSwitch == ShowFileline {

		}
	}
}

// getsss is a reference for information
func getsss[K any, V any](m *sync.Map, key K, def V) V {
	val, _ := m.LoadOrStore(key, def)
	return val.(V)
}
