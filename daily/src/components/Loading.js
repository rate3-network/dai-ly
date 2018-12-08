import React, { Component } from 'react';
import Lottie from 'react-lottie';

import biking from '../assets/biking_is_cool.json';
// animation created by Yue XI
// https://www.lottiefiles.com/1446-bikingiscool
const defaultOptions = {
  animationData: biking,
  loop: true,
};

class LoadingAnim extends Component {
  state = {
    // startFrame: 0,
    // endFrame: 120,
    speed: 1,
  }
  componentDidMount() {
    this.setFrames();
  }

  setFrames() {
    this.setState({ endFrame: this.state.endFrame, speed: this.state.speed });
  }

  render() {
    return (
      <Lottie
        options={defaultOptions}
        height={80}
        width={80}
        speed={this.state.speed}
        segments={[this.state.startFrame, this.state.endFrame]}
        isClickToPauseDisabled
      />
    );
  }
}


export default LoadingAnim;
