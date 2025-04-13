package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	RiskAssessmentTypefinancial   = "financial"
	RiskAssessmentTypelegal       = "legal"
	RiskAssessmentTypeoperational = "operational"
	RiskAssessmentTypecompliance  = "compliance"
	RiskAssessmentTypecredit      = "credit"
	RiskAssessmentTypefraud       = "fraud"
)

type RiskAssessment struct {
	ContractID      Contract `json:"contract"`        // Unique identifier for the contract
	Score           int      `json:"score"`           // Risk score (0-100)
	Level           string   `json:"level"`           // Risk level (low, medium, high)
	Recommendations []string `json:"recommendations"` // Recommendations for risk mitigation
	Details         string   `json:"details"`         // Detailed explanation of the risk assessment
}

func NewRiskAssessment(contract Contract, level string, recommendations []string, details string) RiskAssessment {
	// Create a new risk assessment
	riskAssessment := RiskAssessment{
		ContractID:      contract,
		Score:           0,
		Recommendations: recommendations,
		Details:         details,
	}
	return riskAssessment
}
func (ra *RiskAssessment) CalculateRiskScore() {
	prompt := fmt.Sprintf(`
		You are a risk analyst. Based on the following contract details, provide:
		1. A risk score between 0 and 100
		2. A risk level (low, medium, high)
		3. 2-3 recommendations

		Details:
		%s

		Respond in this format:
		Score: <score>
		Level: <level>
		Recommendations:
		- <rec1>
		- <rec2>
		- <rec3>
		`, ra.Details)

	// OpenAI API completion request
	payload := map[string]interface{}{
		"model": "gpt-4", // You can change this to whichever model you are using
		"messages": []map[string]string{
			{
				"role":    "system",
				"content": "You are a helpful risk analyst.",
			},
			{
				"role":    "user",
				"content": prompt,
			},
		},
		"max_tokens":  200,
		"temperature": 0.7,
	}

	data, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Failed to marshal payload:", err)
		return
	}

	// Make the HTTP request to OpenAI API
	resp, err := http.Post("https://api.openai.com/v1/chat/completions", "application/json", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Error calling OpenAI API:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	// Handle the response from the OpenAI API
	var response struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error parsing response:", err)
		return
	}

	// Check if choices are available in the response
	if len(response.Choices) == 0 {
		fmt.Println("No choices returned from the API.")
		return
	}

	// Parse the response
	completion := response.Choices[0].Message.Content
	lines := strings.Split(completion, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "Score:") {
			fmt.Sscanf(line, "Score: %d", &ra.Score)
		}
		if strings.HasPrefix(line, "Level:") {
			ra.Level = strings.TrimSpace(strings.TrimPrefix(line, "Level:"))
		}
		if strings.HasPrefix(line, "- ") {
			ra.Recommendations = append(ra.Recommendations, strings.TrimPrefix(line, "- "))
		}
	}
}
