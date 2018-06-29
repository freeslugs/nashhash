// @flow
import React, { Component } from 'react'
import { Link } from "react-router-dom";
import { Icon, Step, Container, Header, Menu, Card } from 'semantic-ui-react'
import { Well, Grid, Row, Col, FormGroup, FormControl, Button, ButtonToolbar, ToggleButton, ToggleButtonGroup } from 'react-bootstrap'
import styled from 'styled-components'
import testLogo from './img/group.svg'

import GameRegistryAPI from './api/GameRegistryAPI.js'
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



class CommitForm extends Component<props> {
  state = {
    guess: null,
    loading: false
  }

  componentDidMount() {
    this.props.setParentState({ state: "COMMIT" })
  }

  loading() {
    const { web3, accounts, gameregistry, gameaddresses } = this.props;
    return !(web3 && accounts && accounts.length > 0 && gameaddresses.length > 0)
    //Add statement to reroute if gameaddresses is null
  }

  commit = async () => {
    const web3 = this.props.web3;
    const account = this.props.accounts[0];
    const gameaddresses = this.props.gameaddresses;
    const gametype = this.props.gametype;
    const stake = this.props.stake;

    const registryAPI = new GameRegistryAPI(web3, gameaddresses);

    const game = await registryAPI.configureGame(gametype, stake);

    const gameAPI = new API(web3.utils, () => {}, game);

    const hashKey = /*this.props.hashKey || */ srs({length: 50});
    this.props.setParentState({ hashKey })

    const guess = this.state.guess

    if(guess == null || guess.length == 0) {
      console.log('no guess!')
      return false;
    }

    this.setState({loading: true})
    try {
      await gameAPI.commitGuess(account, guess, hashKey);
    } catch(e) {
      console.log(`error: ${e}`)
    }

    this.setState({loading: false})
    this.props.setParentState({ guess: guess, state: "COMMITTED" })

    if(gametype == "TwoThirds"){
      this.props.history.push('/games/two-thirds/committed')
    }
    else if(gametype == "LowestUnique"){
      this.props.history.push('/games/lowest-unique/committed')
    }
  }

  render() {
    return (
      <div className="wrapper">
      {/*  <Header as='h1'>{gametitle(gameType)}</Header>
        <Header as='h2'>{gameinfo(gameType)}</Header>   */}

        <Grid className="game-info-grid">
          <Col className="game-info-logo" xs={12} sm={12} md={4}>
            <Well bsClass="game-drawing-card">
              <img src={testLogo} className="two-thirds-logo"/>
            </Well>
          </Col>
          <Col className="game-info-text" xs={12} sm={12} md={8}>
            <Well bsClass="game-info-card">
              <h1>2/3 Average</h1>
              <p>Select a number between 0-100 with the intention of guessing 2/3 of the average guess. </p>
              <p>The winner receives the total prize pool (10 * game stake) minus a 5% platform commission. </p>
              <p> <span>Example:</span> <br/>
                  The sum of the 10 players’ guesses is 653.
                  The average guess is therefore 650/10, 65.
                  2/3 of the average guess is 65*2/3, 43.29.
                  The winner is the person who guessed closest to 43.29.
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
                    <h1>Pick a number from 1-100</h1>
                </Col>

                <Col className="pick-number-entry" xs={12} sm={12} md={6}>
                      <FormGroup controlId="number-entry-field">
                        <FormControl
                          type="text"
                          placeholder="Enter guess"
                        />
                      </FormGroup>
                </Col>

              </Row>

              <Row className="pick-stake">
                <Col className="pick-stake-info" xs={12} sm={12} md={6}>

                    <h1>Choose your stakes</h1>

                </Col>
                <Col className="pick-stake-entry" xs={12} sm={12} md={6}>

                    <ButtonToolbar>
                      <ToggleButtonGroup type="radio" name="options" defaultValue={1}>
                        <ToggleButton className="stake-button" value={1}>0.01 ETH</ToggleButton>
                        <ToggleButton className="stake-button" value={2}>0.1 ETH</ToggleButton>
                        <ToggleButton className="stake-button" value={3}>1 ETH</ToggleButton>
                      </ToggleButtonGroup>
                    </ButtonToolbar>

                </Col>
              </Row>
            </Grid>
          </Well>

          <div className="submit-container">
            <Button className="submit-guess-button" type="submit">Submit Guess</Button>
            <p>You’ll see a MetaMask pop-up asking to approve a
            transaction for this amount, plus gas costs. <br/>Submit to
            send your bet to the blockchain.</p>
          </div>
          </FormGroup>
        </form>
      </div>
    )
  }
}
export default CommitForm;
