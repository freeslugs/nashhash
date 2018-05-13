// @flow
import React, { Component } from 'react';
import { BrowserRouter as Router, Route, Link } from 'react-router-dom';
import getWeb3 from './utils/getWeb3';

import Landing from './Landing'
import SelectPool from './SelectPool'

import { Button, Container, Header, Menu, Card } from 'semantic-ui-react'
import styled from 'styled-components';

const Logo = styled(Menu.Item)`
  font-family: 'Lobster Two', cursive;
  font-size: 2rem;
  padding-bottom: 12px !important;
`;

const activeItem = 'home'

const FullPage = styled.div`
  height: 100%;
  min-height: 100vh;
  width: 100%;
  padding: 3em 0;
`;

type props = {};

class App extends Component<props> {
  render() {
    return (
      <Router>
        <FullPage>
          <Container>
            <Menu pointing secondary>
              <Logo name='NashHash' as={Link} to="/" onClick={this.handleItemClick} />
              <Menu.Menu position='right'>
                <Menu.Item name='FAQ' active={activeItem === 'FAQ'} onClick={this.handleItemClick} />
                <Menu.Item name='friends' active={activeItem === 'friends'} onClick={this.handleItemClick} />
                <Menu.Item name='logout' active={activeItem === 'logout'} onClick={this.handleItemClick} />
              </Menu.Menu>
            </Menu>
        
            <Route exact path="/" render={(props) => ( <Landing {...props} {...this.state} /> )} />
            <Route path="/games/" render={(props) => ( <SelectPool {...props} {...this.state} /> )} />
          </Container>
        </FullPage>
      </Router>
    );
  }
}

export default App;
