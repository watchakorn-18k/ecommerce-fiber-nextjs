package gateways

import (
	"go-ecommerce/domain/entities"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h HTTPGateway) GetAllUserData(ctx *fiber.Ctx) error {

	data, err := h.UserService.GetAllUser()
	if err != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(entities.ResponseModel{Message: "cannot get all users data"})
	}
	return ctx.Status(fiber.StatusOK).JSON(entities.ResponseModel{Message: "success", Data: data})
}

func (h HTTPGateway) CreateNewUserAccount(ctx *fiber.Ctx) error {

	var bodyData entities.UserDataFormat
	if err := ctx.BodyParser(&bodyData); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(entities.ResponseMessage{Message: "invalid json body"})
	}

	status := h.UserService.InsertNewAccount(&bodyData)
	log.Println("status after ", status)

	if !status {
		return ctx.Status(fiber.StatusForbidden).JSON(entities.ResponseModel{Message: "cannot insert new user account."})
	}
	return ctx.Status(fiber.StatusOK).JSON(entities.ResponseModel{Message: "success"})
}

func (h HTTPGateway) GetUserById(ctx *fiber.Ctx) error {

	params := ctx.Queries()

	user_id := params["id"]

	data, err := h.UserService.GetUserByID(user_id)

	if err != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(entities.ResponseModel{Message: "cannot get user data"})
	}
	return ctx.Status(fiber.StatusOK).JSON(entities.ResponseModel{Message: "success", Data: data})
}

func (h HTTPGateway) UpdateUser(ctx *fiber.Ctx) error {
	userData := new(entities.UserDataFormat)

	err := ctx.BodyParser(&userData)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(entities.ResponseModel{Message: "invalid json body"})
	}

	params := ctx.Queries()

	if len(params) <= 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(entities.ResponseModel{Message: "user id not fill"})
	}

	user_id := params["id"]

	err = h.UserService.UpdateUser(user_id, userData)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(entities.ResponseModel{Message: "cannot update user data"})
	}

	return ctx.Status(fiber.StatusOK).JSON(entities.ResponseModel{Message: "success", Data: userData})
}

func (h HTTPGateway) DeleteUser(ctx *fiber.Ctx) error {

	params := ctx.Queries()

	if len(params) <= 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(entities.ResponseModel{Message: "user id not fill"})
	}

	user_id := params["id"]

	err := h.UserService.DeleteUser(user_id)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(entities.ResponseModel{Message: "cannot delete user data"})
	}

	return ctx.Status(fiber.StatusOK).JSON(entities.ResponseModel{Message: "success"})
}

func (h HTTPGateway) GetAddressByUserID(ctx *fiber.Ctx) error {

	params := ctx.Queries()

	if len(params) <= 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(entities.ResponseModel{Message: "user id not fill"})
	}

	user_id := params["id"]

	data, err := h.AddressService.GetAddressByUserID(user_id)

	if err != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(entities.ResponseModel{Message: "cannot get address data"})
	}
	return ctx.Status(fiber.StatusOK).JSON(entities.ResponseModel{Message: "success", Data: data})
}

func (h HTTPGateway) AddNewAddress(ctx *fiber.Ctx) error {

	var bodyData entities.AddressData
	if err := ctx.BodyParser(&bodyData); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(entities.ResponseMessage{Message: "invalid json body"})
	}

	params := ctx.Queries()
	if len(params) <= 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(entities.ResponseModel{Message: "user id not fill"})
	}

	user_id := params["id"]

	status := h.AddressService.InsertNewAddress(user_id, &bodyData)
	log.Println("status after ", status)

	if !status {
		return ctx.Status(fiber.StatusForbidden).JSON(entities.ResponseModel{Message: "cannot insert new address account."})
	}
	return ctx.Status(fiber.StatusOK).JSON(entities.ResponseModel{Message: "success"})
}

