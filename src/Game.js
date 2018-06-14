// @flow
import React, { Component } from 'react'
import { Route, Link } from "react-router-dom";
import SelectPool from './SelectPool'
import CommitForm from './CommitForm'
import Committed from './Committed'
import RevealForm from './RevealForm'
import Revealed from './Revealed'
import PayoutPage from './PayoutPage'


type props = {};

class Game extends Component<props> {
  state = {
    state: null, // COMMIT, COMMITTED, REVEAL, REVEALED, PAYOUT
    stake: null, // 1, 0.1, 0.001
    guess: null,
    hashKey: null
  }

  componentDidMount() {
    const json = JSON.parse(localStorage.getItem("state"));
    // console.log(`LS state: ${JSON.stringify(json)}`)
    // console.log("mounted")

    const { state, stake, guess, hashKey } = json;
    let newState = {}
    Object.keys(json).forEach((key) => {
      if(!this.state[key]) {
        newState[key] = json[key]
      }
    })
    
    this.setState(newState)
    // this.setState({ state, stake, guess, hashKey })
    // if(state == "COMMIT")
    //   this.props.history.push('/games/two-thirds/commit')
    // else if(state == "COMMITTED")
    //   this.props.history.push('/games/two-thirds/committed')
  }

  componentWillUpdate(nextProps, nextState) {
    // if diff, write to local storage
    if(this.state != nextState) {
      // console.log(`componentWillUpdate ${JSON.stringify(nextState)}`)
      localStorage.setItem("state", JSON.stringify(nextState));
    }
  }

  setParentState = (newState) => {
    this.setState(newState, () => {
      // console.log(`set parent state ${JSON.stringify(this.state)}`)
      localStorage.setItem("state", JSON.stringify(this.state));  
    });
  }

  render() {
    return (
      <div>
        <Route exact path="/games/two-thirds/" render={(props) => ( 
          <SelectPool setParentState={this.setParentState} {...props} {...this.state} /> 
        )} />
        <Route exact path="/games/two-thirds/commit" render={(props) => ( 
          <CommitForm setParentState={this.setParentState} {...this.props} {...this.state} /> 
        )} />
        <Route exact path="/games/two-thirds/committed" render={(props) => ( <Committed setParentState={this.setParentState} {...props} {...this.props} {...this.state} /> )} />
        <Route exact path="/games/two-thirds/reveal" render={(props) => ( <RevealForm {...this.props} {...this.state} /> )} />
        <Route exact path="/games/two-thirds/revealed" render={(props) => ( <Revealed setParentState={this.setParentState} {...props} {...this.props} {...this.state} /> )} />
        <Route exact path="/games/two-thirds/payout" render={(props) => ( <PayoutPage {...this.props} {...this.state} /> )} />
      </div>
    )
  }
}

export default Game;