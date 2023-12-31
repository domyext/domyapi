package webhook

import (
	"math/rand"
	"time"
)

// Initialize a global Rand variable with a secure seed
var globalRand = rand.New(rand.NewSource(time.Now().UnixNano()))
