// SPDX-License-Identifier: Apache-2.0
pragma solidity 0.8.17;
import {IHooks} from "./IHooks.sol";
import "hardhat/console.sol";

contract SpendingLimitsModule is IHooks {
    // smartAccount=>tokenAddress=>{Struct}

    function preHook(
        address target,
        uint256 value,
        bytes memory data
    ) external {
        console.log("preHook called at", address(this));
    }

    function postHook(
        address target,
        uint256 value,
        bytes memory data
    ) external {
        console.log("postHook called at", address(this));
    }
}
