package setup

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/zigzagalex/gator/internal/database"
)

// EnsureInitialUser prompts once and inserts a user if the table is empty.
// It returns the created user (or the first existing one) and nil on success.
func EnsureInitialUser(q *database.Queries) (database.User, error) {
	ctx := context.TODO()

	users, err := q.GetUsers(ctx)
	if err != nil {
		return database.User{}, fmt.Errorf("check users: %w", err)
	}
	if len(users) > 0 {
		return users[0], nil // already have at least one
	}

	// prompt on the normal terminal
	fmt.Print("It looks like this is your first run.\nEnter a user name: ")
	reader := bufio.NewReader(os.Stdin)
	nameRaw, _ := reader.ReadString('\n')
	name := strings.TrimSpace(nameRaw)
	if name == "" {
		name = "gator" // sensible default
	}

	user, err := q.CreateUser(ctx, database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	})
	if err != nil {
		return database.User{}, fmt.Errorf("insert user: %w", err)
	}
	return user, nil
}
