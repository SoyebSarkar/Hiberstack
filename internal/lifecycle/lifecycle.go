package lifecycle

import (
	"time"
)

func ShouldOffload(lastAccess time.Time, ttl time.Duration) bool {
	return time.Since(lastAccess) > ttl
}
