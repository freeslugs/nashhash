// @flow
import React, { Component } from 'react'
import { Link } from "react-router-dom";
import { LinkContainer } from "react-router-bootstrap"
import { Container, Header, Menu, Card } from 'semantic-ui-react'
import styled from 'styled-components';
import { Well, Button, Grid, Row, Col } from 'react-bootstrap'
import './Landing.css'
import fairLogo from './img/group.svg'

import styles from './test.css';

const Landing = () => <div>
  <div class="banner">
    <Header as='h1'>Game theory Games on Ethereum</Header>      
    <Header as='h2'>If we asked them what they wanted, they would have said more interns.</Header>      
  </div>
  <div class="games">
    <Header as='h2'>Play now!</Header>   
    <div class="cardcontainer">
     
      <Well bsClass="card">  
        <div class="cardcontent">
          <h1>2/3 Average</h1>
          <h3>10 player game</h3>
          <h3>Select a number between 0-100 with the intention of guessing 2/3 of the average guess.</h3>
          <LinkContainer to="/games/two-thirds">
            <Button bsClass="play">Play Game</Button>
          </LinkContainer>
        </div>
      </Well>

      <Well bsClass="card">
        <div class="cardcontent">
          <h1>Lowest Unique Number</h1>
          <h3>10 player game</h3>
          <h3>Choose the lowest number that no one else has picked.</h3>
          <LinkContainer to="/potato">
            <Button bsClass="play">Play Game</Button>
          </LinkContainer>
        </div>
      </Well>
    </div>
  </div>
  <br/>
  <br/>
  <br/>
  <div>
  <Grid>
    <Row className="show-grid">
      <Col sm={6} md={5}>
        <Well bsClass=".facts-drawing-card">
        <img src={fairLogo} bsClass="faq-logos"/>
        </Well>
      </Col>
      <Col sm={6} md={5}>
        <Well bsClass="facts-card">
        <h1>Provably Fair Games</h1>
        <h2>Open source game mechanics for all to look at.</h2>
        </Well>
      </Col>
    </Row>
    <Row className="show-grid">
      <Col sm={6} md={5}>
        <Well bsClass="facts-card">
        <h1>User Funds Safe</h1>
        <h2>No third party custody of user assets. Winners are paid automatically.</h2>
        </Well>
      </Col>
       <Col sm={6} md={5}>
        <Well bsClass=".facts-drawing-card">
        <img src={fairLogo} bsClass="faq-logos"/>
        </Well>
      </Col>
    </Row>
    <Row className="show-grid">
      <Col sm={6} md={5}>
        <Well bsClass=".facts-drawing-card">
        <img src={fairLogo} bsClass="faq-logos"/>
        </Well>
      </Col>
      <Col sm={6} md={5}>
        <Well bsClass="facts-card">
        <h1>Win NashPoints</h1>
        <h2>All players receive NPS, with top 100 NPS holders entered into a weekly prize pool.</h2>
        </Well>
      </Col>
    </Row>
  </Grid>
  </div>
</div>

export default Landing;