package api

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	db "github.com/STAMBOULI-ABDELKARIM/car_repair_shop/db/sqlc"
	"github.com/gin-gonic/gin"
)

// swagger:model CustomerResponse
type CustomerResponse struct {
	// The ID of a Customer
	// example: 1 2 3 4 5
	ID int64 `json:"id"`
	// The Name of a Customer
	// example: Karim Stam
	FullName string `json:"full_name"`
	// The PhoneNumber of a Customer
	// example: +2131122334455
	PhoneNumber string `json:"phone_number"`
	// The time a Customer was created
	// example: 2021-05-25T00:53:16.535668Z
	CreatedAt time.Time `json:"created_at"`
}

// swagger:model createCustomerRequest
type createCustomerRequest struct {
	// The Name of a Customer
	// example: Karim Stam
	FullName string `json:"fullName" binding:"required"`
	// The PhoneNumber for a Customer
	// example: +2131122334455
	PhoneNumber string `json:"phoneNumber" binding:"required"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

// createCustomer godoc
// @Summary Create new Customer
// @Description Create a new Customer
// @ID create-Customer
// @Tags Customer
// @Accept  json
// @Produce  json
// @Param Body body createCustomerRequest true "The body to create a Customer"
// @Success 200 {object} CustomerResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /customers [post]
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

// swagger:model getCustomerRequest
type getCustomerRequest struct {
	// The id of a thing
	// in:path
	ID int64 `uri:"id" binding:"required"`
}

// getCustomer godoc
// @Summary  GET Customer
// @Description  GET  Customer by it's id
// @Tags Customer
// @ID get-Customer
// @Accept  json
// @Produce  json
// @Param id path string true  "The id to get a Customer"
// @Success 200 {object} CustomerResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /customers/{id} [get]
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

type CustomersResponse struct {
	Customers []CustomerResponse `json:"customers"`
} // @name CustomersResponse

// swagger:model ListCustomersRequest
type ListCustomersRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

// listCustomers godoc
// @Summary list all Customers
// @Description Create GET list of all Customers
// @Tags Customer
// @ID list-Customer
// @Accept  json
// @Produce  json
// @Param Body body ListCustomersRequest true "The body to list all Customers by pagination"
// @Success 200 {object} CustomersResponse
// @Success 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /customers [get]
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

// deleteCustomer godoc
// @Summary DELETE a Customer
// @Description use this api to delete a customer by it's id
// @Tags Customer
// @ID delete-Customer
// @Accept  json
// @Produce  json
// @Param id path string true  "The id to delete a Customer"
// @Success 204 string deleted
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /customers/{id} [delete]
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

// swagger:model updateCustomerRequest
type updateCustomerRequest struct {
	// The Name of a Customer
	// example: Karim Stam
	FullName string `json:"fullName"`
	// The PhoneNumber for a Customer
	// example: +2131122334455
	PhoneNumber string `json:"phoneNumber"`
}

// updateCustomer godoc
// @Summary update  Customer
// @Description update a  Customer
// @Tags Customer
// @ID update-Customer
// @Accept  json
// @Produce  json
// @Param id path string true  "The id to get a Customer"
// @Param Body body createCustomerRequest true "The body to create a Customer"
// @Success 200 {object} CustomersResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /customers/{id} [put]
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
