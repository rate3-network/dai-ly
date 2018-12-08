import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import Fade from '@material-ui/core/Fade';
import Button from '@material-ui/core/Button';
import { inject, observer } from 'mobx-react';
import { Link } from 'react-router-dom';

import Loading from '../components/Loading';
import Success from '../components/Success';
import Failure from '../components/Failure';

const parseSpeed = (speed) => {
  switch (`${speed}`) {
    case '0':
      return 'normal';
    case '1':
      return 'fast';
    case '2':
      return 'super fast';
    default:
      return 'normal';
  }
};
const styles = (theme) => {
  return ({
    verticleSpacer: {
      marginTop: '0.5em',
    },
    col: {
      display: 'flex',
      flexDirection: 'column',
      justifyContent: 'center',
      alignItems: 'center',
    },
    status: {
      display: 'flex',
      flexDirection: 'column',
      justifyContent: 'center',
      alignItems: 'center',
      color: '#666',
      fontWeight: 500,
      fontSize: '1.2em',
      textAlign: 'center',
    },
    address: {
      fontWeight: 500,
      fontSize: '0.8em',
      color: '#888',
    },
    hash: {
      fontWeight: 500,
      fontSize: '0.55em',
      color: '#888',
    },
  });
};

@inject('RootStore') @observer
class SendProgress extends Component {
  componentDidMount() {
    console.log(`
      animation created by Yue XI
      https://www.lottiefiles.com/1446-bikingiscool
      `);
  }
  parseStatus = (s) => {
    const status = `${s}`;
    if (status.startsWith('2')) {
      return 2;
    } else if (status.startsWith('4') || status.startsWith('5')) {
      return 3;
    }
    return 1;
  }
  render() {
    const { classes, RootStore } = this.props;
    const { networkStore } = RootStore;
    const { status } = networkStore;
    const verticleSpacer = () => {
      return <div className={classes.verticleSpacer} />;
    };
    const statusParsed = this.parseStatus(status);
    return (
      <div>
        <div>
          <Fade in={statusParsed === 1} timeout={1000}>
            <div className={classes.col}>
              {statusParsed === 1 && <Loading />}
            </div>
          </Fade>
          <Fade in={statusParsed === 2} timeout={1000}>
            <div className={classes.col}>
              {verticleSpacer()}
              {verticleSpacer()}
              {verticleSpacer()}
              {statusParsed === 2 && <Success />}
            </div>
          </Fade>
          <Fade in={statusParsed === 3} timeout={1000}>
            <div className={classes.col}>
              {verticleSpacer()}
              {verticleSpacer()}
              {verticleSpacer()}
              {statusParsed === 3 && <Failure />}
            </div>
          </Fade>
        </div>
        {statusParsed === 1 &&
          <div className={classes.status}>
            Sending ${networkStore.amount} DAI to
            <span className={classes.address}>{networkStore.receiver}</span>
             with {parseSpeed(networkStore.speed)} speed
          </div>
        }
        {statusParsed === 2 &&
          <div className={classes.status}>
            Successfully sent ${networkStore.amount} DAI to
            <span className={classes.address}>{networkStore.receiver}</span>
             with {parseSpeed(networkStore.speed)} speed
            {verticleSpacer()}
            {verticleSpacer()}
            Transaction Hash
            <span className={classes.hash}>{networkStore.submittedHash}</span>
            {verticleSpacer()}
            {verticleSpacer()}
            <Link to="/send" style={{ textDecoration: 'none' }}>
              <Button variant="contained" color="primary">Send Another One</Button>
            </Link>
          </div>
        }
        {statusParsed === 3 &&
          <div className={classes.status}>
            Failed to send ${networkStore.amount}.00 to
            <span className={classes.address}>{networkStore.receiver}</span>
             with {parseSpeed(networkStore.speed)} speed
            {verticleSpacer()}
            {verticleSpacer()}
            <Link to="/send" style={{ textDecoration: 'none' }}>
              <Button variant="contained" color="primary">Try Again</Button>
            </Link>
          </div>
        }
      </div>
    );
  }
}
SendProgress.wrappedComponent.propTypes = {
  RootStore: PropTypes.object.isRequired,
};
SendProgress.propTypes = {
  classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(SendProgress);
