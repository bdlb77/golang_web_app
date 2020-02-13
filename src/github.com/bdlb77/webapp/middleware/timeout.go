package middleware

import (
	"context"
	"net/http"
	"time"
)

type TimeoutMiddleware struct {
	Next http.Handler
}

func (tm *TimeoutMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if tm.Next == nil {
		tm.Next = http.DefaultServeMux
	}
	// create context
	ctx := r.Context()
	// withTImeout.. create timeout and create new context
	ctx, _ = context.WithTimeout(ctx, 3*time.Second)
	// pass new ctx to request
	r.WithContext(ctx)

	// create channel
	ch := make(chan struct{})
	// send go routine
	go func() {
		// send regular response to next handler
		tm.Next.ServeHTTP(w, r)
		// pull out value (classic struct value to show that request finished)
		ch <- struct{}{}
	}()

	select {
	// if 1st go channel finishes 1st , return regularly
	// if ctx.Done() channel outputs(finishes) 1st.. respond with 408
	case <-ch:
		return
	case <-ctx.Done():
		w.WriteHeader(http.StatusRequestTimeout)
	}
}
