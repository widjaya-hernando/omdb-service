package server_test

import (
	"context"
	"fmt"
	"log"
	"net"
	"testing"

	pb "github.com/widjaya-hernando/prod-trans-user-service/gen"
	"github.com/widjaya-hernando/prod-trans-user-service/lib/database_transaction"
	"github.com/widjaya-hernando/prod-trans-user-service/transactions/repository"
	"github.com/widjaya-hernando/prod-trans-user-service/transactions/server"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	ctx := context.Background()
	db := setupDatabase()
	var tM database_transaction.Client
	repo := repository.NewRepo(db)

	srv, _ := server.NewServer(ctx, db, repo, tM)

	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	pb.RegisterTransactionsAPIServer(s, srv)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func setupDatabase() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		"localhost",
		"postgres",
		"postgres",
		"warpin_assesment",
		"5432",
	)

	fmt.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
		//Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Println("DSN: ", dsn)
		log.Println("Setup Database Error: ", err)
		panic(err)
	}

	return db
}

func TestGetTransaction(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := pb.NewTransactionsAPIClient(conn)
	req := &pb.GetTransactionRequest{}
	req.Id = 1

	resp, err := client.GetTransaction(ctx, req)
	if err != nil {
		t.Fatalf("SayHello failed: %v", err)
	}
	log.Printf("Response: %+v", resp)
	// Test for output here.
	assert.Equal(t, resp.Transaction.Id, req.Id)
}
