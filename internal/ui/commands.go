package ui

import (
	"context"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
	"github.com/zigzagalex/gator/internal/database"
)

type errorMsg struct{ Error string }

type usersFetchedMsg struct{ Users []database.User }

func fetchUsersCmd(q *database.Queries) tea.Cmd {
	return func() tea.Msg {
		users, err := q.GetUsers(context.TODO())
		if err != nil {
			return errorMsg{err.Error()}
		}
		return usersFetchedMsg{users}
	}
}

type feedsFetchedMsg struct {
	Feeds []database.GetFeedFollowsForUserRow
}

func fetchFollowedFeedsCmd(q *database.Queries, userName string) tea.Cmd {
	return func() tea.Msg {
		feedFollows, err := q.GetFeedFollowsForUser(context.TODO(), userName)
		if err != nil {
			return errorMsg{err.Error()}
		}
		return feedsFetchedMsg{feedFollows}
	}
}

type postsFetchedMsg struct{ Posts []database.Post }

func fetchPostsCmd(q *database.Queries, userId uuid.UUID, feedId uuid.UUID) tea.Cmd {
	return func() tea.Msg {
		posts, err := q.GetFeedPosts(context.TODO(), database.GetFeedPostsParams{
			UserID: userId,
			FeedID: feedId,
		})
		if err != nil {
			return errorMsg{err.Error()}
		}
		return postsFetchedMsg{posts}
	}
}

type OpenedPostMsg struct{ Error error }

func postOpenedPostCmd(q *database.Queries, userId uuid.UUID, feedId uuid.UUID, postId uuid.UUID) tea.Cmd {
	return func() tea.Msg {
		_, err := q.CreateOpenedPost(context.TODO(), database.CreateOpenedPostParams{
			ID:       uuid.New(),
			OpenedAt: time.Now(),
			UserID:   userId,
			FeedID:   feedId,
			PostID:   postId,
		})
		if err != nil {
			return errorMsg{err.Error()}
		}
		return OpenedPostMsg{nil}
	}
}

type CreateUserMsg struct {Error error}

func createUsersCmd(q *database.Queries, userName string) tea.Cmd {
	return func () tea.Msg {
		_, err := q.CreateUser(context.TODO(), database.CreateUserParams{
			ID: uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name: userName,
		})
		if err != nil {
			return errorMsg{err.Error()}
		}
		return CreateUserMsg{nil}
	}
}