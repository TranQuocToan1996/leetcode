package uti

import (
	"encoding/json"
	"log"
	"runtime"
)

func MemoryAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	log.Printf("Memory: %d KB\n", m.Alloc/1024)
	buf, _ := json.MarshalIndent(m, "", "\t")
	if len(buf) > 0 {
		log.Println(string(buf))
	}
}
