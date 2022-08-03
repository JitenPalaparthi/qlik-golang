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

func (c *Contact) GetBy() func(*gin.Context) {
	return func(ctx *gin.Context) {
		id, ok := ctx.Params.Get("id")
		if !ok {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": 400, "innerError": "invalid id parameter", "message": "invalid id parameter"})
			return
		}

		contact, err := c.IContact.GetBy(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": 400, "innerError": err.Error(), "message": "bad request"})
			return
		}
		ctx.JSON(http.StatusOK, contact)
	}
}

func (c *Contact) DeleteBy() func(*gin.Context) {
	return func(ctx *gin.Context) {
		id, ok := ctx.Params.Get("id")
		if !ok {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": 400, "innerError": "invalid id parameter", "message": "invalid id parameter"})
			return
		}

		noOfRecords, err := c.IContact.DeleteBy(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": 400, "innerError": err.Error(), "message": "bad request"})
			return
		}
		ctx.JSON(http.StatusAccepted, noOfRecords)
	}
}

func (c *Contact) UpdateBy() func(*gin.Context) {
	return func(ctx *gin.Context) {
		id, ok := ctx.Params.Get("id")
		if !ok {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": 400, "innerError": "invalid id parameter", "message": "invalid id parameter"})
			return
		}
		data := make(map[string]interface{})
		err := ctx.Bind(&data)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": 400, "innerError": err.Error(), "message": "bad request"})
			return
		}

		contact, err := c.IContact.UpdateBy(id, data)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": 400, "innerError": err.Error(), "message": "bad request"})
			return
		}
		ctx.JSON(http.StatusOK, contact)
	}
}
