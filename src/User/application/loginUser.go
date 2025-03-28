package application

import (
	"api/src/User/domain"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// LoginUserService maneja la autenticación de usuarios
type LoginUserService struct {
	repo domain.UserRepository
}

// NewLoginUserService crea una nueva instancia del servicio de login
func NewLoginUserService(repo domain.UserRepository) *LoginUserService {
	return &LoginUserService{repo: repo}
}

// LoginUser autentica un usuario y genera un token JWT si es válido
func (s *LoginUserService) LoginUser(correo, contraseña string) (string, string, error) {
	// Buscar usuario por correo y contraseña (el repositorio debe implementar esta firma)
	user, err := s.repo.LoginUser(correo, contraseña)
	if err != nil {
		return "", "", fmt.Errorf("error al buscar usuario: %v", err)
	}
	if user == nil {
		return "", "", errors.New("usuario no encontrado")
	}

	// Verificar contraseña con bcrypt
	err = bcrypt.CompareHashAndPassword([]byte(user.Contraseña), []byte(contraseña))
	if err != nil {
		return "", "", errors.New("contraseña incorrecta")
	}

	// Generar token JWT
	token, err := generateJWT(user)
	if err != nil {
		return "", "", fmt.Errorf("error al generar token: %v", err)
	}

	// Retornar el token y el ID del usuario
	return token, user.ID.Hex(), nil
}

// generateJWT genera un token JWT con la información del usuario
func generateJWT(user *domain.User) (string, error) {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		return "", errors.New("JWT_SECRET no está configurado")
	}

	// Crear los claims del token
	claims := jwt.MapClaims{
		"user_id": user.ID.Hex(),
		"correo":  user.Correo,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Expira en 24 horas
	}

	// Crear el token con claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}
