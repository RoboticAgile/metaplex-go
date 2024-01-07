// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package me_launchpad

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/RoboticAgile/solana-go"
	ag_format "github.com/RoboticAgile/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// InitializeConfig is the `initializeConfig` instruction.
type InitializeConfig struct {
	Args *InitializeConfigArgs

	// [0] = [WRITE, SIGNER] config
	//
	// [1] = [] authority
	//
	// [2] = [WRITE, SIGNER] payer
	//
	// [3] = [] rent
	//
	// [4] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewInitializeConfigInstructionBuilder creates a new `InitializeConfig` instruction builder.
func NewInitializeConfigInstructionBuilder() *InitializeConfig {
	nd := &InitializeConfig{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 5),
	}
	return nd
}

// SetArgs sets the "args" parameter.
func (inst *InitializeConfig) SetArgs(args InitializeConfigArgs) *InitializeConfig {
	inst.Args = &args
	return inst
}

// SetConfigAccount sets the "config" account.
func (inst *InitializeConfig) SetConfigAccount(config ag_solanago.PublicKey) *InitializeConfig {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(config).WRITE().SIGNER()
	return inst
}

// GetConfigAccount gets the "config" account.
func (inst *InitializeConfig) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAuthorityAccount sets the "authority" account.
func (inst *InitializeConfig) SetAuthorityAccount(authority ag_solanago.PublicKey) *InitializeConfig {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(authority)
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *InitializeConfig) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetPayerAccount sets the "payer" account.
func (inst *InitializeConfig) SetPayerAccount(payer ag_solanago.PublicKey) *InitializeConfig {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(payer).WRITE().SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
func (inst *InitializeConfig) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetRentAccount sets the "rent" account.
func (inst *InitializeConfig) SetRentAccount(rent ag_solanago.PublicKey) *InitializeConfig {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *InitializeConfig) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *InitializeConfig) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *InitializeConfig {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *InitializeConfig) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

func (inst InitializeConfig) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_InitializeConfig,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst InitializeConfig) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *InitializeConfig) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Args == nil {
			return errors.New("Args parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Config is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Rent is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *InitializeConfig) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("InitializeConfig")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Args", *inst.Args))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=5]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("       config", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("    authority", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("        payer", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("         rent", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("systemProgram", inst.AccountMetaSlice.Get(4)))
					})
				})
		})
}

func (obj InitializeConfig) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Args` param:
	err = encoder.Encode(obj.Args)
	if err != nil {
		return err
	}
	return nil
}
func (obj *InitializeConfig) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Args`:
	err = decoder.Decode(&obj.Args)
	if err != nil {
		return err
	}
	return nil
}

// NewInitializeConfigInstruction declares a new InitializeConfig instruction with the provided parameters and accounts.
func NewInitializeConfigInstruction(
	// Parameters:
	args InitializeConfigArgs,
	// Accounts:
	config ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	rent ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *InitializeConfig {
	return NewInitializeConfigInstructionBuilder().
		SetArgs(args).
		SetConfigAccount(config).
		SetAuthorityAccount(authority).
		SetPayerAccount(payer).
		SetRentAccount(rent).
		SetSystemProgramAccount(systemProgram)
}
