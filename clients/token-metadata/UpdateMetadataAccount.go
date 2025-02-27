// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package token_metadata

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/RoboticAgile/solana-go"
	ag_format "github.com/RoboticAgile/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Update a Metadata
type UpdateMetadataAccount struct {
	Args *UpdateMetadataAccountArgs

	// [0] = [WRITE] metadata
	// ··········· Metadata account
	//
	// [1] = [SIGNER] updateAuthorityKey
	// ··········· Update authority key
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewUpdateMetadataAccountInstructionBuilder creates a new `UpdateMetadataAccount` instruction builder.
func NewUpdateMetadataAccountInstructionBuilder() *UpdateMetadataAccount {
	nd := &UpdateMetadataAccount{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 2),
	}
	return nd
}

// SetArgs sets the "args" parameter.
func (inst *UpdateMetadataAccount) SetArgs(args UpdateMetadataAccountArgs) *UpdateMetadataAccount {
	inst.Args = &args
	return inst
}

// SetMetadataAccount sets the "metadata" account.
// Metadata account
func (inst *UpdateMetadataAccount) SetMetadataAccount(metadata ag_solanago.PublicKey) *UpdateMetadataAccount {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(metadata).WRITE()
	return inst
}

// GetMetadataAccount gets the "metadata" account.
// Metadata account
func (inst *UpdateMetadataAccount) GetMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetUpdateAuthorityKeyAccount sets the "updateAuthorityKey" account.
// Update authority key
func (inst *UpdateMetadataAccount) SetUpdateAuthorityKeyAccount(updateAuthorityKey ag_solanago.PublicKey) *UpdateMetadataAccount {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(updateAuthorityKey).SIGNER()
	return inst
}

// GetUpdateAuthorityKeyAccount gets the "updateAuthorityKey" account.
// Update authority key
func (inst *UpdateMetadataAccount) GetUpdateAuthorityKeyAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

func (inst UpdateMetadataAccount) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_UpdateMetadataAccount),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UpdateMetadataAccount) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UpdateMetadataAccount) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Args == nil {
			return errors.New("Args parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Metadata is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.UpdateAuthorityKey is not set")
		}
	}
	return nil
}

func (inst *UpdateMetadataAccount) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UpdateMetadataAccount")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Args", *inst.Args))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=2]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("          metadata", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("updateAuthorityKey", inst.AccountMetaSlice.Get(1)))
					})
				})
		})
}

func (obj UpdateMetadataAccount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Args` param:
	err = encoder.Encode(obj.Args)
	if err != nil {
		return err
	}
	return nil
}
func (obj *UpdateMetadataAccount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Args`:
	err = decoder.Decode(&obj.Args)
	if err != nil {
		return err
	}
	return nil
}

// NewUpdateMetadataAccountInstruction declares a new UpdateMetadataAccount instruction with the provided parameters and accounts.
func NewUpdateMetadataAccountInstruction(
	// Parameters:
	args UpdateMetadataAccountArgs,
	// Accounts:
	metadata ag_solanago.PublicKey,
	updateAuthorityKey ag_solanago.PublicKey) *UpdateMetadataAccount {
	return NewUpdateMetadataAccountInstructionBuilder().
		SetArgs(args).
		SetMetadataAccount(metadata).
		SetUpdateAuthorityKeyAccount(updateAuthorityKey)
}
