import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';
import { inject, observer } from 'mobx-react';

import HistoryRow from '../components/HistoryRow';

import convert from '../assets/convert.png';

const styles = (theme) => {
  return ({
    root: {
      paddingTop: '3em',

      display: 'flex',
      flexDirection: 'column',
      alignItems: 'center',
    },
    balance: {
      fontSize: '1.2em',
      fontWeight: 500,
      color: '#444',
      marginTop: '0.8em',
    },
    topRow: {
      width: '100%',
      display: 'flex',
      padding: '1em 0',
      justifyContent: 'space-around',
    },
    balanceTitle: {
      fontSize: '1.1em',
      fontWeight: 500,
      color: '#444',
    },
    tttt: {
      alignSelf: 'flex-start',
      padding: '1em',
      color: '#555',
    },
    histSection: {
      display: 'flex',
      flexDirection: 'column',
      minWidth: '100vw',
      minHeight: '100vh',
      backgroundColor: 'white',
    },
  });
};
@inject('RootStore') @observer
class Delegate extends Component {
  componentDidMount() {
    this.props.RootStore.setRole(2);

    this.props.RootStore.networkStore.createProxy();
  }
  componentWillUnmount() {
    this.props.RootStore.setRole(1);
  }
  render() {
    const { classes } = this.props;
    return (
      <div className={classes.root}>
        <img width="80%" src={convert} alt="convert" />
        <div className={classes.topRow}>
          <div className={classes.balanceTitle}>
            1000.00 DAI
          </div>
          <div className={classes.balanceTitle}>
            11.03 ETH
          </div>
        </div>
        {/* <div className={}/> */}
        {/* 11.03 ETH */}
        <Button variant="contained" color="secondary">Swap</Button>
        <span className={classes.tttt}>History</span>
        <div className={classes.histSection}>
          {this.props.RootStore.networkStore.historyEvents.map((item) => {
            return (
              <HistoryRow key={item.transactionHash} ev={item} />
            );
          })}
        </div>
      </div>
    );
  }
}

Delegate.propTypes = {
  classes: PropTypes.object.isRequired,
};
Delegate.wrappedComponent.propTypes = {
  RootStore: PropTypes.object.isRequired,
};
export default withStyles(styles)(Delegate);
