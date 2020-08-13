package main

import (
	"runtime"
	"strconv"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/sirupsen/logrus"
)

func printMemoryStats() {
	mem := memStats()

	logrus.Infof("\u001b[33m---- Memory Dump ----\u001b[39m")
	logrus.Infof("Allocated: %s", humanize.Bytes(mem.Alloc))
	logrus.Infof("Total Allocated: %s", humanize.Bytes(mem.TotalAlloc))
	logrus.Infof("Memory Allocations: %d", mem.Mallocs)
	logrus.Infof("Memory Frees: %d", mem.Frees)
	logrus.Infof("Heap Allocated: %s", humanize.Bytes(mem.HeapAlloc))
	logrus.Infof("Heap System: %s", humanize.Bytes(mem.HeapSys))
	logrus.Infof("Heap In Use: %s", humanize.Bytes(mem.HeapInuse))
	logrus.Infof("Heap Idle: %s", humanize.Bytes(mem.HeapIdle))
	logrus.Infof("Heap OS Related: %s", humanize.Bytes(mem.HeapReleased))
	logrus.Infof("Heap Objects: %s", humanize.Bytes(mem.HeapObjects))
	logrus.Infof("Stack In Use: %s", humanize.Bytes(mem.StackInuse))
	logrus.Infof("Stack System: %s", humanize.Bytes(mem.StackSys))
	logrus.Infof("Stack Span In Use: %s", humanize.Bytes(mem.MSpanInuse))
	logrus.Infof("Stack Cache In Use: %s", humanize.Bytes(mem.MCacheInuse))
	logrus.Infof("Next GC cycle: %s", humanizeNano(mem.NextGC))
	logrus.Infof("Last GC cycle: %s", humanize.Time(time.Unix(0, int64(mem.LastGC))))
	logrus.Infof("\u001b[33m---- rMemory Stats ----\u001b[39m")
}

func memStats() runtime.MemStats {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	return mem
}

func humanizeNano(n uint64) string {
	var suffix string

	switch {
	case n > 1e9:
		n /= 1e9
		suffix = "s"
	case n > 1e6:
		n /= 1e6
		suffix = "ms"
	case n > 1e3:
		n /= 1e3
		suffix = "us"
	default:
		suffix = "ns"
	}

	return strconv.Itoa(int(n)) + suffix
}
