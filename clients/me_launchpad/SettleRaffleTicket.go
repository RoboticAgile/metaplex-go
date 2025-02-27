// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package me_launchpad

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/RoboticAgile/solana-go"
	ag_format "github.com/RoboticAgile/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// SettleRaffleTicket is the `settleRaffleTicket` instruction.
type SettleRaffleTicket struct {
	CurrTime *int64

	// [0] = [] config
	//
	// [1] = [WRITE] candyMachine
	//
	// [2] = [WRITE, SIGNER] payer
	//
	// [3] = [] launchStagesInfo
	//
	// [4] = [WRITE] raffleTicket
	//
	// [5] = [WRITE] raffleEscrow
	// ··········· ATA where SPL tokens are held in escrow, owned by raffle_ticket
	//
	// [6] = [WRITE] payTo
	// ··········· ATA owned by wallet authority that we pay to if we win
	//
	// [7] = [WRITE] refundTo
	// ··········· or ATA owned by raffle_payer to refund SPL
	//
	// [8] = [WRITE] tokenAta
	//
	// [9] = [WRITE] rafflePayer
	//
	// [10] = [WRITE] orderInfo
	//
	// [11] = [SIGNER] notary
	//
	// [12] = [WRITE] metadata
	//
	// [13] = [WRITE, SIGNER] mint
	//
	// [14] = [SIGNER] updateAuthority
	//
	// [15] = [WRITE] masterEdition
	//
	// [16] = [] slotHashes
	//
	// [17] = [] tokenMetadataProgram
	//
	// [18] = [] associatedTokenProgram
	//
	// [19] = [] tokenProgram
	//
	// [20] = [] systemProgram
	//
	// [21] = [] rent
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewSettleRaffleTicketInstructionBuilder creates a new `SettleRaffleTicket` instruction builder.
func NewSettleRaffleTicketInstructionBuilder() *SettleRaffleTicket {
	nd := &SettleRaffleTicket{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 22),
	}
	return nd
}

// SetCurrTime sets the "currTime" parameter.
func (inst *SettleRaffleTicket) SetCurrTime(currTime int64) *SettleRaffleTicket {
	inst.CurrTime = &currTime
	return inst
}

// SetConfigAccount sets the "config" account.
func (inst *SettleRaffleTicket) SetConfigAccount(config ag_solanago.PublicKey) *SettleRaffleTicket {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(config)
	return inst
}

// GetConfigAccount gets the "config" account.
func (inst *SettleRaffleTicket) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetCandyMachineAccount sets the "candyMachine" account.
func (inst *SettleRaffleTicket) SetCandyMachineAccount(candyMachine ag_solanago.PublicKey) *SettleRaffleTicket {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(candyMachine).WRITE()
	return inst
}

