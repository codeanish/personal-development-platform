package routes

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func CreateRepository(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "CreateRepository",
		"env": os.Getenv("API_TOKEN"),
	})
}

func GetRepositories(c *gin.Context) {
	url := os.Getenv("GITHUB_API_URL")
	method := "GET"

	client := &http.Client {}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error creating request",
		})
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")
	req.Header.Add("Authorization", "token " + os.Getenv("GITHUB_API_TOKEN"))
	
	res, err := client.Do(req)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error making request",
		})
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error reading response",
		})
		return
	}
	
	c.JSON(200, gin.H{
		"message": "GetRepositories",
		"body": string(body),
	})
}