package notes

import (
	"context"
	"errors"
	"sync"
	"time"
)

// Note is a sample record. It exists so the API, database, and clients
// share one vertical slice. Replace it when the product has a real resource.
type Note struct {
	ID        string    `json:"id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

var ErrUnavailable = errors.New("store unavailable")

// Store is the boundary a future service split would keep.
// HTTP handlers depend on this interface, not on Postgres.
type Store interface {
	Ping(ctx context.Context) error
	List(ctx context.Context) ([]Note, error)
	Create(ctx context.Context, body string) (Note, error)
}

// Memory is an in-process store for tests and for running the API with no database.
type Memory struct {
	mu    sync.Mutex
	notes []Note
	now   func() time.Time
	ids   func() string
}

func NewMemory() *Memory {
	seq := 0
	return &Memory{
		now: time.Now,
		ids: func() string {
			seq++
			return time.Now().UTC().Format("20060102150405") + "-" + itoa(seq)
		},
	}
}

func (m *Memory) Ping(context.Context) error { return nil }

func (m *Memory) List(context.Context) ([]Note, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	out := make([]Note, len(m.notes))
	copy(out, m.notes)
	return out, nil
}

func (m *Memory) Create(_ context.Context, body string) (Note, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	note := Note{ID: m.ids(), Body: body, CreatedAt: m.now().UTC()}
	m.notes = append([]Note{note}, m.notes...)
	return note, nil
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	var buf [12]byte
	i := len(buf)
	for n > 0 {
		i--
		buf[i] = byte('0' + n%10)
		n /= 10
	}
	return string(buf[i:])
}
