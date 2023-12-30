// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package auction_house

import (
	"bytes"
	"fmt"
	ag_spew "github.com/davecgh/go-spew/spew"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/RoboticAgile/solana-go"
	ag_text "github.com/RoboticAgile/solana-go/text"
	ag_treeout "github.com/gagliardetto/treeout"
)

var ProgramID ag_solanago.PublicKey = ag_solanago.MustPublicKeyFromBase58("hausS13jsjafwWwGqZTUQRmWyvyxn9EQpqMwV1PBBmk")

func SetProgramID(pubkey ag_solanago.PublicKey) {
	ProgramID = pubkey
	ag_solanago.RegisterInstructionDecoder(ProgramID, registryDecodeInstruction)
}

const ProgramName = "AuctionHouse"

func init() {
	if !ProgramID.IsZero() {
		ag_solanago.RegisterInstructionDecoder(ProgramID, registryDecodeInstruction)
	}
}

var (
	Instruction_WithdrawFromFee = ag_binary.TypeID([8]byte{179, 208, 190, 154, 32, 179, 19, 59})

	Instruction_WithdrawFromTreasury = ag_binary.TypeID([8]byte{0, 164, 86, 76, 56, 72, 12, 170})

	Instruction_UpdateAuctionHouse = ag_binary.TypeID([8]byte{84, 215, 2, 172, 241, 0, 245, 219})

	Instruction_CreateAuctionHouse = ag_binary.TypeID([8]byte{221, 66, 242, 159, 249, 206, 134, 241})

	Instruction_Buy = ag_binary.TypeID([8]byte{102, 6, 61, 18, 1, 218, 235, 234})

	Instruction_AuctioneerBuy = ag_binary.TypeID([8]byte{17, 106, 133, 46, 229, 48, 45, 208})

	Instruction_PublicBuy = ag_binary.TypeID([8]byte{169, 84, 218, 35, 42, 206, 16, 171})

	Instruction_AuctioneerPublicBuy = ag_binary.TypeID([8]byte{221, 239, 99, 240, 86, 46, 213, 126})

	Instruction_Cancel = ag_binary.TypeID([8]byte{232, 219, 223, 41, 219, 236, 220, 190})

	Instruction_AuctioneerCancel = ag_binary.TypeID([8]byte{197, 97, 152, 196, 115, 204, 64, 215})

	Instruction_Deposit = ag_binary.TypeID([8]byte{242, 35, 198, 137, 82, 225, 242, 182})

	Instruction_AuctioneerDeposit = ag_binary.TypeID([8]byte{79, 122, 37, 162, 120, 173, 57, 127})

	Instruction_ExecuteSale = ag_binary.TypeID([8]byte{37, 74, 217, 157, 79, 49, 35, 6})

	Instruction_ExecutePartialSale = ag_binary.TypeID([8]byte{163, 18, 35, 157, 49, 164, 203, 133})

	Instruction_AuctioneerExecuteSale = ag_binary.TypeID([8]byte{68, 125, 32, 65, 251, 43, 35, 53})

	Instruction_AuctioneerExecutePartialSale = ag_binary.TypeID([8]byte{9, 44, 46, 15, 161, 143, 21, 54})

	Instruction_Sell = ag_binary.TypeID([8]byte{51, 230, 133, 164, 1, 127, 131, 173})

	Instruction_AuctioneerSell = ag_binary.TypeID([8]byte{251, 60, 142, 195, 121, 203, 26, 183})

	Instruction_Withdraw = ag_binary.TypeID([8]byte{183, 18, 70, 156, 148, 109, 161, 34})

	Instruction_AuctioneerWithdraw = ag_binary.TypeID([8]byte{85, 166, 219, 110, 168, 143, 180, 236})

	Instruction_CloseEscrowAccount = ag_binary.TypeID([8]byte{209, 42, 208, 179, 140, 78, 18, 43})

	Instruction_DelegateAuctioneer = ag_binary.TypeID([8]byte{106, 178, 12, 122, 74, 173, 251, 222})

	Instruction_UpdateAuctioneer = ag_binary.TypeID([8]byte{103, 255, 80, 234, 94, 56, 168, 208})

	Instruction_PrintListingReceipt = ag_binary.TypeID([8]byte{207, 107, 44, 160, 75, 222, 195, 27})

	Instruction_CancelListingReceipt = ag_binary.TypeID([8]byte{171, 59, 138, 126, 246, 189, 91, 11})

	Instruction_PrintBidReceipt = ag_binary.TypeID([8]byte{94, 249, 90, 230, 239, 64, 68, 218})

	Instruction_CancelBidReceipt = ag_binary.TypeID([8]byte{246, 108, 27, 229, 220, 42, 176, 43})

	Instruction_PrintPurchaseReceipt = ag_binary.TypeID([8]byte{227, 154, 251, 7, 180, 56, 100, 143})
)

