// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package nft_candy_machine

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/RoboticAgile/solana-go"
	ag_format "github.com/RoboticAgile/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// UpdateCandyMachine is the `updateCandyMachine` instruction.
type UpdateCandyMachine struct {
	Price      *uint64 `bin:"optional"`
	GoLiveDate *int64  `bin:"optional"`

	// [0] = [WRITE] candyMachine
	//
	// [1] = [SIGNER] authority
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewUpdateCandyMachineInstructionBuilder creates a new `UpdateCandyMachine` instruction builder.
func NewUpdateCandyMachineInstructionBuilder() *UpdateCandyMachine {
	nd := &UpdateCandyMachine{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 2),
	}
	return nd
}

// SetPrice sets the "price" parameter.
func (inst *UpdateCandyMachine) SetPrice(price uint64) *UpdateCandyMachine {
	inst.Price = &price
	return inst
}

// SetGoLiveDate sets the "goLiveDate" parameter.
func (inst *UpdateCandyMachine) SetGoLiveDate(goLiveDate int64) *UpdateCandyMachine {
	inst.GoLiveDate = &goLiveDate
	return inst
}

// SetCandyMachineAccount sets the "candyMachine" account.
func (inst *UpdateCandyMachine) SetCandyMachineAccount(candyMachine ag_solanago.PublicKey) *UpdateCandyMachine {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(candyMachine).WRITE()
	return inst
}

// GetCandyMachineAccount gets the "candyMachine" account.
func (inst *UpdateCandyMachine) GetCandyMachineAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetAuthorityAccount sets the "authority" account.
func (inst *UpdateCandyMachine) SetAuthorityAccount(authority ag_solanago.PublicKey) *UpdateCandyMachine {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *UpdateCandyMachine) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

func (inst UpdateCandyMachine) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_UpdateCandyMachine,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UpdateCandyMachine) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UpdateCandyMachine) Validate() error {
	// Check whether all (required) parameters are set:
	{
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.CandyMachine is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Authority is not set")
		}
	}
	return nil
}

func (inst *UpdateCandyMachine) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UpdateCandyMachine")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("     Price (OPT)", inst.Price))
						paramsBranch.Child(ag_format.Param("GoLiveDate (OPT)", inst.GoLiveDate))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=2]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("candyMachine", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("   authority", inst.AccountMetaSlice[1]))
					})
				})
		})
}

func (obj UpdateCandyMachine) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Price` param (optional):
	{
		if obj.Price == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.Price)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `GoLiveDate` param (optional):
	{
		if obj.GoLiveDate == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.GoLiveDate)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (obj *UpdateCandyMachine) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Price` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.Price)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `GoLiveDate` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.GoLiveDate)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// NewUpdateCandyMachineInstruction declares a new UpdateCandyMachine instruction with the provided parameters and accounts.
func NewUpdateCandyMachineInstruction(
	// Parameters:
	price uint64,
	goLiveDate int64,
	// Accounts:
	candyMachine ag_solanago.PublicKey,
	authority ag_solanago.PublicKey) *UpdateCandyMachine {
	return NewUpdateCandyMachineInstructionBuilder().
		SetPrice(price).
		SetGoLiveDate(goLiveDate).
		SetCandyMachineAccount(candyMachine).
		SetAuthorityAccount(authority)
}
