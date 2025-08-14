package middleware

import (
	"net/http"
	"slices"
	"sync"
	"time"
)

type ApiKey struct {
	Key      string
	Requests int
	mu       sync.Mutex
	LastSeen time.Time
}

type Limiter struct {
	Requests int
	ApiKeys  map[string]*ApiKey
	Name     string
	mu       sync.Mutex
}

func NewApiKey(key string, requests int) *ApiKey {
	return &ApiKey{
		Key:      key,
		Requests: requests,
		LastSeen: time.Now(),
	}
}

func (l *Limiter) addApiKey(apiKey *ApiKey) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.ApiKeys[apiKey.Key] = apiKey
}

func (a *ApiKey) SetRequests(requests int) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.Requests = requests
}

func (a *ApiKey) SetLastSeen(time time.Time) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.LastSeen = time
}

func NewLimiter(requests int, name string) Limiter {
	return Limiter{
		Requests: requests,
		ApiKeys:  make(map[string]*ApiKey),
		Name:     name,
	}
}

func (l *Limiter) SetRequest(limit int) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.Requests = limit
}

func (a *ApiKey) addReq(newReq int) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.Requests += newReq
}

func (l *Limiter) inLimiter(apiKey string) *ApiKey {
	l.mu.Lock()
	defer l.mu.Unlock()
	if k, ok := l.ApiKeys[apiKey]; ok {
		return k
	}
	return nil
}

func ApiKeyHandler(apiKey string, apiLimiter *Limiter, validApiKeys []string, limit int) int {
	found := slices.Contains(validApiKeys, apiKey)

	if found {
		foundApiKey := apiLimiter.inLimiter(apiKey)
		if foundApiKey == nil {
			newApiKey := NewApiKey(apiKey, 1)
			apiLimiter.addApiKey(newApiKey)
			return 200
		} else if foundApiKey.Requests < limit {
			foundApiKey.addReq(1)
			return 200
		}
		return 429
	}

	return 404
}

// ReqLimiter is a basic request limiter middleware.
func ReqLimiter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//apiKey := r.Header.Get("APIKEY")
		next.ServeHTTP(w, r)
	})
}
