const DelegatableDai = artifacts.require("./DelegatableDai.sol");
const web3js = require('web3'); // web3 v1

const delay = ms => new Promise(resolve => setTimeout(resolve, ms));

const asyncFunction = async function asyncFunction() {

  const fromAddr = '0x2b522cabe9950d1153c26c1b399b293caa99fcf9';
  const toAddr = '0x3644b986b3f5ba3cb8d5627a22465942f8e06d09';
  const delegate = '0x47a793d7d0aa5727095c3fe132a6c1a46804c8d2';
  const token = '0x87313e56553eb21b4c46908defe47ad018bbfa64';

  await DelegatableDai.at(token).then(function(inst) {
    dai = inst;
  });
  const amountToRecipient = 100;
  const amountToDelegate = 1;
  const nonce = web3.eth.getTransactionCount(fromAddr);

  // signature will consists of V, R, S
  // r = signature[0:64]
  // s = signature[64:128]
  // v = signature[128:130]
  // https://github.com/ethereum/wiki/wiki/JavaScript-API#web3ethsign
  const signature = web3.eth.sign(fromAddr, web3js.utils.soliditySha3(nonce, fromAddr, toAddr, delegate, amountToRecipient, amountToDelegate, token));
  console.log('signature', signature);

  dai.transferPreSigned(signature, toAddr, amountToRecipient, amountToDelegate, nonce, { from: delegate });
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