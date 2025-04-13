package internal

type BankingDetails struct {
	AccountNumber      string             `json:"account_number"`      // Account number of the client
	BankName           string             `json:"bank_name"`           // Name of the bank
	IBAN               string             `json:"iban"`                // International Bank Account Number
	SwiftCode          string             `json:"swift_code"`          // SWIFT code of the bank
	Balance            string             `json:"balance"`             // Current balance of the account
	TransactionHistory []Transaction      `json:"transaction_history"` // Historical transactions
	Investments        []Investment       `json:"investments"`         // Investments made by the client
	Loans              []Loan             `json:"loans"`               // Loans taken by the client
	InsurancePolicies  []InsurancePolicy  `json:"insurance_policies"`  // Insurance policies held by the client
	RealEstate         []RealEstate       `json:"real_estate"`         // Real estate owned by the client
	BusinessInterests  []BusinessInterest `json:"business_interests"`  // Business interests of the client
	OtherAssets        []OtherAsset       `json:"other_assets"`        // Other assets owned by the client
	OtherLiabilities   []OtherLiability   `json:"other_liabilities"`   // Other liabilities of the client
	OtherIncome        []OtherIncome      `json:"other_income"`        // Other income sources of the client
	OtherExpenses      []OtherExpense     `json:"other_expenses"`      // Other expenses of the client
	OtherDebts         []OtherDebt        `json:"other_debts"`         // Other debts of the client

}

type Transaction struct {
	ID            string   `json:"id"`             // Unique identifier for the transaction
	Description   string   `json:"description"`    // Description of the transaction
	Amount        string   `json:"amount"`         // Amount of the transaction
	Date          string   `json:"date"`           // Date of the transaction
	Type          string   `json:"type"`           // Type of the transaction (e.g., credit, debit)
	Category      string   `json:"category"`       // Category of the transaction (e.g., groceries, rent, etc.)
	Tags          []string `json:"tags"`           // Tags associated with the transaction
	Notes         string   `json:"notes"`          // Notes associated with the transaction
	Location      string   `json:"location"`       // Location of the transaction
	PaymentMethod string   `json:"payment_method"` // Payment method used for the transaction
	Currency      string   `json:"currency"`       // Currency of the transaction
	ExchangeRate  string   `json:"exchange_rate"`  // Exchange rate used for the transaction
}

type Investment struct {
	ID          string `json:"id"`          // Unique identifier for the investment
	Type        string `json:"type"`        // Type of investment (e.g., stocks, bonds)
	Amount      string `json:"amount"`      // Amount invested
	Currency    string `json:"currency"`    // Currency of the investment
	StartDate   string `json:"start_date"`  // Date investment started
	EndDate     string `json:"end_date"`    // Date investment ended (if applicable)
	Description string `json:"description"` // Description or details of the investment
	Institution string `json:"institution"` // Institution managing the investment
}

type Loan struct {
	ID              string `json:"id"`               // Unique identifier for the loan
	Type            string `json:"type"`             // Type of loan (e.g., personal, mortgage)
	Principal       string `json:"principal"`        // Original loan amount
	Outstanding     string `json:"outstanding"`      // Outstanding balance
	InterestRate    string `json:"interest_rate"`    // Interest rate
	StartDate       string `json:"start_date"`       // Loan start date
	EndDate         string `json:"end_date"`         // Loan end date
	Lender          string `json:"lender"`           // Name of the lender
	PaymentSchedule string `json:"payment_schedule"` // Payment schedule details
}

type InsurancePolicy struct {
	ID            string `json:"id"`            // Unique identifier for the insurance policy
	Provider      string `json:"provider"`      // Insurance provider name
	Type          string `json:"type"`          // Type of policy (e.g., health, life)
	Coverage      string `json:"coverage"`      // Coverage amount
	Premium       string `json:"premium"`       // Premium amount
	StartDate     string `json:"start_date"`    // Policy start date
	EndDate       string `json:"end_date"`      // Policy end date
	Beneficiaries string `json:"beneficiaries"` // Named beneficiaries
}

type RealEstate struct {
	ID           string `json:"id"`            // Unique identifier for the property
	Address      string `json:"address"`       // Property address
	Type         string `json:"type"`          // Type of property (e.g., residential, commercial)
	Value        string `json:"value"`         // Current market value
	PurchaseDate string `json:"purchase_date"` // Date of purchase
	Ownership    string `json:"ownership"`     // Ownership status (e.g., full, joint)
}

type BusinessInterest struct {
	ID           string `json:"id"`            // Unique identifier for the business interest
	BusinessName string `json:"business_name"` // Name of the business
	Ownership    string `json:"ownership"`     // Ownership percentage
	Value        string `json:"value"`         // Valuation of the interest
	Role         string `json:"role"`          // Role in the business (e.g., investor, partner)
	StartDate    string `json:"start_date"`    // Date the interest began
}

type OtherAsset struct {
	ID           string `json:"id"`            // Unique identifier for the asset
	Description  string `json:"description"`   // Description of the asset
	Value        string `json:"value"`         // Value of the asset
	Type         string `json:"type"`          // Type/category of the asset
	AcquiredDate string `json:"acquired_date"` // Date of acquisition
}