// GetCandyMachineAccount gets the "candyMachine" account.
func (inst *SettleRaffleTicket) GetCandyMachineAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetPayerAccount sets the "payer" account.
func (inst *SettleRaffleTicket) SetPayerAccount(payer ag_solanago.PublicKey) *SettleRaffleTicket {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(payer).WRITE().SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
func (inst *SettleRaffleTicket) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetLaunchStagesInfoAccount sets the "launchStagesInfo" account.
func (inst *SettleRaffleTicket) SetLaunchStagesInfoAccount(launchStagesInfo ag_solanago.PublicKey) *SettleRaffleTicket {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(launchStagesInfo)
	return inst
}

// GetLaunchStagesInfoAccount gets the "launchStagesInfo" account.
func (inst *SettleRaffleTicket) GetLaunchStagesInfoAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetRaffleTicketAccount sets the "raffleTicket" account.
func (inst *SettleRaffleTicket) SetRaffleTicketAccount(raffleTicket ag_solanago.PublicKey) *SettleRaffleTicket {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(raffleTicket).WRITE()
	return inst
}

// GetRaffleTicketAccount gets the "raffleTicket" account.
func (inst *SettleRaffleTicket) GetRaffleTicketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetRaffleEscrowAccount sets the "raffleEscrow" account.
// ATA where SPL tokens are held in escrow, owned by raffle_ticket
func (inst *SettleRaffleTicket) SetRaffleEscrowAccount(raffleEscrow ag_solanago.PublicKey) *SettleRaffleTicket {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(raffleEscrow).WRITE()
	return inst
}

// GetRaffleEscrowAccount gets the "raffleEscrow" account.
// ATA where SPL tokens are held in escrow, owned by raffle_ticket
func (inst *SettleRaffleTicket) GetRaffleEscrowAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetPayToAccount sets the "payTo" account.
// ATA owned by wallet authority that we pay to if we win
func (inst *SettleRaffleTicket) SetPayToAccount(payTo ag_solanago.PublicKey) *SettleRaffleTicket {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(payTo).WRITE()
	return inst
}

// GetPayToAccount gets the "payTo" account.
// ATA owned by wallet authority that we pay to if we win
func (inst *SettleRaffleTicket) GetPayToAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetRefundToAccount sets the "refundTo" account.
// or ATA owned by raffle_payer to refund SPL
func (inst *SettleRaffleTicket) SetRefundToAccount(refundTo ag_solanago.PublicKey) *SettleRaffleTicket {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(refundTo).WRITE()
	return inst
}

// GetRefundToAccount gets the "refundTo" account.
// or ATA owned by raffle_payer to refund SPL
func (inst *SettleRaffleTicket) GetRefundToAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetTokenAtaAccount sets the "tokenAta" account.
func (inst *SettleRaffleTicket) SetTokenAtaAccount(tokenAta ag_solanago.PublicKey) *SettleRaffleTicket {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(tokenAta).WRITE()
	return inst
}

// GetTokenAtaAccount gets the "tokenAta" account.
func (inst *SettleRaffleTicket) GetTokenAtaAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetRafflePayerAccount sets the "rafflePayer" account.
func (inst *SettleRaffleTicket) SetRafflePayerAccount(rafflePayer ag_solanago.PublicKey) *SettleRaffleTicket {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(rafflePayer).WRITE()
	return inst
}

// GetRafflePayerAccount gets the "rafflePayer" account.
func (inst *SettleRaffleTicket) GetRafflePayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetOrderInfoAccount sets the "orderInfo" account.
func (inst *SettleRaffleTicket) SetOrderInfoAccount(orderInfo ag_solanago.PublicKey) *SettleRaffleTicket {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(orderInfo).WRITE()
	return inst
}

// GetOrderInfoAccount gets the "orderInfo" account.
func (inst *SettleRaffleTicket) GetOrderInfoAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetNotaryAccount sets the "notary" account.
func (inst *SettleRaffleTicket) SetNotaryAccount(notary ag_solanago.PublicKey) *SettleRaffleTicket {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(notary).SIGNER()
	return inst
}

// GetNotaryAccount gets the "notary" account.
func (inst *SettleRaffleTicket) GetNotaryAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetMetadataAccount sets the "metadata" account.
func (inst *SettleRaffleTicket) SetMetadataAccount(metadata ag_solanago.PublicKey) *SettleRaffleTicket {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(metadata).WRITE()
	return inst
}

// GetMetadataAccount gets the "metadata" account.
func (inst *SettleRaffleTicket) GetMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

// SetMintAccount sets the "mint" account.
func (inst *SettleRaffleTicket) SetMintAccount(mint ag_solanago.PublicKey) *SettleRaffleTicket {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(mint).WRITE().SIGNER()
	return inst
}

// GetMintAccount gets the "mint" account.
func (inst *SettleRaffleTicket) GetMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(13)
}

// SetUpdateAuthorityAccount sets the "updateAuthority" account.
func (inst *SettleRaffleTicket) SetUpdateAuthorityAccount(updateAuthority ag_solanago.PublicKey) *SettleRaffleTicket {
	inst.AccountMetaSlice[14] = ag_solanago.Meta(updateAuthority).SIGNER()
	return inst
}

// GetUpdateAuthorityAccount gets the "updateAuthority" account.
func (inst *SettleRaffleTicket) GetUpdateAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(14)
}

