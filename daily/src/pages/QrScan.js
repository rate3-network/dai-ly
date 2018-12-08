import React, { Component } from 'react';
import QrReader from 'react-qr-reader';
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import { withRouter } from 'react-router';
import { inject, observer } from 'mobx-react';

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

@inject('RootStore') @observer
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
      const cleaned = data.replace(/\//g, '');
      this.setState({
        result: cleaned,
      });
      setTimeout(() => {
        window.navigator.vibrate(600);
        this.props.RootStore.setRecipient(cleaned);
        this.props.history.push(`./send/${cleaned}`);
      }, 800);
    }
  }
  handleError(err) {
    console.log(err);
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
QrScan.wrappedComponent.propTypes = {
  RootStore: PropTypes.object.isRequired,
};
QrScan.propTypes = {
  classes: PropTypes.object.isRequired,
  history: PropTypes.object.isRequired,
};
export default withRouter(withStyles(styles)(QrScan));
