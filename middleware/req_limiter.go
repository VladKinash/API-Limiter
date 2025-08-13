package middleware
import (
	"net/http"
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
	ApiKeys  map[ApiKey]int
	Name     string
	mu       sync.Mutex
}

func NewApiKey(key string, requests int) ApiKey {
	return ApiKey{
		Key:      key,
		Requests: requests,
		LastSeen: time.Now(),
	}
}

func (a *ApiKey) SetRequests(requests int){
	a.mu.Lock()
	defer a.mu.Unlock()
	a.Requests = requests
}

func (a *ApiKey) SetLastSeen(time time.Time){
	a.mu.Lock()
	defer a.mu.Unlock()
	a.LastSeen = time
}



func NewLimiter(requests int, name string) Limiter {
	return Limiter{
		Requests: requests,
		ApiKeys:  nil,
		Name:     name,
	}
}

func (l *Limiter) SetRequest(limit int){
	l.mu.Lock()
	defer l.mu.Unlock()
	l.Requests = limit
}

func ApiKeyHandler(apiKey string) bool{
	return true
}

// ReqLimiter is a basic request limiter middleware.
func ReqLimiter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//apiKey := r.Header.Get("APIKEY")
		next.ServeHTTP(w, r)
	})
}