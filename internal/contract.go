package internal

import (
	"fmt"

	"github.com/google/uuid"
)

// Contract represents a contract with a unique ID, a name, and a description.
// This tool will be used to generate a contract for a specific use case.

const (
	// ContractType represents the type of contract.
	ContractTypemortgage    = "mortgage"
	ContractTypelease       = "lease"
	ContractTypeinsurance   = "insurance"
	ContractTypepurchase    = "purchase"
	ContractTypepartnership = "partnership"
)

type Contract struct {
	ID          string         `json:"id"`            // Unique identifier for the contract
	Name        string         `json:"name"`          // Name of the contract
	Description string         `json:"description"`   // Description of the contract
	Type        string         `json:"type"`          // Type of the contract (e.g., mortgage, lease, etc.)
	AskingInUSD int            `json:"asking_in_usd"` // Asking price in USD
	Parties     []ContractPart `json:"parties"`       // List of parties involved in the contract
}

// GenerateContract creates a new contract based on the specified type.
func GenerateContract(contractType string) (Contract, error) {
	// Validate the contract type
	if contractType != ContractTypemortgage && contractType != ContractTypelease &&
		contractType != ContractTypeinsurance && contractType != ContractTypepurchase && contractType != ContractTypepartnership {
		return Contract{}, fmt.Errorf("invalid contract type: %s", contractType)
	}
	// Generate a unique ID for the contract
	contractID := uuid.New().String()
	// Create the contract
	contract := Contract{
		ID:          fmt.Sprintf("contract-%s", contractID),
		Name:        "Contract for " + contractType,
		Description: "This is a contract for " + contractType,
		Type:        contractType,
	}
	// Return the generated contract
	return contract, nil
}

// AddParty adds a party to the contract.
func (c *Contract) AddParty(party ContractPart) {
	c.Parties = append(c.Parties, party)
}

// RemoveParty removes a party from the contract by name.
func (c *Contract) RemoveParty(name string) bool {
	for i, party := range c.Parties {
		if party.Name == name {
			c.Parties = append(c.Parties[:i], c.Parties[i+1:]...)
			return true
		}
	}
	return false
}

// GetParty retrieves a party from the contract by name.
func (c *Contract) GetParty(name string) (*ContractPart, bool) {
	for _, party := range c.Parties {
		if party.Name == name {
			return &party, true
		}
	}
	return nil, false
}

// ListPartyNames returns a list of names of all parties in the contract.
func (c *Contract) ListPartyNames() []string {
	names := make([]string, len(c.Parties))
	for i, party := range c.Parties {
		names[i] = party.Name
	}
	return names
}

// Summary returns a brief summary of the contract.
func (c *Contract) Summary() string {
	return fmt.Sprintf("Contract ID: %s\nName: %s\nType: %s\nParties Involved: %d", c.ID, c.Name, c.Type, len(c.Parties))
}
