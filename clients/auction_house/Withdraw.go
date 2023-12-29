// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package auction_house

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/desperatee/solana-go"
	ag_format "github.com/desperatee/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Withdraw is the `withdraw` instruction.
type Withdraw struct {
	EscrowPaymentBump *uint8
	Amount            *uint64

	// [0] = [] wallet
	//
	// [1] = [WRITE] receiptAccount
	//
	// [2] = [WRITE] escrowPaymentAccount
	//
	// [3] = [] treasuryMint
	//
	// [4] = [] authority
	//
	// [5] = [] auctionHouse
	//
	// [6] = [WRITE] auctionHouseFeeAccount
	//
	// [7] = [] tokenProgram
	//
	// [8] = [] systemProgram
	//
	// [9] = [] ataProgram
	//
	// [10] = [] rent
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewWithdrawInstructionBuilder creates a new `Withdraw` instruction builder.
func NewWithdrawInstructionBuilder() *Withdraw {
	nd := &Withdraw{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 11),
	}
	return nd
}

// SetEscrowPaymentBump sets the "escrowPaymentBump" parameter.
func (inst *Withdraw) SetEscrowPaymentBump(escrowPaymentBump uint8) *Withdraw {
	inst.EscrowPaymentBump = &escrowPaymentBump
	return inst
}

// SetAmount sets the "amount" parameter.
func (inst *Withdraw) SetAmount(amount uint64) *Withdraw {
	inst.Amount = &amount
	return inst
}

// SetWalletAccount sets the "wallet" account.
func (inst *Withdraw) SetWalletAccount(wallet ag_solanago.PublicKey) *Withdraw {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(wallet)
	return inst
}

// GetWalletAccount gets the "wallet" account.
func (inst *Withdraw) GetWalletAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetReceiptAccountAccount sets the "receiptAccount" account.
func (inst *Withdraw) SetReceiptAccountAccount(receiptAccount ag_solanago.PublicKey) *Withdraw {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(receiptAccount).WRITE()
	return inst
}

// GetReceiptAccountAccount gets the "receiptAccount" account.
func (inst *Withdraw) GetReceiptAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetEscrowPaymentAccountAccount sets the "escrowPaymentAccount" account.
func (inst *Withdraw) SetEscrowPaymentAccountAccount(escrowPaymentAccount ag_solanago.PublicKey) *Withdraw {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(escrowPaymentAccount).WRITE()
	return inst
}

// GetEscrowPaymentAccountAccount gets the "escrowPaymentAccount" account.
func (inst *Withdraw) GetEscrowPaymentAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetTreasuryMintAccount sets the "treasuryMint" account.
func (inst *Withdraw) SetTreasuryMintAccount(treasuryMint ag_solanago.PublicKey) *Withdraw {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(treasuryMint)
	return inst
}

// GetTreasuryMintAccount gets the "treasuryMint" account.
func (inst *Withdraw) GetTreasuryMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetAuthorityAccount sets the "authority" account.
func (inst *Withdraw) SetAuthorityAccount(authority ag_solanago.PublicKey) *Withdraw {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(authority)
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *Withdraw) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetAuctionHouseAccount sets the "auctionHouse" account.
func (inst *Withdraw) SetAuctionHouseAccount(auctionHouse ag_solanago.PublicKey) *Withdraw {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(auctionHouse)
	return inst
}

// GetAuctionHouseAccount gets the "auctionHouse" account.
func (inst *Withdraw) GetAuctionHouseAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetAuctionHouseFeeAccountAccount sets the "auctionHouseFeeAccount" account.
func (inst *Withdraw) SetAuctionHouseFeeAccountAccount(auctionHouseFeeAccount ag_solanago.PublicKey) *Withdraw {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(auctionHouseFeeAccount).WRITE()
	return inst
}

// GetAuctionHouseFeeAccountAccount gets the "auctionHouseFeeAccount" account.
func (inst *Withdraw) GetAuctionHouseFeeAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *Withdraw) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *Withdraw {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *Withdraw) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *Withdraw) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *Withdraw {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *Withdraw) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetAtaProgramAccount sets the "ataProgram" account.
func (inst *Withdraw) SetAtaProgramAccount(ataProgram ag_solanago.PublicKey) *Withdraw {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(ataProgram)
	return inst
}

// GetAtaProgramAccount gets the "ataProgram" account.
func (inst *Withdraw) GetAtaProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetRentAccount sets the "rent" account.
func (inst *Withdraw) SetRentAccount(rent ag_solanago.PublicKey) *Withdraw {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *Withdraw) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

func (inst Withdraw) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_Withdraw,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst Withdraw) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *Withdraw) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.EscrowPaymentBump == nil {
			return errors.New("EscrowPaymentBump parameter is not set")
		}
		if inst.Amount == nil {
			return errors.New("Amount parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Wallet is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.ReceiptAccount is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.EscrowPaymentAccount is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.TreasuryMint is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.AuctionHouse is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.AuctionHouseFeeAccount is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.AtaProgram is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.Rent is not set")
		}
	}
	return nil
}

func (inst *Withdraw) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("Withdraw")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("EscrowPaymentBump", *inst.EscrowPaymentBump))
						paramsBranch.Child(ag_format.Param("           Amount", *inst.Amount))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=11]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("         wallet", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("        receipt", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("  escrowPayment", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("   treasuryMint", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("      authority", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("   auctionHouse", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("auctionHouseFee", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("   tokenProgram", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("  systemProgram", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("     ataProgram", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("           rent", inst.AccountMetaSlice.Get(10)))
					})
				})
		})
}

func (obj Withdraw) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `EscrowPaymentBump` param:
	err = encoder.Encode(obj.EscrowPaymentBump)
	if err != nil {
		return err
	}
	// Serialize `Amount` param:
	err = encoder.Encode(obj.Amount)
	if err != nil {
		return err
	}
	return nil
}
func (obj *Withdraw) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `EscrowPaymentBump`:
	err = decoder.Decode(&obj.EscrowPaymentBump)
	if err != nil {
		return err
	}
	// Deserialize `Amount`:
	err = decoder.Decode(&obj.Amount)
	if err != nil {
		return err
	}
	return nil
}

// NewWithdrawInstruction declares a new Withdraw instruction with the provided parameters and accounts.
func NewWithdrawInstruction(
	// Parameters:
	escrowPaymentBump uint8,
	amount uint64,
	// Accounts:
	wallet ag_solanago.PublicKey,
	receiptAccount ag_solanago.PublicKey,
	escrowPaymentAccount ag_solanago.PublicKey,
	treasuryMint ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	auctionHouse ag_solanago.PublicKey,
	auctionHouseFeeAccount ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	ataProgram ag_solanago.PublicKey,
	rent ag_solanago.PublicKey) *Withdraw {
	return NewWithdrawInstructionBuilder().
		SetEscrowPaymentBump(escrowPaymentBump).
		SetAmount(amount).
		SetWalletAccount(wallet).
		SetReceiptAccountAccount(receiptAccount).
		SetEscrowPaymentAccountAccount(escrowPaymentAccount).
		SetTreasuryMintAccount(treasuryMint).
		SetAuthorityAccount(authority).
		SetAuctionHouseAccount(auctionHouse).
		SetAuctionHouseFeeAccountAccount(auctionHouseFeeAccount).
		SetTokenProgramAccount(tokenProgram).
		SetSystemProgramAccount(systemProgram).
		SetAtaProgramAccount(ataProgram).
		SetRentAccount(rent)
}
