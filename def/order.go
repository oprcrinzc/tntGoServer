package def

const (
	StatusPending   = "pending"
	StatusConfirmed = "confirmed"
	StatusCompleted = "completed"
)

type Order struct {
	Time     int
	Customer string
	Content  string
	File     []string
	Color    string
	Material string
	Status   string
}

type Orders []Order
