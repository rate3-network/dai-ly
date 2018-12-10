import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';
import { inject, observer } from 'mobx-react';
import { Link } from 'react-router-dom';
import Dai from '../assets/Dai.png';
import HistoryRow from '../components/HistoryRow';

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
    pending: {
      fontSize: '0.8em',
      fontWeight: 500,
      color: '#9A9A9A',
      margin: '0.2em 0',
    },
    btnGroup: {
      margin: '1em 0',
      display: 'flex',
      // flexDirection: 'column',
      alignItems: 'center',
    },
    spacer: {
      margin: '0 0.5em',
    },
    btn: {
      boxShadow: 'none',
      width: '8em',
    },
    noDeco: {
      textDecoration: 'none',
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
const Main = inject('RootStore')(observer((props) => {
  const { classes } = props;
  return (
    <div className={classes.root}>
      <img src={Dai} alt="Dai Icon" />
      <span className={classes.balance}>
        {props.RootStore.networkStore.balance} DAI
      </span>
      <span className={classes.pending}>
        ___ DAI Pending
      </span>
      <div className={classes.btnGroup}>
        <Link className={classes.noDeco} to="/send">
          <Button className={classes.btn} size="small" variant="contained" color="primary">
            Send
          </Button>
        </Link>
        <span className={classes.spacer} />
        <Link className={classes.noDeco} to="/scan">
          <Button className={classes.btn} size="small" variant="contained" color="primary">
            Scan
          </Button>
        </Link>
      </div>
      <span className={classes.tttt}>History</span>
      <div className={classes.histSection}>
        {props.RootStore.networkStore.historyEvents.map((item) => {
          return (
            <HistoryRow key={item.transactionHash} ev={item} />
          );
        })}
      </div>
    </div>
  );
}));

Main.propTypes = {
  classes: PropTypes.object.isRequired,
};
Main.wrappedComponent.propTypes = {
  RootStore: PropTypes.object.isRequired,
};
export default withStyles(styles)(Main);
