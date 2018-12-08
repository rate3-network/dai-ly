import React, { PureComponent } from 'react';
import Lottie from 'react-lottie';

import biking from '../assets/biking_is_cool.json';
// animation created by Yue XI
// https://www.lottiefiles.com/1446-bikingiscool
const defaultOptions = {
  animationData: biking,
  loop: true,
};
/* eslint react/prefer-stateless-function: "off" */
class LoadingAnim extends PureComponent {
  render() {
    return (
      <Lottie
        options={defaultOptions}
        height={200}
        isClickToPauseDisabled
      />
    );
  }
}


export default LoadingAnim;
