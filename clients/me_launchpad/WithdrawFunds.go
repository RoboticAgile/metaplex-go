// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package me_launchpad

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/RoboticAgile/solana-go"
	ag_format "github.com/RoboticAgile/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// WithdrawFunds is the `withdrawFunds` instruction.
type WithdrawFunds struct {

	// [0] = [WRITE] config
	//
	// [1] = [SIGNER] authority
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewWithdrawFundsInstructionBuilder creates a new `WithdrawFunds` instruction builder.
func NewWithdrawFundsInstructionBuilder() *WithdrawFunds {
	nd := &WithdrawFunds{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 2),
	}
	return nd
}

// SetConfigAccount sets the "config" account.
func (inst *WithdrawFunds) SetConfigAccount(config ag_solanago.PublicKey) *WithdrawFunds {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(config).WRITE()
	return inst
}

// GetConfigAccount gets the "config" account.
func (inst *WithdrawFunds) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAuthorityAccount sets the "authority" account.
func (inst *WithdrawFunds) SetAuthorityAccount(authority ag_solanago.PublicKey) *WithdrawFunds {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *WithdrawFunds) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

func (inst WithdrawFunds) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_WithdrawFunds,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst WithdrawFunds) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *WithdrawFunds) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Config is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Authority is not set")
		}
	}
	return nil
}

func (inst *WithdrawFunds) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("WithdrawFunds")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=2]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("   config", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("authority", inst.AccountMetaSlice.Get(1)))
					})
				})
		})
}

func (obj WithdrawFunds) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *WithdrawFunds) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewWithdrawFundsInstruction declares a new WithdrawFunds instruction with the provided parameters and accounts.
func NewWithdrawFundsInstruction(
	// Accounts:
	config ag_solanago.PublicKey,
	authority ag_solanago.PublicKey) *WithdrawFunds {
	return NewWithdrawFundsInstructionBuilder().
		SetConfigAccount(config).
		SetAuthorityAccount(authority)
}
