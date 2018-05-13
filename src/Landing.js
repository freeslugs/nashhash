// @flow
import React, { Component } from 'react'
import { Link } from "react-router-dom";
import { Button, Container, Header, Menu, Card } from 'semantic-ui-react'
import styled from 'styled-components';
import './Landing.css'

const Banner = styled(Container)`
  h1.ui.header {
    margin-top: 3em;
    margin-bottom: 0;
    font-size: 3em;
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
`;

const MoreComingSoon = styled(Card)`
  .content {
    display: flex;
    align-items: center;
    justify-content: center;
  }
  background: repeating-linear-gradient(
    -55deg,
    #ffffff,
    #ffffff 10px,
    #E0E1E2 10px,
    #E0E1E2 20px
  ) !important;
`;

const Landing = () => <div>
  <Banner>
    <Header as='h1'>Game theory Games on Ethereum</Header>      
    <Header as='h2'>If we asked them what they wanted, they would have said more interns.</Header>      
  </Banner>

  <Games>
    <Header as='h2'>Play now!</Header>  

    <Card.Group itemsPerRow={4}>
      <Card color='blue' as={Link} to="/games/two-thirds">  
        <Card.Content>
          <Card.Header>2/3 Average</Card.Header>
          <Card.Meta>10 player game</Card.Meta>
          <Card.Description>Select a number between 0-100 with the intention of guessing 2/3 of the average guess.</Card.Description>
        </Card.Content>
      </Card>

      <Card color='green' as={Link} to="/games/lowest-unique">
        <Card.Content>
          <Card.Header>Lowest Unique Number</Card.Header>
          <Card.Meta>10 player game</Card.Meta>
          <Card.Description>Choose the lowest number that no one else has picked.</Card.Description>
        </Card.Content>
      </Card>

      <Card color='purple' as={Link} to="/games/chicken">
        <Card.Content>
          <Card.Header>Chicken</Card.Header>
          <Card.Meta>2 player game</Card.Meta>
          <Card.Description>Two players and a countdown timer: whoever clicks stop first, loses.</Card.Description>
        </Card.Content>
      </Card>

      <MoreComingSoon color='red' className="invalid">
        <Card.Content>
          <Card.Header>More coming soon</Card.Header>
        </Card.Content>
      </MoreComingSoon>
    </Card.Group>
  </Games>
</div>

export default Landing;