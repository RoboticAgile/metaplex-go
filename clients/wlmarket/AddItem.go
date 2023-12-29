// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package wlmarket

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/desperatee/solana-go"
	ag_format "github.com/desperatee/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// AddItem is the `addItem` instruction.
type AddItem struct {
	Bump *uint8
	Foxy *bool

	// [0] = [WRITE] item
	//
	// [1] = [SIGNER] owner
	//
	// [2] = [] mint
	//
	// [3] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewAddItemInstructionBuilder creates a new `AddItem` instruction builder.
func NewAddItemInstructionBuilder() *AddItem {
	nd := &AddItem{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 4),
	}
	return nd
}

// SetBump sets the "bump" parameter.
func (inst *AddItem) SetBump(bump uint8) *AddItem {
	inst.Bump = &bump
	return inst
}

// SetFoxy sets the "foxy" parameter.
func (inst *AddItem) SetFoxy(foxy bool) *AddItem {
	inst.Foxy = &foxy
	return inst
}

// SetItemAccount sets the "item" account.
func (inst *AddItem) SetItemAccount(item ag_solanago.PublicKey) *AddItem {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(item).WRITE()
	return inst
}

// GetItemAccount gets the "item" account.
func (inst *AddItem) GetItemAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetOwnerAccount sets the "owner" account.
func (inst *AddItem) SetOwnerAccount(owner ag_solanago.PublicKey) *AddItem {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(owner).SIGNER()
	return inst
}

// GetOwnerAccount gets the "owner" account.
func (inst *AddItem) GetOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetMintAccount sets the "mint" account.
func (inst *AddItem) SetMintAccount(mint ag_solanago.PublicKey) *AddItem {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(mint)
	return inst
}

// GetMintAccount gets the "mint" account.
func (inst *AddItem) GetMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *AddItem) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *AddItem {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *AddItem) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

func (inst AddItem) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_AddItem,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst AddItem) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *AddItem) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Bump == nil {
			return errors.New("Bump parameter is not set")
		}
		if inst.Foxy == nil {
			return errors.New("Foxy parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Item is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Owner is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Mint is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *AddItem) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("AddItem")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Bump", *inst.Bump))
						paramsBranch.Child(ag_format.Param("Foxy", *inst.Foxy))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=4]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("         item", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("        owner", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("         mint", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("systemProgram", inst.AccountMetaSlice.Get(3)))
					})
				})
		})
}

func (obj AddItem) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Bump` param:
	err = encoder.Encode(obj.Bump)
	if err != nil {
		return err
	}
	// Serialize `Foxy` param:
	err = encoder.Encode(obj.Foxy)
	if err != nil {
		return err
	}
	return nil
}
func (obj *AddItem) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Bump`:
	err = decoder.Decode(&obj.Bump)
	if err != nil {
		return err
	}
	// Deserialize `Foxy`:
	err = decoder.Decode(&obj.Foxy)
	if err != nil {
		return err
	}
	return nil
}

// NewAddItemInstruction declares a new AddItem instruction with the provided parameters and accounts.
func NewAddItemInstruction(
	// Parameters:
	bump uint8,
	foxy bool,
	// Accounts:
	item ag_solanago.PublicKey,
	owner ag_solanago.PublicKey,
	mint ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *AddItem {
	return NewAddItemInstructionBuilder().
		SetBump(bump).
		SetFoxy(foxy).
		SetItemAccount(item).
		SetOwnerAccount(owner).
		SetMintAccount(mint).
		SetSystemProgramAccount(systemProgram)
}
