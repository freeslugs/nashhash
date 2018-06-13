pragma solidity ^0.4.23;

/*

!!!!!!!!!! LOGICAL IMPROVEMENTS !!!!!!!!!!!!
1) Send money to HOME address only once the fees reach a certain amount.
2) I have this idea of introducing a system of rounds so our players get a receit.
They can then use this receit to look up info about their game.

*/


/*
Dear all, games should inherit form this contract because this contract has the commit/reveal protocol
The specific game will have to only define two functions:

-- guessCheck(string guess): the function has to error out in if the guess is not compliant with the rules
-- findWinners(): the function determined the winners and distributes the payouts. 

*/

import "./Pausable.sol";
import "./GameHelper.sol";


contract Game is Pausable, GameHelper {

    uint public birthBlock = block.number;
    
    enum GameState {COMMIT_STATE, REVEAL_STATE, PAYOUT_STATE}

    // This is the idea.
    struct State {
        GameState gameState;
        uint gameStateDebug;
        uint currNumberCommits;
        uint currNumberReveals;
        uint commitStageStartBlock;
        uint revealStageStartBlock;
        uint round; // not implemented

    }

    struct Config {

        address FEE_ADDRESS;
        uint GAME_FEE_PERCENT;
        uint STAKE_SIZE;

        uint GAME_STAGE_LENGTH;
        uint MAX_PLAYERS;

        address NPT_ADDRESS;

    }

    struct GameInfo {
        address[] lastWinners;
        uint lastWinnersLength;
        uint lastPrize;
    }

    Config internal config;
    State internal state;
    GameInfo internal info;


    // Commit/Reveal Protocol vars
    mapping (address => bytes32) public commits;
    address[] internal commitsKeys;

    mapping (address => string) public gameData;
    address[] internal gameDataKeys;


    constructor(
        address _feeAddress,
        uint _gameFeePercent,
        uint _stakeSize,
        uint _maxp, 
        uint _gameStageLength,
        address _nptAddress) public {


        owner = msg.sender;

        config.GAME_STAGE_LENGTH = _gameStageLength;
        config.GAME_FEE_PERCENT = _gameFeePercent;
        config.MAX_PLAYERS = _maxp;
        config.STAKE_SIZE = _stakeSize;//1 ether;
        config.FEE_ADDRESS = _feeAddress;
        config.NPT_ADDRESS = _nptAddress;

        state.gameState = GameState.COMMIT_STATE;
        state.currNumberCommits = 0;
        state.currNumberReveals = 0;
        state.commitStageStartBlock = ~uint(0) - config.GAME_STAGE_LENGTH;
        state.revealStageStartBlock = ~uint(0) - config.GAME_STAGE_LENGTH;


        info.lastPrize = 0;

        commitsKeys = new address[](_maxp);
        gameDataKeys = new address[](_maxp);

        info.lastWinners = new address[](_maxp);
        info.lastWinnersLength = 0;



    }

    // Contrcact public API
    function getGameState() public view returns(uint) {
        return state.gameStateDebug;
    }

    function getCurrentCommits() public view returns(uint) {
        return state.currNumberCommits;
    }

    function getCurrentReveals() public view returns(uint) {
        return state.currNumberReveals;
    }

    function getStakeSize() public view returns(uint) {
        return config.STAKE_SIZE;
    }

    function setStakeSize(uint new_stake) public onlyOwner {
        config.STAKE_SIZE = new_stake;
    }

    function getNumberOfWinners() public view returns(uint) {
        return info.lastWinnersLength;
    }

    function getLastWinners(uint i) public view returns(address){
        require(i < info.lastWinnersLength);
        return info.lastWinners[i];
    }

    function getLastPrize() public view returns(uint){
        return info.lastPrize;
    }

    function getGameFee() public view returns(uint){
        return config.GAME_FEE_PERCENT;
    }

    function getMaxPlayers() public view returns(uint){
        return config.MAX_PLAYERS;
    }

    function setMaxPlayers(uint new_max) public onlyOwner {
        config.MAX_PLAYERS = new_max;
    }

    function getCommitStageStartBlock() public view returns(uint) {
        return state.commitStageStartBlock;
    }

    function getRevealStageStartBlock() public view returns(uint) {
        return state.revealStageStartBlock;
    }

    //ONLY FOR DEBUGGING PURPOSES! REMOVE LATER!!
    function setNPTAddress(address npt_addr) public onlyOwner {
        config.NPT_ADDRESS = npt_addr;
    }

    /*
        The following two functions are the users gaming interface.
            -- Call commit to commit a hash of your guess for the game. Its a hash, since
            you probably dont want other players to see your guess
            -- Call reveal to reveal your guess. You will not participate in the
            round if you forget to reveal your guess, but your stake will still become
            someone's prize! Make sure you reveal.
    */
    function commit(bytes32 hashedCommit) public payable whenNotPaused {
        
        require(state.gameState == GameState.COMMIT_STATE);

        require(msg.value == config.STAKE_SIZE);

        //Make sure this is first and only commit. 
        require(commits[msg.sender] == bytes32(0x0));

        commits[msg.sender] = hashedCommit;
        commitsKeys[state.currNumberCommits] = msg.sender;
        state.currNumberCommits++;

        // Start the 'commit stage timer', to protect the players in case the
        // game master goes rogue
        if (state.currNumberCommits == 1) {
            state.commitStageStartBlock = block.number;
        }

        // If we received the MAX_PLAYER number of commits, it is time for
        // us to change state.
        if (state.currNumberCommits == config.MAX_PLAYERS) {
            toRevealState();
        }
    }

    function reveal(string guess, string random) public whenNotPaused {
        
        require(state.gameState == GameState.REVEAL_STATE);
        
        // Function checks if the guess fits the requirement of the game
        checkGuess(guess);

        // Check that the hashes match
        require(commits[msg.sender] == keccak256(guess, random));

        //Prevents user from revealing twice because above require will fail.
        delete commits[msg.sender];

        // When they do, we add the revealed guess to game data
        gameData[msg.sender] = guess;
        gameDataKeys[state.currNumberReveals] = msg.sender;
        state.currNumberReveals++;

        if(state.currNumberReveals == config.MAX_PLAYERS){
            toPayoutState();
        }
    }



    /*
        The following three functions are GameMaster controlled functions.
        GM is responsible for the state transitions of the contracts.
    */
    function forceToRevealState() public onlyOwner whenNotPaused {
        require(state.gameState == GameState.COMMIT_STATE);
        //require(state.commitStageStartBlock + config.GAME_STAGE_LENGTH <= block.number);
        toRevealState();
    }

    function forceToPayoutState() public onlyOwner whenNotPaused {
        require(state.gameState == GameState.REVEAL_STATE);
        //require(state.revealStageStartBlock + config.GAME_STAGE_LENGTH <= block.number);
        toPayoutState();
    }

    function payout() public onlyOwner whenNotPaused {
        require(state.gameState == GameState.PAYOUT_STATE);
        findWinners();
        toCommitState();
    } 

    /* 
        Function resets the game contract state to deployment state.
        This is an emergency function.
    */
    function resetGame() public onlyOwner whenNotPaused {  
        
        toCommitState();

        info.lastPrize = 0;
        
        uint i;
        for (i = 0; i < info.lastWinnersLength; i++){
            delete info.lastWinners[i];
        }
        info.lastWinnersLength = 0;
    }



    /* 
        Function parforms a payout of money to the winners.
        The function has to be called by the child contract once the winners are found
        i.e the function has to be called in the end of findWinners().
        TODO: Can performPayout call be better placed?
    */
    function performPayout(address[] winners, uint numWinners, uint prize) internal {
        uint i = 0;

        info.lastWinnersLength = numWinners;
        for(i = 0; i < numWinners; i++){
            winners[i].transfer(prize); 
            info.lastWinners[i] = winners[i];
        }
        info.lastPrize = prize;
    } 





    /* 
        These are the abstract functions that the inheriting game
        contracts will have to define. 
    */
    function checkGuess(string guess) private;
    function findWinners() private;






    /* Contract state transition helpers */
    event CommitsSubmitted();
    function toRevealState() internal {
        state.gameState = GameState.REVEAL_STATE;
        state.gameStateDebug = 1;
        state.revealStageStartBlock = block.number;
        emit CommitsSubmitted();

    }

    event RevealsSubmitted();
    function toPayoutState() internal {
        state.gameState = GameState.PAYOUT_STATE;
        state.gameStateDebug = 2;
        emit RevealsSubmitted();
    }

    event NewRoundStarted();
    function toCommitState() internal {
        // Set state
        state.gameState = GameState.COMMIT_STATE;
        state.gameStateDebug = 0;

        // Cleanup the commits
        uint i;
        for(i = 0; i < state.currNumberCommits; i++){
            delete commits[commitsKeys[i]];
        }
        
        state.currNumberCommits = 0;
        state.currNumberReveals = 0;

        // TODO: Find a better way to do this
        state.commitStageStartBlock = ~uint(0) - config.GAME_STAGE_LENGTH;
        state.revealStageStartBlock = ~uint(0) - config.GAME_STAGE_LENGTH;

        emit NewRoundStarted();
    }




}