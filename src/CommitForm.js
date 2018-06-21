// @flow
import React, { Component } from 'react'
import { Route, Link } from "react-router-dom";
import { Segment, Button, Checkbox, Header, Form } from 'semantic-ui-react'
import styled from 'styled-components';
import API from './api/Game.js'
import srs from 'secure-random-string'

import GameABI from './contracts/Game.json'

const contract = require('truffle-contract');
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

class CommitForm extends Component<props> {
  state = {
    guess: null,
    loading: false
  }

  componentDidMount() {
    this.props.setParentState({ state: "COMMIT" })
  }

  loading() {
    const { web3, accounts, gameregistry, gameaddresses } = this.props;
    return !(web3 && accounts && accounts.length > 0 && gameaddresses.length > 0)
  }

  commit = async () => {
    const web3 = this.props.web3;
    const account = this.props.accounts[0];
    //const game = this.props.game;
    let game

    const gameaddresses = this.props.gameaddresses;
    const gametype = this.props.gametype;
    const stake = this.props.stake;

    if(gametype=="TwoThirds"){
      if(stake==0.01){
        game = contract(GameABI, gameaddresses[2]);
      }
      else if(stake==0.1){
        game = contract(GameABI, gameaddresses[1]);
      }
      else if(stake==1){
        game = contract(GameABI, gameaddresses[0]);
      }
    } else if(gametype=="LowestUnique")
    {
      if(stake==0.01){
        game = contract(GameABI, gameaddresses[5]);
      }
      else if(stake==0.1){
        game = contract(GameABI, gameaddresses[4]);
      }
      else if(stake==1){
        game = contract(GameABI, gameaddresses[3]);
      }      
    }
    
    const api = new API(web3.utils, () => {}, game);

    const hashKey = /*this.props.hashKey || */ srs({length: 50});
    this.props.setParentState({ hashKey })

    const guess = this.state.guess
    if(guess == null || guess.length == 0) {
      console.log('no guess!')
      return false;
    }
    this.setState({loading: true})
    try {
      await api.commitGuess(account, guess, hashKey);
    } catch(e) {
      console.log(`error: ${e}`)
    }
    this.setState({loading: false})
    this.props.setParentState({ guess: guess, state: "COMMITTED" })
    this.props.history.push('/games/two-thirds/committed')
  }

  render() {
    return (
      <Wrapper>
        <Form>
          <Header as='h2'>Make a guess!</Header>     
          <Form.Field>
            <label>Guess</label>
            <input placeholder='5' type="number" onChange={(e) => this.setState({guess: e.target.value})} />
          </Form.Field>
          <Button disabled={this.loading()} loading={this.state.loading} color="purple" onClick={this.commit} type='submit'>Submit</Button>
        </Form>      
      </Wrapper>
    )
  }
}

export default CommitForm;