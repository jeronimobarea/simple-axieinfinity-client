package slp

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Checkpoint is a paid mutator transaction binding the contract method 0xd3392ddf.
//
// Solidity: function checkpoint(address _owner, uint256 _amount, uint256 _createdAt, bytes _signature) returns(uint256 _balance)
func (_Slp *SlpTransactor) Checkpoint(opts *bind.TransactOpts, _owner common.Address, _amount *big.Int, _createdAt *big.Int, _signature []byte) (*types.Transaction, error) {
	return _Slp.contract.Transact(opts, "checkpoint", _owner, _amount, _createdAt, _signature)
}

// Checkpoint is a paid mutator transaction binding the contract method 0xd3392ddf.
//
// Solidity: function checkpoint(address _owner, uint256 _amount, uint256 _createdAt, bytes _signature) returns(uint256 _balance)
func (_Slp *SlpSession) Checkpoint(_owner common.Address, _amount *big.Int, _createdAt *big.Int, _signature []byte) (*types.Transaction, error) {
	return _Slp.Contract.Checkpoint(&_Slp.TransactOpts, _owner, _amount, _createdAt, _signature)
}

// Checkpoint is a paid mutator transaction binding the contract method 0xd3392ddf.
//
// Solidity: function checkpoint(address _owner, uint256 _amount, uint256 _createdAt, bytes _signature) returns(uint256 _balance)
func (_Slp *SlpTransactorSession) Checkpoint(_owner common.Address, _amount *big.Int, _createdAt *big.Int, _signature []byte) (*types.Transaction, error) {
	return _Slp.Contract.Checkpoint(&_Slp.TransactOpts, _owner, _amount, _createdAt, _signature)
}
