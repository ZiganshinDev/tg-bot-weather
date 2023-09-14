package tgbot

import "sync"

type User struct {
	mu     sync.Mutex
	states map[int64]int
}

func NewUserState() *User {
	return &User{
		states: make(map[int64]int),
	}
}

func (us *User) SetState(userId int64, state int) {
	us.mu.Lock()
	defer us.mu.Unlock()
	us.states[userId] = state
}

func (us *User) GetState(userId int64) (int, bool) {
	us.mu.Lock()
	defer us.mu.Unlock()
	state, ok := us.states[userId]
	return state, ok
}
