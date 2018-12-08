import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import { inject, observer } from 'mobx-react';

import Loading from '../components/Loading';

const styles = (theme) => {
  return ({

  });
};

@inject('RootStore') @observer
class SendProgress extends Component {
  componentDidMount() {

  }
  render() {
    const { classes, RootStore } = this.props;
    return (
      <div>
        <Loading />
        progress
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
