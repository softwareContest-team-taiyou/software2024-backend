package infrastructure

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	todov1 "github.com/softwareContest-team-taiyou/software2024-backend/gen/go/v1/todo"
	"github.com/softwareContest-team-taiyou/software2024-backend/internal/domain/repository"
	"github.com/softwareContest-team-taiyou/software2024-backend/internal/handler"
	"github.com/softwareContest-team-taiyou/software2024-backend/internal/usecase"
	"github.com/softwareContest-team-taiyou/software2024-backend/middleware/auth0"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
)

func Router() {
	port := os.Getenv("PORT")
	domain := os.Getenv("AUTH0_DOMAIN")
	clientID := os.Getenv("AUTH0_CLIENT_ID")
	jwks, err := auth0.FetchJWKS(domain)
	 if err != nil {
        log.Fatal(err)
    }
    // domain, clientID, 公開鍵を元にJWTMiddlewareを作成する
   

	databaseHandler := NewDatabaseHandler()

	if port == "" {
		log.Fatal("PORT environment variable not set.")
	}
	// 直接ポート文字列を使用
	address := fmt.Sprintf(":%s", port)
	listener, err := net.Listen("tcp",  address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	todoHandler := handler.NewTodoHandler(usecase.NewTodoUsecase(repository.NewTodoRepository(databaseHandler)))
	opts := []grpc_zap.Option{
		grpc_zap.WithDurationField(func(duration time.Duration) zapcore.Field {
				return zap.Int64("grpc.time_ns", duration.Nanoseconds())
		}),
	}
	var zapLogger *zap.Logger
	if os.Getenv("ENV") == "development" {
		// Development environment: detailed logging
		zapLogger, _ = zap.NewDevelopment()
	} else {
		// Production environment: minimal logging
		zapLogger, _ = zap.NewProduction()
	}
	grpc_zap.ReplaceGrpcLogger(zapLogger)
	srv := grpc.NewServer(   
		grpc_middleware.WithUnaryServerChain(
        grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
        grpc_zap.UnaryServerInterceptor(zapLogger, opts...), 
		     
	),grpc.UnaryInterceptor(auth0.AuthInterceptor(domain,clientID,jwks)),)

	todov1.RegisterTodoServiceServer(srv, todoHandler)
		// ログを出力するmiddlewareを実行
	
	if err := srv.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
	}
    
}