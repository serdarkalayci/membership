package rest

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nicholasjackson/env"
	"github.com/serdarkalayci/membership/api/adapters/comm/htmx"
	"github.com/serdarkalayci/membership/api/application"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	stdout "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	oteltrace "go.opentelemetry.io/otel/trace"
)

var tracer = otel.Tracer("membership-server")

// RestServer handler for getting and updating Ratings
type RestServer struct {
	server *http.Server
	dbContext *application.DataContext
	TracerProvider 	 *sdktrace.TracerProvider
}

// NewRestServer returns a new APIContext handler with the given logger
func NewRestServer(dbContext *application.DataContext) (*RestServer) {
	restServer := &RestServer{
		dbContext: dbContext,
	}	
	return restServer

}

// RunServer prepares the server and runs it
func (restServer *RestServer) RunServer(bindAddress *string) (error) {
	// tp, err := initTracer()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// restServer.TracerProvider = tp
	// defer func() {
	// 	if err := tp.Shutdown(context.Background()); err != nil {
	// 		log.Printf("Error shutting down tracer provider: %v", err)
	// 	}
	// }()
	engine := gin.New()
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5500"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
		  return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	  }))
	engine.Use(otelgin.Middleware("membership-server"))
	var frontEndEnabled = env.Bool("FRONTEND_ENABLED", false, false, "Enable or disable the frontend")
	*frontEndEnabled = true
	if frontEndEnabled != nil && *frontEndEnabled{
		htmx.SetWebRoutes(engine, restServer.dbContext)
	}
	engine.GET("/api/member", restServer.getMembers)
	httpSrv := &http.Server {
        Addr:    *bindAddress,
        Handler: engine,
    }
	// Documentation handler
	// opts := openapimw.RedocOpts{SpecURL: "/swagger.yaml"}
	// sh := openapimw.Redoc(opts, nil)
	// getR.Handle("/docs", sh)
	// getR.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))
	restServer.server = httpSrv
	return httpSrv.ListenAndServe()

}

// Shutdown gracefully shuts down the server
func (restServer *RestServer) Shutdown(ctx context.Context) {
	restServer.server.Shutdown(ctx)
}

func (rs *RestServer)getMembers(c *gin.Context) {
	pageSize := 10
	pageNum := 1
	if c.Query("pageSize") != "" {
		size, err := strconv.Atoi(c.Query("pageSize")); if err == nil {
			pageSize = size
		}
	}
	if c.Query("pageNum") != "" {
		num, err := strconv.Atoi(c.Query("pageNum")); if err == nil {
			pageNum = num
		}
	}
	ms := application.NewMemberService(rs.dbContext)
	members, _, err := ms.ListMembers(pageSize, pageNum, "", 0, 0)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, members)
}

func getUser(c *gin.Context, id string) string {
	// Pass the built-in `context.Context` object from http.Request to OpenTelemetry APIs
	// where required. It is available from gin.Context.Request.Context()
	_, span := tracer.Start(c.Request.Context(), "getUser", oteltrace.WithAttributes(attribute.String("id", id)))
	defer span.End()
	if id == "123" {
		return "otelgin tester"
	}
	return "unknown"
}
// createSpan extracts the span from the request if exists or creates a new one using openTelemetry. Span with the given name and returns it
func createSpan(ctx context.Context, opName string, r *http.Request) (context.Context, trace.Span) {
	spanContext := otel.GetTextMapPropagator().Extract(
		ctx,
		propagation.HeaderCarrier(r.Header))

	ctx, span := otel.Tracer("BookStore").Start(
		spanContext,
		opName,
	)
	return ctx, span
}

func initTracer() (*sdktrace.TracerProvider, error) {
	exporter, err := stdout.New(stdout.WithPrettyPrint())
	if err != nil {
		return nil, err
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return tp, nil
}