// SetMasterEditionAccount sets the "masterEdition" account.
func (inst *SettleRaffleTicket) SetMasterEditionAccount(masterEdition ag_solanago.PublicKey) *SettleRaffleTicket {
	inst.AccountMetaSlice[15] = ag_solanago.Meta(masterEdition).WRITE()
	return inst
}

// GetMasterEditionAccount gets the "masterEdition" account.
func (inst *SettleRaffleTicket) GetMasterEditionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(15)
}

// SetSlotHashesAccount sets the "slotHashes" account.
func (inst *SettleRaffleTicket) SetSlotHashesAccount(slotHashes ag_solanago.PublicKey) *SettleRaffleTicket {
	inst.AccountMetaSlice[16] = ag_solanago.Meta(slotHashes)
	return inst
}

// GetSlotHashesAccount gets the "slotHashes" account.
func (inst *SettleRaffleTicket) GetSlotHashesAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(16)
}

// SetTokenMetadataProgramAccount sets the "tokenMetadataProgram" account.
func (inst *SettleRaffleTicket) SetTokenMetadataProgramAccount(tokenMetadataProgram ag_solanago.PublicKey) *SettleRaffleTicket {
	inst.AccountMetaSlice[17] = ag_solanago.Meta(tokenMetadataProgram)
	return inst
}

// GetTokenMetadataProgramAccount gets the "tokenMetadataProgram" account.
func (inst *SettleRaffleTicket) GetTokenMetadataProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(17)
}

// SetAssociatedTokenProgramAccount sets the "associatedTokenProgram" account.
func (inst *SettleRaffleTicket) SetAssociatedTokenProgramAccount(associatedTokenProgram ag_solanago.PublicKey) *SettleRaffleTicket {
	inst.AccountMetaSlice[18] = ag_solanago.Meta(associatedTokenProgram)
	return inst
}

// GetAssociatedTokenProgramAccount gets the "associatedTokenProgram" account.
func (inst *SettleRaffleTicket) GetAssociatedTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(18)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *SettleRaffleTicket) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *SettleRaffleTicket {
	inst.AccountMetaSlice[19] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *SettleRaffleTicket) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(19)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *SettleRaffleTicket) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *SettleRaffleTicket {
	inst.AccountMetaSlice[20] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *SettleRaffleTicket) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(20)
}

// SetRentAccount sets the "rent" account.
func (inst *SettleRaffleTicket) SetRentAccount(rent ag_solanago.PublicKey) *SettleRaffleTicket {
	inst.AccountMetaSlice[21] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *SettleRaffleTicket) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(21)
}

func (inst SettleRaffleTicket) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_SettleRaffleTicket,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst SettleRaffleTicket) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *SettleRaffleTicket) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.CurrTime == nil {
			return errors.New("CurrTime parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Config is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.CandyMachine is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.LaunchStagesInfo is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.RaffleTicket is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.RaffleEscrow is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.PayTo is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.RefundTo is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.TokenAta is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.RafflePayer is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.OrderInfo is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.Notary is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.Metadata is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.Mint is not set")
		}
		if inst.AccountMetaSlice[14] == nil {
			return errors.New("accounts.UpdateAuthority is not set")
		}
		if inst.AccountMetaSlice[15] == nil {
			return errors.New("accounts.MasterEdition is not set")
		}
		if inst.AccountMetaSlice[16] == nil {
			return errors.New("accounts.SlotHashes is not set")
		}
		if inst.AccountMetaSlice[17] == nil {
			return errors.New("accounts.TokenMetadataProgram is not set")
		}
		if inst.AccountMetaSlice[18] == nil {
			return errors.New("accounts.AssociatedTokenProgram is not set")
		}
		if inst.AccountMetaSlice[19] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[20] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[21] == nil {
			return errors.New("accounts.Rent is not set")
		}
	}
	return nil
}

