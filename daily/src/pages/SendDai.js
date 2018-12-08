import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import { withRouter } from 'react-router';
import Card from '@material-ui/core/Card';
import Slider from '@material-ui/lab/Slider';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import InputAdornment from '@material-ui/core/InputAdornment';
import { inject, observer } from 'mobx-react';

const styles = theme => ({
  root: {
    display: 'flex',
    flexDirection: 'column',
    // justifyContent: 'center',
    alignItems: 'center',
    paddingTop: '1.5em',
  },
  verticleSpacer: {
    marginTop: '0.5em',
  },
  top: {
    alignSelf: 'flex-start',
    paddingLeft: '1em',
    display: 'flex',
    flexDirection: 'column',
  },
  title: {
    fontSize: '2em',
    fontWeight: 500,
    color: '#6E6E6E',
  },
  subTitle: {
    fontSize: '1em',
    fontWeight: 500,
    color: '#999',
    marginBottom: '2em',
  },
  container: {
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    margin: '0.5em 0',
  },
  cardRoot: {
    width: 380,
    minWidth: 380,
    maxWidth: 380,
    overflow: 'hidden',
  },
  card: {
    padding: '1em',
  },
  cardTitle: {
    fontSize: '1.5em',
    fontWeight: 500,
    color: '#555',
  },
  address: {
    marginTop: '0.5em',
    fontWeight: 500,
    fontSize: '0.9em',
    color: '#888',
  },
  textField: {
    flexBasis: 200,
    minWidth: 350,
  },
  slider: {
    padding: '22px 0px',
    margin: '0 20px',
    touchAction: 'none',
  },
});

@inject('RootStore') @observer
class SendDai extends Component {
  state = {
    amount: '',
    address: this.props.RootStore.recipient,
    speed: 0,
  }
  componentDidMount() {
    console.log(this.props.match);
  }
  handleChange(e) {
    this.setState({
      amount: e.target.value,
    });
  }
  handleAddressChange(e) {
    this.setState({
      address: e.target.value,
    });
  }
  handleSliderChange = (event, value) => {
    this.setState({ speed: value });
  };
  render() {
    const { classes } = this.props;
    const verticleSpacer = () => {
      return <div className={classes.verticleSpacer} />;
    };
    return (
      <div className={classes.root}>
        <div className={classes.top}>
          <span className={classes.title}>Send DAI </span>
          <span className={classes.subTitle}>Send to any wallet and pay gas in Dai too</span>

          <div className={classes.container}>
            <Card className={classes.cardRoot}>
              <div className={classes.card}>
                <div className={classes.cardTitle}>
                  From:
                </div>
                <div className={classes.address}>
                  0xE4Bfd8b40e78e539eb59719Ad695D0D0132FA502
                </div>
                {verticleSpacer()}
                <div className={classes.cardTitle}>
                  Balance:
                </div>
                <div className={classes.address}>
                  $1000.00 Dai
                </div>

              </div>
            </Card>
          </div>

          <div className={classes.container}>
            <Card className={classes.cardRoot}>
              <div className={classes.card}>
                <div className={classes.cardTitle}>
                  To:
                </div>
                <div className={classes.address}>
                  <TextField
                    id="standard-name"
                    label="Address"
                    InputProps={{
                      className: classes.address,
                    }}
                    className={classes.textField}
                    value={this.state.address}
                    onChange={this.handleAddressChange.bind(this)}
                    margin="normal"
                  />
                </div>
                {verticleSpacer()}
                <div className={classes.cardTitle}>
                  Amount:
                </div>
                {/* <div className={classes.address}> */}
                <TextField
                  id="standard-name"
                  label="Amount"
                  className={classes.textField}
                  InputProps={{
                    className: classes.address,
                    endAdornment: <InputAdornment position="end">Dai</InputAdornment>,
                  }}
                  value={this.state.amount}
                  onChange={this.handleChange.bind(this)}
                  margin="normal"
                />
                {/* </div> */}
              </div>
            </Card>
          </div>

          <div className={classes.container}>
            <Card className={classes.cardRoot}>
              <div className={classes.card}>
                <div className={classes.cardTitle}>
                  Options
                </div>
                {verticleSpacer()}
                <div className={classes.address}>
                  Transaction Speed:
                  {this.state.speed === 0 && ' Slow'}
                  {this.state.speed === 1 && ' Fast'}
                  {this.state.speed === 2 && ' Super Fast'}
                </div>
                <Slider
                  classes={{ container: classes.slider }}
                  value={this.state.speed}
                  min={0}
                  max={2}
                  step={1}
                  onChange={this.handleSliderChange}
                />
              </div>
            </Card>
            {verticleSpacer()}
            {verticleSpacer()}
            {verticleSpacer()}
            <Button
              color="primary"
              variant="contained"
              onClick={() => {
                this.props.RootStore.networkStore.send(
                  this.state.amount,
                  this.state.address,
                  this.state.speed,
                );
                this.props.history.push('/progress');
              }}
            >
              Send
            </Button>

          </div>
        </div>
      </div>
    );
  }
}
SendDai.wrappedComponent.propTypes = {
  RootStore: PropTypes.object.isRequired,
};
SendDai.propTypes = {
  classes: PropTypes.object.isRequired,
  match: PropTypes.object.isRequired,
  history: PropTypes.object.isRequired,
};
export default withRouter(withStyles(styles)(SendDai));
