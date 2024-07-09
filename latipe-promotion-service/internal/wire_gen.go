// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package server

import (
	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/hellofresh/health-go/v5"
	"google.golang.org/grpc"
	"latipe-promotion-services/config"
	"latipe-promotion-services/internal/adapter/storeserv"
	"latipe-promotion-services/internal/adapter/userserv"
	"latipe-promotion-services/internal/api"
	"latipe-promotion-services/internal/domain/repos"
	"latipe-promotion-services/internal/grpcservice/interceptor"
	"latipe-promotion-services/internal/grpcservice/vouchergrpc"
	"latipe-promotion-services/internal/middleware"
	"latipe-promotion-services/internal/publisher/purchaseCreate"
	"latipe-promotion-services/internal/router"
	"latipe-promotion-services/internal/services/voucherserv"
	"latipe-promotion-services/internal/subs/createPurchase"
	"latipe-promotion-services/pkgs/healthcheck"
	"latipe-promotion-services/pkgs/mongodb"
	"latipe-promotion-services/pkgs/rabbitclient"
	"latipe-promotion-services/pkgs/response"
)

// Injectors from server.go:

func New() (*Server, error) {
	configConfig, err := config.NewConfig()
	if err != nil {
		return nil, err
	}
	mongoClient, err := mongodb.OpenMongoDBConnection(configConfig)
	if err != nil {
		return nil, err
	}
	voucherRepository := repos.NewVoucherRepos(mongoClient)
	connection := rabbitclient.NewRabbitClientConnection(configConfig)
	replyPurchaseTransactionPub := purchaseCreate.NewReplyPurchaseTransactionPub(configConfig, connection)
	voucherService := voucherserv.NewVoucherService(voucherRepository, replyPurchaseTransactionPub)
	voucherServiceServer := vouchergrpc.NewVoucherServerGRPC(voucherService)
	voucherHandle := api.NewVoucherHandler(voucherService)
	userService := userserv.NewUserService(configConfig)
	storeService := storeserv.NewStoreServiceAdapter(configConfig)
	authMiddleware := middleware.NewAuthMiddleware(userService, storeService)
	voucherRouter := router.NewVoucherRouter(voucherHandle, authMiddleware)
	grpcInterceptor := interceptor.NewGrpcInterceptor(configConfig)
	purchaseCreateSubscriber := createPurchase.NewPurchaseCreateSubscriber(configConfig, voucherService, connection)
	purchaseRollbackSubscriber := createPurchase.NewPurchaseRollbackSubscriber(configConfig, voucherService, connection)
	server := NewServer(configConfig, voucherServiceServer, voucherRouter, grpcInterceptor, purchaseCreateSubscriber, purchaseRollbackSubscriber)
	return server, nil
}

// server.go:

type Server struct {
	app                  *fiber.App
	cfg                  *config.Config
	grpcServ             *grpc.Server
	purchaseSubs         *createPurchase.PurchaseCreateSubscriber
	rollbackPurchaseSubs *createPurchase.PurchaseRollbackSubscriber
}

func NewServer(
	cfg *config.Config,
	voucherGrpc vouchergrpc.VoucherServiceServer,
	voucherRouter router.VoucherRouter,
	grpcInterceptor *interceptor.GrpcInterceptor,
	purchaseSubs *createPurchase.PurchaseCreateSubscriber,
	rollbackPurchaseSubs *createPurchase.PurchaseRollbackSubscriber) *Server {

	app := fiber.New(fiber.Config{
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		JSONDecoder:  sonic.Unmarshal,
		JSONEncoder:  sonic.Marshal,
		ErrorHandler: responses.CustomErrorHandler,
	})

	recoverConfig := recover2.ConfigDefault
	app.Use(recover2.New(recoverConfig))

	app.Use(logger.New())
	api2 := app.Group("/api")
	v1 := api2.Group("/v1")

	basicAuthConfig := basicauth.Config{
		Users: map[string]string{
			cfg.Metrics.Username: cfg.Metrics.Password,
		},
	}

	prometheus := fiberprometheus.New("promotion-services")
	prometheus.RegisterAt(app, cfg.Metrics.Host, basicauth.New(basicAuthConfig))
	app.Use(prometheus.Middleware)
	app.Use(logger.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		s := struct {
			Message string `json:"message"`
			Version string `json:"version"`
		}{
			Message: "Promotion services was developed by TienDat",
			Version: "v1.0.1",
		}
		return ctx.JSON(s)
	})

	h, _ := healthService.NewHealthCheckService(cfg)
	app.Get("/health", basicauth.New(basicAuthConfig), adaptor.HTTPHandlerFunc(h.HandlerFunc))
	app.Use(healthcheck.New(healthcheck.Config{
		LivenessProbe: func(c *fiber.Ctx) bool {
			return true
		},
		LivenessEndpoint: "/liveness",
		ReadinessProbe: func(c *fiber.Ctx) bool {
			result := h.Measure(c.Context())
			return result.Status == health.StatusOK
		},
		ReadinessEndpoint: "/readiness",
	}))

	app.Get(cfg.Metrics.FiberDashboard, basicauth.New(basicAuthConfig), monitor.New(monitor.Config{Title: "Promotion Services Metrics Page (Fiber)"}))
	h, _ = healthService.NewHealthCheckService(cfg)
	app.Get("/health", basicauth.New(basicAuthConfig), adaptor.HTTPHandlerFunc(h.HandlerFunc))
	app.Use(healthcheck.New(healthcheck.Config{
		LivenessProbe: func(c *fiber.Ctx) bool {
			return true
		},
		LivenessEndpoint: "/liveness",
		ReadinessProbe: func(c *fiber.Ctx) bool {
			result := h.Measure(c.Context())
			return result.Status == health.StatusOK
		},
		ReadinessEndpoint: "/readiness",
	}))

	app.Get(cfg.Metrics.FiberDashboard, basicauth.New(basicAuthConfig), monitor.New(monitor.Config{Title: "Promotion Services Metrics Page (Fiber)"}))

	voucherRouter.Init(&v1)

	grpcServ := grpc.NewServer(grpc.UnaryInterceptor(grpcInterceptor.MiddlewareUnaryRequest))
	vouchergrpc.RegisterVoucherServiceServer(grpcServ, voucherGrpc)

	return &Server{
		cfg:                  cfg,
		app:                  app,
		purchaseSubs:         purchaseSubs,
		rollbackPurchaseSubs: rollbackPurchaseSubs,
		grpcServ:             grpcServ,
	}
}

func (serv Server) App() *fiber.App {
	return serv.app
}

func (serv Server) Config() *config.Config {
	return serv.cfg
}

func (serv Server) GrpcServ() *grpc.Server {
	return serv.grpcServ
}

func (serv Server) CommitPurchaseTransactionSubscriber() *createPurchase.PurchaseCreateSubscriber {
	return serv.purchaseSubs
}

func (serv Server) RollbackPurchaseTransactionSubscriber() *createPurchase.PurchaseRollbackSubscriber {
	return serv.rollbackPurchaseSubs
}