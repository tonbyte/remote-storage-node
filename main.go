package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/tonbyte/remote-storage-node/config"
	"github.com/tonbyte/remote-storage-node/storage"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"golang.org/x/exp/slices"
)

var Version = "0.1"

func validateBagID(bagID string) bool {
	valid, err := regexp.Match(`^([abcdefABCDEF0-9]*)?`, []byte(bagID))
	if !valid || err != nil {
		log.Warn(fmt.Sprintf("invalid bag ID - %e", err))
		return false
	}

	return true
}

func filterIP(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if !slices.Contains(config.StorageConfig.WhitelistIPs, c.RealIP()) {
			message := fmt.Sprintf("IP address %s not allowed", c.RealIP())
			log.Info(message)
			return echo.NewHTTPError(http.StatusUnauthorized, message)
		}

		return next(c)
	}
}

func main() {
	e := echo.New()
	e.Use(filterIP)
	config.LoadConfig()

	e.GET("/status", func(c echo.Context) error {
		log.Info("/status")

		return c.JSON(http.StatusOK, echo.Map{
			"version": Version,
		})
	})

	e.GET("/addBag/:bagID", func(c echo.Context) error {
		bagID := c.Param("bagID")
		log.Info("/addBag/", bagID)

		if !validateBagID(bagID) {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "invalid bag ID",
			})
		}

		if !storage.AddBag(bagID) {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "can not add bag. Check logs",
			})
		}

		return c.JSON(http.StatusOK, echo.Map{
			"message": "ok",
		})
	})

	e.GET("/removeBag/:bagID", func(c echo.Context) error {
		bagID := c.Param("bagID")
		log.Info("/removeBag/", bagID)

		if !validateBagID(bagID) {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "invalid bag ID",
			})
		}

		if !storage.RemoveBag(bagID) {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "can not remove bag. Check logs",
			})
		}

		return c.JSON(http.StatusOK, echo.Map{
			"message": "ok",
		})
	})

	e.Logger.Fatal(e.Start(":" + strconv.Itoa(config.StorageConfig.Port)))
}
