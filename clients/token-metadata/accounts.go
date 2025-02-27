// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package token_metadata

import (
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/RoboticAgile/solana-go"
)

type Key ag_binary.BorshEnum

const (
	KeyUninitialized Key = iota
	KeyEditionV1
	KeyMasterEditionV1
	KeyReservationListV1
	KeyMetadataV1
	KeyReservationListV2
	KeyMasterEditionV2
	KeyEditionMarker
	KeyUseAuthorityRecord
	KeyCollectionAuthorityRecord
)

func (value Key) String() string {
	switch value {
	case KeyUninitialized:
		return "Uninitialized"
	case KeyEditionV1:
		return "EditionV1"
	case KeyMasterEditionV1:
		return "MasterEditionV1"
	case KeyReservationListV1:
		return "ReservationListV1"
	case KeyMetadataV1:
		return "MetadataV1"
	case KeyReservationListV2:
		return "ReservationListV2"
	case KeyMasterEditionV2:
		return "MasterEditionV2"
	case KeyEditionMarker:
		return "EditionMarker"
	case KeyUseAuthorityRecord:
		return "UseAuthorityRecord"
	case KeyCollectionAuthorityRecord:
		return "CollectionAuthorityRecord"
	default:
		return ""
	}
}

type Data struct {
	// The name of the asset
	Name string

	// The symbol for the asset
	Symbol string

	// URI pointing to JSON representing the asset
	Uri string

	// Royalty basis points that goes to creators in secondary sales (0-10000)
	SellerFeeBasisPoints uint16

	// Array of creators, optional
	Creators *[]Creator `bin:"optional"`
}

func (obj Data) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Name` param:
	err = encoder.Encode(obj.Name)
	if err != nil {
		return err
	}
	// Serialize `Symbol` param:
	err = encoder.Encode(obj.Symbol)
	if err != nil {
		return err
	}
	// Serialize `Uri` param:
	err = encoder.Encode(obj.Uri)
	if err != nil {
		return err
	}
	// Serialize `SellerFeeBasisPoints` param:
	err = encoder.Encode(obj.SellerFeeBasisPoints)
	if err != nil {
		return err
	}
	// Serialize `Creators` param (optional):
	{
		if obj.Creators == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.Creators)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (obj *Data) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Name`:
	err = decoder.Decode(&obj.Name)
	if err != nil {
		return err
	}
	// Deserialize `Symbol`:
	err = decoder.Decode(&obj.Symbol)
	if err != nil {
		return err
	}
	// Deserialize `Uri`:
	err = decoder.Decode(&obj.Uri)
	if err != nil {
		return err
	}
	// Deserialize `SellerFeeBasisPoints`:
	err = decoder.Decode(&obj.SellerFeeBasisPoints)
	if err != nil {
		return err
	}
	// Deserialize `Creators` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.Creators)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type DataV2 struct {
	// The name of the asset
	Name string

	// The symbol for the asset
	Symbol string

	// URI pointing to JSON representing the asset
	Uri string

	// Royalty basis points that goes to creators in secondary sales (0-10000)
	SellerFeeBasisPoints uint16

	// Array of creators, optional
	Creators *[]Creator `bin:"optional"`

	// Collection
	Collection *Collection `bin:"optional"`

	// Uses
	Uses *Uses `bin:"optional"`
}

func (obj DataV2) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Name` param:
	err = encoder.Encode(obj.Name)
	if err != nil {
		return err
	}
	// Serialize `Symbol` param:
	err = encoder.Encode(obj.Symbol)
	if err != nil {
		return err
	}
	// Serialize `Uri` param:
	err = encoder.Encode(obj.Uri)
	if err != nil {
		return err
	}
	// Serialize `SellerFeeBasisPoints` param:
	err = encoder.Encode(obj.SellerFeeBasisPoints)
	if err != nil {
		return err
	}
	// Serialize `Creators` param (optional):
	{
		if obj.Creators == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.Creators)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `Collection` param (optional):
	{
		if obj.Collection == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.Collection)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `Uses` param (optional):
	{
		if obj.Uses == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.Uses)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (obj *DataV2) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Name`:
	err = decoder.Decode(&obj.Name)
	if err != nil {
		return err
	}
	// Deserialize `Symbol`:
	err = decoder.Decode(&obj.Symbol)
	if err != nil {
		return err
	}
	// Deserialize `Uri`:
	err = decoder.Decode(&obj.Uri)
	if err != nil {
		return err
	}
	// Deserialize `SellerFeeBasisPoints`:
	err = decoder.Decode(&obj.SellerFeeBasisPoints)
	if err != nil {
		return err
	}
	// Deserialize `Creators` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.Creators)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `Collection` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.Collection)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `Uses` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.Uses)
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

	// In percentages, NOT basis points ;) Watch out!
	Share uint8
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

