// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package token_metadata

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/RoboticAgile/solana-go"
	ag_format "github.com/RoboticAgile/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Revoke account to call `utilize` on this NFT
type RevokeUseAuthority struct {

	// [0] = [WRITE] useAuthorityRecordPDA
	// ··········· Use Authority Record PDA
	//
	// [1] = [WRITE] ownedToken
	// ··········· Owned Token Account Of Mint
	//
	// [2] = [SIGNER] owner
	// ··········· Owner
	//
	// [3] = [SIGNER] payer
	// ··········· Payer
	//
	// [4] = [] useAuthority
	// ··········· A Use Authority
	//
	// [5] = [] metadata
	// ··········· Metadata account
	//
	// [6] = [] metadataMint
	// ··········· Mint of Metadata
	//
	// [7] = [] tokenProgram
	// ··········· Token program
	//
	// [8] = [] system
	// ··········· System program
	//
	// [9] = [] rent
	// ··········· Rent info
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewRevokeUseAuthorityInstructionBuilder creates a new `RevokeUseAuthority` instruction builder.
func NewRevokeUseAuthorityInstructionBuilder() *RevokeUseAuthority {
	nd := &RevokeUseAuthority{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 10),
	}
	return nd
}

// SetUseAuthorityRecordPDAAccount sets the "useAuthorityRecordPDA" account.
// Use Authority Record PDA
func (inst *RevokeUseAuthority) SetUseAuthorityRecordPDAAccount(useAuthorityRecordPDA ag_solanago.PublicKey) *RevokeUseAuthority {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(useAuthorityRecordPDA).WRITE()
	return inst
}

// GetUseAuthorityRecordPDAAccount gets the "useAuthorityRecordPDA" account.
// Use Authority Record PDA
func (inst *RevokeUseAuthority) GetUseAuthorityRecordPDAAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetOwnedTokenAccount sets the "ownedToken" account.
// Owned Token Account Of Mint
func (inst *RevokeUseAuthority) SetOwnedTokenAccount(ownedToken ag_solanago.PublicKey) *RevokeUseAuthority {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(ownedToken).WRITE()
	return inst
}

// GetOwnedTokenAccount gets the "ownedToken" account.
// Owned Token Account Of Mint
func (inst *RevokeUseAuthority) GetOwnedTokenAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetOwnerAccount sets the "owner" account.
// Owner
func (inst *RevokeUseAuthority) SetOwnerAccount(owner ag_solanago.PublicKey) *RevokeUseAuthority {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(owner).SIGNER()
	return inst
}

// GetOwnerAccount gets the "owner" account.
// Owner
func (inst *RevokeUseAuthority) GetOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetPayerAccount sets the "payer" account.
// Payer
func (inst *RevokeUseAuthority) SetPayerAccount(payer ag_solanago.PublicKey) *RevokeUseAuthority {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(payer).SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
// Payer
func (inst *RevokeUseAuthority) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetUseAuthorityAccount sets the "useAuthority" account.
// A Use Authority
func (inst *RevokeUseAuthority) SetUseAuthorityAccount(useAuthority ag_solanago.PublicKey) *RevokeUseAuthority {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(useAuthority)
	return inst
}

// GetUseAuthorityAccount gets the "useAuthority" account.
// A Use Authority
func (inst *RevokeUseAuthority) GetUseAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetMetadataAccount sets the "metadata" account.
// Metadata account
func (inst *RevokeUseAuthority) SetMetadataAccount(metadata ag_solanago.PublicKey) *RevokeUseAuthority {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(metadata)
	return inst
}

// GetMetadataAccount gets the "metadata" account.
// Metadata account
func (inst *RevokeUseAuthority) GetMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetMetadataMintAccount sets the "metadataMint" account.
// Mint of Metadata
func (inst *RevokeUseAuthority) SetMetadataMintAccount(metadataMint ag_solanago.PublicKey) *RevokeUseAuthority {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(metadataMint)
	return inst
}

// GetMetadataMintAccount gets the "metadataMint" account.
// Mint of Metadata
func (inst *RevokeUseAuthority) GetMetadataMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
// Token program
func (inst *RevokeUseAuthority) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *RevokeUseAuthority {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
// Token program
func (inst *RevokeUseAuthority) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetSystemAccount sets the "system" account.
// System program
func (inst *RevokeUseAuthority) SetSystemAccount(system ag_solanago.PublicKey) *RevokeUseAuthority {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(system)
	return inst
}

// GetSystemAccount gets the "system" account.
// System program
func (inst *RevokeUseAuthority) GetSystemAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetRentAccount sets the "rent" account.
// Rent info
func (inst *RevokeUseAuthority) SetRentAccount(rent ag_solanago.PublicKey) *RevokeUseAuthority {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
// Rent info
func (inst *RevokeUseAuthority) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

func (inst RevokeUseAuthority) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_RevokeUseAuthority),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst RevokeUseAuthority) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *RevokeUseAuthority) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.UseAuthorityRecordPDA is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.OwnedToken is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Owner is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.UseAuthority is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.Metadata is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.MetadataMint is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.System is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.Rent is not set")
		}
	}
	return nil
}

func (inst *RevokeUseAuthority) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("RevokeUseAuthority")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=10]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("useAuthorityRecordPDA", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("           ownedToken", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("                owner", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("                payer", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("         useAuthority", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("             metadata", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("         metadataMint", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("         tokenProgram", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("               system", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("                 rent", inst.AccountMetaSlice.Get(9)))
					})
				})
		})
}

func (obj RevokeUseAuthority) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *RevokeUseAuthority) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewRevokeUseAuthorityInstruction declares a new RevokeUseAuthority instruction with the provided parameters and accounts.
func NewRevokeUseAuthorityInstruction(
	// Accounts:
	useAuthorityRecordPDA ag_solanago.PublicKey,
	ownedToken ag_solanago.PublicKey,
	owner ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	useAuthority ag_solanago.PublicKey,
	metadata ag_solanago.PublicKey,
	metadataMint ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	system ag_solanago.PublicKey,
	rent ag_solanago.PublicKey) *RevokeUseAuthority {
	return NewRevokeUseAuthorityInstructionBuilder().
		SetUseAuthorityRecordPDAAccount(useAuthorityRecordPDA).
		SetOwnedTokenAccount(ownedToken).
		SetOwnerAccount(owner).
		SetPayerAccount(payer).
		SetUseAuthorityAccount(useAuthority).
		SetMetadataAccount(metadata).
		SetMetadataMintAccount(metadataMint).
		SetTokenProgramAccount(tokenProgram).
		SetSystemAccount(system).
		SetRentAccount(rent)
}
