// @flow
import React, { Component } from 'react'
import { Link } from "react-router-dom";
import { Icon, Step, Container, Header, Menu, Card } from 'semantic-ui-react'
import { Well, Grid, Row, Col, FormGroup, FormControl, Button, ButtonToolbar, ToggleButton, ToggleButtonGroup } from 'react-bootstrap'
import styled from 'styled-components' 
import testLogo from './img/group.svg'
import "./CommitForm.css"

function gametitle(gtyp) {
  if(gtyp == "TwoThirds"){
    return "2/3 Average"
  }
  else if(gtyp == "LowestUnique"){
    return "Lowest Unique Number"
  }
}

function gameinfo(gtyp) {
  if(gtyp == "TwoThirds"){
    return "Select a number between 0-100 with the intention of guessing 2/3 of the average guess."
  }
  else if(gtyp == "LowestUnique"){
    return "Select the lowest number (0 or greater) that no one else has picked."
  }
}

const SelectPool = ({setParentState, gameType, history}) => <div className="wrapper">
{/*  <Header as='h1'>{gametitle(gameType)}</Header>      
  <Header as='h2'>{gameinfo(gameType)}</Header>   */}    

  <Grid className="game-info-grid">
    <Col xs={12} sm={12} md={4}>
      <Well bsClass="game-drawing-card">
        <img src={testLogo} className="two-thirds-logo"/>
      </Well>
    </Col>
    <Col xs={12} sm={12} md={8}>
      <Well bsClass="game-info-card">
        <h1>2/3 Average</h1>
        <p>Select a number between 0-100 with the intention of guessing 2/3 of the average guess. </p>
        <p>The winner receives the total prize pool (10 * game stake) minus a 5% platform commissio</p>
        <p> Example:
            The sum of the 10 players’ guesses is 653 
            The average guess is therefore 650/10, 65
            2/3 of the average guess is 65*2/3, 43.29
            The winner is the person who guessed closest to 43.29
        </p>
      </Well>
    </Col>
  </Grid>
  <form>
    <FormGroup controlId="submit-commit-form">
    <Well className="game-options">
      <Grid>
        <Row className="pick-number">
          <Col className="pick-number-info" xs={12} sm={12} md={6}>
            <Well className="number-info-card">
              <h1>Pick a number from 1-100</h1>
            </Well>
          </Col>
          <Col className="pick-number-entry" xs={12} sm={12} md={6}>
            <Well className="number-entry-card">
                <FormGroup controlId="number-entry-field">
                  <FormControl
                    type="text"
                    placeholder="Enter guess"
                  />
                </FormGroup>
            </Well>
          </Col>
        </Row>
        <Row className="pick-stake">
          <Col className="pick-stake-info" xs={12} sm={12} md={6}>
            <Well className="stake-info-card">
              <h1>Choose your stakes</h1>
            </Well> 
          </Col>
          <Col className="pick-stake-entry" xs={12} sm={12} md={6}>
            <Well className="stake-entry-card">
              <ButtonToolbar>
                <ToggleButtonGroup type="radio" name="options" defaultValue={1}>
                  <ToggleButton className="stake-button" value={1}>0.01 ETH</ToggleButton>
                  <ToggleButton className="stake-button" value={2}>0.1 ETH</ToggleButton>
                  <ToggleButton className="stake-button" value={3}>1 ETH</ToggleButton>
                </ToggleButtonGroup>
              </ButtonToolbar>
            </Well>
          </Col>
        </Row>
      </Grid>
    </Well>

    <div className="submit-container">
      <Button type="submit">Submit Guess</Button>
      <p>You’ll see a MetaMask pop-up asking to approve a 
      transaction for this amount, plus gas costs. Submit to 
      send your bet to the blockchain.</p>
    </div>
    </FormGroup>
  </form>

{/*
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
*/}
</div>

export default SelectPool;