// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package token_metadata

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/RoboticAgile/solana-go"
	ag_format "github.com/RoboticAgile/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// If a MetadataAccount Has a Collection allow the UpdateAuthority of the Collection to Verify the NFT Belongs in the Collection
type VerifyCollection struct {

	// [0] = [WRITE] metadata
	// ··········· Metadata account
	//
	// [1] = [SIGNER] collectionUpdateAuthority
	// ··········· Collection Update authority
	//
	// [2] = [SIGNER] payer
	// ··········· payer
	//
	// [3] = [] collectionMint
	// ··········· Mint of the Collection
	//
	// [4] = [] collectionMetadata
	// ··········· Metadata Account of the Collection
	//
	// [5] = [] masterEditionV2
	// ··········· MasterEdition2 Account of the Collection Token
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewVerifyCollectionInstructionBuilder creates a new `VerifyCollection` instruction builder.
func NewVerifyCollectionInstructionBuilder() *VerifyCollection {
	nd := &VerifyCollection{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 6),
	}
	return nd
}

// SetMetadataAccount sets the "metadata" account.
// Metadata account
func (inst *VerifyCollection) SetMetadataAccount(metadata ag_solanago.PublicKey) *VerifyCollection {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(metadata).WRITE()
	return inst
}

// GetMetadataAccount gets the "metadata" account.
// Metadata account
func (inst *VerifyCollection) GetMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetCollectionUpdateAuthorityAccount sets the "collectionUpdateAuthority" account.
// Collection Update authority
func (inst *VerifyCollection) SetCollectionUpdateAuthorityAccount(collectionUpdateAuthority ag_solanago.PublicKey) *VerifyCollection {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(collectionUpdateAuthority).SIGNER()
	return inst
}

// GetCollectionUpdateAuthorityAccount gets the "collectionUpdateAuthority" account.
// Collection Update authority
func (inst *VerifyCollection) GetCollectionUpdateAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetPayerAccount sets the "payer" account.
// payer
func (inst *VerifyCollection) SetPayerAccount(payer ag_solanago.PublicKey) *VerifyCollection {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(payer).SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
// payer
func (inst *VerifyCollection) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetCollectionMintAccount sets the "collectionMint" account.
// Mint of the Collection
func (inst *VerifyCollection) SetCollectionMintAccount(collectionMint ag_solanago.PublicKey) *VerifyCollection {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(collectionMint)
	return inst
}

// GetCollectionMintAccount gets the "collectionMint" account.
// Mint of the Collection
func (inst *VerifyCollection) GetCollectionMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetCollectionMetadataAccount sets the "collectionMetadata" account.
// Metadata Account of the Collection
func (inst *VerifyCollection) SetCollectionMetadataAccount(collectionMetadata ag_solanago.PublicKey) *VerifyCollection {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(collectionMetadata)
	return inst
}

// GetCollectionMetadataAccount gets the "collectionMetadata" account.
// Metadata Account of the Collection
func (inst *VerifyCollection) GetCollectionMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetMasterEditionV2Account sets the "masterEditionV2" account.
// MasterEdition2 Account of the Collection Token
func (inst *VerifyCollection) SetMasterEditionV2Account(masterEditionV2 ag_solanago.PublicKey) *VerifyCollection {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(masterEditionV2)
	return inst
}

// GetMasterEditionV2Account gets the "masterEditionV2" account.
// MasterEdition2 Account of the Collection Token
func (inst *VerifyCollection) GetMasterEditionV2Account() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

func (inst VerifyCollection) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_VerifyCollection),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst VerifyCollection) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *VerifyCollection) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Metadata is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.CollectionUpdateAuthority is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.CollectionMint is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.CollectionMetadata is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.MasterEditionV2 is not set")
		}
	}
	return nil
}

func (inst *VerifyCollection) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("VerifyCollection")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=6]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                 metadata", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("collectionUpdateAuthority", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("                    payer", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("           collectionMint", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("       collectionMetadata", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("          masterEditionV2", inst.AccountMetaSlice.Get(5)))
					})
				})
		})
}

func (obj VerifyCollection) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *VerifyCollection) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewVerifyCollectionInstruction declares a new VerifyCollection instruction with the provided parameters and accounts.
func NewVerifyCollectionInstruction(
	// Accounts:
	metadata ag_solanago.PublicKey,
	collectionUpdateAuthority ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	collectionMint ag_solanago.PublicKey,
	collectionMetadata ag_solanago.PublicKey,
	masterEditionV2 ag_solanago.PublicKey) *VerifyCollection {
	return NewVerifyCollectionInstructionBuilder().
		SetMetadataAccount(metadata).
		SetCollectionUpdateAuthorityAccount(collectionUpdateAuthority).
		SetPayerAccount(payer).
		SetCollectionMintAccount(collectionMint).
		SetCollectionMetadataAccount(collectionMetadata).
		SetMasterEditionV2Account(masterEditionV2)
}
