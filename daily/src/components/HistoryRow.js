import React from 'react';
import PropTypes from 'prop-types';
import Divider from '@material-ui/core/Divider';
import { withStyles } from '@material-ui/core/styles';

const styles = (theme) => {
  return ({
    root: {
      width: '100%',
      backgroundColor: 'white',
      color: '#555',
    },
  });
};
const HistoryRow = (props) => {
  const { classes } = props;
  return (
    <div className={classes.root}>
      <Divider />
      test
    </div>
  );
};

HistoryRow.propTypes = {
  classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(HistoryRow);
