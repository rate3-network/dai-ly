import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import { withRouter } from 'react-router';

const styles = theme => ({
  root: {
    display: 'flex',
    flexDirection: 'column',
    // justifyContent: 'center',
    alignItems: 'center',
  },
  scanner: {
    maxWidth: 500,
  },
});

class SendDai extends Component {
  componentDidMount() {
    console.log(this.props.match);
  }
  render() {
    const { classes } = this.props;
    return (
      <div className={classes.root}>
        send to {this.props.match.params.address}
      </div>
    );
  }
}
SendDai.propTypes = {
  classes: PropTypes.object.isRequired,
  match: PropTypes.object.isRequired,
};
export default withRouter(withStyles(styles)(SendDai));
