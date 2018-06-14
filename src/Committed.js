// @flow
import React, { Component } from 'react'
import { Route, Link } from "react-router-dom";
import { Progress, Segment, Button, Checkbox, Header, Form } from 'semantic-ui-react'
import styled from 'styled-components';
import API from './api/Game.js'
import srs from 'secure-random-string'

var Web3Utils = require('web3-utils');

const Wrapper = styled(Segment)`
  max-width: 500px;
  margin: 0 auto !important; 
  margin-top: 10% !important;

  h2.ui.header {
    font-size: 1.7em;
    font-weight: normal;
  }
`;

type props = {};

class Committed extends Component<props> {
  state = { 
    percent: 0,
    interval: null,
    inProgress: false
  }

  componentDidMount = async () => {
    const interval = setInterval(this.poll, 2000);
    this.setState({ interval })
  }

  componentWillUnmount = async () => {
    if(this.state.interval)
      clearInterval(this.state.interval);
  }

  poll = async () => {
    if(this.state.inProgress)
      return false
    console.info('polling')
    if(!(this.props.game && this.props.accounts.length > 0 && this.props.web3)) {
      console.info('loading content')
      return false
    }
    const web3 = this.props.web3;
    const account = this.props.accounts[0];
    const game = this.props.game;
    const api = new API(web3.utils, () => {}, game);
    // poll contract for # of players 
    let current, total, percent
    this.setState({ inProgress: true })
    try {
      current = await api.getCurrentCommits()
      total = await api.getMaxPlayers()
      percent = (current / total) * 100;
    } catch(e) {
      console.log(e)
      return false
    } finally {
      this.setState({ inProgress: false })
    }
    console.info(`current: ${current}, total: ${total}, percent: ${percent}`)
    this.setState({ percent })
    if(current == total) { // if complete
      setTimeout(()=>{
        clearInterval(this.state.interval);
        console.log('done')
        this.setState({interval: null})
        this.props.setParentState({ state: "REVEAL" })
        this.props.history.push('/games/two-thirds/reveal')
      }, 2);
    }
  }

  render() {
    return (
      <Wrapper>
        <Form>
          <Header as='h2'>Waiting for other to guess.</Header>     
          <Progress percent={this.state.percent} indicating progress />
        </Form>      
      </Wrapper>
    )
  }
}

export default Committed;