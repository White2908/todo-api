package middleware

import (
	"log"
	"time"

	"github.com/go-fuego/fuego"
	"github.com/google/uuid"
)

func Logger() func(next func(fuego.ContextNoBody) error) func(fuego.ContextNoBody) error {
	return func(next func(fuego.ContextNoBody) error) func(fuego.ContextNoBody) error {
		return func(c fuego.ContextNoBody) error {
			start := time.Now() //start time of request

			err := next(c) //call next handler

			//Get log info
			log.Printf(
				"[%s] %s %s %s %v",
				time.Now().Format("2006-01-02 15:04:05"),
				c.Request().Method,
				c.Request().URL.Path,
				c.Request().RemoteAddr,
				time.Since(start),
			)

			return err
		}
	}
}

// Recover middleware

// Recover là middleware bắt panic và trả về lỗi 500
func Recover() func(next func(fuego.ContextNoBody) error) func(fuego.ContextNoBody) error {
	return func(next func(fuego.ContextNoBody) error) func(fuego.ContextNoBody) error {
		return func(c fuego.ContextNoBody) (err error) {
			// defer để bắt panic nếu có
			defer func() {
				if r := recover(); r != nil {

					log.Printf("Recovered from panic: %v", r)
					err = fuego.InternalServerError{
						Title:  "Internal Server Error",
						Detail: "An unexpected error occurred",
					}
				}
			}()

			return next(c)
		}
	}
}

// RequestID middleware
func RequestID() func(next func(fuego.ContextNoBody) error) func(fuego.ContextNoBody) error {
	return func(next func(fuego.ContextNoBody) error) func(fuego.ContextNoBody) error {
		return func(c fuego.ContextNoBody) error {
			requestID := c.Request().Header.Get("X-Request-ID")
			if requestID == "" {
				requestID = uuid.New().String()
			}
			c.Response().Header().Set("X-Request-ID", requestID)

			return next(c)
		}
	}
}
