// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package token_vault

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/RoboticAgile/solana-go"
	ag_format "github.com/RoboticAgile/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// If in combine state, authority on vault can hit this to withdrawal some of a token type from a safety deposit box.
// Once fractional supply is zero and all tokens have been removed this action will take vault to Deactivated
type WithdrawTokenFromSafetyDepositBox struct {
	Args *AmountArgs

	// [0] = [WRITE] initializedDestination
	// ··········· Initialized Destination account for the tokens being withdrawn
	//
	// [1] = [WRITE] safetyDepositBox
	// ··········· The safety deposit box account key for the tokens
	//
	// [2] = [WRITE] storeKey
	// ··········· The store key on the safety deposit box account
	//
	// [3] = [WRITE] initializedCombinedTokenVault
	// ··········· The initialized combined token vault
	//
	// [4] = [] fractionMint
	// ··········· Fraction mint
	//
	// [5] = [SIGNER] vaultAuthority
	// ··········· Authority of vault
	//
	// [6] = [] pdaBasedTransferAuthority
	// ··········· PDA-based Transfer authority to move the tokens from the store to the destination seed [PREFIX, program_id]
	//
	// [7] = [] tokenProgram
	// ··········· Token program
	//
	// [8] = [] rentSysvar
	// ··········· Rent sysvar
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewWithdrawTokenFromSafetyDepositBoxInstructionBuilder creates a new `WithdrawTokenFromSafetyDepositBox` instruction builder.
func NewWithdrawTokenFromSafetyDepositBoxInstructionBuilder() *WithdrawTokenFromSafetyDepositBox {
	nd := &WithdrawTokenFromSafetyDepositBox{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 9),
	}
	return nd
}

// SetArgs sets the "args" parameter.
func (inst *WithdrawTokenFromSafetyDepositBox) SetArgs(args AmountArgs) *WithdrawTokenFromSafetyDepositBox {
	inst.Args = &args
	return inst
}

// SetInitializedDestinationAccount sets the "initializedDestination" account.
// Initialized Destination account for the tokens being withdrawn
func (inst *WithdrawTokenFromSafetyDepositBox) SetInitializedDestinationAccount(initializedDestination ag_solanago.PublicKey) *WithdrawTokenFromSafetyDepositBox {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(initializedDestination).WRITE()
	return inst
}

// GetInitializedDestinationAccount gets the "initializedDestination" account.
// Initialized Destination account for the tokens being withdrawn
func (inst *WithdrawTokenFromSafetyDepositBox) GetInitializedDestinationAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetSafetyDepositBoxAccount sets the "safetyDepositBox" account.
// The safety deposit box account key for the tokens
func (inst *WithdrawTokenFromSafetyDepositBox) SetSafetyDepositBoxAccount(safetyDepositBox ag_solanago.PublicKey) *WithdrawTokenFromSafetyDepositBox {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(safetyDepositBox).WRITE()
	return inst
}

// GetSafetyDepositBoxAccount gets the "safetyDepositBox" account.
// The safety deposit box account key for the tokens
func (inst *WithdrawTokenFromSafetyDepositBox) GetSafetyDepositBoxAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetStoreKeyAccount sets the "storeKey" account.
// The store key on the safety deposit box account
func (inst *WithdrawTokenFromSafetyDepositBox) SetStoreKeyAccount(storeKey ag_solanago.PublicKey) *WithdrawTokenFromSafetyDepositBox {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(storeKey).WRITE()
	return inst
}

// GetStoreKeyAccount gets the "storeKey" account.
// The store key on the safety deposit box account
func (inst *WithdrawTokenFromSafetyDepositBox) GetStoreKeyAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetInitializedCombinedTokenVaultAccount sets the "initializedCombinedTokenVault" account.
// The initialized combined token vault
func (inst *WithdrawTokenFromSafetyDepositBox) SetInitializedCombinedTokenVaultAccount(initializedCombinedTokenVault ag_solanago.PublicKey) *WithdrawTokenFromSafetyDepositBox {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(initializedCombinedTokenVault).WRITE()
	return inst
}

// GetInitializedCombinedTokenVaultAccount gets the "initializedCombinedTokenVault" account.
// The initialized combined token vault
func (inst *WithdrawTokenFromSafetyDepositBox) GetInitializedCombinedTokenVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetFractionMintAccount sets the "fractionMint" account.
// Fraction mint
func (inst *WithdrawTokenFromSafetyDepositBox) SetFractionMintAccount(fractionMint ag_solanago.PublicKey) *WithdrawTokenFromSafetyDepositBox {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(fractionMint)
	return inst
}

