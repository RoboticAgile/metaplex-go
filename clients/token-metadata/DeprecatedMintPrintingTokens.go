// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package token_metadata

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/RoboticAgile/solana-go"
	ag_format "github.com/RoboticAgile/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Using your update authority, mint printing tokens for your master edition.
type DeprecatedMintPrintingTokens struct {
	Args *MintPrintingTokensViaTokenArgs

	// [0] = [WRITE] destinationAccount
	// ··········· Destination account
	//
	// [1] = [WRITE] printingMint
	// ··········· Printing mint
	//
	// [2] = [SIGNER] updateAuthority
	// ··········· Update authority
	//
	// [3] = [] metadataKeyPDA
	// ··········· Metadata key (pda of ['metadata', program id, mint id])
	//
	// [4] = [] masterEditionV1
	// ··········· Master Edition V1 key (pda of ['metadata', program id, mint id, 'edition'])
	//
	// [5] = [] tokenProgram
	// ··········· Token program
	//
	// [6] = [] rent
	// ··········· Rent
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewDeprecatedMintPrintingTokensInstructionBuilder creates a new `DeprecatedMintPrintingTokens` instruction builder.
func NewDeprecatedMintPrintingTokensInstructionBuilder() *DeprecatedMintPrintingTokens {
	nd := &DeprecatedMintPrintingTokens{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 7),
	}
	return nd
}

// SetArgs sets the "args" parameter.
func (inst *DeprecatedMintPrintingTokens) SetArgs(args MintPrintingTokensViaTokenArgs) *DeprecatedMintPrintingTokens {
	inst.Args = &args
	return inst
}

// SetDestinationAccount sets the "destinationAccount" account.
// Destination account
func (inst *DeprecatedMintPrintingTokens) SetDestinationAccount(destinationAccount ag_solanago.PublicKey) *DeprecatedMintPrintingTokens {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(destinationAccount).WRITE()
	return inst
}

// GetDestinationAccount gets the "destinationAccount" account.
// Destination account
func (inst *DeprecatedMintPrintingTokens) GetDestinationAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetPrintingMintAccount sets the "printingMint" account.
// Printing mint
func (inst *DeprecatedMintPrintingTokens) SetPrintingMintAccount(printingMint ag_solanago.PublicKey) *DeprecatedMintPrintingTokens {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(printingMint).WRITE()
	return inst
}

// GetPrintingMintAccount gets the "printingMint" account.
// Printing mint
func (inst *DeprecatedMintPrintingTokens) GetPrintingMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetUpdateAuthorityAccount sets the "updateAuthority" account.
// Update authority
func (inst *DeprecatedMintPrintingTokens) SetUpdateAuthorityAccount(updateAuthority ag_solanago.PublicKey) *DeprecatedMintPrintingTokens {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(updateAuthority).SIGNER()
	return inst
}

// GetUpdateAuthorityAccount gets the "updateAuthority" account.
// Update authority
func (inst *DeprecatedMintPrintingTokens) GetUpdateAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetMetadataKeyPDAAccount sets the "metadataKeyPDA" account.
// Metadata key (pda of ['metadata', program id, mint id])
func (inst *DeprecatedMintPrintingTokens) SetMetadataKeyPDAAccount(metadataKeyPDA ag_solanago.PublicKey) *DeprecatedMintPrintingTokens {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(metadataKeyPDA)
	return inst
}

// GetMetadataKeyPDAAccount gets the "metadataKeyPDA" account.
// Metadata key (pda of ['metadata', program id, mint id])
func (inst *DeprecatedMintPrintingTokens) GetMetadataKeyPDAAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetMasterEditionV1Account sets the "masterEditionV1" account.
// Master Edition V1 key (pda of ['metadata', program id, mint id, 'edition'])
func (inst *DeprecatedMintPrintingTokens) SetMasterEditionV1Account(masterEditionV1 ag_solanago.PublicKey) *DeprecatedMintPrintingTokens {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(masterEditionV1)
	return inst
}

// GetMasterEditionV1Account gets the "masterEditionV1" account.
// Master Edition V1 key (pda of ['metadata', program id, mint id, 'edition'])
func (inst *DeprecatedMintPrintingTokens) GetMasterEditionV1Account() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
// Token program
func (inst *DeprecatedMintPrintingTokens) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *DeprecatedMintPrintingTokens {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
// Token program
func (inst *DeprecatedMintPrintingTokens) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetRentAccount sets the "rent" account.
// Rent
func (inst *DeprecatedMintPrintingTokens) SetRentAccount(rent ag_solanago.PublicKey) *DeprecatedMintPrintingTokens {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
// Rent
func (inst *DeprecatedMintPrintingTokens) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

func (inst DeprecatedMintPrintingTokens) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_DeprecatedMintPrintingTokens),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst DeprecatedMintPrintingTokens) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *DeprecatedMintPrintingTokens) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Args == nil {
			return errors.New("Args parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.DestinationAccount is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.PrintingMint is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.UpdateAuthority is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.MetadataKeyPDA is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.MasterEditionV1 is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.Rent is not set")
		}
	}
	return nil
}

func (inst *DeprecatedMintPrintingTokens) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("DeprecatedMintPrintingTokens")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Args", *inst.Args))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=7]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("    destination", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("   printingMint", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("updateAuthority", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta(" metadataKeyPDA", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("masterEditionV1", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("   tokenProgram", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("           rent", inst.AccountMetaSlice.Get(6)))
					})
				})
		})
}

func (obj DeprecatedMintPrintingTokens) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Args` param:
	err = encoder.Encode(obj.Args)
	if err != nil {
		return err
	}
	return nil
}
func (obj *DeprecatedMintPrintingTokens) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Args`:
	err = decoder.Decode(&obj.Args)
	if err != nil {
		return err
	}
	return nil
}

// NewDeprecatedMintPrintingTokensInstruction declares a new DeprecatedMintPrintingTokens instruction with the provided parameters and accounts.
func NewDeprecatedMintPrintingTokensInstruction(
	// Parameters:
	args MintPrintingTokensViaTokenArgs,
	// Accounts:
	destinationAccount ag_solanago.PublicKey,
	printingMint ag_solanago.PublicKey,
	updateAuthority ag_solanago.PublicKey,
	metadataKeyPDA ag_solanago.PublicKey,
	masterEditionV1 ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	rent ag_solanago.PublicKey) *DeprecatedMintPrintingTokens {
	return NewDeprecatedMintPrintingTokensInstructionBuilder().
		SetArgs(args).
		SetDestinationAccount(destinationAccount).
		SetPrintingMintAccount(printingMint).
		SetUpdateAuthorityAccount(updateAuthority).
		SetMetadataKeyPDAAccount(metadataKeyPDA).
		SetMasterEditionV1Account(masterEditionV1).
		SetTokenProgramAccount(tokenProgram).
		SetRentAccount(rent)
}
