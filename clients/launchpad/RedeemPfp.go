// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package launchpad

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/RoboticAgile/solana-go"
	ag_format "github.com/RoboticAgile/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// RedeemPfp is the `redeemPfp` instruction.
type RedeemPfp struct {

	// [0] = [WRITE, SIGNER] initializer
	//
	// [1] = [WRITE] collection
	//
	// [2] = [] nftMint
	//
	// [3] = [WRITE] initializerNftTa
	//
	// [4] = [WRITE] collectionNftTa
	//
	// [5] = [] pfpMint
	//
	// [6] = [WRITE] initializerPfpTa
	//
	// [7] = [WRITE] collectionPfpTa
	//
	// [8] = [WRITE] nftMetadata
	//
	// [9] = [WRITE] nftMasterEdition
	//
	// [10] = [] mplTokenMetadata
	//
	// [11] = [] tokenProgram
	//
	// [12] = [] associatedTokenProgram
	//
	// [13] = [] systemProgram
	//
	// [14] = [] rent
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewRedeemPfpInstructionBuilder creates a new `RedeemPfp` instruction builder.
func NewRedeemPfpInstructionBuilder() *RedeemPfp {
	nd := &RedeemPfp{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 15),
	}
	return nd
}

// SetInitializerAccount sets the "initializer" account.
func (inst *RedeemPfp) SetInitializerAccount(initializer ag_solanago.PublicKey) *RedeemPfp {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(initializer).WRITE().SIGNER()
	return inst
}

// GetInitializerAccount gets the "initializer" account.
func (inst *RedeemPfp) GetInitializerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetCollectionAccount sets the "collection" account.
func (inst *RedeemPfp) SetCollectionAccount(collection ag_solanago.PublicKey) *RedeemPfp {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(collection).WRITE()
	return inst
}

// GetCollectionAccount gets the "collection" account.
func (inst *RedeemPfp) GetCollectionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetNftMintAccount sets the "nftMint" account.
func (inst *RedeemPfp) SetNftMintAccount(nftMint ag_solanago.PublicKey) *RedeemPfp {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(nftMint)
	return inst
}

// GetNftMintAccount gets the "nftMint" account.
func (inst *RedeemPfp) GetNftMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetInitializerNftTaAccount sets the "initializerNftTa" account.
func (inst *RedeemPfp) SetInitializerNftTaAccount(initializerNftTa ag_solanago.PublicKey) *RedeemPfp {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(initializerNftTa).WRITE()
	return inst
}

// GetInitializerNftTaAccount gets the "initializerNftTa" account.
func (inst *RedeemPfp) GetInitializerNftTaAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetCollectionNftTaAccount sets the "collectionNftTa" account.
func (inst *RedeemPfp) SetCollectionNftTaAccount(collectionNftTa ag_solanago.PublicKey) *RedeemPfp {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(collectionNftTa).WRITE()
	return inst
}

// GetCollectionNftTaAccount gets the "collectionNftTa" account.
func (inst *RedeemPfp) GetCollectionNftTaAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetPfpMintAccount sets the "pfpMint" account.
func (inst *RedeemPfp) SetPfpMintAccount(pfpMint ag_solanago.PublicKey) *RedeemPfp {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(pfpMint)
	return inst
}

// GetPfpMintAccount gets the "pfpMint" account.
func (inst *RedeemPfp) GetPfpMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetInitializerPfpTaAccount sets the "initializerPfpTa" account.
func (inst *RedeemPfp) SetInitializerPfpTaAccount(initializerPfpTa ag_solanago.PublicKey) *RedeemPfp {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(initializerPfpTa).WRITE()
	return inst
}

// GetInitializerPfpTaAccount gets the "initializerPfpTa" account.
func (inst *RedeemPfp) GetInitializerPfpTaAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetCollectionPfpTaAccount sets the "collectionPfpTa" account.
func (inst *RedeemPfp) SetCollectionPfpTaAccount(collectionPfpTa ag_solanago.PublicKey) *RedeemPfp {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(collectionPfpTa).WRITE()
	return inst
}

// GetCollectionPfpTaAccount gets the "collectionPfpTa" account.
func (inst *RedeemPfp) GetCollectionPfpTaAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetNftMetadataAccount sets the "nftMetadata" account.
func (inst *RedeemPfp) SetNftMetadataAccount(nftMetadata ag_solanago.PublicKey) *RedeemPfp {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(nftMetadata).WRITE()
	return inst
}

