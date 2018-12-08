import { action, configure, observable } from 'mobx';

configure({ enforceActions: 'always' }); // don't allow state modifications outside actions

class UiStore {
  constructor(rootStore) {
    this.rootStore = rootStore;
  }
  @observable drawerIsOpen = false;

  @action
  setDrawerIsOpen(value) {
    this.drawerIsOpen = value;
  }
  @action
  toggleDrawer() {
    this.drawerIsOpen = !this.drawerIsOpen;
  }
}

export default UiStore;
