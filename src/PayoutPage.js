// @flow
import React, { Component } from 'react'
import { Route, Link } from "react-router-dom";
import { Button, Header, Checkbox, Form } from 'semantic-ui-react'

type props = {};

class PayoutPage extends Component<props> {

  state = {
    text: "Congrats, you won 1 ETH!"
  }
  
  // componentWillMount() {
  //   const account = this.props.accounts[0]

  //   const game = this.props.GameInstance;

  //   // var number_of_winners = await game.num_last_winners();
  //   //   for (i = 0; i < number_of_winners; i++){
  //   //     var winner = await game.last_winners(i);
  //   //     if(winner == account){
  //   //       var prize = await game.last_prize();
  //   //       this.setState({text: `congrats, you won ${prize} ETH!`})
  //   //     }
  //   //   }

  // }

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