// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package nft_candy_machine

import (
	"bytes"
	"fmt"
	ag_spew "github.com/davecgh/go-spew/spew"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/RoboticAgile/solana-go"
	ag_text "github.com/RoboticAgile/solana-go/text"
	ag_treeout "github.com/gagliardetto/treeout"
)

var ProgramID ag_solanago.PublicKey

func SetProgramID(pubkey ag_solanago.PublicKey) {
	ProgramID = pubkey
	ag_solanago.RegisterInstructionDecoder(ProgramID, registryDecodeInstruction)
}

const ProgramName = "NftCandyMachine"

func init() {
	if !ProgramID.IsZero() {
		ag_solanago.RegisterInstructionDecoder(ProgramID, registryDecodeInstruction)
	}
}

var (
	Instruction_MintV6 = ag_binary.TypeID([8]byte{111, 169, 221, 193, 234, 227, 8, 180})

	Instruction_MintV5 = ag_binary.TypeID([8]byte{96, 114, 250, 235, 203, 205, 188, 36})

	Instruction_EditCmV5 = ag_binary.TypeID([8]byte{43, 25, 89, 26, 165, 209, 229, 92})

	Instruction_EditCmV6 = ag_binary.TypeID([8]byte{8, 151, 9, 92, 26, 236, 43, 194})

	Instruction_RevealV5 = ag_binary.TypeID([8]byte{57, 46, 3, 81, 70, 219, 168, 49})

	Instruction_RevealV6 = ag_binary.TypeID([8]byte{96, 31, 203, 62, 198, 112, 162, 73})

	Instruction_AllowUnfreezeV5 = ag_binary.TypeID([8]byte{156, 1, 166, 135, 249, 139, 95, 54})

	Instruction_AllowRevealV5 = ag_binary.TypeID([8]byte{13, 155, 159, 197, 5, 211, 116, 91})

	Instruction_BurnSupplyV5 = ag_binary.TypeID([8]byte{107, 75, 172, 207, 53, 177, 22, 255})

	Instruction_AllowUnfreezeV6 = ag_binary.TypeID([8]byte{218, 161, 193, 67, 230, 24, 153, 227})

	Instruction_AllowRevealV6 = ag_binary.TypeID([8]byte{26, 70, 81, 233, 255, 128, 149, 36})

	Instruction_BurnSupplyV6 = ag_binary.TypeID([8]byte{155, 195, 120, 109, 95, 193, 192, 151})

	Instruction_BurnFrozen = ag_binary.TypeID([8]byte{115, 215, 82, 186, 58, 93, 69, 24})

	Instruction_Migrate = ag_binary.TypeID([8]byte{155, 234, 231, 146, 236, 158, 162, 30})

	Instruction_ThawV5 = ag_binary.TypeID([8]byte{181, 47, 178, 189, 46, 187, 145, 248})

	Instruction_ThawV6 = ag_binary.TypeID([8]byte{235, 243, 216, 226, 135, 47, 129, 80})

	Instruction_InitCmV5 = ag_binary.TypeID([8]byte{179, 164, 16, 182, 42, 27, 249, 75})

	Instruction_InitCmV6 = ag_binary.TypeID([8]byte{82, 15, 232, 184, 21, 89, 180, 143})
)

// InstructionIDToName returns the name of the instruction given its ID.
func InstructionIDToName(id ag_binary.TypeID) string {
	switch id {
	case Instruction_MintV6:
		return "MintV6"
	case Instruction_MintV5:
		return "MintV5"
	case Instruction_EditCmV5:
		return "EditCmV5"
	case Instruction_EditCmV6:
		return "EditCmV6"
	case Instruction_RevealV5:
		return "RevealV5"
	case Instruction_RevealV6:
		return "RevealV6"
	case Instruction_AllowUnfreezeV5:
		return "AllowUnfreezeV5"
	case Instruction_AllowRevealV5:
		return "AllowRevealV5"
	case Instruction_BurnSupplyV5:
		return "BurnSupplyV5"
	case Instruction_AllowUnfreezeV6:
		return "AllowUnfreezeV6"
	case Instruction_AllowRevealV6:
		return "AllowRevealV6"
	case Instruction_BurnSupplyV6:
		return "BurnSupplyV6"
	case Instruction_BurnFrozen:
		return "BurnFrozen"
	case Instruction_Migrate:
		return "Migrate"
	case Instruction_ThawV5:
		return "ThawV5"
	case Instruction_ThawV6:
		return "ThawV6"
	case Instruction_InitCmV5:
		return "InitCmV5"
	case Instruction_InitCmV6:
		return "InitCmV6"
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
			"mint_v6", (*MintV6)(nil),
		},
		{
			"mint_v5", (*MintV5)(nil),
		},
		{
			"edit_cm_v5", (*EditCmV5)(nil),
		},
		{
			"edit_cm_v6", (*EditCmV6)(nil),
		},
		{
			"reveal_v5", (*RevealV5)(nil),
		},
		{
			"reveal_v6", (*RevealV6)(nil),
		},
		{
			"allow_unfreeze_v5", (*AllowUnfreezeV5)(nil),
		},
		{
			"allow_reveal_v5", (*AllowRevealV5)(nil),
		},
		{
			"burn_supply_v5", (*BurnSupplyV5)(nil),
		},
		{
			"allow_unfreeze_v6", (*AllowUnfreezeV6)(nil),
		},
		{
			"allow_reveal_v6", (*AllowRevealV6)(nil),
		},
		{
			"burn_supply_v6", (*BurnSupplyV6)(nil),
		},
		{
			"burn_frozen", (*BurnFrozen)(nil),
		},
		{
			"migrate", (*Migrate)(nil),
		},
		{
			"thaw_v5", (*ThawV5)(nil),
		},
		{
			"thaw_v6", (*ThawV6)(nil),
		},
		{
			"init_cm_v5", (*InitCmV5)(nil),
		},
		{
			"init_cm_v6", (*InitCmV6)(nil),
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
