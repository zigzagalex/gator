package commands

import (
	"context"
	"fmt"
	"os"
)

func HandlerReset(s *State, cmd Command) error {
	err := s.DB.Reset(context.Background())
	if err != nil {
		fmt.Println("DB reset failed")
		os.Exit(1)
	}
	fmt.Println("DB reset successful")
	return nil

}
