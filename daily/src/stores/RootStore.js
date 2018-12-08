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

  @action
  setRecipient(value) {
    this.recipient = value;
  }
}

const SingletonRootStore = new RootStore();

export default SingletonRootStore;
