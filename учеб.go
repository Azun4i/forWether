package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"regexp"
	"strconv"
)

type Numbers struct {
	Numbers string `json:"numbers,omitempty"`
}

type Result struct {
	Result int `json:"result"`
}

func sliceAtoi(sa []string) ([]int, error) {
	si := make([]int, 0, len(sa))
	for _, a := range sa {
		i, err := strconv.Atoi(a)
		if err != nil {
			return si, err
		}
		si = append(si, i)
	}
	return si, nil
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		if c.Request().Header.Get("User-Access") == "" {
			e.Logger.Infof("access denied")
			return c.JSON(http.StatusForbidden, "access denied")
		}
		req := &Numbers{}
		if err := c.Bind(req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if _, err := strconv.Atoi(req.Numbers[len(req.Numbers):]); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "строка дожна заканчиваться числом")
		}
		re := regexp.MustCompile("[0-9]+")
		numbers := re.FindAllString(req.Numbers, -1)
		if len(numbers) == 0 || len(numbers) == 1 {
			return c.JSON(http.StatusBadRequest, "не достаточно чисел")
		}

		symbolsRegexp := regexp.MustCompile("[-+]")
		symbols := symbolsRegexp.FindAllString(req.Numbers, -1)

		intArray, _ := sliceAtoi(numbers)
		result := 0
		for i, num := range intArray {
			if i+1 <= len(symbols)-1 {
				switch symbols[i] {
				case "+":
					result += num
				case "-":
					result -= num
				}
			}

		}
		return c.JSON(http.StatusOK, &Result{
			Result: result,
		})
	})
	e.Logger.Fatal(e.Start(":1323"))
}
