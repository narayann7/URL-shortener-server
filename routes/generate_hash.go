package routes

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/narayann7/gourl/database"
	srv "github.com/narayann7/gourl/services"
)

func GenerateNewHashes(c *fiber.Ctx) error {

	//recover the panic and send as respond
	//with suitable message and status code
	defer srv.CatchErrors(c)

	//rate limiting
	err := database.CounterReducer(c.IP())
	if err != nil {
		panic(srv.AppError{
			Message:   err.Error(),
			ErrorCode: 400,
		})
	}

	//get size from parmas
	size := c.Params("size")
	//convert the size string to int
	sizeInInt, err := strconv.Atoi(size)
	if err != nil {
		panic(srv.AppError{
			Message:   "invaild size",
			ErrorCode: 400,
		})

	}
	if sizeInInt <= 0 || sizeInInt > 10 {
		fmt.Println("sdfug")
		panic(srv.AppError{
			Message:   "invaild size",
			ErrorCode: 400,
		})
	}
	//get the list hash which are not used
	hashList := srv.GetNewHashes(sizeInInt)

	return c.JSON(fiber.Map{
		"urls": hashList,
	})

}
