// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package token_vault

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/desperatee/solana-go"
	ag_format "github.com/desperatee/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Returns shares to the vault if you wish to remove them from circulation.
type AddSharesToTreasury struct {
	Args *NumberOfShareArgs

	// [0] = [WRITE] initializedSourceAccount
	// ··········· Initialized account from which shares will be withdrawn
	//
	// [1] = [WRITE] fractionTreasury
	// ··········· Fraction treasury
	//
	// [2] = [] initializedActiveTokenVault
	// ··········· The initialized active token vault
	//
	// [3] = [SIGNER] transferAuthority
	// ··········· Transfer authority to move tokens from your account to treasury
	//
	// [4] = [SIGNER] vaultAuthority
	// ··········· Authority of vault
	//
	// [5] = [] tokenProgram
	// ··········· Token program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewAddSharesToTreasuryInstructionBuilder creates a new `AddSharesToTreasury` instruction builder.
func NewAddSharesToTreasuryInstructionBuilder() *AddSharesToTreasury {
	nd := &AddSharesToTreasury{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 6),
	}
	return nd
}

// SetArgs sets the "args" parameter.
func (inst *AddSharesToTreasury) SetArgs(args NumberOfShareArgs) *AddSharesToTreasury {
	inst.Args = &args
	return inst
}

// SetInitializedSourceAccount sets the "initializedSourceAccount" account.
// Initialized account from which shares will be withdrawn
func (inst *AddSharesToTreasury) SetInitializedSourceAccount(initializedSourceAccount ag_solanago.PublicKey) *AddSharesToTreasury {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(initializedSourceAccount).WRITE()
	return inst
}

// GetInitializedSourceAccount gets the "initializedSourceAccount" account.
// Initialized account from which shares will be withdrawn
func (inst *AddSharesToTreasury) GetInitializedSourceAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetFractionTreasuryAccount sets the "fractionTreasury" account.
// Fraction treasury
func (inst *AddSharesToTreasury) SetFractionTreasuryAccount(fractionTreasury ag_solanago.PublicKey) *AddSharesToTreasury {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(fractionTreasury).WRITE()
	return inst
}

// GetFractionTreasuryAccount gets the "fractionTreasury" account.
// Fraction treasury
func (inst *AddSharesToTreasury) GetFractionTreasuryAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetInitializedActiveTokenVaultAccount sets the "initializedActiveTokenVault" account.
// The initialized active token vault
func (inst *AddSharesToTreasury) SetInitializedActiveTokenVaultAccount(initializedActiveTokenVault ag_solanago.PublicKey) *AddSharesToTreasury {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(initializedActiveTokenVault)
	return inst
}

// GetInitializedActiveTokenVaultAccount gets the "initializedActiveTokenVault" account.
// The initialized active token vault
func (inst *AddSharesToTreasury) GetInitializedActiveTokenVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetTransferAuthorityAccount sets the "transferAuthority" account.
// Transfer authority to move tokens from your account to treasury
func (inst *AddSharesToTreasury) SetTransferAuthorityAccount(transferAuthority ag_solanago.PublicKey) *AddSharesToTreasury {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(transferAuthority).SIGNER()
	return inst
}

// GetTransferAuthorityAccount gets the "transferAuthority" account.
// Transfer authority to move tokens from your account to treasury
func (inst *AddSharesToTreasury) GetTransferAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetVaultAuthorityAccount sets the "vaultAuthority" account.
// Authority of vault
func (inst *AddSharesToTreasury) SetVaultAuthorityAccount(vaultAuthority ag_solanago.PublicKey) *AddSharesToTreasury {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(vaultAuthority).SIGNER()
	return inst
}

// GetVaultAuthorityAccount gets the "vaultAuthority" account.
// Authority of vault
func (inst *AddSharesToTreasury) GetVaultAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
// Token program
func (inst *AddSharesToTreasury) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *AddSharesToTreasury {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
// Token program
func (inst *AddSharesToTreasury) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

func (inst AddSharesToTreasury) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_AddSharesToTreasury),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst AddSharesToTreasury) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *AddSharesToTreasury) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Args == nil {
			return errors.New("Args parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.InitializedSourceAccount is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.FractionTreasury is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.InitializedActiveTokenVault is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.TransferAuthority is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.VaultAuthority is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
	}
	return nil
}

func (inst *AddSharesToTreasury) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("AddSharesToTreasury")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Args", *inst.Args))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=6]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("          initializedSource", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("           fractionTreasury", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("initializedActiveTokenVault", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("          transferAuthority", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("             vaultAuthority", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("               tokenProgram", inst.AccountMetaSlice.Get(5)))
					})
				})
		})
}

func (obj AddSharesToTreasury) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Args` param:
	err = encoder.Encode(obj.Args)
	if err != nil {
		return err
	}
	return nil
}
func (obj *AddSharesToTreasury) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Args`:
	err = decoder.Decode(&obj.Args)
	if err != nil {
		return err
	}
	return nil
}

// NewAddSharesToTreasuryInstruction declares a new AddSharesToTreasury instruction with the provided parameters and accounts.
func NewAddSharesToTreasuryInstruction(
	// Parameters:
	args NumberOfShareArgs,
	// Accounts:
	initializedSourceAccount ag_solanago.PublicKey,
	fractionTreasury ag_solanago.PublicKey,
	initializedActiveTokenVault ag_solanago.PublicKey,
	transferAuthority ag_solanago.PublicKey,
	vaultAuthority ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey) *AddSharesToTreasury {
	return NewAddSharesToTreasuryInstructionBuilder().
		SetArgs(args).
		SetInitializedSourceAccount(initializedSourceAccount).
		SetFractionTreasuryAccount(fractionTreasury).
		SetInitializedActiveTokenVaultAccount(initializedActiveTokenVault).
		SetTransferAuthorityAccount(transferAuthority).
		SetVaultAuthorityAccount(vaultAuthority).
		SetTokenProgramAccount(tokenProgram)
}
