 <!-- This file is auto-generated. Please do not modify it yourself. -->
# Protobuf Documentation
<a name="top"></a>

## Table of Contents

- [gert/auction/v1beta1/auction.proto](#gert/auction/v1beta1/auction.proto)
    - [BaseAuction](#gert.auction.v1beta1.BaseAuction)
    - [CollateralAuction](#gert.auction.v1beta1.CollateralAuction)
    - [DebtAuction](#gert.auction.v1beta1.DebtAuction)
    - [SurplusAuction](#gert.auction.v1beta1.SurplusAuction)
    - [WeightedAddresses](#gert.auction.v1beta1.WeightedAddresses)
  
- [gert/auction/v1beta1/genesis.proto](#gert/auction/v1beta1/genesis.proto)
    - [GenesisState](#gert.auction.v1beta1.GenesisState)
    - [Params](#gert.auction.v1beta1.Params)
  
- [gert/auction/v1beta1/query.proto](#gert/auction/v1beta1/query.proto)
    - [QueryAuctionRequest](#gert.auction.v1beta1.QueryAuctionRequest)
    - [QueryAuctionResponse](#gert.auction.v1beta1.QueryAuctionResponse)
    - [QueryAuctionsRequest](#gert.auction.v1beta1.QueryAuctionsRequest)
    - [QueryAuctionsResponse](#gert.auction.v1beta1.QueryAuctionsResponse)
    - [QueryNextAuctionIDRequest](#gert.auction.v1beta1.QueryNextAuctionIDRequest)
    - [QueryNextAuctionIDResponse](#gert.auction.v1beta1.QueryNextAuctionIDResponse)
    - [QueryParamsRequest](#gert.auction.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#gert.auction.v1beta1.QueryParamsResponse)
  
    - [Query](#gert.auction.v1beta1.Query)
  
- [gert/auction/v1beta1/tx.proto](#gert/auction/v1beta1/tx.proto)
    - [MsgPlaceBid](#gert.auction.v1beta1.MsgPlaceBid)
    - [MsgPlaceBidResponse](#gert.auction.v1beta1.MsgPlaceBidResponse)
  
    - [Msg](#gert.auction.v1beta1.Msg)
  
- [gert/bep3/v1beta1/bep3.proto](#gert/bep3/v1beta1/bep3.proto)
    - [AssetParam](#gert.bep3.v1beta1.AssetParam)
    - [AssetSupply](#gert.bep3.v1beta1.AssetSupply)
    - [AtomicSwap](#gert.bep3.v1beta1.AtomicSwap)
    - [Params](#gert.bep3.v1beta1.Params)
    - [SupplyLimit](#gert.bep3.v1beta1.SupplyLimit)
  
    - [SwapDirection](#gert.bep3.v1beta1.SwapDirection)
    - [SwapStatus](#gert.bep3.v1beta1.SwapStatus)
  
- [gert/bep3/v1beta1/genesis.proto](#gert/bep3/v1beta1/genesis.proto)
    - [GenesisState](#gert.bep3.v1beta1.GenesisState)
  
- [gert/bep3/v1beta1/query.proto](#gert/bep3/v1beta1/query.proto)
    - [AssetSupplyResponse](#gert.bep3.v1beta1.AssetSupplyResponse)
    - [AtomicSwapResponse](#gert.bep3.v1beta1.AtomicSwapResponse)
    - [QueryAssetSuppliesRequest](#gert.bep3.v1beta1.QueryAssetSuppliesRequest)
    - [QueryAssetSuppliesResponse](#gert.bep3.v1beta1.QueryAssetSuppliesResponse)
    - [QueryAssetSupplyRequest](#gert.bep3.v1beta1.QueryAssetSupplyRequest)
    - [QueryAssetSupplyResponse](#gert.bep3.v1beta1.QueryAssetSupplyResponse)
    - [QueryAtomicSwapRequest](#gert.bep3.v1beta1.QueryAtomicSwapRequest)
    - [QueryAtomicSwapResponse](#gert.bep3.v1beta1.QueryAtomicSwapResponse)
    - [QueryAtomicSwapsRequest](#gert.bep3.v1beta1.QueryAtomicSwapsRequest)
    - [QueryAtomicSwapsResponse](#gert.bep3.v1beta1.QueryAtomicSwapsResponse)
    - [QueryParamsRequest](#gert.bep3.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#gert.bep3.v1beta1.QueryParamsResponse)
  
    - [Query](#gert.bep3.v1beta1.Query)
  
- [gert/bep3/v1beta1/tx.proto](#gert/bep3/v1beta1/tx.proto)
    - [MsgClaimAtomicSwap](#gert.bep3.v1beta1.MsgClaimAtomicSwap)
    - [MsgClaimAtomicSwapResponse](#gert.bep3.v1beta1.MsgClaimAtomicSwapResponse)
    - [MsgCreateAtomicSwap](#gert.bep3.v1beta1.MsgCreateAtomicSwap)
    - [MsgCreateAtomicSwapResponse](#gert.bep3.v1beta1.MsgCreateAtomicSwapResponse)
    - [MsgRefundAtomicSwap](#gert.bep3.v1beta1.MsgRefundAtomicSwap)
    - [MsgRefundAtomicSwapResponse](#gert.bep3.v1beta1.MsgRefundAtomicSwapResponse)
  
    - [Msg](#gert.bep3.v1beta1.Msg)
  
- [gert/cdp/v1beta1/cdp.proto](#gert/cdp/v1beta1/cdp.proto)
    - [CDP](#gert.cdp.v1beta1.CDP)
    - [Deposit](#gert.cdp.v1beta1.Deposit)
    - [OwnerCDPIndex](#gert.cdp.v1beta1.OwnerCDPIndex)
    - [TotalCollateral](#gert.cdp.v1beta1.TotalCollateral)
    - [TotalPrincipal](#gert.cdp.v1beta1.TotalPrincipal)
  
- [gert/cdp/v1beta1/genesis.proto](#gert/cdp/v1beta1/genesis.proto)
    - [CollateralParam](#gert.cdp.v1beta1.CollateralParam)
    - [DebtParam](#gert.cdp.v1beta1.DebtParam)
    - [GenesisAccumulationTime](#gert.cdp.v1beta1.GenesisAccumulationTime)
    - [GenesisState](#gert.cdp.v1beta1.GenesisState)
    - [GenesisTotalPrincipal](#gert.cdp.v1beta1.GenesisTotalPrincipal)
    - [Params](#gert.cdp.v1beta1.Params)
  
- [gert/cdp/v1beta1/query.proto](#gert/cdp/v1beta1/query.proto)
    - [CDPResponse](#gert.cdp.v1beta1.CDPResponse)
    - [QueryAccountsRequest](#gert.cdp.v1beta1.QueryAccountsRequest)
    - [QueryAccountsResponse](#gert.cdp.v1beta1.QueryAccountsResponse)
    - [QueryCdpRequest](#gert.cdp.v1beta1.QueryCdpRequest)
    - [QueryCdpResponse](#gert.cdp.v1beta1.QueryCdpResponse)
    - [QueryCdpsRequest](#gert.cdp.v1beta1.QueryCdpsRequest)
    - [QueryCdpsResponse](#gert.cdp.v1beta1.QueryCdpsResponse)
    - [QueryDepositsRequest](#gert.cdp.v1beta1.QueryDepositsRequest)
    - [QueryDepositsResponse](#gert.cdp.v1beta1.QueryDepositsResponse)
    - [QueryParamsRequest](#gert.cdp.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#gert.cdp.v1beta1.QueryParamsResponse)
    - [QueryTotalCollateralRequest](#gert.cdp.v1beta1.QueryTotalCollateralRequest)
    - [QueryTotalCollateralResponse](#gert.cdp.v1beta1.QueryTotalCollateralResponse)
    - [QueryTotalPrincipalRequest](#gert.cdp.v1beta1.QueryTotalPrincipalRequest)
    - [QueryTotalPrincipalResponse](#gert.cdp.v1beta1.QueryTotalPrincipalResponse)
  
    - [Query](#gert.cdp.v1beta1.Query)
  
- [gert/cdp/v1beta1/tx.proto](#gert/cdp/v1beta1/tx.proto)
    - [MsgCreateCDP](#gert.cdp.v1beta1.MsgCreateCDP)
    - [MsgCreateCDPResponse](#gert.cdp.v1beta1.MsgCreateCDPResponse)
    - [MsgDeposit](#gert.cdp.v1beta1.MsgDeposit)
    - [MsgDepositResponse](#gert.cdp.v1beta1.MsgDepositResponse)
    - [MsgDrawDebt](#gert.cdp.v1beta1.MsgDrawDebt)
    - [MsgDrawDebtResponse](#gert.cdp.v1beta1.MsgDrawDebtResponse)
    - [MsgLiquidate](#gert.cdp.v1beta1.MsgLiquidate)
    - [MsgLiquidateResponse](#gert.cdp.v1beta1.MsgLiquidateResponse)
    - [MsgRepayDebt](#gert.cdp.v1beta1.MsgRepayDebt)
    - [MsgRepayDebtResponse](#gert.cdp.v1beta1.MsgRepayDebtResponse)
    - [MsgWithdraw](#gert.cdp.v1beta1.MsgWithdraw)
    - [MsgWithdrawResponse](#gert.cdp.v1beta1.MsgWithdrawResponse)
  
    - [Msg](#gert.cdp.v1beta1.Msg)
  
- [gert/committee/v1beta1/committee.proto](#gert/committee/v1beta1/committee.proto)
    - [BaseCommittee](#gert.committee.v1beta1.BaseCommittee)
    - [MemberCommittee](#gert.committee.v1beta1.MemberCommittee)
    - [TokenCommittee](#gert.committee.v1beta1.TokenCommittee)
  
    - [TallyOption](#gert.committee.v1beta1.TallyOption)
  
- [gert/committee/v1beta1/genesis.proto](#gert/committee/v1beta1/genesis.proto)
    - [GenesisState](#gert.committee.v1beta1.GenesisState)
    - [Proposal](#gert.committee.v1beta1.Proposal)
    - [Vote](#gert.committee.v1beta1.Vote)
  
    - [VoteType](#gert.committee.v1beta1.VoteType)
  
- [gert/committee/v1beta1/permissions.proto](#gert/committee/v1beta1/permissions.proto)
    - [AllowedParamsChange](#gert.committee.v1beta1.AllowedParamsChange)
    - [GodPermission](#gert.committee.v1beta1.GodPermission)
    - [ParamsChangePermission](#gert.committee.v1beta1.ParamsChangePermission)
    - [SoftwareUpgradePermission](#gert.committee.v1beta1.SoftwareUpgradePermission)
    - [SubparamRequirement](#gert.committee.v1beta1.SubparamRequirement)
    - [TextPermission](#gert.committee.v1beta1.TextPermission)
  
- [gert/committee/v1beta1/proposal.proto](#gert/committee/v1beta1/proposal.proto)
    - [CommitteeChangeProposal](#gert.committee.v1beta1.CommitteeChangeProposal)
    - [CommitteeDeleteProposal](#gert.committee.v1beta1.CommitteeDeleteProposal)
  
- [gert/committee/v1beta1/query.proto](#gert/committee/v1beta1/query.proto)
    - [QueryCommitteeRequest](#gert.committee.v1beta1.QueryCommitteeRequest)
    - [QueryCommitteeResponse](#gert.committee.v1beta1.QueryCommitteeResponse)
    - [QueryCommitteesRequest](#gert.committee.v1beta1.QueryCommitteesRequest)
    - [QueryCommitteesResponse](#gert.committee.v1beta1.QueryCommitteesResponse)
    - [QueryNextProposalIDRequest](#gert.committee.v1beta1.QueryNextProposalIDRequest)
    - [QueryNextProposalIDResponse](#gert.committee.v1beta1.QueryNextProposalIDResponse)
    - [QueryProposalRequest](#gert.committee.v1beta1.QueryProposalRequest)
    - [QueryProposalResponse](#gert.committee.v1beta1.QueryProposalResponse)
    - [QueryProposalsRequest](#gert.committee.v1beta1.QueryProposalsRequest)
    - [QueryProposalsResponse](#gert.committee.v1beta1.QueryProposalsResponse)
    - [QueryRawParamsRequest](#gert.committee.v1beta1.QueryRawParamsRequest)
    - [QueryRawParamsResponse](#gert.committee.v1beta1.QueryRawParamsResponse)
    - [QueryTallyRequest](#gert.committee.v1beta1.QueryTallyRequest)
    - [QueryTallyResponse](#gert.committee.v1beta1.QueryTallyResponse)
    - [QueryVoteRequest](#gert.committee.v1beta1.QueryVoteRequest)
    - [QueryVoteResponse](#gert.committee.v1beta1.QueryVoteResponse)
    - [QueryVotesRequest](#gert.committee.v1beta1.QueryVotesRequest)
    - [QueryVotesResponse](#gert.committee.v1beta1.QueryVotesResponse)
  
    - [Query](#gert.committee.v1beta1.Query)
  
- [gert/committee/v1beta1/tx.proto](#gert/committee/v1beta1/tx.proto)
    - [MsgSubmitProposal](#gert.committee.v1beta1.MsgSubmitProposal)
    - [MsgSubmitProposalResponse](#gert.committee.v1beta1.MsgSubmitProposalResponse)
    - [MsgVote](#gert.committee.v1beta1.MsgVote)
    - [MsgVoteResponse](#gert.committee.v1beta1.MsgVoteResponse)
  
    - [Msg](#gert.committee.v1beta1.Msg)
  
- [gert/evmutil/v1beta1/genesis.proto](#gert/evmutil/v1beta1/genesis.proto)
    - [Account](#gert.evmutil.v1beta1.Account)
    - [GenesisState](#gert.evmutil.v1beta1.GenesisState)
  
- [gert/hard/v1beta1/hard.proto](#gert/hard/v1beta1/hard.proto)
    - [Borrow](#gert.hard.v1beta1.Borrow)
    - [BorrowInterestFactor](#gert.hard.v1beta1.BorrowInterestFactor)
    - [BorrowLimit](#gert.hard.v1beta1.BorrowLimit)
    - [CoinsProto](#gert.hard.v1beta1.CoinsProto)
    - [Deposit](#gert.hard.v1beta1.Deposit)
    - [InterestRateModel](#gert.hard.v1beta1.InterestRateModel)
    - [MoneyMarket](#gert.hard.v1beta1.MoneyMarket)
    - [Params](#gert.hard.v1beta1.Params)
    - [SupplyInterestFactor](#gert.hard.v1beta1.SupplyInterestFactor)
  
- [gert/hard/v1beta1/genesis.proto](#gert/hard/v1beta1/genesis.proto)
    - [GenesisAccumulationTime](#gert.hard.v1beta1.GenesisAccumulationTime)
    - [GenesisState](#gert.hard.v1beta1.GenesisState)
  
- [gert/hard/v1beta1/query.proto](#gert/hard/v1beta1/query.proto)
    - [BorrowInterestFactorResponse](#gert.hard.v1beta1.BorrowInterestFactorResponse)
    - [BorrowResponse](#gert.hard.v1beta1.BorrowResponse)
    - [DepositResponse](#gert.hard.v1beta1.DepositResponse)
    - [InterestFactor](#gert.hard.v1beta1.InterestFactor)
    - [MoneyMarketInterestRate](#gert.hard.v1beta1.MoneyMarketInterestRate)
    - [QueryAccountsRequest](#gert.hard.v1beta1.QueryAccountsRequest)
    - [QueryAccountsResponse](#gert.hard.v1beta1.QueryAccountsResponse)
    - [QueryBorrowsRequest](#gert.hard.v1beta1.QueryBorrowsRequest)
    - [QueryBorrowsResponse](#gert.hard.v1beta1.QueryBorrowsResponse)
    - [QueryDepositsRequest](#gert.hard.v1beta1.QueryDepositsRequest)
    - [QueryDepositsResponse](#gert.hard.v1beta1.QueryDepositsResponse)
    - [QueryInterestFactorsRequest](#gert.hard.v1beta1.QueryInterestFactorsRequest)
    - [QueryInterestFactorsResponse](#gert.hard.v1beta1.QueryInterestFactorsResponse)
    - [QueryInterestRateRequest](#gert.hard.v1beta1.QueryInterestRateRequest)
    - [QueryInterestRateResponse](#gert.hard.v1beta1.QueryInterestRateResponse)
    - [QueryParamsRequest](#gert.hard.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#gert.hard.v1beta1.QueryParamsResponse)
    - [QueryReservesRequest](#gert.hard.v1beta1.QueryReservesRequest)
    - [QueryReservesResponse](#gert.hard.v1beta1.QueryReservesResponse)
    - [QueryTotalBorrowedRequest](#gert.hard.v1beta1.QueryTotalBorrowedRequest)
    - [QueryTotalBorrowedResponse](#gert.hard.v1beta1.QueryTotalBorrowedResponse)
    - [QueryTotalDepositedRequest](#gert.hard.v1beta1.QueryTotalDepositedRequest)
    - [QueryTotalDepositedResponse](#gert.hard.v1beta1.QueryTotalDepositedResponse)
    - [QueryUnsyncedBorrowsRequest](#gert.hard.v1beta1.QueryUnsyncedBorrowsRequest)
    - [QueryUnsyncedBorrowsResponse](#gert.hard.v1beta1.QueryUnsyncedBorrowsResponse)
    - [QueryUnsyncedDepositsRequest](#gert.hard.v1beta1.QueryUnsyncedDepositsRequest)
    - [QueryUnsyncedDepositsResponse](#gert.hard.v1beta1.QueryUnsyncedDepositsResponse)
    - [SupplyInterestFactorResponse](#gert.hard.v1beta1.SupplyInterestFactorResponse)
  
    - [Query](#gert.hard.v1beta1.Query)
  
- [gert/hard/v1beta1/tx.proto](#gert/hard/v1beta1/tx.proto)
    - [MsgBorrow](#gert.hard.v1beta1.MsgBorrow)
    - [MsgBorrowResponse](#gert.hard.v1beta1.MsgBorrowResponse)
    - [MsgDeposit](#gert.hard.v1beta1.MsgDeposit)
    - [MsgDepositResponse](#gert.hard.v1beta1.MsgDepositResponse)
    - [MsgLiquidate](#gert.hard.v1beta1.MsgLiquidate)
    - [MsgLiquidateResponse](#gert.hard.v1beta1.MsgLiquidateResponse)
    - [MsgRepay](#gert.hard.v1beta1.MsgRepay)
    - [MsgRepayResponse](#gert.hard.v1beta1.MsgRepayResponse)
    - [MsgWithdraw](#gert.hard.v1beta1.MsgWithdraw)
    - [MsgWithdrawResponse](#gert.hard.v1beta1.MsgWithdrawResponse)
  
    - [Msg](#gert.hard.v1beta1.Msg)
  
- [gert/incentive/v1beta1/claims.proto](#gert/incentive/v1beta1/claims.proto)
    - [BaseClaim](#gert.incentive.v1beta1.BaseClaim)
    - [BaseMultiClaim](#gert.incentive.v1beta1.BaseMultiClaim)
    - [DelegatorClaim](#gert.incentive.v1beta1.DelegatorClaim)
    - [HardLiquidityProviderClaim](#gert.incentive.v1beta1.HardLiquidityProviderClaim)
    - [MultiRewardIndex](#gert.incentive.v1beta1.MultiRewardIndex)
    - [MultiRewardIndexesProto](#gert.incentive.v1beta1.MultiRewardIndexesProto)
    - [RewardIndex](#gert.incentive.v1beta1.RewardIndex)
    - [RewardIndexesProto](#gert.incentive.v1beta1.RewardIndexesProto)
    - [SavingsClaim](#gert.incentive.v1beta1.SavingsClaim)
    - [SwapClaim](#gert.incentive.v1beta1.SwapClaim)
    - [USDXMintingClaim](#gert.incentive.v1beta1.USDXMintingClaim)
  
- [gert/incentive/v1beta1/params.proto](#gert/incentive/v1beta1/params.proto)
    - [MultiRewardPeriod](#gert.incentive.v1beta1.MultiRewardPeriod)
    - [Multiplier](#gert.incentive.v1beta1.Multiplier)
    - [MultipliersPerDenom](#gert.incentive.v1beta1.MultipliersPerDenom)
    - [Params](#gert.incentive.v1beta1.Params)
    - [RewardPeriod](#gert.incentive.v1beta1.RewardPeriod)
  
- [gert/incentive/v1beta1/genesis.proto](#gert/incentive/v1beta1/genesis.proto)
    - [AccumulationTime](#gert.incentive.v1beta1.AccumulationTime)
    - [GenesisRewardState](#gert.incentive.v1beta1.GenesisRewardState)
    - [GenesisState](#gert.incentive.v1beta1.GenesisState)
  
- [gert/incentive/v1beta1/tx.proto](#gert/incentive/v1beta1/tx.proto)
    - [MsgClaimDelegatorReward](#gert.incentive.v1beta1.MsgClaimDelegatorReward)
    - [MsgClaimDelegatorRewardResponse](#gert.incentive.v1beta1.MsgClaimDelegatorRewardResponse)
    - [MsgClaimHardReward](#gert.incentive.v1beta1.MsgClaimHardReward)
    - [MsgClaimHardRewardResponse](#gert.incentive.v1beta1.MsgClaimHardRewardResponse)
    - [MsgClaimSavingsReward](#gert.incentive.v1beta1.MsgClaimSavingsReward)
    - [MsgClaimSavingsRewardResponse](#gert.incentive.v1beta1.MsgClaimSavingsRewardResponse)
    - [MsgClaimSwapReward](#gert.incentive.v1beta1.MsgClaimSwapReward)
    - [MsgClaimSwapRewardResponse](#gert.incentive.v1beta1.MsgClaimSwapRewardResponse)
    - [MsgClaimUSDXMintingReward](#gert.incentive.v1beta1.MsgClaimUSDXMintingReward)
    - [MsgClaimUSDXMintingRewardResponse](#gert.incentive.v1beta1.MsgClaimUSDXMintingRewardResponse)
    - [Selection](#gert.incentive.v1beta1.Selection)
  
    - [Msg](#gert.incentive.v1beta1.Msg)
  
- [gert/issuance/v1beta1/genesis.proto](#gert/issuance/v1beta1/genesis.proto)
    - [Asset](#gert.issuance.v1beta1.Asset)
    - [AssetSupply](#gert.issuance.v1beta1.AssetSupply)
    - [GenesisState](#gert.issuance.v1beta1.GenesisState)
    - [Params](#gert.issuance.v1beta1.Params)
    - [RateLimit](#gert.issuance.v1beta1.RateLimit)
  
- [gert/issuance/v1beta1/query.proto](#gert/issuance/v1beta1/query.proto)
    - [QueryParamsRequest](#gert.issuance.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#gert.issuance.v1beta1.QueryParamsResponse)
  
    - [Query](#gert.issuance.v1beta1.Query)
  
- [gert/issuance/v1beta1/tx.proto](#gert/issuance/v1beta1/tx.proto)
    - [MsgBlockAddress](#gert.issuance.v1beta1.MsgBlockAddress)
    - [MsgBlockAddressResponse](#gert.issuance.v1beta1.MsgBlockAddressResponse)
    - [MsgIssueTokens](#gert.issuance.v1beta1.MsgIssueTokens)
    - [MsgIssueTokensResponse](#gert.issuance.v1beta1.MsgIssueTokensResponse)
    - [MsgRedeemTokens](#gert.issuance.v1beta1.MsgRedeemTokens)
    - [MsgRedeemTokensResponse](#gert.issuance.v1beta1.MsgRedeemTokensResponse)
    - [MsgSetPauseStatus](#gert.issuance.v1beta1.MsgSetPauseStatus)
    - [MsgSetPauseStatusResponse](#gert.issuance.v1beta1.MsgSetPauseStatusResponse)
    - [MsgUnblockAddress](#gert.issuance.v1beta1.MsgUnblockAddress)
    - [MsgUnblockAddressResponse](#gert.issuance.v1beta1.MsgUnblockAddressResponse)
  
    - [Msg](#gert.issuance.v1beta1.Msg)
  
- [gert/gertdist/v1beta1/params.proto](#gert/gertdist/v1beta1/params.proto)
    - [Params](#gert.gertdist.v1beta1.Params)
    - [Period](#gert.gertdist.v1beta1.Period)
  
- [gert/gertdist/v1beta1/genesis.proto](#gert/gertdist/v1beta1/genesis.proto)
    - [GenesisState](#gert.gertdist.v1beta1.GenesisState)
  
- [gert/gertdist/v1beta1/proposal.proto](#gert/gertdist/v1beta1/proposal.proto)
    - [CommunityPoolMultiSpendProposal](#gert.gertdist.v1beta1.CommunityPoolMultiSpendProposal)
    - [CommunityPoolMultiSpendProposalJSON](#gert.gertdist.v1beta1.CommunityPoolMultiSpendProposalJSON)
    - [MultiSpendRecipient](#gert.gertdist.v1beta1.MultiSpendRecipient)
  
- [gert/gertdist/v1beta1/query.proto](#gert/gertdist/v1beta1/query.proto)
    - [QueryBalanceRequest](#gert.gertdist.v1beta1.QueryBalanceRequest)
    - [QueryBalanceResponse](#gert.gertdist.v1beta1.QueryBalanceResponse)
    - [QueryParamsRequest](#gert.gertdist.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#gert.gertdist.v1beta1.QueryParamsResponse)
  
    - [Query](#gert.gertdist.v1beta1.Query)
  
- [gert/pricefeed/v1beta1/store.proto](#gert/pricefeed/v1beta1/store.proto)
    - [CurrentPrice](#gert.pricefeed.v1beta1.CurrentPrice)
    - [Market](#gert.pricefeed.v1beta1.Market)
    - [Params](#gert.pricefeed.v1beta1.Params)
    - [PostedPrice](#gert.pricefeed.v1beta1.PostedPrice)
  
- [gert/pricefeed/v1beta1/genesis.proto](#gert/pricefeed/v1beta1/genesis.proto)
    - [GenesisState](#gert.pricefeed.v1beta1.GenesisState)
  
- [gert/pricefeed/v1beta1/query.proto](#gert/pricefeed/v1beta1/query.proto)
    - [CurrentPriceResponse](#gert.pricefeed.v1beta1.CurrentPriceResponse)
    - [MarketResponse](#gert.pricefeed.v1beta1.MarketResponse)
    - [PostedPriceResponse](#gert.pricefeed.v1beta1.PostedPriceResponse)
    - [QueryMarketsRequest](#gert.pricefeed.v1beta1.QueryMarketsRequest)
    - [QueryMarketsResponse](#gert.pricefeed.v1beta1.QueryMarketsResponse)
    - [QueryOraclesRequest](#gert.pricefeed.v1beta1.QueryOraclesRequest)
    - [QueryOraclesResponse](#gert.pricefeed.v1beta1.QueryOraclesResponse)
    - [QueryParamsRequest](#gert.pricefeed.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#gert.pricefeed.v1beta1.QueryParamsResponse)
    - [QueryPriceRequest](#gert.pricefeed.v1beta1.QueryPriceRequest)
    - [QueryPriceResponse](#gert.pricefeed.v1beta1.QueryPriceResponse)
    - [QueryPricesRequest](#gert.pricefeed.v1beta1.QueryPricesRequest)
    - [QueryPricesResponse](#gert.pricefeed.v1beta1.QueryPricesResponse)
    - [QueryRawPricesRequest](#gert.pricefeed.v1beta1.QueryRawPricesRequest)
    - [QueryRawPricesResponse](#gert.pricefeed.v1beta1.QueryRawPricesResponse)
  
    - [Query](#gert.pricefeed.v1beta1.Query)
  
- [gert/pricefeed/v1beta1/tx.proto](#gert/pricefeed/v1beta1/tx.proto)
    - [MsgPostPrice](#gert.pricefeed.v1beta1.MsgPostPrice)
    - [MsgPostPriceResponse](#gert.pricefeed.v1beta1.MsgPostPriceResponse)
  
    - [Msg](#gert.pricefeed.v1beta1.Msg)
  
- [gert/savings/v1beta1/store.proto](#gert/savings/v1beta1/store.proto)
    - [Deposit](#gert.savings.v1beta1.Deposit)
    - [Params](#gert.savings.v1beta1.Params)
  
- [gert/savings/v1beta1/genesis.proto](#gert/savings/v1beta1/genesis.proto)
    - [GenesisState](#gert.savings.v1beta1.GenesisState)
  
- [gert/savings/v1beta1/query.proto](#gert/savings/v1beta1/query.proto)
    - [QueryDepositsRequest](#gert.savings.v1beta1.QueryDepositsRequest)
    - [QueryDepositsResponse](#gert.savings.v1beta1.QueryDepositsResponse)
    - [QueryParamsRequest](#gert.savings.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#gert.savings.v1beta1.QueryParamsResponse)
  
    - [Query](#gert.savings.v1beta1.Query)
  
- [gert/savings/v1beta1/tx.proto](#gert/savings/v1beta1/tx.proto)
    - [MsgDeposit](#gert.savings.v1beta1.MsgDeposit)
    - [MsgDepositResponse](#gert.savings.v1beta1.MsgDepositResponse)
    - [MsgWithdraw](#gert.savings.v1beta1.MsgWithdraw)
    - [MsgWithdrawResponse](#gert.savings.v1beta1.MsgWithdrawResponse)
  
    - [Msg](#gert.savings.v1beta1.Msg)
  
- [gert/swap/v1beta1/swap.proto](#gert/swap/v1beta1/swap.proto)
    - [AllowedPool](#gert.swap.v1beta1.AllowedPool)
    - [Params](#gert.swap.v1beta1.Params)
    - [PoolRecord](#gert.swap.v1beta1.PoolRecord)
    - [ShareRecord](#gert.swap.v1beta1.ShareRecord)
  
- [gert/swap/v1beta1/genesis.proto](#gert/swap/v1beta1/genesis.proto)
    - [GenesisState](#gert.swap.v1beta1.GenesisState)
  
- [gert/swap/v1beta1/query.proto](#gert/swap/v1beta1/query.proto)
    - [DepositResponse](#gert.swap.v1beta1.DepositResponse)
    - [PoolResponse](#gert.swap.v1beta1.PoolResponse)
    - [QueryDepositsRequest](#gert.swap.v1beta1.QueryDepositsRequest)
    - [QueryDepositsResponse](#gert.swap.v1beta1.QueryDepositsResponse)
    - [QueryParamsRequest](#gert.swap.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#gert.swap.v1beta1.QueryParamsResponse)
    - [QueryPoolsRequest](#gert.swap.v1beta1.QueryPoolsRequest)
    - [QueryPoolsResponse](#gert.swap.v1beta1.QueryPoolsResponse)
  
    - [Query](#gert.swap.v1beta1.Query)
  
- [gert/swap/v1beta1/tx.proto](#gert/swap/v1beta1/tx.proto)
    - [MsgDeposit](#gert.swap.v1beta1.MsgDeposit)
    - [MsgDepositResponse](#gert.swap.v1beta1.MsgDepositResponse)
    - [MsgSwapExactForTokens](#gert.swap.v1beta1.MsgSwapExactForTokens)
    - [MsgSwapExactForTokensResponse](#gert.swap.v1beta1.MsgSwapExactForTokensResponse)
    - [MsgSwapForExactTokens](#gert.swap.v1beta1.MsgSwapForExactTokens)
    - [MsgSwapForExactTokensResponse](#gert.swap.v1beta1.MsgSwapForExactTokensResponse)
    - [MsgWithdraw](#gert.swap.v1beta1.MsgWithdraw)
    - [MsgWithdrawResponse](#gert.swap.v1beta1.MsgWithdrawResponse)
  
    - [Msg](#gert.swap.v1beta1.Msg)
  
- [Scalar Value Types](#scalar-value-types)



<a name="gert/auction/v1beta1/auction.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gert/auction/v1beta1/auction.proto



<a name="gert.auction.v1beta1.BaseAuction"></a>

### BaseAuction
BaseAuction defines common attributes of all auctions


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [uint64](#uint64) |  |  |
| `initiator` | [string](#string) |  |  |
| `lot` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `bidder` | [bytes](#bytes) |  |  |
| `bid` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `has_received_bids` | [bool](#bool) |  |  |
| `end_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `max_end_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="gert.auction.v1beta1.CollateralAuction"></a>

### CollateralAuction
CollateralAuction is a two phase auction.
Initially, in forward auction phase, bids can be placed up to a max bid.
Then it switches to a reverse auction phase, where the initial amount up for auction is bid down.
Unsold Lot is sent to LotReturns, being divided among the addresses by weight.
Collateral auctions are normally used to sell off collateral seized from CDPs.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_auction` | [BaseAuction](#gert.auction.v1beta1.BaseAuction) |  |  |
| `corresponding_debt` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `max_bid` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `lot_returns` | [WeightedAddresses](#gert.auction.v1beta1.WeightedAddresses) |  |  |






<a name="gert.auction.v1beta1.DebtAuction"></a>

### DebtAuction
DebtAuction is a reverse auction that mints what it pays out.
It is normally used to acquire pegged asset to cover the CDP system's debts that were not covered by selling
collateral.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_auction` | [BaseAuction](#gert.auction.v1beta1.BaseAuction) |  |  |
| `corresponding_debt` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="gert.auction.v1beta1.SurplusAuction"></a>

### SurplusAuction
SurplusAuction is a forward auction that burns what it receives from bids.
It is normally used to sell off excess pegged asset acquired by the CDP system.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_auction` | [BaseAuction](#gert.auction.v1beta1.BaseAuction) |  |  |






<a name="gert.auction.v1beta1.WeightedAddresses"></a>

### WeightedAddresses
WeightedAddresses is a type for storing some addresses and associated weights.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `addresses` | [bytes](#bytes) | repeated |  |
| `weights` | [bytes](#bytes) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="gert/auction/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gert/auction/v1beta1/genesis.proto



<a name="gert.auction.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the auction module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `next_auction_id` | [uint64](#uint64) |  |  |
| `params` | [Params](#gert.auction.v1beta1.Params) |  |  |
| `auctions` | [google.protobuf.Any](#google.protobuf.Any) | repeated | Genesis auctions |






<a name="gert.auction.v1beta1.Params"></a>

### Params
Params defines the parameters for the issuance module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `max_auction_duration` | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |
| `forward_bid_duration` | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |
| `reverse_bid_duration` | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |
| `increment_surplus` | [bytes](#bytes) |  |  |
| `increment_debt` | [bytes](#bytes) |  |  |
| `increment_collateral` | [bytes](#bytes) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="gert/auction/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gert/auction/v1beta1/query.proto



<a name="gert.auction.v1beta1.QueryAuctionRequest"></a>

### QueryAuctionRequest
QueryAuctionRequest is the request type for the Query/Auction RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `auction_id` | [uint64](#uint64) |  |  |






<a name="gert.auction.v1beta1.QueryAuctionResponse"></a>

### QueryAuctionResponse
QueryAuctionResponse is the response type for the Query/Auction RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `auction` | [google.protobuf.Any](#google.protobuf.Any) |  |  |






<a name="gert.auction.v1beta1.QueryAuctionsRequest"></a>

### QueryAuctionsRequest
QueryAuctionsRequest is the request type for the Query/Auctions RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `type` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |
| `denom` | [string](#string) |  |  |
| `phase` | [string](#string) |  |  |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="gert.auction.v1beta1.QueryAuctionsResponse"></a>

### QueryAuctionsResponse
QueryAuctionsResponse is the response type for the Query/Auctions RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `auctions` | [google.protobuf.Any](#google.protobuf.Any) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="gert.auction.v1beta1.QueryNextAuctionIDRequest"></a>

### QueryNextAuctionIDRequest
QueryNextAuctionIDRequest defines the request type for querying x/auction next auction ID.






<a name="gert.auction.v1beta1.QueryNextAuctionIDResponse"></a>

### QueryNextAuctionIDResponse
QueryNextAuctionIDResponse defines the response type for querying x/auction next auction ID.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [uint64](#uint64) |  |  |






<a name="gert.auction.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest defines the request type for querying x/auction parameters.






<a name="gert.auction.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse defines the response type for querying x/auction parameters.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#gert.auction.v1beta1.Params) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="gert.auction.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service for auction module

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#gert.auction.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#gert.auction.v1beta1.QueryParamsResponse) | Params queries all parameters of the auction module. | GET|/gert/auction/v1beta1/params|
| `Auction` | [QueryAuctionRequest](#gert.auction.v1beta1.QueryAuctionRequest) | [QueryAuctionResponse](#gert.auction.v1beta1.QueryAuctionResponse) | Auction queries an individual Auction by auction ID | GET|/gert/auction/v1beta1/auctions/{auction_id}|
| `Auctions` | [QueryAuctionsRequest](#gert.auction.v1beta1.QueryAuctionsRequest) | [QueryAuctionsResponse](#gert.auction.v1beta1.QueryAuctionsResponse) | Auctions queries auctions filtered by asset denom, owner address, phase, and auction type | GET|/gert/auction/v1beta1/auctions|
| `NextAuctionID` | [QueryNextAuctionIDRequest](#gert.auction.v1beta1.QueryNextAuctionIDRequest) | [QueryNextAuctionIDResponse](#gert.auction.v1beta1.QueryNextAuctionIDResponse) | NextAuctionID queries the next auction ID | GET|/gert/auction/v1beta1/next-auction-id|

 <!-- end services -->



<a name="gert/auction/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gert/auction/v1beta1/tx.proto



<a name="gert.auction.v1beta1.MsgPlaceBid"></a>

### MsgPlaceBid
MsgPlaceBid represents a message used by bidders to place bids on auctions


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `auction_id` | [uint64](#uint64) |  |  |
| `bidder` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="gert.auction.v1beta1.MsgPlaceBidResponse"></a>

### MsgPlaceBidResponse
MsgPlaceBidResponse defines the Msg/PlaceBid response type.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="gert.auction.v1beta1.Msg"></a>

### Msg
Msg defines the auction Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `PlaceBid` | [MsgPlaceBid](#gert.auction.v1beta1.MsgPlaceBid) | [MsgPlaceBidResponse](#gert.auction.v1beta1.MsgPlaceBidResponse) | PlaceBid message type used by bidders to place bids on auctions | |

 <!-- end services -->



<a name="gert/bep3/v1beta1/bep3.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gert/bep3/v1beta1/bep3.proto



<a name="gert.bep3.v1beta1.AssetParam"></a>

### AssetParam
AssetParam defines parameters for each bep3 asset.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  | denom represents the denominatin for this asset |
| `coin_id` | [int64](#int64) |  | coin_id represents the registered coin type to use (https://github.com/satoshilabs/slips/blob/master/slip-0044.md) |
| `supply_limit` | [SupplyLimit](#gert.bep3.v1beta1.SupplyLimit) |  | supply_limit defines the maximum supply allowed for the asset - a total or time based rate limit |
| `active` | [bool](#bool) |  | active specifies if the asset is live or paused |
| `deputy_address` | [bytes](#bytes) |  | deputy_address the gert address of the deputy |
| `fixed_fee` | [string](#string) |  | fixed_fee defines the fee for incoming swaps |
| `min_swap_amount` | [string](#string) |  | min_swap_amount defines the minimum amount able to be swapped in a single message |
| `max_swap_amount` | [string](#string) |  | max_swap_amount defines the maximum amount able to be swapped in a single message |
| `min_block_lock` | [uint64](#uint64) |  | min_block_lock defined the minimum blocks to lock |
| `max_block_lock` | [uint64](#uint64) |  | min_block_lock defined the maximum blocks to lock |






<a name="gert.bep3.v1beta1.AssetSupply"></a>

### AssetSupply
AssetSupply defines information about an asset's supply.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `incoming_supply` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | incoming_supply represents the incoming supply of an asset |
| `outgoing_supply` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | outgoing_supply represents the outgoing supply of an asset |
| `current_supply` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | current_supply represents the current on-chain supply of an asset |
| `time_limited_current_supply` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | time_limited_current_supply represents the time limited current supply of an asset |
| `time_elapsed` | [google.protobuf.Duration](#google.protobuf.Duration) |  | time_elapsed represents the time elapsed |






<a name="gert.bep3.v1beta1.AtomicSwap"></a>

### AtomicSwap
AtomicSwap defines an atomic swap between chains for the pricefeed module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | amount represents the amount being swapped |
| `random_number_hash` | [bytes](#bytes) |  | random_number_hash represents the hash of the random number |
| `expire_height` | [uint64](#uint64) |  | expire_height represents the height when the swap expires |
| `timestamp` | [int64](#int64) |  | timestamp represents the timestamp of the swap |
| `sender` | [bytes](#bytes) |  | sender is the gert chain sender of the swap |
| `recipient` | [bytes](#bytes) |  | recipient is the gert chain recipient of the swap |
| `sender_other_chain` | [string](#string) |  | sender_other_chain is the sender on the other chain |
| `recipient_other_chain` | [string](#string) |  | recipient_other_chain is the recipient on the other chain |
| `closed_block` | [int64](#int64) |  | closed_block is the block when the swap is closed |
| `status` | [SwapStatus](#gert.bep3.v1beta1.SwapStatus) |  | status represents the current status of the swap |
| `cross_chain` | [bool](#bool) |  | cross_chain identifies whether the atomic swap is cross chain |
| `direction` | [SwapDirection](#gert.bep3.v1beta1.SwapDirection) |  | direction identifies if the swap is incoming or outgoing |






<a name="gert.bep3.v1beta1.Params"></a>

### Params
Params defines the parameters for the bep3 module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `asset_params` | [AssetParam](#gert.bep3.v1beta1.AssetParam) | repeated | asset_params define the parameters for each bep3 asset |






<a name="gert.bep3.v1beta1.SupplyLimit"></a>

### SupplyLimit
SupplyLimit define the absolute and time-based limits for an assets's supply.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `limit` | [string](#string) |  | limit defines the total supply allowed |
| `time_limited` | [bool](#bool) |  | time_limited enables or disables time based supply limiting |
| `time_period` | [google.protobuf.Duration](#google.protobuf.Duration) |  | time_period specifies the duration that time_based_limit is evalulated |
| `time_based_limit` | [string](#string) |  | time_based_limit defines the maximum supply that can be swapped within time_period |





 <!-- end messages -->


<a name="gert.bep3.v1beta1.SwapDirection"></a>

### SwapDirection
SwapDirection is the direction of an AtomicSwap

| Name | Number | Description |
| ---- | ------ | ----------- |
| SWAP_DIRECTION_UNSPECIFIED | 0 | SWAP_DIRECTION_UNSPECIFIED represents unspecified or invalid swap direcation |
| SWAP_DIRECTION_INCOMING | 1 | SWAP_DIRECTION_INCOMING represents is incoming swap (to the gert chain) |
| SWAP_DIRECTION_OUTGOING | 2 | SWAP_DIRECTION_OUTGOING represents an outgoing swap (from the gert chain) |



<a name="gert.bep3.v1beta1.SwapStatus"></a>

### SwapStatus
SwapStatus is the status of an AtomicSwap

| Name | Number | Description |
| ---- | ------ | ----------- |
| SWAP_STATUS_UNSPECIFIED | 0 | SWAP_STATUS_UNSPECIFIED represents an unspecified status |
| SWAP_STATUS_OPEN | 1 | SWAP_STATUS_OPEN represents an open swap |
| SWAP_STATUS_COMPLETED | 2 | SWAP_STATUS_COMPLETED represents a completed swap |
| SWAP_STATUS_EXPIRED | 3 | SWAP_STATUS_EXPIRED represents an expired swap |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="gert/bep3/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gert/bep3/v1beta1/genesis.proto



<a name="gert.bep3.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the pricefeed module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#gert.bep3.v1beta1.Params) |  | params defines all the paramaters of the module. |
| `atomic_swaps` | [AtomicSwap](#gert.bep3.v1beta1.AtomicSwap) | repeated | atomic_swaps represents the state of stored atomic swaps |
| `supplies` | [AssetSupply](#gert.bep3.v1beta1.AssetSupply) | repeated | supplies represents the supply information of each atomic swap |
| `previous_block_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | previous_block_time represents the time of the previous block |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="gert/bep3/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gert/bep3/v1beta1/query.proto



<a name="gert.bep3.v1beta1.AssetSupplyResponse"></a>

### AssetSupplyResponse
AssetSupplyResponse defines information about an asset's supply.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `incoming_supply` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | incoming_supply represents the incoming supply of an asset |
| `outgoing_supply` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | outgoing_supply represents the outgoing supply of an asset |
| `current_supply` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | current_supply represents the current on-chain supply of an asset |
| `time_limited_current_supply` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | time_limited_current_supply represents the time limited current supply of an asset |
| `time_elapsed` | [google.protobuf.Duration](#google.protobuf.Duration) |  | time_elapsed represents the time elapsed |






<a name="gert.bep3.v1beta1.AtomicSwapResponse"></a>

### AtomicSwapResponse
AtomicSwapResponse represents the returned atomic swap properties


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [string](#string) |  | id represents the id of the atomic swap |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | amount represents the amount being swapped |
| `random_number_hash` | [string](#string) |  | random_number_hash represents the hash of the random number |
| `expire_height` | [uint64](#uint64) |  | expire_height represents the height when the swap expires |
| `timestamp` | [int64](#int64) |  | timestamp represents the timestamp of the swap |
| `sender` | [string](#string) |  | sender is the gert chain sender of the swap |
| `recipient` | [string](#string) |  | recipient is the gert chain recipient of the swap |
| `sender_other_chain` | [string](#string) |  | sender_other_chain is the sender on the other chain |
| `recipient_other_chain` | [string](#string) |  | recipient_other_chain is the recipient on the other chain |
| `closed_block` | [int64](#int64) |  | closed_block is the block when the swap is closed |
| `status` | [SwapStatus](#gert.bep3.v1beta1.SwapStatus) |  | status represents the current status of the swap |
| `cross_chain` | [bool](#bool) |  | cross_chain identifies whether the atomic swap is cross chain |
| `direction` | [SwapDirection](#gert.bep3.v1beta1.SwapDirection) |  | direction identifies if the swap is incoming or outgoing |






<a name="gert.bep3.v1beta1.QueryAssetSuppliesRequest"></a>

### QueryAssetSuppliesRequest
QueryAssetSuppliesRequest is the request type for the Query/AssetSupplies RPC method.






<a name="gert.bep3.v1beta1.QueryAssetSuppliesResponse"></a>

### QueryAssetSuppliesResponse
QueryAssetSuppliesResponse is the response type for the Query/AssetSupplies RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `asset_supplies` | [AssetSupplyResponse](#gert.bep3.v1beta1.AssetSupplyResponse) | repeated | asset_supplies represents the supplies of returned assets |






<a name="gert.bep3.v1beta1.QueryAssetSupplyRequest"></a>

### QueryAssetSupplyRequest
QueryAssetSupplyRequest is the request type for the Query/AssetSupply RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  | denom filters the asset response for the specified denom |






<a name="gert.bep3.v1beta1.QueryAssetSupplyResponse"></a>

### QueryAssetSupplyResponse
QueryAssetSupplyResponse is the response type for the Query/AssetSupply RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `asset_supply` | [AssetSupplyResponse](#gert.bep3.v1beta1.AssetSupplyResponse) |  | asset_supply represents the supply of the asset |






<a name="gert.bep3.v1beta1.QueryAtomicSwapRequest"></a>

### QueryAtomicSwapRequest
QueryAtomicSwapRequest is the request type for the Query/AtomicSwap RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `swap_id` | [string](#string) |  | swap_id represents the id of the swap to query |






<a name="gert.bep3.v1beta1.QueryAtomicSwapResponse"></a>

### QueryAtomicSwapResponse
QueryAtomicSwapResponse is the response type for the Query/AtomicSwap RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `atomic_swap` | [AtomicSwapResponse](#gert.bep3.v1beta1.AtomicSwapResponse) |  |  |






<a name="gert.bep3.v1beta1.QueryAtomicSwapsRequest"></a>

### QueryAtomicSwapsRequest
QueryAtomicSwapsRequest is the request type for the Query/AtomicSwaps RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `involve` | [string](#string) |  | involve filters by address |
| `expiration` | [uint64](#uint64) |  | expiration filters by expiration block height |
| `status` | [SwapStatus](#gert.bep3.v1beta1.SwapStatus) |  | status filters by swap status |
| `direction` | [SwapDirection](#gert.bep3.v1beta1.SwapDirection) |  | direction fitlers by swap direction |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="gert.bep3.v1beta1.QueryAtomicSwapsResponse"></a>

### QueryAtomicSwapsResponse
QueryAtomicSwapsResponse is the response type for the Query/AtomicSwaps RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `atomic_swaps` | [AtomicSwapResponse](#gert.bep3.v1beta1.AtomicSwapResponse) | repeated | atomic_swap represents the returned atomic swaps for the request |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="gert.bep3.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest defines the request type for querying x/bep3 parameters.






<a name="gert.bep3.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse defines the response type for querying x/bep3 parameters.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#gert.bep3.v1beta1.Params) |  | params represents the parameters of the module |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="gert.bep3.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service for bep3 module

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#gert.bep3.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#gert.bep3.v1beta1.QueryParamsResponse) | Params queries module params | GET|/gert/bep3/v1beta1/params|
| `AssetSupply` | [QueryAssetSupplyRequest](#gert.bep3.v1beta1.QueryAssetSupplyRequest) | [QueryAssetSupplyResponse](#gert.bep3.v1beta1.QueryAssetSupplyResponse) | AssetSupply queries info about an asset's supply | GET|/gert/bep3/v1beta1/assetsupply/{denom}|
| `AssetSupplies` | [QueryAssetSuppliesRequest](#gert.bep3.v1beta1.QueryAssetSuppliesRequest) | [QueryAssetSuppliesResponse](#gert.bep3.v1beta1.QueryAssetSuppliesResponse) | AssetSupplies queries a list of asset supplies | GET|/gert/bep3/v1beta1/assetsupplies|
| `AtomicSwap` | [QueryAtomicSwapRequest](#gert.bep3.v1beta1.QueryAtomicSwapRequest) | [QueryAtomicSwapResponse](#gert.bep3.v1beta1.QueryAtomicSwapResponse) | AtomicSwap queries info about an atomic swap | GET|/gert/bep3/v1beta1/atomicswap/{swap_id}|
| `AtomicSwaps` | [QueryAtomicSwapsRequest](#gert.bep3.v1beta1.QueryAtomicSwapsRequest) | [QueryAtomicSwapsResponse](#gert.bep3.v1beta1.QueryAtomicSwapsResponse) | AtomicSwaps queries a list of atomic swaps | GET|/gert/bep3/v1beta1/atomicswaps|

 <!-- end services -->



<a name="gert/bep3/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gert/bep3/v1beta1/tx.proto



<a name="gert.bep3.v1beta1.MsgClaimAtomicSwap"></a>

### MsgClaimAtomicSwap
MsgClaimAtomicSwap defines the Msg/ClaimAtomicSwap request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `from` | [string](#string) |  |  |
| `swap_id` | [string](#string) |  |  |
| `random_number` | [string](#string) |  |  |






<a name="gert.bep3.v1beta1.MsgClaimAtomicSwapResponse"></a>

### MsgClaimAtomicSwapResponse
MsgClaimAtomicSwapResponse defines the Msg/ClaimAtomicSwap response type.






<a name="gert.bep3.v1beta1.MsgCreateAtomicSwap"></a>

### MsgCreateAtomicSwap
MsgCreateAtomicSwap defines the Msg/CreateAtomicSwap request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `from` | [string](#string) |  |  |
| `to` | [string](#string) |  |  |
| `recipient_other_chain` | [string](#string) |  |  |
| `sender_other_chain` | [string](#string) |  |  |
| `random_number_hash` | [string](#string) |  |  |
| `timestamp` | [int64](#int64) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| `height_span` | [uint64](#uint64) |  |  |






<a name="gert.bep3.v1beta1.MsgCreateAtomicSwapResponse"></a>

### MsgCreateAtomicSwapResponse
MsgCreateAtomicSwapResponse defines the Msg/CreateAtomicSwap response type.






<a name="gert.bep3.v1beta1.MsgRefundAtomicSwap"></a>

### MsgRefundAtomicSwap
MsgRefundAtomicSwap defines the Msg/RefundAtomicSwap request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `from` | [string](#string) |  |  |
| `swap_id` | [string](#string) |  |  |






<a name="gert.bep3.v1beta1.MsgRefundAtomicSwapResponse"></a>

### MsgRefundAtomicSwapResponse
MsgRefundAtomicSwapResponse defines the Msg/RefundAtomicSwap response type.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="gert.bep3.v1beta1.Msg"></a>

### Msg
Msg defines the bep3 Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `CreateAtomicSwap` | [MsgCreateAtomicSwap](#gert.bep3.v1beta1.MsgCreateAtomicSwap) | [MsgCreateAtomicSwapResponse](#gert.bep3.v1beta1.MsgCreateAtomicSwapResponse) | CreateAtomicSwap defines a method for creating an atomic swap | |
| `ClaimAtomicSwap` | [MsgClaimAtomicSwap](#gert.bep3.v1beta1.MsgClaimAtomicSwap) | [MsgClaimAtomicSwapResponse](#gert.bep3.v1beta1.MsgClaimAtomicSwapResponse) | ClaimAtomicSwap defines a method for claiming an atomic swap | |
| `RefundAtomicSwap` | [MsgRefundAtomicSwap](#gert.bep3.v1beta1.MsgRefundAtomicSwap) | [MsgRefundAtomicSwapResponse](#gert.bep3.v1beta1.MsgRefundAtomicSwapResponse) | RefundAtomicSwap defines a method for refunding an atomic swap | |

 <!-- end services -->



<a name="gert/cdp/v1beta1/cdp.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gert/cdp/v1beta1/cdp.proto



<a name="gert.cdp.v1beta1.CDP"></a>

### CDP
CDP defines the state of a single collateralized debt position.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [uint64](#uint64) |  |  |
| `owner` | [bytes](#bytes) |  |  |
| `type` | [string](#string) |  |  |
| `collateral` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `principal` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `accumulated_fees` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `fees_updated` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `interest_factor` | [string](#string) |  |  |






<a name="gert.cdp.v1beta1.Deposit"></a>

### Deposit
Deposit defines an amount of coins deposited by an account to a cdp


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `cdp_id` | [uint64](#uint64) |  |  |
| `depositor` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="gert.cdp.v1beta1.OwnerCDPIndex"></a>

### OwnerCDPIndex
OwnerCDPIndex defines the cdp ids for a single cdp owner


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `cdp_ids` | [uint64](#uint64) | repeated |  |






<a name="gert.cdp.v1beta1.TotalCollateral"></a>

### TotalCollateral
TotalCollateral defines the total collateral of a given collateral type


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_type` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="gert.cdp.v1beta1.TotalPrincipal"></a>

### TotalPrincipal
TotalPrincipal defines the total principal of a given collateral type


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_type` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="gert/cdp/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gert/cdp/v1beta1/genesis.proto



<a name="gert.cdp.v1beta1.CollateralParam"></a>

### CollateralParam
CollateralParam defines governance parameters for each collateral type within the cdp module


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `type` | [string](#string) |  |  |
| `liquidation_ratio` | [string](#string) |  |  |
| `debt_limit` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `stability_fee` | [string](#string) |  |  |
| `auction_size` | [string](#string) |  |  |
| `liquidation_penalty` | [string](#string) |  |  |
| `spot_market_id` | [string](#string) |  |  |
| `liquidation_market_id` | [string](#string) |  |  |
| `keeper_reward_percentage` | [string](#string) |  |  |
| `check_collateralization_index_count` | [string](#string) |  |  |
| `conversion_factor` | [string](#string) |  |  |






<a name="gert.cdp.v1beta1.DebtParam"></a>

### DebtParam
DebtParam defines governance params for debt assets


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `reference_asset` | [string](#string) |  |  |
| `conversion_factor` | [string](#string) |  |  |
| `debt_floor` | [string](#string) |  |  |






<a name="gert.cdp.v1beta1.GenesisAccumulationTime"></a>

### GenesisAccumulationTime
GenesisAccumulationTime defines the previous distribution time and its corresponding denom


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_type` | [string](#string) |  |  |
| `previous_accumulation_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `interest_factor` | [string](#string) |  |  |






<a name="gert.cdp.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the cdp module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#gert.cdp.v1beta1.Params) |  | params defines all the paramaters of the module. |
| `cdps` | [CDP](#gert.cdp.v1beta1.CDP) | repeated |  |
| `deposits` | [Deposit](#gert.cdp.v1beta1.Deposit) | repeated |  |
| `starting_cdp_id` | [uint64](#uint64) |  |  |
| `debt_denom` | [string](#string) |  |  |
| `gov_denom` | [string](#string) |  |  |
| `previous_accumulation_times` | [GenesisAccumulationTime](#gert.cdp.v1beta1.GenesisAccumulationTime) | repeated |  |
| `total_principals` | [GenesisTotalPrincipal](#gert.cdp.v1beta1.GenesisTotalPrincipal) | repeated |  |






<a name="gert.cdp.v1beta1.GenesisTotalPrincipal"></a>

### GenesisTotalPrincipal
GenesisTotalPrincipal defines the total principal and its corresponding collateral type


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_type` | [string](#string) |  |  |
| `total_principal` | [string](#string) |  |  |






<a name="gert.cdp.v1beta1.Params"></a>

### Params
Params defines the parameters for the cdp module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_params` | [CollateralParam](#gert.cdp.v1beta1.CollateralParam) | repeated |  |
| `debt_param` | [DebtParam](#gert.cdp.v1beta1.DebtParam) |  |  |
| `global_debt_limit` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `surplus_auction_threshold` | [string](#string) |  |  |
| `surplus_auction_lot` | [string](#string) |  |  |
| `debt_auction_threshold` | [string](#string) |  |  |
| `debt_auction_lot` | [string](#string) |  |  |
| `circuit_breaker` | [bool](#bool) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="gert/cdp/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gert/cdp/v1beta1/query.proto



<a name="gert.cdp.v1beta1.CDPResponse"></a>

### CDPResponse
CDPResponse defines the state of a single collateralized debt position.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [uint64](#uint64) |  |  |
| `owner` | [string](#string) |  |  |
| `type` | [string](#string) |  |  |
| `collateral` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `principal` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `accumulated_fees` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `fees_updated` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `interest_factor` | [string](#string) |  |  |
| `collateral_value` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `collateralization_ratio` | [string](#string) |  |  |






<a name="gert.cdp.v1beta1.QueryAccountsRequest"></a>

### QueryAccountsRequest
QueryAccountsRequest defines the request type for the Query/Accounts RPC method.






<a name="gert.cdp.v1beta1.QueryAccountsResponse"></a>

### QueryAccountsResponse
QueryAccountsResponse defines the response type for the Query/Accounts RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `accounts` | [cosmos.auth.v1beta1.ModuleAccount](#cosmos.auth.v1beta1.ModuleAccount) | repeated |  |






<a name="gert.cdp.v1beta1.QueryCdpRequest"></a>

### QueryCdpRequest
QueryCdpRequest defines the request type for the Query/Cdp RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_type` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |






<a name="gert.cdp.v1beta1.QueryCdpResponse"></a>

### QueryCdpResponse
QueryCdpResponse defines the response type for the Query/Cdp RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `cdp` | [CDPResponse](#gert.cdp.v1beta1.CDPResponse) |  |  |






<a name="gert.cdp.v1beta1.QueryCdpsRequest"></a>

### QueryCdpsRequest
QueryCdpsRequest is the params for a filtered CDP query, the request type for the Query/Cdps RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_type` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |
| `id` | [uint64](#uint64) |  |  |
| `ratio` | [string](#string) |  | sdk.Dec as a string |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="gert.cdp.v1beta1.QueryCdpsResponse"></a>

### QueryCdpsResponse
QueryCdpsResponse defines the response type for the Query/Cdps RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `cdps` | [CDPResponse](#gert.cdp.v1beta1.CDPResponse) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="gert.cdp.v1beta1.QueryDepositsRequest"></a>

### QueryDepositsRequest
QueryDepositsRequest defines the request type for the Query/Deposits RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_type` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |






<a name="gert.cdp.v1beta1.QueryDepositsResponse"></a>

### QueryDepositsResponse
QueryDepositsResponse defines the response type for the Query/Deposits RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `deposits` | [Deposit](#gert.cdp.v1beta1.Deposit) | repeated |  |






<a name="gert.cdp.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest defines the request type for the Query/Params RPC method.






<a name="gert.cdp.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse defines the response type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#gert.cdp.v1beta1.Params) |  |  |






<a name="gert.cdp.v1beta1.QueryTotalCollateralRequest"></a>

### QueryTotalCollateralRequest
QueryTotalCollateralRequest defines the request type for the Query/TotalCollateral RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_type` | [string](#string) |  |  |






<a name="gert.cdp.v1beta1.QueryTotalCollateralResponse"></a>

### QueryTotalCollateralResponse
QueryTotalCollateralResponse defines the response type for the Query/TotalCollateral RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `total_collateral` | [TotalCollateral](#gert.cdp.v1beta1.TotalCollateral) | repeated |  |






<a name="gert.cdp.v1beta1.QueryTotalPrincipalRequest"></a>

### QueryTotalPrincipalRequest
QueryTotalPrincipalRequest defines the request type for the Query/TotalPrincipal RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_type` | [string](#string) |  |  |






<a name="gert.cdp.v1beta1.QueryTotalPrincipalResponse"></a>

### QueryTotalPrincipalResponse
QueryTotalPrincipalResponse defines the response type for the Query/TotalPrincipal RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `total_principal` | [TotalPrincipal](#gert.cdp.v1beta1.TotalPrincipal) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="gert.cdp.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service for cdp module

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#gert.cdp.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#gert.cdp.v1beta1.QueryParamsResponse) | Params queries all parameters of the cdp module. | GET|/gert/cdp/v1beta1/params|
| `Accounts` | [QueryAccountsRequest](#gert.cdp.v1beta1.QueryAccountsRequest) | [QueryAccountsResponse](#gert.cdp.v1beta1.QueryAccountsResponse) | Accounts queries the CDP module accounts. | GET|/gert/cdp/v1beta1/accounts|
| `TotalPrincipal` | [QueryTotalPrincipalRequest](#gert.cdp.v1beta1.QueryTotalPrincipalRequest) | [QueryTotalPrincipalResponse](#gert.cdp.v1beta1.QueryTotalPrincipalResponse) | TotalPrincipal queries the total principal of a given collateral type. | GET|/gert/cdp/v1beta1/totalPrincipal|
| `TotalCollateral` | [QueryTotalCollateralRequest](#gert.cdp.v1beta1.QueryTotalCollateralRequest) | [QueryTotalCollateralResponse](#gert.cdp.v1beta1.QueryTotalCollateralResponse) | TotalCollateral queries the total collateral of a given collateral type. | GET|/gert/cdp/v1beta1/totalCollateral|
| `Cdps` | [QueryCdpsRequest](#gert.cdp.v1beta1.QueryCdpsRequest) | [QueryCdpsResponse](#gert.cdp.v1beta1.QueryCdpsResponse) | Cdps queries all active CDPs. | GET|/gert/cdp/v1beta1/cdps|
| `Cdp` | [QueryCdpRequest](#gert.cdp.v1beta1.QueryCdpRequest) | [QueryCdpResponse](#gert.cdp.v1beta1.QueryCdpResponse) | Cdp queries a CDP with the input owner address and collateral type. | GET|/gert/cdp/v1beta1/cdps/{owner}/{collateral_type}|
| `Deposits` | [QueryDepositsRequest](#gert.cdp.v1beta1.QueryDepositsRequest) | [QueryDepositsResponse](#gert.cdp.v1beta1.QueryDepositsResponse) | Deposits queries deposits associated with the CDP owned by an address for a collateral type. | GET|/gert/cdp/v1beta1/cdps/deposits/{owner}/{collateral_type}|

 <!-- end services -->



<a name="gert/cdp/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gert/cdp/v1beta1/tx.proto



<a name="gert.cdp.v1beta1.MsgCreateCDP"></a>

### MsgCreateCDP
MsgCreateCDP defines a message to create a new CDP.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `collateral` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `principal` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `collateral_type` | [string](#string) |  |  |






<a name="gert.cdp.v1beta1.MsgCreateCDPResponse"></a>

### MsgCreateCDPResponse
MsgCreateCDPResponse defines the Msg/CreateCDP response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `cdp_id` | [uint64](#uint64) |  |  |






<a name="gert.cdp.v1beta1.MsgDeposit"></a>

### MsgDeposit
MsgDeposit defines a message to deposit to a CDP.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `depositor` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |
| `collateral` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `collateral_type` | [string](#string) |  |  |






<a name="gert.cdp.v1beta1.MsgDepositResponse"></a>

### MsgDepositResponse
MsgDepositResponse defines the Msg/Deposit response type.






<a name="gert.cdp.v1beta1.MsgDrawDebt"></a>

### MsgDrawDebt
MsgDrawDebt defines a message to draw debt from a CDP.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `collateral_type` | [string](#string) |  |  |
| `principal` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="gert.cdp.v1beta1.MsgDrawDebtResponse"></a>

### MsgDrawDebtResponse
MsgDrawDebtResponse defines the Msg/DrawDebt response type.






<a name="gert.cdp.v1beta1.MsgLiquidate"></a>

### MsgLiquidate
MsgLiquidate defines a message to attempt to liquidate a CDP whos
collateralization ratio is under its liquidation ratio.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `keeper` | [string](#string) |  |  |
| `borrower` | [string](#string) |  |  |
| `collateral_type` | [string](#string) |  |  |






<a name="gert.cdp.v1beta1.MsgLiquidateResponse"></a>

### MsgLiquidateResponse
MsgLiquidateResponse defines the Msg/Liquidate response type.






<a name="gert.cdp.v1beta1.MsgRepayDebt"></a>

### MsgRepayDebt
MsgRepayDebt defines a message to repay debt from a CDP.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `collateral_type` | [string](#string) |  |  |
| `payment` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="gert.cdp.v1beta1.MsgRepayDebtResponse"></a>

### MsgRepayDebtResponse
MsgRepayDebtResponse defines the Msg/RepayDebt response type.






<a name="gert.cdp.v1beta1.MsgWithdraw"></a>

### MsgWithdraw
MsgWithdraw defines a message to withdraw collateral from a CDP.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `depositor` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |
| `collateral` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `collateral_type` | [string](#string) |  |  |






<a name="gert.cdp.v1beta1.MsgWithdrawResponse"></a>

### MsgWithdrawResponse
MsgWithdrawResponse defines the Msg/Withdraw response type.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="gert.cdp.v1beta1.Msg"></a>

### Msg
Msg defines the cdp Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `CreateCDP` | [MsgCreateCDP](#gert.cdp.v1beta1.MsgCreateCDP) | [MsgCreateCDPResponse](#gert.cdp.v1beta1.MsgCreateCDPResponse) | CreateCDP defines a method to create a new CDP. | |
| `Deposit` | [MsgDeposit](#gert.cdp.v1beta1.MsgDeposit) | [MsgDepositResponse](#gert.cdp.v1beta1.MsgDepositResponse) | Deposit defines a method to deposit to a CDP. | |
| `Withdraw` | [MsgWithdraw](#gert.cdp.v1beta1.MsgWithdraw) | [MsgWithdrawResponse](#gert.cdp.v1beta1.MsgWithdrawResponse) | Withdraw defines a method to withdraw collateral from a CDP. | |
| `DrawDebt` | [MsgDrawDebt](#gert.cdp.v1beta1.MsgDrawDebt) | [MsgDrawDebtResponse](#gert.cdp.v1beta1.MsgDrawDebtResponse) | DrawDebt defines a method to draw debt from a CDP. | |
| `RepayDebt` | [MsgRepayDebt](#gert.cdp.v1beta1.MsgRepayDebt) | [MsgRepayDebtResponse](#gert.cdp.v1beta1.MsgRepayDebtResponse) | RepayDebt defines a method to repay debt from a CDP. | |
| `Liquidate` | [MsgLiquidate](#gert.cdp.v1beta1.MsgLiquidate) | [MsgLiquidateResponse](#gert.cdp.v1beta1.MsgLiquidateResponse) | Liquidate defines a method to attempt to liquidate a CDP whos collateralization ratio is under its liquidation ratio. | |

 <!-- end services -->



<a name="gert/committee/v1beta1/committee.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gert/committee/v1beta1/committee.proto



<a name="gert.committee.v1beta1.BaseCommittee"></a>

### BaseCommittee
BaseCommittee is a common type shared by all Committees


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [uint64](#uint64) |  |  |
| `description` | [string](#string) |  |  |
| `members` | [bytes](#bytes) | repeated |  |
| `permissions` | [google.protobuf.Any](#google.protobuf.Any) | repeated |  |
| `vote_threshold` | [string](#string) |  | Smallest percentage that must vote for a proposal to pass |
| `proposal_duration` | [google.protobuf.Duration](#google.protobuf.Duration) |  | The length of time a proposal remains active for. Proposals will close earlier if they get enough votes. |
| `tally_option` | [TallyOption](#gert.committee.v1beta1.TallyOption) |  |  |






<a name="gert.committee.v1beta1.MemberCommittee"></a>

### MemberCommittee
MemberCommittee is an alias of BaseCommittee


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_committee` | [BaseCommittee](#gert.committee.v1beta1.BaseCommittee) |  |  |






<a name="gert.committee.v1beta1.TokenCommittee"></a>

### TokenCommittee
TokenCommittee supports voting on proposals by token holders


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_committee` | [BaseCommittee](#gert.committee.v1beta1.BaseCommittee) |  |  |
| `quorum` | [string](#string) |  |  |
| `tally_denom` | [string](#string) |  |  |





 <!-- end messages -->


<a name="gert.committee.v1beta1.TallyOption"></a>

### TallyOption
TallyOption enumerates the valid types of a tally.

| Name | Number | Description |
| ---- | ------ | ----------- |
| TALLY_OPTION_UNSPECIFIED | 0 | TALLY_OPTION_UNSPECIFIED defines a null tally option. |
| TALLY_OPTION_FIRST_PAST_THE_POST | 1 | Votes are tallied each block and the proposal passes as soon as the vote threshold is reached |
| TALLY_OPTION_DEADLINE | 2 | Votes are tallied exactly once, when the deadline time is reached |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="gert/committee/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gert/committee/v1beta1/genesis.proto



<a name="gert.committee.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the committee module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `next_proposal_id` | [uint64](#uint64) |  |  |
| `committees` | [google.protobuf.Any](#google.protobuf.Any) | repeated |  |
| `proposals` | [Proposal](#gert.committee.v1beta1.Proposal) | repeated |  |
| `votes` | [Vote](#gert.committee.v1beta1.Vote) | repeated |  |






<a name="gert.committee.v1beta1.Proposal"></a>

### Proposal
Proposal is an internal record of a governance proposal submitted to a committee.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `content` | [google.protobuf.Any](#google.protobuf.Any) |  |  |
| `id` | [uint64](#uint64) |  |  |
| `committee_id` | [uint64](#uint64) |  |  |
| `deadline` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="gert.committee.v1beta1.Vote"></a>

### Vote
Vote is an internal record of a single governance vote.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  |  |
| `voter` | [bytes](#bytes) |  |  |
| `vote_type` | [VoteType](#gert.committee.v1beta1.VoteType) |  |  |





 <!-- end messages -->


<a name="gert.committee.v1beta1.VoteType"></a>

### VoteType
VoteType enumerates the valid types of a vote.

| Name | Number | Description |
| ---- | ------ | ----------- |
| VOTE_TYPE_UNSPECIFIED | 0 | VOTE_TYPE_UNSPECIFIED defines a no-op vote option. |
| VOTE_TYPE_YES | 1 | VOTE_TYPE_YES defines a yes vote option. |
| VOTE_TYPE_NO | 2 | VOTE_TYPE_NO defines a no vote option. |
| VOTE_TYPE_ABSTAIN | 3 | VOTE_TYPE_ABSTAIN defines an abstain vote option. |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="gert/committee/v1beta1/permissions.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gert/committee/v1beta1/permissions.proto



<a name="gert.committee.v1beta1.AllowedParamsChange"></a>

### AllowedParamsChange
AllowedParamsChange contains data on the allowed parameter changes for subspace, key, and sub params requirements.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `subspace` | [string](#string) |  |  |
| `key` | [string](#string) |  |  |
| `single_subparam_allowed_attrs` | [string](#string) | repeated | Requirements for when the subparam value is a single record. This contains list of allowed attribute keys that can be changed on the subparam record. |
| `multi_subparams_requirements` | [SubparamRequirement](#gert.committee.v1beta1.SubparamRequirement) | repeated | Requirements for when the subparam value is a list of records. The requirements contains requirements for each record in the list. |






<a name="gert.committee.v1beta1.GodPermission"></a>

### GodPermission
GodPermission allows any governance proposal. It is used mainly for testing.






<a name="gert.committee.v1beta1.ParamsChangePermission"></a>

### ParamsChangePermission
ParamsChangePermission allows any parameter or sub parameter change proposal.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `allowed_params_changes` | [AllowedParamsChange](#gert.committee.v1beta1.AllowedParamsChange) | repeated |  |






<a name="gert.committee.v1beta1.SoftwareUpgradePermission"></a>

### SoftwareUpgradePermission
SoftwareUpgradePermission permission type for software upgrade proposals






<a name="gert.committee.v1beta1.SubparamRequirement"></a>

### SubparamRequirement
SubparamRequirement contains requirements for a single record in a subparam value list


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key` | [string](#string) |  | The required attr key of the param record. |
| `val` | [string](#string) |  | The required param value for the param record key. The key and value is used to match to the target param record. |
| `allowed_subparam_attr_changes` | [string](#string) | repeated | The sub param attrs that are allowed to be changed. |






<a name="gert.committee.v1beta1.TextPermission"></a>

### TextPermission
TextPermission allows any text governance proposal.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="gert/committee/v1beta1/proposal.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gert/committee/v1beta1/proposal.proto



<a name="gert.committee.v1beta1.CommitteeChangeProposal"></a>

### CommitteeChangeProposal
CommitteeChangeProposal is a gov proposal for creating a new committee or modifying an existing one.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  |  |
| `description` | [string](#string) |  |  |
| `new_committee` | [google.protobuf.Any](#google.protobuf.Any) |  |  |






<a name="gert.committee.v1beta1.CommitteeDeleteProposal"></a>

### CommitteeDeleteProposal
CommitteeDeleteProposal is a gov proposal for removing a committee.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  |  |
| `description` | [string](#string) |  |  |
| `committee_id` | [uint64](#uint64) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="gert/committee/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gert/committee/v1beta1/query.proto



<a name="gert.committee.v1beta1.QueryCommitteeRequest"></a>

### QueryCommitteeRequest
QueryCommitteeRequest defines the request type for querying x/committee committee.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `committee_id` | [uint64](#uint64) |  |  |






<a name="gert.committee.v1beta1.QueryCommitteeResponse"></a>

### QueryCommitteeResponse
QueryCommitteeResponse defines the response type for querying x/committee committee.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `committee` | [google.protobuf.Any](#google.protobuf.Any) |  |  |






<a name="gert.committee.v1beta1.QueryCommitteesRequest"></a>

### QueryCommitteesRequest
QueryCommitteesRequest defines the request type for querying x/committee committees.






<a name="gert.committee.v1beta1.QueryCommitteesResponse"></a>

### QueryCommitteesResponse
QueryCommitteesResponse defines the response type for querying x/committee committees.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `committees` | [google.protobuf.Any](#google.protobuf.Any) | repeated |  |






<a name="gert.committee.v1beta1.QueryNextProposalIDRequest"></a>

### QueryNextProposalIDRequest
QueryNextProposalIDRequest defines the request type for querying x/committee NextProposalID.






<a name="gert.committee.v1beta1.QueryNextProposalIDResponse"></a>

### QueryNextProposalIDResponse
QueryNextProposalIDRequest defines the response type for querying x/committee NextProposalID.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `next_proposal_id` | [uint64](#uint64) |  |  |






<a name="gert.committee.v1beta1.QueryProposalRequest"></a>

### QueryProposalRequest
QueryProposalRequest defines the request type for querying x/committee proposal.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  |  |






<a name="gert.committee.v1beta1.QueryProposalResponse"></a>

### QueryProposalResponse
QueryProposalResponse defines the response type for querying x/committee proposal.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pub_proposal` | [google.protobuf.Any](#google.protobuf.Any) |  |  |
| `id` | [uint64](#uint64) |  |  |
| `committee_id` | [uint64](#uint64) |  |  |
| `deadline` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="gert.committee.v1beta1.QueryProposalsRequest"></a>

### QueryProposalsRequest
QueryProposalsRequest defines the request type for querying x/committee proposals.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `committee_id` | [uint64](#uint64) |  |  |






<a name="gert.committee.v1beta1.QueryProposalsResponse"></a>

### QueryProposalsResponse
QueryProposalsResponse defines the response type for querying x/committee proposals.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposals` | [QueryProposalResponse](#gert.committee.v1beta1.QueryProposalResponse) | repeated |  |






<a name="gert.committee.v1beta1.QueryRawParamsRequest"></a>

### QueryRawParamsRequest
QueryRawParamsRequest defines the request type for querying x/committee raw params.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `subspace` | [string](#string) |  |  |
| `key` | [string](#string) |  |  |






<a name="gert.committee.v1beta1.QueryRawParamsResponse"></a>

### QueryRawParamsResponse
QueryRawParamsResponse defines the response type for querying x/committee raw params.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `raw_data` | [string](#string) |  |  |






<a name="gert.committee.v1beta1.QueryTallyRequest"></a>

### QueryTallyRequest
QueryTallyRequest defines the request type for querying x/committee tally.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  |  |






<a name="gert.committee.v1beta1.QueryTallyResponse"></a>

### QueryTallyResponse
QueryTallyResponse defines the response type for querying x/committee tally.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  |  |
| `yes_votes` | [string](#string) |  |  |
| `no_votes` | [string](#string) |  |  |
| `current_votes` | [string](#string) |  |  |
| `possible_votes` | [string](#string) |  |  |
| `vote_threshold` | [string](#string) |  |  |
| `quorum` | [string](#string) |  |  |






<a name="gert.committee.v1beta1.QueryVoteRequest"></a>

### QueryVoteRequest
QueryVoteRequest defines the request type for querying x/committee vote.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  |  |
| `voter` | [string](#string) |  |  |






<a name="gert.committee.v1beta1.QueryVoteResponse"></a>

### QueryVoteResponse
QueryVoteResponse defines the response type for querying x/committee vote.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  |  |
| `voter` | [string](#string) |  |  |
| `vote_type` | [VoteType](#gert.committee.v1beta1.VoteType) |  |  |






<a name="gert.committee.v1beta1.QueryVotesRequest"></a>

### QueryVotesRequest
QueryVotesRequest defines the request type for querying x/committee votes.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  |  |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="gert.committee.v1beta1.QueryVotesResponse"></a>

### QueryVotesResponse
QueryVotesResponse defines the response type for querying x/committee votes.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `votes` | [QueryVoteResponse](#gert.committee.v1beta1.QueryVoteResponse) | repeated | votes defined the queried votes. |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="gert.committee.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service for committee module

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Committees` | [QueryCommitteesRequest](#gert.committee.v1beta1.QueryCommitteesRequest) | [QueryCommitteesResponse](#gert.committee.v1beta1.QueryCommitteesResponse) | Committees queries all committess of the committee module. | GET|/gert/committee/v1beta1/committees|
| `Committee` | [QueryCommitteeRequest](#gert.committee.v1beta1.QueryCommitteeRequest) | [QueryCommitteeResponse](#gert.committee.v1beta1.QueryCommitteeResponse) | Committee queries a committee based on committee ID. | GET|/gert/committee/v1beta1/committees/{committee_id}|
| `Proposals` | [QueryProposalsRequest](#gert.committee.v1beta1.QueryProposalsRequest) | [QueryProposalsResponse](#gert.committee.v1beta1.QueryProposalsResponse) | Proposals queries proposals based on committee ID. | GET|/gert/committee/v1beta1/proposals|
| `Proposal` | [QueryProposalRequest](#gert.committee.v1beta1.QueryProposalRequest) | [QueryProposalResponse](#gert.committee.v1beta1.QueryProposalResponse) | Deposits queries a proposal based on proposal ID. | GET|/gert/committee/v1beta1/proposals/{proposal_id}|
| `NextProposalID` | [QueryNextProposalIDRequest](#gert.committee.v1beta1.QueryNextProposalIDRequest) | [QueryNextProposalIDResponse](#gert.committee.v1beta1.QueryNextProposalIDResponse) | NextProposalID queries the next proposal ID of the committee module. | GET|/gert/committee/v1beta1/next-proposal-id|
| `Votes` | [QueryVotesRequest](#gert.committee.v1beta1.QueryVotesRequest) | [QueryVotesResponse](#gert.committee.v1beta1.QueryVotesResponse) | Votes queries all votes for a single proposal ID. | GET|/gert/committee/v1beta1/proposals/{proposal_id}/votes|
| `Vote` | [QueryVoteRequest](#gert.committee.v1beta1.QueryVoteRequest) | [QueryVoteResponse](#gert.committee.v1beta1.QueryVoteResponse) | Vote queries the vote of a single voter for a single proposal ID. | GET|/gert/committee/v1beta1/proposals/{proposal_id}/votes/{voter}|
| `Tally` | [QueryTallyRequest](#gert.committee.v1beta1.QueryTallyRequest) | [QueryTallyResponse](#gert.committee.v1beta1.QueryTallyResponse) | Tally queries the tally of a single proposal ID. | GET|/gert/committee/v1beta1/proposals/{proposal_id}/tally|
| `RawParams` | [QueryRawParamsRequest](#gert.committee.v1beta1.QueryRawParamsRequest) | [QueryRawParamsResponse](#gert.committee.v1beta1.QueryRawParamsResponse) | RawParams queries the raw params data of any subspace and key. | GET|/gert/committee/v1beta1/raw-params|

 <!-- end services -->



<a name="gert/committee/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gert/committee/v1beta1/tx.proto



<a name="gert.committee.v1beta1.MsgSubmitProposal"></a>

### MsgSubmitProposal
MsgSubmitProposal is used by committee members to create a new proposal that they can vote on.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pub_proposal` | [google.protobuf.Any](#google.protobuf.Any) |  |  |
| `proposer` | [string](#string) |  |  |
| `committee_id` | [uint64](#uint64) |  |  |






<a name="gert.committee.v1beta1.MsgSubmitProposalResponse"></a>

### MsgSubmitProposalResponse
MsgSubmitProposalResponse defines the SubmitProposal response type


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  |  |






<a name="gert.committee.v1beta1.MsgVote"></a>

### MsgVote
MsgVote is submitted by committee members to vote on proposals.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  |  |
| `voter` | [string](#string) |  |  |
| `vote_type` | [VoteType](#gert.committee.v1beta1.VoteType) |  |  |






<a name="gert.committee.v1beta1.MsgVoteResponse"></a>

### MsgVoteResponse
MsgVoteResponse defines the Vote response type





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="gert.committee.v1beta1.Msg"></a>

### Msg
Msg defines the committee Msg service

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `SubmitProposal` | [MsgSubmitProposal](#gert.committee.v1beta1.MsgSubmitProposal) | [MsgSubmitProposalResponse](#gert.committee.v1beta1.MsgSubmitProposalResponse) | SubmitProposal defines a method for submitting a committee proposal | |
| `Vote` | [MsgVote](#gert.committee.v1beta1.MsgVote) | [MsgVoteResponse](#gert.committee.v1beta1.MsgVoteResponse) | Vote defines a method for voting on a proposal | |

 <!-- end services -->



<a name="gert/evmutil/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gert/evmutil/v1beta1/genesis.proto



<a name="gert.evmutil.v1beta1.Account"></a>

### Account
BalanceAccount defines an account in the evmutil module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [bytes](#bytes) |  |  |
| `balance` | [string](#string) |  | balance indicates the amount of agert owned by the address. |






<a name="gert.evmutil.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the evmutil module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `accounts` | [Account](#gert.evmutil.v1beta1.Account) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="gert/hard/v1beta1/hard.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gert/hard/v1beta1/hard.proto



<a name="gert.hard.v1beta1.Borrow"></a>

### Borrow
Borrow defines an amount of coins borrowed from a hard module account.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `borrower` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| `index` | [BorrowInterestFactor](#gert.hard.v1beta1.BorrowInterestFactor) | repeated |  |






<a name="gert.hard.v1beta1.BorrowInterestFactor"></a>

### BorrowInterestFactor
BorrowInterestFactor defines an individual borrow interest factor.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `value` | [string](#string) |  |  |






<a name="gert.hard.v1beta1.BorrowLimit"></a>

### BorrowLimit
BorrowLimit enforces restrictions on a money market.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `has_max_limit` | [bool](#bool) |  |  |
| `maximum_limit` | [string](#string) |  |  |
| `loan_to_value` | [string](#string) |  |  |






<a name="gert.hard.v1beta1.CoinsProto"></a>

### CoinsProto
CoinsProto defines a Protobuf wrapper around a Coins slice


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `coins` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="gert.hard.v1beta1.Deposit"></a>

### Deposit
Deposit defines an amount of coins deposited into a hard module account.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `depositor` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| `index` | [SupplyInterestFactor](#gert.hard.v1beta1.SupplyInterestFactor) | repeated |  |






<a name="gert.hard.v1beta1.InterestRateModel"></a>

### InterestRateModel
InterestRateModel contains information about an asset's interest rate.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_rate_apy` | [string](#string) |  |  |
| `base_multiplier` | [string](#string) |  |  |
| `kink` | [string](#string) |  |  |
| `jump_multiplier` | [string](#string) |  |  |






<a name="gert.hard.v1beta1.MoneyMarket"></a>

### MoneyMarket
MoneyMarket is a money market for an individual asset.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `borrow_limit` | [BorrowLimit](#gert.hard.v1beta1.BorrowLimit) |  |  |
| `spot_market_id` | [string](#string) |  |  |
| `conversion_factor` | [string](#string) |  |  |
| `interest_rate_model` | [InterestRateModel](#gert.hard.v1beta1.InterestRateModel) |  |  |
| `reserve_factor` | [string](#string) |  |  |
| `keeper_reward_percentage` | [string](#string) |  |  |






<a name="gert.hard.v1beta1.Params"></a>

### Params
Params defines the parameters for the hard module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `money_markets` | [MoneyMarket](#gert.hard.v1beta1.MoneyMarket) | repeated |  |
| `minimum_borrow_usd_value` | [string](#string) |  |  |






<a name="gert.hard.v1beta1.SupplyInterestFactor"></a>

### SupplyInterestFactor
SupplyInterestFactor defines an individual borrow interest factor.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `value` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="gert/hard/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gert/hard/v1beta1/genesis.proto



<a name="gert.hard.v1beta1.GenesisAccumulationTime"></a>

### GenesisAccumulationTime
GenesisAccumulationTime stores the previous distribution time and its corresponding denom.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_type` | [string](#string) |  |  |
| `previous_accumulation_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `supply_interest_factor` | [string](#string) |  |  |
| `borrow_interest_factor` | [string](#string) |  |  |






<a name="gert.hard.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the hard module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#gert.hard.v1beta1.Params) |  |  |
| `previous_accumulation_times` | [GenesisAccumulationTime](#gert.hard.v1beta1.GenesisAccumulationTime) | repeated |  |
| `deposits` | [Deposit](#gert.hard.v1beta1.Deposit) | repeated |  |
| `borrows` | [Borrow](#gert.hard.v1beta1.Borrow) | repeated |  |
| `total_supplied` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| `total_borrowed` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| `total_reserves` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="gert/hard/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gert/hard/v1beta1/query.proto



<a name="gert.hard.v1beta1.BorrowInterestFactorResponse"></a>

### BorrowInterestFactorResponse
BorrowInterestFactorResponse defines an individual borrow interest factor.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `value` | [string](#string) |  | sdk.Dec as string |






<a name="gert.hard.v1beta1.BorrowResponse"></a>

### BorrowResponse
BorrowResponse defines an amount of coins borrowed from a hard module account.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `borrower` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| `index` | [BorrowInterestFactorResponse](#gert.hard.v1beta1.BorrowInterestFactorResponse) | repeated |  |






<a name="gert.hard.v1beta1.DepositResponse"></a>

### DepositResponse
DepositResponse defines an amount of coins deposited into a hard module account.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `depositor` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| `index` | [SupplyInterestFactorResponse](#gert.hard.v1beta1.SupplyInterestFactorResponse) | repeated |  |






<a name="gert.hard.v1beta1.InterestFactor"></a>

### InterestFactor
InterestFactor is a unique type returned by interest factor queries


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `borrow_interest_factor` | [string](#string) |  | sdk.Dec as String |
| `supply_interest_factor` | [string](#string) |  | sdk.Dec as String |






<a name="gert.hard.v1beta1.MoneyMarketInterestRate"></a>

### MoneyMarketInterestRate
MoneyMarketInterestRate is a unique type returned by interest rate queries


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `supply_interest_rate` | [string](#string) |  | sdk.Dec as String |
| `borrow_interest_rate` | [string](#string) |  | sdk.Dec as String |






<a name="gert.hard.v1beta1.QueryAccountsRequest"></a>

### QueryAccountsRequest
QueryAccountsRequest is the request type for the Query/Accounts RPC method.






<a name="gert.hard.v1beta1.QueryAccountsResponse"></a>

### QueryAccountsResponse
QueryAccountsResponse is the response type for the Query/Accounts RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `accounts` | [cosmos.auth.v1beta1.ModuleAccount](#cosmos.auth.v1beta1.ModuleAccount) | repeated |  |






<a name="gert.hard.v1beta1.QueryBorrowsRequest"></a>

### QueryBorrowsRequest
QueryBorrowsRequest is the request type for the Query/Borrows RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="gert.hard.v1beta1.QueryBorrowsResponse"></a>

### QueryBorrowsResponse
QueryBorrowsResponse is the response type for the Query/Borrows RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `borrows` | [BorrowResponse](#gert.hard.v1beta1.BorrowResponse) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="gert.hard.v1beta1.QueryDepositsRequest"></a>

### QueryDepositsRequest
QueryDepositsRequest is the request type for the Query/Deposits RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="gert.hard.v1beta1.QueryDepositsResponse"></a>

### QueryDepositsResponse
QueryDepositsResponse is the response type for the Query/Deposits RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `deposits` | [DepositResponse](#gert.hard.v1beta1.DepositResponse) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="gert.hard.v1beta1.QueryInterestFactorsRequest"></a>

### QueryInterestFactorsRequest
QueryInterestFactorsRequest is the request type for the Query/InterestFactors RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |






<a name="gert.hard.v1beta1.QueryInterestFactorsResponse"></a>

### QueryInterestFactorsResponse
QueryInterestFactorsResponse is the response type for the Query/InterestFactors RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `interest_factors` | [InterestFactor](#gert.hard.v1beta1.InterestFactor) | repeated |  |






<a name="gert.hard.v1beta1.QueryInterestRateRequest"></a>

### QueryInterestRateRequest
QueryInterestRateRequest is the request type for the Query/InterestRate RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |






<a name="gert.hard.v1beta1.QueryInterestRateResponse"></a>

### QueryInterestRateResponse
QueryInterestRateResponse is the response type for the Query/InterestRate RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `interest_rates` | [MoneyMarketInterestRate](#gert.hard.v1beta1.MoneyMarketInterestRate) | repeated |  |






<a name="gert.hard.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is the request type for the Query/Params RPC method.






<a name="gert.hard.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is the response type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#gert.hard.v1beta1.Params) |  |  |






<a name="gert.hard.v1beta1.QueryReservesRequest"></a>

### QueryReservesRequest
QueryReservesRequest is the request type for the Query/Reserves RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |






<a name="gert.hard.v1beta1.QueryReservesResponse"></a>

### QueryReservesResponse
QueryReservesResponse is the response type for the Query/Reserves RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="gert.hard.v1beta1.QueryTotalBorrowedRequest"></a>

### QueryTotalBorrowedRequest
QueryTotalBorrowedRequest is the request type for the Query/TotalBorrowed RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |






<a name="gert.hard.v1beta1.QueryTotalBorrowedResponse"></a>

### QueryTotalBorrowedResponse
QueryTotalBorrowedResponse is the response type for the Query/TotalBorrowed RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `borrowed_coins` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="gert.hard.v1beta1.QueryTotalDepositedRequest"></a>

### QueryTotalDepositedRequest
QueryTotalDepositedRequest is the request type for the Query/TotalDeposited RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |






<a name="gert.hard.v1beta1.QueryTotalDepositedResponse"></a>

### QueryTotalDepositedResponse
QueryTotalDepositedResponse is the response type for the Query/TotalDeposited RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `supplied_coins` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="gert.hard.v1beta1.QueryUnsyncedBorrowsRequest"></a>

### QueryUnsyncedBorrowsRequest
QueryUnsyncedBorrowsRequest is the request type for the Query/UnsyncedBorrows RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="gert.hard.v1beta1.QueryUnsyncedBorrowsResponse"></a>

### QueryUnsyncedBorrowsResponse
QueryUnsyncedBorrowsResponse is the response type for the Query/UnsyncedBorrows RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `borrows` | [BorrowResponse](#gert.hard.v1beta1.BorrowResponse) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="gert.hard.v1beta1.QueryUnsyncedDepositsRequest"></a>

### QueryUnsyncedDepositsRequest
QueryUnsyncedDepositsRequest is the request type for the Query/UnsyncedDeposits RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="gert.hard.v1beta1.QueryUnsyncedDepositsResponse"></a>

### QueryUnsyncedDepositsResponse
QueryUnsyncedDepositsResponse is the response type for the Query/UnsyncedDeposits RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `deposits` | [DepositResponse](#gert.hard.v1beta1.DepositResponse) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="gert.hard.v1beta1.SupplyInterestFactorResponse"></a>

### SupplyInterestFactorResponse
SupplyInterestFactorResponse defines an individual borrow interest factor.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `value` | [string](#string) |  | sdk.Dec as string |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="gert.hard.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service for bep3 module.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#gert.hard.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#gert.hard.v1beta1.QueryParamsResponse) | Params queries module params. | GET|/gert/hard/v1beta1/params|
| `Accounts` | [QueryAccountsRequest](#gert.hard.v1beta1.QueryAccountsRequest) | [QueryAccountsResponse](#gert.hard.v1beta1.QueryAccountsResponse) | Accounts queries module accounts. | GET|/gert/hard/v1beta1/accounts|
| `Deposits` | [QueryDepositsRequest](#gert.hard.v1beta1.QueryDepositsRequest) | [QueryDepositsResponse](#gert.hard.v1beta1.QueryDepositsResponse) | Deposits queries hard deposits. | GET|/gert/hard/v1beta1/deposits|
| `UnsyncedDeposits` | [QueryUnsyncedDepositsRequest](#gert.hard.v1beta1.QueryUnsyncedDepositsRequest) | [QueryUnsyncedDepositsResponse](#gert.hard.v1beta1.QueryUnsyncedDepositsResponse) | UnsyncedDeposits queries unsynced deposits. | GET|/gert/hard/v1beta1/unsynced-deposits|
| `TotalDeposited` | [QueryTotalDepositedRequest](#gert.hard.v1beta1.QueryTotalDepositedRequest) | [QueryTotalDepositedResponse](#gert.hard.v1beta1.QueryTotalDepositedResponse) | TotalDeposited queries total coins deposited to hard liquidity pools. | GET|/gert/hard/v1beta1/total-deposited/{denom}|
| `Borrows` | [QueryBorrowsRequest](#gert.hard.v1beta1.QueryBorrowsRequest) | [QueryBorrowsResponse](#gert.hard.v1beta1.QueryBorrowsResponse) | Borrows queries hard borrows. | GET|/gert/hard/v1beta1/borrows|
| `UnsyncedBorrows` | [QueryUnsyncedBorrowsRequest](#gert.hard.v1beta1.QueryUnsyncedBorrowsRequest) | [QueryUnsyncedBorrowsResponse](#gert.hard.v1beta1.QueryUnsyncedBorrowsResponse) | UnsyncedBorrows queries unsynced borrows. | GET|/gert/hard/v1beta1/unsynced-borrows|
| `TotalBorrowed` | [QueryTotalBorrowedRequest](#gert.hard.v1beta1.QueryTotalBorrowedRequest) | [QueryTotalBorrowedResponse](#gert.hard.v1beta1.QueryTotalBorrowedResponse) | TotalBorrowed queries total coins borrowed from hard liquidity pools. | GET|/gert/hard/v1beta1/total-borrowed/{denom}|
| `InterestRate` | [QueryInterestRateRequest](#gert.hard.v1beta1.QueryInterestRateRequest) | [QueryInterestRateResponse](#gert.hard.v1beta1.QueryInterestRateResponse) | InterestRate queries the hard module interest rates. | GET|/gert/hard/v1beta1/interest-rate/{denom}|
| `Reserves` | [QueryReservesRequest](#gert.hard.v1beta1.QueryReservesRequest) | [QueryReservesResponse](#gert.hard.v1beta1.QueryReservesResponse) | Reserves queries total hard reserve coins. | GET|/gert/hard/v1beta1/reserves/{denom}|
| `InterestFactors` | [QueryInterestFactorsRequest](#gert.hard.v1beta1.QueryInterestFactorsRequest) | [QueryInterestFactorsResponse](#gert.hard.v1beta1.QueryInterestFactorsResponse) | InterestFactors queries hard module interest factors. | GET|/gert/hard/v1beta1/interest-factors/{denom}|

 <!-- end services -->



<a name="gert/hard/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gert/hard/v1beta1/tx.proto



<a name="gert.hard.v1beta1.MsgBorrow"></a>

### MsgBorrow
MsgBorrow defines the Msg/Borrow request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `borrower` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="gert.hard.v1beta1.MsgBorrowResponse"></a>

### MsgBorrowResponse
MsgBorrowResponse defines the Msg/Borrow response type.






<a name="gert.hard.v1beta1.MsgDeposit"></a>

### MsgDeposit
MsgDeposit defines the Msg/Deposit request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `depositor` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="gert.hard.v1beta1.MsgDepositResponse"></a>

### MsgDepositResponse
MsgDepositResponse defines the Msg/Deposit response type.






<a name="gert.hard.v1beta1.MsgLiquidate"></a>

### MsgLiquidate
MsgLiquidate defines the Msg/Liquidate request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `keeper` | [string](#string) |  |  |
| `borrower` | [string](#string) |  |  |






<a name="gert.hard.v1beta1.MsgLiquidateResponse"></a>

### MsgLiquidateResponse
MsgLiquidateResponse defines the Msg/Liquidate response type.






<a name="gert.hard.v1beta1.MsgRepay"></a>

### MsgRepay
MsgRepay defines the Msg/Repay request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="gert.hard.v1beta1.MsgRepayResponse"></a>

### MsgRepayResponse
MsgRepayResponse defines the Msg/Repay response type.






<a name="gert.hard.v1beta1.MsgWithdraw"></a>

### MsgWithdraw
MsgWithdraw defines the Msg/Withdraw request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `depositor` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="gert.hard.v1beta1.MsgWithdrawResponse"></a>

### MsgWithdrawResponse
MsgWithdrawResponse defines the Msg/Withdraw response type.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="gert.hard.v1beta1.Msg"></a>

### Msg
Msg defines the hard Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Deposit` | [MsgDeposit](#gert.hard.v1beta1.MsgDeposit) | [MsgDepositResponse](#gert.hard.v1beta1.MsgDepositResponse) | Deposit defines a method for depositing funds to hard liquidity pool. | |
| `Withdraw` | [MsgWithdraw](#gert.hard.v1beta1.MsgWithdraw) | [MsgWithdrawResponse](#gert.hard.v1beta1.MsgWithdrawResponse) | Withdraw defines a method for withdrawing funds from hard liquidity pool. | |
| `Borrow` | [MsgBorrow](#gert.hard.v1beta1.MsgBorrow) | [MsgBorrowResponse](#gert.hard.v1beta1.MsgBorrowResponse) | Borrow defines a method for borrowing funds from hard liquidity pool. | |
| `Repay` | [MsgRepay](#gert.hard.v1beta1.MsgRepay) | [MsgRepayResponse](#gert.hard.v1beta1.MsgRepayResponse) | Repay defines a method for repaying funds borrowed from hard liquidity pool. | |
| `Liquidate` | [MsgLiquidate](#gert.hard.v1beta1.MsgLiquidate) | [MsgLiquidateResponse](#gert.hard.v1beta1.MsgLiquidateResponse) | Liquidate defines a method for attempting to liquidate a borrower that is over their loan-to-value. | |

 <!-- end services -->



<a name="gert/incentive/v1beta1/claims.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gert/incentive/v1beta1/claims.proto



<a name="gert.incentive.v1beta1.BaseClaim"></a>

### BaseClaim
BaseClaim is a claim with a single reward coin types


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `owner` | [bytes](#bytes) |  |  |
| `reward` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="gert.incentive.v1beta1.BaseMultiClaim"></a>

### BaseMultiClaim
BaseMultiClaim is a claim with multiple reward coin types


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `owner` | [bytes](#bytes) |  |  |
| `reward` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="gert.incentive.v1beta1.DelegatorClaim"></a>

### DelegatorClaim
DelegatorClaim stores delegation rewards that can be claimed by owner


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_claim` | [BaseMultiClaim](#gert.incentive.v1beta1.BaseMultiClaim) |  |  |
| `reward_indexes` | [MultiRewardIndex](#gert.incentive.v1beta1.MultiRewardIndex) | repeated |  |






<a name="gert.incentive.v1beta1.HardLiquidityProviderClaim"></a>

### HardLiquidityProviderClaim
HardLiquidityProviderClaim stores the hard liquidity provider rewards that can be claimed by owner


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_claim` | [BaseMultiClaim](#gert.incentive.v1beta1.BaseMultiClaim) |  |  |
| `supply_reward_indexes` | [MultiRewardIndex](#gert.incentive.v1beta1.MultiRewardIndex) | repeated |  |
| `borrow_reward_indexes` | [MultiRewardIndex](#gert.incentive.v1beta1.MultiRewardIndex) | repeated |  |






<a name="gert.incentive.v1beta1.MultiRewardIndex"></a>

### MultiRewardIndex
MultiRewardIndex stores reward accumulation information on multiple reward types


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_type` | [string](#string) |  |  |
| `reward_indexes` | [RewardIndex](#gert.incentive.v1beta1.RewardIndex) | repeated |  |






<a name="gert.incentive.v1beta1.MultiRewardIndexesProto"></a>

### MultiRewardIndexesProto
MultiRewardIndexesProto defines a Protobuf wrapper around a MultiRewardIndexes slice


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `multi_reward_indexes` | [MultiRewardIndex](#gert.incentive.v1beta1.MultiRewardIndex) | repeated |  |






<a name="gert.incentive.v1beta1.RewardIndex"></a>

### RewardIndex
RewardIndex stores reward accumulation information


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_type` | [string](#string) |  |  |
| `reward_factor` | [bytes](#bytes) |  |  |






<a name="gert.incentive.v1beta1.RewardIndexesProto"></a>

### RewardIndexesProto
RewardIndexesProto defines a Protobuf wrapper around a RewardIndexes slice


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `reward_indexes` | [RewardIndex](#gert.incentive.v1beta1.RewardIndex) | repeated |  |






<a name="gert.incentive.v1beta1.SavingsClaim"></a>

### SavingsClaim
SavingsClaim stores the savings rewards that can be claimed by owner


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_claim` | [BaseMultiClaim](#gert.incentive.v1beta1.BaseMultiClaim) |  |  |
| `reward_indexes` | [MultiRewardIndex](#gert.incentive.v1beta1.MultiRewardIndex) | repeated |  |






<a name="gert.incentive.v1beta1.SwapClaim"></a>

### SwapClaim
SwapClaim stores the swap rewards that can be claimed by owner


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_claim` | [BaseMultiClaim](#gert.incentive.v1beta1.BaseMultiClaim) |  |  |
| `reward_indexes` | [MultiRewardIndex](#gert.incentive.v1beta1.MultiRewardIndex) | repeated |  |






<a name="gert.incentive.v1beta1.USDXMintingClaim"></a>

### USDXMintingClaim
USDXMintingClaim is for USDX minting rewards


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_claim` | [BaseClaim](#gert.incentive.v1beta1.BaseClaim) |  |  |
| `reward_indexes` | [RewardIndex](#gert.incentive.v1beta1.RewardIndex) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="gert/incentive/v1beta1/params.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gert/incentive/v1beta1/params.proto



<a name="gert.incentive.v1beta1.MultiRewardPeriod"></a>

### MultiRewardPeriod
MultiRewardPeriod supports multiple reward types


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `active` | [bool](#bool) |  |  |
| `collateral_type` | [string](#string) |  |  |
| `start` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `end` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `rewards_per_second` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="gert.incentive.v1beta1.Multiplier"></a>

### Multiplier
Multiplier amount the claim rewards get increased by, along with how long the claim rewards are locked


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `name` | [string](#string) |  |  |
| `months_lockup` | [int64](#int64) |  |  |
| `factor` | [bytes](#bytes) |  |  |






<a name="gert.incentive.v1beta1.MultipliersPerDenom"></a>

### MultipliersPerDenom
MultipliersPerDenom is a map of denoms to a set of multipliers


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `multipliers` | [Multiplier](#gert.incentive.v1beta1.Multiplier) | repeated |  |






<a name="gert.incentive.v1beta1.Params"></a>

### Params
Params


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `usdx_minting_reward_periods` | [RewardPeriod](#gert.incentive.v1beta1.RewardPeriod) | repeated |  |
| `hard_supply_reward_periods` | [MultiRewardPeriod](#gert.incentive.v1beta1.MultiRewardPeriod) | repeated |  |
| `hard_borrow_reward_periods` | [MultiRewardPeriod](#gert.incentive.v1beta1.MultiRewardPeriod) | repeated |  |
| `delegator_reward_periods` | [MultiRewardPeriod](#gert.incentive.v1beta1.MultiRewardPeriod) | repeated |  |
| `swap_reward_periods` | [MultiRewardPeriod](#gert.incentive.v1beta1.MultiRewardPeriod) | repeated |  |
| `claim_multipliers` | [MultipliersPerDenom](#gert.incentive.v1beta1.MultipliersPerDenom) | repeated |  |
| `claim_end` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `savings_reward_periods` | [MultiRewardPeriod](#gert.incentive.v1beta1.MultiRewardPeriod) | repeated |  |






<a name="gert.incentive.v1beta1.RewardPeriod"></a>

### RewardPeriod
RewardPeriod stores the state of an ongoing reward


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `active` | [bool](#bool) |  |  |
| `collateral_type` | [string](#string) |  |  |
| `start` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `end` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `rewards_per_second` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="gert/incentive/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gert/incentive/v1beta1/genesis.proto



<a name="gert.incentive.v1beta1.AccumulationTime"></a>

### AccumulationTime
AccumulationTime stores the previous reward distribution time and its corresponding collateral type


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_type` | [string](#string) |  |  |
| `previous_accumulation_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="gert.incentive.v1beta1.GenesisRewardState"></a>

### GenesisRewardState
GenesisRewardState groups together the global state for a particular reward so it can be exported in genesis.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `accumulation_times` | [AccumulationTime](#gert.incentive.v1beta1.AccumulationTime) | repeated |  |
| `multi_reward_indexes` | [MultiRewardIndex](#gert.incentive.v1beta1.MultiRewardIndex) | repeated |  |






<a name="gert.incentive.v1beta1.GenesisState"></a>

### GenesisState
GenesisState is the state that must be provided at genesis.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#gert.incentive.v1beta1.Params) |  |  |
| `usdx_reward_state` | [GenesisRewardState](#gert.incentive.v1beta1.GenesisRewardState) |  |  |
| `hard_supply_reward_state` | [GenesisRewardState](#gert.incentive.v1beta1.GenesisRewardState) |  |  |
| `hard_borrow_reward_state` | [GenesisRewardState](#gert.incentive.v1beta1.GenesisRewardState) |  |  |
| `delegator_reward_state` | [GenesisRewardState](#gert.incentive.v1beta1.GenesisRewardState) |  |  |
| `swap_reward_state` | [GenesisRewardState](#gert.incentive.v1beta1.GenesisRewardState) |  |  |
| `usdx_minting_claims` | [USDXMintingClaim](#gert.incentive.v1beta1.USDXMintingClaim) | repeated |  |
| `hard_liquidity_provider_claims` | [HardLiquidityProviderClaim](#gert.incentive.v1beta1.HardLiquidityProviderClaim) | repeated |  |
| `delegator_claims` | [DelegatorClaim](#gert.incentive.v1beta1.DelegatorClaim) | repeated |  |
| `swap_claims` | [SwapClaim](#gert.incentive.v1beta1.SwapClaim) | repeated |  |
| `savings_reward_state` | [GenesisRewardState](#gert.incentive.v1beta1.GenesisRewardState) |  |  |
| `savings_claims` | [SavingsClaim](#gert.incentive.v1beta1.SavingsClaim) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="gert/incentive/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gert/incentive/v1beta1/tx.proto



<a name="gert.incentive.v1beta1.MsgClaimDelegatorReward"></a>

### MsgClaimDelegatorReward
MsgClaimDelegatorReward message type used to claim delegator rewards


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `denoms_to_claim` | [Selection](#gert.incentive.v1beta1.Selection) | repeated |  |






<a name="gert.incentive.v1beta1.MsgClaimDelegatorRewardResponse"></a>

### MsgClaimDelegatorRewardResponse
MsgClaimDelegatorRewardResponse defines the Msg/ClaimDelegatorReward response type.






<a name="gert.incentive.v1beta1.MsgClaimHardReward"></a>

### MsgClaimHardReward
MsgClaimHardReward message type used to claim Hard liquidity provider rewards


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `denoms_to_claim` | [Selection](#gert.incentive.v1beta1.Selection) | repeated |  |






<a name="gert.incentive.v1beta1.MsgClaimHardRewardResponse"></a>

### MsgClaimHardRewardResponse
MsgClaimHardRewardResponse defines the Msg/ClaimHardReward response type.






<a name="gert.incentive.v1beta1.MsgClaimSavingsReward"></a>

### MsgClaimSavingsReward
MsgClaimSavingsReward message type used to claim savings rewards


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `denoms_to_claim` | [Selection](#gert.incentive.v1beta1.Selection) | repeated |  |






<a name="gert.incentive.v1beta1.MsgClaimSavingsRewardResponse"></a>

### MsgClaimSavingsRewardResponse
MsgClaimSavingsRewardResponse defines the Msg/ClaimSavingsReward response type.






<a name="gert.incentive.v1beta1.MsgClaimSwapReward"></a>

### MsgClaimSwapReward
MsgClaimSwapReward message type used to claim delegator rewards


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `denoms_to_claim` | [Selection](#gert.incentive.v1beta1.Selection) | repeated |  |






<a name="gert.incentive.v1beta1.MsgClaimSwapRewardResponse"></a>

### MsgClaimSwapRewardResponse
MsgClaimSwapRewardResponse defines the Msg/ClaimSwapReward response type.






<a name="gert.incentive.v1beta1.MsgClaimUSDXMintingReward"></a>

### MsgClaimUSDXMintingReward
MsgClaimUSDXMintingReward message type used to claim USDX minting rewards


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `multiplier_name` | [string](#string) |  |  |






<a name="gert.incentive.v1beta1.MsgClaimUSDXMintingRewardResponse"></a>

### MsgClaimUSDXMintingRewardResponse
MsgClaimUSDXMintingRewardResponse defines the Msg/ClaimUSDXMintingReward response type.






<a name="gert.incentive.v1beta1.Selection"></a>

### Selection
Selection is a pair of denom and multiplier name. It holds the choice of multiplier a user makes when they claim a
denom.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `multiplier_name` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="gert.incentive.v1beta1.Msg"></a>

### Msg
Msg defines the incentive Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `ClaimUSDXMintingReward` | [MsgClaimUSDXMintingReward](#gert.incentive.v1beta1.MsgClaimUSDXMintingReward) | [MsgClaimUSDXMintingRewardResponse](#gert.incentive.v1beta1.MsgClaimUSDXMintingRewardResponse) | ClaimUSDXMintingReward is a message type used to claim USDX minting rewards | |
| `ClaimHardReward` | [MsgClaimHardReward](#gert.incentive.v1beta1.MsgClaimHardReward) | [MsgClaimHardRewardResponse](#gert.incentive.v1beta1.MsgClaimHardRewardResponse) | ClaimHardReward is a message type used to claim Hard liquidity provider rewards | |
| `ClaimDelegatorReward` | [MsgClaimDelegatorReward](#gert.incentive.v1beta1.MsgClaimDelegatorReward) | [MsgClaimDelegatorRewardResponse](#gert.incentive.v1beta1.MsgClaimDelegatorRewardResponse) | ClaimDelegatorReward is a message type used to claim delegator rewards | |
| `ClaimSwapReward` | [MsgClaimSwapReward](#gert.incentive.v1beta1.MsgClaimSwapReward) | [MsgClaimSwapRewardResponse](#gert.incentive.v1beta1.MsgClaimSwapRewardResponse) | ClaimSwapReward is a message type used to claim delegator rewards | |
| `ClaimSavingsReward` | [MsgClaimSavingsReward](#gert.incentive.v1beta1.MsgClaimSavingsReward) | [MsgClaimSavingsRewardResponse](#gert.incentive.v1beta1.MsgClaimSavingsRewardResponse) | ClaimSavingsReward is a message type used to claim savings rewards | |

 <!-- end services -->



<a name="gert/issuance/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gert/issuance/v1beta1/genesis.proto



<a name="gert.issuance.v1beta1.Asset"></a>

### Asset
Asset type for assets in the issuance module


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `owner` | [string](#string) |  |  |
| `denom` | [string](#string) |  |  |
| `blocked_addresses` | [string](#string) | repeated |  |
| `paused` | [bool](#bool) |  |  |
| `blockable` | [bool](#bool) |  |  |
| `rate_limit` | [RateLimit](#gert.issuance.v1beta1.RateLimit) |  |  |






<a name="gert.issuance.v1beta1.AssetSupply"></a>

### AssetSupply
AssetSupply contains information about an asset's rate-limited supply (the
total supply of the asset is tracked in the top-level supply module)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `current_supply` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `time_elapsed` | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |






<a name="gert.issuance.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the issuance module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#gert.issuance.v1beta1.Params) |  | params defines all the paramaters of the module. |
| `supplies` | [AssetSupply](#gert.issuance.v1beta1.AssetSupply) | repeated |  |






<a name="gert.issuance.v1beta1.Params"></a>

### Params
Params defines the parameters for the issuance module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `assets` | [Asset](#gert.issuance.v1beta1.Asset) | repeated |  |






<a name="gert.issuance.v1beta1.RateLimit"></a>

### RateLimit
RateLimit parameters for rate-limiting the supply of an issued asset


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `active` | [bool](#bool) |  |  |
| `limit` | [bytes](#bytes) |  |  |
| `time_period` | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="gert/issuance/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gert/issuance/v1beta1/query.proto



<a name="gert.issuance.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest defines the request type for querying x/issuance parameters.






<a name="gert.issuance.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse defines the response type for querying x/issuance parameters.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#gert.issuance.v1beta1.Params) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="gert.issuance.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service for issuance module

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#gert.issuance.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#gert.issuance.v1beta1.QueryParamsResponse) | Params queries all parameters of the issuance module. | GET|/gert/issuance/v1beta1/params|

 <!-- end services -->



<a name="gert/issuance/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gert/issuance/v1beta1/tx.proto



<a name="gert.issuance.v1beta1.MsgBlockAddress"></a>

### MsgBlockAddress
MsgBlockAddress represents a message used by the issuer to block an address from holding or transferring tokens


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `denom` | [string](#string) |  |  |
| `blocked_address` | [string](#string) |  |  |






<a name="gert.issuance.v1beta1.MsgBlockAddressResponse"></a>

### MsgBlockAddressResponse
MsgBlockAddressResponse defines the Msg/BlockAddress response type.






<a name="gert.issuance.v1beta1.MsgIssueTokens"></a>

### MsgIssueTokens
MsgIssueTokens represents a message used by the issuer to issue new tokens


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `tokens` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `receiver` | [string](#string) |  |  |






<a name="gert.issuance.v1beta1.MsgIssueTokensResponse"></a>

### MsgIssueTokensResponse
MsgIssueTokensResponse defines the Msg/IssueTokens response type.






<a name="gert.issuance.v1beta1.MsgRedeemTokens"></a>

### MsgRedeemTokens
MsgRedeemTokens represents a message used by the issuer to redeem (burn) tokens


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `tokens` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="gert.issuance.v1beta1.MsgRedeemTokensResponse"></a>

### MsgRedeemTokensResponse
MsgRedeemTokensResponse defines the Msg/RedeemTokens response type.






<a name="gert.issuance.v1beta1.MsgSetPauseStatus"></a>

### MsgSetPauseStatus
MsgSetPauseStatus message type used by the issuer to pause or unpause status


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `denom` | [string](#string) |  |  |
| `status` | [bool](#bool) |  |  |






<a name="gert.issuance.v1beta1.MsgSetPauseStatusResponse"></a>

### MsgSetPauseStatusResponse
MsgSetPauseStatusResponse defines the Msg/SetPauseStatus response type.






<a name="gert.issuance.v1beta1.MsgUnblockAddress"></a>

### MsgUnblockAddress
MsgUnblockAddress message type used by the issuer to unblock an address from holding or transferring tokens


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `denom` | [string](#string) |  |  |
| `blocked_address` | [string](#string) |  |  |






<a name="gert.issuance.v1beta1.MsgUnblockAddressResponse"></a>

### MsgUnblockAddressResponse
MsgUnblockAddressResponse defines the Msg/UnblockAddress response type.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="gert.issuance.v1beta1.Msg"></a>

### Msg
Msg defines the issuance Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `IssueTokens` | [MsgIssueTokens](#gert.issuance.v1beta1.MsgIssueTokens) | [MsgIssueTokensResponse](#gert.issuance.v1beta1.MsgIssueTokensResponse) | IssueTokens message type used by the issuer to issue new tokens | |
| `RedeemTokens` | [MsgRedeemTokens](#gert.issuance.v1beta1.MsgRedeemTokens) | [MsgRedeemTokensResponse](#gert.issuance.v1beta1.MsgRedeemTokensResponse) | RedeemTokens message type used by the issuer to redeem (burn) tokens | |
| `BlockAddress` | [MsgBlockAddress](#gert.issuance.v1beta1.MsgBlockAddress) | [MsgBlockAddressResponse](#gert.issuance.v1beta1.MsgBlockAddressResponse) | BlockAddress message type used by the issuer to block an address from holding or transferring tokens | |
| `UnblockAddress` | [MsgUnblockAddress](#gert.issuance.v1beta1.MsgUnblockAddress) | [MsgUnblockAddressResponse](#gert.issuance.v1beta1.MsgUnblockAddressResponse) | UnblockAddress message type used by the issuer to unblock an address from holding or transferring tokens | |
| `SetPauseStatus` | [MsgSetPauseStatus](#gert.issuance.v1beta1.MsgSetPauseStatus) | [MsgSetPauseStatusResponse](#gert.issuance.v1beta1.MsgSetPauseStatusResponse) | SetPauseStatus message type used to pause or unpause status | |

 <!-- end services -->



<a name="gert/gertdist/v1beta1/params.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gert/gertdist/v1beta1/params.proto



<a name="gert.gertdist.v1beta1.Params"></a>

### Params
Params governance parameters for gertdist module


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `active` | [bool](#bool) |  |  |
| `periods` | [Period](#gert.gertdist.v1beta1.Period) | repeated |  |






<a name="gert.gertdist.v1beta1.Period"></a>

### Period
Period stores the specified start and end dates, and the inflation, expressed as a decimal
representing the yearly APR of gert tokens that will be minted during that period


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `start` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | example "2020-03-01T15:20:00Z" |
| `end` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | example "2020-06-01T15:20:00Z" |
| `inflation` | [bytes](#bytes) |  | example "1.000000003022265980" - 10% inflation |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="gert/gertdist/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gert/gertdist/v1beta1/genesis.proto



<a name="gert.gertdist.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the gertdist module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#gert.gertdist.v1beta1.Params) |  |  |
| `previous_block_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="gert/gertdist/v1beta1/proposal.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gert/gertdist/v1beta1/proposal.proto



<a name="gert.gertdist.v1beta1.CommunityPoolMultiSpendProposal"></a>

### CommunityPoolMultiSpendProposal
CommunityPoolMultiSpendProposal spends from the community pool by sending to one or more
addresses


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  |  |
| `description` | [string](#string) |  |  |
| `recipient_list` | [MultiSpendRecipient](#gert.gertdist.v1beta1.MultiSpendRecipient) | repeated |  |






<a name="gert.gertdist.v1beta1.CommunityPoolMultiSpendProposalJSON"></a>

### CommunityPoolMultiSpendProposalJSON
CommunityPoolMultiSpendProposalJSON defines a CommunityPoolMultiSpendProposal with a deposit


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  |  |
| `description` | [string](#string) |  |  |
| `recipient_list` | [MultiSpendRecipient](#gert.gertdist.v1beta1.MultiSpendRecipient) | repeated |  |
| `deposit` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="gert.gertdist.v1beta1.MultiSpendRecipient"></a>

### MultiSpendRecipient
MultiSpendRecipient defines a recipient and the amount of coins they are receiving


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="gert/gertdist/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gert/gertdist/v1beta1/query.proto



<a name="gert.gertdist.v1beta1.QueryBalanceRequest"></a>

### QueryBalanceRequest
QueryBalanceRequest defines the request type for querying x/gertdist balance.






<a name="gert.gertdist.v1beta1.QueryBalanceResponse"></a>

### QueryBalanceResponse
QueryBalanceResponse defines the response type for querying x/gertdist balance.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `coins` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="gert.gertdist.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest defines the request type for querying x/gertdist parameters.






<a name="gert.gertdist.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse defines the response type for querying x/gertdist parameters.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#gert.gertdist.v1beta1.Params) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="gert.gertdist.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#gert.gertdist.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#gert.gertdist.v1beta1.QueryParamsResponse) | Params queries the parameters of x/gertdist module. | GET|/gert/gertdist/v1beta1/parameters|
| `Balance` | [QueryBalanceRequest](#gert.gertdist.v1beta1.QueryBalanceRequest) | [QueryBalanceResponse](#gert.gertdist.v1beta1.QueryBalanceResponse) | Balance queries the balance of all coins of x/gertdist module. | GET|/gert/gertdist/v1beta1/balance|

 <!-- end services -->



<a name="gert/pricefeed/v1beta1/store.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gert/pricefeed/v1beta1/store.proto



<a name="gert.pricefeed.v1beta1.CurrentPrice"></a>

### CurrentPrice
CurrentPrice defines a current price for a particular market in the pricefeed
module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `market_id` | [string](#string) |  |  |
| `price` | [string](#string) |  |  |






<a name="gert.pricefeed.v1beta1.Market"></a>

### Market
Market defines an asset in the pricefeed.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `market_id` | [string](#string) |  |  |
| `base_asset` | [string](#string) |  |  |
| `quote_asset` | [string](#string) |  |  |
| `oracles` | [bytes](#bytes) | repeated |  |
| `active` | [bool](#bool) |  |  |






<a name="gert.pricefeed.v1beta1.Params"></a>

### Params
Params defines the parameters for the pricefeed module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `markets` | [Market](#gert.pricefeed.v1beta1.Market) | repeated |  |






<a name="gert.pricefeed.v1beta1.PostedPrice"></a>

### PostedPrice
PostedPrice defines a price for market posted by a specific oracle.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `market_id` | [string](#string) |  |  |
| `oracle_address` | [bytes](#bytes) |  |  |
| `price` | [string](#string) |  |  |
| `expiry` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="gert/pricefeed/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gert/pricefeed/v1beta1/genesis.proto



<a name="gert.pricefeed.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the pricefeed module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#gert.pricefeed.v1beta1.Params) |  | params defines all the paramaters of the module. |
| `posted_prices` | [PostedPrice](#gert.pricefeed.v1beta1.PostedPrice) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="gert/pricefeed/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gert/pricefeed/v1beta1/query.proto



<a name="gert.pricefeed.v1beta1.CurrentPriceResponse"></a>

### CurrentPriceResponse
CurrentPriceResponse defines a current price for a particular market in the pricefeed
module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `market_id` | [string](#string) |  |  |
| `price` | [string](#string) |  |  |






<a name="gert.pricefeed.v1beta1.MarketResponse"></a>

### MarketResponse
MarketResponse defines an asset in the pricefeed.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `market_id` | [string](#string) |  |  |
| `base_asset` | [string](#string) |  |  |
| `quote_asset` | [string](#string) |  |  |
| `oracles` | [string](#string) | repeated |  |
| `active` | [bool](#bool) |  |  |






<a name="gert.pricefeed.v1beta1.PostedPriceResponse"></a>

### PostedPriceResponse
PostedPriceResponse defines a price for market posted by a specific oracle.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `market_id` | [string](#string) |  |  |
| `oracle_address` | [string](#string) |  |  |
| `price` | [string](#string) |  |  |
| `expiry` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="gert.pricefeed.v1beta1.QueryMarketsRequest"></a>

### QueryMarketsRequest
QueryMarketsRequest is the request type for the Query/Markets RPC method.






<a name="gert.pricefeed.v1beta1.QueryMarketsResponse"></a>

### QueryMarketsResponse
QueryMarketsResponse is the response type for the Query/Markets RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `markets` | [MarketResponse](#gert.pricefeed.v1beta1.MarketResponse) | repeated | List of markets |






<a name="gert.pricefeed.v1beta1.QueryOraclesRequest"></a>

### QueryOraclesRequest
QueryOraclesRequest is the request type for the Query/Oracles RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `market_id` | [string](#string) |  |  |






<a name="gert.pricefeed.v1beta1.QueryOraclesResponse"></a>

### QueryOraclesResponse
QueryOraclesResponse is the response type for the Query/Oracles RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `oracles` | [string](#string) | repeated | List of oracle addresses |






<a name="gert.pricefeed.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest defines the request type for querying x/pricefeed
parameters.






<a name="gert.pricefeed.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse defines the response type for querying x/pricefeed
parameters.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#gert.pricefeed.v1beta1.Params) |  |  |






<a name="gert.pricefeed.v1beta1.QueryPriceRequest"></a>

### QueryPriceRequest
QueryPriceRequest is the request type for the Query/PriceRequest RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `market_id` | [string](#string) |  |  |






<a name="gert.pricefeed.v1beta1.QueryPriceResponse"></a>

### QueryPriceResponse
QueryPriceResponse is the response type for the Query/Prices RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `price` | [CurrentPriceResponse](#gert.pricefeed.v1beta1.CurrentPriceResponse) |  |  |






<a name="gert.pricefeed.v1beta1.QueryPricesRequest"></a>

### QueryPricesRequest
QueryPricesRequest is the request type for the Query/Prices RPC method.






<a name="gert.pricefeed.v1beta1.QueryPricesResponse"></a>

### QueryPricesResponse
QueryPricesResponse is the response type for the Query/Prices RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `prices` | [CurrentPriceResponse](#gert.pricefeed.v1beta1.CurrentPriceResponse) | repeated |  |






<a name="gert.pricefeed.v1beta1.QueryRawPricesRequest"></a>

### QueryRawPricesRequest
QueryRawPricesRequest is the request type for the Query/RawPrices RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `market_id` | [string](#string) |  |  |






<a name="gert.pricefeed.v1beta1.QueryRawPricesResponse"></a>

### QueryRawPricesResponse
QueryRawPricesResponse is the response type for the Query/RawPrices RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `raw_prices` | [PostedPriceResponse](#gert.pricefeed.v1beta1.PostedPriceResponse) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="gert.pricefeed.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service for pricefeed module

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#gert.pricefeed.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#gert.pricefeed.v1beta1.QueryParamsResponse) | Params queries all parameters of the pricefeed module. | GET|/gert/pricefeed/v1beta1/params|
| `Price` | [QueryPriceRequest](#gert.pricefeed.v1beta1.QueryPriceRequest) | [QueryPriceResponse](#gert.pricefeed.v1beta1.QueryPriceResponse) | Price queries price details based on a market | GET|/gert/pricefeed/v1beta1/prices/{market_id}|
| `Prices` | [QueryPricesRequest](#gert.pricefeed.v1beta1.QueryPricesRequest) | [QueryPricesResponse](#gert.pricefeed.v1beta1.QueryPricesResponse) | Prices queries all prices | GET|/gert/pricefeed/v1beta1/prices|
| `RawPrices` | [QueryRawPricesRequest](#gert.pricefeed.v1beta1.QueryRawPricesRequest) | [QueryRawPricesResponse](#gert.pricefeed.v1beta1.QueryRawPricesResponse) | RawPrices queries all raw prices based on a market | GET|/gert/pricefeed/v1beta1/rawprices/{market_id}|
| `Oracles` | [QueryOraclesRequest](#gert.pricefeed.v1beta1.QueryOraclesRequest) | [QueryOraclesResponse](#gert.pricefeed.v1beta1.QueryOraclesResponse) | Oracles queries all oracles based on a market | GET|/gert/pricefeed/v1beta1/oracles/{market_id}|
| `Markets` | [QueryMarketsRequest](#gert.pricefeed.v1beta1.QueryMarketsRequest) | [QueryMarketsResponse](#gert.pricefeed.v1beta1.QueryMarketsResponse) | Markets queries all markets | GET|/gert/pricefeed/v1beta1/markets|

 <!-- end services -->



<a name="gert/pricefeed/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gert/pricefeed/v1beta1/tx.proto



<a name="gert.pricefeed.v1beta1.MsgPostPrice"></a>

### MsgPostPrice
MsgPostPrice represents a method for creating a new post price


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `from` | [string](#string) |  | address of client |
| `market_id` | [string](#string) |  |  |
| `price` | [string](#string) |  |  |
| `expiry` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="gert.pricefeed.v1beta1.MsgPostPriceResponse"></a>

### MsgPostPriceResponse
MsgPostPriceResponse defines the Msg/PostPrice response type.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="gert.pricefeed.v1beta1.Msg"></a>

### Msg
Msg defines the pricefeed Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `PostPrice` | [MsgPostPrice](#gert.pricefeed.v1beta1.MsgPostPrice) | [MsgPostPriceResponse](#gert.pricefeed.v1beta1.MsgPostPriceResponse) | PostPrice defines a method for creating a new post price | |

 <!-- end services -->



<a name="gert/savings/v1beta1/store.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gert/savings/v1beta1/store.proto



<a name="gert.savings.v1beta1.Deposit"></a>

### Deposit
Deposit defines an amount of coins deposited into a savings module account.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `depositor` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="gert.savings.v1beta1.Params"></a>

### Params
Params defines the parameters for the savings module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `supported_denoms` | [string](#string) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="gert/savings/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gert/savings/v1beta1/genesis.proto



<a name="gert.savings.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the savings module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#gert.savings.v1beta1.Params) |  | params defines all the parameters of the module. |
| `deposits` | [Deposit](#gert.savings.v1beta1.Deposit) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="gert/savings/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gert/savings/v1beta1/query.proto



<a name="gert.savings.v1beta1.QueryDepositsRequest"></a>

### QueryDepositsRequest
QueryDepositsRequest defines the request type for querying x/savings
deposits.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="gert.savings.v1beta1.QueryDepositsResponse"></a>

### QueryDepositsResponse
QueryDepositsResponse defines the response type for querying x/savings
deposits.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `deposits` | [Deposit](#gert.savings.v1beta1.Deposit) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="gert.savings.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest defines the request type for querying x/savings
parameters.






<a name="gert.savings.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse defines the response type for querying x/savings
parameters.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#gert.savings.v1beta1.Params) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="gert.savings.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service for savings module

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#gert.savings.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#gert.savings.v1beta1.QueryParamsResponse) | Params queries all parameters of the savings module. | GET|/gert/savings/v1beta1/params|
| `Deposits` | [QueryDepositsRequest](#gert.savings.v1beta1.QueryDepositsRequest) | [QueryDepositsResponse](#gert.savings.v1beta1.QueryDepositsResponse) | Deposits queries savings deposits. | GET|/gert/savings/v1beta1/deposits|

 <!-- end services -->



<a name="gert/savings/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gert/savings/v1beta1/tx.proto



<a name="gert.savings.v1beta1.MsgDeposit"></a>

### MsgDeposit
MsgDeposit defines the Msg/Deposit request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `depositor` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="gert.savings.v1beta1.MsgDepositResponse"></a>

### MsgDepositResponse
MsgDepositResponse defines the Msg/Deposit response type.






<a name="gert.savings.v1beta1.MsgWithdraw"></a>

### MsgWithdraw
MsgWithdraw defines the Msg/Withdraw request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `depositor` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="gert.savings.v1beta1.MsgWithdrawResponse"></a>

### MsgWithdrawResponse
MsgWithdrawResponse defines the Msg/Withdraw response type.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="gert.savings.v1beta1.Msg"></a>

### Msg
Msg defines the savings Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Deposit` | [MsgDeposit](#gert.savings.v1beta1.MsgDeposit) | [MsgDepositResponse](#gert.savings.v1beta1.MsgDepositResponse) | Deposit defines a method for depositing funds to the savings module account | |
| `Withdraw` | [MsgWithdraw](#gert.savings.v1beta1.MsgWithdraw) | [MsgWithdrawResponse](#gert.savings.v1beta1.MsgWithdrawResponse) | Withdraw defines a method for withdrawing funds to the savings module account | |

 <!-- end services -->



<a name="gert/swap/v1beta1/swap.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gert/swap/v1beta1/swap.proto



<a name="gert.swap.v1beta1.AllowedPool"></a>

### AllowedPool
AllowedPool defines a pool that is allowed to be created


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `token_a` | [string](#string) |  | token_a represents the a token allowed |
| `token_b` | [string](#string) |  | token_b represents the b token allowed |






<a name="gert.swap.v1beta1.Params"></a>

### Params
Params defines the parameters for the swap module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `allowed_pools` | [AllowedPool](#gert.swap.v1beta1.AllowedPool) | repeated | allowed_pools defines that pools that are allowed to be created |
| `swap_fee` | [string](#string) |  | swap_fee defines the swap fee for all pools |






<a name="gert.swap.v1beta1.PoolRecord"></a>

### PoolRecord
PoolRecord represents the state of a liquidity pool
and is used to store the state of a denominated pool


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pool_id` | [string](#string) |  | pool_id represents the unique id of the pool |
| `reserves_a` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | reserves_a is the a token coin reserves |
| `reserves_b` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | reserves_b is the a token coin reserves |
| `total_shares` | [string](#string) |  | total_shares is the total distrubuted shares of the pool |






<a name="gert.swap.v1beta1.ShareRecord"></a>

### ShareRecord
ShareRecord stores the shares owned for a depositor and pool


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `depositor` | [bytes](#bytes) |  | depositor represents the owner of the shares |
| `pool_id` | [string](#string) |  | pool_id represents the pool the shares belong to |
| `shares_owned` | [string](#string) |  | shares_owned represents the number of shares owned by depsoitor for the pool_id |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="gert/swap/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gert/swap/v1beta1/genesis.proto



<a name="gert.swap.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the swap module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#gert.swap.v1beta1.Params) |  | params defines all the paramaters related to swap |
| `pool_records` | [PoolRecord](#gert.swap.v1beta1.PoolRecord) | repeated | pool_records defines the available pools |
| `share_records` | [ShareRecord](#gert.swap.v1beta1.ShareRecord) | repeated | share_records defines the owned shares of each pool |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="gert/swap/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gert/swap/v1beta1/query.proto



<a name="gert.swap.v1beta1.DepositResponse"></a>

### DepositResponse
DepositResponse defines a single deposit query response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `depositor` | [string](#string) |  | depositor represents the owner of the deposit |
| `pool_id` | [string](#string) |  | pool_id represents the pool the deposit is for |
| `shares_owned` | [string](#string) |  | shares_owned presents the shares owned by the depositor for the pool |
| `shares_value` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | shares_value represents the coin value of the shares_owned |






<a name="gert.swap.v1beta1.PoolResponse"></a>

### PoolResponse
Pool represents the state of a single pool


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `name` | [string](#string) |  | name represents the name of the pool |
| `coins` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | coins represents the total reserves of the pool |
| `total_shares` | [string](#string) |  | total_shares represents the total shares of the pool |






<a name="gert.swap.v1beta1.QueryDepositsRequest"></a>

### QueryDepositsRequest
QueryDepositsRequest is the request type for the Query/Deposits RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `owner` | [string](#string) |  | owner optionally filters deposits by owner |
| `pool_id` | [string](#string) |  | pool_id optionally fitlers deposits by pool id |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="gert.swap.v1beta1.QueryDepositsResponse"></a>

### QueryDepositsResponse
QueryDepositsResponse is the response type for the Query/Deposits RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `deposits` | [DepositResponse](#gert.swap.v1beta1.DepositResponse) | repeated | deposits returns the deposits matching the requested parameters |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="gert.swap.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest defines the request type for querying x/swap parameters.






<a name="gert.swap.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse defines the response type for querying x/swap parameters.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#gert.swap.v1beta1.Params) |  | params represents the swap module parameters |






<a name="gert.swap.v1beta1.QueryPoolsRequest"></a>

### QueryPoolsRequest
QueryPoolsRequest is the request type for the Query/Pools RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pool_id` | [string](#string) |  | pool_id filters pools by id |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="gert.swap.v1beta1.QueryPoolsResponse"></a>

### QueryPoolsResponse
QueryPoolsResponse is the response type for the Query/Pools RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pools` | [PoolResponse](#gert.swap.v1beta1.PoolResponse) | repeated | pools represents returned pools |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="gert.swap.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service for swap module

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#gert.swap.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#gert.swap.v1beta1.QueryParamsResponse) | Params queries all parameters of the swap module. | GET|/gert/swap/v1beta1/params|
| `Pools` | [QueryPoolsRequest](#gert.swap.v1beta1.QueryPoolsRequest) | [QueryPoolsResponse](#gert.swap.v1beta1.QueryPoolsResponse) | Pools queries pools based on pool ID | GET|/gert/swap/v1beta1/pools|
| `Deposits` | [QueryDepositsRequest](#gert.swap.v1beta1.QueryDepositsRequest) | [QueryDepositsResponse](#gert.swap.v1beta1.QueryDepositsResponse) | Deposits queries deposit details based on owner address and pool | GET|/gert/swap/v1beta1/deposits|

 <!-- end services -->



<a name="gert/swap/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gert/swap/v1beta1/tx.proto



<a name="gert.swap.v1beta1.MsgDeposit"></a>

### MsgDeposit
MsgDeposit represents a message for depositing liquidity into a pool


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `depositor` | [string](#string) |  | depositor represents the address to deposit funds from |
| `token_a` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | token_a represents one token of deposit pair |
| `token_b` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | token_b represents one token of deposit pair |
| `slippage` | [string](#string) |  | slippage represents the max decimal percentage price change |
| `deadline` | [int64](#int64) |  | deadline represents the unix timestamp to complete the deposit by |






<a name="gert.swap.v1beta1.MsgDepositResponse"></a>

### MsgDepositResponse
MsgDepositResponse defines the Msg/Deposit response type.






<a name="gert.swap.v1beta1.MsgSwapExactForTokens"></a>

### MsgSwapExactForTokens
MsgSwapExactForTokens represents a message for trading exact coinA for coinB


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `requester` | [string](#string) |  | represents the address swaping the tokens |
| `exact_token_a` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | exact_token_a represents the exact amount to swap for token_b |
| `token_b` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | token_b represents the desired token_b to swap for |
| `slippage` | [string](#string) |  | slippage represents the maximum change in token_b allowed |
| `deadline` | [int64](#int64) |  | deadline represents the unix timestamp to complete the swap by |






<a name="gert.swap.v1beta1.MsgSwapExactForTokensResponse"></a>

### MsgSwapExactForTokensResponse
MsgSwapExactForTokensResponse defines the Msg/SwapExactForTokens response
type.






<a name="gert.swap.v1beta1.MsgSwapForExactTokens"></a>

### MsgSwapForExactTokens
MsgSwapForExactTokens represents a message for trading coinA for an exact
coinB


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `requester` | [string](#string) |  | represents the address swaping the tokens |
| `token_a` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | token_a represents the desired token_a to swap for |
| `exact_token_b` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | exact_token_b represents the exact token b amount to swap for token a |
| `slippage` | [string](#string) |  | slippage represents the maximum change in token_a allowed |
| `deadline` | [int64](#int64) |  | deadline represents the unix timestamp to complete the swap by |






<a name="gert.swap.v1beta1.MsgSwapForExactTokensResponse"></a>

### MsgSwapForExactTokensResponse
MsgSwapForExactTokensResponse defines the Msg/SwapForExactTokensResponse
response type.






<a name="gert.swap.v1beta1.MsgWithdraw"></a>

### MsgWithdraw
MsgWithdraw represents a message for withdrawing liquidity from a pool


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `from` | [string](#string) |  | from represents the address we are withdrawing for |
| `shares` | [string](#string) |  | shares represents the amount of shares to withdraw |
| `min_token_a` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | min_token_a represents the minimum a token to withdraw |
| `min_token_b` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | min_token_a represents the minimum a token to withdraw |
| `deadline` | [int64](#int64) |  | deadline represents the unix timestamp to complete the withdraw by |






<a name="gert.swap.v1beta1.MsgWithdrawResponse"></a>

### MsgWithdrawResponse
MsgWithdrawResponse defines the Msg/Withdraw response type.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="gert.swap.v1beta1.Msg"></a>

### Msg
Msg defines the swap Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Deposit` | [MsgDeposit](#gert.swap.v1beta1.MsgDeposit) | [MsgDepositResponse](#gert.swap.v1beta1.MsgDepositResponse) | Deposit defines a method for depositing liquidity into a pool | |
| `Withdraw` | [MsgWithdraw](#gert.swap.v1beta1.MsgWithdraw) | [MsgWithdrawResponse](#gert.swap.v1beta1.MsgWithdrawResponse) | Withdraw defines a method for withdrawing liquidity into a pool | |
| `SwapExactForTokens` | [MsgSwapExactForTokens](#gert.swap.v1beta1.MsgSwapExactForTokens) | [MsgSwapExactForTokensResponse](#gert.swap.v1beta1.MsgSwapExactForTokensResponse) | SwapExactForTokens represents a message for trading exact coinA for coinB | |
| `SwapForExactTokens` | [MsgSwapForExactTokens](#gert.swap.v1beta1.MsgSwapForExactTokens) | [MsgSwapForExactTokensResponse](#gert.swap.v1beta1.MsgSwapForExactTokensResponse) | SwapForExactTokens represents a message for trading coinA for an exact coinB | |

 <!-- end services -->



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

