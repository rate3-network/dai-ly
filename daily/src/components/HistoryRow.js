import React from 'react';
import PropTypes from 'prop-types';
import Divider from '@material-ui/core/Divider';
import { withStyles } from '@material-ui/core/styles';

import { fromTokenAmount, shorten } from '../stores/convert';

const styles = (theme) => {
  return ({
    root: {
      width: '100%',
      backgroundColor: 'white',
      color: '#555',
    },
    container: {
      padding: '1em 1em',
      display: 'grid',
      gridTemplateColumns: '3fr 1.5fr',
      gridGap: '8px',
      width: '100%',
      gridTemplateAreas: `
      'a b'
      'c d'
      `,
    },
    a: {
      gridArea: 'a',
    },
    b: {
      gridArea: 'b',
      color: '#777',
      fontSize: '0.8em',
    },
    c: {
      gridArea: 'c',
    },
    d: {
      gridArea: 'd',
      color: '#777',
      fontSize: '0.8em',
    },
  });
};
const HistoryRow = (props) => {
  const { classes, ev } = props;
  return (
    <div className={classes.root}>
      <Divider />
      <div className={classes.container}>
        <span className={classes.a} onClick={() => { window.open(`https://ropsten.etherscan.io/search?q=${ev.transactionHash}`) }}>{shorten(ev.transactionHash)}......</span>
        <span className={classes.b}>${fromTokenAmount(ev.returnValues.amount, 2)} DAI</span>
        <span className={classes.c}>Success</span>
        <span className={classes.d}>${fromTokenAmount(ev.returnValues.fee, 2)} DAI</span>
      </div>

    </div>
  );
};

HistoryRow.propTypes = {
  classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(HistoryRow);
