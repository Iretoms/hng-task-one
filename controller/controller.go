package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/Iretoms/hng-task-one/model"
	"github.com/Iretoms/hng-task-one/response"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func HelloCall() gin.HandlerFunc {
	return func(c *gin.Context) {
		visitor_name := c.DefaultQuery("visitor_name", "Guest")
		clientIp := c.ClientIP()
		location := getLoc(clientIp)
		temp := getTemp(location)

		c.JSON(http.StatusOK, response.HelloResponse{ClientIp: clientIp, Location: location, Greeting: fmt.Sprintf("Hello, %v!, the temperature is %v degrees celsius in %v", visitor_name, temp, location)})
	}
}

func getLoc(ip string) string {
	response, err := http.Get(fmt.Sprintf("https://ipv4.geojs.io/v1/ip/geo.js?ip=%v", ip))

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	jsonStr := extractJSON(string(responseData))

	var responseObject []model.GeoData

	err = json.Unmarshal([]byte(jsonStr), &responseObject)
	if err != nil {
		log.Fatal(err)
	}

	if len(responseObject) > 0 {
		return responseObject[0].Location
	}

	return ""
}

func extractJSON(jsonp string) string {

	jsonp = strings.TrimPrefix(jsonp, "geoip(")
	jsonp = strings.TrimSuffix(jsonp, ")")

	re := regexp.MustCompile(`\[.*\]`)
	return re.FindString(jsonp)
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
