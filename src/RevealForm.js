// @flow
import React, { Component } from 'react'
import { Route, Link } from "react-router-dom";
import { Button, Header, Checkbox, Form } from 'semantic-ui-react'
import styled from 'styled-components';
import API from './api/Game.js'

var Web3Utils = require('web3-utils');

type props = {};

class RevealForm extends Component<props> {

  state = {
    loading: false
  }

  loading() {
    const { web3, accounts, game } = this.props;
    return !(web3 && accounts && accounts.length > 0 && game)
  }

  reveal = async () => { 
    const web3 = this.props.web3;
    const account = this.props.accounts[0];
    const game = this.props.game;
    
    const api = new API(web3.utils, () => {}, game);

    this.setState({loading: true})
    try {
      await api.revealGuess(account, this.props.guess, this.props.hashKey);
    } catch(e) {
      console.log(e)
    }
    this.setState({loading: false})
    this.props.setParentState({ state: "REVEALED" })
    this.props.history.push('/games/two-thirds/revealed') 
  }

  render() {
    return (
      <div>
        <Header as='h2'>Reveal your guess!</Header>      
        <Button disabled={this.loading()} loading={this.state.loading} color="blue" onClick={this.reveal} type='submit'>Reveal</Button>
      </div>
    )
  }
}

export default RevealForm;