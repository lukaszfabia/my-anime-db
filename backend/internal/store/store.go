package store

import (
	"errors"
	"sync"
	"time"
)

type VerificationStore struct {
	sync.Mutex
	verificationCodes map[string]string
}

func NewVerificationStore() *VerificationStore {
	return &VerificationStore{
		verificationCodes: make(map[string]string),
	}
}

func (vs *VerificationStore) Set(email string, pin string) {
	vs.Lock()
	defer vs.Unlock()
	vs.verificationCodes[email] = pin
	go vs.setExpiration(email, time.Minute*2) // 2 minutes
}

func (vs *VerificationStore) Compare(inputPin string, email string) error {
	vs.Lock()
	defer vs.Unlock()
	pin, exists := vs.verificationCodes[email]

	if pin != inputPin || !exists {
		return errors.New("invalid pin")
	}

	return nil
}

func (vs *VerificationStore) Delete(email string) {
	vs.Lock()
	defer vs.Unlock()
	delete(vs.verificationCodes, email)
}

func (vs *VerificationStore) setExpiration(email string, duration time.Duration) {
	time.Sleep(duration)
	vs.Delete(email)
}
