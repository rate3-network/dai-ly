import React, { Component } from 'react';
import Lottie from 'react-lottie';

import submit from '../assets/submit.json';
// https://www.lottiefiles.com/627-loading-success-failed

const noLoopOptions = {
  animationData: submit,
  loop: false,
  // autoPlay: true,
};
class SuccessAnim extends Component {
  state = {
    startFrame: 240,
    endFrame: 400,
    speed: 1.5,
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
        options={noLoopOptions}
        height={200}
        speed={this.state.speed}
        segments={[this.state.startFrame, this.state.endFrame]}
        isClickToPauseDisabled
      />
    );
  }
}


export default SuccessAnim;
