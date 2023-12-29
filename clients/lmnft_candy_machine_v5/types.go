// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package nft_candy_machine

import (
	"fmt"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/desperatee/solana-go"
)

type PublicMode struct {
	MintsPerUser *uint32 `bin:"optional"`
}

func (obj PublicMode) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `MintsPerUser` param (optional):
	{
		if obj.MintsPerUser == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.MintsPerUser)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (obj *PublicMode) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `MintsPerUser` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.MintsPerUser)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type WhitelistTokenMode struct {
	TokenMint ag_solanago.PublicKey
	Burn      bool
}

func (obj WhitelistTokenMode) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `TokenMint` param:
	err = encoder.Encode(obj.TokenMint)
	if err != nil {
		return err
	}
	// Serialize `Burn` param:
	err = encoder.Encode(obj.Burn)
	if err != nil {
		return err
	}
	return nil
}

func (obj *WhitelistTokenMode) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `TokenMint`:
	err = decoder.Decode(&obj.TokenMint)
	if err != nil {
		return err
	}
	// Deserialize `Burn`:
	err = decoder.Decode(&obj.Burn)
	if err != nil {
		return err
	}
	return nil
}

type NFTHolderMode struct {
	VerifiedCreator ag_solanago.PublicKey
	MintsPerNft     uint32
}

func (obj NFTHolderMode) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `VerifiedCreator` param:
	err = encoder.Encode(obj.VerifiedCreator)
	if err != nil {
		return err
	}
	// Serialize `MintsPerNft` param:
	err = encoder.Encode(obj.MintsPerNft)
	if err != nil {
		return err
	}
	return nil
}

func (obj *NFTHolderMode) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `VerifiedCreator`:
	err = decoder.Decode(&obj.VerifiedCreator)
	if err != nil {
		return err
	}
	// Deserialize `MintsPerNft`:
	err = decoder.Decode(&obj.MintsPerNft)
	if err != nil {
		return err
	}
	return nil
}

type WhitelistWalletListMode struct {
	MintsPerUser uint32
	MerkleRoot   [32]uint8
}

func (obj WhitelistWalletListMode) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `MintsPerUser` param:
	err = encoder.Encode(obj.MintsPerUser)
	if err != nil {
		return err
	}
	// Serialize `MerkleRoot` param:
	err = encoder.Encode(obj.MerkleRoot)
	if err != nil {
		return err
	}
	return nil
}

func (obj *WhitelistWalletListMode) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `MintsPerUser`:
	err = decoder.Decode(&obj.MintsPerUser)
	if err != nil {
		return err
	}
	// Deserialize `MerkleRoot`:
	err = decoder.Decode(&obj.MerkleRoot)
	if err != nil {
		return err
	}
	return nil
}

type SaleFaze struct {
	Start         int64
	Price         uint64
	Currency      *ag_solanago.PublicKey `bin:"optional"`
	WhitelistMode WhitelistMode
	Name          string
}

