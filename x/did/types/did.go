package types

import (
	fmt "fmt"
	"log"
	"regexp"
	"strings"

	"github.com/btcsuite/btcutil/base58"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gogo/protobuf/proto"
	"github.com/tendermint/tendermint/libs/bytes"

	sdkcodec "github.com/cosmos/cosmos-sdk/codec"
	"github.com/tendermint/tendermint/crypto"

	"soarchain/x/did/errors"
)

const (
	DIDMethod     = "soar"
	Base58Charset = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
)

const (
	ContextDidV1 = "https://www.w3.org/ns/did/v1"
)

const (
	MaxVerificationMethodIDLen = 128
)

const (
	JSONWEBKEY_2020 = "JsonWebKey2020"
	ES256K_2019     = "EcdsaSecp256k1VerificationKey2019"
	ES256K_2018     = "Secp256k1VerificationKey2018"
	ED25519_2018    = "Ed25519VerificationKey2018"
	BLS1281G1_2020  = "Bls12381G1Key2020"
	BLS1281G2_2020  = "Bls12381G2Key2020"
	GPG_2020        = "GpgVerificationKey2020"
	RSA_2018        = "RsaVerificationKey2018"
	X25519_2019     = "X25519KeyAgreementKey2019"
	SS256K_2019     = "SchnorrSecp256k1VerificationKey2019"
	ES256K_R_2020   = "EcdsaSecp256k1RecoveryMethod2020"
)

func ValidateKeyType(keyType string) bool {
	switch keyType {
	case JSONWEBKEY_2020,
		ES256K_2019,
		ES256K_2018,
		ED25519_2018,
		BLS1281G1_2020,
		BLS1281G2_2020,
		GPG_2020,
		RSA_2018,
		X25519_2019,
		SS256K_2019,
		ES256K_R_2020:
		return true
	}

	if keyType == "" {
		return false
	}
	log.Printf("[warn] unknown key type: %s\n", keyType)
	return true
}

type JSONStringOrStrings []string

func ParseDid(str string) (string, error) {
	did := str
	if !ValidateDid(did) {
		return "", sdkerrors.Wrap(sdkerrors.ErrNotFound, "[ParseDid] failed. Make sure you are using a valid did format.")
	}
	return did, nil
}

func ValidateDid(did string) bool {
	pattern := fmt.Sprintf("^%s$", didRegex())
	matched, err := regexp.MatchString(pattern, did)
	if err != nil {
		return false
	}
	return matched
}

func Sign(signableData sdkcodec.ProtoMarshaler, seq uint64, privKey crypto.PrivKey) ([]byte, error) {
	signBytes, err := mustGetSignBytesWithSeq(signableData, seq)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[Sign] failed. Make sure you are using a valid signable data.")
	}
	return privKey.Sign(signBytes)
}

func (strings JSONStringOrStrings) protoType() *Strings {
	values := make([]string, 0, len(strings))

	for _, s := range strings {
		values = append(values, s)
	}

	return &Strings{values}
}

func (strings JSONStringOrStrings) Marshal() ([]byte, error) {
	return proto.Marshal(strings.protoType())
}

func (strings *JSONStringOrStrings) MarshalTo(data []byte) (n int, err error) {
	return strings.protoType().MarshalTo(data)
}

func (strings *JSONStringOrStrings) Unmarshal(data []byte) error {
	protoType := &Strings{}
	if err := proto.Unmarshal(data, protoType); err != nil {
		return err
	}

	*strings = protoType.Values
	return nil
}

func (strings JSONStringOrStrings) Size() int {
	return strings.protoType().Size()
}

func ParseDidDocument(str string) (string, error) {
	did := str
	if !ValidateDid(did) {
		return "", sdkerrors.Wrap(sdkerrors.ErrNotFound, "[ParseDidDocument] failed. Make sure you are using a valid did format.")
	}
	return did, nil
}

func didRegex() string {
	return fmt.Sprintf("did:%s:[%s]{32,44}", DIDMethod, Base58Charset)
}

func EmptyDid(did string) bool {
	return did == ""
}

func (doc ClientDid) Empty() bool {
	return EmptyDid(doc.Id)
}

func EmptyDids(strings []string) bool {
	if len(strings) == 0 {
		return true
	}

	for _, did := range strings {
		if !EmptyDid(did) {
			return false
		}
	}

	return true
}

