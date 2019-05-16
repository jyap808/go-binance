package binance

type Assets struct {
	Code              string  `json:"assetCode"`
	Name              string  `json:"assetName"`
	WithdrawalFee     float64 `json:"transactionFee"`
	MinimumWithdrawal float64 `json:"minProductWithdraw,string"`
	Confirmations     int     `json:"confirmTimes,string"`
	DepositEnabled    bool    `json:"enableCharge"`
	WithdrawalEnabled bool    `json:"enableWithdraw"`
}

// [{"id":"10","assetCode":"BNB","assetName":"Binance Coin","unit":"","transactionFee":0.005,"commissionRate":0.0,"freeAuditWithdrawAmt":20000.0,"freeUserChargeAmount":28156.0,"minProductWithdraw":"0.100000000000000000","withdrawIntegerMultiple":"0E-18","confirmTimes":"1","chargeLockConfirmTimes":null,"createTime":null,"test":0,"url":"https://explorer.binance.org/tx/","addressUrl":"https://explorer.binance.org/address/","blockUrl":"https://explorer.binance.org/block/","enableCharge":true,"enableWithdraw":true,"regEx":"^(bnb1)[0-9a-z]{38}$","regExTag":"^[0-9A-Za-z\\-_]{1,120}$","gas":1.0,"parentCode":"BNB","isLegalMoney":false,"reconciliationAmount":20.0,"seqNum":"0","chineseName":"Binance Coin","cnLink":"https://binance.zendesk.com/hc/zh-cn/articles/115000497111-%E5%B8%81%E5%AE%89%E5%B8%81-BNB-","enLink":"https://binance.zendesk.com/hc/en-us/articles/115000497111-Binance-Coin-BNB-","logoUrl":"https://ex.bnbstatic.com/images/20170912/image_1505205843840.png","fullLogoUrl":"https://ex.bnbstatic.com/images/20170912/image_1505205843840.png","forceStatus":true,"resetAddressStatus":false,"chargeDescCn":null,"chargeDescEn":null,"assetLabel":null,"sameAddress":true,"depositTipStatus":false,"dynamicFeeStatus":false,"depositTipEn":null,"depositTipCn":null,"assetLabelEn":null,"supportMarket":null,"feeReferenceAsset":"","feeRate":null,"feeDigit":2,"assetDigit":8,"legalMoney":false}
