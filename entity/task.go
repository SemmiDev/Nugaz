package entity

type (
	EisenHowerMatrix string
	Epoch            int64
)

const (
	UrgentAndImportant       EisenHowerMatrix = "Urgent And Important"         // quadrant 1 (do it now)
	ImportantButNotUrgent    EisenHowerMatrix = "Important But Not Urgent"     // quadrant 2 (plan it)
	UrgentButNotImportant    EisenHowerMatrix = "Urgent But Not Important"     // quadrant 3 (delegate)
	NotUrgentAndNotImportant EisenHowerMatrix = "Not Urgent And Not Important" // quadrant 4 (drop it)
)

type Task struct {
	ID          string           `db:"id"`
	Title       string           `db:"title"`
	Description string           `db:"description"`
	IsDone      bool             `db:"is_done"`
	IsOver      bool             `db:"is_over"`
	Duration    int64            `db:"duration"`
	Matrix      EisenHowerMatrix `db:"matrix"`
	StartAt     Epoch            `db:"start_at"`
	Due         Epoch            `db:"due"`
	CreatedAt   Epoch            `db:"created_at"`
	UpdatedAt   Epoch            `db:"updated_at"`
}
