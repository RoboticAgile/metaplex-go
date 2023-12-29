// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package nft_candy_machine

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/desperatee/solana-go"
	ag_format "github.com/desperatee/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// RevealV5 is the `revealV5` instruction.
type RevealV5 struct {

	// [0] = [WRITE] candyMachine
	//
	// [1] = [WRITE] metadata
	//
	// [2] = [] tokenMetadataProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewRevealV5InstructionBuilder creates a new `RevealV5` instruction builder.
func NewRevealV5InstructionBuilder() *RevealV5 {
	nd := &RevealV5{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetCandyMachineAccount sets the "candyMachine" account.
func (inst *RevealV5) SetCandyMachineAccount(candyMachine ag_solanago.PublicKey) *RevealV5 {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(candyMachine).WRITE()
	return inst
}

// GetCandyMachineAccount gets the "candyMachine" account.
func (inst *RevealV5) GetCandyMachineAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetMetadataAccount sets the "metadata" account.
func (inst *RevealV5) SetMetadataAccount(metadata ag_solanago.PublicKey) *RevealV5 {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(metadata).WRITE()
	return inst
}

// GetMetadataAccount gets the "metadata" account.
func (inst *RevealV5) GetMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetTokenMetadataProgramAccount sets the "tokenMetadataProgram" account.
func (inst *RevealV5) SetTokenMetadataProgramAccount(tokenMetadataProgram ag_solanago.PublicKey) *RevealV5 {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(tokenMetadataProgram)
	return inst
}

// GetTokenMetadataProgramAccount gets the "tokenMetadataProgram" account.
func (inst *RevealV5) GetTokenMetadataProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

func (inst RevealV5) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_RevealV5,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst RevealV5) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *RevealV5) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.CandyMachine is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Metadata is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.TokenMetadataProgram is not set")
		}
	}
	return nil
}

func (inst *RevealV5) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("RevealV5")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("        candyMachine", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("            metadata", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("tokenMetadataProgram", inst.AccountMetaSlice.Get(2)))
					})
				})
		})
}

func (obj RevealV5) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *RevealV5) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewRevealV5Instruction declares a new RevealV5 instruction with the provided parameters and accounts.
func NewRevealV5Instruction(
	// Accounts:
	candyMachine ag_solanago.PublicKey,
	metadata ag_solanago.PublicKey,
	tokenMetadataProgram ag_solanago.PublicKey) *RevealV5 {
	return NewRevealV5InstructionBuilder().
		SetCandyMachineAccount(candyMachine).
		SetMetadataAccount(metadata).
		SetTokenMetadataProgramAccount(tokenMetadataProgram)
}
