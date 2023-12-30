// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package gumdrop

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/RoboticAgile/solana-go"
	ag_format "github.com/RoboticAgile/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// ClaimEdition is the `claimEdition` instruction.
type ClaimEdition struct {
	ClaimBump      *uint8
	Index          *uint64
	Amount         *uint64
	Edition        *uint64
	ClaimantSecret *ag_solanago.PublicKey
	Proof          *[][32]uint8

	// [0] = [WRITE] distributor
	//
	// [1] = [WRITE] claimCount
	//
	// [2] = [SIGNER] temporal
	//
	// [3] = [SIGNER] payer
	//
	// [4] = [WRITE] metadataNewMetadata
	//
	// [5] = [WRITE] metadataNewEdition
	//
	// [6] = [WRITE] metadataMasterEdition
	//
	// [7] = [WRITE] metadataNewMint
	//
	// [8] = [WRITE] metadataEditionMarkPda
	//
	// [9] = [SIGNER] metadataNewMintAuthority
	//
	// [10] = [] metadataMasterTokenAccount
	//
	// [11] = [] metadataNewUpdateAuthority
	//
	// [12] = [] metadataMasterMetadata
	//
	// [13] = [] metadataMasterMint
	//
	// [14] = [] systemProgram
	//
	// [15] = [] tokenProgram
	//
	// [16] = [] tokenMetadataProgram
	//
	// [17] = [] rent
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewClaimEditionInstructionBuilder creates a new `ClaimEdition` instruction builder.
func NewClaimEditionInstructionBuilder() *ClaimEdition {
	nd := &ClaimEdition{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 18),
	}
	return nd
}

// SetClaimBump sets the "claimBump" parameter.
func (inst *ClaimEdition) SetClaimBump(claimBump uint8) *ClaimEdition {
	inst.ClaimBump = &claimBump
	return inst
}

// SetIndex sets the "index" parameter.
func (inst *ClaimEdition) SetIndex(index uint64) *ClaimEdition {
	inst.Index = &index
	return inst
}

// SetAmount sets the "amount" parameter.
func (inst *ClaimEdition) SetAmount(amount uint64) *ClaimEdition {
	inst.Amount = &amount
	return inst
}

// SetEdition sets the "edition" parameter.
func (inst *ClaimEdition) SetEdition(edition uint64) *ClaimEdition {
	inst.Edition = &edition
	return inst
}

// SetClaimantSecret sets the "claimantSecret" parameter.
func (inst *ClaimEdition) SetClaimantSecret(claimantSecret ag_solanago.PublicKey) *ClaimEdition {
	inst.ClaimantSecret = &claimantSecret
	return inst
}

// SetProof sets the "proof" parameter.
func (inst *ClaimEdition) SetProof(proof [][32]uint8) *ClaimEdition {
	inst.Proof = &proof
	return inst
}

// SetDistributorAccount sets the "distributor" account.
func (inst *ClaimEdition) SetDistributorAccount(distributor ag_solanago.PublicKey) *ClaimEdition {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(distributor).WRITE()
	return inst
}

// GetDistributorAccount gets the "distributor" account.
func (inst *ClaimEdition) GetDistributorAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetClaimCountAccount sets the "claimCount" account.
func (inst *ClaimEdition) SetClaimCountAccount(claimCount ag_solanago.PublicKey) *ClaimEdition {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(claimCount).WRITE()
	return inst
}

// GetClaimCountAccount gets the "claimCount" account.
func (inst *ClaimEdition) GetClaimCountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetTemporalAccount sets the "temporal" account.
func (inst *ClaimEdition) SetTemporalAccount(temporal ag_solanago.PublicKey) *ClaimEdition {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(temporal).SIGNER()
	return inst
}