type UseMethod ag_binary.BorshEnum

const (
	UseMethodBurn UseMethod = iota
	UseMethodMultiple
	UseMethodSingle
)

func (value UseMethod) String() string {
	switch value {
	case UseMethodBurn:
		return "Burn"
	case UseMethodMultiple:
		return "Multiple"
	case UseMethodSingle:
		return "Single"
	default:
		return ""
	}
}

type Uses struct {
	UseMethod UseMethod
	Remaining uint64
	Total     uint64
}

func (obj Uses) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `UseMethod` param:
	err = encoder.Encode(obj.UseMethod)
	if err != nil {
		return err
	}
	// Serialize `Remaining` param:
	err = encoder.Encode(obj.Remaining)
	if err != nil {
		return err
	}
	// Serialize `Total` param:
	err = encoder.Encode(obj.Total)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Uses) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `UseMethod`:
	err = decoder.Decode(&obj.UseMethod)
	if err != nil {
		return err
	}
	// Deserialize `Remaining`:
	err = decoder.Decode(&obj.Remaining)
	if err != nil {
		return err
	}
	// Deserialize `Total`:
	err = decoder.Decode(&obj.Total)
	if err != nil {
		return err
	}
	return nil
}

type TokenStandard ag_binary.BorshEnum

const (
	// This is a master edition
	TokenStandardNonFungible TokenStandard = iota

	// A token with metadata that can also have attrributes
	TokenStandardFungibleAsset

	// A token with simple metadata
	TokenStandardFungible

	// This is a limited edition
	TokenStandardNonFungibleEdition
)

func (value TokenStandard) String() string {
	switch value {
	case TokenStandardNonFungible:
		return "NonFungible"
	case TokenStandardFungibleAsset:
		return "FungibleAsset"
	case TokenStandardFungible:
		return "Fungible"
	case TokenStandardNonFungibleEdition:
		return "NonFungibleEdition"
	default:
		return ""
	}
}

type UseAuthorityRecord struct {
	Key         Key
	AllowedUses uint64
}

func (obj UseAuthorityRecord) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Key` param:
	err = encoder.Encode(obj.Key)
	if err != nil {
		return err
	}
	// Serialize `AllowedUses` param:
	err = encoder.Encode(obj.AllowedUses)
	if err != nil {
		return err
	}
	return nil
}

func (obj *UseAuthorityRecord) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Key`:
	err = decoder.Decode(&obj.Key)
	if err != nil {
		return err
	}
	// Deserialize `AllowedUses`:
	err = decoder.Decode(&obj.AllowedUses)
	if err != nil {
		return err
	}
	return nil
}

type CollectionAuthorityRecord struct {
	Key Key
}

func (obj CollectionAuthorityRecord) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Key` param:
	err = encoder.Encode(obj.Key)
	if err != nil {
		return err
	}
	return nil
}

func (obj *CollectionAuthorityRecord) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Key`:
	err = decoder.Decode(&obj.Key)
	if err != nil {
		return err
	}
	return nil
}

type Collection struct {
	Verified bool
	Key      ag_solanago.PublicKey
}

func (obj Collection) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Verified` param:
	err = encoder.Encode(obj.Verified)
	if err != nil {
		return err
	}
	// Serialize `Key` param:
	err = encoder.Encode(obj.Key)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Collection) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Verified`:
	err = decoder.Decode(&obj.Verified)
	if err != nil {
		return err
	}
	// Deserialize `Key`:
	err = decoder.Decode(&obj.Key)
	if err != nil {
		return err
	}
	return nil
}

