package api

import (
	"database/sql"
	"net/http"
	"strconv"

	db "github.com/STAMBOULI-ABDELKARIM/car_repair_shop/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createCustomerRequest struct {
	FullName    string `json:"fullName" binding:"required"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
}
type Response struct {
	Message string `json:"message"`
}

// Paths Information

// createCustomer godoc
// @Summary Create new Customer
// @Description Create new Customer
// @Tags customer,create
// @Accept  json
// @Produce  json
// @Param FullName formData string true "FullName"
// @Param PhoneNumber formData string true "PhoneNumber"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 401 {object} Response
// @Router /customers [get]
func (server *Server) createCustomer(ctx *gin.Context) {
	var req createCustomerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateCustomerParams{
		FullName:    req.FullName,
		PhoneNumber: req.PhoneNumber,
	}
	customer, err := server.store.CreateCustomer(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, customer)
}

type getCustomerRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

func (server *Server) getCustomer(ctx *gin.Context) {
	var req getCustomerRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	customer, err := server.store.GetCustomer(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, customer)

}

type ListCustomersRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listCustomers(ctx *gin.Context) {
	var req ListCustomersRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.ListCustomersParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}
	customers, err := server.store.ListCustomers(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, customers)

}

type deleteCustomerRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

func (server *Server) deleteCustomer(ctx *gin.Context) {
	var req deleteCustomerRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteCustomer(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusNoContent, "deleted")

}

type updateCustomerRequest struct {
	FullName    string `json:"fullName"`
	PhoneNumber string `json:"phoneNumber"`
}

func (server *Server) updateCustomer(ctx *gin.Context) {
	var req updateCustomerRequest
	userID := ctx.Param("id")
	user, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	customer, err := server.store.GetCustomer(ctx, user)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateCustomerParams{
		ID:          customer.ID,
		FullName:    req.FullName,
		PhoneNumber: req.PhoneNumber,
	}

	customer2, err := server.store.UpdateCustomer(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, customer2)

}