// InstructionIDToName returns the name of the instruction given its ID.
func InstructionIDToName(id ag_binary.TypeID) string {
	switch id {
	case Instruction_WithdrawFromFee:
		return "WithdrawFromFee"
	case Instruction_WithdrawFromTreasury:
		return "WithdrawFromTreasury"
	case Instruction_UpdateAuctionHouse:
		return "UpdateAuctionHouse"
	case Instruction_CreateAuctionHouse:
		return "CreateAuctionHouse"
	case Instruction_Buy:
		return "Buy"
	case Instruction_AuctioneerBuy:
		return "AuctioneerBuy"
	case Instruction_PublicBuy:
		return "PublicBuy"
	case Instruction_AuctioneerPublicBuy:
		return "AuctioneerPublicBuy"
	case Instruction_Cancel:
		return "Cancel"
	case Instruction_AuctioneerCancel:
		return "AuctioneerCancel"
	case Instruction_Deposit:
		return "Deposit"
	case Instruction_AuctioneerDeposit:
		return "AuctioneerDeposit"
	case Instruction_ExecuteSale:
		return "ExecuteSale"
	case Instruction_ExecutePartialSale:
		return "ExecutePartialSale"
	case Instruction_AuctioneerExecuteSale:
		return "AuctioneerExecuteSale"
	case Instruction_AuctioneerExecutePartialSale:
		return "AuctioneerExecutePartialSale"
	case Instruction_Sell:
		return "Sell"
	case Instruction_AuctioneerSell:
		return "AuctioneerSell"
	case Instruction_Withdraw:
		return "Withdraw"
	case Instruction_AuctioneerWithdraw:
		return "AuctioneerWithdraw"
	case Instruction_CloseEscrowAccount:
		return "CloseEscrowAccount"
	case Instruction_DelegateAuctioneer:
		return "DelegateAuctioneer"
	case Instruction_UpdateAuctioneer:
		return "UpdateAuctioneer"
	case Instruction_PrintListingReceipt:
		return "PrintListingReceipt"
	case Instruction_CancelListingReceipt:
		return "CancelListingReceipt"
	case Instruction_PrintBidReceipt:
		return "PrintBidReceipt"
	case Instruction_CancelBidReceipt:
		return "CancelBidReceipt"
	case Instruction_PrintPurchaseReceipt:
		return "PrintPurchaseReceipt"
	default:
		return ""
	}
}

type Instruction struct {
	ag_binary.BaseVariant
}

func (inst *Instruction) EncodeToTree(parent ag_treeout.Branches) {
	if enToTree, ok := inst.Impl.(ag_text.EncodableToTree); ok {
		enToTree.EncodeToTree(parent)
	} else {
		parent.Child(ag_spew.Sdump(inst))
	}
}

