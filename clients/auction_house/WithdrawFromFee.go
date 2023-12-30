// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package auction_house

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/RoboticAgile/solana-go"
	ag_format "github.com/RoboticAgile/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// WithdrawFromFee is the `withdrawFromFee` instruction.
type WithdrawFromFee struct {
	Amount *uint64

	// [0] = [SIGNER] authority
	//
	// [1] = [WRITE] feeWithdrawalDestination
	//
	// [2] = [WRITE] auctionHouseFeeAccount
	//
	// [3] = [WRITE] auctionHouse
	//
	// [4] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewWithdrawFromFeeInstructionBuilder creates a new `WithdrawFromFee` instruction builder.
func NewWithdrawFromFeeInstructionBuilder() *WithdrawFromFee {
	nd := &WithdrawFromFee{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 5),
	}
	return nd
}

// SetAmount sets the "amount" parameter.
func (inst *WithdrawFromFee) SetAmount(amount uint64) *WithdrawFromFee {
	inst.Amount = &amount
	return inst
}

// SetAuthorityAccount sets the "authority" account.
func (inst *WithdrawFromFee) SetAuthorityAccount(authority ag_solanago.PublicKey) *WithdrawFromFee {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *WithdrawFromFee) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetFeeWithdrawalDestinationAccount sets the "feeWithdrawalDestination" account.
func (inst *WithdrawFromFee) SetFeeWithdrawalDestinationAccount(feeWithdrawalDestination ag_solanago.PublicKey) *WithdrawFromFee {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(feeWithdrawalDestination).WRITE()
	return inst
}

// GetFeeWithdrawalDestinationAccount gets the "feeWithdrawalDestination" account.
func (inst *WithdrawFromFee) GetFeeWithdrawalDestinationAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetAuctionHouseFeeAccountAccount sets the "auctionHouseFeeAccount" account.
func (inst *WithdrawFromFee) SetAuctionHouseFeeAccountAccount(auctionHouseFeeAccount ag_solanago.PublicKey) *WithdrawFromFee {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(auctionHouseFeeAccount).WRITE()
	return inst
}

// GetAuctionHouseFeeAccountAccount gets the "auctionHouseFeeAccount" account.
func (inst *WithdrawFromFee) GetAuctionHouseFeeAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetAuctionHouseAccount sets the "auctionHouse" account.
func (inst *WithdrawFromFee) SetAuctionHouseAccount(auctionHouse ag_solanago.PublicKey) *WithdrawFromFee {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(auctionHouse).WRITE()
	return inst
}

// GetAuctionHouseAccount gets the "auctionHouse" account.
func (inst *WithdrawFromFee) GetAuctionHouseAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *WithdrawFromFee) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *WithdrawFromFee {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *WithdrawFromFee) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

func (inst WithdrawFromFee) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_WithdrawFromFee,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst WithdrawFromFee) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *WithdrawFromFee) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Amount == nil {
			return errors.New("Amount parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.FeeWithdrawalDestination is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.AuctionHouseFeeAccount is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.AuctionHouse is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *WithdrawFromFee) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("WithdrawFromFee")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Amount", *inst.Amount))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=5]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("               authority", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("feeWithdrawalDestination", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("         auctionHouseFee", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("            auctionHouse", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("           systemProgram", inst.AccountMetaSlice.Get(4)))
					})
				})
		})
}

func (obj WithdrawFromFee) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Amount` param:
	err = encoder.Encode(obj.Amount)
	if err != nil {
		return err
	}
	return nil
}
func (obj *WithdrawFromFee) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Amount`:
	err = decoder.Decode(&obj.Amount)
	if err != nil {
		return err
	}
	return nil
}

// NewWithdrawFromFeeInstruction declares a new WithdrawFromFee instruction with the provided parameters and accounts.
func NewWithdrawFromFeeInstruction(
	// Parameters:
	amount uint64,
	// Accounts:
	authority ag_solanago.PublicKey,
	feeWithdrawalDestination ag_solanago.PublicKey,
	auctionHouseFeeAccount ag_solanago.PublicKey,
	auctionHouse ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *WithdrawFromFee {
	return NewWithdrawFromFeeInstructionBuilder().
		SetAmount(amount).
		SetAuthorityAccount(authority).
		SetFeeWithdrawalDestinationAccount(feeWithdrawalDestination).
		SetAuctionHouseFeeAccountAccount(auctionHouseFeeAccount).
		SetAuctionHouseAccount(auctionHouse).
		SetSystemProgramAccount(systemProgram)
}
