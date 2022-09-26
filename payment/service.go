package payment

import (
	"ngevent-api/event"
	"ngevent-api/user"

	midtrans "github.com/veritrans/go-midtrans"
)

type service struct {
	eventRepository event.Repository
}

type Service interface {
	GetPaymentURL(transaction Transaction, user user.User) (string, error)
}

func NewService(eventRepository event.Repository) *service {
	return &service{eventRepository}
}

func (s *service) GetPaymentURL(transaction Transaction, user user.User) (string, error) {
	midclient := midtrans.NewClient()
	midclient.ServerKey = ""
	midclient.ClientKey = ""
	midclient.APIEnvType = midtrans.Sandbox

	snapGateway := midtrans.SnapGateway{
		Client: midclient,
	}

	snapReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			Email: user.Email,
			FName: user.Name,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  transaction.OrderId,
			GrossAmt: int64(transaction.Amount),
		},
	}

	snapTokenresp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "", err
	}

	return snapTokenresp.RedirectURL, nil
}
