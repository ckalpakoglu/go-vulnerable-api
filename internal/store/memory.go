package store

import (
	"fmt"
	"sync"
	"time"

	"github.com/cenk-kalpakoglu/go-vulnerable-api/internal/models"
	"github.com/google/uuid"
)

type MemoryStore struct {
	mu    sync.RWMutex
	users map[string]models.User // key: email
	otp   []models.OTP
}

func NewMemoryStore() (*MemoryStore, error) {
	ms := &MemoryStore{users: make(map[string]models.User)}
	if err := ms.seedUsers(); err != nil {
		return nil, err
	}
	return ms, nil
}

func (s *MemoryStore) seedUsers() error {
	users := []models.User{
		{ID: uuid.New().String(), Email: "admin", Password: "admin123", Mobile: "15551230000"},
		{ID: uuid.New().String(), Email: "user1@example.com", Password: "11223344", Mobile: "15551230001"},
		{ID: uuid.New().String(), Email: "user2@example.com", Password: "11223355", Mobile: "15551230002"},
		{ID: uuid.New().String(), Email: "user3@example.com", Password: "11223366", Mobile: "15551230003"},
	}

	fmt.Println(users)

	s.mu.Lock()
	defer s.mu.Unlock()
	for _, u := range users {
		s.users[u.Email] = u
	}
	return nil
}

func (s *MemoryStore) Authenticate(email, password string) (models.User, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	u, ok := s.users[email]
	if !ok {
		return models.User{}, false
	}
	if u.Password != password {
		return models.User{}, false
	}
	return u, true
}

func (s *MemoryStore) ListUsers() []models.User {
	s.mu.RLock()
	defer s.mu.RUnlock()
	result := make([]models.User, 0, len(s.users))
	for _, u := range s.users {
		result = append(result, u)
	}
	return result
}

func (s *MemoryStore) GetUserByID(id string) (models.User, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, u := range s.users {
		if u.ID == id {
			return u, true
		}
	}
	return models.User{}, false
}

func (s *MemoryStore) GetUserByMobile(mobile string) (models.User, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, u := range s.users {
		if u.Mobile == mobile {
			return u, true
		}
	}
	return models.User{}, false
}

func (s *MemoryStore) CheckUserAndMobile(email, mobile string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, u := range s.users {
		if u.Email == email && u.Mobile == mobile {
			return true
		}
	}
	return false
}

func (s *MemoryStore) AddNewOTP(token string) (models.OTP, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	otp := models.OTP{
		Value:     token,
		Expired:   false,
		CreatedAt: time.Now(),
	}

	s.otp = append(s.otp, otp)

	return otp, true
}
