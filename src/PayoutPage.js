// @flow
import React, { Component } from 'react'
import { Route, Link } from "react-router-dom";
import { Button, Header, Checkbox, Form } from 'semantic-ui-react'

type props = {};

class PayoutPage extends Component<props> {

  state = {
    text: "loading..."
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

    const diff = this.props.balance - balance / 1000000000000000000;

    // this.setState({text: `diff: ${diff}`})

    if(diff > 0) {
      this.setState({text: `Congrats, you won ${diff} ETH!`})
    } else {
      this.setState({text: `You lost.`});
    }

    // console.log(web3)

    // console.log(game.num_last_winners)

    // var prize = await game.get_();
    // console.log('prize: ' + prize);


    // // var prize = await game.last_prize();
    // // console.log('prize: ' + prize);

    // console.log(game)

    // var number_of_winners = await game.num_last_winners();
    // for (var i = 0; i < number_of_winners; i++){
    //   var winner = await game.last_winners(i);
    //   // if(winner == account){
    //   //   var prize = await game.last_prize();
    //   //   this.setState({text: `congrats, you won ${prize} ETH!`})
    //   // }


    //   console.log("winner: " + winner)
    // }

  }

  render() {
    return (
      <div>
        <Header as='h2'>{this.state.text}</Header>      
        {/*<Button onClick={() => this.props.history.push('/games/two-thirds/payout') }type='submit'>Reveal</Button>*/}
      </div>
    )
  }
}

export default PayoutPage;