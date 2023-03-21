import { expect } from "chai";
import { ethers } from "hardhat";
import {
  SmartAccount,
  SmartAccountFactory,
  EntryPoint,
  MockToken,
  MultiSend,
  StorageSetter,
  DefaultCallbackHandler,
  SignMessageLib,
} from "../../typechain";
import { encodeTransfer, encodeSignMessage } from "../testUtils";
import {
  SafeTransaction,
  Transaction,
  FeeRefund,
  safeSignTypedData,
  buildContractSignature,
  buildSafeTransaction,
  EOA_CONTROLLED_FLOW,
} from "../../../src/utils/execution";
import { deployContract } from "../../utils/setupHelper";

export const AddressZero = ethers.constants.AddressZero;

describe("EIP-1271 Signatures Tests", function () {
  let baseImpl: SmartAccount;
  let walletFactory: SmartAccountFactory;
  let entryPoint: EntryPoint;
  let token: MockToken;
  let multiSend: MultiSend;
  let storage: StorageSetter;
  let signMessageLib: SignMessageLib;
  let owner: string;
  let bob: string;
  let charlie: string;
  let hacker: string;
  let signerSmartAccount: any;
  let mainSmartAccount: any;
  let handler: DefaultCallbackHandler;
  let accounts: any;
  let smartAccountInitialNativeTokenBalance: any;

  /* const domainType = [
    { name: "name", type: "string" },
    { name: "version", type: "string" },
    { name: "verifyingContract", type: "address" },
    { name: "salt", type: "bytes32" },
  ]; */

  beforeEach(async () => {
    accounts = await ethers.getSigners();

    owner = await accounts[0].getAddress();
    bob = await accounts[1].getAddress();
    charlie = await accounts[2].getAddress();
    hacker = await accounts[3].getAddress();
    // const owner = "0x7306aC7A32eb690232De81a9FFB44Bb346026faB";

    const EntryPoint = await ethers.getContractFactory("EntryPoint");
    entryPoint = await EntryPoint.deploy();
    await entryPoint.deployed();
    // console.log("Entry point deployed at: ", entryPoint.address);

    const DefaultHandler = await ethers.getContractFactory(
      "DefaultCallbackHandler"
    );
    handler = await DefaultHandler.deploy();
    await handler.deployed();
    // console.log("Default callback handler deployed at: ", handler.address);

    const BaseImplementation = await ethers.getContractFactory("SmartAccount");
    baseImpl = await BaseImplementation.deploy(entryPoint.address);
    await baseImpl.deployed();
    // console.log("base wallet impl deployed at: ", baseImpl.address);

    const WalletFactory = await ethers.getContractFactory(
      "SmartAccountFactory"
    );
    walletFactory = await WalletFactory.deploy(baseImpl.address);
    await walletFactory.deployed();
    // console.log("wallet factory deployed at: ", walletFactory.address);

    const MockToken = await ethers.getContractFactory("MockToken");
    token = await MockToken.deploy();
    await token.deployed();
    // console.log("Test token deployed at: ", token.address);

    const Storage = await ethers.getContractFactory("StorageSetter");
    storage = await Storage.deploy();
    // console.log("storage setter contract deployed at: ", storage.address);

    const MultiSend = await ethers.getContractFactory("MultiSend");
    multiSend = await MultiSend.deploy();
    // console.log("Multisend helper contract deployed at: ", multiSend.address);

    const SignMessageLib = await ethers.getContractFactory("SignMessageLib");
    signMessageLib = await SignMessageLib.deploy();

    // console.log("mint tokens to owner address..");
    await token.mint(owner, ethers.utils.parseEther("1000000"));

    const deployWalletIndex = 0;

    // console.log("Owner of Signer Smart Account is ", owner);
    // Deploy Signer Smart Account owned by Owner
    const signerSmartAccountAddress =
      await walletFactory.getAddressForCounterFactualAccount(
        owner,
        deployWalletIndex
      );

    await walletFactory.deployCounterFactualAccount(owner, deployWalletIndex);

    signerSmartAccount = await ethers.getContractAt(
      "contracts/smart-contract-wallet/SmartAccount.sol:SmartAccount",
      signerSmartAccountAddress
    );

    // Deploy Main Smart Account owned by SignerSmartAccount
    const mainSmartAccountAddress =
      await walletFactory.getAddressForCounterFactualAccount(
        signerSmartAccountAddress,
        deployWalletIndex
      );

    await walletFactory.deployCounterFactualAccount(
      signerSmartAccountAddress,
      deployWalletIndex
    );

    mainSmartAccount = await ethers.getContractAt(
      "contracts/smart-contract-wallet/SmartAccount.sol:SmartAccount",
      mainSmartAccountAddress
    );

    smartAccountInitialNativeTokenBalance = ethers.utils.parseEther("5");

    await accounts[1].sendTransaction({
      from: bob,
      to: signerSmartAccountAddress,
      value: smartAccountInitialNativeTokenBalance,
    });

    await accounts[1].sendTransaction({
      from: bob,
      to: mainSmartAccountAddress,
      value: smartAccountInitialNativeTokenBalance,
    });
  });

  it("Can execute tx with a valid 1271 signature", async function () {
    // transfer 100 tokens to Main Smart Account and Signer Smart Account
    await token
      .connect(accounts[0])
      .transfer(mainSmartAccount.address, ethers.utils.parseEther("100"));

    await token
      .connect(accounts[0])
      .transfer(signerSmartAccount.address, ethers.utils.parseEther("100"));

    const tokensToBeTransferred = ethers.utils.parseEther("10");

    const safeTx: SafeTransaction = buildSafeTransaction({
      to: token.address,
      // value: ethers.utils.parseEther("1"),
      data: encodeTransfer(charlie, tokensToBeTransferred.toString()),
      nonce: await mainSmartAccount.getNonce(EOA_CONTROLLED_FLOW),
    });

    const transaction: Transaction = {
      to: safeTx.to,
      value: safeTx.value,
      data: safeTx.data,
      operation: safeTx.operation,
      targetTxGas: safeTx.targetTxGas,
    };
    const refundInfo: FeeRefund = {
      baseGas: safeTx.baseGas,
      gasPrice: safeTx.gasPrice,
      tokenGasPriceFactor: safeTx.tokenGasPriceFactor,
      gasToken: safeTx.gasToken,
      refundReceiver: safeTx.refundReceiver,
    };

    const chainId = await mainSmartAccount.getChainId();

    // BUILD 1271 SIGNATURE BY SIGNER SMART ACCOUNT

    const { signer, data } = await safeSignTypedData(
      accounts[0], // owner
      mainSmartAccount,
      safeTx,
      chainId
    );

    const signature = buildContractSignature(signerSmartAccount.address, data);

    await expect(
      mainSmartAccount
        .connect(accounts[1])
        .execTransaction_S6W(transaction, refundInfo, signature)
    ).to.emit(mainSmartAccount, "ExecutionSuccess");

    expect(await token.balanceOf(charlie)).to.equal(tokensToBeTransferred);
  });

  it("Fallback handler reverts if called directly", async function () {
    const dataHash = ethers.utils.keccak256("0xbaddad");
    await expect(handler.isValidSignature(dataHash, "0x")).to.be.reverted;
  });

  // TODO: move from here to fallback-handler.specs.ts
  it("Fallback handler handles tokensReceived", async () => {
    await handler.callStatic.tokensReceived(
      AddressZero,
      AddressZero,
      AddressZero,
      0,
      "0x",
      "0x"
    );
  });

  it("Fallback handler handles onERC721Received", async () => {
    await expect(
      await handler.callStatic.onERC721Received(
        AddressZero,
        AddressZero,
        0,
        "0x"
      )
    ).to.be.eq("0x150b7a02");
  });

  it("Fallback handler handles onERC1155Received", async () => {
    await expect(
      await handler.callStatic.onERC1155Received(
        AddressZero,
        AddressZero,
        0,
        0,
        "0x"
      )
    ).to.be.eq("0xf23a6e61");
  });

  it("Fallback handler handles onERC1155BatchReceived", async () => {
    await expect(
      await handler.callStatic.onERC1155BatchReceived(
        AddressZero,
        AddressZero,
        [],
        [],
        "0x"
      )
    ).to.be.eq("0xbc197c81");
  });

  it("Fallback handler returns 0xffffffff if the message has not been signed", async function () {
    const dataHash = ethers.utils.keccak256("0xdeafbeef");

    const smartAccountWithHandlerInterface = await ethers.getContractAt(
      "contracts/smart-contract-wallet/handler/DefaultCallbackHandler.sol:DefaultCallbackHandler",
      signerSmartAccount.address
    );

    const value = await smartAccountWithHandlerInterface.callStatic[
      "isValidSignature(bytes32,bytes)"
    ](dataHash, "0x");
    const notMagicValue = "0xffffffff";
    expect(value).to.be.equal(notMagicValue);
  });

  it("Fallback handler returns 0xffffffff if signature is not valid", async function () {
    const dataHash = ethers.utils.keccak256("0xdeafbeef");
    const invalidSignature = "0xabcdefdecafcafe0";

    const smartAccountWithHandlerInterface = await ethers.getContractAt(
      "contracts/smart-contract-wallet/handler/DefaultCallbackHandler.sol:DefaultCallbackHandler",
      signerSmartAccount.address
    );

    const value = await smartAccountWithHandlerInterface.callStatic[
      "isValidSignature(bytes32,bytes)"
    ](dataHash, invalidSignature);
    const notMagicValue = "0xffffffff";
    expect(value).to.be.equal(notMagicValue);
  });

  it("Fallback handler returns magic value if message has been signed", async function () {
    const delegateCall = 1;
    const dataToSign = "0xdeafbeefdecaf0";

    // prepare tx to call signMessageLib.signMessage from signerSmartAccount
    // through delegate call
    const safeTx: SafeTransaction = buildSafeTransaction({
      to: signMessageLib.address,
      data: encodeSignMessage(dataToSign),
      operation: delegateCall,
      nonce: await signerSmartAccount.getNonce(EOA_CONTROLLED_FLOW),
    });

    const transaction: Transaction = {
      to: safeTx.to,
      value: safeTx.value,
      data: safeTx.data,
      operation: safeTx.operation,
      targetTxGas: safeTx.targetTxGas,
    };
    const refundInfo: FeeRefund = {
      baseGas: safeTx.baseGas,
      gasPrice: safeTx.gasPrice,
      tokenGasPriceFactor: safeTx.tokenGasPriceFactor,
      gasToken: safeTx.gasToken,
      refundReceiver: safeTx.refundReceiver,
    };

    const chainId = await mainSmartAccount.getChainId();

    const { signer, data } = await safeSignTypedData(
      accounts[0],
      signerSmartAccount,
      safeTx,
      chainId
    );

    let signature = "0x";
    signature += data.slice(2);

    const txSignMessage = await signerSmartAccount
      .connect(accounts[0])
      .execTransaction_S6W(transaction, refundInfo, signature);
    const receipt = await txSignMessage.wait();

    const smartAccountWithHandlerInterface = await ethers.getContractAt(
      "contracts/smart-contract-wallet/handler/DefaultCallbackHandler.sol:DefaultCallbackHandler",
      signerSmartAccount.address
    );

    const dataHash = await smartAccountWithHandlerInterface.callStatic[
      "getMessageHash(bytes)"
    ](dataToSign);

    const eip1271MagicValue = "0x1626ba7e";
    const value = await smartAccountWithHandlerInterface.callStatic[
      "isValidSignature(bytes32,bytes)"
    ](dataHash, "0x");
    expect(value).to.be.equal(eip1271MagicValue);
  });

  it("Fallback handler returns magic value if correct signature has been provided directly to Smart Account", async function () {
    const message = "Some message from dApp";
    const signature = await accounts[0].signMessage(message);

    // since .signMessage actually signs the message hash prepended by
    // \x19Ethereum Signed Message:\n" and the length of the message
    // we use .hashMessage to get message hash to verify against
    const messageHash = ethers.utils.hashMessage(message);

    const smartAccountWithHandlerInterface = await ethers.getContractAt(
      "contracts/smart-contract-wallet/handler/DefaultCallbackHandler.sol:DefaultCallbackHandler",
      signerSmartAccount.address
    );

    const eip1271MagicValue = "0x1626ba7e";
    const value = await smartAccountWithHandlerInterface.callStatic[
      "isValidSignature(bytes32,bytes)"
    ](messageHash, signature);
    expect(value).to.be.equal(eip1271MagicValue);
  });

  it("Wont let the transaction to go through with manipulated signer contract address", async function () {
    const deployWalletIndex = 1;
    const BaseImplementation = await ethers.getContractFactory("SmartAccount");

    // Deploy Signer Smart Account 2 owned by Owner
    const signerSmartAccount2Address =
      await walletFactory.getAddressForCounterFactualAccount(
        owner,
        deployWalletIndex
      );

    await walletFactory.deployCounterFactualAccount(owner, deployWalletIndex);

    const signerSmartAccount2 = await ethers.getContractAt(
      "contracts/smart-contract-wallet/SmartAccount.sol:SmartAccount",
      signerSmartAccount2Address
    );

    // Deploy Main Smart Account 2 owned by SignerSmartAccount 2
    const mainSmartAccount2Address =
      await walletFactory.getAddressForCounterFactualAccount(
        signerSmartAccount2Address,
        deployWalletIndex
      );

    await walletFactory.deployCounterFactualAccount(
      signerSmartAccount2Address,
      deployWalletIndex
    );

    const mainSmartAccount2 = await ethers.getContractAt(
      "contracts/smart-contract-wallet/SmartAccount.sol:SmartAccount",
      mainSmartAccount2Address
    );

    await token
      .connect(accounts[0])
      .transfer(mainSmartAccount.address, ethers.utils.parseEther("100"));

    await token
      .connect(accounts[0])
      .transfer(signerSmartAccount.address, ethers.utils.parseEther("100"));

    await token
      .connect(accounts[0])
      .transfer(signerSmartAccount2.address, ethers.utils.parseEther("100"));

    await token
      .connect(accounts[0])
      .transfer(mainSmartAccount2.address, ethers.utils.parseEther("100"));

    expect(await token.balanceOf(charlie)).to.equal(0);

    const tokensToBeTransferred = ethers.utils.parseEther("10");

    // TX TO TRANSFER 10 tokens FROM mainSmartAccount to Charlie
    const safeTx: SafeTransaction = buildSafeTransaction({
      to: token.address,
      data: encodeTransfer(charlie, tokensToBeTransferred.toString()),
      nonce: await mainSmartAccount.getNonce(EOA_CONTROLLED_FLOW),
    });

    const transaction: Transaction = {
      to: safeTx.to,
      value: safeTx.value,
      data: safeTx.data,
      operation: safeTx.operation,
      targetTxGas: safeTx.targetTxGas,
    };
    const refundInfo: FeeRefund = {
      baseGas: safeTx.baseGas,
      gasPrice: safeTx.gasPrice,
      tokenGasPriceFactor: safeTx.tokenGasPriceFactor,
      gasToken: safeTx.gasToken,
      refundReceiver: safeTx.refundReceiver,
    };

    const chainId = await mainSmartAccount.getChainId();

    // BUILD 1271 SIGNATURE BY SIGNER SMART ACCOUNT

    const { signer, data } = await safeSignTypedData(
      accounts[0], // owner
      mainSmartAccount,
      safeTx,
      chainId
    );

    const signature = buildContractSignature(signerSmartAccount.address, data);

    // MANIPULATE SIGNATURE TO SET signerContract2 as verifier
    const addressToInsert = signerSmartAccount2.address.slice(2);

    const manipulatedSignature = signature.replace(
      signature.substring(26, 66),
      addressToInsert
    );

    // Expect can not use this signature on main smart account 2, even despite
    // it is owned by signer smart account 2, that is owned by the owner (original signer)
    await expect(
      mainSmartAccount2
        .connect(accounts[1])
        .execTransaction_S6W(transaction, refundInfo, manipulatedSignature)
    ).to.be.revertedWith("WrongContractSignature");

    expect(await token.balanceOf(charlie)).to.equal(0);
  });

  it("0x exploit 1271 | Reverts if trying to use 1271 signature instead of EOA signature", async function () {
    // i.e. trying to call isVslidSignature from EOA won't return magic value
    // See https://samczsun.com/the-0x-vulnerability-explained/
    await token
      .connect(accounts[0])
      .transfer(signerSmartAccount.address, ethers.utils.parseEther("100"));

    const tokensToBeTransferred = ethers.utils.parseEther("10");

    const safeTx: SafeTransaction = buildSafeTransaction({
      to: token.address,
      // value: ethers.utils.parseEther("1"),
      data: encodeTransfer(charlie, tokensToBeTransferred.toString()),
      nonce: await mainSmartAccount.getNonce(EOA_CONTROLLED_FLOW),
    });

    const transaction: Transaction = {
      to: safeTx.to,
      value: safeTx.value,
      data: safeTx.data,
      operation: safeTx.operation,
      targetTxGas: safeTx.targetTxGas,
    };
    const refundInfo: FeeRefund = {
      baseGas: safeTx.baseGas,
      gasPrice: safeTx.gasPrice,
      tokenGasPriceFactor: safeTx.tokenGasPriceFactor,
      gasToken: safeTx.gasToken,
      refundReceiver: safeTx.refundReceiver,
    };

    const chainId = await mainSmartAccount.getChainId();

    // BUILD 1271 SIGNATURE BY OWNER

    const { signer, data } = await safeSignTypedData(
      accounts[0], // owner
      mainSmartAccount,
      safeTx,
      chainId
    );

    const fakeSignature = buildContractSignature(accounts[0].address, data);

    await expect(
      signerSmartAccount
        .connect(accounts[1])
        .execTransaction_S6W(transaction, refundInfo, fakeSignature)
    ).to.be.reverted;
  });

  it("0x exploit 1271 | Reverts if isValidSignature changes the state", async function () {
    const deployWalletIndex = 0;
    const BaseImplementation = await ethers.getContractFactory("SmartAccount");

    const source = `
            contract Test {
                bool public changeState;
                uint256 public nonce;
                function isValidSignature(bytes32 _dataHash, bytes memory _signature) public returns (bytes4) {
                    if (changeState) {
                        nonce = nonce + 1;
                    }
                    return 0x1626ba7e;
                }
    
                function shouldChangeState(bool value) public {
                    changeState = value;
                }
            }`;
    const testValidator = await deployContract(accounts[0], source);
    await testValidator.shouldChangeState(true);

    // Deploy Main Smart Account 2 owned by testValidator
    const mainSmartAccount2Address =
      await walletFactory.getAddressForCounterFactualAccount(
        testValidator.address,
        deployWalletIndex
      );

    await walletFactory.deployCounterFactualAccount(
      testValidator.address,
      deployWalletIndex
    );

    const mainSmartAccount2 = await ethers.getContractAt(
      "contracts/smart-contract-wallet/SmartAccount.sol:SmartAccount",
      mainSmartAccount2Address
    );

    await token
      .connect(accounts[0])
      .transfer(mainSmartAccount2.address, ethers.utils.parseEther("100"));

    expect(await token.balanceOf(charlie)).to.equal(0);

    const tokensToBeTransferred = ethers.utils.parseEther("10");

    const safeTx: SafeTransaction = buildSafeTransaction({
      to: token.address,
      data: encodeTransfer(charlie, tokensToBeTransferred.toString()),
      nonce: await mainSmartAccount2.getNonce(EOA_CONTROLLED_FLOW),
    });

    const transaction: Transaction = {
      to: safeTx.to,
      value: safeTx.value,
      data: safeTx.data,
      operation: safeTx.operation,
      targetTxGas: safeTx.targetTxGas,
    };
    const refundInfo: FeeRefund = {
      baseGas: safeTx.baseGas,
      gasPrice: safeTx.gasPrice,
      tokenGasPriceFactor: safeTx.tokenGasPriceFactor,
      gasToken: safeTx.gasToken,
      refundReceiver: safeTx.refundReceiver,
    };

    const chainId = await mainSmartAccount2.getChainId();

    // BUILD 1271 SIGNATURE BY test validator

    const { signer, data } = await safeSignTypedData(
      accounts[0], // owner
      mainSmartAccount2,
      safeTx,
      chainId
    );

    const signature = buildContractSignature(testValidator.address, data);

    await expect(
      mainSmartAccount2
        .connect(accounts[1])
        .execTransaction_S6W(transaction, refundInfo, signature)
    ).to.be.reverted;
    expect(await token.balanceOf(charlie)).to.equal(0);

    await testValidator.shouldChangeState(false);

    await expect(
      mainSmartAccount2
        .connect(accounts[1])
        .execTransaction_S6W(transaction, refundInfo, signature)
    ).to.emit(mainSmartAccount2, "ExecutionSuccess");

    expect(await token.balanceOf(charlie)).to.equal(tokensToBeTransferred);
  });
});