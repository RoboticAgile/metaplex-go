// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package me_launchpad

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/RoboticAgile/solana-go"
	ag_format "github.com/RoboticAgile/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// MintNftMip1 is the `mintNftMip1` instruction.
type MintNftMip1 struct {
	WalletLimitBump *uint8
	InOrder         *bool
	UserLimit       *uint8 `bin:"optional"`
	CurrTime        *int64

	// [0] = [] config
	//
	// [1] = [WRITE] candyMachine
	//
	// [2] = [] mintReceiver
	//
	// [3] = [WRITE] candyMachineWalletAuthority
	//
	// [4] = [WRITE, SIGNER] payer
	//
	// [5] = [WRITE] launchStagesInfo
	//
	// [6] = [WRITE] payFrom
	//
	// [7] = [WRITE] payTo
	//
	// [8] = [] notary
	//
	// [9] = [WRITE] metadata
	//
	// [10] = [WRITE, SIGNER] mint
	//
	// [11] = [WRITE] tokenAta
	//
	// [12] = [WRITE] masterEdition
	//
	// [13] = [WRITE] tokenRecord
	//
	// [14] = [WRITE] walletLimitInfo
	//
	// [15] = [WRITE] orderInfo
	//
	// [16] = [] slotHashes
	//
	// [17] = [] tokenMetadataProgram
	//
	// [18] = [] tokenProgram
	//
	// [19] = [] systemProgram
	//
	// [20] = [] associatedTokenProgram
	//
	// [21] = [] ruleSet
	//
	// [22] = [] authorizationRulesProgram
	//
	// [23] = [] instructions
	//
	// [24] = [] rent
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewMintNftMip1InstructionBuilder creates a new `MintNftMip1` instruction builder.
func NewMintNftMip1InstructionBuilder() *MintNftMip1 {
	nd := &MintNftMip1{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 25),
	}
	return nd
}

// SetWalletLimitBump sets the "walletLimitBump" parameter.
func (inst *MintNftMip1) SetWalletLimitBump(walletLimitBump uint8) *MintNftMip1 {
	inst.WalletLimitBump = &walletLimitBump
	return inst
}

// SetInOrder sets the "inOrder" parameter.
func (inst *MintNftMip1) SetInOrder(inOrder bool) *MintNftMip1 {
	inst.InOrder = &inOrder
	return inst
}

// SetUserLimit sets the "userLimit" parameter.
func (inst *MintNftMip1) SetUserLimit(userLimit uint8) *MintNftMip1 {
	inst.UserLimit = &userLimit
	return inst
}

// SetCurrTime sets the "currTime" parameter.
func (inst *MintNftMip1) SetCurrTime(currTime int64) *MintNftMip1 {
	inst.CurrTime = &currTime
	return inst
}

// SetConfigAccount sets the "config" account.
func (inst *MintNftMip1) SetConfigAccount(config ag_solanago.PublicKey) *MintNftMip1 {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(config)
	return inst
}

// GetConfigAccount gets the "config" account.
func (inst *MintNftMip1) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetCandyMachineAccount sets the "candyMachine" account.
func (inst *MintNftMip1) SetCandyMachineAccount(candyMachine ag_solanago.PublicKey) *MintNftMip1 {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(candyMachine).WRITE()
	return inst
}

// GetCandyMachineAccount gets the "candyMachine" account.
func (inst *MintNftMip1) GetCandyMachineAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetMintReceiverAccount sets the "mintReceiver" account.
func (inst *MintNftMip1) SetMintReceiverAccount(mintReceiver ag_solanago.PublicKey) *MintNftMip1 {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(mintReceiver)
	return inst
}

// GetMintReceiverAccount gets the "mintReceiver" account.
func (inst *MintNftMip1) GetMintReceiverAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetCandyMachineWalletAuthorityAccount sets the "candyMachineWalletAuthority" account.
func (inst *MintNftMip1) SetCandyMachineWalletAuthorityAccount(candyMachineWalletAuthority ag_solanago.PublicKey) *MintNftMip1 {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(candyMachineWalletAuthority).WRITE()
	return inst
}