type Metadata struct {
	Key             Key
	UpdateAuthority ag_solanago.PublicKey
	Mint            ag_solanago.PublicKey
	Data            Data

	// Immutable, once flipped, all sales of this metadata are considered secondary.
	PrimarySaleHappened bool

	// Whether or not the data struct is mutable, default is not
	IsMutable bool

	// nonce for easy calculation of editions, if present
	EditionNonce *uint8 `bin:"optional"`

	// Since we cannot easily change Metadata, we add the new DataV2 fields here at the end.
	TokenStandard *TokenStandard `bin:"optional"`

	// Collection
	Collection *Collection `bin:"optional"`

	// Uses
	Uses *Uses `bin:"optional"`
}

func (obj Metadata) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Key` param:
	err = encoder.Encode(obj.Key)
	if err != nil {
		return err
	}
	// Serialize `UpdateAuthority` param:
	err = encoder.Encode(obj.UpdateAuthority)
	if err != nil {
		return err
	}
	// Serialize `Mint` param:
	err = encoder.Encode(obj.Mint)
	if err != nil {
		return err
	}
	// Serialize `Data` param:
	err = encoder.Encode(obj.Data)
	if err != nil {
		return err
	}
	// Serialize `PrimarySaleHappened` param:
	err = encoder.Encode(obj.PrimarySaleHappened)
	if err != nil {
		return err
	}
	// Serialize `IsMutable` param:
	err = encoder.Encode(obj.IsMutable)
	if err != nil {
		return err
	}
	// Serialize `EditionNonce` param (optional):
	{
		if obj.EditionNonce == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.EditionNonce)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `TokenStandard` param (optional):
	{
		if obj.TokenStandard == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.TokenStandard)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `Collection` param (optional):
	{
		if obj.Collection == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.Collection)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `Uses` param (optional):
	{
		if obj.Uses == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.Uses)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (obj *Metadata) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Key`:
	err = decoder.Decode(&obj.Key)
	if err != nil {
		return err
	}
	// Deserialize `UpdateAuthority`:
	err = decoder.Decode(&obj.UpdateAuthority)
	if err != nil {
		return err
	}
	// Deserialize `Mint`:
	err = decoder.Decode(&obj.Mint)
	if err != nil {
		return err
	}
	// Deserialize `Data`:
	err = decoder.Decode(&obj.Data)
	if err != nil {
		return err
	}
	// Deserialize `PrimarySaleHappened`:
	err = decoder.Decode(&obj.PrimarySaleHappened)
	if err != nil {
		return err
	}
	// Deserialize `IsMutable`:
	err = decoder.Decode(&obj.IsMutable)
	if err != nil {
		return err
	}
	// Deserialize `EditionNonce` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.EditionNonce)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `TokenStandard` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.TokenStandard)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `Collection` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.Collection)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `Uses` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.Uses)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type MasterEditionV2 struct {
	Key       Key
	Supply    uint64
	MaxSupply *uint64 `bin:"optional"`
}

func (obj MasterEditionV2) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Key` param:
	err = encoder.Encode(obj.Key)
	if err != nil {
		return err
	}
	// Serialize `Supply` param:
	err = encoder.Encode(obj.Supply)
	if err != nil {
		return err
	}
	// Serialize `MaxSupply` param (optional):
	{
		if obj.MaxSupply == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.MaxSupply)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (obj *MasterEditionV2) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Key`:
	err = decoder.Decode(&obj.Key)
	if err != nil {
		return err
	}
	// Deserialize `Supply`:
	err = decoder.Decode(&obj.Supply)
	if err != nil {
		return err
	}
	// Deserialize `MaxSupply` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.MaxSupply)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type MasterEditionV1 struct {
	Key       Key
	Supply    uint64
	MaxSupply *uint64 `bin:"optional"`

	// Can be used to mint tokens that give one-time permission to mint a single limited edition.
	PrintingMint ag_solanago.PublicKey

	// If you don't know how many printing tokens you are going to need, but you do know
	// you are going to need some amount in the future, you can use a token from this mint.
	// Coming back to token metadata with one of these tokens allows you to mint (one time)
	// any number of printing tokens you want. This is used for instance by Auction Manager
	// with participation NFTs, where we dont know how many people will bid and need participation
	// printing tokens to redeem, so we give it ONE of these tokens to use after the auction is over,
	// because when the auction begins we just dont know how many printing tokens we will need,
	// but at the end we will. At the end it then burns this token with token-metadata to
	// get the printing tokens it needs to give to bidders. Each bidder then redeems a printing token
	// to get their limited editions.
	OneTimePrintingAuthorizationMint ag_solanago.PublicKey
}

