package app

import (
	"fmt"

	"github.com/keksShmeks1/notRate/pkg/client"
)

type Server struct {
	Amount int64
}

func InitServer(amount int64) *Server {
	return &Server{Amount: amount}
}

func (s *Server) GetRate() error {
	rate, err := client.GetNotRate(s.Amount)
	if err != nil {
		return err
	}

	fmt.Printf("Not amount %d in RUB is: %.3f \n", s.Amount, rate)
	return nil
}
