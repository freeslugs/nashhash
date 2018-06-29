// @flow
import React, { Component } from 'react'
import { Route, Link } from "react-router-dom";
import { Progress, Segment, Button, Checkbox, Header, Form } from 'semantic-ui-react'
import styled from 'styled-components';
import API from './api/Game.js'
import GameRegistryAPI from './api/GameRegistryAPI.js'
import srs from 'secure-random-string'
import './Committed.css'

var Web3Utils = require('web3-utils');

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
    const gametype = this.props.gametype;
    const gameaddresses = this.props.gameaddresses;
    const stake = this.props.stake;

    const registryAPI = new GameRegistryAPI(web3, gameaddresses);
    const game = await registryAPI.configureGame(gametype, stake);
    const gameAPI = new API(web3.utils, () => {}, game);
    // poll contract for # of players 
    let current, total, percent, cur_state
    this.setState({ inProgress: true })
    try {
      //DISCUSS CHANGING THIS WITH KEVIN  
      current = await gameAPI.getCurrentCommits()
      total = await gameAPI.getMaxPlayers()
      cur_state = await gameAPI.getGameState()  
      percent = (current / total) * 100;
    } catch(e) {
      console.log(e)
      return false
    } finally {
      this.setState({ inProgress: false })
    }
    console.info(`current: ${current}, total: ${total}, percent: ${percent}`)
    this.setState({ percent })
    if(cur_state == 1) { // if game has transitioned to payout state
      setTimeout(()=>{
        clearInterval(this.state.interval);
        console.log('done')
        this.setState({interval: null})
        this.props.setParentState({ state: "REVEAL" })
        if(gametype == "TwoThirds"){
          this.props.history.push('/games/two-thirds/reveal')
        }
        else if(gametype == "LowestUnique"){
          this.props.history.push('/games/lowest-unique/reveal')
        }
      }, 2);
    }
  }

  render() {
    return (
      <div className="committed-wrapper">
          <div className="commit-animation"></div>
          <h2>Your guess is in!</h2>
          <p>Now we wait a little for everyone else’s bets to be processed. Once that is complete, you will see the reveal button above, to see how well you did this time around!</p>
      </div>
    )
  }
}

export default Committed;