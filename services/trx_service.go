package services

import (
	"MiniProject-MuhammadRendyansyahDwiAnggara/entities"
	"MiniProject-MuhammadRendyansyahDwiAnggara/repositories"
	"errors"
)

type TransactionService interface {
	Create(userID uint, req entities.CreateTransactionRequest) (uint, error)
	GetMyTransactions(userID uint, search string, page, limit int) ([]entities.Transaction, error)
	GetByID(id, userID uint) (*entities.Transaction, error)
}

type transactionService struct {
	repo repositories.TransactionRepository
}

func NewTransactionService(repo repositories.TransactionRepository) TransactionService {
	return &transactionService{repo}
}

func (s *transactionService) Create(userID uint, req entities.CreateTransactionRequest) (uint, error) {
	return s.repo.CreateTransaction(userID, req)
}

func (s *transactionService) GetMyTransactions(userID uint, search string, page, limit int) ([]entities.Transaction, error) {
	offset := 0
	if page > 0 && limit > 0 {
		offset = (page - 1) * limit
	}
	return s.repo.FindAllByUserID(userID, search, limit, offset)
}

func (s *transactionService) GetByID(id, userID uint) (*entities.Transaction, error) {
	trx, err := s.repo.FindByIDAndUserID(id, userID)
	if err != nil {
		return nil, errors.New("No Data Trx")
	}
	return trx, nil
}