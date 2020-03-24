package entity

// Define a beneficiary type
// @Agency the number of your agency
// @Account the number of your account
// @Wallet the number of your wallet
// @TypeAccount
type Beneficiary struct {
	Agency      string
	Account     string
	TypeAccount string
	Agreement   string
	Wallet      string
}
