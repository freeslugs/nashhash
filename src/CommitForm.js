// @flow
import React, { Component } from 'react'
import { Route, Link } from "react-router-dom";
import { Segment, Button, Checkbox, Header, Form } from 'semantic-ui-react'
import styled from 'styled-components';
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
    local_guess: null,
    loading: false
  }

  commit = async () => {
    const account = this.props.accounts[0]
    // console.log(account)
    const game = this.props.GameInstance;
    const bet = await game.BET_SIZE();
    const hash = Web3Utils.soliditySha3({type: 'string', value: this.state.local_guess}, {type: 'string', value: "3"});

    this.props.setGuess(this.state.local_guess)

    // const curr_number_bets = await game.curr_number_bets();
    // console.log(parseInt(curr_number_bets))
    // await game.set_MAX_PLAYERS(1, { from: account, gasPrice: 80000000000 });
    // await game.commit(hash, { value: bet, from: donor });
    // await game.reveal("66", "3", {from: donor});

    this.setState({loading: true})
    try {
      await game.commit(hash, { value: bet, from: account, gasPrice: 80000000000 });
    } catch(e) {
      console.log(e)
    }
    this.setState({loading: false})


    // assert.equal(curr_number_bets, 1, "Number of bets did not increment");
    // assert.equal(guess_commit, hash, "Hashes do not match");

    this.props.history.push('/games/two-thirds/reveal')
  }

  render() {
    return (
      <Wrapper>
        <Form>
          <Header as='h2'>Make a guess!</Header>     
          <Form.Field>
            <label>Guess</label>
            <input placeholder='5' type="number" onChange={(e) => this.setState({local_guess: e.target.value})} />
          </Form.Field>
          <Button loading={this.state.loading} color="purple" onClick={this.commit} type='submit'>Submit</Button>
        </Form>      
      </Wrapper>
    )
  }
}

export default CommitForm;