func (obj SaleFaze) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Start` param:
	err = encoder.Encode(obj.Start)
	if err != nil {
		return err
	}
	// Serialize `Price` param:
	err = encoder.Encode(obj.Price)
	if err != nil {
		return err
	}
	// Serialize `Currency` param (optional):
	{
		if obj.Currency == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.Currency)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `WhitelistMode` param:
	{
		tmp := whitelistModeContainer{}
		switch realvalue := obj.WhitelistMode.(type) {
		case *WhitelistModeWalletBased:
			tmp.Enum = 0
			tmp.WalletBased = *realvalue
		case *WhitelistModeTokenBased:
			tmp.Enum = 1
			tmp.TokenBased = *realvalue
		case *WhitelistModePublic:
			tmp.Enum = 2
			tmp.Public = *realvalue
		case *WhitelistModeNFTBased:
			tmp.Enum = 3
			tmp.NFTBased = *realvalue
		}
		err := encoder.Encode(tmp)
		if err != nil {
			return err
		}
	}
	// Serialize `Name` param:
	err = encoder.Encode(obj.Name)
	if err != nil {
		return err
	}
	return nil
}

func (obj *SaleFaze) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Start`:
	err = decoder.Decode(&obj.Start)
	if err != nil {
		return err
	}
	// Deserialize `Price`:
	err = decoder.Decode(&obj.Price)
	if err != nil {
		return err
	}
	// Deserialize `Currency` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.Currency)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `WhitelistMode`:
	{
		tmp := new(whitelistModeContainer)
		err := decoder.Decode(tmp)
		if err != nil {
			return err
		}
		switch tmp.Enum {
		case 0:
			obj.WhitelistMode = &tmp.WalletBased
		case 1:
			obj.WhitelistMode = &tmp.TokenBased
		case 2:
			obj.WhitelistMode = &tmp.Public
		case 3:
			obj.WhitelistMode = &tmp.NFTBased
		default:
			return fmt.Errorf("unknown enum index: %v", tmp.Enum)
		}
	}
	// Deserialize `Name`:
	err = decoder.Decode(&obj.Name)
	if err != nil {
		return err
	}
	return nil
}

type CandyMachineDataV3 struct {
	ItemsAvailable       uint64
	GoLiveDate           int64
	Symbol               string
	SellerFeeBasisPoints uint16
	Creators             []Creator
	IsMutable            bool
	RetainAuthority      bool
	BaseUrl              string
	SaleFazes            []SaleFaze
}

func (obj CandyMachineDataV3) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `ItemsAvailable` param:
	err = encoder.Encode(obj.ItemsAvailable)
	if err != nil {
		return err
	}
	// Serialize `GoLiveDate` param:
	err = encoder.Encode(obj.GoLiveDate)
	if err != nil {
		return err
	}
	// Serialize `Symbol` param:
	err = encoder.Encode(obj.Symbol)
	if err != nil {
		return err
	}
	// Serialize `SellerFeeBasisPoints` param:
	err = encoder.Encode(obj.SellerFeeBasisPoints)
	if err != nil {
		return err
	}
	// Serialize `Creators` param:
	err = encoder.Encode(obj.Creators)
	if err != nil {
		return err
	}
	// Serialize `IsMutable` param:
	err = encoder.Encode(obj.IsMutable)
	if err != nil {
		return err
	}
	// Serialize `RetainAuthority` param:
	err = encoder.Encode(obj.RetainAuthority)
	if err != nil {
		return err
	}
	// Serialize `BaseUrl` param:
	err = encoder.Encode(obj.BaseUrl)
	if err != nil {
		return err
	}
	// Serialize `SaleFazes` param:
	err = encoder.Encode(obj.SaleFazes)
	if err != nil {
		return err
	}
	return nil
}

func (obj *CandyMachineDataV3) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `ItemsAvailable`:
	err = decoder.Decode(&obj.ItemsAvailable)
	if err != nil {
		return err
	}
	// Deserialize `GoLiveDate`:
	err = decoder.Decode(&obj.GoLiveDate)
	if err != nil {
		return err
	}
	// Deserialize `Symbol`:
	err = decoder.Decode(&obj.Symbol)
	if err != nil {
		return err
	}
	// Deserialize `SellerFeeBasisPoints`:
	err = decoder.Decode(&obj.SellerFeeBasisPoints)
	if err != nil {
		return err
	}
	// Deserialize `Creators`:
	err = decoder.Decode(&obj.Creators)
	if err != nil {
		return err
	}
	// Deserialize `IsMutable`:
	err = decoder.Decode(&obj.IsMutable)
	if err != nil {
		return err
	}
	// Deserialize `RetainAuthority`:
	err = decoder.Decode(&obj.RetainAuthority)
	if err != nil {
		return err
	}
	// Deserialize `BaseUrl`:
	err = decoder.Decode(&obj.BaseUrl)
	if err != nil {
		return err
	}
	// Deserialize `SaleFazes`:
	err = decoder.Decode(&obj.SaleFazes)
	if err != nil {
		return err
	}
	return nil
}

