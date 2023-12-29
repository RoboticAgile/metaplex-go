// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package editions_program_solana

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/desperatee/solana-go"
	ag_format "github.com/desperatee/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// CancelBuynowSale is the `cancelBuynowSale` instruction.
type CancelBuynowSale struct {

	// [0] = [WRITE, SIGNER] seller
	//
	// [1] = [WRITE] saleStateAccount
	//
	// [2] = [WRITE] mintKey
	//
	// [3] = [WRITE] sellerMintTokenAccount
	//
	// [4] = [WRITE] exchgDepositAuthority
	//
	// [5] = [WRITE] exchgDepositAccount
	//
	// [6] = [WRITE] cardinalManager
	//
	// [7] = [WRITE] cardinalMintCounter
	//
	// [8] = [] instructions
	//
	// [9] = [] tokenProgram
	//
	// [10] = [] associatedTokenProgram
	//
	// [11] = [] cardinalTokenManagerProgram
	//
	// [12] = [] systemProgram
	//
	// [13] = [] rent
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewCancelBuynowSaleInstructionBuilder creates a new `CancelBuynowSale` instruction builder.
func NewCancelBuynowSaleInstructionBuilder() *CancelBuynowSale {
	nd := &CancelBuynowSale{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 14),
	}
	return nd
}

// SetSellerAccount sets the "seller" account.
func (inst *CancelBuynowSale) SetSellerAccount(seller ag_solanago.PublicKey) *CancelBuynowSale {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(seller).WRITE().SIGNER()
	return inst
}

// GetSellerAccount gets the "seller" account.
func (inst *CancelBuynowSale) GetSellerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetSaleStateAccountAccount sets the "saleStateAccount" account.
func (inst *CancelBuynowSale) SetSaleStateAccountAccount(saleStateAccount ag_solanago.PublicKey) *CancelBuynowSale {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(saleStateAccount).WRITE()
	return inst
}

// GetSaleStateAccountAccount gets the "saleStateAccount" account.
func (inst *CancelBuynowSale) GetSaleStateAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetMintKeyAccount sets the "mintKey" account.
func (inst *CancelBuynowSale) SetMintKeyAccount(mintKey ag_solanago.PublicKey) *CancelBuynowSale {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(mintKey).WRITE()
	return inst
}

// GetMintKeyAccount gets the "mintKey" account.
func (inst *CancelBuynowSale) GetMintKeyAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetSellerMintTokenAccountAccount sets the "sellerMintTokenAccount" account.
func (inst *CancelBuynowSale) SetSellerMintTokenAccountAccount(sellerMintTokenAccount ag_solanago.PublicKey) *CancelBuynowSale {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(sellerMintTokenAccount).WRITE()
	return inst
}

// GetSellerMintTokenAccountAccount gets the "sellerMintTokenAccount" account.
func (inst *CancelBuynowSale) GetSellerMintTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetExchgDepositAuthorityAccount sets the "exchgDepositAuthority" account.
func (inst *CancelBuynowSale) SetExchgDepositAuthorityAccount(exchgDepositAuthority ag_solanago.PublicKey) *CancelBuynowSale {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(exchgDepositAuthority).WRITE()
	return inst
}

// GetExchgDepositAuthorityAccount gets the "exchgDepositAuthority" account.
func (inst *CancelBuynowSale) GetExchgDepositAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetExchgDepositAccountAccount sets the "exchgDepositAccount" account.
func (inst *CancelBuynowSale) SetExchgDepositAccountAccount(exchgDepositAccount ag_solanago.PublicKey) *CancelBuynowSale {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(exchgDepositAccount).WRITE()
	return inst
}

// GetExchgDepositAccountAccount gets the "exchgDepositAccount" account.
func (inst *CancelBuynowSale) GetExchgDepositAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetCardinalManagerAccount sets the "cardinalManager" account.
func (inst *CancelBuynowSale) SetCardinalManagerAccount(cardinalManager ag_solanago.PublicKey) *CancelBuynowSale {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(cardinalManager).WRITE()
	return inst
}

// GetCardinalManagerAccount gets the "cardinalManager" account.
func (inst *CancelBuynowSale) GetCardinalManagerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetCardinalMintCounterAccount sets the "cardinalMintCounter" account.
func (inst *CancelBuynowSale) SetCardinalMintCounterAccount(cardinalMintCounter ag_solanago.PublicKey) *CancelBuynowSale {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(cardinalMintCounter).WRITE()
	return inst
}

