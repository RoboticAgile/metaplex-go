// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package candy_machine

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/RoboticAgile/solana-go"
	ag_format "github.com/RoboticAgile/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// ClosePermissionedSettings is the `closePermissionedSettings` instruction.
type ClosePermissionedSettings struct {

	// [0] = [WRITE] candyMachine
	//
	// [1] = [SIGNER] authority
	//
	// [2] = [WRITE] permissionedSettings
	//
	// [3] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewClosePermissionedSettingsInstructionBuilder creates a new `ClosePermissionedSettings` instruction builder.
func NewClosePermissionedSettingsInstructionBuilder() *ClosePermissionedSettings {
	nd := &ClosePermissionedSettings{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 4),
	}
	return nd
}

// SetCandyMachineAccount sets the "candyMachine" account.
func (inst *ClosePermissionedSettings) SetCandyMachineAccount(candyMachine ag_solanago.PublicKey) *ClosePermissionedSettings {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(candyMachine).WRITE()
	return inst
}

// GetCandyMachineAccount gets the "candyMachine" account.
func (inst *ClosePermissionedSettings) GetCandyMachineAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAuthorityAccount sets the "authority" account.
func (inst *ClosePermissionedSettings) SetAuthorityAccount(authority ag_solanago.PublicKey) *ClosePermissionedSettings {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *ClosePermissionedSettings) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetPermissionedSettingsAccount sets the "permissionedSettings" account.
func (inst *ClosePermissionedSettings) SetPermissionedSettingsAccount(permissionedSettings ag_solanago.PublicKey) *ClosePermissionedSettings {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(permissionedSettings).WRITE()
	return inst
}

// GetPermissionedSettingsAccount gets the "permissionedSettings" account.
func (inst *ClosePermissionedSettings) GetPermissionedSettingsAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *ClosePermissionedSettings) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *ClosePermissionedSettings {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *ClosePermissionedSettings) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

func (inst ClosePermissionedSettings) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_ClosePermissionedSettings,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst ClosePermissionedSettings) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *ClosePermissionedSettings) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.CandyMachine is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.PermissionedSettings is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *ClosePermissionedSettings) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("ClosePermissionedSettings")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=4]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("        candyMachine", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("           authority", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("permissionedSettings", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("       systemProgram", inst.AccountMetaSlice.Get(3)))
					})
				})
		})
}

func (obj ClosePermissionedSettings) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *ClosePermissionedSettings) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewClosePermissionedSettingsInstruction declares a new ClosePermissionedSettings instruction with the provided parameters and accounts.
func NewClosePermissionedSettingsInstruction(
	// Accounts:
	candyMachine ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	permissionedSettings ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *ClosePermissionedSettings {
	return NewClosePermissionedSettingsInstructionBuilder().
		SetCandyMachineAccount(candyMachine).
		SetAuthorityAccount(authority).
		SetPermissionedSettingsAccount(permissionedSettings).
		SetSystemProgramAccount(systemProgram)
}
