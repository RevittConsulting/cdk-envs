package buckets

import (
	"context"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Key string

const (
	BucketNameCTX Key = "bucketName"
)

var (
	LastBucketName string
)

func (h *HttpHandler) TryGetBucket(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		bucketName := chi.URLParam(r, "bucketName")
		if bucketName != "" {
			LastBucketName = bucketName
		}
		ctx = context.WithValue(ctx, BucketNameCTX, LastBucketName)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
