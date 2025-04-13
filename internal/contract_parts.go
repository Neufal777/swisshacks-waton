package internal

import (
	"fmt"

	"github.com/google/uuid"
)

// ContractParts represents a part of a contract with it's financial details.
// inclues  historical banking transactions, and other financial data.
// This tool will be used to generate a contract for a specific use case.
// and asses the risk of the contract.

type ClientDetails struct {
	Name      string `json:"name"`       // Name of the client
	Surname   string `json:"surname"`    // Surname of the client
	Email     string `json:"email"`      // Email of the client
	BirthDate string `json:"birth_date"` // Birth date of the client
	Address   string `json:"address"`    // Address of the client
	Phone     string `json:"phone"`      // Phone number of the client
	SSN       string `json:"ssn"`        // Social Security Number of the client
}

type ContractPart struct {
	ID              string         `json:"id"`
	Name            string         `json:"name"`            // Name of the contract part
	ClientDetails   ClientDetails  `json:"client_details"`  // Details of the client involved in the contract
	BankindgDetails BankingDetails `json:"banking_details"` // Banking details of the client
}

func NewContractPart(name string, clientDetails ClientDetails, bankingDetails BankingDetails) ContractPart {
	// Generate a unique ID for the contract part
	contractPartID := uuid.New().String()
	contractPart := ContractPart{
		ID:              fmt.Sprintf("contract-part-%s", contractPartID),
		Name:            name,
		ClientDetails:   clientDetails,
		BankindgDetails: bankingDetails,
	}
	return contractPart
}