func ValidateDIDs(strings []string) bool {
	if EmptyDids(strings) {
		return false
	}

	for _, did := range strings {
		if !ValidateDid(did) {
			return false
		}
	}

	return true
}

func ValidateContexts(contexts []string) bool {
	if len(contexts) == 0 || contexts[0] != ContextDidV1 {
		return false
	}

	set := make(map[string]struct{}, len(contexts))
	for _, context := range contexts {
		_, dup := set[context]
		if dup || !ValidateContext(context) {
			return false
		}
		set[context] = struct{}{}
	}
	return true
}

func ValidateContext(context string) bool {
	return context != ""
}

func (v VerificationRelationship) hasDedicatedMethod() bool {
	return v.GetVerificationMethod() != nil
}

func ValidateVerificationMethodID(verificationMethodID string, did string) bool {
	prefix := fmt.Sprintf("%v#", did)
	if !strings.HasPrefix(verificationMethodID, prefix) {
		return false
	}

	if len(verificationMethodID)-len(prefix) > MaxVerificationMethodIDLen {
		return false
	}

	suffix := verificationMethodID[len(prefix):]
	matched, err := regexp.MatchString(`^\S+$`, suffix)
	if err != nil {
		return false
	}
	return matched
}

func (pk VerificationMethod) Valid(did string) bool {
	if !ValidateVerificationMethodID(pk.Id, did) || !ValidateKeyType(pk.Type) {
		return false
	}

	pattern := fmt.Sprintf("^[%s]+$", Base58Charset)
	matched, err := regexp.MatchString(pattern, pk.PublicKeyBase58)
	if err != nil {
		return false
	}
	return matched
}

func (v VerificationRelationship) Valid(did string) bool {
	if v.hasDedicatedMethod() {
		return v.GetVerificationMethod().Valid(did)
	} else {
		return ValidateVerificationMethodID(v.GetVerificationMethodId(), did)
	}
}

func (doc ClientDid) validVerificationRelationships(relationships []VerificationRelationship) bool {
	for _, relationship := range relationships {
		if !relationship.Valid(doc.Id) {
			return false
		}
		if !relationship.hasDedicatedMethod() {
			if _, ok := doc.VerificationMethodByID(relationship.GetVerificationMethodId()); !ok {
				return false
			}
		}
	}
	return true
}

func (s Service) Valid() bool {
	return s.Id != "" && s.Type != "" && s.ServiceEndpoint != ""
}

func (doc ClientDid) VerificationMethodByID(id string) (VerificationMethod, bool) {
	for _, verificationMethod := range doc.VerificationMethods {
		if verificationMethod.Id == id {
			return *verificationMethod, true
		}
	}
	return VerificationMethod{}, false
}

func (doc ClientDid) Valid() bool {
	if doc.Empty() {
		return true
	}

	if !ValidateDid(doc.Id) || doc.VerificationMethods == nil || doc.Authentications == nil {
		return false
	}

	for _, verificationMethod := range doc.VerificationMethods {
		if !verificationMethod.Valid(doc.Id) {
			return false
		}
	}

	if !doc.validVerificationRelationships(doc.Authentications) {
		return false
	}

	for _, service := range doc.Services {
		if !service.Valid() {
			return false
		}
	}

	return true
}

func (d ClientDidWithSeq) Valid() bool {
	return d.Document.Valid()
}

func NewVerificationMethodID(did string, name string) string {
	return fmt.Sprintf("%v#%s", did, name)
}

func NewVerificationMethod(id string, keyType string, controller string, publicKeyBase58 []byte) VerificationMethod {
	return VerificationMethod{
		Id:              id,
		Type:            keyType,
		Controller:      controller,
		PublicKeyBase58: base58.Encode(publicKeyBase58),
	}
}

func NewVerificationRelationship(verificationMethodID string) VerificationRelationship {
	return VerificationRelationship{
		Content: &VerificationRelationship_VerificationMethodId{VerificationMethodId: verificationMethodID},
	}
}

type ClientDidDocumentOption func(opts *ClientDid)

func NewClientDidDocument(id string, index string, address string, soarchainType string, pids []bool, opts ...ClientDidDocumentOption) ClientDid {
	doc := ClientDid{
		Id:      id,
		PubKey:  index,
		Address: address,
		Type:    soarchainType,
	}

	for _, opt := range opts {
		opt(&doc)
	}
	return doc
}

