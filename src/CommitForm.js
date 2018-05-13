// @flow
import React, { Component } from 'react'
import { Route, Link } from "react-router-dom";
import { Segment, Button, Checkbox, Header, Form } from 'semantic-ui-react'
import styled from 'styled-components';

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
    guess: null
  }

  commit = async () => {
    const account = this.props.accounts[0]

    const game = this.props.GameInstance;
    const bet = await game.BET_SIZE();
    const hash = this.props.web3.utils.soliditySha3({type: 'string', value: this.state.guess}, {type: 'string', value: "3"});

    await game.commit(hash, { value: bet, from: account });

    const curr_number_bets = await game.curr_number_bets();
    console.log(parseInt(curr_number_bets))
    const guess_commit = await game.commits(account);
    console.log(guess_commit)

    // assert.equal(curr_number_bets, 1, "Number of bets did not increment");
    // assert.equal(guess_commit, hash, "Hashes do not match");

    // this.props.history.push('/ games/two-thirds/reveal')
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
          <Button color="purple" onClick={this.commit}type='submit'>Submit</Button>
        </Form>      
      </Wrapper>
    )
  }
}

export default CommitForm;