// Package pushgateway
package pushgateway

import (
	"bytes"
	"fmt"
	"time"

	"github.com/jellydator/ttlcache/v3"
)
// version set pushgateway version
const version = "202406091624"

// P struct data
type P struct {
	cache *ttlcache.Cache[string, string]
}

// New create New client
func New() *P {
	cache := ttlcache.New[string, string](
		ttlcache.WithTTL[string, string](1 * time.Minute),
	)
	return &P{cache: cache}
}

// Set store metrics data
func (p *P) Set(job, instance string, data []byte) {
	entersep := []byte("#")[0]
	bracketsep := []byte("} ")
	lines := bytes.Split(data, []byte("\n"))
	metrics := ""
	for _, line := range lines {
		if len(line) == 0 || line[0] == entersep {
			continue
		}
		b := bytes.Split(line, bracketsep)
		if len(b) != 2 {
			continue
		}
		metric := fmt.Sprintf("%s,exported_instance=\"%s\",exported_job=\"%s\"} %s\n", b[0], instance, job, b[1])
		metrics += metric
	}
	p.cache.Set(job+instance, metrics, 1*time.Minute)
}

// Format get all metrics data
func (p *P) Format() []byte {
	l := fmt.Sprintf("# version %s\n# time %s.\n", version, time.Now())
	for _, v := range p.cache.Keys() {
		l += p.cache.Get(v).Value()
	}
	return []byte(l)
}