func (inst *SettleRaffleTicket) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("SettleRaffleTicket")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("CurrTime", *inst.CurrTime))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=22]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                config", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("          candyMachine", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("                 payer", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("      launchStagesInfo", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("          raffleTicket", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("          raffleEscrow", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("                 payTo", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("              refundTo", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("              tokenAta", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("           rafflePayer", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("             orderInfo", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("                notary", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("              metadata", inst.AccountMetaSlice.Get(12)))
						accountsBranch.Child(ag_format.Meta("                  mint", inst.AccountMetaSlice.Get(13)))
						accountsBranch.Child(ag_format.Meta("       updateAuthority", inst.AccountMetaSlice.Get(14)))
						accountsBranch.Child(ag_format.Meta("         masterEdition", inst.AccountMetaSlice.Get(15)))
						accountsBranch.Child(ag_format.Meta("            slotHashes", inst.AccountMetaSlice.Get(16)))
						accountsBranch.Child(ag_format.Meta("  tokenMetadataProgram", inst.AccountMetaSlice.Get(17)))
						accountsBranch.Child(ag_format.Meta("associatedTokenProgram", inst.AccountMetaSlice.Get(18)))
						accountsBranch.Child(ag_format.Meta("          tokenProgram", inst.AccountMetaSlice.Get(19)))
						accountsBranch.Child(ag_format.Meta("         systemProgram", inst.AccountMetaSlice.Get(20)))
						accountsBranch.Child(ag_format.Meta("                  rent", inst.AccountMetaSlice.Get(21)))
					})
				})
		})
}

func (obj SettleRaffleTicket) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `CurrTime` param:
	err = encoder.Encode(obj.CurrTime)
	if err != nil {
		return err
	}
	return nil
}
func (obj *SettleRaffleTicket) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `CurrTime`:
	err = decoder.Decode(&obj.CurrTime)
	if err != nil {
		return err
	}
	return nil
}

// NewSettleRaffleTicketInstruction declares a new SettleRaffleTicket instruction with the provided parameters and accounts.
func NewSettleRaffleTicketInstruction(
	// Parameters:
	currTime int64,
	// Accounts:
	config ag_solanago.PublicKey,
	candyMachine ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	launchStagesInfo ag_solanago.PublicKey,
	raffleTicket ag_solanago.PublicKey,
	raffleEscrow ag_solanago.PublicKey,
	payTo ag_solanago.PublicKey,
	refundTo ag_solanago.PublicKey,
	tokenAta ag_solanago.PublicKey,
	rafflePayer ag_solanago.PublicKey,
	orderInfo ag_solanago.PublicKey,
	notary ag_solanago.PublicKey,
	metadata ag_solanago.PublicKey,
	mint ag_solanago.PublicKey,
	updateAuthority ag_solanago.PublicKey,
	masterEdition ag_solanago.PublicKey,
	slotHashes ag_solanago.PublicKey,
	tokenMetadataProgram ag_solanago.PublicKey,
	associatedTokenProgram ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	rent ag_solanago.PublicKey) *SettleRaffleTicket {
	return NewSettleRaffleTicketInstructionBuilder().
		SetCurrTime(currTime).
		SetConfigAccount(config).
		SetCandyMachineAccount(candyMachine).
		SetPayerAccount(payer).
		SetLaunchStagesInfoAccount(launchStagesInfo).
		SetRaffleTicketAccount(raffleTicket).
		SetRaffleEscrowAccount(raffleEscrow).
		SetPayToAccount(payTo).
		SetRefundToAccount(refundTo).
		SetTokenAtaAccount(tokenAta).
		SetRafflePayerAccount(rafflePayer).
		SetOrderInfoAccount(orderInfo).
		SetNotaryAccount(notary).
		SetMetadataAccount(metadata).
		SetMintAccount(mint).
		SetUpdateAuthorityAccount(updateAuthority).
		SetMasterEditionAccount(masterEdition).
		SetSlotHashesAccount(slotHashes).
		SetTokenMetadataProgramAccount(tokenMetadataProgram).
		SetAssociatedTokenProgramAccount(associatedTokenProgram).
		SetTokenProgramAccount(tokenProgram).
		SetSystemProgramAccount(systemProgram).
		SetRentAccount(rent)
}