// GetNftMetadataAccount gets the "nftMetadata" account.
func (inst *RedeemPfp) GetNftMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetNftMasterEditionAccount sets the "nftMasterEdition" account.
func (inst *RedeemPfp) SetNftMasterEditionAccount(nftMasterEdition ag_solanago.PublicKey) *RedeemPfp {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(nftMasterEdition).WRITE()
	return inst
}

// GetNftMasterEditionAccount gets the "nftMasterEdition" account.
func (inst *RedeemPfp) GetNftMasterEditionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetMplTokenMetadataAccount sets the "mplTokenMetadata" account.
func (inst *RedeemPfp) SetMplTokenMetadataAccount(mplTokenMetadata ag_solanago.PublicKey) *RedeemPfp {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(mplTokenMetadata)
	return inst
}

// GetMplTokenMetadataAccount gets the "mplTokenMetadata" account.
func (inst *RedeemPfp) GetMplTokenMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *RedeemPfp) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *RedeemPfp {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *RedeemPfp) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetAssociatedTokenProgramAccount sets the "associatedTokenProgram" account.
func (inst *RedeemPfp) SetAssociatedTokenProgramAccount(associatedTokenProgram ag_solanago.PublicKey) *RedeemPfp {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(associatedTokenProgram)
	return inst
}

// GetAssociatedTokenProgramAccount gets the "associatedTokenProgram" account.
func (inst *RedeemPfp) GetAssociatedTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *RedeemPfp) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *RedeemPfp {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *RedeemPfp) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(13)
}

// SetRentAccount sets the "rent" account.
func (inst *RedeemPfp) SetRentAccount(rent ag_solanago.PublicKey) *RedeemPfp {
	inst.AccountMetaSlice[14] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *RedeemPfp) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(14)
}

func (inst RedeemPfp) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_RedeemPfp,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst RedeemPfp) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *RedeemPfp) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Initializer is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Collection is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.NftMint is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.InitializerNftTa is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.CollectionNftTa is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.PfpMint is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.InitializerPfpTa is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.CollectionPfpTa is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.NftMetadata is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.NftMasterEdition is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.MplTokenMetadata is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.AssociatedTokenProgram is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[14] == nil {
			return errors.New("accounts.Rent is not set")
		}
	}
	return nil
}

func (inst *RedeemPfp) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("RedeemPfp")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=15]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("           initializer", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("            collection", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("               nftMint", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("      initializerNftTa", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("       collectionNftTa", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("               pfpMint", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("      initializerPfpTa", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("       collectionPfpTa", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("           nftMetadata", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("      nftMasterEdition", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("      mplTokenMetadata", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("          tokenProgram", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("associatedTokenProgram", inst.AccountMetaSlice.Get(12)))
						accountsBranch.Child(ag_format.Meta("         systemProgram", inst.AccountMetaSlice.Get(13)))
						accountsBranch.Child(ag_format.Meta("                  rent", inst.AccountMetaSlice.Get(14)))
					})
				})
		})
}

func (obj RedeemPfp) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *RedeemPfp) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewRedeemPfpInstruction declares a new RedeemPfp instruction with the provided parameters and accounts.
func NewRedeemPfpInstruction(
	// Accounts:
	initializer ag_solanago.PublicKey,
	collection ag_solanago.PublicKey,
	nftMint ag_solanago.PublicKey,
	initializerNftTa ag_solanago.PublicKey,
	collectionNftTa ag_solanago.PublicKey,
	pfpMint ag_solanago.PublicKey,
	initializerPfpTa ag_solanago.PublicKey,
	collectionPfpTa ag_solanago.PublicKey,
	nftMetadata ag_solanago.PublicKey,
	nftMasterEdition ag_solanago.PublicKey,
	mplTokenMetadata ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	associatedTokenProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	rent ag_solanago.PublicKey) *RedeemPfp {
	return NewRedeemPfpInstructionBuilder().
		SetInitializerAccount(initializer).
		SetCollectionAccount(collection).
		SetNftMintAccount(nftMint).
		SetInitializerNftTaAccount(initializerNftTa).
		SetCollectionNftTaAccount(collectionNftTa).
		SetPfpMintAccount(pfpMint).
		SetInitializerPfpTaAccount(initializerPfpTa).
		SetCollectionPfpTaAccount(collectionPfpTa).
		SetNftMetadataAccount(nftMetadata).
		SetNftMasterEditionAccount(nftMasterEdition).
		SetMplTokenMetadataAccount(mplTokenMetadata).
		SetTokenProgramAccount(tokenProgram).
		SetAssociatedTokenProgramAccount(associatedTokenProgram).
		SetSystemProgramAccount(systemProgram).
		SetRentAccount(rent)
}