type CandyMachineDataV2 struct {
	Price                uint64
	ItemsAvailable       uint64
	GoLiveDate           int64
	Symbol               string
	SellerFeeBasisPoints uint16
	Creators             []Creator
	IsMutable            bool
	RetainAuthority      bool
	BaseUrl              string
	MintsPerUser         *uint32      `bin:"optional"`
	Whitelist            *WhitelistV2 `bin:"optional"`
}

func (obj CandyMachineDataV2) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Price` param:
	err = encoder.Encode(obj.Price)
	if err != nil {
		return err
	}
	// Serialize `ItemsAvailable` param:
	err = encoder.Encode(obj.ItemsAvailable)
	if err != nil {
		return err
	}
	// Serialize `GoLiveDate` param:
	err = encoder.Encode(obj.GoLiveDate)
	if err != nil {
		return err
	}
	// Serialize `Symbol` param:
	err = encoder.Encode(obj.Symbol)
	if err != nil {
		return err
	}
	// Serialize `SellerFeeBasisPoints` param:
	err = encoder.Encode(obj.SellerFeeBasisPoints)
	if err != nil {
		return err
	}
	// Serialize `Creators` param:
	err = encoder.Encode(obj.Creators)
	if err != nil {
		return err
	}
	// Serialize `IsMutable` param:
	err = encoder.Encode(obj.IsMutable)
	if err != nil {
		return err
	}
	// Serialize `RetainAuthority` param:
	err = encoder.Encode(obj.RetainAuthority)
	if err != nil {
		return err
	}
	// Serialize `BaseUrl` param:
	err = encoder.Encode(obj.BaseUrl)
	if err != nil {
		return err
	}
	// Serialize `MintsPerUser` param (optional):
	{
		if obj.MintsPerUser == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.MintsPerUser)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `Whitelist` param (optional):
	{
		if obj.Whitelist == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.Whitelist)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (obj *CandyMachineDataV2) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Price`:
	err = decoder.Decode(&obj.Price)
	if err != nil {
		return err
	}
	// Deserialize `ItemsAvailable`:
	err = decoder.Decode(&obj.ItemsAvailable)
	if err != nil {
		return err
	}
	// Deserialize `GoLiveDate`:
	err = decoder.Decode(&obj.GoLiveDate)
	if err != nil {
		return err
	}
	// Deserialize `Symbol`:
	err = decoder.Decode(&obj.Symbol)
	if err != nil {
		return err
	}
	// Deserialize `SellerFeeBasisPoints`:
	err = decoder.Decode(&obj.SellerFeeBasisPoints)
	if err != nil {
		return err
	}
	// Deserialize `Creators`:
	err = decoder.Decode(&obj.Creators)
	if err != nil {
		return err
	}
	// Deserialize `IsMutable`:
	err = decoder.Decode(&obj.IsMutable)
	if err != nil {
		return err
	}
	// Deserialize `RetainAuthority`:
	err = decoder.Decode(&obj.RetainAuthority)
	if err != nil {
		return err
	}
	// Deserialize `BaseUrl`:
	err = decoder.Decode(&obj.BaseUrl)
	if err != nil {
		return err
	}
	// Deserialize `MintsPerUser` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.MintsPerUser)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `Whitelist` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.Whitelist)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type WhitelistV2 struct {
	MerkleRoot   [32]uint8
	MintsPerUser *uint32 `bin:"optional"`
	GoLiveDate   int64
	Price        *uint64 `bin:"optional"`
}

