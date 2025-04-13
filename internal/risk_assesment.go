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
	Contract        Contract `json:"contract"`        // Unique identifier for the contract
	Score           int      `json:"score"`           // Risk score (0-100)
	Level           string   `json:"level"`           // Risk level (low, medium, high)
	Recommendations []string `json:"recommendations"` // Recommendations for risk mitigation
	Details         string   `json:"details"`         // Detailed explanation of the risk assessment
}

func NewRiskAssessment(contract Contract, level string, recommendations []string, details string) RiskAssessment {
	// Create a new risk assessment
	riskAssessment := RiskAssessment{
		Contract:        contract,
		Score:           0,
		Recommendations: recommendations,
		Details:         details,
	}
	return riskAssessment
}
func (ra *RiskAssessment) CalculateRiskScore() {
	// Building a complex prompt based on contract and banking details
	var contractDetailsBuilder strings.Builder
	var totalAmountInUSD float64

	contractDetailsBuilder.WriteString(fmt.Sprintf("Contract ID: %s\n", ra.Contract.ID))
	contractDetailsBuilder.WriteString(fmt.Sprintf("Contract Name: %s\n", ra.Contract.Name))
	contractDetailsBuilder.WriteString(fmt.Sprintf("Contract Description: %s\n", ra.Contract.Description))
	contractDetailsBuilder.WriteString(fmt.Sprintf("Contract Type: %s\n", ra.Contract.Type))

	// Calculate the total asking amount in USD
	contractDetailsBuilder.WriteString("Parties Involved:\n")
	for _, party := range ra.Contract.Parties {
		contractDetailsBuilder.WriteString(fmt.Sprintf("- Party: %s, Client: %s %s, Email: %s, SSN: %s\n", party.Name, party.ClientDetails.Name, party.ClientDetails.Surname, party.ClientDetails.Email, party.ClientDetails.SSN))
		contractDetailsBuilder.WriteString(fmt.Sprintf("  - Address: %s\n", party.ClientDetails.Address))
		contractDetailsBuilder.WriteString(fmt.Sprintf("  - Phone: %s\n", party.ClientDetails.Phone))
		contractDetailsBuilder.WriteString("  Banking Details:\n")
		contractDetailsBuilder.WriteString(fmt.Sprintf("    - Bank: %s, Account: %s, IBAN: %s, Balance: %f\n", party.BankingDetails.BankName, party.BankingDetails.AccountNumber, party.BankingDetails.IBAN, party.BankingDetails.Balance))

		// Add details for transactions, loans, investments
		contractDetailsBuilder.WriteString("    - Recent Transactions:\n")
		for _, transaction := range party.BankingDetails.TransactionHistory {
			amount := transaction.Amount // Ensure amounts are parsed as float64
			totalAmountInUSD += amount   // Adding to total amount
			contractDetailsBuilder.WriteString(fmt.Sprintf("      - %s: %s %s %s, Amount: %f\n", transaction.Date, transaction.Description, transaction.Category, transaction.Location, transaction.Amount))
		}
		contractDetailsBuilder.WriteString("    - Investments:\n")
		for _, investment := range party.BankingDetails.Investments {
			amount := investment.Amount
			totalAmountInUSD += amount // Adding to total amount
			contractDetailsBuilder.WriteString(fmt.Sprintf("      - Investment in %s: Amount %.2f, Currency: %s\n", investment.Description, investment.Amount, investment.Currency))
		}
		contractDetailsBuilder.WriteString("    - Loans:\n")
		for _, loan := range party.BankingDetails.Loans {
			amount := loan.Principal
			totalAmountInUSD += amount // Adding to total amount
			contractDetailsBuilder.WriteString(fmt.Sprintf("      - Loan Type: %s, Amount: %.2f, Outstanding: %f, Interest Rate: %f\n", loan.Type, loan.Principal, loan.Outstanding, loan.InterestRate))
		}
	}

	// Asking in USD (calculated total amount)
	contractDetailsBuilder.WriteString(fmt.Sprintf("Total Asking in USD: %.2f\n", totalAmountInUSD))

	prompt := fmt.Sprintf(`
	This is a simulation for educational purposes and involves no real financial data. You are a risk analyst tasked with evaluating a hypothetical contract based on the details provided. Based on the following contract details, please provide the following:
	1. A risk score between 0 and 100 (0 = low risk, 100 = high risk).
	2. A risk level: low, medium, or high.
	3. 2-3 specific recommendations on how to mitigate the risks identified.
	4. A final decision: "Accept Credit", "Assess Personally", or "Too Risky". Make the decision based on the financial stability and debts of the parties involved. This is for a learning exercise and not for real-world implementation.
	
	Contract Details:
	%s
	
	Please respond in the following exact format:
	- Score: <score>
	- Level: <level>
	Recommendations:
	- <rec1>
	- <rec2>
	- <rec3>
	Decision: <decision>
	`, contractDetailsBuilder.String())

	// Print the prompt for debugging purposes
	println("Llama 3.2 model:")
	println(contractDetailsBuilder.String())

	// Prepare the payload for the local Llama 3.2 model
	payload := map[string]interface{}{
		"model": "llama3.2:1b", // Adjust to your local model's identifier
		"messages": []map[string]string{
			{"role": "system", "content": "You are a helpful risk analyst."},
			{"role": "user", "content": prompt},
		},
		"max_tokens":  200,
		"temperature": 0.7,
	}

	// Marshal the payload into JSON
	data, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Failed to marshal payload:", err)
		return
	}

	// Create a new HTTP request to the local server
	req, err := http.NewRequest("POST", "http://localhost:1337/v1/chat/completions", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Failed to create request:", err)
		return
	}

	// Set the necessary headers
	req.Header.Set("Content-Type", "application/json")

	// Initialize the HTTP client and make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error calling local Llama API:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Handle the response from the local Llama API
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

	// Parse the response content
	completion := response.Choices[0].Message.Content
	lines := strings.Split(completion, "\n")
	println("------------------------------------------------------------")
	for _, line := range lines {
		// Clean up the line
		line = strings.ReplaceAll(line, "*", "")
		line = strings.TrimSpace(line)
		println(line) // Print each line for debugging

		// Parse risk score
		if strings.HasPrefix(line, "Risk Score:") {
			// Extract the risk score
			fmt.Sscanf(line, "Risk Score: %d", &ra.Score)

			// Determine the risk level based on the risk score
			if ra.Score <= 30 {
				ra.Level = "Low"
			} else if ra.Score <= 70 {
				ra.Level = "Medium"
			} else {
				ra.Level = "High"
			}
		}

		// Parse recommendations
		if strings.HasPrefix(line, "Given the risk score of") {
			// Starting point for recommendations, might need custom parsing logic here
			// Assuming next lines are recommendations:
			// Iterate until you hit the decision part or a blank line
			var recommendations []string
			for _, recommendationLine := range lines {
				if strings.HasPrefix(recommendationLine, "I recommend") {
					recommendations = append(recommendations, strings.TrimSpace(recommendationLine))
				} else {
					break
				}
			}
			ra.Recommendations = append(ra.Recommendations, recommendations...)
		}

		// Parse decision
		if strings.HasPrefix(line, "Decision:") {
			ra.Details = strings.TrimSpace(strings.TrimPrefix(line, "Decision:"))
		}
	}

	// Print the results
	fmt.Printf("Risk Score: %d\n", ra.Score)
	fmt.Printf("Risk Level: %s\n", ra.Level)
	fmt.Println("Recommendations:")
	for _, rec := range ra.Recommendations {
		fmt.Printf("- %s\n", rec)
	}
	fmt.Printf("Final Decision: %s\n", ra.Details)
}
