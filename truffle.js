require('dotenv').config()

require('babel-register')
require('babel-polyfill')

var HDWalletProvider = require("truffle-hdwallet-provider");
var words = process.env.WORDS;

if(words == undefined) {
  console.log(`
⚠️  Important ⚠️  : You are missing WORDS in the .env file. Check readme on howto. 
`)
}

module.exports = {
  mocha: {
    "scripts": {
      "test": "mocha --require babel-register"
    }
  },
  networks: {
    development: {
      host: "127.0.0.1",
      port: 9545,
      network_id: "*" // Match any network id
    },
    azure: {
      network_id: "*", // Match any network id
      provider: function() {
        return new HDWalletProvider(words, 'http://23.101.141.196:8545')
      },
      gas: 1000000
    },
    rinkeby: {
      network_id: 4,
      provider: function() {
        return new HDWalletProvider(words, 'https://rinkeby.infura.io/SpARW7NJTyVseQE3d8Bs')
      },
      "gas": 4000000,
      "gasPrice": 4000000000
    }
  }
};