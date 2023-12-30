// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package candy_machine_core

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/RoboticAgile/solana-go"
	ag_format "github.com/RoboticAgile/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Initialize is the `initialize` instruction.
type Initialize struct {
	Data *CandyMachineData

	// [0] = [WRITE] candyMachine
	//
	// [1] = [WRITE] authorityPda
	//
	// [2] = [] authority
	//
	// [3] = [SIGNER] payer
	//
	// [4] = [] collectionMetadata
	//
	// [5] = [] collectionMint
	//
	// [6] = [] collectionMasterEdition
	//
	// [7] = [WRITE, SIGNER] collectionUpdateAuthority
	//
	// [8] = [WRITE] collectionAuthorityRecord
	//
	// [9] = [] tokenMetadataProgram
	//
	// [10] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewInitializeInstructionBuilder creates a new `Initialize` instruction builder.
func NewInitializeInstructionBuilder() *Initialize {
	nd := &Initialize{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 11),
	}
	return nd
}

// SetData sets the "data" parameter.
func (inst *Initialize) SetData(data CandyMachineData) *Initialize {
	inst.Data = &data
	return inst
}

// SetCandyMachineAccount sets the "candyMachine" account.
func (inst *Initialize) SetCandyMachineAccount(candyMachine ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(candyMachine).WRITE()
	return inst
}

// GetCandyMachineAccount gets the "candyMachine" account.
func (inst *Initialize) GetCandyMachineAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAuthorityPdaAccount sets the "authorityPda" account.
func (inst *Initialize) SetAuthorityPdaAccount(authorityPda ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(authorityPda).WRITE()
	return inst
}

// GetAuthorityPdaAccount gets the "authorityPda" account.
func (inst *Initialize) GetAuthorityPdaAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetAuthorityAccount sets the "authority" account.
func (inst *Initialize) SetAuthorityAccount(authority ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(authority)
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *Initialize) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetPayerAccount sets the "payer" account.
func (inst *Initialize) SetPayerAccount(payer ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(payer).SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
func (inst *Initialize) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetCollectionMetadataAccount sets the "collectionMetadata" account.
func (inst *Initialize) SetCollectionMetadataAccount(collectionMetadata ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(collectionMetadata)
	return inst
}

// GetCollectionMetadataAccount gets the "collectionMetadata" account.
func (inst *Initialize) GetCollectionMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetCollectionMintAccount sets the "collectionMint" account.
func (inst *Initialize) SetCollectionMintAccount(collectionMint ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(collectionMint)
	return inst
}

// GetCollectionMintAccount gets the "collectionMint" account.
func (inst *Initialize) GetCollectionMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetCollectionMasterEditionAccount sets the "collectionMasterEdition" account.
func (inst *Initialize) SetCollectionMasterEditionAccount(collectionMasterEdition ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(collectionMasterEdition)
	return inst
}

// GetCollectionMasterEditionAccount gets the "collectionMasterEdition" account.
func (inst *Initialize) GetCollectionMasterEditionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetCollectionUpdateAuthorityAccount sets the "collectionUpdateAuthority" account.
func (inst *Initialize) SetCollectionUpdateAuthorityAccount(collectionUpdateAuthority ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(collectionUpdateAuthority).WRITE().SIGNER()
	return inst
}

// GetCollectionUpdateAuthorityAccount gets the "collectionUpdateAuthority" account.
func (inst *Initialize) GetCollectionUpdateAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetCollectionAuthorityRecordAccount sets the "collectionAuthorityRecord" account.
func (inst *Initialize) SetCollectionAuthorityRecordAccount(collectionAuthorityRecord ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(collectionAuthorityRecord).WRITE()
	return inst
}

// GetCollectionAuthorityRecordAccount gets the "collectionAuthorityRecord" account.
func (inst *Initialize) GetCollectionAuthorityRecordAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetTokenMetadataProgramAccount sets the "tokenMetadataProgram" account.
func (inst *Initialize) SetTokenMetadataProgramAccount(tokenMetadataProgram ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(tokenMetadataProgram)
	return inst
}

// GetTokenMetadataProgramAccount gets the "tokenMetadataProgram" account.
func (inst *Initialize) GetTokenMetadataProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *Initialize) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *Initialize) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

func (inst Initialize) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_Initialize,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst Initialize) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *Initialize) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Data == nil {
			return errors.New("Data parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.CandyMachine is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.AuthorityPda is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.CollectionMetadata is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.CollectionMint is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.CollectionMasterEdition is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.CollectionUpdateAuthority is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.CollectionAuthorityRecord is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.TokenMetadataProgram is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *Initialize) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("Initialize")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Data", *inst.Data))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=11]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("             candyMachine", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("             authorityPda", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("                authority", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("                    payer", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("       collectionMetadata", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("           collectionMint", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("  collectionMasterEdition", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("collectionUpdateAuthority", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("collectionAuthorityRecord", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("     tokenMetadataProgram", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("            systemProgram", inst.AccountMetaSlice.Get(10)))
					})
				})
		})
}

func (obj Initialize) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Data` param:
	err = encoder.Encode(obj.Data)
	if err != nil {
		return err
	}
	return nil
}
func (obj *Initialize) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Data`:
	err = decoder.Decode(&obj.Data)
	if err != nil {
		return err
	}
	return nil
}

// NewInitializeInstruction declares a new Initialize instruction with the provided parameters and accounts.
func NewInitializeInstruction(
	// Parameters:
	data CandyMachineData,
	// Accounts:
	candyMachine ag_solanago.PublicKey,
	authorityPda ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	collectionMetadata ag_solanago.PublicKey,
	collectionMint ag_solanago.PublicKey,
	collectionMasterEdition ag_solanago.PublicKey,
	collectionUpdateAuthority ag_solanago.PublicKey,
	collectionAuthorityRecord ag_solanago.PublicKey,
	tokenMetadataProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *Initialize {
	return NewInitializeInstructionBuilder().
		SetData(data).
		SetCandyMachineAccount(candyMachine).
		SetAuthorityPdaAccount(authorityPda).
		SetAuthorityAccount(authority).
		SetPayerAccount(payer).
		SetCollectionMetadataAccount(collectionMetadata).
		SetCollectionMintAccount(collectionMint).
		SetCollectionMasterEditionAccount(collectionMasterEdition).
		SetCollectionUpdateAuthorityAccount(collectionUpdateAuthority).
		SetCollectionAuthorityRecordAccount(collectionAuthorityRecord).
		SetTokenMetadataProgramAccount(tokenMetadataProgram).
		SetSystemProgramAccount(systemProgram)
}
