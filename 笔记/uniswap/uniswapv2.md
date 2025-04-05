## 合约包含4个关键模块


### UniswapV2Factory
作用: 负责创建和管理交易对（资金池）。
主要功能:
constructor(address _feeToSetter): 初始化合约，设置手续费管理员地址。
createPair(address tokenA, address tokenB): 创建一个新的交易对。
setFeeTo(address _feeTo): 设置手续费接收地址。
setFeeToSetter(address _feeToSetter): 设置手续费管理员地址。
getPair(address tokenA, address tokenB): 获取已存在的交易对地址。
allPairs(uint256): 获取所有交易对的地址。
allPairsLength(): 获取所有交易对的数量。



### UniswapV2Pair
作用: 每个交易对都有一个对应的 UniswapV2Pair 合约，管理该交易对的流动性、交易和价格。
主要功能:
initialize(address _token0, address _token1): 初始化交易对。
mint(address to): 添加流动性并铸造 LP 代币。
burn(address to): 移除流动性并销毁 LP 代币。
swap(uint256 amount0Out, uint256 amount1Out, address to, bytes calldata data): 执行代币交换。
getReserves(): 获取交易对的储备量。
sync(): 同步储备量。



### UniswapV2ERC20
作用: 继承自 ERC20 标准，用于管理流动性提供者的 LP 代币。
主要功能:
totalSupply(): 获取总供应量。
balanceOf(address owner): 获取某个地址的余额。
transfer(address to, uint256 value): 转账。
approve(address spender, uint256 value): 授权。
transferFrom(address from, address to, uint256 value): 从一个地址转账到另一个地址。



### UniswapV2Router02
作用: 提供一系列接口，方便用户与核心合约进行交互，执行复杂的交易操作。
主要功能:
addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline): 添加流动性。
removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline): 移除流动性。
swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] calldata path, address to, uint256 deadline): 交换代币。
getAmountsOut(uint256 amountIn, address[] memory path): 获取预期输出量。
getAmountsIn(uint256 amountOut, address[] memory path): 获取预期输入量。



辅助合约和库
除了上述核心合约，Uniswap V2 还包含一些辅助合约和库，用于提供额外的功能和工具：

SafeMath
作用: 提供安全的数学运算，防止溢出和下溢。
主要功能:
add(uint256 a, uint256 b): 安全加法。
sub(uint256 a, uint256 b): 安全减法。
mul(uint256 a, uint256 b): 安全乘法。
div(uint256 a, uint256 b): 安全除法。
UQ112x112
作用: 提供固定点数运算，用于精确的价格计算。
主要功能:
encode(uint112 y): 将无符号整数编码为 UQ112x112 类型。
decode(uint224 x): 将 UQ112x112 类型解码为无符号整数。
uqdiv(uint224 x, uint112 y): 固定点数除法。
TransferHelper
作用: 提供安全的代币转移函数。
主要功能:
safeApprove(address token, address to, uint256 value): 安全授权。
safeTransfer(address token, address to, uint256 value): 安全转账。
safeTransferFrom(address token, address from, address to, uint256 value): 安全从一个地址转账到另一个地址。


## 创建币对合约过程
工厂合约是核心合约：

创建过程，
1.调用 createPair 函数:
用户或合约调用 UniswapV2Factory 合约的 createPair 函数，传入两个代币的地址（tokenA 和 tokenB）。
createPair 函数会检查两个代币地址是否相同、是否为零地址、以及该交易对是否已经存在。
2.创建 UniswapV2Pair 合约:
如果所有检查通过，createPair 函数会使用 create2 操作码创建一个新的 UniswapV2Pair 合约。
create2 操作码允许通过一个盐值（salt）来确定新合约的地址，从而使得地址可预测。
3.初始化 UniswapV2Pair 合约:
新创建的 UniswapV2Pair 合约会调用其 initialize 函数，传入两个代币的地址，完成初始化。
initialize 函数会设置 token0 和 token1，并确保合约准备好处理流动性管理和交易。
4.记录交易对:
UniswapV2Factory 合约会将新创建的交易对地址记录到 getPair 映射和 allPairs 数组中。
触发 PairCreated 事件，通知其他合约和用户新的交易对已创建。

