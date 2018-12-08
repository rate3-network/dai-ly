import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import { inject, observer } from 'mobx-react';


const styles = (theme) => {
  return ({

  });
};
const Main = inject('RootStore')(observer((props) => {
  const { classes } = props;
  console.log('test');
  return (
    <div>
      test
      {props.RootStore.networkStore.hello}
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
