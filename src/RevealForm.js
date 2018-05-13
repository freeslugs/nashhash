// @flow
import React, { Component } from 'react'
import { Route, Link } from "react-router-dom";
import { Button, Header, Checkbox, Form } from 'semantic-ui-react'

type props = {};

class RevealForm extends Component<props> {

  render() {
    return (
      <div>
        <Header as='h2'>Reveal your guess!</Header>      
        <Button onClick={() => this.props.history.push('/games/two-thirds/payout') }type='submit'>Reveal</Button>
      </div>
    )
  }
}

export default RevealForm;