const DelegatableDai = artifacts.require("./DelegatableDai.sol");
const web3js = require('web3'); // web3 v1

const delay = ms => new Promise(resolve => setTimeout(resolve, ms));

const asyncFunction = async function asyncFunction() {

  const fromAddr = '0x2b522cabe9950d1153c26c1b399b293caa99fcf9';
  const toAddr = '0x3644b986b3f5ba3cb8d5627a22465942f8e06d09';
  const delegate = '0x47a793d7d0aa5727095c3fe132a6c1a46804c8d2';
  const token = '0x9ea55b285ddeab2db64d4c2624efec1aa5875283';

  await DelegatableDai.at(token).then(function(inst) {
    dai = inst;
  });

  const amountToRecipient = 100;
  const amountToDelegate = 2;
  const nonce = web3.eth.getTransactionCount(fromAddr);

  // signature will consists of V, R, S
  // r = signature[0:64]
  // s = signature[64:128]
  // v = signature[128:130]
  // https://github.com/ethereum/wiki/wiki/JavaScript-API#web3ethsign
  const hash = await dai.transferPreSignedHashing(token, toAddr, amountToRecipient, amountToDelegate, nonce, { from: fromAddr });
  const signature = web3.eth.sign(fromAddr, hash);

  const fromAddrBeforeDaiBalance = await dai.balanceOf(fromAddr);
  const toAddrBeforeDaiBalance = await dai.balanceOf(toAddr);
  const delegateBeforeDaiBalance = await dai.balanceOf(delegate);
  const fromAddrBeforeEthBalance = await web3.eth.getBalance(fromAddr);
  const toAddrBeforeEthBalance = await web3.eth.getBalance(toAddr);
  const delegateBeforeEthBalance = await web3.eth.getBalance(delegate);
  console.log('fromAddr before Dai balance', fromAddrBeforeDaiBalance);
  console.log('toAddr before Dai balance', toAddrBeforeDaiBalance);
  console.log('delegate before Dai balance', delegateBeforeDaiBalance);
  console.log('fromAddr before Eth balance', fromAddrBeforeEthBalance);
  console.log('toAddr before Eth balance', toAddrBeforeEthBalance);
  console.log('delegate before Eth balance', delegateBeforeEthBalance);

  const test = await dai.transferPreSigned(signature, toAddr, amountToRecipient, amountToDelegate, nonce, { from: delegate });

  const fromAddrAfterDaiBalance = await dai.balanceOf(fromAddr);
  const toAddrAfterDaiBalance = await dai.balanceOf(toAddr);
  const delegateAfterDaiBalance = await dai.balanceOf(delegate);
  const fromAddrAfterEthBalance = await web3.eth.getBalance(fromAddr);
  const toAddrAfterEthBalance = await web3.eth.getBalance(toAddr);
  const delegateAfterEthBalance = await web3.eth.getBalance(delegate);
  console.log('fromAddr after Dai balance', fromAddrAfterDaiBalance);
  console.log('toAddr after Dai balance', toAddrAfterDaiBalance);
  console.log('delegate after Dai balance', delegateAfterDaiBalance);
  console.log('fromAddr after Eth balance', fromAddrAfterEthBalance);
  console.log('toAddr after Eth balance', toAddrAfterEthBalance);
  console.log('delegate after Eth balance', delegateAfterEthBalance);
};

module.exports = function(callback) {
  asyncFunction();
};

// Accounts
// (0) 0x2b522cabe9950d1153c26c1b399b293caa99fcf9 (~500 ETH)
// (1) 0x3644b986b3f5ba3cb8d5627a22465942f8e06d09 (~500 ETH)
// (2) 0x9e8f633d0c46ed7170ef3b30e291c64a91a49c7e (~500 ETH)
// (3) 0xd1c3bf2f2fd296249228734299290cf8616c1e7c (~500 ETH)
// (4) 0x47a793d7d0aa5727095c3fe132a6c1a46804c8d2 (~500 ETH)
// (5) 0x0d95ebb4874f17157e40635c19dbc6e9b0bfdb03 (~500 ETH)
// (6) 0x5243b5970f327c328b2739dec88abc46fae8931a (~500 ETH)
// (7) 0xe1a1d3637ee02391ac4035e72456ca7448c73fd4 (~500 ETH)
// (8) 0x1cf1919d91cebab2e56a5c0cc7180bb54ed4f3f6 (~500 ETH)
// (9) 0x0339f40494265dcdd5ab47a3b55183a28952744b (~500 ETH)