// GetTemporalAccount gets the "temporal" account.
func (inst *ClaimEdition) GetTemporalAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetPayerAccount sets the "payer" account.
func (inst *ClaimEdition) SetPayerAccount(payer ag_solanago.PublicKey) *ClaimEdition {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(payer).SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
func (inst *ClaimEdition) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetMetadataNewMetadataAccount sets the "metadataNewMetadata" account.
func (inst *ClaimEdition) SetMetadataNewMetadataAccount(metadataNewMetadata ag_solanago.PublicKey) *ClaimEdition {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(metadataNewMetadata).WRITE()
	return inst
}

// GetMetadataNewMetadataAccount gets the "metadataNewMetadata" account.
func (inst *ClaimEdition) GetMetadataNewMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetMetadataNewEditionAccount sets the "metadataNewEdition" account.
func (inst *ClaimEdition) SetMetadataNewEditionAccount(metadataNewEdition ag_solanago.PublicKey) *ClaimEdition {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(metadataNewEdition).WRITE()
	return inst
}

// GetMetadataNewEditionAccount gets the "metadataNewEdition" account.
func (inst *ClaimEdition) GetMetadataNewEditionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetMetadataMasterEditionAccount sets the "metadataMasterEdition" account.
func (inst *ClaimEdition) SetMetadataMasterEditionAccount(metadataMasterEdition ag_solanago.PublicKey) *ClaimEdition {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(metadataMasterEdition).WRITE()
	return inst
}

// GetMetadataMasterEditionAccount gets the "metadataMasterEdition" account.
func (inst *ClaimEdition) GetMetadataMasterEditionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetMetadataNewMintAccount sets the "metadataNewMint" account.
func (inst *ClaimEdition) SetMetadataNewMintAccount(metadataNewMint ag_solanago.PublicKey) *ClaimEdition {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(metadataNewMint).WRITE()
	return inst
}

// GetMetadataNewMintAccount gets the "metadataNewMint" account.
func (inst *ClaimEdition) GetMetadataNewMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetMetadataEditionMarkPdaAccount sets the "metadataEditionMarkPda" account.
func (inst *ClaimEdition) SetMetadataEditionMarkPdaAccount(metadataEditionMarkPda ag_solanago.PublicKey) *ClaimEdition {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(metadataEditionMarkPda).WRITE()
	return inst
}

// GetMetadataEditionMarkPdaAccount gets the "metadataEditionMarkPda" account.
func (inst *ClaimEdition) GetMetadataEditionMarkPdaAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetMetadataNewMintAuthorityAccount sets the "metadataNewMintAuthority" account.
func (inst *ClaimEdition) SetMetadataNewMintAuthorityAccount(metadataNewMintAuthority ag_solanago.PublicKey) *ClaimEdition {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(metadataNewMintAuthority).SIGNER()
	return inst
}

// GetMetadataNewMintAuthorityAccount gets the "metadataNewMintAuthority" account.
func (inst *ClaimEdition) GetMetadataNewMintAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetMetadataMasterTokenAccountAccount sets the "metadataMasterTokenAccount" account.
func (inst *ClaimEdition) SetMetadataMasterTokenAccountAccount(metadataMasterTokenAccount ag_solanago.PublicKey) *ClaimEdition {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(metadataMasterTokenAccount)
	return inst
}

// GetMetadataMasterTokenAccountAccount gets the "metadataMasterTokenAccount" account.
func (inst *ClaimEdition) GetMetadataMasterTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetMetadataNewUpdateAuthorityAccount sets the "metadataNewUpdateAuthority" account.
func (inst *ClaimEdition) SetMetadataNewUpdateAuthorityAccount(metadataNewUpdateAuthority ag_solanago.PublicKey) *ClaimEdition {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(metadataNewUpdateAuthority)
	return inst
}

// GetMetadataNewUpdateAuthorityAccount gets the "metadataNewUpdateAuthority" account.
func (inst *ClaimEdition) GetMetadataNewUpdateAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetMetadataMasterMetadataAccount sets the "metadataMasterMetadata" account.
func (inst *ClaimEdition) SetMetadataMasterMetadataAccount(metadataMasterMetadata ag_solanago.PublicKey) *ClaimEdition {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(metadataMasterMetadata)
	return inst
}

// GetMetadataMasterMetadataAccount gets the "metadataMasterMetadata" account.
func (inst *ClaimEdition) GetMetadataMasterMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

// SetMetadataMasterMintAccount sets the "metadataMasterMint" account.
func (inst *ClaimEdition) SetMetadataMasterMintAccount(metadataMasterMint ag_solanago.PublicKey) *ClaimEdition {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(metadataMasterMint)
	return inst
}

// GetMetadataMasterMintAccount gets the "metadataMasterMint" account.
func (inst *ClaimEdition) GetMetadataMasterMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(13)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *ClaimEdition) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *ClaimEdition {
	inst.AccountMetaSlice[14] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *ClaimEdition) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(14)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *ClaimEdition) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *ClaimEdition {
	inst.AccountMetaSlice[15] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *ClaimEdition) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(15)
}

