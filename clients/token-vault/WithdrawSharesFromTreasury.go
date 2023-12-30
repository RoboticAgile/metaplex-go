// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package token_vault

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/RoboticAgile/solana-go"
	ag_format "github.com/RoboticAgile/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Withdraws shares from the treasury to a desired account.
type WithdrawSharesFromTreasury struct {
	Args *NumberOfShareArgs

	// [0] = [WRITE] initializedDestination
	// ··········· Initialized Destination account for the shares being withdrawn
	//
	// [1] = [WRITE] fractionTreasury
	// ··········· Fraction treasury
	//
	// [2] = [] initializedActiveTokenVault
	// ··········· The initialized active token vault
	//
	// [3] = [] pdaBasedTransferAuthority
	// ··········· PDA-based Transfer authority to move tokens from treasury to your destination[PREFIX, program_id]
	//
	// [4] = [SIGNER] vaultAuthority
	// ··········· Authority of vault
	//
	// [5] = [] tokenProgram
	// ··········· Token program
	//
	// [6] = [] rentSysvar
	// ··········· Rent sysvar
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewWithdrawSharesFromTreasuryInstructionBuilder creates a new `WithdrawSharesFromTreasury` instruction builder.
func NewWithdrawSharesFromTreasuryInstructionBuilder() *WithdrawSharesFromTreasury {
	nd := &WithdrawSharesFromTreasury{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 7),
	}
	return nd
}

// SetArgs sets the "args" parameter.
func (inst *WithdrawSharesFromTreasury) SetArgs(args NumberOfShareArgs) *WithdrawSharesFromTreasury {
	inst.Args = &args
	return inst
}

// SetInitializedDestinationAccount sets the "initializedDestination" account.
// Initialized Destination account for the shares being withdrawn
func (inst *WithdrawSharesFromTreasury) SetInitializedDestinationAccount(initializedDestination ag_solanago.PublicKey) *WithdrawSharesFromTreasury {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(initializedDestination).WRITE()
	return inst
}

// GetInitializedDestinationAccount gets the "initializedDestination" account.
// Initialized Destination account for the shares being withdrawn
func (inst *WithdrawSharesFromTreasury) GetInitializedDestinationAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetFractionTreasuryAccount sets the "fractionTreasury" account.
// Fraction treasury
func (inst *WithdrawSharesFromTreasury) SetFractionTreasuryAccount(fractionTreasury ag_solanago.PublicKey) *WithdrawSharesFromTreasury {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(fractionTreasury).WRITE()
	return inst
}

// GetFractionTreasuryAccount gets the "fractionTreasury" account.
// Fraction treasury
func (inst *WithdrawSharesFromTreasury) GetFractionTreasuryAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetInitializedActiveTokenVaultAccount sets the "initializedActiveTokenVault" account.
// The initialized active token vault
func (inst *WithdrawSharesFromTreasury) SetInitializedActiveTokenVaultAccount(initializedActiveTokenVault ag_solanago.PublicKey) *WithdrawSharesFromTreasury {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(initializedActiveTokenVault)
	return inst
}

// GetInitializedActiveTokenVaultAccount gets the "initializedActiveTokenVault" account.
// The initialized active token vault
func (inst *WithdrawSharesFromTreasury) GetInitializedActiveTokenVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetPdaBasedTransferAuthorityAccount sets the "pdaBasedTransferAuthority" account.
// PDA-based Transfer authority to move tokens from treasury to your destination[PREFIX, program_id]
func (inst *WithdrawSharesFromTreasury) SetPdaBasedTransferAuthorityAccount(pdaBasedTransferAuthority ag_solanago.PublicKey) *WithdrawSharesFromTreasury {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(pdaBasedTransferAuthority)
	return inst
}

// GetPdaBasedTransferAuthorityAccount gets the "pdaBasedTransferAuthority" account.
// PDA-based Transfer authority to move tokens from treasury to your destination[PREFIX, program_id]
func (inst *WithdrawSharesFromTreasury) GetPdaBasedTransferAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetVaultAuthorityAccount sets the "vaultAuthority" account.
// Authority of vault
func (inst *WithdrawSharesFromTreasury) SetVaultAuthorityAccount(vaultAuthority ag_solanago.PublicKey) *WithdrawSharesFromTreasury {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(vaultAuthority).SIGNER()
	return inst
}

// GetVaultAuthorityAccount gets the "vaultAuthority" account.
// Authority of vault
func (inst *WithdrawSharesFromTreasury) GetVaultAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
// Token program
func (inst *WithdrawSharesFromTreasury) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *WithdrawSharesFromTreasury {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
// Token program
func (inst *WithdrawSharesFromTreasury) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetRentSysvarAccount sets the "rentSysvar" account.
// Rent sysvar
func (inst *WithdrawSharesFromTreasury) SetRentSysvarAccount(rentSysvar ag_solanago.PublicKey) *WithdrawSharesFromTreasury {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(rentSysvar)
	return inst
}

// GetRentSysvarAccount gets the "rentSysvar" account.
// Rent sysvar
func (inst *WithdrawSharesFromTreasury) GetRentSysvarAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

func (inst WithdrawSharesFromTreasury) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_WithdrawSharesFromTreasury),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst WithdrawSharesFromTreasury) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *WithdrawSharesFromTreasury) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Args == nil {
			return errors.New("Args parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.InitializedDestination is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.FractionTreasury is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.InitializedActiveTokenVault is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.PdaBasedTransferAuthority is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.VaultAuthority is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.RentSysvar is not set")
		}
	}
	return nil
}

func (inst *WithdrawSharesFromTreasury) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("WithdrawSharesFromTreasury")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Args", *inst.Args))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=7]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("     initializedDestination", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("           fractionTreasury", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("initializedActiveTokenVault", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("  pdaBasedTransferAuthority", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("             vaultAuthority", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("               tokenProgram", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("                 rentSysvar", inst.AccountMetaSlice.Get(6)))
					})
				})
		})
}

func (obj WithdrawSharesFromTreasury) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Args` param:
	err = encoder.Encode(obj.Args)
	if err != nil {
		return err
	}
	return nil
}
func (obj *WithdrawSharesFromTreasury) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Args`:
	err = decoder.Decode(&obj.Args)
	if err != nil {
		return err
	}
	return nil
}

// NewWithdrawSharesFromTreasuryInstruction declares a new WithdrawSharesFromTreasury instruction with the provided parameters and accounts.
func NewWithdrawSharesFromTreasuryInstruction(
	// Parameters:
	args NumberOfShareArgs,
	// Accounts:
	initializedDestination ag_solanago.PublicKey,
	fractionTreasury ag_solanago.PublicKey,
	initializedActiveTokenVault ag_solanago.PublicKey,
	pdaBasedTransferAuthority ag_solanago.PublicKey,
	vaultAuthority ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	rentSysvar ag_solanago.PublicKey) *WithdrawSharesFromTreasury {
	return NewWithdrawSharesFromTreasuryInstructionBuilder().
		SetArgs(args).
		SetInitializedDestinationAccount(initializedDestination).
		SetFractionTreasuryAccount(fractionTreasury).
		SetInitializedActiveTokenVaultAccount(initializedActiveTokenVault).
		SetPdaBasedTransferAuthorityAccount(pdaBasedTransferAuthority).
		SetVaultAuthorityAccount(vaultAuthority).
		SetTokenProgramAccount(tokenProgram).
		SetRentSysvarAccount(rentSysvar)
}
