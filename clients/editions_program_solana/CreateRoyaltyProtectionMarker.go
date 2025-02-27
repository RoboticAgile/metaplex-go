// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package editions_program_solana

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/RoboticAgile/solana-go"
	ag_format "github.com/RoboticAgile/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// CreateRoyaltyProtectionMarker is the `createRoyaltyProtectionMarker` instruction.
type CreateRoyaltyProtectionMarker struct {

	// [0] = [WRITE, SIGNER] buyer
	//
	// [1] = [WRITE] royaltyProtectionMarker
	//
	// [2] = [] mintKey
	//
	// [3] = [] exchgDepositAuthority
	//
	// [4] = [] systemProgram
	//
	// [5] = [] rent
	//
	// [6] = [] instructions
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewCreateRoyaltyProtectionMarkerInstructionBuilder creates a new `CreateRoyaltyProtectionMarker` instruction builder.
func NewCreateRoyaltyProtectionMarkerInstructionBuilder() *CreateRoyaltyProtectionMarker {
	nd := &CreateRoyaltyProtectionMarker{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 7),
	}
	return nd
}

// SetBuyerAccount sets the "buyer" account.
func (inst *CreateRoyaltyProtectionMarker) SetBuyerAccount(buyer ag_solanago.PublicKey) *CreateRoyaltyProtectionMarker {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(buyer).WRITE().SIGNER()
	return inst
}

// GetBuyerAccount gets the "buyer" account.
func (inst *CreateRoyaltyProtectionMarker) GetBuyerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetRoyaltyProtectionMarkerAccount sets the "royaltyProtectionMarker" account.
func (inst *CreateRoyaltyProtectionMarker) SetRoyaltyProtectionMarkerAccount(royaltyProtectionMarker ag_solanago.PublicKey) *CreateRoyaltyProtectionMarker {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(royaltyProtectionMarker).WRITE()
	return inst
}

// GetRoyaltyProtectionMarkerAccount gets the "royaltyProtectionMarker" account.
func (inst *CreateRoyaltyProtectionMarker) GetRoyaltyProtectionMarkerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetMintKeyAccount sets the "mintKey" account.
func (inst *CreateRoyaltyProtectionMarker) SetMintKeyAccount(mintKey ag_solanago.PublicKey) *CreateRoyaltyProtectionMarker {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(mintKey)
	return inst
}

// GetMintKeyAccount gets the "mintKey" account.
func (inst *CreateRoyaltyProtectionMarker) GetMintKeyAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetExchgDepositAuthorityAccount sets the "exchgDepositAuthority" account.
func (inst *CreateRoyaltyProtectionMarker) SetExchgDepositAuthorityAccount(exchgDepositAuthority ag_solanago.PublicKey) *CreateRoyaltyProtectionMarker {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(exchgDepositAuthority)
	return inst
}

// GetExchgDepositAuthorityAccount gets the "exchgDepositAuthority" account.
func (inst *CreateRoyaltyProtectionMarker) GetExchgDepositAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *CreateRoyaltyProtectionMarker) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *CreateRoyaltyProtectionMarker {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *CreateRoyaltyProtectionMarker) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetRentAccount sets the "rent" account.
func (inst *CreateRoyaltyProtectionMarker) SetRentAccount(rent ag_solanago.PublicKey) *CreateRoyaltyProtectionMarker {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *CreateRoyaltyProtectionMarker) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetInstructionsAccount sets the "instructions" account.
func (inst *CreateRoyaltyProtectionMarker) SetInstructionsAccount(instructions ag_solanago.PublicKey) *CreateRoyaltyProtectionMarker {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(instructions)
	return inst
}

// GetInstructionsAccount gets the "instructions" account.
func (inst *CreateRoyaltyProtectionMarker) GetInstructionsAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

func (inst CreateRoyaltyProtectionMarker) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_CreateRoyaltyProtectionMarker,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst CreateRoyaltyProtectionMarker) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *CreateRoyaltyProtectionMarker) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Buyer is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.RoyaltyProtectionMarker is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.MintKey is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.ExchgDepositAuthority is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.Rent is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.Instructions is not set")
		}
	}
	return nil
}

func (inst *CreateRoyaltyProtectionMarker) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("CreateRoyaltyProtectionMarker")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=7]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                  buyer", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("royaltyProtectionMarker", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("                mintKey", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("  exchgDepositAuthority", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("          systemProgram", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("                   rent", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("           instructions", inst.AccountMetaSlice.Get(6)))
					})
				})
		})
}

func (obj CreateRoyaltyProtectionMarker) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *CreateRoyaltyProtectionMarker) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewCreateRoyaltyProtectionMarkerInstruction declares a new CreateRoyaltyProtectionMarker instruction with the provided parameters and accounts.
func NewCreateRoyaltyProtectionMarkerInstruction(
	// Accounts:
	buyer ag_solanago.PublicKey,
	royaltyProtectionMarker ag_solanago.PublicKey,
	mintKey ag_solanago.PublicKey,
	exchgDepositAuthority ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	rent ag_solanago.PublicKey,
	instructions ag_solanago.PublicKey) *CreateRoyaltyProtectionMarker {
	return NewCreateRoyaltyProtectionMarkerInstructionBuilder().
		SetBuyerAccount(buyer).
		SetRoyaltyProtectionMarkerAccount(royaltyProtectionMarker).
		SetMintKeyAccount(mintKey).
		SetExchgDepositAuthorityAccount(exchgDepositAuthority).
		SetSystemProgramAccount(systemProgram).
		SetRentAccount(rent).
		SetInstructionsAccount(instructions)
}
