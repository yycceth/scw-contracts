// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package smartAccount

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// UserOperation is an auto generated low-level Go binding around an user-defined struct.
type UserOperation struct {
	Sender               common.Address
	Nonce                *big.Int
	InitCode             []byte
	CallData             []byte
	CallGasLimit         *big.Int
	VerificationGasLimit *big.Int
	PreVerificationGas   *big.Int
	MaxFeePerGas         *big.Int
	MaxPriorityFeePerGas *big.Int
	PaymasterAndData     []byte
	Signature            []byte
}

// SmartAccountMetaData contains all meta data concerning the SmartAccount contract.
var SmartAccountMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"anEntryPoint\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseImplementationCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"CallerIsNotAnEntryPoint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"CallerIsNotEntryPoint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"CallerIsNotEntryPointOrOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"CallerIsNotEntryPointOrSelf\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"CallerIsNotOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"CallerIsNotSelf\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DelegateCallsOnly\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EntryPointCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HandlerCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementationAddress\",\"type\":\"address\"}],\"name\":\"InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"MixedAuthFail\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"ModuleAlreadyEnabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"expectedModule\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"returnedModule\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"prevModule\",\"type\":\"address\"}],\"name\":\"ModuleAndPrevModuleMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"ModuleCannotBeZeroOrSentinel\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"ModuleNotEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ModulesAlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ModulesSetupExecutionFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerCanNotBeSelf\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerProvidedIsSame\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToZeroAddressAttempt\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valueLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"funcLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"operationLength\",\"type\":\"uint256\"}],\"name\":\"WrongBatchProvided\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"contractSignature\",\"type\":\"bytes\"}],\"name\":\"WrongContractSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"uintS\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"contractSignatureLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"signatureLength\",\"type\":\"uint256\"}],\"name\":\"WrongContractSignatureFormat\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"moduleAddressProvided\",\"type\":\"address\"}],\"name\":\"WrongValidationModule\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousHandler\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"handler\",\"type\":\"address\"}],\"name\":\"ChangedFallbackHandler\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"DisabledModule\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"EnabledModule\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"enumEnum.Operation\",\"name\":\"operation\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"txGas\",\"type\":\"uint256\"}],\"name\":\"ExecutionFailure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"ExecutionFromModuleFailure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"ExecutionFromModuleSuccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"enumEnum.Operation\",\"name\":\"operation\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"txGas\",\"type\":\"uint256\"}],\"name\":\"ExecutionSuccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldImplementation\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"ImplementationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"enumEnum.Operation\",\"name\":\"operation\",\"type\":\"uint8\"}],\"name\":\"ModuleTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SmartAccountReceivedNativeToken\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addDeposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prevModule\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"disableModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"enableModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"entryPoint\",\"outputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"to\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"value\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"},{\"internalType\":\"enumEnum.Operation[]\",\"name\":\"operations\",\"type\":\"uint8[]\"}],\"name\":\"execBatchTransactionFromModule\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"enumEnum.Operation\",\"name\":\"operation\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"txGas\",\"type\":\"uint256\"}],\"name\":\"execTransactionFromModule\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"enumEnum.Operation\",\"name\":\"operation\",\"type\":\"uint8\"}],\"name\":\"execTransactionFromModule\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"enumEnum.Operation\",\"name\":\"operation\",\"type\":\"uint8\"}],\"name\":\"execTransactionFromModuleReturnData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dest\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"func\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"dest\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"value\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"func\",\"type\":\"bytes[]\"}],\"name\":\"executeBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"dest\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"value\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"func\",\"type\":\"bytes[]\"}],\"name\":\"executeBatch_y6U\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dest\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"func\",\"type\":\"bytes\"}],\"name\":\"execute_ncC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFallbackHandler\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_handler\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_implementation\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"start\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pageSize\",\"type\":\"uint256\"}],\"name\":\"getModulesPaginated\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"array\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"next\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"handler\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"moduleSetupContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"moduleSetupData\",\"type\":\"bytes\"}],\"name\":\"init\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"isModuleEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint192\",\"name\":\"_key\",\"type\":\"uint192\"}],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"noncesDeprecated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ownerDeprecated\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"handler\",\"type\":\"address\"}],\"name\":\"setFallbackHandler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"setupContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"setupData\",\"type\":\"bytes\"}],\"name\":\"setupAndEnableModule\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_implementation\",\"type\":\"address\"}],\"name\":\"updateImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"callGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"missingAccountFunds\",\"type\":\"uint256\"}],\"name\":\"validateUserOp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"validationData\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawDepositTo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// SmartAccountABI is the input ABI used to generate the binding from.
// Deprecated: Use SmartAccountMetaData.ABI instead.
var SmartAccountABI = SmartAccountMetaData.ABI

// SmartAccount is an auto generated Go binding around an Ethereum contract.
type SmartAccount struct {
	SmartAccountCaller     // Read-only binding to the contract
	SmartAccountTransactor // Write-only binding to the contract
	SmartAccountFilterer   // Log filterer for contract events
}

// SmartAccountCaller is an auto generated read-only Go binding around an Ethereum contract.
type SmartAccountCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmartAccountTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SmartAccountTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmartAccountFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SmartAccountFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmartAccountSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SmartAccountSession struct {
	Contract     *SmartAccount     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SmartAccountCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SmartAccountCallerSession struct {
	Contract *SmartAccountCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// SmartAccountTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SmartAccountTransactorSession struct {
	Contract     *SmartAccountTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// SmartAccountRaw is an auto generated low-level Go binding around an Ethereum contract.
type SmartAccountRaw struct {
	Contract *SmartAccount // Generic contract binding to access the raw methods on
}

// SmartAccountCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SmartAccountCallerRaw struct {
	Contract *SmartAccountCaller // Generic read-only contract binding to access the raw methods on
}

// SmartAccountTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SmartAccountTransactorRaw struct {
	Contract *SmartAccountTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSmartAccount creates a new instance of SmartAccount, bound to a specific deployed contract.
func NewSmartAccount(address common.Address, backend bind.ContractBackend) (*SmartAccount, error) {
	contract, err := bindSmartAccount(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SmartAccount{SmartAccountCaller: SmartAccountCaller{contract: contract}, SmartAccountTransactor: SmartAccountTransactor{contract: contract}, SmartAccountFilterer: SmartAccountFilterer{contract: contract}}, nil
}

// NewSmartAccountCaller creates a new read-only instance of SmartAccount, bound to a specific deployed contract.
func NewSmartAccountCaller(address common.Address, caller bind.ContractCaller) (*SmartAccountCaller, error) {
	contract, err := bindSmartAccount(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SmartAccountCaller{contract: contract}, nil
}

// NewSmartAccountTransactor creates a new write-only instance of SmartAccount, bound to a specific deployed contract.
func NewSmartAccountTransactor(address common.Address, transactor bind.ContractTransactor) (*SmartAccountTransactor, error) {
	contract, err := bindSmartAccount(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SmartAccountTransactor{contract: contract}, nil
}

// NewSmartAccountFilterer creates a new log filterer instance of SmartAccount, bound to a specific deployed contract.
func NewSmartAccountFilterer(address common.Address, filterer bind.ContractFilterer) (*SmartAccountFilterer, error) {
	contract, err := bindSmartAccount(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SmartAccountFilterer{contract: contract}, nil
}

// bindSmartAccount binds a generic wrapper to an already deployed contract.
func bindSmartAccount(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SmartAccountMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SmartAccount *SmartAccountRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SmartAccount.Contract.SmartAccountCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SmartAccount *SmartAccountRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SmartAccount.Contract.SmartAccountTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SmartAccount *SmartAccountRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SmartAccount.Contract.SmartAccountTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SmartAccount *SmartAccountCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SmartAccount.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SmartAccount *SmartAccountTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SmartAccount.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SmartAccount *SmartAccountTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SmartAccount.Contract.contract.Transact(opts, method, params...)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_SmartAccount *SmartAccountCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SmartAccount.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_SmartAccount *SmartAccountSession) VERSION() (string, error) {
	return _SmartAccount.Contract.VERSION(&_SmartAccount.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_SmartAccount *SmartAccountCallerSession) VERSION() (string, error) {
	return _SmartAccount.Contract.VERSION(&_SmartAccount.CallOpts)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_SmartAccount *SmartAccountCaller) EntryPoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SmartAccount.contract.Call(opts, &out, "entryPoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_SmartAccount *SmartAccountSession) EntryPoint() (common.Address, error) {
	return _SmartAccount.Contract.EntryPoint(&_SmartAccount.CallOpts)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_SmartAccount *SmartAccountCallerSession) EntryPoint() (common.Address, error) {
	return _SmartAccount.Contract.EntryPoint(&_SmartAccount.CallOpts)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_SmartAccount *SmartAccountCaller) GetDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SmartAccount.contract.Call(opts, &out, "getDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_SmartAccount *SmartAccountSession) GetDeposit() (*big.Int, error) {
	return _SmartAccount.Contract.GetDeposit(&_SmartAccount.CallOpts)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_SmartAccount *SmartAccountCallerSession) GetDeposit() (*big.Int, error) {
	return _SmartAccount.Contract.GetDeposit(&_SmartAccount.CallOpts)
}

// GetFallbackHandler is a free data retrieval call binding the contract method 0x856dfd99.
//
// Solidity: function getFallbackHandler() view returns(address _handler)
func (_SmartAccount *SmartAccountCaller) GetFallbackHandler(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SmartAccount.contract.Call(opts, &out, "getFallbackHandler")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFallbackHandler is a free data retrieval call binding the contract method 0x856dfd99.
//
// Solidity: function getFallbackHandler() view returns(address _handler)
func (_SmartAccount *SmartAccountSession) GetFallbackHandler() (common.Address, error) {
	return _SmartAccount.Contract.GetFallbackHandler(&_SmartAccount.CallOpts)
}

// GetFallbackHandler is a free data retrieval call binding the contract method 0x856dfd99.
//
// Solidity: function getFallbackHandler() view returns(address _handler)
func (_SmartAccount *SmartAccountCallerSession) GetFallbackHandler() (common.Address, error) {
	return _SmartAccount.Contract.GetFallbackHandler(&_SmartAccount.CallOpts)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address _implementation)
func (_SmartAccount *SmartAccountCaller) GetImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SmartAccount.contract.Call(opts, &out, "getImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address _implementation)
func (_SmartAccount *SmartAccountSession) GetImplementation() (common.Address, error) {
	return _SmartAccount.Contract.GetImplementation(&_SmartAccount.CallOpts)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address _implementation)
func (_SmartAccount *SmartAccountCallerSession) GetImplementation() (common.Address, error) {
	return _SmartAccount.Contract.GetImplementation(&_SmartAccount.CallOpts)
}

// GetModulesPaginated is a free data retrieval call binding the contract method 0xcc2f8452.
//
// Solidity: function getModulesPaginated(address start, uint256 pageSize) view returns(address[] array, address next)
func (_SmartAccount *SmartAccountCaller) GetModulesPaginated(opts *bind.CallOpts, start common.Address, pageSize *big.Int) (struct {
	Array []common.Address
	Next  common.Address
}, error) {
	var out []interface{}
	err := _SmartAccount.contract.Call(opts, &out, "getModulesPaginated", start, pageSize)

	outstruct := new(struct {
		Array []common.Address
		Next  common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Array = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Next = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// GetModulesPaginated is a free data retrieval call binding the contract method 0xcc2f8452.
//
// Solidity: function getModulesPaginated(address start, uint256 pageSize) view returns(address[] array, address next)
func (_SmartAccount *SmartAccountSession) GetModulesPaginated(start common.Address, pageSize *big.Int) (struct {
	Array []common.Address
	Next  common.Address
}, error) {
	return _SmartAccount.Contract.GetModulesPaginated(&_SmartAccount.CallOpts, start, pageSize)
}

// GetModulesPaginated is a free data retrieval call binding the contract method 0xcc2f8452.
//
// Solidity: function getModulesPaginated(address start, uint256 pageSize) view returns(address[] array, address next)
func (_SmartAccount *SmartAccountCallerSession) GetModulesPaginated(start common.Address, pageSize *big.Int) (struct {
	Array []common.Address
	Next  common.Address
}, error) {
	return _SmartAccount.Contract.GetModulesPaginated(&_SmartAccount.CallOpts, start, pageSize)
}

// IsModuleEnabled is a free data retrieval call binding the contract method 0x2d9ad53d.
//
// Solidity: function isModuleEnabled(address module) view returns(bool)
func (_SmartAccount *SmartAccountCaller) IsModuleEnabled(opts *bind.CallOpts, module common.Address) (bool, error) {
	var out []interface{}
	err := _SmartAccount.contract.Call(opts, &out, "isModuleEnabled", module)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsModuleEnabled is a free data retrieval call binding the contract method 0x2d9ad53d.
//
// Solidity: function isModuleEnabled(address module) view returns(bool)
func (_SmartAccount *SmartAccountSession) IsModuleEnabled(module common.Address) (bool, error) {
	return _SmartAccount.Contract.IsModuleEnabled(&_SmartAccount.CallOpts, module)
}

// IsModuleEnabled is a free data retrieval call binding the contract method 0x2d9ad53d.
//
// Solidity: function isModuleEnabled(address module) view returns(bool)
func (_SmartAccount *SmartAccountCallerSession) IsModuleEnabled(module common.Address) (bool, error) {
	return _SmartAccount.Contract.IsModuleEnabled(&_SmartAccount.CallOpts, module)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 dataHash, bytes signature) view returns(bytes4)
func (_SmartAccount *SmartAccountCaller) IsValidSignature(opts *bind.CallOpts, dataHash [32]byte, signature []byte) ([4]byte, error) {
	var out []interface{}
	err := _SmartAccount.contract.Call(opts, &out, "isValidSignature", dataHash, signature)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 dataHash, bytes signature) view returns(bytes4)
func (_SmartAccount *SmartAccountSession) IsValidSignature(dataHash [32]byte, signature []byte) ([4]byte, error) {
	return _SmartAccount.Contract.IsValidSignature(&_SmartAccount.CallOpts, dataHash, signature)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 dataHash, bytes signature) view returns(bytes4)
func (_SmartAccount *SmartAccountCallerSession) IsValidSignature(dataHash [32]byte, signature []byte) ([4]byte, error) {
	return _SmartAccount.Contract.IsValidSignature(&_SmartAccount.CallOpts, dataHash, signature)
}

// Nonce is a free data retrieval call binding the contract method 0xd86f2b3c.
//
// Solidity: function nonce(uint192 _key) view returns(uint256)
func (_SmartAccount *SmartAccountCaller) Nonce(opts *bind.CallOpts, _key *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SmartAccount.contract.Call(opts, &out, "nonce", _key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0xd86f2b3c.
//
// Solidity: function nonce(uint192 _key) view returns(uint256)
func (_SmartAccount *SmartAccountSession) Nonce(_key *big.Int) (*big.Int, error) {
	return _SmartAccount.Contract.Nonce(&_SmartAccount.CallOpts, _key)
}

// Nonce is a free data retrieval call binding the contract method 0xd86f2b3c.
//
// Solidity: function nonce(uint192 _key) view returns(uint256)
func (_SmartAccount *SmartAccountCallerSession) Nonce(_key *big.Int) (*big.Int, error) {
	return _SmartAccount.Contract.Nonce(&_SmartAccount.CallOpts, _key)
}

// NoncesDeprecated is a free data retrieval call binding the contract method 0xf33623b1.
//
// Solidity: function noncesDeprecated(uint256 ) view returns(uint256)
func (_SmartAccount *SmartAccountCaller) NoncesDeprecated(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SmartAccount.contract.Call(opts, &out, "noncesDeprecated", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NoncesDeprecated is a free data retrieval call binding the contract method 0xf33623b1.
//
// Solidity: function noncesDeprecated(uint256 ) view returns(uint256)
func (_SmartAccount *SmartAccountSession) NoncesDeprecated(arg0 *big.Int) (*big.Int, error) {
	return _SmartAccount.Contract.NoncesDeprecated(&_SmartAccount.CallOpts, arg0)
}

// NoncesDeprecated is a free data retrieval call binding the contract method 0xf33623b1.
//
// Solidity: function noncesDeprecated(uint256 ) view returns(uint256)
func (_SmartAccount *SmartAccountCallerSession) NoncesDeprecated(arg0 *big.Int) (*big.Int, error) {
	return _SmartAccount.Contract.NoncesDeprecated(&_SmartAccount.CallOpts, arg0)
}

// OwnerDeprecated is a free data retrieval call binding the contract method 0x6424e9fe.
//
// Solidity: function ownerDeprecated() view returns(address)
func (_SmartAccount *SmartAccountCaller) OwnerDeprecated(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SmartAccount.contract.Call(opts, &out, "ownerDeprecated")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerDeprecated is a free data retrieval call binding the contract method 0x6424e9fe.
//
// Solidity: function ownerDeprecated() view returns(address)
func (_SmartAccount *SmartAccountSession) OwnerDeprecated() (common.Address, error) {
	return _SmartAccount.Contract.OwnerDeprecated(&_SmartAccount.CallOpts)
}

// OwnerDeprecated is a free data retrieval call binding the contract method 0x6424e9fe.
//
// Solidity: function ownerDeprecated() view returns(address)
func (_SmartAccount *SmartAccountCallerSession) OwnerDeprecated() (common.Address, error) {
	return _SmartAccount.Contract.OwnerDeprecated(&_SmartAccount.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_SmartAccount *SmartAccountCaller) SupportsInterface(opts *bind.CallOpts, _interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _SmartAccount.contract.Call(opts, &out, "supportsInterface", _interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_SmartAccount *SmartAccountSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _SmartAccount.Contract.SupportsInterface(&_SmartAccount.CallOpts, _interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_SmartAccount *SmartAccountCallerSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _SmartAccount.Contract.SupportsInterface(&_SmartAccount.CallOpts, _interfaceId)
}

// AddDeposit is a paid mutator transaction binding the contract method 0x4a58db19.
//
// Solidity: function addDeposit() payable returns()
func (_SmartAccount *SmartAccountTransactor) AddDeposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SmartAccount.contract.Transact(opts, "addDeposit")
}

// AddDeposit is a paid mutator transaction binding the contract method 0x4a58db19.
//
// Solidity: function addDeposit() payable returns()
func (_SmartAccount *SmartAccountSession) AddDeposit() (*types.Transaction, error) {
	return _SmartAccount.Contract.AddDeposit(&_SmartAccount.TransactOpts)
}

// AddDeposit is a paid mutator transaction binding the contract method 0x4a58db19.
//
// Solidity: function addDeposit() payable returns()
func (_SmartAccount *SmartAccountTransactorSession) AddDeposit() (*types.Transaction, error) {
	return _SmartAccount.Contract.AddDeposit(&_SmartAccount.TransactOpts)
}

// DisableModule is a paid mutator transaction binding the contract method 0xe009cfde.
//
// Solidity: function disableModule(address prevModule, address module) returns()
func (_SmartAccount *SmartAccountTransactor) DisableModule(opts *bind.TransactOpts, prevModule common.Address, module common.Address) (*types.Transaction, error) {
	return _SmartAccount.contract.Transact(opts, "disableModule", prevModule, module)
}

// DisableModule is a paid mutator transaction binding the contract method 0xe009cfde.
//
// Solidity: function disableModule(address prevModule, address module) returns()
func (_SmartAccount *SmartAccountSession) DisableModule(prevModule common.Address, module common.Address) (*types.Transaction, error) {
	return _SmartAccount.Contract.DisableModule(&_SmartAccount.TransactOpts, prevModule, module)
}

// DisableModule is a paid mutator transaction binding the contract method 0xe009cfde.
//
// Solidity: function disableModule(address prevModule, address module) returns()
func (_SmartAccount *SmartAccountTransactorSession) DisableModule(prevModule common.Address, module common.Address) (*types.Transaction, error) {
	return _SmartAccount.Contract.DisableModule(&_SmartAccount.TransactOpts, prevModule, module)
}

// EnableModule is a paid mutator transaction binding the contract method 0x610b5925.
//
// Solidity: function enableModule(address module) returns()
func (_SmartAccount *SmartAccountTransactor) EnableModule(opts *bind.TransactOpts, module common.Address) (*types.Transaction, error) {
	return _SmartAccount.contract.Transact(opts, "enableModule", module)
}

// EnableModule is a paid mutator transaction binding the contract method 0x610b5925.
//
// Solidity: function enableModule(address module) returns()
func (_SmartAccount *SmartAccountSession) EnableModule(module common.Address) (*types.Transaction, error) {
	return _SmartAccount.Contract.EnableModule(&_SmartAccount.TransactOpts, module)
}

// EnableModule is a paid mutator transaction binding the contract method 0x610b5925.
//
// Solidity: function enableModule(address module) returns()
func (_SmartAccount *SmartAccountTransactorSession) EnableModule(module common.Address) (*types.Transaction, error) {
	return _SmartAccount.Contract.EnableModule(&_SmartAccount.TransactOpts, module)
}

// ExecBatchTransactionFromModule is a paid mutator transaction binding the contract method 0xacfdf503.
//
// Solidity: function execBatchTransactionFromModule(address[] to, uint256[] value, bytes[] data, uint8[] operations) returns(bool success)
func (_SmartAccount *SmartAccountTransactor) ExecBatchTransactionFromModule(opts *bind.TransactOpts, to []common.Address, value []*big.Int, data [][]byte, operations []uint8) (*types.Transaction, error) {
	return _SmartAccount.contract.Transact(opts, "execBatchTransactionFromModule", to, value, data, operations)
}

// ExecBatchTransactionFromModule is a paid mutator transaction binding the contract method 0xacfdf503.
//
// Solidity: function execBatchTransactionFromModule(address[] to, uint256[] value, bytes[] data, uint8[] operations) returns(bool success)
func (_SmartAccount *SmartAccountSession) ExecBatchTransactionFromModule(to []common.Address, value []*big.Int, data [][]byte, operations []uint8) (*types.Transaction, error) {
	return _SmartAccount.Contract.ExecBatchTransactionFromModule(&_SmartAccount.TransactOpts, to, value, data, operations)
}

// ExecBatchTransactionFromModule is a paid mutator transaction binding the contract method 0xacfdf503.
//
// Solidity: function execBatchTransactionFromModule(address[] to, uint256[] value, bytes[] data, uint8[] operations) returns(bool success)
func (_SmartAccount *SmartAccountTransactorSession) ExecBatchTransactionFromModule(to []common.Address, value []*big.Int, data [][]byte, operations []uint8) (*types.Transaction, error) {
	return _SmartAccount.Contract.ExecBatchTransactionFromModule(&_SmartAccount.TransactOpts, to, value, data, operations)
}

// ExecTransactionFromModule is a paid mutator transaction binding the contract method 0x21632045.
//
// Solidity: function execTransactionFromModule(address to, uint256 value, bytes data, uint8 operation, uint256 txGas) returns(bool success)
func (_SmartAccount *SmartAccountTransactor) ExecTransactionFromModule(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte, operation uint8, txGas *big.Int) (*types.Transaction, error) {
	return _SmartAccount.contract.Transact(opts, "execTransactionFromModule", to, value, data, operation, txGas)
}

// ExecTransactionFromModule is a paid mutator transaction binding the contract method 0x21632045.
//
// Solidity: function execTransactionFromModule(address to, uint256 value, bytes data, uint8 operation, uint256 txGas) returns(bool success)
func (_SmartAccount *SmartAccountSession) ExecTransactionFromModule(to common.Address, value *big.Int, data []byte, operation uint8, txGas *big.Int) (*types.Transaction, error) {
	return _SmartAccount.Contract.ExecTransactionFromModule(&_SmartAccount.TransactOpts, to, value, data, operation, txGas)
}

// ExecTransactionFromModule is a paid mutator transaction binding the contract method 0x21632045.
//
// Solidity: function execTransactionFromModule(address to, uint256 value, bytes data, uint8 operation, uint256 txGas) returns(bool success)
func (_SmartAccount *SmartAccountTransactorSession) ExecTransactionFromModule(to common.Address, value *big.Int, data []byte, operation uint8, txGas *big.Int) (*types.Transaction, error) {
	return _SmartAccount.Contract.ExecTransactionFromModule(&_SmartAccount.TransactOpts, to, value, data, operation, txGas)
}

// ExecTransactionFromModule0 is a paid mutator transaction binding the contract method 0x468721a7.
//
// Solidity: function execTransactionFromModule(address to, uint256 value, bytes data, uint8 operation) returns(bool)
func (_SmartAccount *SmartAccountTransactor) ExecTransactionFromModule0(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _SmartAccount.contract.Transact(opts, "execTransactionFromModule0", to, value, data, operation)
}

// ExecTransactionFromModule0 is a paid mutator transaction binding the contract method 0x468721a7.
//
// Solidity: function execTransactionFromModule(address to, uint256 value, bytes data, uint8 operation) returns(bool)
func (_SmartAccount *SmartAccountSession) ExecTransactionFromModule0(to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _SmartAccount.Contract.ExecTransactionFromModule0(&_SmartAccount.TransactOpts, to, value, data, operation)
}

// ExecTransactionFromModule0 is a paid mutator transaction binding the contract method 0x468721a7.
//
// Solidity: function execTransactionFromModule(address to, uint256 value, bytes data, uint8 operation) returns(bool)
func (_SmartAccount *SmartAccountTransactorSession) ExecTransactionFromModule0(to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _SmartAccount.Contract.ExecTransactionFromModule0(&_SmartAccount.TransactOpts, to, value, data, operation)
}

// ExecTransactionFromModuleReturnData is a paid mutator transaction binding the contract method 0x5229073f.
//
// Solidity: function execTransactionFromModuleReturnData(address to, uint256 value, bytes data, uint8 operation) returns(bool success, bytes returnData)
func (_SmartAccount *SmartAccountTransactor) ExecTransactionFromModuleReturnData(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _SmartAccount.contract.Transact(opts, "execTransactionFromModuleReturnData", to, value, data, operation)
}

// ExecTransactionFromModuleReturnData is a paid mutator transaction binding the contract method 0x5229073f.
//
// Solidity: function execTransactionFromModuleReturnData(address to, uint256 value, bytes data, uint8 operation) returns(bool success, bytes returnData)
func (_SmartAccount *SmartAccountSession) ExecTransactionFromModuleReturnData(to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _SmartAccount.Contract.ExecTransactionFromModuleReturnData(&_SmartAccount.TransactOpts, to, value, data, operation)
}

// ExecTransactionFromModuleReturnData is a paid mutator transaction binding the contract method 0x5229073f.
//
// Solidity: function execTransactionFromModuleReturnData(address to, uint256 value, bytes data, uint8 operation) returns(bool success, bytes returnData)
func (_SmartAccount *SmartAccountTransactorSession) ExecTransactionFromModuleReturnData(to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _SmartAccount.Contract.ExecTransactionFromModuleReturnData(&_SmartAccount.TransactOpts, to, value, data, operation)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address dest, uint256 value, bytes func) returns()
func (_SmartAccount *SmartAccountTransactor) Execute(opts *bind.TransactOpts, dest common.Address, value *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _SmartAccount.contract.Transact(opts, "execute", dest, value, arg2)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address dest, uint256 value, bytes func) returns()
func (_SmartAccount *SmartAccountSession) Execute(dest common.Address, value *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _SmartAccount.Contract.Execute(&_SmartAccount.TransactOpts, dest, value, arg2)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address dest, uint256 value, bytes func) returns()
func (_SmartAccount *SmartAccountTransactorSession) Execute(dest common.Address, value *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _SmartAccount.Contract.Execute(&_SmartAccount.TransactOpts, dest, value, arg2)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x47e1da2a.
//
// Solidity: function executeBatch(address[] dest, uint256[] value, bytes[] func) returns()
func (_SmartAccount *SmartAccountTransactor) ExecuteBatch(opts *bind.TransactOpts, dest []common.Address, value []*big.Int, arg2 [][]byte) (*types.Transaction, error) {
	return _SmartAccount.contract.Transact(opts, "executeBatch", dest, value, arg2)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x47e1da2a.
//
// Solidity: function executeBatch(address[] dest, uint256[] value, bytes[] func) returns()
func (_SmartAccount *SmartAccountSession) ExecuteBatch(dest []common.Address, value []*big.Int, arg2 [][]byte) (*types.Transaction, error) {
	return _SmartAccount.Contract.ExecuteBatch(&_SmartAccount.TransactOpts, dest, value, arg2)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x47e1da2a.
//
// Solidity: function executeBatch(address[] dest, uint256[] value, bytes[] func) returns()
func (_SmartAccount *SmartAccountTransactorSession) ExecuteBatch(dest []common.Address, value []*big.Int, arg2 [][]byte) (*types.Transaction, error) {
	return _SmartAccount.Contract.ExecuteBatch(&_SmartAccount.TransactOpts, dest, value, arg2)
}

// ExecuteBatchY6U is a paid mutator transaction binding the contract method 0x00004680.
//
// Solidity: function executeBatch_y6U(address[] dest, uint256[] value, bytes[] func) returns()
func (_SmartAccount *SmartAccountTransactor) ExecuteBatchY6U(opts *bind.TransactOpts, dest []common.Address, value []*big.Int, arg2 [][]byte) (*types.Transaction, error) {
	return _SmartAccount.contract.Transact(opts, "executeBatch_y6U", dest, value, arg2)
}

// ExecuteBatchY6U is a paid mutator transaction binding the contract method 0x00004680.
//
// Solidity: function executeBatch_y6U(address[] dest, uint256[] value, bytes[] func) returns()
func (_SmartAccount *SmartAccountSession) ExecuteBatchY6U(dest []common.Address, value []*big.Int, arg2 [][]byte) (*types.Transaction, error) {
	return _SmartAccount.Contract.ExecuteBatchY6U(&_SmartAccount.TransactOpts, dest, value, arg2)
}

// ExecuteBatchY6U is a paid mutator transaction binding the contract method 0x00004680.
//
// Solidity: function executeBatch_y6U(address[] dest, uint256[] value, bytes[] func) returns()
func (_SmartAccount *SmartAccountTransactorSession) ExecuteBatchY6U(dest []common.Address, value []*big.Int, arg2 [][]byte) (*types.Transaction, error) {
	return _SmartAccount.Contract.ExecuteBatchY6U(&_SmartAccount.TransactOpts, dest, value, arg2)
}

// ExecuteNcC is a paid mutator transaction binding the contract method 0x0000189a.
//
// Solidity: function execute_ncC(address dest, uint256 value, bytes func) returns()
func (_SmartAccount *SmartAccountTransactor) ExecuteNcC(opts *bind.TransactOpts, dest common.Address, value *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _SmartAccount.contract.Transact(opts, "execute_ncC", dest, value, arg2)
}

// ExecuteNcC is a paid mutator transaction binding the contract method 0x0000189a.
//
// Solidity: function execute_ncC(address dest, uint256 value, bytes func) returns()
func (_SmartAccount *SmartAccountSession) ExecuteNcC(dest common.Address, value *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _SmartAccount.Contract.ExecuteNcC(&_SmartAccount.TransactOpts, dest, value, arg2)
}

// ExecuteNcC is a paid mutator transaction binding the contract method 0x0000189a.
//
// Solidity: function execute_ncC(address dest, uint256 value, bytes func) returns()
func (_SmartAccount *SmartAccountTransactorSession) ExecuteNcC(dest common.Address, value *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _SmartAccount.Contract.ExecuteNcC(&_SmartAccount.TransactOpts, dest, value, arg2)
}

// Init is a paid mutator transaction binding the contract method 0x378dfd8e.
//
// Solidity: function init(address handler, address moduleSetupContract, bytes moduleSetupData) returns(address)
func (_SmartAccount *SmartAccountTransactor) Init(opts *bind.TransactOpts, handler common.Address, moduleSetupContract common.Address, moduleSetupData []byte) (*types.Transaction, error) {
	return _SmartAccount.contract.Transact(opts, "init", handler, moduleSetupContract, moduleSetupData)
}

// Init is a paid mutator transaction binding the contract method 0x378dfd8e.
//
// Solidity: function init(address handler, address moduleSetupContract, bytes moduleSetupData) returns(address)
func (_SmartAccount *SmartAccountSession) Init(handler common.Address, moduleSetupContract common.Address, moduleSetupData []byte) (*types.Transaction, error) {
	return _SmartAccount.Contract.Init(&_SmartAccount.TransactOpts, handler, moduleSetupContract, moduleSetupData)
}

// Init is a paid mutator transaction binding the contract method 0x378dfd8e.
//
// Solidity: function init(address handler, address moduleSetupContract, bytes moduleSetupData) returns(address)
func (_SmartAccount *SmartAccountTransactorSession) Init(handler common.Address, moduleSetupContract common.Address, moduleSetupData []byte) (*types.Transaction, error) {
	return _SmartAccount.Contract.Init(&_SmartAccount.TransactOpts, handler, moduleSetupContract, moduleSetupData)
}

// SetFallbackHandler is a paid mutator transaction binding the contract method 0xf08a0323.
//
// Solidity: function setFallbackHandler(address handler) returns()
func (_SmartAccount *SmartAccountTransactor) SetFallbackHandler(opts *bind.TransactOpts, handler common.Address) (*types.Transaction, error) {
	return _SmartAccount.contract.Transact(opts, "setFallbackHandler", handler)
}

// SetFallbackHandler is a paid mutator transaction binding the contract method 0xf08a0323.
//
// Solidity: function setFallbackHandler(address handler) returns()
func (_SmartAccount *SmartAccountSession) SetFallbackHandler(handler common.Address) (*types.Transaction, error) {
	return _SmartAccount.Contract.SetFallbackHandler(&_SmartAccount.TransactOpts, handler)
}

// SetFallbackHandler is a paid mutator transaction binding the contract method 0xf08a0323.
//
// Solidity: function setFallbackHandler(address handler) returns()
func (_SmartAccount *SmartAccountTransactorSession) SetFallbackHandler(handler common.Address) (*types.Transaction, error) {
	return _SmartAccount.Contract.SetFallbackHandler(&_SmartAccount.TransactOpts, handler)
}

// SetupAndEnableModule is a paid mutator transaction binding the contract method 0x5305dd27.
//
// Solidity: function setupAndEnableModule(address setupContract, bytes setupData) returns(address)
func (_SmartAccount *SmartAccountTransactor) SetupAndEnableModule(opts *bind.TransactOpts, setupContract common.Address, setupData []byte) (*types.Transaction, error) {
	return _SmartAccount.contract.Transact(opts, "setupAndEnableModule", setupContract, setupData)
}

// SetupAndEnableModule is a paid mutator transaction binding the contract method 0x5305dd27.
//
// Solidity: function setupAndEnableModule(address setupContract, bytes setupData) returns(address)
func (_SmartAccount *SmartAccountSession) SetupAndEnableModule(setupContract common.Address, setupData []byte) (*types.Transaction, error) {
	return _SmartAccount.Contract.SetupAndEnableModule(&_SmartAccount.TransactOpts, setupContract, setupData)
}

// SetupAndEnableModule is a paid mutator transaction binding the contract method 0x5305dd27.
//
// Solidity: function setupAndEnableModule(address setupContract, bytes setupData) returns(address)
func (_SmartAccount *SmartAccountTransactorSession) SetupAndEnableModule(setupContract common.Address, setupData []byte) (*types.Transaction, error) {
	return _SmartAccount.Contract.SetupAndEnableModule(&_SmartAccount.TransactOpts, setupContract, setupData)
}

// UpdateImplementation is a paid mutator transaction binding the contract method 0x025b22bc.
//
// Solidity: function updateImplementation(address _implementation) returns()
func (_SmartAccount *SmartAccountTransactor) UpdateImplementation(opts *bind.TransactOpts, _implementation common.Address) (*types.Transaction, error) {
	return _SmartAccount.contract.Transact(opts, "updateImplementation", _implementation)
}

// UpdateImplementation is a paid mutator transaction binding the contract method 0x025b22bc.
//
// Solidity: function updateImplementation(address _implementation) returns()
func (_SmartAccount *SmartAccountSession) UpdateImplementation(_implementation common.Address) (*types.Transaction, error) {
	return _SmartAccount.Contract.UpdateImplementation(&_SmartAccount.TransactOpts, _implementation)
}

// UpdateImplementation is a paid mutator transaction binding the contract method 0x025b22bc.
//
// Solidity: function updateImplementation(address _implementation) returns()
func (_SmartAccount *SmartAccountTransactorSession) UpdateImplementation(_implementation common.Address) (*types.Transaction, error) {
	return _SmartAccount.Contract.UpdateImplementation(&_SmartAccount.TransactOpts, _implementation)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x3a871cdd.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_SmartAccount *SmartAccountTransactor) ValidateUserOp(opts *bind.TransactOpts, userOp UserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _SmartAccount.contract.Transact(opts, "validateUserOp", userOp, userOpHash, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x3a871cdd.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_SmartAccount *SmartAccountSession) ValidateUserOp(userOp UserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _SmartAccount.Contract.ValidateUserOp(&_SmartAccount.TransactOpts, userOp, userOpHash, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x3a871cdd.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_SmartAccount *SmartAccountTransactorSession) ValidateUserOp(userOp UserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _SmartAccount.Contract.ValidateUserOp(&_SmartAccount.TransactOpts, userOp, userOpHash, missingAccountFunds)
}

// WithdrawDepositTo is a paid mutator transaction binding the contract method 0x4d44560d.
//
// Solidity: function withdrawDepositTo(address withdrawAddress, uint256 amount) payable returns()
func (_SmartAccount *SmartAccountTransactor) WithdrawDepositTo(opts *bind.TransactOpts, withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SmartAccount.contract.Transact(opts, "withdrawDepositTo", withdrawAddress, amount)
}

// WithdrawDepositTo is a paid mutator transaction binding the contract method 0x4d44560d.
//
// Solidity: function withdrawDepositTo(address withdrawAddress, uint256 amount) payable returns()
func (_SmartAccount *SmartAccountSession) WithdrawDepositTo(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SmartAccount.Contract.WithdrawDepositTo(&_SmartAccount.TransactOpts, withdrawAddress, amount)
}

// WithdrawDepositTo is a paid mutator transaction binding the contract method 0x4d44560d.
//
// Solidity: function withdrawDepositTo(address withdrawAddress, uint256 amount) payable returns()
func (_SmartAccount *SmartAccountTransactorSession) WithdrawDepositTo(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SmartAccount.Contract.WithdrawDepositTo(&_SmartAccount.TransactOpts, withdrawAddress, amount)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_SmartAccount *SmartAccountTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _SmartAccount.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_SmartAccount *SmartAccountSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _SmartAccount.Contract.Fallback(&_SmartAccount.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_SmartAccount *SmartAccountTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _SmartAccount.Contract.Fallback(&_SmartAccount.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SmartAccount *SmartAccountTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SmartAccount.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SmartAccount *SmartAccountSession) Receive() (*types.Transaction, error) {
	return _SmartAccount.Contract.Receive(&_SmartAccount.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SmartAccount *SmartAccountTransactorSession) Receive() (*types.Transaction, error) {
	return _SmartAccount.Contract.Receive(&_SmartAccount.TransactOpts)
}

// SmartAccountChangedFallbackHandlerIterator is returned from FilterChangedFallbackHandler and is used to iterate over the raw logs and unpacked data for ChangedFallbackHandler events raised by the SmartAccount contract.
type SmartAccountChangedFallbackHandlerIterator struct {
	Event *SmartAccountChangedFallbackHandler // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SmartAccountChangedFallbackHandlerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartAccountChangedFallbackHandler)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SmartAccountChangedFallbackHandler)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SmartAccountChangedFallbackHandlerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartAccountChangedFallbackHandlerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartAccountChangedFallbackHandler represents a ChangedFallbackHandler event raised by the SmartAccount contract.
type SmartAccountChangedFallbackHandler struct {
	PreviousHandler common.Address
	Handler         common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterChangedFallbackHandler is a free log retrieval operation binding the contract event 0x06be9a1bea257286cf2afa8205ed494ca9d6a4b41aa58d04238deebada20fb0c.
//
// Solidity: event ChangedFallbackHandler(address indexed previousHandler, address indexed handler)
func (_SmartAccount *SmartAccountFilterer) FilterChangedFallbackHandler(opts *bind.FilterOpts, previousHandler []common.Address, handler []common.Address) (*SmartAccountChangedFallbackHandlerIterator, error) {

	var previousHandlerRule []interface{}
	for _, previousHandlerItem := range previousHandler {
		previousHandlerRule = append(previousHandlerRule, previousHandlerItem)
	}
	var handlerRule []interface{}
	for _, handlerItem := range handler {
		handlerRule = append(handlerRule, handlerItem)
	}

	logs, sub, err := _SmartAccount.contract.FilterLogs(opts, "ChangedFallbackHandler", previousHandlerRule, handlerRule)
	if err != nil {
		return nil, err
	}
	return &SmartAccountChangedFallbackHandlerIterator{contract: _SmartAccount.contract, event: "ChangedFallbackHandler", logs: logs, sub: sub}, nil
}

// WatchChangedFallbackHandler is a free log subscription operation binding the contract event 0x06be9a1bea257286cf2afa8205ed494ca9d6a4b41aa58d04238deebada20fb0c.
//
// Solidity: event ChangedFallbackHandler(address indexed previousHandler, address indexed handler)
func (_SmartAccount *SmartAccountFilterer) WatchChangedFallbackHandler(opts *bind.WatchOpts, sink chan<- *SmartAccountChangedFallbackHandler, previousHandler []common.Address, handler []common.Address) (event.Subscription, error) {

	var previousHandlerRule []interface{}
	for _, previousHandlerItem := range previousHandler {
		previousHandlerRule = append(previousHandlerRule, previousHandlerItem)
	}
	var handlerRule []interface{}
	for _, handlerItem := range handler {
		handlerRule = append(handlerRule, handlerItem)
	}

	logs, sub, err := _SmartAccount.contract.WatchLogs(opts, "ChangedFallbackHandler", previousHandlerRule, handlerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartAccountChangedFallbackHandler)
				if err := _SmartAccount.contract.UnpackLog(event, "ChangedFallbackHandler", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseChangedFallbackHandler is a log parse operation binding the contract event 0x06be9a1bea257286cf2afa8205ed494ca9d6a4b41aa58d04238deebada20fb0c.
//
// Solidity: event ChangedFallbackHandler(address indexed previousHandler, address indexed handler)
func (_SmartAccount *SmartAccountFilterer) ParseChangedFallbackHandler(log types.Log) (*SmartAccountChangedFallbackHandler, error) {
	event := new(SmartAccountChangedFallbackHandler)
	if err := _SmartAccount.contract.UnpackLog(event, "ChangedFallbackHandler", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmartAccountDisabledModuleIterator is returned from FilterDisabledModule and is used to iterate over the raw logs and unpacked data for DisabledModule events raised by the SmartAccount contract.
type SmartAccountDisabledModuleIterator struct {
	Event *SmartAccountDisabledModule // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SmartAccountDisabledModuleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartAccountDisabledModule)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SmartAccountDisabledModule)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SmartAccountDisabledModuleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartAccountDisabledModuleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartAccountDisabledModule represents a DisabledModule event raised by the SmartAccount contract.
type SmartAccountDisabledModule struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDisabledModule is a free log retrieval operation binding the contract event 0xaab4fa2b463f581b2b32cb3b7e3b704b9ce37cc209b5fb4d77e593ace4054276.
//
// Solidity: event DisabledModule(address module)
func (_SmartAccount *SmartAccountFilterer) FilterDisabledModule(opts *bind.FilterOpts) (*SmartAccountDisabledModuleIterator, error) {

	logs, sub, err := _SmartAccount.contract.FilterLogs(opts, "DisabledModule")
	if err != nil {
		return nil, err
	}
	return &SmartAccountDisabledModuleIterator{contract: _SmartAccount.contract, event: "DisabledModule", logs: logs, sub: sub}, nil
}

// WatchDisabledModule is a free log subscription operation binding the contract event 0xaab4fa2b463f581b2b32cb3b7e3b704b9ce37cc209b5fb4d77e593ace4054276.
//
// Solidity: event DisabledModule(address module)
func (_SmartAccount *SmartAccountFilterer) WatchDisabledModule(opts *bind.WatchOpts, sink chan<- *SmartAccountDisabledModule) (event.Subscription, error) {

	logs, sub, err := _SmartAccount.contract.WatchLogs(opts, "DisabledModule")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartAccountDisabledModule)
				if err := _SmartAccount.contract.UnpackLog(event, "DisabledModule", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDisabledModule is a log parse operation binding the contract event 0xaab4fa2b463f581b2b32cb3b7e3b704b9ce37cc209b5fb4d77e593ace4054276.
//
// Solidity: event DisabledModule(address module)
func (_SmartAccount *SmartAccountFilterer) ParseDisabledModule(log types.Log) (*SmartAccountDisabledModule, error) {
	event := new(SmartAccountDisabledModule)
	if err := _SmartAccount.contract.UnpackLog(event, "DisabledModule", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmartAccountEnabledModuleIterator is returned from FilterEnabledModule and is used to iterate over the raw logs and unpacked data for EnabledModule events raised by the SmartAccount contract.
type SmartAccountEnabledModuleIterator struct {
	Event *SmartAccountEnabledModule // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SmartAccountEnabledModuleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartAccountEnabledModule)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SmartAccountEnabledModule)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SmartAccountEnabledModuleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartAccountEnabledModuleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartAccountEnabledModule represents a EnabledModule event raised by the SmartAccount contract.
type SmartAccountEnabledModule struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEnabledModule is a free log retrieval operation binding the contract event 0xecdf3a3effea5783a3c4c2140e677577666428d44ed9d474a0b3a4c9943f8440.
//
// Solidity: event EnabledModule(address module)
func (_SmartAccount *SmartAccountFilterer) FilterEnabledModule(opts *bind.FilterOpts) (*SmartAccountEnabledModuleIterator, error) {

	logs, sub, err := _SmartAccount.contract.FilterLogs(opts, "EnabledModule")
	if err != nil {
		return nil, err
	}
	return &SmartAccountEnabledModuleIterator{contract: _SmartAccount.contract, event: "EnabledModule", logs: logs, sub: sub}, nil
}

// WatchEnabledModule is a free log subscription operation binding the contract event 0xecdf3a3effea5783a3c4c2140e677577666428d44ed9d474a0b3a4c9943f8440.
//
// Solidity: event EnabledModule(address module)
func (_SmartAccount *SmartAccountFilterer) WatchEnabledModule(opts *bind.WatchOpts, sink chan<- *SmartAccountEnabledModule) (event.Subscription, error) {

	logs, sub, err := _SmartAccount.contract.WatchLogs(opts, "EnabledModule")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartAccountEnabledModule)
				if err := _SmartAccount.contract.UnpackLog(event, "EnabledModule", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEnabledModule is a log parse operation binding the contract event 0xecdf3a3effea5783a3c4c2140e677577666428d44ed9d474a0b3a4c9943f8440.
//
// Solidity: event EnabledModule(address module)
func (_SmartAccount *SmartAccountFilterer) ParseEnabledModule(log types.Log) (*SmartAccountEnabledModule, error) {
	event := new(SmartAccountEnabledModule)
	if err := _SmartAccount.contract.UnpackLog(event, "EnabledModule", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmartAccountExecutionFailureIterator is returned from FilterExecutionFailure and is used to iterate over the raw logs and unpacked data for ExecutionFailure events raised by the SmartAccount contract.
type SmartAccountExecutionFailureIterator struct {
	Event *SmartAccountExecutionFailure // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SmartAccountExecutionFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartAccountExecutionFailure)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SmartAccountExecutionFailure)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SmartAccountExecutionFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartAccountExecutionFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartAccountExecutionFailure represents a ExecutionFailure event raised by the SmartAccount contract.
type SmartAccountExecutionFailure struct {
	To        common.Address
	Value     *big.Int
	Data      common.Hash
	Operation uint8
	TxGas     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterExecutionFailure is a free log retrieval operation binding the contract event 0x3ddd038f78c876172d5dbfd730b14c9f8692dfa197ef104eaac6df3f85a0874a.
//
// Solidity: event ExecutionFailure(address indexed to, uint256 indexed value, bytes indexed data, uint8 operation, uint256 txGas)
func (_SmartAccount *SmartAccountFilterer) FilterExecutionFailure(opts *bind.FilterOpts, to []common.Address, value []*big.Int, data [][]byte) (*SmartAccountExecutionFailureIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}
	var dataRule []interface{}
	for _, dataItem := range data {
		dataRule = append(dataRule, dataItem)
	}

	logs, sub, err := _SmartAccount.contract.FilterLogs(opts, "ExecutionFailure", toRule, valueRule, dataRule)
	if err != nil {
		return nil, err
	}
	return &SmartAccountExecutionFailureIterator{contract: _SmartAccount.contract, event: "ExecutionFailure", logs: logs, sub: sub}, nil
}

// WatchExecutionFailure is a free log subscription operation binding the contract event 0x3ddd038f78c876172d5dbfd730b14c9f8692dfa197ef104eaac6df3f85a0874a.
//
// Solidity: event ExecutionFailure(address indexed to, uint256 indexed value, bytes indexed data, uint8 operation, uint256 txGas)
func (_SmartAccount *SmartAccountFilterer) WatchExecutionFailure(opts *bind.WatchOpts, sink chan<- *SmartAccountExecutionFailure, to []common.Address, value []*big.Int, data [][]byte) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}
	var dataRule []interface{}
	for _, dataItem := range data {
		dataRule = append(dataRule, dataItem)
	}

	logs, sub, err := _SmartAccount.contract.WatchLogs(opts, "ExecutionFailure", toRule, valueRule, dataRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartAccountExecutionFailure)
				if err := _SmartAccount.contract.UnpackLog(event, "ExecutionFailure", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExecutionFailure is a log parse operation binding the contract event 0x3ddd038f78c876172d5dbfd730b14c9f8692dfa197ef104eaac6df3f85a0874a.
//
// Solidity: event ExecutionFailure(address indexed to, uint256 indexed value, bytes indexed data, uint8 operation, uint256 txGas)
func (_SmartAccount *SmartAccountFilterer) ParseExecutionFailure(log types.Log) (*SmartAccountExecutionFailure, error) {
	event := new(SmartAccountExecutionFailure)
	if err := _SmartAccount.contract.UnpackLog(event, "ExecutionFailure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmartAccountExecutionFromModuleFailureIterator is returned from FilterExecutionFromModuleFailure and is used to iterate over the raw logs and unpacked data for ExecutionFromModuleFailure events raised by the SmartAccount contract.
type SmartAccountExecutionFromModuleFailureIterator struct {
	Event *SmartAccountExecutionFromModuleFailure // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SmartAccountExecutionFromModuleFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartAccountExecutionFromModuleFailure)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SmartAccountExecutionFromModuleFailure)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SmartAccountExecutionFromModuleFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartAccountExecutionFromModuleFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartAccountExecutionFromModuleFailure represents a ExecutionFromModuleFailure event raised by the SmartAccount contract.
type SmartAccountExecutionFromModuleFailure struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExecutionFromModuleFailure is a free log retrieval operation binding the contract event 0xacd2c8702804128fdb0db2bb49f6d127dd0181c13fd45dbfe16de0930e2bd375.
//
// Solidity: event ExecutionFromModuleFailure(address indexed module)
func (_SmartAccount *SmartAccountFilterer) FilterExecutionFromModuleFailure(opts *bind.FilterOpts, module []common.Address) (*SmartAccountExecutionFromModuleFailureIterator, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _SmartAccount.contract.FilterLogs(opts, "ExecutionFromModuleFailure", moduleRule)
	if err != nil {
		return nil, err
	}
	return &SmartAccountExecutionFromModuleFailureIterator{contract: _SmartAccount.contract, event: "ExecutionFromModuleFailure", logs: logs, sub: sub}, nil
}

// WatchExecutionFromModuleFailure is a free log subscription operation binding the contract event 0xacd2c8702804128fdb0db2bb49f6d127dd0181c13fd45dbfe16de0930e2bd375.
//
// Solidity: event ExecutionFromModuleFailure(address indexed module)
func (_SmartAccount *SmartAccountFilterer) WatchExecutionFromModuleFailure(opts *bind.WatchOpts, sink chan<- *SmartAccountExecutionFromModuleFailure, module []common.Address) (event.Subscription, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _SmartAccount.contract.WatchLogs(opts, "ExecutionFromModuleFailure", moduleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartAccountExecutionFromModuleFailure)
				if err := _SmartAccount.contract.UnpackLog(event, "ExecutionFromModuleFailure", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExecutionFromModuleFailure is a log parse operation binding the contract event 0xacd2c8702804128fdb0db2bb49f6d127dd0181c13fd45dbfe16de0930e2bd375.
//
// Solidity: event ExecutionFromModuleFailure(address indexed module)
func (_SmartAccount *SmartAccountFilterer) ParseExecutionFromModuleFailure(log types.Log) (*SmartAccountExecutionFromModuleFailure, error) {
	event := new(SmartAccountExecutionFromModuleFailure)
	if err := _SmartAccount.contract.UnpackLog(event, "ExecutionFromModuleFailure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmartAccountExecutionFromModuleSuccessIterator is returned from FilterExecutionFromModuleSuccess and is used to iterate over the raw logs and unpacked data for ExecutionFromModuleSuccess events raised by the SmartAccount contract.
type SmartAccountExecutionFromModuleSuccessIterator struct {
	Event *SmartAccountExecutionFromModuleSuccess // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SmartAccountExecutionFromModuleSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartAccountExecutionFromModuleSuccess)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SmartAccountExecutionFromModuleSuccess)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SmartAccountExecutionFromModuleSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartAccountExecutionFromModuleSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartAccountExecutionFromModuleSuccess represents a ExecutionFromModuleSuccess event raised by the SmartAccount contract.
type SmartAccountExecutionFromModuleSuccess struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExecutionFromModuleSuccess is a free log retrieval operation binding the contract event 0x6895c13664aa4f67288b25d7a21d7aaa34916e355fb9b6fae0a139a9085becb8.
//
// Solidity: event ExecutionFromModuleSuccess(address indexed module)
func (_SmartAccount *SmartAccountFilterer) FilterExecutionFromModuleSuccess(opts *bind.FilterOpts, module []common.Address) (*SmartAccountExecutionFromModuleSuccessIterator, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _SmartAccount.contract.FilterLogs(opts, "ExecutionFromModuleSuccess", moduleRule)
	if err != nil {
		return nil, err
	}
	return &SmartAccountExecutionFromModuleSuccessIterator{contract: _SmartAccount.contract, event: "ExecutionFromModuleSuccess", logs: logs, sub: sub}, nil
}

// WatchExecutionFromModuleSuccess is a free log subscription operation binding the contract event 0x6895c13664aa4f67288b25d7a21d7aaa34916e355fb9b6fae0a139a9085becb8.
//
// Solidity: event ExecutionFromModuleSuccess(address indexed module)
func (_SmartAccount *SmartAccountFilterer) WatchExecutionFromModuleSuccess(opts *bind.WatchOpts, sink chan<- *SmartAccountExecutionFromModuleSuccess, module []common.Address) (event.Subscription, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _SmartAccount.contract.WatchLogs(opts, "ExecutionFromModuleSuccess", moduleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartAccountExecutionFromModuleSuccess)
				if err := _SmartAccount.contract.UnpackLog(event, "ExecutionFromModuleSuccess", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExecutionFromModuleSuccess is a log parse operation binding the contract event 0x6895c13664aa4f67288b25d7a21d7aaa34916e355fb9b6fae0a139a9085becb8.
//
// Solidity: event ExecutionFromModuleSuccess(address indexed module)
func (_SmartAccount *SmartAccountFilterer) ParseExecutionFromModuleSuccess(log types.Log) (*SmartAccountExecutionFromModuleSuccess, error) {
	event := new(SmartAccountExecutionFromModuleSuccess)
	if err := _SmartAccount.contract.UnpackLog(event, "ExecutionFromModuleSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmartAccountExecutionSuccessIterator is returned from FilterExecutionSuccess and is used to iterate over the raw logs and unpacked data for ExecutionSuccess events raised by the SmartAccount contract.
type SmartAccountExecutionSuccessIterator struct {
	Event *SmartAccountExecutionSuccess // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SmartAccountExecutionSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartAccountExecutionSuccess)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SmartAccountExecutionSuccess)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SmartAccountExecutionSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartAccountExecutionSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartAccountExecutionSuccess represents a ExecutionSuccess event raised by the SmartAccount contract.
type SmartAccountExecutionSuccess struct {
	To        common.Address
	Value     *big.Int
	Data      common.Hash
	Operation uint8
	TxGas     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterExecutionSuccess is a free log retrieval operation binding the contract event 0x81d12fffced46c214dfae8ab8fa0b9f7b69f70c9d500e33f612f2105deb261ee.
//
// Solidity: event ExecutionSuccess(address indexed to, uint256 indexed value, bytes indexed data, uint8 operation, uint256 txGas)
func (_SmartAccount *SmartAccountFilterer) FilterExecutionSuccess(opts *bind.FilterOpts, to []common.Address, value []*big.Int, data [][]byte) (*SmartAccountExecutionSuccessIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}
	var dataRule []interface{}
	for _, dataItem := range data {
		dataRule = append(dataRule, dataItem)
	}

	logs, sub, err := _SmartAccount.contract.FilterLogs(opts, "ExecutionSuccess", toRule, valueRule, dataRule)
	if err != nil {
		return nil, err
	}
	return &SmartAccountExecutionSuccessIterator{contract: _SmartAccount.contract, event: "ExecutionSuccess", logs: logs, sub: sub}, nil
}

// WatchExecutionSuccess is a free log subscription operation binding the contract event 0x81d12fffced46c214dfae8ab8fa0b9f7b69f70c9d500e33f612f2105deb261ee.
//
// Solidity: event ExecutionSuccess(address indexed to, uint256 indexed value, bytes indexed data, uint8 operation, uint256 txGas)
func (_SmartAccount *SmartAccountFilterer) WatchExecutionSuccess(opts *bind.WatchOpts, sink chan<- *SmartAccountExecutionSuccess, to []common.Address, value []*big.Int, data [][]byte) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}
	var dataRule []interface{}
	for _, dataItem := range data {
		dataRule = append(dataRule, dataItem)
	}

	logs, sub, err := _SmartAccount.contract.WatchLogs(opts, "ExecutionSuccess", toRule, valueRule, dataRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartAccountExecutionSuccess)
				if err := _SmartAccount.contract.UnpackLog(event, "ExecutionSuccess", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExecutionSuccess is a log parse operation binding the contract event 0x81d12fffced46c214dfae8ab8fa0b9f7b69f70c9d500e33f612f2105deb261ee.
//
// Solidity: event ExecutionSuccess(address indexed to, uint256 indexed value, bytes indexed data, uint8 operation, uint256 txGas)
func (_SmartAccount *SmartAccountFilterer) ParseExecutionSuccess(log types.Log) (*SmartAccountExecutionSuccess, error) {
	event := new(SmartAccountExecutionSuccess)
	if err := _SmartAccount.contract.UnpackLog(event, "ExecutionSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmartAccountImplementationUpdatedIterator is returned from FilterImplementationUpdated and is used to iterate over the raw logs and unpacked data for ImplementationUpdated events raised by the SmartAccount contract.
type SmartAccountImplementationUpdatedIterator struct {
	Event *SmartAccountImplementationUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SmartAccountImplementationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartAccountImplementationUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SmartAccountImplementationUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SmartAccountImplementationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartAccountImplementationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartAccountImplementationUpdated represents a ImplementationUpdated event raised by the SmartAccount contract.
type SmartAccountImplementationUpdated struct {
	OldImplementation common.Address
	NewImplementation common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterImplementationUpdated is a free log retrieval operation binding the contract event 0xaa3f731066a578e5f39b4215468d826cdd15373cbc0dfc9cb9bdc649718ef7da.
//
// Solidity: event ImplementationUpdated(address indexed oldImplementation, address indexed newImplementation)
func (_SmartAccount *SmartAccountFilterer) FilterImplementationUpdated(opts *bind.FilterOpts, oldImplementation []common.Address, newImplementation []common.Address) (*SmartAccountImplementationUpdatedIterator, error) {

	var oldImplementationRule []interface{}
	for _, oldImplementationItem := range oldImplementation {
		oldImplementationRule = append(oldImplementationRule, oldImplementationItem)
	}
	var newImplementationRule []interface{}
	for _, newImplementationItem := range newImplementation {
		newImplementationRule = append(newImplementationRule, newImplementationItem)
	}

	logs, sub, err := _SmartAccount.contract.FilterLogs(opts, "ImplementationUpdated", oldImplementationRule, newImplementationRule)
	if err != nil {
		return nil, err
	}
	return &SmartAccountImplementationUpdatedIterator{contract: _SmartAccount.contract, event: "ImplementationUpdated", logs: logs, sub: sub}, nil
}

// WatchImplementationUpdated is a free log subscription operation binding the contract event 0xaa3f731066a578e5f39b4215468d826cdd15373cbc0dfc9cb9bdc649718ef7da.
//
// Solidity: event ImplementationUpdated(address indexed oldImplementation, address indexed newImplementation)
func (_SmartAccount *SmartAccountFilterer) WatchImplementationUpdated(opts *bind.WatchOpts, sink chan<- *SmartAccountImplementationUpdated, oldImplementation []common.Address, newImplementation []common.Address) (event.Subscription, error) {

	var oldImplementationRule []interface{}
	for _, oldImplementationItem := range oldImplementation {
		oldImplementationRule = append(oldImplementationRule, oldImplementationItem)
	}
	var newImplementationRule []interface{}
	for _, newImplementationItem := range newImplementation {
		newImplementationRule = append(newImplementationRule, newImplementationItem)
	}

	logs, sub, err := _SmartAccount.contract.WatchLogs(opts, "ImplementationUpdated", oldImplementationRule, newImplementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartAccountImplementationUpdated)
				if err := _SmartAccount.contract.UnpackLog(event, "ImplementationUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseImplementationUpdated is a log parse operation binding the contract event 0xaa3f731066a578e5f39b4215468d826cdd15373cbc0dfc9cb9bdc649718ef7da.
//
// Solidity: event ImplementationUpdated(address indexed oldImplementation, address indexed newImplementation)
func (_SmartAccount *SmartAccountFilterer) ParseImplementationUpdated(log types.Log) (*SmartAccountImplementationUpdated, error) {
	event := new(SmartAccountImplementationUpdated)
	if err := _SmartAccount.contract.UnpackLog(event, "ImplementationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmartAccountModuleTransactionIterator is returned from FilterModuleTransaction and is used to iterate over the raw logs and unpacked data for ModuleTransaction events raised by the SmartAccount contract.
type SmartAccountModuleTransactionIterator struct {
	Event *SmartAccountModuleTransaction // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SmartAccountModuleTransactionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartAccountModuleTransaction)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SmartAccountModuleTransaction)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SmartAccountModuleTransactionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartAccountModuleTransactionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartAccountModuleTransaction represents a ModuleTransaction event raised by the SmartAccount contract.
type SmartAccountModuleTransaction struct {
	Module    common.Address
	To        common.Address
	Value     *big.Int
	Data      []byte
	Operation uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterModuleTransaction is a free log retrieval operation binding the contract event 0x8c014e41cffd68ba64f3e7830b8b2e4ee860509d8deab25ebbcbba2f0405e2da.
//
// Solidity: event ModuleTransaction(address module, address to, uint256 value, bytes data, uint8 operation)
func (_SmartAccount *SmartAccountFilterer) FilterModuleTransaction(opts *bind.FilterOpts) (*SmartAccountModuleTransactionIterator, error) {

	logs, sub, err := _SmartAccount.contract.FilterLogs(opts, "ModuleTransaction")
	if err != nil {
		return nil, err
	}
	return &SmartAccountModuleTransactionIterator{contract: _SmartAccount.contract, event: "ModuleTransaction", logs: logs, sub: sub}, nil
}

// WatchModuleTransaction is a free log subscription operation binding the contract event 0x8c014e41cffd68ba64f3e7830b8b2e4ee860509d8deab25ebbcbba2f0405e2da.
//
// Solidity: event ModuleTransaction(address module, address to, uint256 value, bytes data, uint8 operation)
func (_SmartAccount *SmartAccountFilterer) WatchModuleTransaction(opts *bind.WatchOpts, sink chan<- *SmartAccountModuleTransaction) (event.Subscription, error) {

	logs, sub, err := _SmartAccount.contract.WatchLogs(opts, "ModuleTransaction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartAccountModuleTransaction)
				if err := _SmartAccount.contract.UnpackLog(event, "ModuleTransaction", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseModuleTransaction is a log parse operation binding the contract event 0x8c014e41cffd68ba64f3e7830b8b2e4ee860509d8deab25ebbcbba2f0405e2da.
//
// Solidity: event ModuleTransaction(address module, address to, uint256 value, bytes data, uint8 operation)
func (_SmartAccount *SmartAccountFilterer) ParseModuleTransaction(log types.Log) (*SmartAccountModuleTransaction, error) {
	event := new(SmartAccountModuleTransaction)
	if err := _SmartAccount.contract.UnpackLog(event, "ModuleTransaction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmartAccountSmartAccountReceivedNativeTokenIterator is returned from FilterSmartAccountReceivedNativeToken and is used to iterate over the raw logs and unpacked data for SmartAccountReceivedNativeToken events raised by the SmartAccount contract.
type SmartAccountSmartAccountReceivedNativeTokenIterator struct {
	Event *SmartAccountSmartAccountReceivedNativeToken // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SmartAccountSmartAccountReceivedNativeTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartAccountSmartAccountReceivedNativeToken)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SmartAccountSmartAccountReceivedNativeToken)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SmartAccountSmartAccountReceivedNativeTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartAccountSmartAccountReceivedNativeTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartAccountSmartAccountReceivedNativeToken represents a SmartAccountReceivedNativeToken event raised by the SmartAccount contract.
type SmartAccountSmartAccountReceivedNativeToken struct {
	Sender common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSmartAccountReceivedNativeToken is a free log retrieval operation binding the contract event 0x00d05ab44e279ac59e855cb75dc2ae23b200ad994797b6f1f028f96a46ecce02.
//
// Solidity: event SmartAccountReceivedNativeToken(address indexed sender, uint256 indexed value)
func (_SmartAccount *SmartAccountFilterer) FilterSmartAccountReceivedNativeToken(opts *bind.FilterOpts, sender []common.Address, value []*big.Int) (*SmartAccountSmartAccountReceivedNativeTokenIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _SmartAccount.contract.FilterLogs(opts, "SmartAccountReceivedNativeToken", senderRule, valueRule)
	if err != nil {
		return nil, err
	}
	return &SmartAccountSmartAccountReceivedNativeTokenIterator{contract: _SmartAccount.contract, event: "SmartAccountReceivedNativeToken", logs: logs, sub: sub}, nil
}

// WatchSmartAccountReceivedNativeToken is a free log subscription operation binding the contract event 0x00d05ab44e279ac59e855cb75dc2ae23b200ad994797b6f1f028f96a46ecce02.
//
// Solidity: event SmartAccountReceivedNativeToken(address indexed sender, uint256 indexed value)
func (_SmartAccount *SmartAccountFilterer) WatchSmartAccountReceivedNativeToken(opts *bind.WatchOpts, sink chan<- *SmartAccountSmartAccountReceivedNativeToken, sender []common.Address, value []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _SmartAccount.contract.WatchLogs(opts, "SmartAccountReceivedNativeToken", senderRule, valueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartAccountSmartAccountReceivedNativeToken)
				if err := _SmartAccount.contract.UnpackLog(event, "SmartAccountReceivedNativeToken", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSmartAccountReceivedNativeToken is a log parse operation binding the contract event 0x00d05ab44e279ac59e855cb75dc2ae23b200ad994797b6f1f028f96a46ecce02.
//
// Solidity: event SmartAccountReceivedNativeToken(address indexed sender, uint256 indexed value)
func (_SmartAccount *SmartAccountFilterer) ParseSmartAccountReceivedNativeToken(log types.Log) (*SmartAccountSmartAccountReceivedNativeToken, error) {
	event := new(SmartAccountSmartAccountReceivedNativeToken)
	if err := _SmartAccount.contract.UnpackLog(event, "SmartAccountReceivedNativeToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
