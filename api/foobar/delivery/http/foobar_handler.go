package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/muhammadisa/restful-api-boilerplate/api/foobar"
	"github.com/muhammadisa/restful-api-boilerplate/api/models"
	"github.com/muhammadisa/restful-api-boilerplate/api/response"
	"github.com/muhammadisa/restful-api-boilerplate/api/utils/message"
)

// FoobarHandler struct
type FoobarHandler struct {
	fBUsecase foobar.Usecase
}

// NewFoobarHandler initialize enpoints
func NewFoobarHandler(e *echo.Echo, fBu foobar.Usecase) {
	handler := &FoobarHandler{
		fBUsecase: fBu,
	}
	e.GET("foobars/", handler.Fetch)
	e.GET("foobar/:id", handler.GetByID)
	e.POST("foobar/", handler.Store)
	e.PATCH("foobar/update/:id", handler.Update)
	e.DELETE("foobar/delete/:id", handler.Delete)
}

// Fetch foobar data
func (fB *FoobarHandler) Fetch(c echo.Context) error {
	var err error

	fBars, err := fB.fBUsecase.Fetch()
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.Response{
			StatusCode: http.StatusUnprocessableEntity,
			Message: message.Message{
				IsSuccess:       false,
				HTTPMethod:      "GET",
				TargetModelName: "foobar",
				WithID:          0,
			}.GenerateMessage(),
			Data: nil,
		})
	}

	return c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message: message.Message{
			IsSuccess:       true,
			HTTPMethod:      "GET",
			TargetModelName: "foobar",
			WithID:          0,
		}.GenerateMessage(),
		Data: fBars,
	})
}

// GetByID foobar data
func (fB *FoobarHandler) GetByID(c echo.Context) error {
	var err error

	idFBar, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadGateway, response.Response{
			StatusCode: http.StatusBadGateway,
			Message: message.Message{
				IsSuccess:       false,
				HTTPMethod:      "GET",
				TargetModelName: "foobar",
				WithID:          0,
			}.GenerateMessage(),
			Data: nil,
		})
	}

	id := uint64(idFBar)
	fBar, err := fB.fBUsecase.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.Response{
			StatusCode: http.StatusNotFound,
			Message: message.Message{
				IsSuccess:       false,
				HTTPMethod:      "GET",
				TargetModelName: "foobar",
				WithID:          0,
			}.GenerateMessage(),
			Data: nil,
		})
	}

	return c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message: message.Message{
			IsSuccess:       true,
			HTTPMethod:      "GET",
			TargetModelName: "foobar",
			WithID:          id,
		}.GenerateMessage(),
		Data: fBar,
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
			Message: message.Message{
				IsSuccess:       false,
				HTTPMethod:      "POST",
				TargetModelName: "foobar",
				WithID:          0,
			}.GenerateMessage(),
			Data: nil,
		})
	}

	err = fB.fBUsecase.Store(&fooBar)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message: message.Message{
				IsSuccess:       false,
				HTTPMethod:      "POST",
				TargetModelName: "foobar",
				WithID:          0,
			}.GenerateMessage(),
			Data: nil,
		})
	}

	return c.JSON(http.StatusCreated, response.Response{
		StatusCode: http.StatusCreated,
		Message: message.Message{
			IsSuccess:       true,
			HTTPMethod:      "POST",
			TargetModelName: "foobar",
			WithID:          0,
		}.GenerateMessage(),
		Data: fooBar,
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
			Message: message.Message{
				IsSuccess:       false,
				HTTPMethod:      "PATCH",
				TargetModelName: "foobar",
				WithID:          0,
			}.GenerateMessage(),
			Data: nil,
		})
	}

	_, err = fB.fBUsecase.GetByID(fooBar.ID)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.Response{
			StatusCode: http.StatusNotFound,
			Message: message.Message{
				IsSuccess:       false,
				HTTPMethod:      "PATCH",
				TargetModelName: "foobar",
				WithID:          0,
			}.GenerateMessage(),
			Data: nil,
		})
	}

	err = fB.fBUsecase.Update(&fooBar)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message: message.Message{
				IsSuccess:       false,
				HTTPMethod:      "PATCH",
				TargetModelName: "foobar",
				WithID:          0,
			}.GenerateMessage(),
			Data: nil,
		})
	}

	return c.JSON(http.StatusCreated, response.Response{
		StatusCode: http.StatusCreated,
		Message: message.Message{
			IsSuccess:       true,
			HTTPMethod:      "PATCH",
			TargetModelName: "foobar",
			WithID:          0,
		}.GenerateMessage(),
		Data: fooBar,
	})
}

// Delete foobar data
func (fB *FoobarHandler) Delete(c echo.Context) error {
	var err error

	idFBar, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadGateway, response.Response{
			StatusCode: http.StatusBadGateway,
			Message: message.Message{
				IsSuccess:       false,
				HTTPMethod:      "DELETE",
				TargetModelName: "foobar",
				WithID:          0,
			}.GenerateMessage(),
			Data: nil,
		})
	}

	id := uint64(idFBar)
	fBar, err := fB.fBUsecase.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.Response{
			StatusCode: http.StatusNotFound,
			Message: message.Message{
				IsSuccess:       false,
				HTTPMethod:      "DELETE",
				TargetModelName: "foobar",
				WithID:          0,
			}.GenerateMessage(),
			Data: nil,
		})
	}

	err = fB.fBUsecase.Delete(fBar.ID)
	if err != nil {
		return c.JSON(http.StatusBadGateway, response.Response{
			StatusCode: http.StatusBadGateway,
			Message: message.Message{
				IsSuccess:       false,
				HTTPMethod:      "DELETE",
				TargetModelName: "foobar",
				WithID:          0,
			}.GenerateMessage(),
			Data: nil,
		})
	}

	return c.JSON(http.StatusCreated, response.Response{
		StatusCode: http.StatusCreated,
		Message: message.Message{
			IsSuccess:       true,
			HTTPMethod:      "DELETE",
			TargetModelName: "foobar",
			WithID:          id,
		}.GenerateMessage(),
		Data: fBar.ID,
	})
}
