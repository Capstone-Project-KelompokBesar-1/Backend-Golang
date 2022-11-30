package users

import (
	"net/http"
	"ourgym/controllers"
	"ourgym/dto"
	"ourgym/models"
	"ourgym/services"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{
		userService,
	}
}

func (uc *UserController) GetAll(c echo.Context) error {
	name := c.QueryParam("name")

	usersData := uc.userService.GetAll(name)

	users := []dto.DTOUser{}

	for _, user := range usersData {
		users = append(users, user.ConvertToDTO())
	}

	return c.JSON(http.StatusOK, controllers.Response(http.StatusOK, "Success Get Users By Name", users))
}

func (uc *UserController) GetOneByFilter(c echo.Context) error {
	var id string = c.Param("id")

	user := uc.userService.GetByID(id)

	if user.ID == 0 {
		return c.JSON(http.StatusNotFound, controllers.Response(http.StatusNotFound, "User Not Found", ""))
	}

	return c.JSON(http.StatusOK, controllers.Response(http.StatusOK, "User Found", user.ConvertToDTO()))
}

func (uc *UserController) Create(c echo.Context) error {
	input := models.User{}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, controllers.Response(http.StatusBadRequest, "Failed", ""))
	}

	if err := input.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, controllers.Response(http.StatusBadRequest, "Request invalid", ""))
	}

	user := uc.userService.Create(input)

	return c.JSON(http.StatusOK, controllers.Response(http.StatusOK, "Success Created User", user))
}

func (uc *UserController) Update(c echo.Context) error {
	input := models.User{}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusNotFound, controllers.Response(http.StatusBadRequest, "Failed", ""))
	}

	if err := input.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, controllers.Response(http.StatusBadRequest, "Request invalid", ""))
	}

	var userId string = c.Param("id")

	user := uc.userService.Update(userId, input)

	return c.JSON(http.StatusOK, controllers.Response(http.StatusOK, "Success Update User", user.ConvertToDTO()))
}

func (uc *UserController) Delete(c echo.Context) error {
	var userId string = c.Param("id")

	isSuccess := uc.userService.Delete(userId)

	if !isSuccess {
		return c.JSON(http.StatusNotFound, controllers.Response(http.StatusNotFound, "User Not Found", ""))
	}

	return c.JSON(http.StatusOK, controllers.Response(http.StatusOK, "User Success Deleted", ""))
}

func (uc *UserController) DeleteMany(c echo.Context) error {
	ids := c.QueryParam("ids")

	isSuccess := uc.userService.DeleteMany(ids)

	if !isSuccess {
		return c.JSON(http.StatusNotFound, controllers.Response(http.StatusNotFound, "Users Not Found", ""))
	}

	return c.JSON(http.StatusOK, controllers.Response(http.StatusOK, "Users Success Deleted", ""))
}
