// @flow
import React, { Component } from 'react';
import { BrowserRouter as Router, Route, Link } from 'react-router-dom';
import getWeb3 from './utils/getWeb3';
// import './App.css';

import Landing from './Landing'

type props = {};

class App extends Component<props> {
  render() {
    return (
      <Router>
        <Route exact path="/" render={(props) => ( <Landing {...props} {...this.state} /> )} />
      </Router>
    );
  }
}

export default App;
