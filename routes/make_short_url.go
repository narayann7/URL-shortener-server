package routes

import (
	"encoding/json"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/narayann7/gourl/database"
	"github.com/narayann7/gourl/models"
	srv "github.com/narayann7/gourl/services"
)

func MakeShortUrl(c *fiber.Ctx) error {
	rdb1 := database.DatabaseInit(1)
	defer rdb1.Close()
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

	//create new request Struct
	reqBody := new(models.Request)
	//parce the json string to Request Struct
	err = json.Unmarshal(c.Body(), reqBody)
	//if error painc
	if err != nil {
		panic(srv.AppError{
			RealMessage:   err.Error(),
			Message:       "Unable to parse the body",
			ErrorCode:     fiber.StatusInternalServerError,
			ErrorLocation: "from MakeShortUrl, json.Unmarshal",
		})
	}
	//	check for valid url
	if !srv.IsVaildUrl(reqBody.Url) {
		errorMessage := ""
		if reqBody.Url == "" {
			errorMessage = "url parameter is require"
		} else {
			errorMessage = "not a valid url"
		}
		panic(srv.AppError{
			Message:   errorMessage,
			ErrorCode: 400,
		})

	}

	//check for vaild expiry
	if reqBody.Expiry == 0 {
		//if expiry is 0 by default url is vaild for 1440 minutes which is 1 day
		reqBody.Expiry = time.Minute * 1440
	} else {
		if vaild := srv.CheckForVaildExpiry(&reqBody.Expiry); !vaild {
			panic(srv.AppError{
				Message:   "invaild expiry. expiry will be in minutes",
				ErrorCode: 400,
			})
		}
	}

	if len(reqBody.CustomShortUrl) != 0 {
		//check for vaild custom url
		if isVaildCustomUrl := srv.CheckForVaildCustomUrl(reqBody.CustomShortUrl); !isVaildCustomUrl {
			panic(srv.AppError{
				Message:       "invaild custom Url",
				RealMessage:   "custom Url should consist of Alphabets and digits and size of url should 8 in length",
				ErrorLocation: "CheckForVaildCustomUrl",
				ErrorCode:     400,
			})
		} else if !database.CheckIsHashUnique(reqBody.CustomShortUrl) {
			//if its valid check that its already used or not
			panic(srv.AppError{
				Message:   "custom url is already in use",
				ErrorCode: 400,
			})
		}
	} else {
		reqBody.CustomShortUrl = ""
		//create a unique hash for the url
		urlHash := database.GetNewHash()
		reqBody.CustomShortUrl = urlHash
	}

	//check for domain error
	if err := srv.CheckForDominError(reqBody.Url); !err {
		panic(srv.AppError{
			Message:   "this url is not allowed",
			ErrorCode: 400,
		})
	}
	//enforce http
	reqBody.Url = srv.EnforceHttp(reqBody.Url)
	//store the url with url hash and expiry in db
	if result := database.SetUrlInDb(reqBody.Url, reqBody.CustomShortUrl, reqBody.Expiry); !result {
		panic(srv.AppError{
			Message:   "Something went wrong",
			ErrorCode: 500,
		})
	}
	timeRemaining, _ := rdb1.TTL(database.Ctx, c.IP()).Result()

	resBody := models.Responce{
		Url:           reqBody.Url,
		NewUrl:        "http://" + os.Getenv("DOMAIN") + "/" + reqBody.CustomShortUrl,
		Expiry:        reqBody.Expiry / time.Minute,
		RateRemaining: timeRemaining.String(),
	}

	return c.JSON(fiber.Map{
		"result": resBody,
	})

}
