// @flow
import React, { Component } from 'react';
import { BrowserRouter as Router, Route, Link } from 'react-router-dom';
import getWeb3 from './utils/getWeb3';

import Landing from './Landing'
import Game from './Game'

import { Button, Container, Header, Card, Divider, Icon } from 'semantic-ui-react'
import { Navbar, Nav, NavItem, MenuItem, NavDropdown } from 'react-bootstrap'
import styled from 'styled-components';

import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

import "./App.css"

import API from './api/Game.js'

var Web3Utils = require('web3-utils');

const GameABI = require('./contracts/Game.json');
const GameRegistryABI = require('./contracts/GameRegistry.json')
const contract = require('truffle-contract');

/* const Logo = styled(Menu.Item)`
  font-family: 'Lobster Two', cursive;
  font-size: 2rem;
  padding-bottom: 12px !important;
`;
*/

const activeItem = 'home'

type props = {};

class App extends Component<props> {
  state = {
    web3: null,
    accounts: null,
    gameregistry: null,
    gameaddresses: null
  }

  async componentDidMount() { 
    const results = await getWeb3;
    const web3 = results.web3;
    this.setState({ web3 })

    //const Game = contract(GameABI)
    //Game.setProvider(web3.currentProvider)

    const GameRegistryContract = contract(GameRegistryABI)
    GameRegistryContract.setProvider(web3.currentProvider)

    const accounts = await new Promise(function(resolve, reject) {
      web3.eth.getAccounts( function (err, accounts) { resolve(accounts) })
    });
    this.setState({ accounts: accounts });

    let instance, addresses
    const network = await web3.eth.net.getNetworkType();

    if(accounts.length == 0) {
      toast.error('Be sure to log into Metamask.', {
        position: "top-right",
        autoClose: false,
        hideProgressBar: false,
        closeOnClick: true,
        draggable: false,
        draggablePercent: 0
      })
      return false
    }
    if(network === "private") { // localhost:9545
      //instance = await Game.deployed(); 
      instance = await GameRegistryContract.deployed(); 
      addresses = await instance.getGameAdresses();
      toast('Configured with local network. Success!', {
        position: "top-right",
        autoClose: true,
        hideProgressBar: false,
        closeOnClick: true,
        draggable: false,
        draggablePercent: 0
      })
    } else if(network === "rinkeby") {
      //instance = Game.at("0xec9a2508a775d34d49625a860829f5733fbd4bc6") 
      instance = await GameRegistryContract.deployed();
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
    this.setState({ gameregistry: instance, gameaddresses: addresses});
  }

  //Fix this later
  /*
  resetGame = async () =>  {
    const { web3, accounts, game } = this.state;
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
  */

  render() {
    return (
      <Router>
        <div className="fullpage">
          <ToastContainer />
          <div className="fullcontainer">
            <Navbar inverse collapseOnSelect>
              <Navbar.Header>
                <Navbar.Brand>
                  <a href="#brand">Nashhash</a>
                </Navbar.Brand>
                <Navbar.Toggle />
              </Navbar.Header>
              <Navbar.Collapse>
                <Nav pullRight>
                  <NavItem eventKey={1} href="#">
                    FAQ
                  </NavItem>
                  <NavItem eventKey={2} href="#">
                    Wallet
                  </NavItem>
                </Nav>
              </Navbar.Collapse>
            </Navbar>
        
            <Route exact path="/" render={(props) => ( <Landing {...props} {...this.state} /> )} />
            <Route path="/games/two-thirds" render={(props) => ( <Game GameType={"TwoThirds"} {...props} {...this.state} /> )} />
            <Route path="/games/lowest-unique" render={(props) => ( <Game GameType={"LowestUnique"} {...props} {...this.state} /> )} />

            <Divider section />
            <div>
            <Navbar collapseOnSelect>
              <Navbar.Collapse>
                <Nav>
                  <NavItem eventKey={1} href="#">
                    Privacy Policy
                  </NavItem>
                  <NavItem eventKey={2} href="#">
                    Terms of Use
                  </NavItem>
                </Nav>
                <Nav pullRight>
                  <NavItem eventKey={1} href="#">
                    Telegram
                  </NavItem>
                  <NavItem eventKey={2} href="#">
                    Twitter
                  </NavItem>
                  <NavItem eventKey={2} href="#">
                    GitHub
                  </NavItem>
                  <NavItem eventKey={2} href="#">
                    Made with __ in NYC
                  </NavItem>
                </Nav>
              </Navbar.Collapse>
            </Navbar>
            </div>
          </div>
        </div>
      </Router>
    );
  }
}

export default App;
