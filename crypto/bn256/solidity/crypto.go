// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package solidity

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// SolidityABI is the input ABI used to generate the binding from.
const SolidityABI = "[{\"inputs\":[{\"internalType\":\"uint256[2][]\",\"name\":\"sigs\",\"type\":\"uint256[2][]\"},{\"internalType\":\"uint256[]\",\"name\":\"indices\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"invArray\",\"type\":\"uint256[]\"}],\"name\":\"AggregateSignatures\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"HashToG1\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"h\",\"type\":\"uint256[2]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2][]\",\"name\":\"pointsG1\",\"type\":\"uint256[2][]\"},{\"internalType\":\"uint256[]\",\"name\":\"indices\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"invArray\",\"type\":\"uint256[]\"}],\"name\":\"LagrangeInterpolationG1\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"privK\",\"type\":\"uint256\"}],\"name\":\"Sign\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"sig\",\"type\":\"uint256[2]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"uint256[2]\",\"name\":\"sig\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[4]\",\"name\":\"pubK\",\"type\":\"uint256[4]\"}],\"name\":\"Verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"v\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"input\",\"type\":\"uint256[2]\"}],\"name\":\"safeSigningPoint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// SolidityBin is the compiled bytecode used for deploying new contracts.
var SolidityBin = "0x608060405234801561001057600080fd5b5061196a806100206000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c8063063675871461006757806316534acd146101745780636bdf477c1461038e57806395add79c1461043457806396d95a6f146104d8578063f022e06114610523575b600080fd5b610160600480360360e081101561007d57600080fd5b810190602081018135600160201b81111561009757600080fd5b8201836020820111156100a957600080fd5b803590602001918460018302840111600160201b831117156100ca57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080518082018252939695948181019493509150600290839083908082843760009201919091525050604080516080818101909252929594938181019392509060049083908390808284376000920191909152509194506107029350505050565b604080519115158252519081900360200190f35b6103536004803603608081101561018a57600080fd5b810190602081018135600160201b8111156101a457600080fd5b8201836020820111156101b657600080fd5b803590602001918460408302840111600160201b831117156101d757600080fd5b9190808060200260200160405190810160405280939291908181526020016000905b8282101561023757604080518082018252908084028701906002908390839080828437600092019190915250505081526001909101906020016101f9565b50939695946020810194503592505050600160201b81111561025857600080fd5b82018360208201111561026a57600080fd5b803590602001918460208302840111600160201b8311171561028b57600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092958435959094909350604081019250602001359050600160201b8111156102e257600080fd5b8201836020820111156102f457600080fd5b803590602001918460208302840111600160201b8311171561031557600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610881945050505050565b6040518082600260200280838360005b8381101561037b578181015183820152602001610363565b5050505090500191505060405180910390f35b610353600480360360408110156103a457600080fd5b810190602081018135600160201b8111156103be57600080fd5b8201836020820111156103d057600080fd5b803590602001918460018302840111600160201b831117156103f157600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250610a71915050565b6103536004803603602081101561044a57600080fd5b810190602081018135600160201b81111561046457600080fd5b82018360208201111561047657600080fd5b803590602001918460018302840111600160201b8311171561049757600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610aaf945050505050565b610160600480360360408110156104ee57600080fd5b60408051808201825291830192918183019183906002908390839080828437600092019190915250919450610bcc9350505050565b6103536004803603608081101561053957600080fd5b810190602081018135600160201b81111561055357600080fd5b82018360208201111561056557600080fd5b803590602001918460408302840111600160201b8311171561058657600080fd5b9190808060200260200160405190810160405280939291908181526020016000905b828210156105e657604080518082018252908084028701906002908390839080828437600092019190915250505081526001909101906020016105a8565b50939695946020810194503592505050600160201b81111561060757600080fd5b82018360208201111561061957600080fd5b803590602001918460208302840111600160201b8311171561063a57600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092958435959094909350604081019250602001359050600160201b81111561069157600080fd5b8201836020820111156106a357600080fd5b803590602001918460208302840111600160201b831117156106c457600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610bf4945050505050565b600061070c6116e0565b61071585610aaf565b90506108786040518061018001604052808660006002811061073357fe5b602002015181526020018660016002811061074a57fe5b602002015181526020017f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c281526020017f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed81526020017f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec81526020017f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d8152602001836000600281106107f957fe5b602002015181526020018360016002811061081057fe5b602002015181526020018560006004811061082757fe5b602002015181526020018560016004811061083e57fe5b602002015181526020018560026004811061085557fe5b602002015181526020018560036004811061086c57fe5b60200201519052610cec565b95945050505050565b6108896116e0565b83518551146108c95760405162461bcd60e51b815260040180806020018281038252602c815260200180611832602c913960400191505060405180910390fd5b6108d16116e0565b60008082526020820181905280808080806108ea6116e0565b600096505b8b51871015610a60578b878151811061090457fe5b602002602001015194508a87111561091b57610a60565b60019250600095505b8b5186101561098a578b868151811061093957fe5b602002602001015193508a8611156109505761098a565b8484141561095d5761097f565b61096884868c610d62565b915060008051602061185e83398151915282840992505b600190950194610924565b8c878151811061099657fe5b602002602001015190506109e56040518060600160405280836000600281106109bb57fe5b60200201518152602001836001600281106109d257fe5b6020020151815260200185815250610e2c565b9050610a5360405180608001604052808a600060028110610a0257fe5b602002015181526020018a600160028110610a1957fe5b6020020151815260200183600060028110610a3057fe5b6020020151815260200183600160028110610a4757fe5b60200201519052610e84565b97506001909601956108ef565b50959b9a5050505050505050505050565b610a796116e0565b610a816116e0565b610a8a84610aaf565b9050610aa76040518060600160405280836000600281106109bb57fe5b949350505050565b610ab76116e0565b6000610ac88382600160f81b610eec565b90506000610adf84600160f91b600360f81b610eec565b9050610ae96116e0565b610af283611063565b9050610afc6116e0565b610b0583611063565b9050610b39604051806080016040528084600060028110610b2257fe5b6020020151815260200184600160028110610a1957fe5b9450610b44856113a0565b610b7f5760405162461bcd60e51b815260040180806020018281038252602981526020018061171d6029913960400191505060405180910390fd5b610b8885610bcc565b610bc35760405162461bcd60e51b815260040180806020018281038252602a81526020018061187e602a913960400191505060405180910390fd5b50505050919050565b80516000901580610bde575081516001145b15610beb57506000610bef565b5060015b919050565b610bfc6116e0565b8351855114610c3c5760405162461bcd60e51b81526004018080602001828103825260358152602001806118a86035913960400191505060405180910390fd5b82855111610c7b5760405162461bcd60e51b815260040180806020018281038252603a8152602001806117ce603a913960400191505060405180910390fd5b6000610c86856113e0565b9050610c928382611450565b610ccd5760405162461bcd60e51b815260040180806020018281038252602a815260200180611808602a913960400191505060405180910390fd5b610cd56116e0565b610ce187878787610881565b979650505050505050565b6000610cf66116fe565b60006020826101808660085afa905080610d57576040805162461bcd60e51b815260206004820152601d60248201527f656c6c69707469632063757276652070616972696e67206661696c6564000000604482015290519081900360640190fd5b505160011492915050565b600082841415610da35760405162461bcd60e51b81526004018080602001828103825260348152602001806118dd6034913960400191505060405180910390fd5b83600084821115610db75750838503610df0565b60008051602061185e8339815191527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000083099150508484035b836001820381518110610dff57fe5b6020026020010151905060008051602061185e83398151915280610e1f57fe5b8183099695505050505050565b610e346116e0565b600060408260608560075afa905080610e7e5760405162461bcd60e51b81526004018080602001828103825260248152602001806119116024913960400191505060405180910390fd5b50919050565b610e8c6116e0565b600060408260808560065afa905080610e7e576040805162461bcd60e51b815260206004820152601e60248201527f656c6c6970746963206375727665206164646974696f6e206661696c65640000604482015290519081900360640190fd5b600080838560405160200180836001600160f81b0319166001600160f81b031916815260010182805190602001908083835b60208310610f3d5780518252601f199092019160209182019101610f1e565b6001836020036101000a038019825116818451168082178552505050505050905001925050506040516020818303038152906040528051906020012060001c90506000838660405160200180836001600160f81b0319166001600160f81b031916815260010182805190602001908083835b60208310610fce5780518252601f199092019160209182019101610faf565b6001836020036101000a038019825116818451168082178552505050505050905001925050506040516020818303038152906040528051906020012060001c90506000805160206117ae8339815191528061102557fe5b816000805160206117ae8339815191527f0e0a77c19a07df2f666ea36f7879462c0a78eb28f5c70b3dd35d438dc58f0d9d8509089695505050505050565b61106b6116e0565b600080806000805160206117ae83398151915285860992506000805160206117ae8339815191526004840891506000805160206117ae83398151915282840990506110b581611507565b905060006000805160206117ae83398151915283840990506000805160206117ae83398151915283820990506000805160206117ae833981519152848509935060006000805160206117ae8339815191528577b3c4d79d41a91759a9e4c7e359b6b89eaec68e62effffffd0990506000805160206117ae833981519152838209905061114081611548565b90506000805160206117ae8339815191527759e26bcea0d48bacd4f263f1acdb5c4f5763473177fffffe8208905060006000805160206117ae83398151915260018308905061118e81611548565b905060006000805160206117ae833981519152847f2042def740cbc01bd03583cf0100e593ba56470b9af68708d2c05d64905353850990506000805160206117ae83398151915285820990506111e381611548565b90506000805160206117ae83398151915260018208905060006000805160206117ae83398151915284850990506000805160206117ae83398151915284820990506000805160206117ae83398151915260038208905060006112448261156b565b90506000805160206117ae83398151915284850991506000805160206117ae83398151915284830991506000805160206117ae833981519152600383089150600061128e8361156b565b9050600160046002198301600019850102058101906000908214156112b45750866112c7565b81600214156112c45750856112c7565b50845b6000805160206117ae83398151915281820994506000805160206117ae83398151915281860994506000805160206117ae83398151915260038608945061130d856115c9565b9450600061131a8f611604565b90506000805160206117ae833981519152818709955061134d6040518060400160405280848152602001888152506113a0565b6113885760405162461bcd60e51b815260040180806020018281038252602481526020018061178a6024913960400191505060405180910390fd5b508c5250505060208901525095979650505050505050565b80516000906000805160206117ae833981519152906003908290819080098551090860208301516000805160206117ae8339815191529080091492915050565b60008151600014156113f457506000610bef565b6000808360008151811061140457fe5b602002602001015190506000600190505b84518110156114485784818151811061142a57fe5b6020026020010151925081831115611440578291505b600101611415565b509392505050565b600080808060018561145e57fe5b86516001870311156114a15760405162461bcd60e51b81526004018080602001828103825260448152602001806117466044913960600191505060405180910390fd5b600193505b858410156114fd578660018503815181106114bd57fe5b6020026020010151925060008051602061185e833981519152806114dd57fe5b8385099150816001146114f2575060006114fd565b6001909301926114a6565b9695505050505050565b6000611542827f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd456000805160206117ae833981519152611655565b92915050565b60008161155757506000610bef565b506000805160206117ae8339815191520390565b6000806115a7837f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea36000805160206117ae833981519152611655565b905080156115c057600181600116600202039150610e7e565b50600092915050565b6000611542827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f526000805160206117ae833981519152611655565b60017f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea3821115610bef57507f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd46919050565b60008060405160208152602080820152602060408201528560608201528460808201528360a082015260208160c08360055afa90519250905080611448576040805162461bcd60e51b815260206004820152601d60248201527f6d6f64756c6172206578706f6e656e74696174696f6e2066616c696564000000604482015290519081900360640190fd5b60405180604001604052806002906020820280368337509192915050565b6040518060200160405280600190602082028036833750919291505056fe496e76616c6964206861736820706f696e743a206e6f74206f6e20656c6c6970746963206375727665636865636b496e7665727365733a20696e73756666696369656e7420696e76657273657320666f722067726f7570207369676e61747572652063616c63756c6174696f6e496e76616c696420706f696e743a206e6f74206f6e20656c6c697074696320637572766530644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd474661696c656420746f206d656574207265717569726564206e756d626572206f66207369676e61747572657320666f72207468726573686f6c64696e76417272617920646f6573206e6f7420696e636c75646520636f727265637420696e7665727365734d69736d61746368206265747765656e20706f696e7473473120616e6420696e64696365732061727261797330644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000144616e6765726f7573206861736820706f696e743a206e6f74207361666520666f72207369676e696e674d69736d61746368206265747765656e206c656e677468206f66207369676e61747572657320616e6420696e6465782061727261794d7573742068617665206b20213d206a207768656e20636f6d707574696e6720526a207061727469616c20636f6e7374616e7473656c6c6970746963206375727665206d756c7469706c69636174696f6e206661696c6564a2646970667358221220a2460524c27cc11f4710e74d79856d5e7c2052f2c724bd2cc381e71c516dcb9f64736f6c63430006090033"