// SetTokenMetadataProgramAccount sets the "tokenMetadataProgram" account.
func (inst *ClaimEdition) SetTokenMetadataProgramAccount(tokenMetadataProgram ag_solanago.PublicKey) *ClaimEdition {
	inst.AccountMetaSlice[16] = ag_solanago.Meta(tokenMetadataProgram)
	return inst
}

// GetTokenMetadataProgramAccount gets the "tokenMetadataProgram" account.
func (inst *ClaimEdition) GetTokenMetadataProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(16)
}

// SetRentAccount sets the "rent" account.
func (inst *ClaimEdition) SetRentAccount(rent ag_solanago.PublicKey) *ClaimEdition {
	inst.AccountMetaSlice[17] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *ClaimEdition) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(17)
}

func (inst ClaimEdition) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_ClaimEdition,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst ClaimEdition) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *ClaimEdition) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.ClaimBump == nil {
			return errors.New("ClaimBump parameter is not set")
		}
		if inst.Index == nil {
			return errors.New("Index parameter is not set")
		}
		if inst.Amount == nil {
			return errors.New("Amount parameter is not set")
		}
		if inst.Edition == nil {
			return errors.New("Edition parameter is not set")
		}
		if inst.ClaimantSecret == nil {
			return errors.New("ClaimantSecret parameter is not set")
		}
		if inst.Proof == nil {
			return errors.New("Proof parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Distributor is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.ClaimCount is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Temporal is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.MetadataNewMetadata is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.MetadataNewEdition is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.MetadataMasterEdition is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.MetadataNewMint is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.MetadataEditionMarkPda is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.MetadataNewMintAuthority is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.MetadataMasterTokenAccount is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.MetadataNewUpdateAuthority is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.MetadataMasterMetadata is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.MetadataMasterMint is not set")
		}
		if inst.AccountMetaSlice[14] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[15] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[16] == nil {
			return errors.New("accounts.TokenMetadataProgram is not set")
		}
		if inst.AccountMetaSlice[17] == nil {
			return errors.New("accounts.Rent is not set")
		}
	}
	return nil
}