// GetFractionMintAccount gets the "fractionMint" account.
// Fraction mint
func (inst *WithdrawTokenFromSafetyDepositBox) GetFractionMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetVaultAuthorityAccount sets the "vaultAuthority" account.
// Authority of vault
func (inst *WithdrawTokenFromSafetyDepositBox) SetVaultAuthorityAccount(vaultAuthority ag_solanago.PublicKey) *WithdrawTokenFromSafetyDepositBox {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(vaultAuthority).SIGNER()
	return inst
}

// GetVaultAuthorityAccount gets the "vaultAuthority" account.
// Authority of vault
func (inst *WithdrawTokenFromSafetyDepositBox) GetVaultAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetPdaBasedTransferAuthorityAccount sets the "pdaBasedTransferAuthority" account.
// PDA-based Transfer authority to move the tokens from the store to the destination seed [PREFIX, program_id]
func (inst *WithdrawTokenFromSafetyDepositBox) SetPdaBasedTransferAuthorityAccount(pdaBasedTransferAuthority ag_solanago.PublicKey) *WithdrawTokenFromSafetyDepositBox {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(pdaBasedTransferAuthority)
	return inst
}

// GetPdaBasedTransferAuthorityAccount gets the "pdaBasedTransferAuthority" account.
// PDA-based Transfer authority to move the tokens from the store to the destination seed [PREFIX, program_id]
func (inst *WithdrawTokenFromSafetyDepositBox) GetPdaBasedTransferAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
// Token program
func (inst *WithdrawTokenFromSafetyDepositBox) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *WithdrawTokenFromSafetyDepositBox {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
// Token program
func (inst *WithdrawTokenFromSafetyDepositBox) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetRentSysvarAccount sets the "rentSysvar" account.
// Rent sysvar
func (inst *WithdrawTokenFromSafetyDepositBox) SetRentSysvarAccount(rentSysvar ag_solanago.PublicKey) *WithdrawTokenFromSafetyDepositBox {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(rentSysvar)
	return inst
}

// GetRentSysvarAccount gets the "rentSysvar" account.
// Rent sysvar
func (inst *WithdrawTokenFromSafetyDepositBox) GetRentSysvarAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

func (inst WithdrawTokenFromSafetyDepositBox) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_WithdrawTokenFromSafetyDepositBox),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst WithdrawTokenFromSafetyDepositBox) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *WithdrawTokenFromSafetyDepositBox) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Args == nil {
			return errors.New("Args parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.InitializedDestination is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.SafetyDepositBox is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.StoreKey is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.InitializedCombinedTokenVault is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.FractionMint is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.VaultAuthority is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.PdaBasedTransferAuthority is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.RentSysvar is not set")
		}
	}
	return nil
}

func (inst *WithdrawTokenFromSafetyDepositBox) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("WithdrawTokenFromSafetyDepositBox")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Args", *inst.Args))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=9]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("       initializedDestination", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("             safetyDepositBox", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("                     storeKey", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("initializedCombinedTokenVault", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("                 fractionMint", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("               vaultAuthority", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("    pdaBasedTransferAuthority", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("                 tokenProgram", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("                   rentSysvar", inst.AccountMetaSlice.Get(8)))
					})
				})
		})
}

func (obj WithdrawTokenFromSafetyDepositBox) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Args` param:
	err = encoder.Encode(obj.Args)
	if err != nil {
		return err
	}
	return nil
}
func (obj *WithdrawTokenFromSafetyDepositBox) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Args`:
	err = decoder.Decode(&obj.Args)
	if err != nil {
		return err
	}
	return nil
}

// NewWithdrawTokenFromSafetyDepositBoxInstruction declares a new WithdrawTokenFromSafetyDepositBox instruction with the provided parameters and accounts.
func NewWithdrawTokenFromSafetyDepositBoxInstruction(
	// Parameters:
	args AmountArgs,
	// Accounts:
	initializedDestination ag_solanago.PublicKey,
	safetyDepositBox ag_solanago.PublicKey,
	storeKey ag_solanago.PublicKey,
	initializedCombinedTokenVault ag_solanago.PublicKey,
	fractionMint ag_solanago.PublicKey,
	vaultAuthority ag_solanago.PublicKey,
	pdaBasedTransferAuthority ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	rentSysvar ag_solanago.PublicKey) *WithdrawTokenFromSafetyDepositBox {
	return NewWithdrawTokenFromSafetyDepositBoxInstructionBuilder().
		SetArgs(args).
		SetInitializedDestinationAccount(initializedDestination).
		SetSafetyDepositBoxAccount(safetyDepositBox).
		SetStoreKeyAccount(storeKey).
		SetInitializedCombinedTokenVaultAccount(initializedCombinedTokenVault).
		SetFractionMintAccount(fractionMint).
		SetVaultAuthorityAccount(vaultAuthority).
		SetPdaBasedTransferAuthorityAccount(pdaBasedTransferAuthority).
		SetTokenProgramAccount(tokenProgram).
		SetRentSysvarAccount(rentSysvar)
}