func (obj MasterEditionV1) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Key` param:
	err = encoder.Encode(obj.Key)
	if err != nil {
		return err
	}
	// Serialize `Supply` param:
	err = encoder.Encode(obj.Supply)
	if err != nil {
		return err
	}
	// Serialize `MaxSupply` param (optional):
	{
		if obj.MaxSupply == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.MaxSupply)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `PrintingMint` param:
	err = encoder.Encode(obj.PrintingMint)
	if err != nil {
		return err
	}
	// Serialize `OneTimePrintingAuthorizationMint` param:
	err = encoder.Encode(obj.OneTimePrintingAuthorizationMint)
	if err != nil {
		return err
	}
	return nil
}

func (obj *MasterEditionV1) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Key`:
	err = decoder.Decode(&obj.Key)
	if err != nil {
		return err
	}
	// Deserialize `Supply`:
	err = decoder.Decode(&obj.Supply)
	if err != nil {
		return err
	}
	// Deserialize `MaxSupply` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.MaxSupply)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `PrintingMint`:
	err = decoder.Decode(&obj.PrintingMint)
	if err != nil {
		return err
	}
	// Deserialize `OneTimePrintingAuthorizationMint`:
	err = decoder.Decode(&obj.OneTimePrintingAuthorizationMint)
	if err != nil {
		return err
	}
	return nil
}

type Edition struct {
	Key Key

	// Points at MasterEdition struct
	Parent ag_solanago.PublicKey

	// Starting at 0 for master record, this is incremented for each edition minted.
	Edition uint64
}

func (obj Edition) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Key` param:
	err = encoder.Encode(obj.Key)
	if err != nil {
		return err
	}
	// Serialize `Parent` param:
	err = encoder.Encode(obj.Parent)
	if err != nil {
		return err
	}
	// Serialize `Edition` param:
	err = encoder.Encode(obj.Edition)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Edition) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Key`:
	err = decoder.Decode(&obj.Key)
	if err != nil {
		return err
	}
	// Deserialize `Parent`:
	err = decoder.Decode(&obj.Parent)
	if err != nil {
		return err
	}
	// Deserialize `Edition`:
	err = decoder.Decode(&obj.Edition)
	if err != nil {
		return err
	}
	return nil
}

type ReservationListV2 struct {
	Key Key

	// Present for reverse lookups
	MasterEdition ag_solanago.PublicKey

	// What supply counter was on master_edition when this reservation was created.
	SupplySnapshot *uint64 `bin:"optional"`
	Reservations   []Reservation

	// How many reservations there are going to be, given on first set_reservation call
	TotalReservationSpots uint64

	// Cached count of reservation spots in the reservation vec to save on CPU.
	CurrentReservationSpots uint64
}

func (obj ReservationListV2) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Key` param:
	err = encoder.Encode(obj.Key)
	if err != nil {
		return err
	}
	// Serialize `MasterEdition` param:
	err = encoder.Encode(obj.MasterEdition)
	if err != nil {
		return err
	}
	// Serialize `SupplySnapshot` param (optional):
	{
		if obj.SupplySnapshot == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.SupplySnapshot)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `Reservations` param:
	err = encoder.Encode(obj.Reservations)
	if err != nil {
		return err
	}
	// Serialize `TotalReservationSpots` param:
	err = encoder.Encode(obj.TotalReservationSpots)
	if err != nil {
		return err
	}
	// Serialize `CurrentReservationSpots` param:
	err = encoder.Encode(obj.CurrentReservationSpots)
	if err != nil {
		return err
	}
	return nil
}

func (obj *ReservationListV2) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Key`:
	err = decoder.Decode(&obj.Key)
	if err != nil {
		return err
	}
	// Deserialize `MasterEdition`:
	err = decoder.Decode(&obj.MasterEdition)
	if err != nil {
		return err
	}
	// Deserialize `SupplySnapshot` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.SupplySnapshot)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `Reservations`:
	err = decoder.Decode(&obj.Reservations)
	if err != nil {
		return err
	}
	// Deserialize `TotalReservationSpots`:
	err = decoder.Decode(&obj.TotalReservationSpots)
	if err != nil {
		return err
	}
	// Deserialize `CurrentReservationSpots`:
	err = decoder.Decode(&obj.CurrentReservationSpots)
	if err != nil {
		return err
	}
	return nil
}

