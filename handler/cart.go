package handler

import (
	"Cart/database"
	"Cart/models"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// validState is a custom validator function that checks if the cart state is either "PENDING" or "COMPLETED"
func validState(fl validator.FieldLevel) bool {
	state := fl.Field().String()
	return state == "PENDING" || state == "COMPLETED"
}

func CreateCart(c echo.Context) error {
	userID := GetUserIDFromContext(c)

	cartInput := new(models.CartInput)
	if err := c.Bind(cartInput); err != nil {
		return err
	}

	cart := models.Cart{
		UserID: userID,
		Data:   cartInput.Data,
		State:  cartInput.State,
	}

	if err := validate.Struct(&cart); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	result := database.DB.Create(&cart)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, result.Error)
	}

	return c.JSON(http.StatusCreated, cart)
}

func UpdateCart(c echo.Context) error {
	userID := GetUserIDFromContext(c)

	cartID, _ := strconv.Atoi(c.Param("id"))

	var existingCart models.Cart
	if err := database.DB.Where("user_id = ? AND id = ?", userID, cartID).First(&existingCart).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Cart not found")
	}

	if existingCart.State == "COMPLETED" {
		return c.JSON(http.StatusForbidden, "Cannot update a completed cart")
	}

	cartInput := new(models.CartInput)
	if err := c.Bind(cartInput); err != nil {
		return err
	}

	existingCart.Data = cartInput.Data
	existingCart.State = cartInput.State

	if err := validate.Struct(&existingCart); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	database.DB.Save(&existingCart)
	return c.JSON(http.StatusOK, existingCart)
}

func GetAllCarts(c echo.Context) error {
	userID := GetUserIDFromContext(c)

	var carts []models.Cart
	result := database.DB.Where("user_id = ?", userID).Find(&carts)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, result.Error)
	}

	return c.JSON(http.StatusOK, carts)
}

func GetCart(c echo.Context) error {
	userID := GetUserIDFromContext(c)

	cartID, _ := strconv.Atoi(c.Param("id"))

	var cart models.Cart
	result := database.DB.Where("user_id = ? AND id = ?", userID, cartID).First(&cart)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, "Cart not found")
	}

	return c.JSON(http.StatusOK, cart)
}

func DeleteCart(c echo.Context) error {
	userID := GetUserIDFromContext(c)

	cartID, _ := strconv.Atoi(c.Param("id"))

	var cart models.Cart
	if err := database.DB.Where("user_id = ? AND id = ?", userID, cartID).Delete(&cart).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Cart not found")
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Cart deleted"})
}
	
