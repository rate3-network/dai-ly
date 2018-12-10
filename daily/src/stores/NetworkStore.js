import { configure, observable, action } from 'mobx';
import Web3 from 'web3';
import axios from 'axios';
import contractRaw from './built.json';
import { toTokenAmount, fromTokenAmount } from './convert';

configure({ enforceActions: 'always' }); // don't allow state modifications outside actions

class NetworkStore {
  constructor(rootStore) {
    this.rootStore = rootStore;
    window.axios = axios;
  }
  @observable hello = 'gello';
  @observable web3 = null;

  @observable response = {};
  /**
   * status
   * -1 idle
   * 0 loading
   * 2** success
   * 4** or 5** fail
   */
  @observable status = -1;
  @observable receiver = '0x9bd961ef9e0a0a4cd9234608d2b1d10ba7a1352b';
  @observable speed = 0;
  @observable amount = 1;

  @observable submittedHash = '';

  @observable proxyContract = {};
  @observable historyEvents = [];
  @observable balance = '';
  @action
  setBalance(value) {
    this.balance = value;
  }
  @action
  setHistoryEvents(value) {
    this.historyEvents = value;
  }
  @action
  setProxyContract(value) {
    this.proxyContract = value;
  }
  @action
  setSubmittedHash(value) {
    this.submittedHash = value;
  }
  @action
  async init() {
    const web3 = new Web3(new Web3.providers.WebsocketProvider('wss://ropsten.infura.io/ws'));
    this.web3 = web3;
    window.web3 = web3;
    const { abi } = contractRaw;
    const token = '0x479c13c7614560648836bd1a2b3c6cc3ea3edd00';
    const tokenContract = new web3.eth.Contract(abi, token);
    const sender = '0x446C4201924ec3C9CAc04c0f18bEA09D752255C3';
    window.tokenContract = tokenContract;
    const eventList = await tokenContract.getPastEvents('TransferPreSigned', { fromBlock: 0, toBlock: 'latest' });
    // const filteredEvents = eventList.filter(ev => ev.event === 'TransferPreSigned');
    console.log(eventList);
    this.setHistoryEvents(eventList);
    const balance = await tokenContract.methods.balanceOf(sender).call();
    console.log(balance);
    const bn = new web3.utils.BN(balance);
    const inEth = web3.utils.fromWei(bn);
    this.setBalance(inEth);
  }
  @action
  setResponse(value) {
    this.response = value;
  }
  @action
  setStatus(value) {
    this.status = value;
  }

