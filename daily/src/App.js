import React, { Component } from 'react';
import { Provider } from 'mobx-react';
import { createMuiTheme, MuiThemeProvider } from '@material-ui/core/styles';
import { HashRouter, Switch, Route } from 'react-router-dom';

import './App.css';
import Main from './pages/Main';
import QrScan from './pages/QrScan';
import SendProgress from './pages/SendProgress';
import SendDai from './pages/SendDai';
import Delegate from './pages/Delegate';
import RootStore from './stores/RootStore';
import AppDrawer from './components/SwipableDrawer';
import AppNavBar from './components/AppNavBar';

const stores = {
  RootStore,
};
const theme = createMuiTheme({
  typography: {
    useNextVariants: true,
  },
  palette: {
    primary: {
      main: '#F2B350',
      contrastText: '#fff',
    },
    secondary: {
      main: '#5EB1BF',
      contrastText: '#fff',
    },
    // secondary: '#5EB1BF',
  },
});

class App extends Component {
  componentWillMount() {

  }
  render() {
    return (
      <Provider {...stores}>
        <MuiThemeProvider theme={theme}>
          <div className="desktop2mobile">
            <AppNavBar />
            <AppDrawer />
            <div className="app">
              <HashRouter basename="/">
                <Switch>
                  <Route
                    exact
                    path="/delegate"
                    component={Delegate}
                  />
                  <Route
                    exact
                    path="/"
                    component={Main}
                  />
                  <Route
                    exact
                    path="/scan"
                    component={QrScan}
                  />
                  <Route
                    // exact
                    path="/send/"
                    component={SendDai}
                  />
                  <Route
                    exact
                    path="/progress"
                    component={SendProgress}
                  />
                </Switch>
              </HashRouter>
            </div>
            {/* <Main /> */}
          </div>
        </MuiThemeProvider>
      </Provider>
    );
  }
}

export default App;