func (h HTTPGateway) UpdateAddress(ctx *fiber.Ctx) error {
	addressData := new(entities.AddressData)

	err := ctx.BodyParser(&addressData)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(entities.ResponseModel{Message: "invalid json body"})
	}

	params := ctx.Queries()

	if len(params) <= 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(entities.ResponseModel{Message: "address id not fill"})
	}

	user_id := params["id"]
	index := params["index"]

	i, err := strconv.Atoi(index)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(entities.ResponseModel{Message: "invalid address index"})
	}

	err = h.AddressService.UpdateAddress(user_id, i, addressData)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(entities.ResponseModel{Message: "cannot update address data"})
	}

	return ctx.Status(fiber.StatusOK).JSON(entities.ResponseModel{Message: "success", Data: addressData})
}

func (h HTTPGateway) DeleteAddress(ctx *fiber.Ctx) error {

	params := ctx.Queries()

	if len(params) <= 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(entities.ResponseModel{Message: "user id not fill"})
	}

	user_id := params["id"]
	index := params["index"]

	i, err := strconv.Atoi(index)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(entities.ResponseModel{Message: "invalid address index"})
	}

	err = h.AddressService.DeleteAddress(user_id, i)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(entities.ResponseModel{Message: "cannot delete address data"})
	}

	return ctx.Status(fiber.StatusOK).JSON(entities.ResponseModel{Message: "success"})
}

func (h HTTPGateway) GetOrdersByUserID(ctx *fiber.Ctx) error {

	params := ctx.Queries()

	if len(params) <= 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(entities.ResponseModel{Message: "user id not fill"})
	}

	user_id := params["id"]

	data, err := h.CartService.GetOrdersByUserID(user_id)

	if err != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(entities.ResponseModel{Message: "cannot get cart data"})
	}
	return ctx.Status(fiber.StatusOK).JSON(entities.ResponseModel{Message: "success", Data: data})
}

func (h HTTPGateway) AddtoCart(ctx *fiber.Ctx) error {

	var bodyData entities.OrderData

	if err := ctx.BodyParser(&bodyData); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(entities.ResponseMessage{Message: "invalid json body"})
	}

	params := ctx.Queries()

	if len(params) <= 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(entities.ResponseModel{Message: "user id not fill"})
	}

	user_id := params["id"]

	err := h.CartService.AddtoCart(user_id, &bodyData)

	if err != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(entities.ResponseModel{Message: "cannot insert new order account."})
	}
	return ctx.Status(fiber.StatusOK).JSON(entities.ResponseModel{Message: "success"})

}

func (h HTTPGateway) UpdateOrder(ctx *fiber.Ctx) error {
	cartData := new(entities.OrderData)

	err := ctx.BodyParser(&cartData)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(entities.ResponseModel{Message: "invalid json body"})
	}

	params := ctx.Queries()

	if len(params) <= 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(entities.ResponseModel{Message: "cart id not fill"})
	}

	user_id := params["id"]
	index := params["index"]

	i, err := strconv.Atoi(index)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(entities.ResponseModel{Message: "invalid cart index"})
	}

	err = h.CartService.UpdateOrder(user_id, i, cartData)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(entities.ResponseModel{Message: "cannot update cart data"})
	}

	return ctx.Status(fiber.StatusOK).JSON(entities.ResponseModel{Message: "success", Data: cartData})
}

func (h HTTPGateway) DeleteOrder(ctx *fiber.Ctx) error {

	params := ctx.Queries()

	if len(params) <= 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(entities.ResponseModel{Message: "user id not fill"})
	}

	user_id := params["id"]
	index := params["index"]

	i, err := strconv.Atoi(index)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(entities.ResponseModel{Message: "invalid cart index"})
	}

	err = h.CartService.DeleteAddress(user_id, i)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(entities.ResponseModel{Message: "cannot delete cart data"})
	}

	return ctx.Status(fiber.StatusOK).JSON(entities.ResponseModel{Message: "success"})
}
