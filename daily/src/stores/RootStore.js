import { configure, observable } from 'mobx';
import NetworkStore from './NetworkStore';
import UiStore from './UiStore';

configure({ enforceActions: 'always' }); // don't allow state modifications outside actions

class RootStore {
  constructor() {
    this.networkStore.init();
  }
  @observable networkStore = new NetworkStore(this);
  @observable uiStore = new UiStore(this);
}

const SingletonRootStore = new RootStore();

export default SingletonRootStore;
