package handlers

import (
    "WISP/internal/core/domain"
    "fmt"
    "github.com/gin-gonic/gin"
)

func (h *Handler) SendTestEmails(c *gin.Context) {
    messages := make([]domain.EmailMessage, 100)
    for i := 0; i < 100; i++ {
        messages[i] = domain.EmailMessage{
            From:    "test@example.com",
            To:      fmt.Sprintf("user%d@example.com", i),
            Subject: fmt.Sprintf("Test Email %d", i),
            Body:    fmt.Sprintf("This is test email number %d", i),
        }
    }

    if err := h.Producer.PublishEmails(messages); err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, gin.H{"message": "100 emails queued successfully"})
}