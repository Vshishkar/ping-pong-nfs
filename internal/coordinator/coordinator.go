package coordinator

import "sync"

type Coordinator struct {
	mu sync.RWMutex
}
