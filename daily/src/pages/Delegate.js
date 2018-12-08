import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import { inject, observer } from 'mobx-react';

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
