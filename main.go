package main

import (
	"github.com/swisshacks-waton/internal"
)

func main() {
	// Generate a mortgage contract
	contract, err := internal.GenerateContract(internal.ContractTypemortgage)
	if err != nil {
		panic(err)
	}

	// Create client details
	client := internal.ClientDetails{
		Name:      "John",
		Surname:   "Doe",
		Email:     "john.doe@example.com",
		BirthDate: "1985-04-12",
		Address:   "123 Main St, Anytown, USA",
		Phone:     "+1-555-123-4567",
		SSN:       "123-45-6789",
	}

	// Create banking details
	banking := internal.NewBankingDetails("123456789", "Bank of America", "US12345678901234567890", "BOFAUS3N", 50000.00)

	// Add transactions
	banking.AddTransaction(internal.Transaction{
		ID:            "txn-001",
		Description:   "Salary deposit",
		Amount:        7000.00,
		Date:          "2025-03-01",
		Type:          "credit",
		Category:      "income",
		Tags:          []string{"salary", "monthly"},
		Notes:         "Monthly salary",
		Location:      "Anytown",
		PaymentMethod: "wire",
		Currency:      "USD",
		ExchangeRate:  1.00,
	})
	banking.AddTransaction(internal.Transaction{
		ID:            "txn-002",
		Description:   "Rent payment",
		Amount:        2000.00,
		Date:          "2025-03-03",
		Type:          "debit",
		Category:      "housing",
		Tags:          []string{"rent"},
		Notes:         "March rent",
		Location:      "Anytown",
		PaymentMethod: "ACH",
		Currency:      "USD",
		ExchangeRate:  1.00,
	})

	// Add investments
	banking.AddInvestment(internal.Investment{
		ID:          "inv-001",
		Type:        "stocks",
		Amount:      15000.00,
		Currency:    "USD",
		StartDate:   "2023-01-01",
		EndDate:     "",
		Description: "Tech portfolio",
		Institution: "Robinhood",
	})

	// Add a loan
	banking.AddLoan(internal.Loan{
		ID:              "loan-001",
		Type:            "mortgage",
		Principal:       300000.00,
		Outstanding:     250000.00,
		InterestRate:    3.5,
		StartDate:       "2020-01-01",
		EndDate:         "2050-01-01",
		Lender:          "Wells Fargo",
		PaymentSchedule: "monthly",
	})

	// Add insurance
	banking.AddInsurancePolicy(internal.InsurancePolicy{
		ID:            "ins-001",
		Provider:      "StateFarm",
		Type:          "life",
		Coverage:      500000.00,
		Premium:       500.00,
		StartDate:     "2022-01-01",
		EndDate:       "2032-01-01",
		Beneficiaries: "Jane Doe",
	})

	// Add real estate
	banking.AddRealEstate(internal.RealEstate{
		ID:           "re-001",
		Address:      "456 Elm St, Anytown, USA",
		Type:         "residential",
		Value:        350000.00,
		PurchaseDate: "2019-05-01",
		Ownership:    "full",
	})

	// Add business interest
	banking.AddBusinessInterest(internal.BusinessInterest{
		ID:           "biz-001",
		BusinessName: "Doe Ventures LLC",
		Ownership:    "50%",
		Value:        100000.00,
		Role:         "partner",
		StartDate:    "2020-06-01",
	})

	// Add other assets
	banking.AddOtherAsset(internal.OtherAsset{
		ID:           "asset-001",
		Description:  "Art collection",
		Value:        25000.00,
		Type:         "collectibles",
		AcquiredDate: "2018-03-15",
	})

	// Add liabilities
	banking.AddOtherLiability(internal.OtherLiability{
		ID:          "liab-001",
		Description: "Credit card debt",
		Amount:      8000.00,
		DueDate:     "2025-04-30",
		Creditor:    "Chase",
	})

	// Add other income
	banking.AddOtherIncome(internal.OtherIncome{
		ID:           "inc-001",
		Description:  "Freelance web development",
		Amount:       2000.00,
		Frequency:    "monthly",
		Source:       "Upwork",
		DateReceived: "2025-03-10",
	})

	// Add other expenses
	banking.AddOtherExpense(internal.OtherExpense{
		ID:          "exp-001",
		Description: "Gym membership",
		Amount:      100.00,
		Category:    "health",
		Date:        "2025-03-05",
	})

	// Add other debts
	banking.AddOtherDebt(internal.OtherDebt{
		ID:          "debt-001",
		Description: "Student loan",
		Amount:      25000.00,
		DueDate:     "2030-12-01",
		Lender:      "Navient",
	})

	// Create contract part and add to contract
	contractPart := internal.NewContractPart(
		"Contract Part A",
		client,
		banking,
	)

	contract.AddParty(contractPart)

	NewRiskAssessment := internal.NewRiskAssessment(
		contract,
		"medium",
		[]string{"Review contract terms", "Consult with a legal advisor"},
		"Risk assessment based on client details and banking information.")

	NewRiskAssessment.CalculateRiskScore()

}
