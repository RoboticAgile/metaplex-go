// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package launchpad

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/desperatee/solana-go"
	ag_format "github.com/desperatee/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// InitStage is the `initStage` instruction.
type InitStage struct {
	Params *InitStageParams

	// [0] = [WRITE, SIGNER] manager
	//
	// [1] = [WRITE] collection
	//
	// [2] = [WRITE] stage
	//
	// [3] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewInitStageInstructionBuilder creates a new `InitStage` instruction builder.
func NewInitStageInstructionBuilder() *InitStage {
	nd := &InitStage{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 4),
	}
	return nd
}

// SetParams sets the "params" parameter.
func (inst *InitStage) SetParams(params InitStageParams) *InitStage {
	inst.Params = &params
	return inst
}

// SetManagerAccount sets the "manager" account.
func (inst *InitStage) SetManagerAccount(manager ag_solanago.PublicKey) *InitStage {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(manager).WRITE().SIGNER()
	return inst
}

// GetManagerAccount gets the "manager" account.
func (inst *InitStage) GetManagerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetCollectionAccount sets the "collection" account.
func (inst *InitStage) SetCollectionAccount(collection ag_solanago.PublicKey) *InitStage {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(collection).WRITE()
	return inst
}

// GetCollectionAccount gets the "collection" account.
func (inst *InitStage) GetCollectionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetStageAccount sets the "stage" account.
func (inst *InitStage) SetStageAccount(stage ag_solanago.PublicKey) *InitStage {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(stage).WRITE()
	return inst
}

// GetStageAccount gets the "stage" account.
func (inst *InitStage) GetStageAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *InitStage) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *InitStage {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *InitStage) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

func (inst InitStage) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_InitStage,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst InitStage) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *InitStage) Validate() error {
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
	}
	return nil
}

func (inst *InitStage) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("InitStage")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Params", *inst.Params))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=4]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("      manager", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("   collection", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("        stage", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("systemProgram", inst.AccountMetaSlice.Get(3)))
					})
				})
		})
}

func (obj InitStage) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Params` param:
	err = encoder.Encode(obj.Params)
	if err != nil {
		return err
	}
	return nil
}
func (obj *InitStage) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Params`:
	err = decoder.Decode(&obj.Params)
	if err != nil {
		return err
	}
	return nil
}

// NewInitStageInstruction declares a new InitStage instruction with the provided parameters and accounts.
func NewInitStageInstruction(
	// Parameters:
	params InitStageParams,
	// Accounts:
	manager ag_solanago.PublicKey,
	collection ag_solanago.PublicKey,
	stage ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *InitStage {
	return NewInitStageInstructionBuilder().
		SetParams(params).
		SetManagerAccount(manager).
		SetCollectionAccount(collection).
		SetStageAccount(stage).
		SetSystemProgramAccount(systemProgram)
}
