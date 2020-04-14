package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/muhammadisa/go-service-boilerplate/api/app/foobar"
	"github.com/muhammadisa/go-service-boilerplate/api/models"
	"github.com/muhammadisa/go-service-boilerplate/api/response"
	"github.com/muhammadisa/go-service-boilerplate/api/utils/message"
	"github.com/muhammadisa/go-service-boilerplate/api/utils/paging"
	uuid "github.com/satori/go.uuid"
)

// FoobarHandler struct
type FoobarHandler struct {
	fBUsecase foobar.Usecase
}

// NewFoobarHandler initialize enpoints
func NewFoobarHandler(e *echo.Group, fBu foobar.Usecase) {
	handler := &FoobarHandler{
		fBUsecase: fBu,
	}
	e.GET("/foobars/", handler.Fetch)
	e.GET("/foobar/:id", handler.GetByID)
	e.POST("/foobar/", handler.Store)
	e.PATCH("/foobar/update/:id", handler.Update)
	e.DELETE("/foobar/delete/:id", handler.Delete)
}

var (
	model = models.Foobar{}
)

// Fetch foobar data
func (fB *FoobarHandler) Fetch(c echo.Context) error {
	var err error

	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	db, fBars, err := fB.fBUsecase.Fetch()
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.Response{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	return c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    message.GenerateMessage(uuid.Nil, "GET", model, true),
		Data:       paging.GetPaginator(db, page, limit, fBars),
	})
}

// GetByID foobar data
func (fB *FoobarHandler) GetByID(c echo.Context) error {
	var err error

	uid, err := uuid.FromString(c.Param("id"))
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
	}
	fBar, err := fB.fBUsecase.GetByID(uid)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.Response{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	return c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    message.GenerateMessage(uid, "GET", model, true),
		Data:       fBar,
	})
}

// Store foobar data
func (fB *FoobarHandler) Store(c echo.Context) error {
	var err error
	var fooBar models.Foobar

	err = c.Bind(&fooBar)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	err = c.Validate(fooBar)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	err = fB.fBUsecase.Store(&fooBar)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	return c.JSON(http.StatusCreated, response.Response{
		StatusCode: http.StatusCreated,
		Message:    message.GenerateMessage(fooBar.ID, "POST", model, true),
		Data:       fooBar,
	})
}

// Update foobar data
func (fB *FoobarHandler) Update(c echo.Context) error {
	var err error
	var fooBar models.Foobar

	err = c.Bind(&fooBar)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	err = c.Validate(fooBar)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    message.GenerateMessage(uuid.Nil, "POST", model, false),
			Data:       err,
		})
	}
	_, err = fB.fBUsecase.GetByID(fooBar.ID)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.Response{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	err = fB.fBUsecase.Update(&fooBar)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	return c.JSON(http.StatusCreated, response.Response{
		StatusCode: http.StatusCreated,
		Message:    message.GenerateMessage(fooBar.ID, "PATCH", model, true),
		Data:       fooBar,
	})
}

// Delete foobar data
func (fB *FoobarHandler) Delete(c echo.Context) error {
	var err error

	uid, err := uuid.FromString(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	fBar, err := fB.fBUsecase.GetByID(uid)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.Response{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	err = fB.fBUsecase.Delete(fBar.ID)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.Response{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	return c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    message.GenerateMessage(uid, "DELETE", model, true),
		Data:       fBar.ID,
	})
}
