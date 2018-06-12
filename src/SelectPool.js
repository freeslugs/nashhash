// @flow
import React, { Component } from 'react'
import { Link } from "react-router-dom";
import { Icon, Step, Button, Container, Header, Menu, Card } from 'semantic-ui-react'
import styled from 'styled-components';
  
const Wrapper = styled.div`
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;

  h1.ui.header {
    margin-top: 1em;
    font-size: 2.7em;
    font-weight: normal;
    text-align: center;
  }

  h2.ui.header {
    font-size: 1.7em;
    margin-top: 1.5em;
    font-weight: normal;
    text-align: center;
  }
`;

const Games = styled.div`
  h2.ui.header {
    font-style: italic;
    font-size: 2em;
    margin-top: 3em;
    margin-bottom: 1.5em;
    font-weight: normal;
    text-align: center;
  }

  display: flex;
  flex-direction: column;
  align-items: center;
  .ui.three.cards { 
    min-width: 700px;
  }

  .ui.cards>.card>.content>.header:not(.ui) {
    font-size: 2em;
  }
`;

const SelectPool = ({ setParentState, history }) => <Wrapper>
  <Header as='h1'>2/3 Average</Header>      
  <Header as='h2'>Select a number between 0-100 with the intention of guessing 2/3 of the average guess.</Header>      

  <Step.Group unstackable>
    <Step>
      <Icon name='sign in' color='orange'/>
      <Step.Content>
        <Step.Title>Commit</Step.Title>
        <Step.Description>Everyone submits a guess.</Step.Description>
      </Step.Content>
    </Step>
    <Step>
      <Icon name='unlock' color='purple' />
      <Step.Content>
        <Step.Title>Reveal</Step.Title>
        <Step.Description>Everyone reveals their guess.</Step.Description>
      </Step.Content>
    </Step>
    <Step>
      <Icon name='child' color='green' />
      <Step.Content>
        <Step.Title>Payout</Step.Title>
        <Step.Description>Find out who won!</Step.Description>
      </Step.Content>
    </Step>
  </Step.Group>

  <Games>
    <Header as='h2'>Select your stake.</Header>  

    <Card.Group itemsPerRow={3}>
      <Card color='blue' onClick={() => { setParentState({stake: 0.01}); history.push('/games/two-thirds/commit') } }>
        <Card.Content>
          <Card.Header>.01 ETH</Card.Header>
          <Card.Description>5 / 10 players</Card.Description>
        </Card.Content>
      </Card>

      <Card color='yellow' onClick={() => { setParentState({stake: 0.1}); history.push('/games/two-thirds/commit') } }>
        <Card.Content>
          <Card.Header>.1 ETH</Card.Header>
          <Card.Description>5 / 10 players</Card.Description>
        </Card.Content>
      </Card>

      <Card color='red' onClick={() => { setParentState({stake: 1}); history.push('/games/two-thirds/commit') } }>
        <Card.Content>
          <Card.Header>1 ETH</Card.Header>
          <Card.Description>5 / 10 players</Card.Description>
        </Card.Content>
      </Card>      
    </Card.Group>
  </Games>
</Wrapper>

export default SelectPool;