  @action
  async createProxy() {
    const ETH = '0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee';
    const DAI = '0xbf76eCA1fbDE4fcec2D4419549dEb560C53A3071';
    const web3Obj = this.web3;
    web3Obj.eth.accounts.wallet.add('0x19180DD02A43EEBC7C071507F53502E7192FD2E228C929854E2417DDF63ADF6C');
    const proxyAbi = [{"constant":false,"inputs":[{"name":"alerter","type":"address"}],"name":"removeAlerter","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"enabled","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"pendingAdmin","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"getOperators","outputs":[{"name":"","type":"address[]"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"src","type":"address"},{"name":"srcAmount","type":"uint256"},{"name":"dest","type":"address"},{"name":"destAddress","type":"address"},{"name":"maxDestAmount","type":"uint256"},{"name":"minConversionRate","type":"uint256"},{"name":"walletId","type":"address"},{"name":"hint","type":"bytes"}],"name":"tradeWithHint","outputs":[{"name":"","type":"uint256"}],"payable":true,"stateMutability":"payable","type":"function"},{"constant":false,"inputs":[{"name":"token","type":"address"},{"name":"srcAmount","type":"uint256"},{"name":"minConversionRate","type":"uint256"}],"name":"swapTokenToEther","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"token","type":"address"},{"name":"amount","type":"uint256"},{"name":"sendTo","type":"address"}],"name":"withdrawToken","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"maxGasPrice","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"newAlerter","type":"address"}],"name":"addAlerter","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"kyberNetworkContract","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"user","type":"address"}],"name":"getUserCapInWei","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"src","type":"address"},{"name":"srcAmount","type":"uint256"},{"name":"dest","type":"address"},{"name":"minConversionRate","type":"uint256"}],"name":"swapTokenToToken","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"newAdmin","type":"address"}],"name":"transferAdmin","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[],"name":"claimAdmin","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"token","type":"address"},{"name":"minConversionRate","type":"uint256"}],"name":"swapEtherToToken","outputs":[{"name":"","type":"uint256"}],"payable":true,"stateMutability":"payable","type":"function"},{"constant":false,"inputs":[{"name":"newAdmin","type":"address"}],"name":"transferAdminQuickly","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"getAlerters","outputs":[{"name":"","type":"address[]"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"src","type":"address"},{"name":"dest","type":"address"},{"name":"srcQty","type":"uint256"}],"name":"getExpectedRate","outputs":[{"name":"expectedRate","type":"uint256"},{"name":"slippageRate","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"user","type":"address"},{"name":"token","type":"address"}],"name":"getUserCapInTokenWei","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"newOperator","type":"address"}],"name":"addOperator","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"_kyberNetworkContract","type":"address"}],"name":"setKyberNetworkContract","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"operator","type":"address"}],"name":"removeOperator","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"field","type":"bytes32"}],"name":"info","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"src","type":"address"},{"name":"srcAmount","type":"uint256"},{"name":"dest","type":"address"},{"name":"destAddress","type":"address"},{"name":"maxDestAmount","type":"uint256"},{"name":"minConversionRate","type":"uint256"},{"name":"walletId","type":"address"}],"name":"trade","outputs":[{"name":"","type":"uint256"}],"payable":true,"stateMutability":"payable","type":"function"},{"constant":false,"inputs":[{"name":"amount","type":"uint256"},{"name":"sendTo","type":"address"}],"name":"withdrawEther","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"token","type":"address"},{"name":"user","type":"address"}],"name":"getBalance","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"admin","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"inputs":[{"name":"_admin","type":"address"}],"payable":false,"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"name":"trader","type":"address"},{"indexed":false,"name":"src","type":"address"},{"indexed":false,"name":"dest","type":"address"},{"indexed":false,"name":"actualSrcAmount","type":"uint256"},{"indexed":false,"name":"actualDestAmount","type":"uint256"}],"name":"ExecuteTrade","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"newNetworkContract","type":"address"},{"indexed":false,"name":"oldNetworkContract","type":"address"}],"name":"KyberNetworkSet","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"token","type":"address"},{"indexed":false,"name":"amount","type":"uint256"},{"indexed":false,"name":"sendTo","type":"address"}],"name":"TokenWithdraw","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"amount","type":"uint256"},{"indexed":false,"name":"sendTo","type":"address"}],"name":"EtherWithdraw","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"pendingAdmin","type":"address"}],"name":"TransferAdminPending","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"newAdmin","type":"address"},{"indexed":false,"name":"previousAdmin","type":"address"}],"name":"AdminClaimed","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"newAlerter","type":"address"},{"indexed":false,"name":"isAdd","type":"bool"}],"name":"AlerterAdded","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"newOperator","type":"address"},{"indexed":false,"name":"isAdd","type":"bool"}],"name":"OperatorAdded","type":"event"}];
    const contract = new web3Obj.eth.Contract(proxyAbi, '0x818E6FECD516Ecc3849DAf6845e3EC868087B755');
    window.proxyContract = contract;
    const rate = await contract.methods.getExpectedRate(ETH, DAI, '10').call();
    console.log(rate);
    // const transactionData = contract.methods.trade(
    //   '0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee', //ERC20 srcToken
    //   '1000000000000000000', //uint srcAmount 1eth
    //   '0xaD6D458402F60fD3Bd25163575031ACDce07538D', //ERC20 destToken
    //   '0xbf76eCA1fbDE4fcec2D4419549dEb560C53A3071', //address destAddress
    //   maximumDestTokenWeiAmount, //uint maxDestAmount
    //   minConversionWeiRate, //uint minConversionRate
    //   0 //uint walletId
    // ).encodeABI();

    // const txReceipt = await web3.eth.sendTransaction({
    //     from: USER_WALLET_ADDRESS, //obtained from your wallet application
    //     to: KYBER_NETWORK_ADDRESS,
    //     data: transactionData,
    //     value: ethTokenWeiAmount, //ADDITIONAL FIELD HERE
    // });

    console.log(txReceipt);
  }
  @action
  async send(amount, rc, speed) {
    this.receiver = rc;
    this.amount = amount;
    this.speed = speed;
    this.setStatus('0');
    const { abi } = contractRaw;
    const { web3 } = this;// no need provider
    const privateKey = '0xb9c6f0ec1a364cb85fea874fd0d75847c987e277dc9c8553d9bff453eae5ed2e';
    // private key
    const sender = '0x446C4201924ec3C9CAc04c0f18bEA09D752255C3';
    // const receiver = '0x9bd961ef9e0a0a4cd9234608d2b1d10ba7a1352b';
    const token = '0x479c13c7614560648836bd1a2b3c6cc3ea3edd00';
    const tokenContract = new web3.eth.Contract(abi, token);
    console.log(tokenContract);
    const sendAmount = toTokenAmount(amount);
    const receiver = rc;
    const feeAmount = toTokenAmount(1);
    const nonce = `${Math.floor(Math.random() * 1000)}`;

    const signature = '';
    const payload = {
      sender,
      receiver,
      token,
      feeAmount,
      sendAmount,
      nonce,
      signature,
    };
    // send transaction
    const response = await axios.post('https://08bab1d2.ngrok.io/api/hash', payload);
    const hashFromResponse = response.data.data;
    const sig = web3.eth.accounts.sign(hashFromResponse, privateKey).signature;
    const payload2 = {
      sender,
      receiver,
      token,
      feeAmount,
      sendAmount,
      nonce,
      signature: sig,
    };
    const response2 = await axios.post('https://08bab1d2.ngrok.io/api/transaction', payload2);
    console.log(response2.data);
    const { txHash } = response2.data.data;
    console.log('txHash', txHash);
    // const response3 = await axios.get(`http://172.17.9.62:5100/api/transactions/${txHash}`);
    // const { submittedHash } = response3.data.data;
    // console.log('submittedHash', response3.data.data);
    const interval = setInterval(async () => {
      const getResponse = await axios.get(`https://08bab1d2.ngrok.io/api/transactions/${txHash}`);
      console.log('get', getResponse.data.data.status);
      if (getResponse.data.data.status === 'SUCCESS') {
        console.log('succuss@!!', getResponse.data.data);
        this.setSubmittedHash(getResponse.data.data.submittedHash);
        this.setStatus('200');
        clearInterval(interval);
      } else if (getResponse.data.data.status === 'FAILED') {
        this.setStatus('400');
        clearInterval(interval);
      }
    }, 3000);
    window.interval = interval;
  }
}

export default NetworkStore;