func (obj WhitelistV2) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `MerkleRoot` param:
	err = encoder.Encode(obj.MerkleRoot)
	if err != nil {
		return err
	}
	// Serialize `MintsPerUser` param (optional):
	{
		if obj.MintsPerUser == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.MintsPerUser)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `GoLiveDate` param:
	err = encoder.Encode(obj.GoLiveDate)
	if err != nil {
		return err
	}
	// Serialize `Price` param (optional):
	{
		if obj.Price == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.Price)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (obj *WhitelistV2) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `MerkleRoot`:
	err = decoder.Decode(&obj.MerkleRoot)
	if err != nil {
		return err
	}
	// Deserialize `MintsPerUser` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.MintsPerUser)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `GoLiveDate`:
	err = decoder.Decode(&obj.GoLiveDate)
	if err != nil {
		return err
	}
	// Deserialize `Price` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.Price)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type CandyMachineData struct {
	Uuid           string
	Price          uint64
	ItemsAvailable uint64
	GoLiveDate     *int64 `bin:"optional"`
}

func (obj CandyMachineData) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Uuid` param:
	err = encoder.Encode(obj.Uuid)
	if err != nil {
		return err
	}
	// Serialize `Price` param:
	err = encoder.Encode(obj.Price)
	if err != nil {
		return err
	}
	// Serialize `ItemsAvailable` param:
	err = encoder.Encode(obj.ItemsAvailable)
	if err != nil {
		return err
	}
	// Serialize `GoLiveDate` param (optional):
	{
		if obj.GoLiveDate == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.GoLiveDate)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (obj *CandyMachineData) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Uuid`:
	err = decoder.Decode(&obj.Uuid)
	if err != nil {
		return err
	}
	// Deserialize `Price`:
	err = decoder.Decode(&obj.Price)
	if err != nil {
		return err
	}
	// Deserialize `ItemsAvailable`:
	err = decoder.Decode(&obj.ItemsAvailable)
	if err != nil {
		return err
	}
	// Deserialize `GoLiveDate` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.GoLiveDate)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type Creator struct {
	Address  ag_solanago.PublicKey
	Verified bool
	Share    uint8
}

func (obj Creator) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Address` param:
	err = encoder.Encode(obj.Address)
	if err != nil {
		return err
	}
	// Serialize `Verified` param:
	err = encoder.Encode(obj.Verified)
	if err != nil {
		return err
	}
	// Serialize `Share` param:
	err = encoder.Encode(obj.Share)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Creator) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Address`:
	err = decoder.Decode(&obj.Address)
	if err != nil {
		return err
	}
	// Deserialize `Verified`:
	err = decoder.Decode(&obj.Verified)
	if err != nil {
		return err
	}
	// Deserialize `Share`:
	err = decoder.Decode(&obj.Share)
	if err != nil {
		return err
	}
	return nil
}

type WhitelistMode interface {
	isWhitelistMode()
}

type whitelistModeContainer struct {
	Enum        ag_binary.BorshEnum `borsh_enum:"true"`
	WalletBased WhitelistModeWalletBased
	TokenBased  WhitelistModeTokenBased
	Public      WhitelistModePublic
	NFTBased    WhitelistModeNFTBased
}

type WhitelistModeWalletBased struct {
	Info WhitelistWalletListMode
}

func (obj WhitelistModeWalletBased) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Info` param:
	err = encoder.Encode(obj.Info)
	if err != nil {
		return err
	}
	return nil
}

func (obj *WhitelistModeWalletBased) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Info`:
	err = decoder.Decode(&obj.Info)
	if err != nil {
		return err
	}
	return nil
}

func (_ *WhitelistModeWalletBased) isWhitelistMode() {}

type WhitelistModeTokenBased struct {
	Info WhitelistTokenMode
}

func (obj WhitelistModeTokenBased) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Info` param:
	err = encoder.Encode(obj.Info)
	if err != nil {
		return err
	}
	return nil
}

func (obj *WhitelistModeTokenBased) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Info`:
	err = decoder.Decode(&obj.Info)
	if err != nil {
		return err
	}
	return nil
}

func (_ *WhitelistModeTokenBased) isWhitelistMode() {}

type WhitelistModePublic struct {
	Info PublicMode
}

