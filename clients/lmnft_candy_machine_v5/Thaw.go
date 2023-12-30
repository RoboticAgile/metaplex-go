// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package nft_candy_machine

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/RoboticAgile/solana-go"
	ag_format "github.com/RoboticAgile/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Thaw is the `thaw` instruction.
type Thaw struct {

	// [0] = [WRITE] candyMachine
	//
	// [1] = [] mint
	//
	// [2] = [WRITE] associated
	//
	// [3] = [WRITE] metadata
	//
	// [4] = [WRITE] masterEdition
	//
	// [5] = [] tokenMetadataProgram
	//
	// [6] = [] tokenProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewThawInstructionBuilder creates a new `Thaw` instruction builder.
func NewThawInstructionBuilder() *Thaw {
	nd := &Thaw{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 7),
	}
	return nd
}

// SetCandyMachineAccount sets the "candyMachine" account.
func (inst *Thaw) SetCandyMachineAccount(candyMachine ag_solanago.PublicKey) *Thaw {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(candyMachine).WRITE()
	return inst
}

// GetCandyMachineAccount gets the "candyMachine" account.
func (inst *Thaw) GetCandyMachineAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetMintAccount sets the "mint" account.
func (inst *Thaw) SetMintAccount(mint ag_solanago.PublicKey) *Thaw {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(mint)
	return inst
}

// GetMintAccount gets the "mint" account.
func (inst *Thaw) GetMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetAssociatedAccount sets the "associated" account.
func (inst *Thaw) SetAssociatedAccount(associated ag_solanago.PublicKey) *Thaw {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(associated).WRITE()
	return inst
}

// GetAssociatedAccount gets the "associated" account.
func (inst *Thaw) GetAssociatedAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetMetadataAccount sets the "metadata" account.
func (inst *Thaw) SetMetadataAccount(metadata ag_solanago.PublicKey) *Thaw {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(metadata).WRITE()
	return inst
}

// GetMetadataAccount gets the "metadata" account.
func (inst *Thaw) GetMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetMasterEditionAccount sets the "masterEdition" account.
func (inst *Thaw) SetMasterEditionAccount(masterEdition ag_solanago.PublicKey) *Thaw {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(masterEdition).WRITE()
	return inst
}

// GetMasterEditionAccount gets the "masterEdition" account.
func (inst *Thaw) GetMasterEditionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetTokenMetadataProgramAccount sets the "tokenMetadataProgram" account.
func (inst *Thaw) SetTokenMetadataProgramAccount(tokenMetadataProgram ag_solanago.PublicKey) *Thaw {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(tokenMetadataProgram)
	return inst
}

// GetTokenMetadataProgramAccount gets the "tokenMetadataProgram" account.
func (inst *Thaw) GetTokenMetadataProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *Thaw) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *Thaw {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *Thaw) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

func (inst Thaw) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_Thaw,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst Thaw) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *Thaw) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.CandyMachine is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Mint is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Associated is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Metadata is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.MasterEdition is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.TokenMetadataProgram is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
	}
	return nil
}

func (inst *Thaw) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("Thaw")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=7]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("        candyMachine", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("                mint", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("          associated", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("            metadata", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("       masterEdition", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("tokenMetadataProgram", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("        tokenProgram", inst.AccountMetaSlice.Get(6)))
					})
				})
		})
}

func (obj Thaw) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *Thaw) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewThawInstruction declares a new Thaw instruction with the provided parameters and accounts.
func NewThawInstruction(
	// Accounts:
	candyMachine ag_solanago.PublicKey,
	mint ag_solanago.PublicKey,
	associated ag_solanago.PublicKey,
	metadata ag_solanago.PublicKey,
	masterEdition ag_solanago.PublicKey,
	tokenMetadataProgram ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey) *Thaw {
	return NewThawInstructionBuilder().
		SetCandyMachineAccount(candyMachine).
		SetMintAccount(mint).
		SetAssociatedAccount(associated).
		SetMetadataAccount(metadata).
		SetMasterEditionAccount(masterEdition).
		SetTokenMetadataProgramAccount(tokenMetadataProgram).
		SetTokenProgramAccount(tokenProgram)
}
