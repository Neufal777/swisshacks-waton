package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/swisshacks-waton/internal"
)

func main() {
	// Load data from the JSON file
	clientData, err := loadClientData("data/client_data.json")
	if err != nil {
		log.Fatalf("Error loading client data: %v", err)
	}

	// Generate a mortgage contract
	contract, err := internal.GenerateContract(internal.ContractTypemortgage)
	if err != nil {
		panic(err)
	}

	// Create client details from the JSON data
	client := internal.ClientDetails{
		Name:      clientData.Client.Name,
		Surname:   clientData.Client.Surname,
		Email:     clientData.Client.Email,
		BirthDate: clientData.Client.BirthDate,
		Address:   clientData.Client.Address,
		Phone:     clientData.Client.Phone,
		SSN:       clientData.Client.SSN,
	}

	// Create banking details from the JSON data
	banking := internal.NewBankingDetails(
		clientData.BankingDetails.AccountNumber,
		clientData.BankingDetails.BankName,
		clientData.BankingDetails.IBAN,
		clientData.BankingDetails.SwiftCode,
		clientData.BankingDetails.Balance,
	)

	// Add transactions from the JSON data
	for _, txn := range clientData.Transactions {
		banking.AddTransaction(internal.Transaction{
			ID:            txn.ID,
			Description:   txn.Description,
			Amount:        txn.Amount,
			Date:          txn.Date,
			Type:          txn.Type,
			Category:      txn.Category,
			Tags:          txn.Tags,
			Notes:         txn.Notes,
			Location:      txn.Location,
			PaymentMethod: txn.PaymentMethod,
			Currency:      txn.Currency,
			ExchangeRate:  txn.ExchangeRate,
		})
	}

	// Add investments from the JSON data
	for _, inv := range clientData.Investments {
		banking.AddInvestment(internal.Investment{
			ID:          inv.ID,
			Type:        inv.Type,
			Amount:      inv.Amount,
			Currency:    inv.Currency,
			StartDate:   inv.StartDate,
			EndDate:     inv.EndDate,
			Description: inv.Description,
			Institution: inv.Institution,
		})
	}

	// Add loans from the JSON data
	for _, loan := range clientData.Loans {
		banking.AddLoan(internal.Loan{
			ID:              loan.ID,
			Type:            loan.Type,
			Principal:       loan.Principal,
			Outstanding:     loan.Outstanding,
			InterestRate:    loan.InterestRate,
			StartDate:       loan.StartDate,
			EndDate:         loan.EndDate,
			Lender:          loan.Lender,
			PaymentSchedule: loan.PaymentSchedule,
		})
	}

	// Add insurance policies from the JSON data
	for _, ins := range clientData.InsurancePolicies {
		banking.AddInsurancePolicy(internal.InsurancePolicy{
			ID:            ins.ID,
			Provider:      ins.Provider,
			Type:          ins.Type,
			Coverage:      ins.Coverage,
			Premium:       ins.Premium,
			StartDate:     ins.StartDate,
			EndDate:       ins.EndDate,
			Beneficiaries: ins.Beneficiaries,
		})
	}

	// Create contract part and add to contract
	contractPart := internal.NewContractPart(
		"Contract Part A",
		client,
		banking,
	)

	contract.AddParty(contractPart)

	// Create a risk assessment
	NewRiskAssessment := internal.NewRiskAssessment(
		contract,
		"medium",
		[]string{"Review contract terms", "Consult with a legal advisor"},
		"Risk assessment based on client details and banking information.",
	)

	NewRiskAssessment.CalculateRiskScore()
}

// Load client data from the JSON file
func loadClientData(filename string) (ClientData, error) {
	var data ClientData
	fileContent, err := ioutil.ReadFile(filename)
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(fileContent, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}

// Define structures to parse the JSON data
type ClientData struct {
	Client            internal.ClientDetails     `json:"client"`
	BankingDetails    internal.BankingDetails    `json:"banking_details"`
	Transactions      []internal.Transaction     `json:"transactions"`
	Investments       []internal.Investment      `json:"investments"`
	Loans             []internal.Loan            `json:"loans"`
	InsurancePolicies []internal.InsurancePolicy `json:"insurance_policies"`
}
