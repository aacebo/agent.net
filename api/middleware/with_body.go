package middleware

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/aacebo/agent.net/api/schemas"
	"github.com/go-chi/render"
)

func WithBody[T any](ctx context.Context, name string) func(http.Handler) http.Handler {
	schemas := ctx.Value("schemas").(schemas.Schemas)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var body T
			ctx := r.Context()

			b, err := io.ReadAll(r.Body)

			if err != nil {
				render.Status(r, 400)
				render.JSON(w, r, "invalid body")
				return
			}

			if err := json.Unmarshal(b, &body); err != nil {
				render.Status(r, 400)
				render.JSON(w, r, err.Error())
				return
			}

			_json := map[string]any{}

			if err := json.Unmarshal(b, &_json); err != nil {
				render.Status(r, 400)
				render.JSON(w, r, err.Error())
				return
			}

			if err := schemas.Validate(name, _json); err != nil {
				render.Status(r, 400)
				render.JSON(w, r, err.Error())
				return
			}

			ctx = context.WithValue(ctx, "body", body)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
