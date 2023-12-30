// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package token_metadata

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/RoboticAgile/solana-go"
	ag_format "github.com/RoboticAgile/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Register a Metadata as a Master Edition V2, which means Edition V2s can be minted.
// Henceforth, no further tokens will be mintable from this primary mint. Will throw an error if more than one
// token exists, and will throw an error if less than one token exists in this primary mint.
type CreateMasterEditionV3 struct {
	Args *CreateMasterEditionArgs

	// [0] = [WRITE] unallocatedEditionV2
	// ··········· Unallocated edition V2 account with address as pda of ['metadata', program id, mint, 'edition']
	//
	// [1] = [WRITE] metadataMint
	// ··········· Metadata mint
	//
	// [2] = [SIGNER] updateAuthority
	// ··········· Update authority
	//
	// [3] = [SIGNER] mintAuthority
	// ··········· Mint authority on the metadata's mint - THIS WILL TRANSFER AUTHORITY AWAY FROM THIS KEY
	//
	// [4] = [SIGNER] payer
	// ··········· payer
	//
	// [5] = [WRITE] metadata
	// ··········· Metadata account
	//
	// [6] = [] tokenProgram
	// ··········· Token program
	//
	// [7] = [] system
	// ··········· System program
	//
	// [8] = [] rent
	// ··········· Rent info
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewCreateMasterEditionV3InstructionBuilder creates a new `CreateMasterEditionV3` instruction builder.
func NewCreateMasterEditionV3InstructionBuilder() *CreateMasterEditionV3 {
	nd := &CreateMasterEditionV3{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 9),
	}
	return nd
}

// SetArgs sets the "args" parameter.
func (inst *CreateMasterEditionV3) SetArgs(args CreateMasterEditionArgs) *CreateMasterEditionV3 {
	inst.Args = &args
	return inst
}

// SetUnallocatedEditionV2Account sets the "unallocatedEditionV2" account.
// Unallocated edition V2 account with address as pda of ['metadata', program id, mint, 'edition']
func (inst *CreateMasterEditionV3) SetUnallocatedEditionV2Account(unallocatedEditionV2 ag_solanago.PublicKey) *CreateMasterEditionV3 {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(unallocatedEditionV2).WRITE()
	return inst
}

// GetUnallocatedEditionV2Account gets the "unallocatedEditionV2" account.
// Unallocated edition V2 account with address as pda of ['metadata', program id, mint, 'edition']
func (inst *CreateMasterEditionV3) GetUnallocatedEditionV2Account() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetMetadataMintAccount sets the "metadataMint" account.
// Metadata mint
func (inst *CreateMasterEditionV3) SetMetadataMintAccount(metadataMint ag_solanago.PublicKey) *CreateMasterEditionV3 {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(metadataMint).WRITE()
	return inst
}

// GetMetadataMintAccount gets the "metadataMint" account.
// Metadata mint
func (inst *CreateMasterEditionV3) GetMetadataMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetUpdateAuthorityAccount sets the "updateAuthority" account.
// Update authority
func (inst *CreateMasterEditionV3) SetUpdateAuthorityAccount(updateAuthority ag_solanago.PublicKey) *CreateMasterEditionV3 {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(updateAuthority).SIGNER()
	return inst
}

// GetUpdateAuthorityAccount gets the "updateAuthority" account.
// Update authority
func (inst *CreateMasterEditionV3) GetUpdateAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetMintAuthorityAccount sets the "mintAuthority" account.
// Mint authority on the metadata's mint - THIS WILL TRANSFER AUTHORITY AWAY FROM THIS KEY
func (inst *CreateMasterEditionV3) SetMintAuthorityAccount(mintAuthority ag_solanago.PublicKey) *CreateMasterEditionV3 {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(mintAuthority).SIGNER()
	return inst
}

// GetMintAuthorityAccount gets the "mintAuthority" account.
// Mint authority on the metadata's mint - THIS WILL TRANSFER AUTHORITY AWAY FROM THIS KEY
func (inst *CreateMasterEditionV3) GetMintAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetPayerAccount sets the "payer" account.
// payer
func (inst *CreateMasterEditionV3) SetPayerAccount(payer ag_solanago.PublicKey) *CreateMasterEditionV3 {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(payer).SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
// payer
func (inst *CreateMasterEditionV3) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetMetadataAccount sets the "metadata" account.
// Metadata account
func (inst *CreateMasterEditionV3) SetMetadataAccount(metadata ag_solanago.PublicKey) *CreateMasterEditionV3 {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(metadata).WRITE()
	return inst
}

// GetMetadataAccount gets the "metadata" account.
// Metadata account
func (inst *CreateMasterEditionV3) GetMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
// Token program
func (inst *CreateMasterEditionV3) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *CreateMasterEditionV3 {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
// Token program
func (inst *CreateMasterEditionV3) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetSystemAccount sets the "system" account.
// System program
func (inst *CreateMasterEditionV3) SetSystemAccount(system ag_solanago.PublicKey) *CreateMasterEditionV3 {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(system)
	return inst
}

// GetSystemAccount gets the "system" account.
// System program
func (inst *CreateMasterEditionV3) GetSystemAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetRentAccount sets the "rent" account.
// Rent info
func (inst *CreateMasterEditionV3) SetRentAccount(rent ag_solanago.PublicKey) *CreateMasterEditionV3 {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
// Rent info
func (inst *CreateMasterEditionV3) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

func (inst CreateMasterEditionV3) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_CreateMasterEditionV3),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst CreateMasterEditionV3) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *CreateMasterEditionV3) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Args == nil {
			return errors.New("Args parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.UnallocatedEditionV2 is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.MetadataMint is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.UpdateAuthority is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.MintAuthority is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.Metadata is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.System is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.Rent is not set")
		}
	}
	return nil
}

func (inst *CreateMasterEditionV3) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("CreateMasterEditionV3")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Args", *inst.Args))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=9]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("unallocatedEditionV2", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("        metadataMint", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("     updateAuthority", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("       mintAuthority", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("               payer", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("            metadata", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("        tokenProgram", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("              system", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("                rent", inst.AccountMetaSlice.Get(8)))
					})
				})
		})
}

func (obj CreateMasterEditionV3) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Args` param:
	err = encoder.Encode(obj.Args)
	if err != nil {
		return err
	}
	return nil
}
func (obj *CreateMasterEditionV3) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Args`:
	err = decoder.Decode(&obj.Args)
	if err != nil {
		return err
	}
	return nil
}

// NewCreateMasterEditionV3Instruction declares a new CreateMasterEditionV3 instruction with the provided parameters and accounts.
func NewCreateMasterEditionV3Instruction(
	// Parameters:
	args CreateMasterEditionArgs,
	// Accounts:
	unallocatedEditionV2 ag_solanago.PublicKey,
	metadataMint ag_solanago.PublicKey,
	updateAuthority ag_solanago.PublicKey,
	mintAuthority ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	metadata ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	system ag_solanago.PublicKey,
	rent ag_solanago.PublicKey) *CreateMasterEditionV3 {
	return NewCreateMasterEditionV3InstructionBuilder().
		SetArgs(args).
		SetUnallocatedEditionV2Account(unallocatedEditionV2).
		SetMetadataMintAccount(metadataMint).
		SetUpdateAuthorityAccount(updateAuthority).
		SetMintAuthorityAccount(mintAuthority).
		SetPayerAccount(payer).
		SetMetadataAccount(metadata).
		SetTokenProgramAccount(tokenProgram).
		SetSystemAccount(system).
		SetRentAccount(rent)
}