func (obj WhitelistModePublic) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Info` param:
	err = encoder.Encode(obj.Info)
	if err != nil {
		return err
	}
	return nil
}

func (obj *WhitelistModePublic) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Info`:
	err = decoder.Decode(&obj.Info)
	if err != nil {
		return err
	}
	return nil
}

func (_ *WhitelistModePublic) isWhitelistMode() {}

type WhitelistModeNFTBased struct {
	Info NFTHolderMode
}

func (obj WhitelistModeNFTBased) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Info` param:
	err = encoder.Encode(obj.Info)
	if err != nil {
		return err
	}
	return nil
}

func (obj *WhitelistModeNFTBased) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Info`:
	err = decoder.Decode(&obj.Info)
	if err != nil {
		return err
	}
	return nil
}

func (_ *WhitelistModeNFTBased) isWhitelistMode() {}

type ErrorCode ag_binary.BorshEnum

const (
	ErrorCodeIncorrectOwner ErrorCode = iota
	ErrorCodeUninitialized
	ErrorCodeMintMismatch
	ErrorCodeIndexGreaterThanLength
	ErrorCodeConfigMustHaveAtleastOneEntry
	ErrorCodeNumericalOverflowError
	ErrorCodeTooManyCreators
	ErrorCodeUuidMustBeExactly6Length
	ErrorCodeNotEnoughTokens
	ErrorCodeNotEnoughSOL
	ErrorCodeTokenTransferFailed
	ErrorCodeCandyMachineEmpty
	ErrorCodeCandyMachineNotLiveYet
	ErrorCodeConfigLineMismatch
	ErrorCodeWhitelistExists
	ErrorCodeWhiteListMissing
	ErrorCodeWrongWhitelist
	ErrorCodeNotWhitelisted
	ErrorCodeInvalidFraction
	ErrorCodeBumpMissing
	ErrorCodePriceViolation
	ErrorCodeNotThawable
	ErrorCodeTeePot
	ErrorCodePresentProof
)

func (value ErrorCode) String() string {
	switch value {
	case ErrorCodeIncorrectOwner:
		return "IncorrectOwner"
	case ErrorCodeUninitialized:
		return "Uninitialized"
	case ErrorCodeMintMismatch:
		return "MintMismatch"
	case ErrorCodeIndexGreaterThanLength:
		return "IndexGreaterThanLength"
	case ErrorCodeConfigMustHaveAtleastOneEntry:
		return "ConfigMustHaveAtleastOneEntry"
	case ErrorCodeNumericalOverflowError:
		return "NumericalOverflowError"
	case ErrorCodeTooManyCreators:
		return "TooManyCreators"
	case ErrorCodeUuidMustBeExactly6Length:
		return "UuidMustBeExactly6Length"
	case ErrorCodeNotEnoughTokens:
		return "NotEnoughTokens"
	case ErrorCodeNotEnoughSOL:
		return "NotEnoughSOL"
	case ErrorCodeTokenTransferFailed:
		return "TokenTransferFailed"
	case ErrorCodeCandyMachineEmpty:
		return "CandyMachineEmpty"
	case ErrorCodeCandyMachineNotLiveYet:
		return "CandyMachineNotLiveYet"
	case ErrorCodeConfigLineMismatch:
		return "ConfigLineMismatch"
	case ErrorCodeWhitelistExists:
		return "WhitelistExists"
	case ErrorCodeWhiteListMissing:
		return "WhiteListMissing"
	case ErrorCodeWrongWhitelist:
		return "WrongWhitelist"
	case ErrorCodeNotWhitelisted:
		return "NotWhitelisted"
	case ErrorCodeInvalidFraction:
		return "InvalidFraction"
	case ErrorCodeBumpMissing:
		return "BumpMissing"
	case ErrorCodePriceViolation:
		return "PriceViolation"
	case ErrorCodeNotThawable:
		return "NotThawable"
	case ErrorCodeTeePot:
		return "TeePot"
	case ErrorCodePresentProof:
		return "PresentProof"
	default:
		return ""
	}
}
