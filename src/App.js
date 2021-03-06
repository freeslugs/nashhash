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

import API from './api/Game.js'

var Web3Utils = require('web3-utils');

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
    game: null
  }

  async componentDidMount() { 
    const results = await getWeb3;
    const web3 = results.web3;
    this.setState({ web3 })

    const Game = contract(GameContract)
    Game.setProvider(web3.currentProvider)

    const accounts = await new Promise(function(resolve, reject) {
      web3.eth.getAccounts( function (err, accounts) { resolve(accounts) })
    });
    this.setState({ accounts: accounts });

    let instance
    const network = await web3.eth.net.getNetworkType();

    if(network === "private") { // localhost:9545
      instance = await Game.deployed();  
      toast('Configured with local network. Success!', {
        position: "top-right",
        autoClose: true,
        hideProgressBar: false,
        closeOnClick: true,
        draggable: false,
        draggablePercent: 0
      })
    } else if(network === "rinkeby") {
      instance = Game.at("0xec9a2508a775d34d49625a860829f5733fbd4bc6") 
      toast('Configured with Rinkeby network. Success!', {
        position: "top-right",
        autoClose: true,
        hideProgressBar: false,
        closeOnClick: true,
        draggable: false,
        draggablePercent: 0
      })
    } else {  
      toast.error('Make sure Metamask is set to Rinkeby.', {
        position: "top-right",
        autoClose: false,
        hideProgressBar: false,
        closeOnClick: true,
        draggable: false,
        draggablePercent: 0
      })
      return false;
    }
  
    this.setState({ game: instance });
  }

  resetGame = async () =>  {
    const { web3, game, accounts } = this.state;
    const api = new API(web3.utils, () => {}, game);
    console.log(api)
    try {
      // await api.setMaxPlayers(2, accounts[0])
      await api.resetGame(accounts[0])
      console.log('fuck yes')
    } catch(e){
      console.log(e)
    }
  }

  render() {
    return (
      <Router>
        <FullPage>
          <ToastContainer />
          <Container>
            <Menu pointing secondary>
              <Logo name='NashHash' as={Link} to="/" />
              <Menu.Menu position='right'>
                <Menu.Item name='FAQ' active={activeItem === 'FAQ'} as={Link} to="/games/two-thirds/payout" />
                <Menu.Item name='friends' active={activeItem === 'friends'} />
                <Menu.Item name='logout' onClick={this.resetGame} active={activeItem === 'logout'} />
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
