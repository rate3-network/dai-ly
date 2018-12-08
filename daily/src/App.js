import React, { Component } from 'react';
import { Provider } from 'mobx-react';
import { createMuiTheme, MuiThemeProvider } from '@material-ui/core/styles';
import { HashRouter, Switch, Route } from 'react-router-dom';

import './App.css';
import Main from './pages/Main';
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
    textPrimary: {
      main: 'white',
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
          <div>
            <AppNavBar />
            <AppDrawer />
            <div className="app">
              <HashRouter basename="/">
                <Switch>
                  <Route
                    exact
                    path="/"
                    component={Main}
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
