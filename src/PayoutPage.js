// @flow
import React, { Component } from 'react'
import { Route, Link } from "react-router-dom";
import { Button, ListGroup, ListGroupItem } from 'react-bootstrap';

import "./PayoutPage.css"

type props = {};

function PayoutMessage(props) {
  console.log(props.yournum);
  if(props.diff > 0){
    return <h2>Congratulations! You won. ETH! Why don't you keep your streak going and play again?</h2>;
  } else {
    return <h2>Your number was <span className="colored-num">{props.yourNum}</span>, unfortunately you didn't win. Better luck next time.</h2>;
  }
}

class PayoutPage extends Component<props> {

  state = {
    winningnumber: "50",
    yournum: "0",
    diff: "0"
  }
  
  componentDidMount = async () => {
    if(this.props.accounts == null) {
      return false
    }
    const account = this.props.accounts[0]
    // console.log('account: ' + account)
    const game = this.props.GameInstance;
    const web3 = this.props.web3; 

    const balance = await web3.eth.getBalance(account);
    console.log('balance : ' + balance)

    const temp_diff = (this.props.balance - balance) / 1000000000000000000;
    this.setState({diff: temp_diff});
  }

  render() {
    return (
      <div className="payout-page">
        <h2>The winning number was...<span className="colored-num">{this.state.winningnumber}</span>!</h2>
        <PayoutMessage diff={this.state.diff} yourNum={this.state.yournum}/>      
        <ListGroup>

          <ListGroupItem className="table-secondary-row">
            <p>Total pool</p>
            <p className="table-right-values">1 ETH</p>
          </ListGroupItem>

          <ListGroupItem className="table-secondary-row">
            <p>Your bet</p>
            <p className="table-right-values">1 ETH</p>
          </ListGroupItem>

          <ListGroupItem className="table-bottom-row">
            <p>Net result</p>
            <p className="table-right-values net-result-value">1 ETH</p>
          </ListGroupItem>

        </ListGroup>
        <Button className="play-again" onClick={() => this.props.history.push('/games/two-thirds/') }type='submit'>Play Again</Button>
      </div>
    )
  }
}

export default PayoutPage;