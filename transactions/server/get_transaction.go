package server

import (
	"context"
	"log"

	pb "github.com/widjaya-hernando/prod-trans-user-service/gen"
	"github.com/widjaya-hernando/prod-trans-user-service/models"
)

func (s *Server) GetTransaction(ctx context.Context, req *pb.GetTransactionRequest) (*pb.GetTransactionResponse, error) {
	var response pb.GetTransactionResponse

	transaction, err := s.repository.GetTransaction(nil, req.GetId())
	if err != nil {
		log.Println("err-get-transaction: ", err)
		return nil, err
	}

	response.Transaction = models.TransactionToGrpcMessage(transaction)

	return &response, nil
}
