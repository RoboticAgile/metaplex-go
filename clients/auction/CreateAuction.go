// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package auction

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/RoboticAgile/solana-go"
	ag_format "github.com/RoboticAgile/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Create a new auction account bound to a resource, initially in a pending state.
type CreateAuction struct {
	Args *CreateAuctionArgs

	// [0] = [SIGNER] creator
	// ··········· The account creating the auction, which is authorised to make changes.
	//
	// [1] = [WRITE] uninitializedAuctionAccount
	// ··········· Uninitialized auction account.
	//
	// [2] = [WRITE] auctionExtendedData
	// ··········· Auction extended data account (pda relative to auction of ['auction', program id, vault key, 'extended']).
	//
	// [3] = [] rentSysvar
	// ··········· Rent sysvar
	//
	// [4] = [] systemAccount
	// ··········· System account
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewCreateAuctionInstructionBuilder creates a new `CreateAuction` instruction builder.
func NewCreateAuctionInstructionBuilder() *CreateAuction {
	nd := &CreateAuction{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 5),
	}
	return nd
}

// SetArgs sets the "args" parameter.
func (inst *CreateAuction) SetArgs(args CreateAuctionArgs) *CreateAuction {
	inst.Args = &args
	return inst
}

// SetCreatorAccount sets the "creator" account.
// The account creating the auction, which is authorised to make changes.
func (inst *CreateAuction) SetCreatorAccount(creator ag_solanago.PublicKey) *CreateAuction {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(creator).SIGNER()
	return inst
}

// GetCreatorAccount gets the "creator" account.
// The account creating the auction, which is authorised to make changes.
func (inst *CreateAuction) GetCreatorAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetUninitializedAuctionAccount sets the "uninitializedAuctionAccount" account.
// Uninitialized auction account.
func (inst *CreateAuction) SetUninitializedAuctionAccount(uninitializedAuctionAccount ag_solanago.PublicKey) *CreateAuction {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(uninitializedAuctionAccount).WRITE()
	return inst
}

// GetUninitializedAuctionAccount gets the "uninitializedAuctionAccount" account.
// Uninitialized auction account.
func (inst *CreateAuction) GetUninitializedAuctionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetAuctionExtendedDataAccount sets the "auctionExtendedData" account.
// Auction extended data account (pda relative to auction of ['auction', program id, vault key, 'extended']).
func (inst *CreateAuction) SetAuctionExtendedDataAccount(auctionExtendedData ag_solanago.PublicKey) *CreateAuction {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(auctionExtendedData).WRITE()
	return inst
}

// GetAuctionExtendedDataAccount gets the "auctionExtendedData" account.
// Auction extended data account (pda relative to auction of ['auction', program id, vault key, 'extended']).
func (inst *CreateAuction) GetAuctionExtendedDataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetRentSysvarAccount sets the "rentSysvar" account.
// Rent sysvar
func (inst *CreateAuction) SetRentSysvarAccount(rentSysvar ag_solanago.PublicKey) *CreateAuction {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(rentSysvar)
	return inst
}

// GetRentSysvarAccount gets the "rentSysvar" account.
// Rent sysvar
func (inst *CreateAuction) GetRentSysvarAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

// SetSystemAccount sets the "systemAccount" account.
// System account
func (inst *CreateAuction) SetSystemAccount(systemAccount ag_solanago.PublicKey) *CreateAuction {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(systemAccount)
	return inst
}

// GetSystemAccount gets the "systemAccount" account.
// System account
func (inst *CreateAuction) GetSystemAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[4]
}

func (inst CreateAuction) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_CreateAuction),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst CreateAuction) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *CreateAuction) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Args == nil {
			return errors.New("Args parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Creator is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.UninitializedAuctionAccount is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.AuctionExtendedData is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.RentSysvar is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.SystemAccount is not set")
		}
	}
	return nil
}

func (inst *CreateAuction) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("CreateAuction")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Args", *inst.Args))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=5]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("             creator", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("uninitializedAuction", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta(" auctionExtendedData", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta("          rentSysvar", inst.AccountMetaSlice[3]))
						accountsBranch.Child(ag_format.Meta("              system", inst.AccountMetaSlice[4]))
					})
				})
		})
}

func (obj CreateAuction) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Args` param:
	err = encoder.Encode(obj.Args)
	if err != nil {
		return err
	}
	return nil
}
func (obj *CreateAuction) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Args`:
	err = decoder.Decode(&obj.Args)
	if err != nil {
		return err
	}
	return nil
}

// NewCreateAuctionInstruction declares a new CreateAuction instruction with the provided parameters and accounts.
func NewCreateAuctionInstruction(
	// Parameters:
	args CreateAuctionArgs,
	// Accounts:
	creator ag_solanago.PublicKey,
	uninitializedAuctionAccount ag_solanago.PublicKey,
	auctionExtendedData ag_solanago.PublicKey,
	rentSysvar ag_solanago.PublicKey,
	systemAccount ag_solanago.PublicKey) *CreateAuction {
	return NewCreateAuctionInstructionBuilder().
		SetArgs(args).
		SetCreatorAccount(creator).
		SetUninitializedAuctionAccount(uninitializedAuctionAccount).
		SetAuctionExtendedDataAccount(auctionExtendedData).
		SetRentSysvarAccount(rentSysvar).
		SetSystemAccount(systemAccount)
}
