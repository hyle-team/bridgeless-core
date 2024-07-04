package vesting

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"time"
)

type Coin struct {
	Denom  string `json:"denom"`
	Amount string `json:"amount"`
}

type BaseAccount struct {
	Address       string      `json:"address"`
	PubKey        interface{} `json:"pub_key"`
	AccountNumber string      `json:"account_number"`
	Sequence      string      `json:"sequence"`
}

type BaseVestingAccount struct {
	BaseAccount      BaseAccount `json:"base_account"`
	OriginalVesting  []Coin      `json:"original_vesting"`
	DelegatedFree    []Coin      `json:"delegated_free"`
	DelegatedVesting []Coin      `json:"delegated_vesting"`
	EndTime          int64       `json:"end_time"`
}

type VestingPeriod struct {
	Length int64  `json:"length"`
	Amount []Coin `json:"amount"`
}

type ContinuousVestingAccount struct {
	BaseVestingAccount BaseVestingAccount `json:"base_vesting_account"`
	StartTime          int64              `json:"start_time"`
	VestingPeriods     []VestingPeriod    `json:"vesting_periods"`
}

func main() {
	// Define the base account and vesting parameters
	baseAccount := BaseAccount{
		Address:       "evmos1c3jr9xa7mzlujm2fedvzkcg34nsn5lm3f73w5h", // Replace with actual address
		PubKey:        nil,
		AccountNumber: "2",
		Sequence:      "0",
	}

	originalVestingAmount := big.NewInt(1000000000000000000) // 1 aevmos
	originalVesting := []Coin{
		{
			Denom:  "aevmos",
			Amount: originalVestingAmount.String(), // 1 aevmos
		},
	}

	// Define vesting schedule parameters
	startTime := time.Unix(167253000, 0)           // Replace with actual start time
	endTime := startTime.AddDate(1, 2, 234).Unix() // One year later
	numberOfPeriods := 123
	periodLength := int64(2592000) // 30 days in seconds

	// Calculate amount per period
	amountPerPeriod := new(big.Int).Div(originalVestingAmount, big.NewInt(int64(numberOfPeriods)))

	// Create vesting periods
	var vestingPeriods []VestingPeriod
	for i := 0; i < numberOfPeriods; i++ {
		vestingPeriods = append(vestingPeriods, VestingPeriod{
			Length: periodLength,
			Amount: []Coin{
				{
					Denom:  "aevmos",
					Amount: amountPerPeriod.String(),
				},
			},
		})
	}

	// Create the continuous vesting account
	continuousVestingAccount := ContinuousVestingAccount{
		BaseVestingAccount: BaseVestingAccount{
			BaseAccount:      baseAccount,
			OriginalVesting:  originalVesting,
			DelegatedFree:    []Coin{},
			DelegatedVesting: []Coin{},
			EndTime:          endTime,
		},
		StartTime:      startTime.Unix(),
		VestingPeriods: vestingPeriods,
	}

	// Convert to JSON
	data, err := json.MarshalIndent(continuousVestingAccount, "", "  ")
	if err != nil {
		fmt.Printf("Error marshalling JSON: %v\n", err)
		return
	}

	// Save to file
	err = ioutil.WriteFile("continuous_vesting.json", data, 0644)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		return
	}

	fmt.Println("Vesting periods file created successfully")
}
