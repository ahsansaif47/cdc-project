package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ahsansaif47/cdc-app/config"
	"github.com/ahsansaif47/cdc-app/http/routes"
	"github.com/ahsansaif47/cdc-app/repository/postgres"
	"github.com/ahsansaif47/cdc-app/repository/redis"
	"github.com/gofiber/fiber/v2"
	fl "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	sdkresource "go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.opentelemetry.io/otel/trace"
)

func initTracer() (*sdktrace.TracerProvider, error) {
	exporter, err := otlptrace.New(context.Background(), otlptracehttp.NewClient())
	if err != nil {
		return nil, err
	}

	// Optionally add resource info from environment (service.name, env, etc.)
	res, err := sdkresource.New(
		context.Background(),
		sdkresource.WithFromEnv(),      // Reads OTEL_RESOURCE_ATTRIBUTES
		sdkresource.WithTelemetrySDK(), // Adds SDK version info
		sdkresource.WithHost(),         // Adds host info
		sdkresource.WithAttributes(semconv.ServiceVersion("v0.1.0")),
	)
	if err != nil {
		log.Printf("failed to create resource: %v", err)
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{}, propagation.Baggage{}))
	return tp, err
}

func add(ctx context.Context, a, b int) int {
	_, span := otel.Tracer("go_manual").Start(ctx,
		"add",
		trace.WithAttributes(attribute.Int("a", a)),
		trace.WithAttributes(attribute.Int("b", b)),
	)

	defer span.End()
	return a + b
}

func calculateSeven(ctx context.Context) int {
	newCtx, span := otel.Tracer("go_manual").Start(ctx, "calculateSeven")
	defer span.End()
	return add(newCtx, 3, add(newCtx, 2, 2))
}

func main() {
	tp, err := initTracer()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("failed to shutdown tracer provider: %v", err)
		}
	}()

	calculateSeven(context.TODO())

	db := postgres.GetDatabaseConnection().Pool
	cache := redis.NewCache()

	appCache := redis.NewUserCache(cache)

	startHTTP(db, appCache)

}

func startHTTP(db *pgxpool.Pool, cache redis.ICacheRepository) {
	app := fiber.New()
	// Add logger middleware
	app.Use(fl.New())

	routes.InitRoutes(app, db, cache)

	port := config.GetConfig().ServerPort
	log.Printf("Fiber server listening on port: %s", port)
	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}
