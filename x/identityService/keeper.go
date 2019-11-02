package identityService

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/bank"

	sdk "github.com/cosmos/cosmos-sdk/types"
)


type Keeper struct {
	coinKeeper bank.Keeper
	storeKey  sdk.StoreKey
	cdc *codec.Codec
}

func (k Keeper) SetCert(ctx sdk.Context, name string, account Identity) {
	if account.Owner.Empty() {
		return
	}
	if !(account.Certifcate){
		return
	}
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(name), k.cdc.MustMarshalBinaryBare(account))
}

func (k Keeper) GetCert(ctx sdk.Context, name string) Identity {
	store := ctx.KVStore(k.storeKey)
	if !store.Has([]byte(name)) {
		return NewIdentity()
	}
	bz := store.Get([]byte(name))
	var identity Identity
	k.cdc.MustUnmarshalBinaryBare(bz, &identity)
	return identity
}

func (k Keeper) ResolveName(ctx sdk.Context, name string) string {
	return k.GetCert(ctx, name).ReadableName
}

func (k Keeper) SetName(ctx sdk.Context, name string, value string) {
	account := k.GetCert(ctx, name)
	account.ReadableName = value
	account.certificate = true
	k.SetCert(ctx, name, account)
}

func (k Keeper) HasOwner(ctx sdk.Context, name string) bool {
	return !k.GetCert(ctx, name).Owner.Empty()
}

func (k Keeper) GetOwner(ctx sdk.Context, name string) sdk.AccAddress {
	return k.GetCert(ctx, name).Owner
}

func (k Keeper) SetOwner(ctx sdk.Context, name string, owner sdk.AccAddress) {
	account := k.GetCert(ctx, name)
	account.Owner = owner
	k.SetCert(ctx, name, account)
}

func (k Keeper) GetVerificationStatus(ctx sdk.Context, name string) bool {
	return k.GetCert(ctx, name).Certificate
}

func (k Keeper) Certify(ctx sdk.Context, name string) {
	account := k.GetCert(ctx, name)
	account.Certificate = true
	k.SetCert(ctx, name, account)
}

func (k Keeper) GetNamesIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, []byte{})
}

func NewKeeper(coinKeeper bank.Keeper, storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		coinKeeper: coinKeeper,
		storeKey:   storeKey,
		cdc:        cdc,
	}
}