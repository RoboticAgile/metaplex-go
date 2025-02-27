// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package launchpad

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/RoboticAgile/solana-go"
	ag_format "github.com/RoboticAgile/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// UpdateStage is the `updateStage` instruction.
type UpdateStage struct {
	Params *UpdateStageParams

	// [0] = [WRITE, SIGNER] manager
	//
	// [1] = [WRITE] collection
	//
	// [2] = [WRITE] stage
	//
	// [3] = [] systemProgram
	//
	// [4] = [] clock
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewUpdateStageInstructionBuilder creates a new `UpdateStage` instruction builder.
func NewUpdateStageInstructionBuilder() *UpdateStage {
	nd := &UpdateStage{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 5),
	}
	return nd
}

// SetParams sets the "params" parameter.
func (inst *UpdateStage) SetParams(params UpdateStageParams) *UpdateStage {
	inst.Params = &params
	return inst
}

// SetManagerAccount sets the "manager" account.
func (inst *UpdateStage) SetManagerAccount(manager ag_solanago.PublicKey) *UpdateStage {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(manager).WRITE().SIGNER()
	return inst
}

// GetManagerAccount gets the "manager" account.
func (inst *UpdateStage) GetManagerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetCollectionAccount sets the "collection" account.
func (inst *UpdateStage) SetCollectionAccount(collection ag_solanago.PublicKey) *UpdateStage {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(collection).WRITE()
	return inst
}

// GetCollectionAccount gets the "collection" account.
func (inst *UpdateStage) GetCollectionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetStageAccount sets the "stage" account.
func (inst *UpdateStage) SetStageAccount(stage ag_solanago.PublicKey) *UpdateStage {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(stage).WRITE()
	return inst
}

// GetStageAccount gets the "stage" account.
func (inst *UpdateStage) GetStageAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *UpdateStage) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *UpdateStage {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *UpdateStage) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetClockAccount sets the "clock" account.
func (inst *UpdateStage) SetClockAccount(clock ag_solanago.PublicKey) *UpdateStage {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(clock)
	return inst
}

// GetClockAccount gets the "clock" account.
func (inst *UpdateStage) GetClockAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

func (inst UpdateStage) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_UpdateStage,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UpdateStage) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UpdateStage) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Params == nil {
			return errors.New("Params parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Manager is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Collection is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Stage is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Clock is not set")
		}
	}
	return nil
}

func (inst *UpdateStage) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UpdateStage")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Params", *inst.Params))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=5]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("      manager", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("   collection", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("        stage", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("systemProgram", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("        clock", inst.AccountMetaSlice.Get(4)))
					})
				})
		})
}

func (obj UpdateStage) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Params` param:
	err = encoder.Encode(obj.Params)
	if err != nil {
		return err
	}
	return nil
}
func (obj *UpdateStage) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Params`:
	err = decoder.Decode(&obj.Params)
	if err != nil {
		return err
	}
	return nil
}

// NewUpdateStageInstruction declares a new UpdateStage instruction with the provided parameters and accounts.
func NewUpdateStageInstruction(
	// Parameters:
	params UpdateStageParams,
	// Accounts:
	manager ag_solanago.PublicKey,
	collection ag_solanago.PublicKey,
	stage ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	clock ag_solanago.PublicKey) *UpdateStage {
	return NewUpdateStageInstructionBuilder().
		SetParams(params).
		SetManagerAccount(manager).
		SetCollectionAccount(collection).
		SetStageAccount(stage).
		SetSystemProgramAccount(systemProgram).
		SetClockAccount(clock)
}
