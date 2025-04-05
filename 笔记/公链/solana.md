solana的账户模型:
solana账户4类。eth账户2类，个人账户和合约账户



普通账户（User Accounts）：这类账户主要用于持有SOL代币和其他资产，并由用户的私钥控制。用户可以通过这些账户发起交易或与其他智能合约交互2。
程序账户（Program Accounts）：也称为智能合约账户，它们包含可执行代码，但只能由Solana运行时环境执行。程序账户不持有状态信息，而是无状态的设计，这意味着每次调用时都需要重新加载必要的上下文2。
数据账户（Data Accounts）：用于存储状态和数据，可以被程序账户读取和写入。数据账户是由程序创建并用来持久化应用逻辑之外的信息，比如用户的余额或者其他业务相关的元数据2。
原生账户（Native Accounts）：指的是Solana自带的一些特殊功能账户，例如“System”、“Stake”以及“Vote”，它们服务于特定的目的，如质押投票权等3。



账户结构：
kv 数据对。可以理解账户的结构是一个表。     比较关键的，账户最大10MB
Public Key（公钥）：类似于以太坊地址，用于唯一识别账户。   全局状态树    地址查找表
Lamports：Solana的最小单位，表示账户中的余额。
Owner：指定拥有该账户的程序公钥（程序ID），只有所有者能够修改账户内的数据。
Executable：布尔值标志，表明账户是否包含可执行程序代码。
Data：存储在账户中的二进制数据，可以是任意格式，取决于应用程序的需求2。




数据帐户
Solana程序是“无状态的”，这意味着程序帐户仅包含程序的可执行字节码。若要存储和修改其他数据，必须创建新帐户。这些帐户通常称为“数据帐户”。

数据帐户可以存储所有者程序代码中定义的任何任意数据。

Data Account
数据帐户
请注意，只有系统程序可以创建新帐户。一旦系统程序创建了一个帐户，它就可以将新帐户的所有权转移到另一个程序。

换句话说，为自定义程序创建数据帐户需要两个步骤：

调用系统程序创建一个帐户，然后将所有权转移给自定义程序

调用现在拥有该帐户的自定义程序，然后初始化程序代码中定义的帐户数据

此数据帐户创建过程通常抽象为单个步骤，但了解基础过程很有帮助。


go-solana-sdk

账户生成：  1.获得随机熵，也叫seed   32位。    
2.由seed生成privateKey。seed sha512获得64位，用前32位在椭圆曲线上生成s点，由s映射生成A点。A点转32位公钥。
privateKey=seed+publicKey


私钥字符串88位 

初始化token账户的关键。
mint账户是token地址。
关联账户owner，生成ate。拥有sol的账户。
为创建账户，需要支付最小的费用。




<!-- token创建 -->





1.创建mit账户。需要feePayer，alice
调用函数system.CreateAccount，token.InitializeMint
包含参数，Decimals，Mint，MintAuth
获得mint地址。
2.创建token关联账户，mint地址，feePayer，alice
方法一：
种子映射，通过mint地址，alice地址获得ate地址。种子变，则ate地址变
common.FindAssociatedTokenAddress(alice.PublicKey, mintPubkey)
调用函数associated_token_account.Create
方法二：
随机创建，获得随机账户。NewAccount
调用system.CreateAccount，token.InitializeAccount
3.mint-to
需要feePayer，alice，ate地址
调用token.MintToChecked，参数：Mint，Auth，Signers，To，Amount
向ate地址创建指定数量的代币。
4.token-transfer
调用token.TransferChecked，参数From，To，Mint，Auth，Amount



