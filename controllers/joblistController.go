package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/GabrielEdwinSP/GolangDeveloperTest/dto"
	"github.com/GabrielEdwinSP/GolangDeveloperTest/initializers"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/schema"
)

func RequestJobList(c *gin.Context) {
	var request *dto.Request
	// var out io.Writer
	var err error

	c.Bind((&request))

	resp, err := http.Get(request.Host)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to get that link",
		})

		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read Body",
		})

		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": "user.ID",
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})

		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	// enc := json.NewEncoder(out)
	// enc.SetIndent("", "    ")

	// json.Unmarshal([]byte(body), response)
	// var response []dto.JobList
	// sb := string(body)
	// // sb, _ := json.MarshalIndent(body, "", "\t")
	// // sb, _ := json.Marshal(obj)
	// c.String(http.StatusOK, sb)

	jsonString := body
	jsonBytes := []byte(jsonString)

	joblist := []dto.JobList{}

	err = json.Unmarshal(jsonBytes, &joblist)
	// fmt.Scanf("%v", joblist)
	c.JSON(http.StatusOK, gin.H{
		"data": joblist,
	})

	result := initializers.DB.Create(&joblist)
	if result.Error != nil {
		c.Status(400)
		return
	}
}

func RequestJobListSearch(c *gin.Context) {
	var request *dto.Request
	var decoder = schema.NewDecoder()

	u, _ := url.Parse(request.Host)

	c.Bind((&request))

	err := decoder.Decode(&request, u.Query())
	if err != nil {
		panic(err)
	}

	// c.ShouldBindUri((&request))

	resp, err := http.Get(request.Host)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to get that link",
		})

		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read Body",
		})

		return
	}

	jsonString := body
	jsonBytes := []byte(jsonString)

	joblist := []dto.JobList{}

	err = json.Unmarshal(jsonBytes, &joblist)
	// fmt.Scanf("%v", joblist)
	c.JSON(http.StatusOK, gin.H{
		"data": joblist,
	})

}