// GetCandyMachineWalletAuthorityAccount gets the "candyMachineWalletAuthority" account.
func (inst *MintNftMip1) GetCandyMachineWalletAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetPayerAccount sets the "payer" account.
func (inst *MintNftMip1) SetPayerAccount(payer ag_solanago.PublicKey) *MintNftMip1 {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(payer).WRITE().SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
func (inst *MintNftMip1) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetLaunchStagesInfoAccount sets the "launchStagesInfo" account.
func (inst *MintNftMip1) SetLaunchStagesInfoAccount(launchStagesInfo ag_solanago.PublicKey) *MintNftMip1 {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(launchStagesInfo).WRITE()
	return inst
}

// GetLaunchStagesInfoAccount gets the "launchStagesInfo" account.
func (inst *MintNftMip1) GetLaunchStagesInfoAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetPayFromAccount sets the "payFrom" account.
func (inst *MintNftMip1) SetPayFromAccount(payFrom ag_solanago.PublicKey) *MintNftMip1 {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(payFrom).WRITE()
	return inst
}

// GetPayFromAccount gets the "payFrom" account.
func (inst *MintNftMip1) GetPayFromAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetPayToAccount sets the "payTo" account.
func (inst *MintNftMip1) SetPayToAccount(payTo ag_solanago.PublicKey) *MintNftMip1 {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(payTo).WRITE()
	return inst
}

// GetPayToAccount gets the "payTo" account.
func (inst *MintNftMip1) GetPayToAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetNotaryAccount sets the "notary" account.
func (inst *MintNftMip1) SetNotaryAccount(notary ag_solanago.PublicKey) *MintNftMip1 {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(notary)
	return inst
}

// GetNotaryAccount gets the "notary" account.
func (inst *MintNftMip1) GetNotaryAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetMetadataAccount sets the "metadata" account.
func (inst *MintNftMip1) SetMetadataAccount(metadata ag_solanago.PublicKey) *MintNftMip1 {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(metadata).WRITE()
	return inst
}

// GetMetadataAccount gets the "metadata" account.
func (inst *MintNftMip1) GetMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetMintAccount sets the "mint" account.
func (inst *MintNftMip1) SetMintAccount(mint ag_solanago.PublicKey) *MintNftMip1 {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(mint).WRITE().SIGNER()
	return inst
}

// GetMintAccount gets the "mint" account.
func (inst *MintNftMip1) GetMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetTokenAtaAccount sets the "tokenAta" account.
func (inst *MintNftMip1) SetTokenAtaAccount(tokenAta ag_solanago.PublicKey) *MintNftMip1 {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(tokenAta).WRITE()
	return inst
}

// GetTokenAtaAccount gets the "tokenAta" account.
func (inst *MintNftMip1) GetTokenAtaAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetMasterEditionAccount sets the "masterEdition" account.
func (inst *MintNftMip1) SetMasterEditionAccount(masterEdition ag_solanago.PublicKey) *MintNftMip1 {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(masterEdition).WRITE()
	return inst
}

// GetMasterEditionAccount gets the "masterEdition" account.
func (inst *MintNftMip1) GetMasterEditionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

// SetTokenRecordAccount sets the "tokenRecord" account.
func (inst *MintNftMip1) SetTokenRecordAccount(tokenRecord ag_solanago.PublicKey) *MintNftMip1 {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(tokenRecord).WRITE()
	return inst
}

// GetTokenRecordAccount gets the "tokenRecord" account.
func (inst *MintNftMip1) GetTokenRecordAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(13)
}

// SetWalletLimitInfoAccount sets the "walletLimitInfo" account.
func (inst *MintNftMip1) SetWalletLimitInfoAccount(walletLimitInfo ag_solanago.PublicKey) *MintNftMip1 {
	inst.AccountMetaSlice[14] = ag_solanago.Meta(walletLimitInfo).WRITE()
	return inst
}

