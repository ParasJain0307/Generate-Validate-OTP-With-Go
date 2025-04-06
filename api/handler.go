package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/ParasJain0307/generate-validate-otp/data"
	"github.com/gin-gonic/gin"
)

const appTimeout = 10 * time.Second

func (app *Config) sendSMS() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		var payload data.OTPData
		defer cancel()

		app.validateBody(c, &payload)

		newData := data.OTPData{
			PhoneNumber: payload.PhoneNumber,
		}

		_, err := app.twilioSendOTP(newData.PhoneNumber)
		if err != nil {
			app.errorJSON(c, err)
			return
		}
		app.writeJSON(c, http.StatusAccepted, "OTP Sent Successfully.")
	}

}

func (app *Config) verifySMS() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		var payload data.VerifyData
		defer cancel()

		app.validateBody(c, &payload)

		newData := data.VerifyData{
			User: payload.User,
			Code: payload.Code,
		}

		err := app.twilioVerifyOTP(newData.User.PhoneNumber, newData.Code)
		fmt.Println("err", err)
		if err != nil {
			app.errorJSON(c, err)
			return
		}
		app.writeJSON(c, http.StatusAccepted, "OTP Verfied Successfully.")
	}
}
