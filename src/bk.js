module.exports = async (deployer, network, accounts) => {
  await deployer.deploy(Game, 10);

  const metaDataFile = `${__dirname}/../build/contracts/Game.json`
  const metaData = require(metaDataFile)
  metaData.networks[deployer.network_id] = {}
  metaData.networks[deployer.network_id].address = Game.address
  fs.writeFileSync(metaDataFile, JSON.stringify(metaData, null, 4))
}

let curr_number_bets = await game.curr_number_bets();
      console.log('current bets: ' + parseInt(curr_number_bets))
      let state = await game.game_state_debug();
      console.log('current state: ' +  parseInt(state))

      // console.log(parseInt(result))



await API.commitGuess(game, donor, this.state.local_guess, "3");

const curr_number_bets = await game.getCurrentCommits();
const guess_commit = await game.commits(donor);

const game = this.props.GameInstance;
const account = this.props.accounts[0]
// console.log(account)

const bet = await game.BET_SIZE();
//Add random number generator
const hash = Web3Utils.soliditySha3({type: 'string', value: this.state.local_guess}, {type: 'string', value: "3"});

this.props.setGuess(this.state.local_guess)

// const curr_number_bets = await game.curr_number_bets();
// console.log(parseInt(curr_number_bets))
// await game.set_MAX_PLAYERS(1, { from: account, gasPrice: 80000000000 });
// await game.commit(hash, { value: bet, from: donor });
// await game.reveal("66", "3", {from: donor});

this.setState({loading: true})
try {
  await game.commit(hash, { value: bet, from: account, gasPrice: 80000000000 });
} catch(e) {
  console.log(e)
}
this.setState({loading: false})


// assert.equal(curr_number_bets, 1, "Number of bets did not increment");
// assert.equal(guess_commit, hash, "Hashes do not match");

this.props.history.push('/games/two-thirds/reveal')