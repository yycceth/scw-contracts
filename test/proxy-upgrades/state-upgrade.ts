import { expect } from "chai";
import { ethers } from "hardhat";
import {
  SmartWallet,
  WalletFactory,
  EntryPoint__factory,
  EntryPoint,
  MockToken,
  MultiSend,
  StorageSetter,
  DefaultCallbackHandler,
} from "../../typechain";
import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { encodeTransfer, encodeTransferFrom } from "../smart-wallet/testUtils";
import {
  buildContractCall,
  MetaTransaction,
  SafeTransaction,
  Transaction,
  FeeRefund,
  executeTx,
  safeSignTypedData,
  safeSignMessage,
  buildSafeTransaction,
  executeContractCallWithSigners,
} from "../../src/utils/execution";
import { buildMultiSendSafeTx } from "../../src/utils/multisend";

export async function deployEntryPoint(
  provider = ethers.provider
): Promise<EntryPoint> {
  const epf = await (await ethers.getContractFactory("EntryPoint")).deploy();
  return EntryPoint__factory.connect(epf.address, provider.getSigner());
}

describe("Upgradeability", function () {
  // TODO
  let baseImpl: SmartWallet;
  let walletFactory: WalletFactory;
  let entryPoint: EntryPoint;
  let token: MockToken;
  let multiSend: MultiSend;
  let storage: StorageSetter;
  let owner: string;
  let bob: string;
  let charlie: string;
  let habibi: string;
  let userSCW: any;
  let handler: DefaultCallbackHandler;
  let accounts: any;

  before(async () => {
    accounts = await ethers.getSigners();
    owner = await accounts[0].getAddress();
    bob = await accounts[1].getAddress();
    charlie = await accounts[2].getAddress();
    habibi = await accounts[3].getAddress();
    // const owner = "0x7306aC7A32eb690232De81a9FFB44Bb346026faB";

    const EntryPoint = await ethers.getContractFactory("EntryPoint");
    entryPoint = await EntryPoint.deploy();
    await entryPoint.deployed();
    console.log("Entry point deployed at: ", entryPoint.address);

    /* const DefaultHandler = await ethers.getContractFactory(
      "DefaultCallbackHandler"
    );
    handler = await DefaultHandler.deploy();
    await handler.deployed();
    console.log("Default callback handler deployed at: ", handler.address); */

    const BaseImplementation = await ethers.getContractFactory("SmartAccount");
    baseImpl = await BaseImplementation.deploy(entryPoint.address);
    await baseImpl.deployed();
    console.log("base wallet impl deployed at: ", baseImpl.address);

    const WalletFactory = await ethers.getContractFactory(
      "SmartAccountFactory"
    );
    walletFactory = await WalletFactory.deploy(baseImpl.address);
    await walletFactory.deployed();
    console.log("wallet factory deployed at: ", walletFactory.address);

    const MockToken = await ethers.getContractFactory("MockToken");
    token = await MockToken.deploy();
    await token.deployed();
    console.log("Test token deployed at: ", token.address);

    const Storage = await ethers.getContractFactory("StorageSetter");
    storage = await Storage.deploy();
    console.log("storage setter contract deployed at: ", storage.address);

    const MultiSend = await ethers.getContractFactory("MultiSend");
    multiSend = await MultiSend.deploy();
    console.log("Multisend helper contract deployed at: ", multiSend.address);

    console.log("mint tokens to owner address..");
    await token.mint(owner, ethers.utils.parseEther("1000000"));
  });

  it("Should deploy a wallet and validate entrypoint", async function () {
    const expected = await walletFactory.getAddressForCounterfactualAccount(
      owner,
      0
    );
    console.log("deploying new wallet..expected address: ", expected);

    await expect(walletFactory.deployCounterFactualAccount(owner, 0))
      .to.emit(walletFactory, "AccountCreation")
      .withArgs(expected, owner, 0);

    userSCW = await ethers.getContractAt(
      "contracts/smart-contract-wallet/SmartAccount.sol:SmartAccount",
      expected
    );

    const entryPointAddress = await userSCW.entryPoint();
    expect(entryPointAddress).to.equal(entryPoint.address);

    await accounts[1].sendTransaction({
      from: bob,
      to: expected,
      value: ethers.utils.parseEther("5"),
    });
  });

  it("Should deploy new implementation and upgrade", async function () {
    // Note that -> To upgrade an entry point we need to deploy new implementation
    // but to update implementation previous entry point can use in it's constructor.
    const BaseImplementation6 = await ethers.getContractFactory(
      "SmartAccount6"
    );
    // deployEntryPoint
    const baseImpl6 = await BaseImplementation6.deploy(entryPoint.address);
    await baseImpl6.deployed();
    console.log("base wallet upgraded impl deployed at: ", baseImpl6.address);

    await expect(
      userSCW.connect(accounts[0]).updateImplementation(baseImpl6.address)
    ).to.emit(userSCW, "ImplementationUpdated");

    userSCW = await ethers.getContractAt(
      "contracts/smart-contract-wallet/test/upgrades/SmartAccount6.sol:SmartAccount6",
      userSCW.address
    );

    // must reinit asap!
    await userSCW.reinit(habibi);

    const entryPointAddress = await userSCW.entryPoint();
    expect(entryPointAddress).to.equal(entryPoint.address);

    const friend = await userSCW.friend();
    expect(friend).to.equal(habibi);
  });

  it("should send a single transacton (EIP712 sign)", async function () {
    await token
      .connect(accounts[0])
      .transfer(userSCW.address, ethers.utils.parseEther("100"));

    const safeTx: SafeTransaction = buildSafeTransaction({
      to: token.address,
      // value: ethers.utils.parseEther("1"),
      data: encodeTransfer(charlie, ethers.utils.parseEther("10").toString()),
      nonce: await userSCW.getNonce(1),
    });

    const chainId = await userSCW.getChainId();
    const { signer, data } = await safeSignTypedData(
      accounts[0],
      userSCW,
      safeTx,
      chainId
    );

    console.log(safeTx);

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

    let signature = "0x";
    signature += data.slice(2);
    await expect(
      userSCW
        .connect(accounts[0])
        .execTransaction_S6W(transaction, refundInfo, signature)
    ).to.emit(userSCW, "ExecutionSuccess");

    expect(await token.balanceOf(charlie)).to.equal(
      ethers.utils.parseEther("10")
    );
  });

  it("should enable only friend to transfer limited eth once", async function () {
    await expect(
      userSCW
        .connect(accounts[1])
        .transferByFriend(bob, ethers.utils.parseEther("1"))
    ).to.be.revertedWith("only friend!");

    await expect(
      userSCW
        .connect(accounts[3])
        .transferByFriend(bob, ethers.utils.parseEther("2"))
    ).to.be.revertedWith("can't exceed");

    await userSCW
      .connect(accounts[3])
      .transferByFriend(bob, ethers.utils.parseEther("1"));

    await expect(
      userSCW
        .connect(accounts[3])
        .transferByFriend(bob, ethers.utils.parseEther("0.5"))
    ).to.be.revertedWith("only once during urgency");
  });
});