func (inst *ClaimEdition) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("ClaimEdition")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=6]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("     ClaimBump", *inst.ClaimBump))
						paramsBranch.Child(ag_format.Param("         Index", *inst.Index))
						paramsBranch.Child(ag_format.Param("        Amount", *inst.Amount))
						paramsBranch.Child(ag_format.Param("       Edition", *inst.Edition))
						paramsBranch.Child(ag_format.Param("ClaimantSecret", *inst.ClaimantSecret))
						paramsBranch.Child(ag_format.Param("         Proof", *inst.Proof))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=18]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("               distributor", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("                claimCount", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("                  temporal", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("                     payer", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("       metadataNewMetadata", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("        metadataNewEdition", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("     metadataMasterEdition", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("           metadataNewMint", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("    metadataEditionMarkPda", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("  metadataNewMintAuthority", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("       metadataMasterToken", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("metadataNewUpdateAuthority", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("    metadataMasterMetadata", inst.AccountMetaSlice.Get(12)))
						accountsBranch.Child(ag_format.Meta("        metadataMasterMint", inst.AccountMetaSlice.Get(13)))
						accountsBranch.Child(ag_format.Meta("             systemProgram", inst.AccountMetaSlice.Get(14)))
						accountsBranch.Child(ag_format.Meta("              tokenProgram", inst.AccountMetaSlice.Get(15)))
						accountsBranch.Child(ag_format.Meta("      tokenMetadataProgram", inst.AccountMetaSlice.Get(16)))
						accountsBranch.Child(ag_format.Meta("                      rent", inst.AccountMetaSlice.Get(17)))
					})
				})
		})
}

func (obj ClaimEdition) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `ClaimBump` param:
	err = encoder.Encode(obj.ClaimBump)
	if err != nil {
		return err
	}
	// Serialize `Index` param:
	err = encoder.Encode(obj.Index)
	if err != nil {
		return err
	}
	// Serialize `Amount` param:
	err = encoder.Encode(obj.Amount)
	if err != nil {
		return err
	}
	// Serialize `Edition` param:
	err = encoder.Encode(obj.Edition)
	if err != nil {
		return err
	}
	// Serialize `ClaimantSecret` param:
	err = encoder.Encode(obj.ClaimantSecret)
	if err != nil {
		return err
	}
	// Serialize `Proof` param:
	err = encoder.Encode(obj.Proof)
	if err != nil {
		return err
	}
	return nil
}
func (obj *ClaimEdition) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `ClaimBump`:
	err = decoder.Decode(&obj.ClaimBump)
	if err != nil {
		return err
	}
	// Deserialize `Index`:
	err = decoder.Decode(&obj.Index)
	if err != nil {
		return err
	}
	// Deserialize `Amount`:
	err = decoder.Decode(&obj.Amount)
	if err != nil {
		return err
	}
	// Deserialize `Edition`:
	err = decoder.Decode(&obj.Edition)
	if err != nil {
		return err
	}
	// Deserialize `ClaimantSecret`:
	err = decoder.Decode(&obj.ClaimantSecret)
	if err != nil {
		return err
	}
	// Deserialize `Proof`:
	err = decoder.Decode(&obj.Proof)
	if err != nil {
		return err
	}
	return nil
}

// NewClaimEditionInstruction declares a new ClaimEdition instruction with the provided parameters and accounts.
func NewClaimEditionInstruction(
	// Parameters:
	claimBump uint8,
	index uint64,
	amount uint64,
	edition uint64,
	claimantSecret ag_solanago.PublicKey,
	proof [][32]uint8,
	// Accounts:
	distributor ag_solanago.PublicKey,
	claimCount ag_solanago.PublicKey,
	temporal ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	metadataNewMetadata ag_solanago.PublicKey,
	metadataNewEdition ag_solanago.PublicKey,
	metadataMasterEdition ag_solanago.PublicKey,
	metadataNewMint ag_solanago.PublicKey,
	metadataEditionMarkPda ag_solanago.PublicKey,
	metadataNewMintAuthority ag_solanago.PublicKey,
	metadataMasterTokenAccount ag_solanago.PublicKey,
	metadataNewUpdateAuthority ag_solanago.PublicKey,
	metadataMasterMetadata ag_solanago.PublicKey,
	metadataMasterMint ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	tokenMetadataProgram ag_solanago.PublicKey,
	rent ag_solanago.PublicKey) *ClaimEdition {
	return NewClaimEditionInstructionBuilder().
		SetClaimBump(claimBump).
		SetIndex(index).
		SetAmount(amount).
		SetEdition(edition).
		SetClaimantSecret(claimantSecret).
		SetProof(proof).
		SetDistributorAccount(distributor).
		SetClaimCountAccount(claimCount).
		SetTemporalAccount(temporal).
		SetPayerAccount(payer).
		SetMetadataNewMetadataAccount(metadataNewMetadata).
		SetMetadataNewEditionAccount(metadataNewEdition).
		SetMetadataMasterEditionAccount(metadataMasterEdition).
		SetMetadataNewMintAccount(metadataNewMint).
		SetMetadataEditionMarkPdaAccount(metadataEditionMarkPda).
		SetMetadataNewMintAuthorityAccount(metadataNewMintAuthority).
		SetMetadataMasterTokenAccountAccount(metadataMasterTokenAccount).
		SetMetadataNewUpdateAuthorityAccount(metadataNewUpdateAuthority).
		SetMetadataMasterMetadataAccount(metadataMasterMetadata).
		SetMetadataMasterMintAccount(metadataMasterMint).
		SetSystemProgramAccount(systemProgram).
		SetTokenProgramAccount(tokenProgram).
		SetTokenMetadataProgramAccount(tokenMetadataProgram).
		SetRentAccount(rent)
}
