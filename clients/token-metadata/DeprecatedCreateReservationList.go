// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package token_metadata

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/desperatee/solana-go"
	ag_format "github.com/desperatee/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Create an empty reservation list for a resource who can come back later as a signer and fill the reservation list
// with reservations to ensure that people who come to get editions get the number they expect. See SetReservationList for more.
type DeprecatedCreateReservationList struct {

	// [0] = [WRITE] pdaForReservationlist
	// ··········· PDA for ReservationList of ['metadata', program id, master edition key, 'reservation', resource-key]
	//
	// [1] = [SIGNER] payer
	// ··········· Payer
	//
	// [2] = [SIGNER] updateAuthority
	// ··········· Update authority
	//
	// [3] = [] masterEditionV1
	// ··········· Master Edition V1 key (pda of ['metadata', program id, mint id, 'edition'])
	//
	// [4] = [] resource
	// ··········· A resource you wish to tie the reservation list to. This is so your later visitors who come to
	// ··········· redeem can derive your reservation list PDA with something they can easily get at. You choose what this should be.
	//
	// [5] = [] metadataKeyPDA
	// ··········· Metadata key (pda of ['metadata', program id, mint id])
	//
	// [6] = [] system
	// ··········· System program
	//
	// [7] = [] rent
	// ··········· Rent info
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewDeprecatedCreateReservationListInstructionBuilder creates a new `DeprecatedCreateReservationList` instruction builder.
func NewDeprecatedCreateReservationListInstructionBuilder() *DeprecatedCreateReservationList {
	nd := &DeprecatedCreateReservationList{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 8),
	}
	return nd
}

// SetPdaForReservationlistAccount sets the "pdaForReservationlist" account.
// PDA for ReservationList of ['metadata', program id, master edition key, 'reservation', resource-key]
func (inst *DeprecatedCreateReservationList) SetPdaForReservationlistAccount(pdaForReservationlist ag_solanago.PublicKey) *DeprecatedCreateReservationList {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(pdaForReservationlist).WRITE()
	return inst
}

// GetPdaForReservationlistAccount gets the "pdaForReservationlist" account.
// PDA for ReservationList of ['metadata', program id, master edition key, 'reservation', resource-key]
func (inst *DeprecatedCreateReservationList) GetPdaForReservationlistAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetPayerAccount sets the "payer" account.
// Payer
func (inst *DeprecatedCreateReservationList) SetPayerAccount(payer ag_solanago.PublicKey) *DeprecatedCreateReservationList {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(payer).SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
// Payer
func (inst *DeprecatedCreateReservationList) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetUpdateAuthorityAccount sets the "updateAuthority" account.
// Update authority
func (inst *DeprecatedCreateReservationList) SetUpdateAuthorityAccount(updateAuthority ag_solanago.PublicKey) *DeprecatedCreateReservationList {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(updateAuthority).SIGNER()
	return inst
}

// GetUpdateAuthorityAccount gets the "updateAuthority" account.
// Update authority
func (inst *DeprecatedCreateReservationList) GetUpdateAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetMasterEditionV1Account sets the "masterEditionV1" account.
// Master Edition V1 key (pda of ['metadata', program id, mint id, 'edition'])
func (inst *DeprecatedCreateReservationList) SetMasterEditionV1Account(masterEditionV1 ag_solanago.PublicKey) *DeprecatedCreateReservationList {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(masterEditionV1)
	return inst
}

// GetMasterEditionV1Account gets the "masterEditionV1" account.
// Master Edition V1 key (pda of ['metadata', program id, mint id, 'edition'])
func (inst *DeprecatedCreateReservationList) GetMasterEditionV1Account() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetResourceAccount sets the "resource" account.
// A resource you wish to tie the reservation list to. This is so your later visitors who come to
// redeem can derive your reservation list PDA with something they can easily get at. You choose what this should be.
func (inst *DeprecatedCreateReservationList) SetResourceAccount(resource ag_solanago.PublicKey) *DeprecatedCreateReservationList {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(resource)
	return inst
}

// GetResourceAccount gets the "resource" account.
// A resource you wish to tie the reservation list to. This is so your later visitors who come to
// redeem can derive your reservation list PDA with something they can easily get at. You choose what this should be.
func (inst *DeprecatedCreateReservationList) GetResourceAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetMetadataKeyPDAAccount sets the "metadataKeyPDA" account.
// Metadata key (pda of ['metadata', program id, mint id])
func (inst *DeprecatedCreateReservationList) SetMetadataKeyPDAAccount(metadataKeyPDA ag_solanago.PublicKey) *DeprecatedCreateReservationList {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(metadataKeyPDA)
	return inst
}

// GetMetadataKeyPDAAccount gets the "metadataKeyPDA" account.
// Metadata key (pda of ['metadata', program id, mint id])
func (inst *DeprecatedCreateReservationList) GetMetadataKeyPDAAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetSystemAccount sets the "system" account.
// System program
func (inst *DeprecatedCreateReservationList) SetSystemAccount(system ag_solanago.PublicKey) *DeprecatedCreateReservationList {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(system)
	return inst
}

// GetSystemAccount gets the "system" account.
// System program
func (inst *DeprecatedCreateReservationList) GetSystemAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetRentAccount sets the "rent" account.
// Rent info
func (inst *DeprecatedCreateReservationList) SetRentAccount(rent ag_solanago.PublicKey) *DeprecatedCreateReservationList {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
// Rent info
func (inst *DeprecatedCreateReservationList) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

func (inst DeprecatedCreateReservationList) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_DeprecatedCreateReservationList),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst DeprecatedCreateReservationList) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *DeprecatedCreateReservationList) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.PdaForReservationlist is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.UpdateAuthority is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.MasterEditionV1 is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Resource is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.MetadataKeyPDA is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.System is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.Rent is not set")
		}
	}
	return nil
}

func (inst *DeprecatedCreateReservationList) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("DeprecatedCreateReservationList")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=8]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("pdaForReservationlist", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("                payer", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("      updateAuthority", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("      masterEditionV1", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("             resource", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("       metadataKeyPDA", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("               system", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("                 rent", inst.AccountMetaSlice.Get(7)))
					})
				})
		})
}

func (obj DeprecatedCreateReservationList) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *DeprecatedCreateReservationList) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewDeprecatedCreateReservationListInstruction declares a new DeprecatedCreateReservationList instruction with the provided parameters and accounts.
func NewDeprecatedCreateReservationListInstruction(
	// Accounts:
	pdaForReservationlist ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	updateAuthority ag_solanago.PublicKey,
	masterEditionV1 ag_solanago.PublicKey,
	resource ag_solanago.PublicKey,
	metadataKeyPDA ag_solanago.PublicKey,
	system ag_solanago.PublicKey,
	rent ag_solanago.PublicKey) *DeprecatedCreateReservationList {
	return NewDeprecatedCreateReservationListInstructionBuilder().
		SetPdaForReservationlistAccount(pdaForReservationlist).
		SetPayerAccount(payer).
		SetUpdateAuthorityAccount(updateAuthority).
		SetMasterEditionV1Account(masterEditionV1).
		SetResourceAccount(resource).
		SetMetadataKeyPDAAccount(metadataKeyPDA).
		SetSystemAccount(system).
		SetRentAccount(rent)
}
