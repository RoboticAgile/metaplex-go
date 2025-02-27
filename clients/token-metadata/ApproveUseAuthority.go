// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package token_metadata

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/RoboticAgile/solana-go"
	ag_format "github.com/RoboticAgile/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Approve another account to call `utilize` on this NFT
type ApproveUseAuthority struct {
	Args *ApproveUseAuthorityArgs

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
	// [7] = [] programAsSigner
	// ··········· Program As Signer (Burner)
	//
	// [8] = [] tokenProgram
	// ··········· Token program
	//
	// [9] = [] system
	// ··········· System program
	//
	// [10] = [] rent
	// ··········· Rent info
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewApproveUseAuthorityInstructionBuilder creates a new `ApproveUseAuthority` instruction builder.
func NewApproveUseAuthorityInstructionBuilder() *ApproveUseAuthority {
	nd := &ApproveUseAuthority{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 11),
	}
	return nd
}

// SetArgs sets the "args" parameter.
func (inst *ApproveUseAuthority) SetArgs(args ApproveUseAuthorityArgs) *ApproveUseAuthority {
	inst.Args = &args
	return inst
}

// SetUseAuthorityRecordPDAAccount sets the "useAuthorityRecordPDA" account.
// Use Authority Record PDA
func (inst *ApproveUseAuthority) SetUseAuthorityRecordPDAAccount(useAuthorityRecordPDA ag_solanago.PublicKey) *ApproveUseAuthority {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(useAuthorityRecordPDA).WRITE()
	return inst
}

// GetUseAuthorityRecordPDAAccount gets the "useAuthorityRecordPDA" account.
// Use Authority Record PDA
func (inst *ApproveUseAuthority) GetUseAuthorityRecordPDAAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetOwnedTokenAccount sets the "ownedToken" account.
// Owned Token Account Of Mint
func (inst *ApproveUseAuthority) SetOwnedTokenAccount(ownedToken ag_solanago.PublicKey) *ApproveUseAuthority {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(ownedToken).WRITE()
	return inst
}

// GetOwnedTokenAccount gets the "ownedToken" account.
// Owned Token Account Of Mint
func (inst *ApproveUseAuthority) GetOwnedTokenAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetOwnerAccount sets the "owner" account.
// Owner
func (inst *ApproveUseAuthority) SetOwnerAccount(owner ag_solanago.PublicKey) *ApproveUseAuthority {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(owner).SIGNER()
	return inst
}

// GetOwnerAccount gets the "owner" account.
// Owner
func (inst *ApproveUseAuthority) GetOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetPayerAccount sets the "payer" account.
// Payer
func (inst *ApproveUseAuthority) SetPayerAccount(payer ag_solanago.PublicKey) *ApproveUseAuthority {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(payer).SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
// Payer
func (inst *ApproveUseAuthority) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetUseAuthorityAccount sets the "useAuthority" account.
// A Use Authority
func (inst *ApproveUseAuthority) SetUseAuthorityAccount(useAuthority ag_solanago.PublicKey) *ApproveUseAuthority {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(useAuthority)
	return inst
}

// GetUseAuthorityAccount gets the "useAuthority" account.
// A Use Authority
func (inst *ApproveUseAuthority) GetUseAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetMetadataAccount sets the "metadata" account.
// Metadata account
func (inst *ApproveUseAuthority) SetMetadataAccount(metadata ag_solanago.PublicKey) *ApproveUseAuthority {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(metadata)
	return inst
}

// GetMetadataAccount gets the "metadata" account.
// Metadata account
func (inst *ApproveUseAuthority) GetMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetMetadataMintAccount sets the "metadataMint" account.
// Mint of Metadata
func (inst *ApproveUseAuthority) SetMetadataMintAccount(metadataMint ag_solanago.PublicKey) *ApproveUseAuthority {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(metadataMint)
	return inst
}

// GetMetadataMintAccount gets the "metadataMint" account.
// Mint of Metadata
func (inst *ApproveUseAuthority) GetMetadataMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetProgramAsSignerAccount sets the "programAsSigner" account.
// Program As Signer (Burner)
func (inst *ApproveUseAuthority) SetProgramAsSignerAccount(programAsSigner ag_solanago.PublicKey) *ApproveUseAuthority {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(programAsSigner)
	return inst
}

// GetProgramAsSignerAccount gets the "programAsSigner" account.
// Program As Signer (Burner)
func (inst *ApproveUseAuthority) GetProgramAsSignerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
// Token program
func (inst *ApproveUseAuthority) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *ApproveUseAuthority {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
// Token program
func (inst *ApproveUseAuthority) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetSystemAccount sets the "system" account.
// System program
func (inst *ApproveUseAuthority) SetSystemAccount(system ag_solanago.PublicKey) *ApproveUseAuthority {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(system)
	return inst
}

// GetSystemAccount gets the "system" account.
// System program
func (inst *ApproveUseAuthority) GetSystemAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetRentAccount sets the "rent" account.
// Rent info
func (inst *ApproveUseAuthority) SetRentAccount(rent ag_solanago.PublicKey) *ApproveUseAuthority {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
// Rent info
func (inst *ApproveUseAuthority) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

func (inst ApproveUseAuthority) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_ApproveUseAuthority),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst ApproveUseAuthority) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *ApproveUseAuthority) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Args == nil {
			return errors.New("Args parameter is not set")
		}
	}

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
			return errors.New("accounts.ProgramAsSigner is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.System is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.Rent is not set")
		}
	}
	return nil
}

func (inst *ApproveUseAuthority) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("ApproveUseAuthority")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Args", *inst.Args))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=11]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("useAuthorityRecordPDA", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("           ownedToken", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("                owner", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("                payer", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("         useAuthority", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("             metadata", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("         metadataMint", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("      programAsSigner", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("         tokenProgram", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("               system", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("                 rent", inst.AccountMetaSlice.Get(10)))
					})
				})
		})
}

func (obj ApproveUseAuthority) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Args` param:
	err = encoder.Encode(obj.Args)
	if err != nil {
		return err
	}
	return nil
}
func (obj *ApproveUseAuthority) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Args`:
	err = decoder.Decode(&obj.Args)
	if err != nil {
		return err
	}
	return nil
}

// NewApproveUseAuthorityInstruction declares a new ApproveUseAuthority instruction with the provided parameters and accounts.
func NewApproveUseAuthorityInstruction(
	// Parameters:
	args ApproveUseAuthorityArgs,
	// Accounts:
	useAuthorityRecordPDA ag_solanago.PublicKey,
	ownedToken ag_solanago.PublicKey,
	owner ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	useAuthority ag_solanago.PublicKey,
	metadata ag_solanago.PublicKey,
	metadataMint ag_solanago.PublicKey,
	programAsSigner ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	system ag_solanago.PublicKey,
	rent ag_solanago.PublicKey) *ApproveUseAuthority {
	return NewApproveUseAuthorityInstructionBuilder().
		SetArgs(args).
		SetUseAuthorityRecordPDAAccount(useAuthorityRecordPDA).
		SetOwnedTokenAccount(ownedToken).
		SetOwnerAccount(owner).
		SetPayerAccount(payer).
		SetUseAuthorityAccount(useAuthority).
		SetMetadataAccount(metadata).
		SetMetadataMintAccount(metadataMint).
		SetProgramAsSignerAccount(programAsSigner).
		SetTokenProgramAccount(tokenProgram).
		SetSystemAccount(system).
		SetRentAccount(rent)
}
