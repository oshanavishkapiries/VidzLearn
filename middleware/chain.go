package middleware

import "net/http"

// ChainMiddlewares applies all middleware in order
func Init(handler http.Handler) http.Handler {
	return RecoverMiddleware(
		RateLimitMiddleware(
			HTTPLoggerMiddleware(
				CustomLoggerMiddleware(
					SecureHeadersMiddleware(
						CORSMiddleware(
							handler,
						),
					),
				),
			),
		),
	)
}
