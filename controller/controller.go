package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/Iretoms/hng-task-one/model"
	"github.com/Iretoms/hng-task-one/response"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func HelloCall() gin.HandlerFunc {
	return func(c *gin.Context) {
		visitor_name := c.DefaultQuery("visitor_name", "Guest")
		clientIp, location := getIpLoc()
		fmt.Printf("clientIp:%v, location:%v", clientIp, location)
		temp := getTemp(location)

		c.JSON(http.StatusOK, response.HelloResponse{ClientIp: clientIp, Location: location, Greeting: fmt.Sprintf("Hello, %v!, the temperature is %v degrees celsius in %v", visitor_name, temp, location)})
	}
}

func getIpLoc() (string, string) {
	response, err := http.Get("https://ipv4.geojs.io/v1/ip/geo.js")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile(`\{.*\}`)
	jsonStr := re.FindString(string(responseData))

	var responseObject model.GeoData

	err = json.Unmarshal([]byte(jsonStr), &responseObject)
	if err != nil {
		log.Fatal(err)
	}

	return responseObject.IP, responseObject.Location
}

func getTemp(loc string) float64 {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("API key not set")
	}

	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%v&q=%v", apiKey, loc)

	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	var responseObject model.TempResponse
	json.Unmarshal(responseData, &responseObject)

	temp := responseObject.CurrentRes.TempCelsius

	return temp
}