type OtherLiability struct {
	ID          string `json:"id"`          // Unique identifier for the liability
	Description string `json:"description"` // Description of the liability
	Amount      string `json:"amount"`      // Amount owed
	DueDate     string `json:"due_date"`    // Due date for repayment
	Creditor    string `json:"creditor"`    // Name of the creditor
}

type OtherIncome struct {
	ID           string `json:"id"`            // Unique identifier for the income source
	Description  string `json:"description"`   // Description of the income source
	Amount       string `json:"amount"`        // Amount received
	Frequency    string `json:"frequency"`     // Frequency (e.g., monthly, annually)
	Source       string `json:"source"`        // Source of the income
	DateReceived string `json:"date_received"` // Date received
}

type OtherExpense struct {
	ID          string `json:"id"`          // Unique identifier for the expense
	Description string `json:"description"` // Description of the expense
	Amount      string `json:"amount"`      // Amount spent
	Category    string `json:"category"`    // Category of the expense
	Date        string `json:"date"`        // Date of the expense
}

type OtherDebt struct {
	ID          string `json:"id"`          // Unique identifier for the debt
	Description string `json:"description"` // Description of the debt
	Amount      string `json:"amount"`      // Amount owed
	DueDate     string `json:"due_date"`    // Due date for repayment
	Lender      string `json:"lender"`      // Name of the lender
}

func NewBankingDetails(accountNumber, bankName, iban, swiftCode, balance string) BankingDetails {
	return BankingDetails{
		AccountNumber: accountNumber,
		BankName:      bankName,
		IBAN:          iban,
		SwiftCode:     swiftCode,
		Balance:       balance,
	}
}

func (bd *BankingDetails) AddTransaction(transaction Transaction) {
	bd.TransactionHistory = append(bd.TransactionHistory, transaction)
}

func (bd *BankingDetails) AddInvestment(investment Investment) {
	bd.Investments = append(bd.Investments, investment)
}

func (bd *BankingDetails) AddLoan(loan Loan) {
	bd.Loans = append(bd.Loans, loan)
}

func (bd *BankingDetails) AddInsurancePolicy(policy InsurancePolicy) {
	bd.InsurancePolicies = append(bd.InsurancePolicies, policy)
}

func (bd *BankingDetails) AddRealEstate(realEstate RealEstate) {
	bd.RealEstate = append(bd.RealEstate, realEstate)
}

func (bd *BankingDetails) AddBusinessInterest(businessInterest BusinessInterest) {
	bd.BusinessInterests = append(bd.BusinessInterests, businessInterest)
}

func (bd *BankingDetails) AddOtherAsset(asset OtherAsset) {
	bd.OtherAssets = append(bd.OtherAssets, asset)
}
func (bd *BankingDetails) AddOtherLiability(liability OtherLiability) {
	bd.OtherLiabilities = append(bd.OtherLiabilities, liability)
}
func (bd *BankingDetails) AddOtherIncome(income OtherIncome) {
	bd.OtherIncome = append(bd.OtherIncome, income)
}
func (bd *BankingDetails) AddOtherExpense(expense OtherExpense) {
	bd.OtherExpenses = append(bd.OtherExpenses, expense)
}
func (bd *BankingDetails) AddOtherDebt(debt OtherDebt) {
	bd.OtherDebts = append(bd.OtherDebts, debt)
}
func (bd *BankingDetails) GetTransactionHistory() []Transaction {
	return bd.TransactionHistory
}
func (bd *BankingDetails) GetInvestments() []Investment {
	return bd.Investments
}
func (bd *BankingDetails) GetLoans() []Loan {
	return bd.Loans
}
func (bd *BankingDetails) GetInsurancePolicies() []InsurancePolicy {
	return bd.InsurancePolicies
}
func (bd *BankingDetails) GetRealEstate() []RealEstate {
	return bd.RealEstate
}
func (bd *BankingDetails) GetBusinessInterests() []BusinessInterest {
	return bd.BusinessInterests
}
func (bd *BankingDetails) GetOtherAssets() []OtherAsset {
	return bd.OtherAssets
}
func (bd *BankingDetails) GetOtherLiabilities() []OtherLiability {
	return bd.OtherLiabilities
}
func (bd *BankingDetails) GetOtherIncome() []OtherIncome {
	return bd.OtherIncome
}
func (bd *BankingDetails) GetOtherExpenses() []OtherExpense {
	return bd.OtherExpenses
}
func (bd *BankingDetails) GetOtherDebts() []OtherDebt {
	return bd.OtherDebts
}
func (bd *BankingDetails) GetAccountNumber() string {
	return bd.AccountNumber
}
func (bd *BankingDetails) GetBankName() string {
	return bd.BankName
}
func (bd *BankingDetails) GetIBAN() string {
	return bd.IBAN
}
func (bd *BankingDetails) GetSwiftCode() string {
	return bd.SwiftCode
}
func (bd *BankingDetails) GetBalance() string {
	return bd.Balance
}
