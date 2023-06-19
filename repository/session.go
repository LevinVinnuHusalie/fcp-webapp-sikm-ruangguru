package repository

import (
	"a21hc3NpZ25tZW50/model"
	"errors"
	"time"

	"gorm.io/gorm"
)

type SessionRepository interface {
	AddSessions(session model.Session) error
	DeleteSession(token string) error
	UpdateSessions(session model.Session) error
	SessionAvailEmail(email string) (model.Session, error)
	SessionAvailToken(token string) (model.Session, error)
	TokenExpired(session model.Session) bool
}

type sessionsRepo struct {
	db *gorm.DB
}

func NewSessionsRepo(db *gorm.DB) *sessionsRepo {
	return &sessionsRepo{db}
}

func (u *sessionsRepo) AddSessions(session model.Session) error {
	if err := u.db.Create(&session).Error; err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (u *sessionsRepo) DeleteSession(token string) error {
	var sss model.Session
	result := u.db.Where("token = ?", token).Delete(&sss)
	if result.RowsAffected == 0 {
		return errors.New("no product updated")
	}
	if result.Error != nil {
		return result.Error
	}
	return nil // TODO: replace this
}

func (u *sessionsRepo) UpdateSessions(session model.Session) error {
	result := u.db.Where("email = ?", session.Email).Updates(&session)
	if result.RowsAffected == 0 {
		return result.Error
	}
	if result.Error != nil {
		return result.Error
	}
	return nil // TODO: replace this
}

func (u *sessionsRepo) SessionAvailEmail(email string) (model.Session, error) {
	var sss model.Session
	result := u.db.Where("email = ?", email).First(&sss)
	if result.Error != nil {
		return model.Session{}, result.Error
	}
	return sss, nil // TODO: replace this
}

func (u *sessionsRepo) SessionAvailToken(token string) (model.Session, error) {
	var sss model.Session
	result := u.db.Where("token = ?", token).First(&sss)
	if result.Error != nil {
		return model.Session{}, result.Error
	}
	return sss, nil
	// TODO: replace this
}

func (u *sessionsRepo) TokenValidity(token string) (model.Session, error) {
	session, err := u.SessionAvailToken(token)
	if err != nil {
		return model.Session{}, err
	}

	if u.TokenExpired(session) {
		err := u.DeleteSession(token)
		if err != nil {
			return model.Session{}, err
		}
		return model.Session{}, err
	}

	return session, nil
}

func (u *sessionsRepo) TokenExpired(session model.Session) bool {
	return session.Expiry.Before(time.Now())
}
