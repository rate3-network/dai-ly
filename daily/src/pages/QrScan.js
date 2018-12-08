import React, { Component } from 'react';
import QrReader from 'react-qr-reader';
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

class QrScan extends Component {
  constructor(props) {
    super(props);
    this.state = {
      delay: 500,
      result: 'No result',
    };
    this.handleScan = this.handleScan.bind(this);
  }
  handleScan(data) {
    console.log(data);
    if (data) {
      this.setState({
        result: data,
      });
      setTimeout(() => {
        window.navigator.vibrate(300);
        this.props.history.push(`./send/${data}`);
      }, 1000);
    }
  }
  handleError(err) {
    console.error(err);
  }
  render() {
    const { classes } = this.props;
    return (
      <div className={classes.root}>
        <QrReader
          className={classes.scanner}
          facingMode="environment"
          delay={this.state.delay}
          onError={this.handleError}
          onScan={this.handleScan}
          style={{ width: '100%' }}
        />
        <p>{this.state.result}</p>
      </div>
    );
  }
}
QrScan.propTypes = {
  classes: PropTypes.object.isRequired,
  history: PropTypes.object.isRequired,
};
export default withRouter(withStyles(styles)(QrScan));
