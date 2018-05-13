// @flow
import React, { Component } from 'react'
import { Route, Link } from "react-router-dom";
import SelectPool from './SelectPool'
import CommitForm from './CommitForm'
import RevealForm from './RevealForm'
import PayoutPage from './PayoutPage'


type props = {};

class Game extends Component<props> {
  state = {
    state: null, // COMMIT, COMMITTED, REVEAL, REVEALED, PAYOUT
    stake: null, // 1, 0.1, 0.001
    guess: null
  }

  setStake = async (stake) => {
    this.setState({ stake, state: "COMMIT" }, function() {
      this.props.history.push('/games/two-thirds/commit')
    })
  }

  render() {
    // const url = this.props.match.url
    return (
      <div>
        <Route exact path="/games/two-thirds/" render={(props) => ( <SelectPool setStake={ (stake) => {
            this.setState({ stake, state: "COMMIT" }, function() {
              props.history.push('/games/two-thirds/commit')
            })
          }
        } {...props} {...this.state} /> )} />
        <Route exact path="/games/two-thirds/commit" render={(props) => ( <CommitForm {...this.props} {...this.state} /> )} />
        <Route exact path="/games/two-thirds/reveal" render={(props) => ( <RevealForm {...this.props} {...this.state} /> )} />
        <Route exact path="/games/two-thirds/payout" render={(props) => ( <PayoutPage {...this.props} {...this.state} /> )} />
      </div>
    )
  }
}

export default Game;