package main

import "sync"

type Coordinator struct {
	mu sync.RWMutex
}
