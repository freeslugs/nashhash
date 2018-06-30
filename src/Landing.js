// @flow
import React, { Component } from 'react'
import { Link } from "react-router-dom";
import { LinkContainer } from "react-router-bootstrap"

import styled from 'styled-components';
import { Well, ListGroup, ListGroupItem, Button, Grid, Row, Col } from 'react-bootstrap'
import './Landing.css'
import fairLogo from './img/group.svg'

const Landing = () => <div>
  <div class="banner">
    <h1>Provably fair betting games, on the Ethereum blockchain.</h1>
    <h2>Game theory games. Real people. Real economic incentives.</h2>
  </div>



    <Grid className="cards-container">

    <Row className="show-grid">

      <Col className="individual-card-container" xs={12} md={6}>
          <Well bsClass="card">
            <div className="gradient-rectangle">
            </div>

             <h1>2/3 Average</h1>
             <h3>10 player game </h3>
             <h3>Select a number between 0-100 with the intention of guessing 2/3 of the average guess</h3>
             <LinkContainer to="/games/two-thirds">
               <Button bsClass="play">Play</Button>
             </LinkContainer>
          </Well>
      </Col>

      <Col className="individual-card-container" xs={12} md={6}>
           <Well bsClass="card">
           <div className="gradient-rectangle">
           </div>
           <div className="card-logo">
           </div>

               <h1>Lowest Unique Number</h1>
               <h3>10 player game</h3>
               <h3>Choose the lowest number that no one else has picked.</h3>
               <LinkContainer to="/games/lowest-unique">
                 <Button bsClass="play">Play</Button>
               </LinkContainer>
           </Well>
       </Col>

     </Row>

 </Grid>


  <p className="more-games-coming"> More exciting games in progress! </p>

  <div>
  <Grid className="value-props">

    <Row className="show-grid value-props-row">
      <Col className="value-props-grid-item" xs={12} sm={12} md={6}>
        <Well className="facts-drawing-card">
        <img src={fairLogo} className="faq-logos"/>
        </Well>
      </Col>
      <Col className="value-props-grid-item" xs={12} sm={12} md={6}>
        <Well className="facts-card">
        <h1>Provably Fair Games</h1>
        <h2>Open source game mechanics for all to look at.</h2>
        </Well>
      </Col>
    </Row>

    <Row className="show-grid value-props-row">
      <Col className="value-props-grid-item" xs={12} sm={12} md={6}>
        <Well className="facts-card">
        <h1>User Funds Safe</h1>
        <h2>No third party custody of user assets. Winners are paid automatically.</h2>
        </Well>
      </Col>
       <Col className="value-props-grid-item" xs={12} sm={12} md={6}>
        <Well className="facts-drawing-card">
        <img src={fairLogo} className="faq-logos"/>
        </Well>
      </Col>
    </Row>

    <Row className="show-grid value-props-row">
      <Col className="value-props-grid-item" xs={12} sm={12} md={6}>
        <Well className="facts-drawing-card">
        <img src={fairLogo} className="faq-logos"/>
        </Well>
      </Col>
      <Col className="value-props-grid-item" xs={12} sm={12} md={6}>
        <Well className="facts-card">
        <h1>Win NashPoints</h1>
        <h2>All players receive NPS, with top 100 NPS holders entered into a weekly prize pool.</h2>
        </Well>
      </Col>
    </Row>

  </Grid>
  </div>
</div>

export default Landing;
