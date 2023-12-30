// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package nft_candy_machine

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/RoboticAgile/solana-go"
	ag_format "github.com/RoboticAgile/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// EditCmV3 is the `editCmV3` instruction.
type EditCmV3 struct {
	Data *CandyMachineDataV2

	// [0] = [SIGNER] authority
	//
	// [1] = [WRITE] candyMachine
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewEditCmV3InstructionBuilder creates a new `EditCmV3` instruction builder.
func NewEditCmV3InstructionBuilder() *EditCmV3 {
	nd := &EditCmV3{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 2),
	}
	return nd
}

// SetData sets the "data" parameter.
func (inst *EditCmV3) SetData(data CandyMachineDataV2) *EditCmV3 {
	inst.Data = &data
	return inst
}

// SetAuthorityAccount sets the "authority" account.
func (inst *EditCmV3) SetAuthorityAccount(authority ag_solanago.PublicKey) *EditCmV3 {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *EditCmV3) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetCandyMachineAccount sets the "candyMachine" account.
func (inst *EditCmV3) SetCandyMachineAccount(candyMachine ag_solanago.PublicKey) *EditCmV3 {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(candyMachine).WRITE()
	return inst
}

// GetCandyMachineAccount gets the "candyMachine" account.
func (inst *EditCmV3) GetCandyMachineAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

func (inst EditCmV3) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_EditCmV3,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst EditCmV3) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *EditCmV3) Validate() error {
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
	}
	return nil
}

func (inst *EditCmV3) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("EditCmV3")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Data", *inst.Data))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=2]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("   authority", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("candyMachine", inst.AccountMetaSlice.Get(1)))
					})
				})
		})
}

func (obj EditCmV3) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Data` param:
	err = encoder.Encode(obj.Data)
	if err != nil {
		return err
	}
	return nil
}
func (obj *EditCmV3) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Data`:
	err = decoder.Decode(&obj.Data)
	if err != nil {
		return err
	}
	return nil
}

// NewEditCmV3Instruction declares a new EditCmV3 instruction with the provided parameters and accounts.
func NewEditCmV3Instruction(
	// Parameters:
	data CandyMachineDataV2,
	// Accounts:
	authority ag_solanago.PublicKey,
	candyMachine ag_solanago.PublicKey) *EditCmV3 {
	return NewEditCmV3InstructionBuilder().
		SetData(data).
		SetAuthorityAccount(authority).
		SetCandyMachineAccount(candyMachine)
}
