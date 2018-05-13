// @flow
import React, { Component } from 'react'
import { Route, Link } from "react-router-dom";
import { Button, Header, Checkbox, Form } from 'semantic-ui-react'
var Web3Utils = require('web3-utils');

type props = {};

class RevealForm extends Component<props> {

  state = {
    loading: false
  }

  reveal = async () => { 

    const account = this.props.accounts[0]

    const game = this.props.GameInstance;
    const bet = await game.BET_SIZE();
    const hash = Web3Utils.soliditySha3({type: 'string', value: this.props.guess}, {type: 'string', value: "3"});

    let curr_number_bets = await game.curr_number_bets();
    console.log(parseInt(curr_number_bets))
    let state = await game.game_state_debug();
    console.log(parseInt(state))

    this.setState({loading: true})
    try {
      await game.reveal(this.props.guess, "3", {from: account, gasPrice: 80000000000 });
    } catch(e) {
      console.log(e)
    }
    this.setState({loading: false})

    curr_number_bets = await game.curr_number_bets();
    console.log(parseInt(curr_number_bets))
    state = await game.game_state_debug();
    console.log(parseInt(state))

    this.props.history.push('/games/two-thirds/payout') 
  }

  render() {
    return (
      <div>
        <Header as='h2'>Reveal your guess!</Header>      
        <Button loading={this.state.loading} color="blue" onClick={this.reveal} type='submit'>Reveal</Button>
      </div>
    )
  }
}

export default RevealForm;