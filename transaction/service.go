package transaction

import (
	"errors"
	"ngevent-api/event"
	"ngevent-api/helper"
	"ngevent-api/payment"
	"ngevent-api/user"
	"strconv"
)

type service struct {
	repository      Repository
	eventRepository event.Repository
	paymentService  payment.Service
}

func NewService(repository Repository, eventRepository event.Repository, paymentService payment.Service) *service {
	return &service{repository, eventRepository, paymentService}
}

type Service interface {
	CreateTransaction(input CreateTransactionInput) (Transaction, error)
	GetTransactions(params helper.Pagination, user user.User) ([]Transaction, *helper.Pagination, error)
	GetTransaction(input GetTransactionDetailInput, user user.User) (Transaction, error)
	ProcessPayment(input TransactionNotificationInput) error
}

func (s *service) CreateTransaction(input CreateTransactionInput) (Transaction, error) {
	findEvent, err := s.eventRepository.FindById(input.EventId)
	if err != nil {
		return Transaction{}, err
	}

	if findEvent.Id == 0 {
		return Transaction{}, errors.New("event not found")
	}

	if findEvent.Quota < 1 {
		return Transaction{}, errors.New("event sold out")
	}

	transaction := Transaction{}
	transaction.EventId = input.EventId
	transaction.OrderId = input.OrderId
	transaction.UserId = input.User.Id
	transaction.Amount = findEvent.Price * input.Qty
	transaction.Qty = input.Qty
	transaction.Status = "pending"
	transaction.Code = input.Code
	transaction.PaymentUrl = input.PaymentUrl

	newTransaction, err := s.repository.Save(transaction)
	if err != nil {
		return newTransaction, err
	}

	paymentTransaction := payment.Transaction{
		OrderId: newTransaction.OrderId,
		Amount:  newTransaction.Amount,
	}

	paymentUrl, err := s.paymentService.GetPaymentURL(paymentTransaction, input.User)
	if err != nil {
		return newTransaction, err
	}

	newTransaction.PaymentUrl = paymentUrl

	newTransaction, err = s.repository.Update(newTransaction)
	if err != nil {
		return newTransaction, err
	}

	return newTransaction, nil
}

func (s *service) GetTransactions(params helper.Pagination, user user.User) ([]Transaction, *helper.Pagination, error) {
	transaction, paginate, err := s.repository.FindAll(params, user.UserType, user.Id)

	if err != nil {
		return transaction, paginate, err
	}

	return transaction, paginate, nil
}

func (s *service) GetTransaction(input GetTransactionDetailInput, user user.User) (Transaction, error) {
	transaction, err := s.repository.FindById(input.Id, user.UserType, user.Id)

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (s *service) ProcessPayment(input TransactionNotificationInput) error {
	var user user.User
	transaction_id, _ := strconv.Atoi(input.OrderID)

	transaction, err := s.repository.FindById(transaction_id, user.UserType, user.Id)
	if err != nil {
		return err
	}

	if input.PaymentType == "credit_card" && input.TransactionStatus == "capture" && input.FraudStatus == "accept" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "settlement" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "deny" || input.TransactionStatus == "expire" || input.TransactionStatus == "cancel" {
		transaction.Status = "cancelled"
	}

	updatedTransaction, err := s.repository.Update(transaction)
	if err != nil {
		return err
	}

	eventDetail, err := s.eventRepository.FindById(updatedTransaction.EventId)
	if err != nil {
		return err
	}

	if updatedTransaction.Status == "paid" {
		eventDetail.Quota = eventDetail.Price - 1

		_, err := s.eventRepository.Update(eventDetail)
		if err != nil {
			return err
		}
	}

	return nil
}
