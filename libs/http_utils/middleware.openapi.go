package http_utils

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getkin/kin-openapi/routers/gorillamux"
)

func handleError(prefix string, err error, w http.ResponseWriter, r *http.Request) {
	fmt.Printf("OpenAPI %s %s\n", prefix, err)
	w.WriteHeader(400)
	w.Write([]byte("Bad request"))
}

func OpenAPIMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx := context.Background()
		loader := &openapi3.Loader{Context: ctx, IsExternalRefsAllowed: true}
		doc, docErr := loader.LoadFromFile(os.Getenv(("OPENAPI_PATH")))

		if docErr != nil {
			handleError("OpenAPI document error", docErr, w, r)
			return
		}

		doc.Validate(ctx)

		router, routerErr := gorillamux.NewRouter(doc)

		if routerErr != nil {
			handleError("Router error", routerErr, w, r)
			return
		}

		route, pathParams, routeErr := router.FindRoute(r)

		if routeErr != nil {
			handleError("Route error", routeErr, w, r)
			return
		}

		err := openapi3filter.ValidateRequest(ctx, &openapi3filter.RequestValidationInput{
			Request:    r,
			PathParams: pathParams,
			Route:      route,
		})

		if err != nil {
			handleError("Validation error", err, w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}
