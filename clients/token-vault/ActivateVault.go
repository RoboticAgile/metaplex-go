// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package token_vault

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/RoboticAgile/solana-go"
	ag_format "github.com/RoboticAgile/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Activates the vault, distributing initial shares into the fraction treasury.
// Tokens can no longer be removed in this state until Combination.
type ActivateVault struct {
	Args *NumberOfShareArgs

	// [0] = [WRITE] initializedInactivatedFractionalizedTokenVault
	// ··········· Initialized inactivated fractionalized token vault
	//
	// [1] = [WRITE] fractionMint
	// ··········· Fraction mint
	//
	// [2] = [WRITE] fractionTreasury
	// ··········· Fraction treasury
	//
	// [3] = [] fractionMintAuthority
	// ··········· Fraction mint authority for the program - seed of [PREFIX, program_id]
	//
	// [4] = [SIGNER] vaultAuthority
	// ··········· Authority on the vault
	//
	// [5] = [] tokenProgram
	// ··········· Token program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewActivateVaultInstructionBuilder creates a new `ActivateVault` instruction builder.
func NewActivateVaultInstructionBuilder() *ActivateVault {
	nd := &ActivateVault{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 6),
	}
	return nd
}

// SetArgs sets the "args" parameter.
func (inst *ActivateVault) SetArgs(args NumberOfShareArgs) *ActivateVault {
	inst.Args = &args
	return inst
}

// SetInitializedInactivatedFractionalizedTokenVaultAccount sets the "initializedInactivatedFractionalizedTokenVault" account.
// Initialized inactivated fractionalized token vault
func (inst *ActivateVault) SetInitializedInactivatedFractionalizedTokenVaultAccount(initializedInactivatedFractionalizedTokenVault ag_solanago.PublicKey) *ActivateVault {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(initializedInactivatedFractionalizedTokenVault).WRITE()
	return inst
}

// GetInitializedInactivatedFractionalizedTokenVaultAccount gets the "initializedInactivatedFractionalizedTokenVault" account.
// Initialized inactivated fractionalized token vault
func (inst *ActivateVault) GetInitializedInactivatedFractionalizedTokenVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetFractionMintAccount sets the "fractionMint" account.
// Fraction mint
func (inst *ActivateVault) SetFractionMintAccount(fractionMint ag_solanago.PublicKey) *ActivateVault {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(fractionMint).WRITE()
	return inst
}

// GetFractionMintAccount gets the "fractionMint" account.
// Fraction mint
func (inst *ActivateVault) GetFractionMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetFractionTreasuryAccount sets the "fractionTreasury" account.
// Fraction treasury
func (inst *ActivateVault) SetFractionTreasuryAccount(fractionTreasury ag_solanago.PublicKey) *ActivateVault {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(fractionTreasury).WRITE()
	return inst
}

// GetFractionTreasuryAccount gets the "fractionTreasury" account.
// Fraction treasury
func (inst *ActivateVault) GetFractionTreasuryAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetFractionMintAuthorityAccount sets the "fractionMintAuthority" account.
// Fraction mint authority for the program - seed of [PREFIX, program_id]
func (inst *ActivateVault) SetFractionMintAuthorityAccount(fractionMintAuthority ag_solanago.PublicKey) *ActivateVault {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(fractionMintAuthority)
	return inst
}

// GetFractionMintAuthorityAccount gets the "fractionMintAuthority" account.
// Fraction mint authority for the program - seed of [PREFIX, program_id]
func (inst *ActivateVault) GetFractionMintAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetVaultAuthorityAccount sets the "vaultAuthority" account.
// Authority on the vault
func (inst *ActivateVault) SetVaultAuthorityAccount(vaultAuthority ag_solanago.PublicKey) *ActivateVault {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(vaultAuthority).SIGNER()
	return inst
}

// GetVaultAuthorityAccount gets the "vaultAuthority" account.
// Authority on the vault
func (inst *ActivateVault) GetVaultAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
// Token program
func (inst *ActivateVault) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *ActivateVault {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
// Token program
func (inst *ActivateVault) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

func (inst ActivateVault) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_ActivateVault),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst ActivateVault) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *ActivateVault) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Args == nil {
			return errors.New("Args parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.InitializedInactivatedFractionalizedTokenVault is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.FractionMint is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.FractionTreasury is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.FractionMintAuthority is not set")
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

func (inst *ActivateVault) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("ActivateVault")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Args", *inst.Args))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=6]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("initializedInactivatedFractionalizedTokenVault", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("                                  fractionMint", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("                              fractionTreasury", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("                         fractionMintAuthority", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("                                vaultAuthority", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("                                  tokenProgram", inst.AccountMetaSlice.Get(5)))
					})
				})
		})
}

func (obj ActivateVault) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Args` param:
	err = encoder.Encode(obj.Args)
	if err != nil {
		return err
	}
	return nil
}
func (obj *ActivateVault) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Args`:
	err = decoder.Decode(&obj.Args)
	if err != nil {
		return err
	}
	return nil
}

// NewActivateVaultInstruction declares a new ActivateVault instruction with the provided parameters and accounts.
func NewActivateVaultInstruction(
	// Parameters:
	args NumberOfShareArgs,
	// Accounts:
	initializedInactivatedFractionalizedTokenVault ag_solanago.PublicKey,
	fractionMint ag_solanago.PublicKey,
	fractionTreasury ag_solanago.PublicKey,
	fractionMintAuthority ag_solanago.PublicKey,
	vaultAuthority ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey) *ActivateVault {
	return NewActivateVaultInstructionBuilder().
		SetArgs(args).
		SetInitializedInactivatedFractionalizedTokenVaultAccount(initializedInactivatedFractionalizedTokenVault).
		SetFractionMintAccount(fractionMint).
		SetFractionTreasuryAccount(fractionTreasury).
		SetFractionMintAuthorityAccount(fractionMintAuthority).
		SetVaultAuthorityAccount(vaultAuthority).
		SetTokenProgramAccount(tokenProgram)
}
