import { configure, observable, action } from 'mobx';
import Web3 from 'web3';

configure({ enforceActions: 'always' }); // don't allow state modifications outside actions

class NetworkStore {
  constructor(rootStore) {
    this.rootStore = rootStore;
  }
  @observable hello = 'gello';
  @observable web3 = null;

  @action
  init() {
    const web3 = new Web3(new Web3.providers.WebsocketProvider('wss://ropsten.infura.io/ws'));
    this.web3 = web3;
    window.web3 = web3;
  }
}

export default NetworkStore;