// GetWalletLimitInfoAccount gets the "walletLimitInfo" account.
func (inst *MintNftMip1) GetWalletLimitInfoAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(14)
}

// SetOrderInfoAccount sets the "orderInfo" account.
func (inst *MintNftMip1) SetOrderInfoAccount(orderInfo ag_solanago.PublicKey) *MintNftMip1 {
	inst.AccountMetaSlice[15] = ag_solanago.Meta(orderInfo).WRITE()
	return inst
}

// GetOrderInfoAccount gets the "orderInfo" account.
func (inst *MintNftMip1) GetOrderInfoAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(15)
}

// SetSlotHashesAccount sets the "slotHashes" account.
func (inst *MintNftMip1) SetSlotHashesAccount(slotHashes ag_solanago.PublicKey) *MintNftMip1 {
	inst.AccountMetaSlice[16] = ag_solanago.Meta(slotHashes)
	return inst
}

// GetSlotHashesAccount gets the "slotHashes" account.
func (inst *MintNftMip1) GetSlotHashesAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(16)
}

// SetTokenMetadataProgramAccount sets the "tokenMetadataProgram" account.
func (inst *MintNftMip1) SetTokenMetadataProgramAccount(tokenMetadataProgram ag_solanago.PublicKey) *MintNftMip1 {
	inst.AccountMetaSlice[17] = ag_solanago.Meta(tokenMetadataProgram)
	return inst
}

// GetTokenMetadataProgramAccount gets the "tokenMetadataProgram" account.
func (inst *MintNftMip1) GetTokenMetadataProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(17)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *MintNftMip1) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *MintNftMip1 {
	inst.AccountMetaSlice[18] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *MintNftMip1) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(18)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *MintNftMip1) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *MintNftMip1 {
	inst.AccountMetaSlice[19] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *MintNftMip1) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(19)
}

// SetAssociatedTokenProgramAccount sets the "associatedTokenProgram" account.
func (inst *MintNftMip1) SetAssociatedTokenProgramAccount(associatedTokenProgram ag_solanago.PublicKey) *MintNftMip1 {
	inst.AccountMetaSlice[20] = ag_solanago.Meta(associatedTokenProgram)
	return inst
}

// GetAssociatedTokenProgramAccount gets the "associatedTokenProgram" account.
func (inst *MintNftMip1) GetAssociatedTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(20)
}

// SetRuleSetAccount sets the "ruleSet" account.
func (inst *MintNftMip1) SetRuleSetAccount(ruleSet ag_solanago.PublicKey) *MintNftMip1 {
	inst.AccountMetaSlice[21] = ag_solanago.Meta(ruleSet)
	return inst
}

// GetRuleSetAccount gets the "ruleSet" account.
func (inst *MintNftMip1) GetRuleSetAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(21)
}

// SetAuthorizationRulesProgramAccount sets the "authorizationRulesProgram" account.
func (inst *MintNftMip1) SetAuthorizationRulesProgramAccount(authorizationRulesProgram ag_solanago.PublicKey) *MintNftMip1 {
	inst.AccountMetaSlice[22] = ag_solanago.Meta(authorizationRulesProgram)
	return inst
}

// GetAuthorizationRulesProgramAccount gets the "authorizationRulesProgram" account.
func (inst *MintNftMip1) GetAuthorizationRulesProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(22)
}

// SetInstructionsAccount sets the "instructions" account.
func (inst *MintNftMip1) SetInstructionsAccount(instructions ag_solanago.PublicKey) *MintNftMip1 {
	inst.AccountMetaSlice[23] = ag_solanago.Meta(instructions)
	return inst
}

// GetInstructionsAccount gets the "instructions" account.
func (inst *MintNftMip1) GetInstructionsAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(23)
}

// SetRentAccount sets the "rent" account.
func (inst *MintNftMip1) SetRentAccount(rent ag_solanago.PublicKey) *MintNftMip1 {
	inst.AccountMetaSlice[24] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *MintNftMip1) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(24)
}

