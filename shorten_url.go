package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Substitua 'SUA_CHAVE_DE_API' pela chave de API real
	apiKey := os.Getenv("apiKey")
	longURL := "https://github.com/joho/godotenv"

	// Endpoint da API do Bitly para encurtar URLs
	endpoint := "https://api-ssl.bitly.com/v4/shorten"

	// Crie um struct para a carga útil (payload) JSON
	payload := map[string]string{
		"long_url": longURL,
	}

	// Converta a carga útil para JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Erro ao converter a carga útil para JSON:", err)
		return
	}

	// Faça a chamada à API do Bitly
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(payloadBytes))
	if err != nil {
		fmt.Println("Erro ao criar a requisição:", err)
		return
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erro ao fazer a chamada à API:", err)
		return
	}
	defer resp.Body.Close()

	// Leia a resposta da API
	var response map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Println("Erro ao decodificar a resposta da API:", err)
		return
	}

	// Verifique se a resposta foi bem-sucedida
	if resp.StatusCode == http.StatusOK {
		shortenedURL := response["link"].(string)
		fmt.Printf("Link encurtado: %s\n", shortenedURL)
	} else {
		fmt.Println("Erro ao encurtar o link. Resposta da API:", response)
	}
}
