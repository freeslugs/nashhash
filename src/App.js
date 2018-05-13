// @flow
import React, { Component } from 'react';
import { BrowserRouter as Router, Route, Link } from 'react-router-dom';
import getWeb3 from './utils/getWeb3';

import Landing from './Landing'
import Game from './Game'

import { Button, Container, Header, Menu, Card } from 'semantic-ui-react'
import styled from 'styled-components';

import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

const GameContract = require('./contracts/Game.json');
const contract = require('truffle-contract');

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
  state = {
    web3: null,
    accounts: null,
    GameInstance: null
  }

  async componentDidMount() { 
    const results = await getWeb3;
    this.setState({ web3: results.web3 })

    const Game = contract(GameContract)
    // const Game = GameContract.at("0xe4bf6b739f547a3d1d44501923048d11721a8d01")
    Game.setProvider(this.state.web3.currentProvider)

    this.state.web3.eth.getAccounts(async (error, accounts) => {

      let instance
      try {
        // console.log(this.stateweb3)
        // instance = await Game.deployed();  
        const network = await this.state.web3.eth.net.getNetworkType();
        if(network !== "rinkeby")
          throw(new Error("Game has not been deployed to detected network (network/artifact mismatch)"))
        instance = Game.at("0xa9a547abb048a35df0956ea9fa1768ceb118d86a") 
        // console.log(instance)
      } catch (e) {
        console.log(e)
        if(e.message == "Game has not been deployed to detected network (network/artifact mismatch)") {
          toast.error('Make sure Metamask is set to Rinkeby.', {
            position: "top-right",
            autoClose: false,
            hideProgressBar: false,
            closeOnClick: true,
            draggable: false,
            draggablePercent: 0
          })
          // return false;
        } else {
          console.log(e)
        }
      }
      this.setState({ accounts: accounts, GameInstance: instance });
      // console.log(instance.curr_number_bets)
      // const result = await instance.curr_number_bets();
      // const result = await instance.BET_SIZE();
      // console.log(parseInt(result))
    })
  }

  render() {
    return (
      <Router>
        <FullPage>
          <ToastContainer />
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
            <Route path="/games/two-thirds" render={(props) => ( <Game {...props} {...this.state} /> )} />
          </Container>
        </FullPage>
      </Router>
    );
  }
}

export default App;