func (inst MintNftMip1) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_MintNftMip1,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst MintNftMip1) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *MintNftMip1) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.WalletLimitBump == nil {
			return errors.New("WalletLimitBump parameter is not set")
		}
		if inst.InOrder == nil {
			return errors.New("InOrder parameter is not set")
		}
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
			return errors.New("accounts.MintReceiver is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.CandyMachineWalletAuthority is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.LaunchStagesInfo is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.PayFrom is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.PayTo is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.Notary is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.Metadata is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.Mint is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.TokenAta is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.MasterEdition is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.TokenRecord is not set")
		}
		if inst.AccountMetaSlice[14] == nil {
			return errors.New("accounts.WalletLimitInfo is not set")
		}
		if inst.AccountMetaSlice[15] == nil {
			return errors.New("accounts.OrderInfo is not set")
		}
		if inst.AccountMetaSlice[16] == nil {
			return errors.New("accounts.SlotHashes is not set")
		}
		if inst.AccountMetaSlice[17] == nil {
			return errors.New("accounts.TokenMetadataProgram is not set")
		}
		if inst.AccountMetaSlice[18] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[19] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[20] == nil {
			return errors.New("accounts.AssociatedTokenProgram is not set")
		}
		if inst.AccountMetaSlice[21] == nil {
			return errors.New("accounts.RuleSet is not set")
		}
		if inst.AccountMetaSlice[22] == nil {
			return errors.New("accounts.AuthorizationRulesProgram is not set")
		}
		if inst.AccountMetaSlice[23] == nil {
			return errors.New("accounts.Instructions is not set")
		}
		if inst.AccountMetaSlice[24] == nil {
			return errors.New("accounts.Rent is not set")
		}
	}
	return nil
}

func (inst *MintNftMip1) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("MintNftMip1")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=4]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("WalletLimitBump", *inst.WalletLimitBump))
						paramsBranch.Child(ag_format.Param("        InOrder", *inst.InOrder))
						paramsBranch.Child(ag_format.Param("      UserLimit (OPT)", inst.UserLimit))
						paramsBranch.Child(ag_format.Param("       CurrTime", *inst.CurrTime))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=25]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                     config", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("               candyMachine", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("               mintReceiver", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("candyMachineWalletAuthority", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("                      payer", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("           launchStagesInfo", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("                    payFrom", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("                      payTo", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("                     notary", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("                   metadata", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("                       mint", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("                   tokenAta", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("              masterEdition", inst.AccountMetaSlice.Get(12)))
						accountsBranch.Child(ag_format.Meta("                tokenRecord", inst.AccountMetaSlice.Get(13)))
						accountsBranch.Child(ag_format.Meta("            walletLimitInfo", inst.AccountMetaSlice.Get(14)))
						accountsBranch.Child(ag_format.Meta("                  orderInfo", inst.AccountMetaSlice.Get(15)))
						accountsBranch.Child(ag_format.Meta("                 slotHashes", inst.AccountMetaSlice.Get(16)))
						accountsBranch.Child(ag_format.Meta("       tokenMetadataProgram", inst.AccountMetaSlice.Get(17)))
						accountsBranch.Child(ag_format.Meta("               tokenProgram", inst.AccountMetaSlice.Get(18)))
						accountsBranch.Child(ag_format.Meta("              systemProgram", inst.AccountMetaSlice.Get(19)))
						accountsBranch.Child(ag_format.Meta("     associatedTokenProgram", inst.AccountMetaSlice.Get(20)))
						accountsBranch.Child(ag_format.Meta("                    ruleSet", inst.AccountMetaSlice.Get(21)))
						accountsBranch.Child(ag_format.Meta("  authorizationRulesProgram", inst.AccountMetaSlice.Get(22)))
						accountsBranch.Child(ag_format.Meta("               instructions", inst.AccountMetaSlice.Get(23)))
						accountsBranch.Child(ag_format.Meta("                       rent", inst.AccountMetaSlice.Get(24)))
					})
				})
		})
}

