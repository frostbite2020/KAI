package services

import (
	"MsKAI/internal/models"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"log"
	"time"

	"gorm.io/gorm"
)

func SaveOrUpdateUser(db *gorm.DB, email string, name string) (*models.User, error) {
	var user models.User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Jika pengguna tidak ditemukan, buat yang baru
			user = models.User{
				Email: email,
				Name:  name,
			}
			if err := db.Create(&user).Error; err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	} else {
		// Perbarui data pengguna jika diperlukan
		user.Name = name
		if err := db.Save(&user).Error; err != nil {
			return nil, err
		}
	}
	return &user, nil
}

func CreateSession(db *gorm.DB, userID uint, duration time.Duration) (*models.Session, error) {
	// Generate random token
	token := generateToken(32)

	session := models.Session{
		UserID:    userID,
		Token:     token,
		ExpiresAt: time.Now().Add(duration),
	}

	if err := db.Create(&session).Error; err != nil {
		return nil, err
	}
	return &session, nil
}

// Function untuk generate token unik
func generateToken(length int) string {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatalf("Failed to generate token: %v", err)
	}
	return base64.URLEncoding.EncodeToString(b)
}
