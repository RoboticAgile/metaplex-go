// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package nft_candy_machine

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/desperatee/solana-go"
	ag_format "github.com/desperatee/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// EditCmV5 is the `editCmV5` instruction.
type EditCmV5 struct {
	Data *CandyMachineDataV3

	// [0] = [SIGNER] authority
	//
	// [1] = [WRITE] candyMachine
	//
	// [2] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewEditCmV5InstructionBuilder creates a new `EditCmV5` instruction builder.
func NewEditCmV5InstructionBuilder() *EditCmV5 {
	nd := &EditCmV5{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetData sets the "data" parameter.
func (inst *EditCmV5) SetData(data CandyMachineDataV3) *EditCmV5 {
	inst.Data = &data
	return inst
}

// SetAuthorityAccount sets the "authority" account.
func (inst *EditCmV5) SetAuthorityAccount(authority ag_solanago.PublicKey) *EditCmV5 {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *EditCmV5) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetCandyMachineAccount sets the "candyMachine" account.
func (inst *EditCmV5) SetCandyMachineAccount(candyMachine ag_solanago.PublicKey) *EditCmV5 {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(candyMachine).WRITE()
	return inst
}

// GetCandyMachineAccount gets the "candyMachine" account.
func (inst *EditCmV5) GetCandyMachineAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *EditCmV5) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *EditCmV5 {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *EditCmV5) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

func (inst EditCmV5) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_EditCmV5,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst EditCmV5) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *EditCmV5) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Data == nil {
			return errors.New("Data parameter is not set")
		}
	}

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

func (inst *EditCmV5) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("EditCmV5")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Data", *inst.Data))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("    authority", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta(" candyMachine", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("systemProgram", inst.AccountMetaSlice.Get(2)))
					})
				})
		})
}

func (obj EditCmV5) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Data` param:
	err = encoder.Encode(obj.Data)
	if err != nil {
		return err
	}
	return nil
}
func (obj *EditCmV5) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Data`:
	err = decoder.Decode(&obj.Data)
	if err != nil {
		return err
	}
	return nil
}

// NewEditCmV5Instruction declares a new EditCmV5 instruction with the provided parameters and accounts.
func NewEditCmV5Instruction(
	// Parameters:
	data CandyMachineDataV3,
	// Accounts:
	authority ag_solanago.PublicKey,
	candyMachine ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *EditCmV5 {
	return NewEditCmV5InstructionBuilder().
		SetData(data).
		SetAuthorityAccount(authority).
		SetCandyMachineAccount(candyMachine).
		SetSystemProgramAccount(systemProgram)
}