// GetCardinalMintCounterAccount gets the "cardinalMintCounter" account.
func (inst *CancelBuynowSale) GetCardinalMintCounterAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetInstructionsAccount sets the "instructions" account.
func (inst *CancelBuynowSale) SetInstructionsAccount(instructions ag_solanago.PublicKey) *CancelBuynowSale {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(instructions)
	return inst
}

// GetInstructionsAccount gets the "instructions" account.
func (inst *CancelBuynowSale) GetInstructionsAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *CancelBuynowSale) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *CancelBuynowSale {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *CancelBuynowSale) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetAssociatedTokenProgramAccount sets the "associatedTokenProgram" account.
func (inst *CancelBuynowSale) SetAssociatedTokenProgramAccount(associatedTokenProgram ag_solanago.PublicKey) *CancelBuynowSale {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(associatedTokenProgram)
	return inst
}

// GetAssociatedTokenProgramAccount gets the "associatedTokenProgram" account.
func (inst *CancelBuynowSale) GetAssociatedTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetCardinalTokenManagerProgramAccount sets the "cardinalTokenManagerProgram" account.
func (inst *CancelBuynowSale) SetCardinalTokenManagerProgramAccount(cardinalTokenManagerProgram ag_solanago.PublicKey) *CancelBuynowSale {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(cardinalTokenManagerProgram)
	return inst
}

// GetCardinalTokenManagerProgramAccount gets the "cardinalTokenManagerProgram" account.
func (inst *CancelBuynowSale) GetCardinalTokenManagerProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *CancelBuynowSale) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *CancelBuynowSale {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *CancelBuynowSale) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

// SetRentAccount sets the "rent" account.
func (inst *CancelBuynowSale) SetRentAccount(rent ag_solanago.PublicKey) *CancelBuynowSale {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *CancelBuynowSale) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(13)
}

func (inst CancelBuynowSale) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_CancelBuynowSale,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst CancelBuynowSale) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *CancelBuynowSale) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Seller is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.SaleStateAccount is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.MintKey is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.SellerMintTokenAccount is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.ExchgDepositAuthority is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.ExchgDepositAccount is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.CardinalManager is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.CardinalMintCounter is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.Instructions is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.AssociatedTokenProgram is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.CardinalTokenManagerProgram is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.Rent is not set")
		}
	}
	return nil
}

func (inst *CancelBuynowSale) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("CancelBuynowSale")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=14]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                     seller", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("                  saleState", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("                    mintKey", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("            sellerMintToken", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("      exchgDepositAuthority", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("               exchgDeposit", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("            cardinalManager", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("        cardinalMintCounter", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("               instructions", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("               tokenProgram", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("     associatedTokenProgram", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("cardinalTokenManagerProgram", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("              systemProgram", inst.AccountMetaSlice.Get(12)))
						accountsBranch.Child(ag_format.Meta("                       rent", inst.AccountMetaSlice.Get(13)))
					})
				})
		})
}

func (obj CancelBuynowSale) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *CancelBuynowSale) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewCancelBuynowSaleInstruction declares a new CancelBuynowSale instruction with the provided parameters and accounts.
func NewCancelBuynowSaleInstruction(
	// Accounts:
	seller ag_solanago.PublicKey,
	saleStateAccount ag_solanago.PublicKey,
	mintKey ag_solanago.PublicKey,
	sellerMintTokenAccount ag_solanago.PublicKey,
	exchgDepositAuthority ag_solanago.PublicKey,
	exchgDepositAccount ag_solanago.PublicKey,
	cardinalManager ag_solanago.PublicKey,
	cardinalMintCounter ag_solanago.PublicKey,
	instructions ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	associatedTokenProgram ag_solanago.PublicKey,
	cardinalTokenManagerProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	rent ag_solanago.PublicKey) *CancelBuynowSale {
	return NewCancelBuynowSaleInstructionBuilder().
		SetSellerAccount(seller).
		SetSaleStateAccountAccount(saleStateAccount).
		SetMintKeyAccount(mintKey).
		SetSellerMintTokenAccountAccount(sellerMintTokenAccount).
		SetExchgDepositAuthorityAccount(exchgDepositAuthority).
		SetExchgDepositAccountAccount(exchgDepositAccount).
		SetCardinalManagerAccount(cardinalManager).
		SetCardinalMintCounterAccount(cardinalMintCounter).
		SetInstructionsAccount(instructions).
		SetTokenProgramAccount(tokenProgram).
		SetAssociatedTokenProgramAccount(associatedTokenProgram).
		SetCardinalTokenManagerProgramAccount(cardinalTokenManagerProgram).
		SetSystemProgramAccount(systemProgram).
		SetRentAccount(rent)
}