func (obj MintNftMip1) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `WalletLimitBump` param:
	err = encoder.Encode(obj.WalletLimitBump)
	if err != nil {
		return err
	}
	// Serialize `InOrder` param:
	err = encoder.Encode(obj.InOrder)
	if err != nil {
		return err
	}
	// Serialize `UserLimit` param (optional):
	{
		if obj.UserLimit == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.UserLimit)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `CurrTime` param:
	err = encoder.Encode(obj.CurrTime)
	if err != nil {
		return err
	}
	return nil
}
func (obj *MintNftMip1) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `WalletLimitBump`:
	err = decoder.Decode(&obj.WalletLimitBump)
	if err != nil {
		return err
	}
	// Deserialize `InOrder`:
	err = decoder.Decode(&obj.InOrder)
	if err != nil {
		return err
	}
	// Deserialize `UserLimit` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.UserLimit)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `CurrTime`:
	err = decoder.Decode(&obj.CurrTime)
	if err != nil {
		return err
	}
	return nil
}

// NewMintNftMip1Instruction declares a new MintNftMip1 instruction with the provided parameters and accounts.
func NewMintNftMip1Instruction(
	// Parameters:
	walletLimitBump uint8,
	inOrder bool,
	userLimit uint8,
	currTime int64,
	// Accounts:
	config ag_solanago.PublicKey,
	candyMachine ag_solanago.PublicKey,
	mintReceiver ag_solanago.PublicKey,
	candyMachineWalletAuthority ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	launchStagesInfo ag_solanago.PublicKey,
	payFrom ag_solanago.PublicKey,
	payTo ag_solanago.PublicKey,
	notary ag_solanago.PublicKey,
	metadata ag_solanago.PublicKey,
	mint ag_solanago.PublicKey,
	tokenAta ag_solanago.PublicKey,
	masterEdition ag_solanago.PublicKey,
	tokenRecord ag_solanago.PublicKey,
	walletLimitInfo ag_solanago.PublicKey,
	orderInfo ag_solanago.PublicKey,
	slotHashes ag_solanago.PublicKey,
	tokenMetadataProgram ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	associatedTokenProgram ag_solanago.PublicKey,
	ruleSet ag_solanago.PublicKey,
	authorizationRulesProgram ag_solanago.PublicKey,
	instructions ag_solanago.PublicKey,
	rent ag_solanago.PublicKey) *MintNftMip1 {
	return NewMintNftMip1InstructionBuilder().
		SetWalletLimitBump(walletLimitBump).
		SetInOrder(inOrder).
		SetUserLimit(userLimit).
		SetCurrTime(currTime).
		SetConfigAccount(config).
		SetCandyMachineAccount(candyMachine).
		SetMintReceiverAccount(mintReceiver).
		SetCandyMachineWalletAuthorityAccount(candyMachineWalletAuthority).
		SetPayerAccount(payer).
		SetLaunchStagesInfoAccount(launchStagesInfo).
		SetPayFromAccount(payFrom).
		SetPayToAccount(payTo).
		SetNotaryAccount(notary).
		SetMetadataAccount(metadata).
		SetMintAccount(mint).
		SetTokenAtaAccount(tokenAta).
		SetMasterEditionAccount(masterEdition).
		SetTokenRecordAccount(tokenRecord).
		SetWalletLimitInfoAccount(walletLimitInfo).
		SetOrderInfoAccount(orderInfo).
		SetSlotHashesAccount(slotHashes).
		SetTokenMetadataProgramAccount(tokenMetadataProgram).
		SetTokenProgramAccount(tokenProgram).
		SetSystemProgramAccount(systemProgram).
		SetAssociatedTokenProgramAccount(associatedTokenProgram).
		SetRuleSetAccount(ruleSet).
		SetAuthorizationRulesProgramAccount(authorizationRulesProgram).
		SetInstructionsAccount(instructions).
		SetRentAccount(rent)
}
