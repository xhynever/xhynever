<!-- 主要参考go-solana和solana-go模块 -->


具体的accountinfo
参数token代币地址。解析代币的铸造者，代币的供应量，小数点。

(*rpc.GetAccountInfoResult)(0xc000119a80)({
 RPCContext: (rpc.RPCContext) {
  Context: (rpc.Context) {
   Slot: (uint64) 323589628
  }
 },
 Value: (*rpc.Account)(0xc0002a2440)({
  Lamports: (uint64) 840695717699,
  Owner: (solana.PublicKey) (len=32 cap=32) TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA,
  Data: (*rpc.DataBytesOrJSON)(0xc0000d8000)({
   rawDataEncoding: (solana.EncodingType) (len=6) "base64",
   asDecodedBinary: (solana.Data) AAAAAFnkpzMo+KIHXFu0C7POimfWZAwz81Y+ImohwO+lC39oDJDkCxqAIwAGAQAAAABZ5KczKPiiB1xbtAuzzopn1mQMM/NWPiJqIcDvpQt/aA==,
   asJSON: (json.RawMessage) <nil>
  }),
  Executable: (bool) false,
  RentEpoch: (*big.Int)(0xc00013f4c0)(18446744073709551615)
 })
})






上述片段中，一般只解析data
bin "github.com/gagliardetto/binary"
var mint token.Mint
err = bin.NewBinDecoder(resp.Value.Data.GetBinary()).Decode(&mint)



获得账户信息。
(*rpc.GetBalanceResult)(0xc0003d2370)({
 RPCContext: (rpc.RPCContext) {
  Context: (rpc.Context) {
   Slot: (uint64) 364119624
  }
 },
 Value: (uint64) 2039280
})
(uint64) 2039280

var lamportsOnAccount = new(big.Float).SetUint64(uint64(out.Value))
只解析value