// DeploySolidity deploys a new Ethereum contract, binding an instance of Solidity to it.
func DeploySolidity(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Solidity, error) {
	parsed, err := abi.JSON(strings.NewReader(SolidityABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SolidityBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Solidity{SolidityCaller: SolidityCaller{contract: contract}, SolidityTransactor: SolidityTransactor{contract: contract}, SolidityFilterer: SolidityFilterer{contract: contract}}, nil
}

// Solidity is an auto generated Go binding around an Ethereum contract.
type Solidity struct {
	SolidityCaller     // Read-only binding to the contract
	SolidityTransactor // Write-only binding to the contract
	SolidityFilterer   // Log filterer for contract events
}

// SolidityCaller is an auto generated read-only Go binding around an Ethereum contract.
type SolidityCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolidityTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SolidityTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolidityFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SolidityFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SoliditySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SoliditySession struct {
	Contract     *Solidity         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SolidityCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SolidityCallerSession struct {
	Contract *SolidityCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SolidityTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SolidityTransactorSession struct {
	Contract     *SolidityTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SolidityRaw is an auto generated low-level Go binding around an Ethereum contract.
type SolidityRaw struct {
	Contract *Solidity // Generic contract binding to access the raw methods on
}

// SolidityCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SolidityCallerRaw struct {
	Contract *SolidityCaller // Generic read-only contract binding to access the raw methods on
}

// SolidityTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SolidityTransactorRaw struct {
	Contract *SolidityTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSolidity creates a new instance of Solidity, bound to a specific deployed contract.
func NewSolidity(address common.Address, backend bind.ContractBackend) (*Solidity, error) {
	contract, err := bindSolidity(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Solidity{SolidityCaller: SolidityCaller{contract: contract}, SolidityTransactor: SolidityTransactor{contract: contract}, SolidityFilterer: SolidityFilterer{contract: contract}}, nil
}

// NewSolidityCaller creates a new read-only instance of Solidity, bound to a specific deployed contract.
func NewSolidityCaller(address common.Address, caller bind.ContractCaller) (*SolidityCaller, error) {
	contract, err := bindSolidity(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SolidityCaller{contract: contract}, nil
}

// NewSolidityTransactor creates a new write-only instance of Solidity, bound to a specific deployed contract.
func NewSolidityTransactor(address common.Address, transactor bind.ContractTransactor) (*SolidityTransactor, error) {
	contract, err := bindSolidity(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SolidityTransactor{contract: contract}, nil
}

// NewSolidityFilterer creates a new log filterer instance of Solidity, bound to a specific deployed contract.
func NewSolidityFilterer(address common.Address, filterer bind.ContractFilterer) (*SolidityFilterer, error) {
	contract, err := bindSolidity(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SolidityFilterer{contract: contract}, nil
}

// bindSolidity binds a generic wrapper to an already deployed contract.
func bindSolidity(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SolidityABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Solidity *SolidityRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Solidity.Contract.SolidityCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Solidity *SolidityRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Solidity.Contract.SolidityTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Solidity *SolidityRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Solidity.Contract.SolidityTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Solidity *SolidityCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Solidity.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Solidity *SolidityTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Solidity.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Solidity *SolidityTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Solidity.Contract.contract.Transact(opts, method, params...)
}

// AggregateSignatures is a free data retrieval call binding the contract method 0xf022e061.
//
// Solidity: function AggregateSignatures(uint256[2][] sigs, uint256[] indices, uint256 threshold, uint256[] invArray) view returns(uint256[2])
func (_Solidity *SolidityCaller) AggregateSignatures(opts *bind.CallOpts, sigs [][2]*big.Int, indices []*big.Int, threshold *big.Int, invArray []*big.Int) ([2]*big.Int, error) {
	var (
		ret0 = new([2]*big.Int)
	)
	out := ret0
	err := _Solidity.contract.Call(opts, out, "AggregateSignatures", sigs, indices, threshold, invArray)
	return *ret0, err
}

// AggregateSignatures is a free data retrieval call binding the contract method 0xf022e061.
//
// Solidity: function AggregateSignatures(uint256[2][] sigs, uint256[] indices, uint256 threshold, uint256[] invArray) view returns(uint256[2])
func (_Solidity *SoliditySession) AggregateSignatures(sigs [][2]*big.Int, indices []*big.Int, threshold *big.Int, invArray []*big.Int) ([2]*big.Int, error) {
	return _Solidity.Contract.AggregateSignatures(&_Solidity.CallOpts, sigs, indices, threshold, invArray)
}

// AggregateSignatures is a free data retrieval call binding the contract method 0xf022e061.
//
// Solidity: function AggregateSignatures(uint256[2][] sigs, uint256[] indices, uint256 threshold, uint256[] invArray) view returns(uint256[2])
func (_Solidity *SolidityCallerSession) AggregateSignatures(sigs [][2]*big.Int, indices []*big.Int, threshold *big.Int, invArray []*big.Int) ([2]*big.Int, error) {
	return _Solidity.Contract.AggregateSignatures(&_Solidity.CallOpts, sigs, indices, threshold, invArray)
}

// HashToG1 is a free data retrieval call binding the contract method 0x95add79c.
//
// Solidity: function HashToG1(bytes message) view returns(uint256[2] h)
func (_Solidity *SolidityCaller) HashToG1(opts *bind.CallOpts, message []byte) ([2]*big.Int, error) {
	var (
		ret0 = new([2]*big.Int)
	)
	out := ret0
	err := _Solidity.contract.Call(opts, out, "HashToG1", message)
	return *ret0, err
}

// HashToG1 is a free data retrieval call binding the contract method 0x95add79c.
//
// Solidity: function HashToG1(bytes message) view returns(uint256[2] h)
func (_Solidity *SoliditySession) HashToG1(message []byte) ([2]*big.Int, error) {
	return _Solidity.Contract.HashToG1(&_Solidity.CallOpts, message)
}

// HashToG1 is a free data retrieval call binding the contract method 0x95add79c.
//
// Solidity: function HashToG1(bytes message) view returns(uint256[2] h)
func (_Solidity *SolidityCallerSession) HashToG1(message []byte) ([2]*big.Int, error) {
	return _Solidity.Contract.HashToG1(&_Solidity.CallOpts, message)
}

// LagrangeInterpolationG1 is a free data retrieval call binding the contract method 0x16534acd.
//
// Solidity: function LagrangeInterpolationG1(uint256[2][] pointsG1, uint256[] indices, uint256 threshold, uint256[] invArray) view returns(uint256[2])
func (_Solidity *SolidityCaller) LagrangeInterpolationG1(opts *bind.CallOpts, pointsG1 [][2]*big.Int, indices []*big.Int, threshold *big.Int, invArray []*big.Int) ([2]*big.Int, error) {
	var (
		ret0 = new([2]*big.Int)
	)
	out := ret0
	err := _Solidity.contract.Call(opts, out, "LagrangeInterpolationG1", pointsG1, indices, threshold, invArray)
	return *ret0, err
}

// LagrangeInterpolationG1 is a free data retrieval call binding the contract method 0x16534acd.
//
// Solidity: function LagrangeInterpolationG1(uint256[2][] pointsG1, uint256[] indices, uint256 threshold, uint256[] invArray) view returns(uint256[2])
func (_Solidity *SoliditySession) LagrangeInterpolationG1(pointsG1 [][2]*big.Int, indices []*big.Int, threshold *big.Int, invArray []*big.Int) ([2]*big.Int, error) {
	return _Solidity.Contract.LagrangeInterpolationG1(&_Solidity.CallOpts, pointsG1, indices, threshold, invArray)
}

// LagrangeInterpolationG1 is a free data retrieval call binding the contract method 0x16534acd.
//
// Solidity: function LagrangeInterpolationG1(uint256[2][] pointsG1, uint256[] indices, uint256 threshold, uint256[] invArray) view returns(uint256[2])
func (_Solidity *SolidityCallerSession) LagrangeInterpolationG1(pointsG1 [][2]*big.Int, indices []*big.Int, threshold *big.Int, invArray []*big.Int) ([2]*big.Int, error) {
	return _Solidity.Contract.LagrangeInterpolationG1(&_Solidity.CallOpts, pointsG1, indices, threshold, invArray)
}

// Sign is a free data retrieval call binding the contract method 0x6bdf477c.
//
// Solidity: function Sign(bytes message, uint256 privK) view returns(uint256[2] sig)
func (_Solidity *SolidityCaller) Sign(opts *bind.CallOpts, message []byte, privK *big.Int) ([2]*big.Int, error) {
	var (
		ret0 = new([2]*big.Int)
	)
	out := ret0
	err := _Solidity.contract.Call(opts, out, "Sign", message, privK)
	return *ret0, err
}

// Sign is a free data retrieval call binding the contract method 0x6bdf477c.
//
// Solidity: function Sign(bytes message, uint256 privK) view returns(uint256[2] sig)
func (_Solidity *SoliditySession) Sign(message []byte, privK *big.Int) ([2]*big.Int, error) {
	return _Solidity.Contract.Sign(&_Solidity.CallOpts, message, privK)
}

// Sign is a free data retrieval call binding the contract method 0x6bdf477c.
//
// Solidity: function Sign(bytes message, uint256 privK) view returns(uint256[2] sig)
func (_Solidity *SolidityCallerSession) Sign(message []byte, privK *big.Int) ([2]*big.Int, error) {
	return _Solidity.Contract.Sign(&_Solidity.CallOpts, message, privK)
}

// Verify is a free data retrieval call binding the contract method 0x06367587.
//
// Solidity: function Verify(bytes message, uint256[2] sig, uint256[4] pubK) view returns(bool v)
func (_Solidity *SolidityCaller) Verify(opts *bind.CallOpts, message []byte, sig [2]*big.Int, pubK [4]*big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Solidity.contract.Call(opts, out, "Verify", message, sig, pubK)
	return *ret0, err
}

// Verify is a free data retrieval call binding the contract method 0x06367587.
//
// Solidity: function Verify(bytes message, uint256[2] sig, uint256[4] pubK) view returns(bool v)
func (_Solidity *SoliditySession) Verify(message []byte, sig [2]*big.Int, pubK [4]*big.Int) (bool, error) {
	return _Solidity.Contract.Verify(&_Solidity.CallOpts, message, sig, pubK)
}

// Verify is a free data retrieval call binding the contract method 0x06367587.
//
// Solidity: function Verify(bytes message, uint256[2] sig, uint256[4] pubK) view returns(bool v)
func (_Solidity *SolidityCallerSession) Verify(message []byte, sig [2]*big.Int, pubK [4]*big.Int) (bool, error) {
	return _Solidity.Contract.Verify(&_Solidity.CallOpts, message, sig, pubK)
}

// SafeSigningPoint is a free data retrieval call binding the contract method 0x96d95a6f.
//
// Solidity: function safeSigningPoint(uint256[2] input) pure returns(bool)
func (_Solidity *SolidityCaller) SafeSigningPoint(opts *bind.CallOpts, input [2]*big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Solidity.contract.Call(opts, out, "safeSigningPoint", input)
	return *ret0, err
}

// SafeSigningPoint is a free data retrieval call binding the contract method 0x96d95a6f.
//
// Solidity: function safeSigningPoint(uint256[2] input) pure returns(bool)
func (_Solidity *SoliditySession) SafeSigningPoint(input [2]*big.Int) (bool, error) {
	return _Solidity.Contract.SafeSigningPoint(&_Solidity.CallOpts, input)
}

// SafeSigningPoint is a free data retrieval call binding the contract method 0x96d95a6f.
//
// Solidity: function safeSigningPoint(uint256[2] input) pure returns(bool)
func (_Solidity *SolidityCallerSession) SafeSigningPoint(input [2]*big.Int) (bool, error) {
	return _Solidity.Contract.SafeSigningPoint(&_Solidity.CallOpts, input)
}