func NewKeys(id string, pubkeyType string, controller string, publicKeyPem string) Keys {
	return Keys{
		Id:           id,
		PubkeyType:   pubkeyType,
		Controller:   controller,
		PublicKeyPem: publicKeyPem,
	}
}

func NewVehicle(vin string) Vehicle {
	return Vehicle{
		Vin: vin,
	}
}

func NewOwner(id string, purchaseDate string) Owner {
	return Owner{
		Id:           id,
		PurchaseDate: purchaseDate,
	}
}

func WithKeys(soarchainPublicKey *Keys) ClientDidDocumentOption {
	return func(opts *ClientDid) {
		opts.Keys = soarchainPublicKey
	}
}

func WithVehicle(vehicle *Vehicle) ClientDidDocumentOption {
	return func(opts *ClientDid) {
		opts.Vehicle = vehicle
	}
}

func WithOwner(owner *Owner) ClientDidDocumentOption {
	return func(opts *ClientDid) {
		opts.Owner = owner
	}
}

func WithVerificationMethods(verificationMethods []*VerificationMethod) ClientDidDocumentOption {
	return func(opts *ClientDid) {
		opts.VerificationMethods = verificationMethods
	}
}

const InitialSequence uint64 = 0

func NewDidDocumentWithSeq(doc *ClientDid, seq uint64) ClientDidWithSeq {
	return ClientDidWithSeq{
		Document: doc,
		Sequence: seq,
	}
}

func WithAuthentications(authentications []VerificationRelationship) ClientDidDocumentOption {
	return func(opts *ClientDid) {
		opts.Authentications = authentications
	}
}

func (d ClientDidWithSeq) Empty() bool {
	return d.Document == nil || d.Document.Empty() && d.Sequence == InitialSequence
}

func (d ClientDidWithSeq) Deactivated() bool {
	return d.Document.Empty() && d.Sequence != InitialSequence
}

func (doc ClientDid) VerificationMethodFrom(relationships []VerificationRelationship, id string) (VerificationMethod, bool) {
	for _, relationship := range relationships {
		if relationship.hasDedicatedMethod() {
			veriMethod := relationship.GetVerificationMethod()
			if veriMethod.Id == id {
				return *veriMethod, true
			}
		} else {
			veriMethodID := relationship.GetVerificationMethodId()
			if veriMethodID == id {
				return doc.VerificationMethodByID(veriMethodID)
			}
		}
	}

	return VerificationMethod{}, false
}

func (d ClientDidWithSeq) Deactivate(newSeq uint64) ClientDidWithSeq {
	return NewDidDocumentWithSeq(&ClientDid{}, newSeq)
}

type Address = bytes.HexBytes

type PubKey interface {
	Address() Address
	Bytes() []byte
	VerifySignature(msg []byte, sig []byte) bool
	Equals(PubKey) bool
	Type() string
}

func Verify(signature []byte, signableData sdkcodec.ProtoMarshaler, seq uint64, pubKey crypto.PubKey) (uint64, bool) {

	signBytes, err := mustGetSignBytesWithSeq(signableData, seq)
	if err != nil {
		return 0, false
	}
	if !pubKey.VerifySignature(signBytes, signature) {
		return 0, false
	}
	return nextSequence(seq), true
}

func mustGetSignBytesWithSeq(signableData sdkcodec.ProtoMarshaler, seq uint64) ([]byte, error) {
	dAtA, err := signableData.Marshal()
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[mustGetSignBytesWithSeq][Marshal] failed. SignableData is not valid.")
	}
	dataWithSeq := ClientDataWithSeq{
		Data:     dAtA,
		Sequence: seq,
	}

	dAtA, err = dataWithSeq.Marshal()

	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, errors.ErrInvalidDidDocument)
	}
	return dAtA, nil
}

func nextSequence(seq uint64) uint64 {
	return seq + 1
}

func ParseVerificationMethodId(id string, did string) (string, error) {
	methodId := id
	if !ValidateVerificationMethodID(id, did) {
		return "", sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "[ParseVerificationMethodId][ValidateVerificationMethodID] failed for verificationMethodID: %v, did: %v", id, did)
	}
	return methodId, nil
}
