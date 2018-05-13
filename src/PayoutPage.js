// @flow
import React, { Component } from 'react'
import { Route, Link } from "react-router-dom";
import { Button, Header, Checkbox, Form } from 'semantic-ui-react'

type props = {};

class PayoutPage extends Component<props> {

  render() {
    return (
      <div>
        <Header as='h2'>Congrats! you won 1 ETH.</Header>      
        {/*<Button onClick={() => this.props.history.push('/games/two-thirds/payout') }type='submit'>Reveal</Button>*/}
      </div>
    )
  }
}

export default PayoutPage;