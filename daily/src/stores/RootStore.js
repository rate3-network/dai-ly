import { action, configure, observable } from 'mobx';
import NetworkStore from './NetworkStore';
import UiStore from './UiStore';

configure({ enforceActions: 'always' }); // don't allow state modifications outside actions

class RootStore {
  constructor() {
    this.networkStore.init();
  }
  @observable networkStore = new NetworkStore(this);
  @observable uiStore = new UiStore(this);

  @observable recipient = '';
  @observable role = 1;
  @observable historyObj = {};
  @action
  setRecipient(value) {
    this.recipient = value;
  }
  @action
  setHistoryObj(value) {
    this.historyObj = value;
  }
  @action
  setRole(value) {
    this.role = value;
  }
}

const SingletonRootStore = new RootStore();

export default SingletonRootStore;
