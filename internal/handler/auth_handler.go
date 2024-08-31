package handler

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nikitsingh/forky/backend/internal/service"
)

type AuthHandler struct {
	service     *service.AuthService
	userService *service.UserService
}

func NewAuthHandler(service *service.AuthService, userService *service.UserService) *AuthHandler {
	return &AuthHandler{service: service, userService: userService}
}

func (h *AuthHandler) CreateMagicLink(c *gin.Context) {
	var req struct {
		Email string `json:"email" binding:"required,email"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		return
	}

	magicLink, err := h.service.CreateMagicLink(c.Request.Context(), req.Email)
	if err != nil {
		log.Println("Failed to create magic link", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create magic link"})
		return
	}

	// TODO: Send the magic link to the user's email and remove from here
	c.JSON(http.StatusOK, gin.H{
		"message": "OTP sent to " + req.Email,
		"email":   magicLink.Email,
		"otp":     magicLink.OTP,
	})
}

func (h *AuthHandler) VerifyMagicLink(c *gin.Context) {
	var req struct {
		Email string `json:"email" binding:"required,email"`
		OTP   string `json:"otp" binding:"required,len=6"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := h.service.VerifyMagicLink(c.Request.Context(), req.Email, req.OTP)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired OTP"})
		return
	}

	user, err := h.userService.GetUserByEmail(c.Request.Context(), req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			user, err = h.userService.CreateUser(c.Request.Context(), req.Email)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
				return
			}
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
			return
		}
	}

	session, err := h.service.CreateSession(c.Request.Context(), user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create session"})
		return
	}

	c.SetCookie("session_token", session.Token.String(), 0, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "OTP verified successfully for " + user.Email})
}
