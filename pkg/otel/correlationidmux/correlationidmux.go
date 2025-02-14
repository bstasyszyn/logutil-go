package correlationidmux

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"github.com/trustbloc/logutil-go/pkg/log"
	"github.com/trustbloc/logutil-go/pkg/otel/api"
	"github.com/trustbloc/logutil-go/pkg/otel/correlationid"
)

var logger = log.New("correlationid-mux")

func Middleware() mux.MiddlewareFunc {
	return func(handler http.Handler) http.Handler {
		return &MuxMiddleware{
			handler: handler,
		}
	}
}

type MuxMiddleware struct {
	handler http.Handler
}

func (tw *MuxMiddleware) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	// TODO: REMOVE!!!
	for key, h := range req.Header {
		fmt.Printf("----- correlationidmux - Header: %s, value: %s\n", key, h)
	}

	correlationID := req.Header.Get(api.CorrelationIDHeader)
	if correlationID != "" {
		// TODO: Change to DEBUG
		logger.Infoc(ctx, "Received HTTP request with correlation ID in header", log.WithCorrelationID(correlationID))

		var err error
		ctx, err = correlationid.SetWithValue(ctx, correlationID)
		if err != nil {
			logger.Warnc(ctx, "Failed to set correlation ID in context", log.WithError(err))
		}
	} else {
		var err error
		ctx, correlationID, err = correlationid.Set(ctx)
		if err != nil {
			logger.Warnc(ctx, "Failed to set correlation ID in context", log.WithError(err))
		} else {
			// TODO: Change to DEBUG
			logger.Infoc(ctx, "Generated new correlation ID since none was found in the HTTP header")
		}
	}

	if correlationID != "" {
		span := trace.SpanFromContext(ctx)
		span.SetAttributes(attribute.String(api.CorrelationIDAttribute, correlationID))
	}

	tw.handler.ServeHTTP(w, req.WithContext(ctx))
}
