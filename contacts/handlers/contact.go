package handlers

import (
	"contacts/interfaces"
	"contacts/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Contact struct {
	IContact interfaces.IContact
}

func (c *Contact) Create() func(*gin.Context) {
	return func(ctx *gin.Context) {
		contact := new(models.Contact)
		// buf, err := io.ReadAll(c.Request.Body)
		// json.Unmarshal(buf, contact)
		err := ctx.Bind(contact)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": 400, "innerError": err.Error(), "message": "bad request"})
			return
		}
		if err := contact.Validate(); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": 400, "innerError": err.Error(), "message": err.Error()})
			return
		}
		contact.Status = "created"
		contact.LastModified = strconv.Itoa(int(time.Now().Unix()))
		if con, err := c.IContact.Create(contact); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": 400, "innerError": err.Error(), "message": "Error in creating contact info"})
			return
		} else {
			ctx.JSON(http.StatusCreated, con)
			return
		}
	}
}
