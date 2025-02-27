// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package wlmarket

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/RoboticAgile/solana-go"
	ag_format "github.com/RoboticAgile/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// CancelOrder is the `cancelOrder` instruction.
type CancelOrder struct {

	// [0] = [WRITE] order
	//
	// [1] = [] mint
	//
	// [2] = [WRITE, SIGNER] owner
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewCancelOrderInstructionBuilder creates a new `CancelOrder` instruction builder.
func NewCancelOrderInstructionBuilder() *CancelOrder {
	nd := &CancelOrder{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetOrderAccount sets the "order" account.
func (inst *CancelOrder) SetOrderAccount(order ag_solanago.PublicKey) *CancelOrder {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(order).WRITE()
	return inst
}

// GetOrderAccount gets the "order" account.
func (inst *CancelOrder) GetOrderAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetMintAccount sets the "mint" account.
func (inst *CancelOrder) SetMintAccount(mint ag_solanago.PublicKey) *CancelOrder {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(mint)
	return inst
}

// GetMintAccount gets the "mint" account.
func (inst *CancelOrder) GetMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetOwnerAccount sets the "owner" account.
func (inst *CancelOrder) SetOwnerAccount(owner ag_solanago.PublicKey) *CancelOrder {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(owner).WRITE().SIGNER()
	return inst
}

// GetOwnerAccount gets the "owner" account.
func (inst *CancelOrder) GetOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

func (inst CancelOrder) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_CancelOrder,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst CancelOrder) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *CancelOrder) Validate() error {
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
	}
	return nil
}

func (inst *CancelOrder) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("CancelOrder")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("order", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta(" mint", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("owner", inst.AccountMetaSlice.Get(2)))
					})
				})
		})
}

func (obj CancelOrder) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *CancelOrder) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewCancelOrderInstruction declares a new CancelOrder instruction with the provided parameters and accounts.
func NewCancelOrderInstruction(
	// Accounts:
	order ag_solanago.PublicKey,
	mint ag_solanago.PublicKey,
	owner ag_solanago.PublicKey) *CancelOrder {
	return NewCancelOrderInstructionBuilder().
		SetOrderAccount(order).
		SetMintAccount(mint).
		SetOwnerAccount(owner)
}
