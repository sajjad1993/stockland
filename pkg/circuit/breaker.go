package circuit

import (
	"context"
	"errors"
	"sync"
	"time"
)

type Circuit func(ctx context.Context) (string, error)

func Breaker(circuit Circuit, failedThreshHold uint) Circuit {
	consecutiveFailure := 0
	lastAttempt := time.Now()
	var m sync.RWMutex

	return func(ctx context.Context) (string, error) {
		m.RLock()
		d := consecutiveFailure - int(failedThreshHold)
		if d >= 0 {
			shouldRetryAt := lastAttempt.Add(time.Second * 2 << d)
			if !time.Now().After(shouldRetryAt){
				m.RUnlock()
				return "",errors.New("service is unavailable")

			}
		}
		m.RUnlock()

		response,err := circuit(ctx)
		m.Lock()
		defer m.Unlock()
		lastAttempt = time.Now()
		if err !=nil{
			consecutiveFailure++
			return response,err
		}
		consecutiveFailure = 0
		return response,nil
	}
}
