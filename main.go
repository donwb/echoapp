package main

import (
	"fmt"
	"log"
	"encoding/json"
	"net/http"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error{
		return c.String(http.StatusOK, "Hello World")
	})
	

	e.GET("/cats", GetCats)
	e.GET("/jsoncats/:data", GetCatsPath)
	e.POST("/cats", AddCat)

	e.Logger.Fatal(e.Start(":8000"))
}

// http://localhost:8000/cats?name=Cali&type=cute
func GetCats(c echo.Context) error {
	catName := c.QueryParam("name")
	catType := c.QueryParam("type")

	return c.String(http.StatusOK, fmt.Sprintf("your cat name is %s\n and type is %s", catName, catType))

}

// http://localhost:8000/cats/json?name=arnold&type=fluffy
// data path variable accepts value as json/string
func GetCatsPath(c echo.Context) error {
	catName := c.QueryParam("name")
	catType := c.QueryParam("type")
	dataType := c.Param("data")

	if dataType == "stirng" {
		return c.String(http.StatusOK, fmt.Sprintf("your cat name is : %s\nand cat type is : %s\n", catName, catType))
	} else if dataType == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"name" : catName,
			"type" : catType})
	} else {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Please specifiy the data type as string or json"})
	}

}


func AddCat(c echo.Context) error {
	type Cat struct {
		Name string `json:"name"`
		Type string `json:"type"`
	}

	cat := &Cat{Name: "Don", Type: "Tabby"}

	defer c.Request().Body.Close()

	log.Printf("%#v", cat)

	//err := json.NewDecoder(c.Request().Body).Decode(&cat)
	data, err := json.Marshal(cat)
	log.Printf("%s", data)

	if err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	log.Printf("this is your cat %#v", cat)
	return c.JSON(http.StatusOK, string(data))

}



