// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package wlmarket

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/RoboticAgile/solana-go"
	ag_format "github.com/RoboticAgile/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// CloseOrder is the `closeOrder` instruction.
type CloseOrder struct {

	// [0] = [WRITE] order
	//
	// [1] = [] mint
	//
	// [2] = [WRITE] owner
	//
	// [3] = [SIGNER] signer
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewCloseOrderInstructionBuilder creates a new `CloseOrder` instruction builder.
func NewCloseOrderInstructionBuilder() *CloseOrder {
	nd := &CloseOrder{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 4),
	}
	return nd
}

// SetOrderAccount sets the "order" account.
func (inst *CloseOrder) SetOrderAccount(order ag_solanago.PublicKey) *CloseOrder {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(order).WRITE()
	return inst
}

// GetOrderAccount gets the "order" account.
func (inst *CloseOrder) GetOrderAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetMintAccount sets the "mint" account.
func (inst *CloseOrder) SetMintAccount(mint ag_solanago.PublicKey) *CloseOrder {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(mint)
	return inst
}

// GetMintAccount gets the "mint" account.
func (inst *CloseOrder) GetMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetOwnerAccount sets the "owner" account.
func (inst *CloseOrder) SetOwnerAccount(owner ag_solanago.PublicKey) *CloseOrder {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(owner).WRITE()
	return inst
}

// GetOwnerAccount gets the "owner" account.
func (inst *CloseOrder) GetOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetSignerAccount sets the "signer" account.
func (inst *CloseOrder) SetSignerAccount(signer ag_solanago.PublicKey) *CloseOrder {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(signer).SIGNER()
	return inst
}

// GetSignerAccount gets the "signer" account.
func (inst *CloseOrder) GetSignerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

func (inst CloseOrder) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_CloseOrder,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst CloseOrder) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *CloseOrder) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Order is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Mint is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Owner is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Signer is not set")
		}
	}
	return nil
}

func (inst *CloseOrder) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("CloseOrder")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=4]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta(" order", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("  mint", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta(" owner", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("signer", inst.AccountMetaSlice.Get(3)))
					})
				})
		})
}

func (obj CloseOrder) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *CloseOrder) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewCloseOrderInstruction declares a new CloseOrder instruction with the provided parameters and accounts.
func NewCloseOrderInstruction(
	// Accounts:
	order ag_solanago.PublicKey,
	mint ag_solanago.PublicKey,
	owner ag_solanago.PublicKey,
	signer ag_solanago.PublicKey) *CloseOrder {
	return NewCloseOrderInstructionBuilder().
		SetOrderAccount(order).
		SetMintAccount(mint).
		SetOwnerAccount(owner).
		SetSignerAccount(signer)
}
