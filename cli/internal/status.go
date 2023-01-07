package internal

type CommentStatus string

const (
	OPEN CommentStatus = "OPEN"
)

func (status CommentStatus) string() string {
	return string(status)
}
