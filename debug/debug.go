package debug

import (
	"fmt"
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

func Write(s string) {
	fmt.Println(s)
}

// get is a reference for information
func get[K any, V any](m *sync.Map, key K, def V) V {
	val, _ := m.LoadOrStore(key, def)
	return val.(V)
}
