// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package nft_candy_machine

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/RoboticAgile/solana-go"
	ag_format "github.com/RoboticAgile/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// AllowUnfreezeV5 is the `allowUnfreezeV5` instruction.
type AllowUnfreezeV5 struct {

	// [0] = [SIGNER] authority
	//
	// [1] = [WRITE] candyMachine
	//
	// [2] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewAllowUnfreezeV5InstructionBuilder creates a new `AllowUnfreezeV5` instruction builder.
func NewAllowUnfreezeV5InstructionBuilder() *AllowUnfreezeV5 {
	nd := &AllowUnfreezeV5{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetAuthorityAccount sets the "authority" account.
func (inst *AllowUnfreezeV5) SetAuthorityAccount(authority ag_solanago.PublicKey) *AllowUnfreezeV5 {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *AllowUnfreezeV5) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetCandyMachineAccount sets the "candyMachine" account.
func (inst *AllowUnfreezeV5) SetCandyMachineAccount(candyMachine ag_solanago.PublicKey) *AllowUnfreezeV5 {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(candyMachine).WRITE()
	return inst
}

// GetCandyMachineAccount gets the "candyMachine" account.
func (inst *AllowUnfreezeV5) GetCandyMachineAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *AllowUnfreezeV5) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *AllowUnfreezeV5 {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *AllowUnfreezeV5) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

func (inst AllowUnfreezeV5) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_AllowUnfreezeV5,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst AllowUnfreezeV5) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *AllowUnfreezeV5) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.CandyMachine is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *AllowUnfreezeV5) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("AllowUnfreezeV5")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("    authority", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta(" candyMachine", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("systemProgram", inst.AccountMetaSlice.Get(2)))
					})
				})
		})
}

func (obj AllowUnfreezeV5) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *AllowUnfreezeV5) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewAllowUnfreezeV5Instruction declares a new AllowUnfreezeV5 instruction with the provided parameters and accounts.
func NewAllowUnfreezeV5Instruction(
	// Accounts:
	authority ag_solanago.PublicKey,
	candyMachine ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *AllowUnfreezeV5 {
	return NewAllowUnfreezeV5InstructionBuilder().
		SetAuthorityAccount(authority).
		SetCandyMachineAccount(candyMachine).
		SetSystemProgramAccount(systemProgram)
}
