package keeper

import (
	"context"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/neutron-org/neutron/v3/x/dex/types"
)

func (k Keeper) LimitOrderTrancheUserAll(
	c context.Context,
	req *types.QueryAllLimitOrderTrancheUserRequest,
) (*types.QueryAllLimitOrderTrancheUserResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var limitOrderTrancheUsers []*types.LimitOrderTrancheUser
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	limitOrderTrancheUserStore := prefix.NewStore(store, types.KeyPrefix(types.LimitOrderTrancheUserKeyPrefix))

	pageRes, err := query.Paginate(limitOrderTrancheUserStore, req.Pagination, func(key, value []byte) error {
		limitOrderTrancheUser := &types.LimitOrderTrancheUser{}
		if err := k.cdc.Unmarshal(value, limitOrderTrancheUser); err != nil {
			return err
		}

		limitOrderTrancheUsers = append(limitOrderTrancheUsers, limitOrderTrancheUser)

		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllLimitOrderTrancheUserResponse{
		LimitOrderTrancheUser: limitOrderTrancheUsers,
		Pagination:            pageRes,
	}, nil
}

func (k Keeper) CalcWithdrawAmount(ctx sdk.Context, trancheUser types.LimitOrderTrancheUser) (amount math.Int, err error) {
	tradePairID, tickIndex := trancheUser.TradePairId, trancheUser.TickIndexTakerToMaker

	tranche, _, found := k.FindLimitOrderTranche(
		ctx,
		&types.LimitOrderTrancheKey{
			TradePairId:           tradePairID,
			TickIndexTakerToMaker: tickIndex,
			TrancheKey:            trancheUser.TrancheKey,
		},
	)

	if found != true {
		return math.ZeroInt(), status.Error(codes.NotFound, "Tranche not found")
	}
	_, amountOut := tranche.Withdraw(&trancheUser)

	return amountOut.TruncateInt(), nil
}
func (k Keeper) LimitOrderTrancheUser(c context.Context,
	req *types.QueryGetLimitOrderTrancheUserRequest,
) (*types.QueryGetLimitOrderTrancheUserResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	trancheUser, found := k.GetLimitOrderTrancheUser(
		ctx,
		req.Address,
		req.TrancheKey,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	resp := &types.QueryGetLimitOrderTrancheUserResponse{LimitOrderTrancheUser: trancheUser}
	if req.CalcWithdrawableAmount {
		withdrawAmt, err := k.CalcWithdrawAmount(ctx, *trancheUser)
		if err != nil {
			return nil, err
		}
		resp.WithdrawableAmount = &withdrawAmt
	}

	return resp, nil
}

func (k Keeper) LimitOrderTrancheUserAllByAddress(
	goCtx context.Context,
	req *types.QueryAllUserLimitOrdersRequest,
) (*types.QueryAllUserLimitOrdersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, err
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	var limitOrderTrancheUserList []*types.LimitOrderTrancheUser
	addressPrefix := types.LimitOrderTrancheUserAddressPrefix(addr.String())
	store := prefix.NewStore(ctx.KVStore(k.storeKey), addressPrefix)

	pageRes, err := query.Paginate(store, req.Pagination, func(key, value []byte) error {
		trancheUser := &types.LimitOrderTrancheUser{}
		if err := k.cdc.Unmarshal(value, trancheUser); err != nil {
			return err
		}

		limitOrderTrancheUserList = append(limitOrderTrancheUserList, trancheUser)

		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllUserLimitOrdersResponse{
		LimitOrders: limitOrderTrancheUserList,
		Pagination:  pageRes,
	}, nil
}