var InstructionImplDef = ag_binary.NewVariantDefinition(
	ag_binary.AnchorTypeIDEncoding,
	[]ag_binary.VariantType{
		{
			"withdraw_from_fee", (*WithdrawFromFee)(nil),
		},
		{
			"withdraw_from_treasury", (*WithdrawFromTreasury)(nil),
		},
		{
			"update_auction_house", (*UpdateAuctionHouse)(nil),
		},
		{
			"create_auction_house", (*CreateAuctionHouse)(nil),
		},
		{
			"buy", (*Buy)(nil),
		},
		{
			"auctioneer_buy", (*AuctioneerBuy)(nil),
		},
		{
			"public_buy", (*PublicBuy)(nil),
		},
		{
			"auctioneer_public_buy", (*AuctioneerPublicBuy)(nil),
		},
		{
			"cancel", (*Cancel)(nil),
		},
		{
			"auctioneer_cancel", (*AuctioneerCancel)(nil),
		},
		{
			"deposit", (*Deposit)(nil),
		},
		{
			"auctioneer_deposit", (*AuctioneerDeposit)(nil),
		},
		{
			"execute_sale", (*ExecuteSale)(nil),
		},
		{
			"execute_partial_sale", (*ExecutePartialSale)(nil),
		},
		{
			"auctioneer_execute_sale", (*AuctioneerExecuteSale)(nil),
		},
		{
			"auctioneer_execute_partial_sale", (*AuctioneerExecutePartialSale)(nil),
		},
		{
			"sell", (*Sell)(nil),
		},
		{
			"auctioneer_sell", (*AuctioneerSell)(nil),
		},
		{
			"withdraw", (*Withdraw)(nil),
		},
		{
			"auctioneer_withdraw", (*AuctioneerWithdraw)(nil),
		},
		{
			"close_escrow_account", (*CloseEscrowAccount)(nil),
		},
		{
			"delegate_auctioneer", (*DelegateAuctioneer)(nil),
		},
		{
			"update_auctioneer", (*UpdateAuctioneer)(nil),
		},
		{
			"print_listing_receipt", (*PrintListingReceipt)(nil),
		},
		{
			"cancel_listing_receipt", (*CancelListingReceipt)(nil),
		},
		{
			"print_bid_receipt", (*PrintBidReceipt)(nil),
		},
		{
			"cancel_bid_receipt", (*CancelBidReceipt)(nil),
		},
		{
			"print_purchase_receipt", (*PrintPurchaseReceipt)(nil),
		},
	},
)

func (inst *Instruction) ProgramID() ag_solanago.PublicKey {
	return ProgramID
}

func (inst *Instruction) Accounts() (out []*ag_solanago.AccountMeta) {
	return inst.Impl.(ag_solanago.AccountsGettable).GetAccounts()
}

func (inst *Instruction) Data() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := ag_binary.NewBorshEncoder(buf).Encode(inst); err != nil {
		return nil, fmt.Errorf("unable to encode instruction: %w", err)
	}
	return buf.Bytes(), nil
}

func (inst *Instruction) TextEncode(encoder *ag_text.Encoder, option *ag_text.Option) error {
	return encoder.Encode(inst.Impl, option)
}

func (inst *Instruction) UnmarshalWithDecoder(decoder *ag_binary.Decoder) error {
	return inst.BaseVariant.UnmarshalBinaryVariant(decoder, InstructionImplDef)
}

func (inst *Instruction) MarshalWithEncoder(encoder *ag_binary.Encoder) error {
	err := encoder.WriteBytes(inst.TypeID.Bytes(), false)
	if err != nil {
		return fmt.Errorf("unable to write variant type: %w", err)
	}
	return encoder.Encode(inst.Impl)
}

func registryDecodeInstruction(accounts []*ag_solanago.AccountMeta, data []byte) (interface{}, error) {
	inst, err := DecodeInstruction(accounts, data)
	if err != nil {
		return nil, err
	}
	return inst, nil
}

func DecodeInstruction(accounts []*ag_solanago.AccountMeta, data []byte) (*Instruction, error) {
	inst := new(Instruction)
	if err := ag_binary.NewBorshDecoder(data).Decode(inst); err != nil {
		return nil, fmt.Errorf("unable to decode instruction: %w", err)
	}
	if v, ok := inst.Impl.(ag_solanago.AccountsSettable); ok {
		err := v.SetAccounts(accounts)
		if err != nil {
			return nil, fmt.Errorf("unable to set accounts for instruction: %w", err)
		}
	}
	return inst, nil
}
