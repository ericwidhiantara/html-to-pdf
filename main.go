package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var gotenbergAPI string

func init() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	gotenbergAPI = os.Getenv("GOTENBERG_API") + "/forms/chromium/convert/html"
	if gotenbergAPI == "" {
		panic("GOTENBERG_API environment variable is not set")
	}
}

func main() {
	router := gin.Default()

	router.POST("/generate-html-pdf", func(c *gin.Context) {
		htmlString := c.PostForm("html")

		if htmlString == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "HTML string is empty"})
			return
		}

		pdfBytes, err := generatePDF(htmlString)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate PDF"})
			return
		}

		// Set response headers
		c.Header("Content-Type", "application/pdf")
		c.Header("Content-Disposition", "attachment; filename=index.pdf")

		// Write PDF bytes to response body
		c.Data(http.StatusOK, "application/pdf", pdfBytes)
	})

	// Run the server
	router.Run(":5000")
}

func generatePDF(htmlString string) ([]byte, error) {
	// Create HTTP multipart request
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("files", "index.html")
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, bytes.NewBufferString(htmlString))
	if err != nil {
		return nil, err
	}
	writer.Close()

	// Send POST request
	request, err := http.NewRequest("POST", gotenbergAPI, body)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Read PDF bytes from response body
	var pdfBuffer bytes.Buffer
	_, err = io.Copy(&pdfBuffer, response.Body)
	if err != nil {
		return nil, err
	}

	return pdfBuffer.Bytes(), nil
}