type Reservation struct {
	Address        ag_solanago.PublicKey
	SpotsRemaining uint64
	TotalSpots     uint64
}

func (obj Reservation) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Address` param:
	err = encoder.Encode(obj.Address)
	if err != nil {
		return err
	}
	// Serialize `SpotsRemaining` param:
	err = encoder.Encode(obj.SpotsRemaining)
	if err != nil {
		return err
	}
	// Serialize `TotalSpots` param:
	err = encoder.Encode(obj.TotalSpots)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Reservation) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Address`:
	err = decoder.Decode(&obj.Address)
	if err != nil {
		return err
	}
	// Deserialize `SpotsRemaining`:
	err = decoder.Decode(&obj.SpotsRemaining)
	if err != nil {
		return err
	}
	// Deserialize `TotalSpots`:
	err = decoder.Decode(&obj.TotalSpots)
	if err != nil {
		return err
	}
	return nil
}

type ReservationListV1 struct {
	Key Key

	// Present for reverse lookups
	MasterEdition ag_solanago.PublicKey

	// What supply counter was on master_edition when this reservation was created.
	SupplySnapshot *uint64 `bin:"optional"`
	Reservations   []ReservationV1
}

func (obj ReservationListV1) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Key` param:
	err = encoder.Encode(obj.Key)
	if err != nil {
		return err
	}
	// Serialize `MasterEdition` param:
	err = encoder.Encode(obj.MasterEdition)
	if err != nil {
		return err
	}
	// Serialize `SupplySnapshot` param (optional):
	{
		if obj.SupplySnapshot == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.SupplySnapshot)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `Reservations` param:
	err = encoder.Encode(obj.Reservations)
	if err != nil {
		return err
	}
	return nil
}

func (obj *ReservationListV1) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Key`:
	err = decoder.Decode(&obj.Key)
	if err != nil {
		return err
	}
	// Deserialize `MasterEdition`:
	err = decoder.Decode(&obj.MasterEdition)
	if err != nil {
		return err
	}
	// Deserialize `SupplySnapshot` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.SupplySnapshot)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `Reservations`:
	err = decoder.Decode(&obj.Reservations)
	if err != nil {
		return err
	}
	return nil
}

type ReservationV1 struct {
	Address        ag_solanago.PublicKey
	SpotsRemaining uint8
	TotalSpots     uint8
}

func (obj ReservationV1) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Address` param:
	err = encoder.Encode(obj.Address)
	if err != nil {
		return err
	}
	// Serialize `SpotsRemaining` param:
	err = encoder.Encode(obj.SpotsRemaining)
	if err != nil {
		return err
	}
	// Serialize `TotalSpots` param:
	err = encoder.Encode(obj.TotalSpots)
	if err != nil {
		return err
	}
	return nil
}

func (obj *ReservationV1) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Address`:
	err = decoder.Decode(&obj.Address)
	if err != nil {
		return err
	}
	// Deserialize `SpotsRemaining`:
	err = decoder.Decode(&obj.SpotsRemaining)
	if err != nil {
		return err
	}
	// Deserialize `TotalSpots`:
	err = decoder.Decode(&obj.TotalSpots)
	if err != nil {
		return err
	}
	return nil
}

type EditionMarker struct {
	Key    Key
	Ledger [31]uint8
}

func (obj EditionMarker) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Key` param:
	err = encoder.Encode(obj.Key)
	if err != nil {
		return err
	}
	// Serialize `Ledger` param:
	err = encoder.Encode(obj.Ledger)
	if err != nil {
		return err
	}
	return nil
}

func (obj *EditionMarker) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Key`:
	err = decoder.Decode(&obj.Key)
	if err != nil {
		return err
	}
	// Deserialize `Ledger`:
	err = decoder.Decode(&obj.Ledger)
	if err != nil {
		return err
	}
	return nil
}
