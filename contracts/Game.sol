pragma solidity ^0.4.23;



/*
!!!!!!!!!!! KNOWN BUGS !!!!!!!!!!!!!!!
1) 2/3 average consistent with the js results. BUT. The way we calculate it
in solidity seems to be rounded down at the taking of the average step, and then
again rounded down after division by 3. It is better to fix this because this skews
the average ever so slightly south. 
2) Protect against repeated reveal call. This is just simply broken as of now.
3) Reveal is broken. Someone can commit hash in the previous round of the game,
and the keep revealing the same number without staking the bet in all the next rounds.
4) In the current implementaition the user that is last to reveal will pay for the whole findWinners
function. What happens if the transaction runs out of gas....

!!!!!!!!!!!! POTENTIAL BUGS !!!!!!!!!!!!!!!!
1) Ether being brought over to the next round due to rounding issues. 
Check that this is not the case. 

!!!!!!!!!! LOGICAL IMPROVEMENTS !!!!!!!!!!!!
1) Send money to HOME address only once the fees reach a certain amount.
2) I have this idea of introducing a system of rounds so our players get a receit.
They can then use this receit to look up info about their game.

!!!!!!!!!! PERFORMANCE INMPROVEMENTS !!!!!!!!!!!!!!!!
1) Change the body of findWinners to use an in memory array to find the winners. 
Then copy those winners to info.lastWinners

!!!!!!!!!! REFACTORING GOALS !!!!!!!!!!!!!!!!!
1) Get a goddamn constructor in here with a reasonable amount of arguments
2) Maybe put all the state related variable into a State struct? Work around that object. 
I started the idea....
3) Give normal names to everything. Also, change to cammelcase. 
Apparently, the lower_case_underscore is not that popular in Solidity...


*/

import "./Pausable.sol";
import "./GameHelper.sol";


contract Game is Pausable, GameHelper {
    
    enum GameState {COMMIT_STATE, REVEAL_STATE, PAYOUT_STATE}

    // This is the idea.
    struct State {
        GameState gameState;
        uint gameStateDebug;
        uint currNumberCommits;
        uint currNumberReveals;
        uint finalCommitBlock;
    }

    struct Config {
        uint REVEAL_PERIOD;
        uint MAX_PLAYERS;
        uint MIN_GUESS;
        uint MAX_GUESS;

        address FEE_ADDRESS;
        uint GAME_FEE_PERCENT;
        uint STAKE_SIZE;
    }

    struct GameInfo {
        address[] lastWinners;
        uint lastPrize;
    }

    Config internal config;
    State internal state;
    GameInfo internal info;

    constructor(uint maxp) public {
        owner = msg.sender;

        config.REVEAL_PERIOD = 5;
        config.GAME_FEE_PERCENT = 5;
        config.MAX_PLAYERS = maxp;
        config.MIN_GUESS = 0;
        config.MAX_GUESS = 100;
        config.STAKE_SIZE = 1 ether;
        config.FEE_ADDRESS = 0x2540099e9ed04aF369d557a40da2D8f9c2ab928D;

        state.gameState = GameState.COMMIT_STATE;
        state.currNumberCommits = 0;
        state.currNumberReveals = 0;
        state.finalCommitBlock = 0;

        info.lastPrize = 0;
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

    function getNumberOfWinners() public view returns(uint) {
        return info.lastWinners.length;
    }

    function getLastWinners(uint i) public view returns(address){
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


    // Commit/Reveal Protocol vars
    mapping (address => bytes32) public commits;
    mapping (address => uint) public game_data;
    address[] internal player_addrs;
    //address[] internal winners;

    // UI vars
    //address[] public last_winners;

    ////// DEBUG vars and debug functions
    uint public average23 = 0;

    function set_MAX_PLAYERS(uint new_val) public onlyOwner {
        config.MAX_PLAYERS = new_val;
    }

    // function is used to trigger a payout in a situation where somone
    // forgets to send the reveal.
    // function trigger_payout() public onlyOwner {
    //     require(game_state == GameState.REVEAL_STATE);

    //     // If the REVEAL_PERIOD blocks has gone by, while unfair, 
    //     // keep the money of nonrevealers, play the game with the
    //     // rest of the players.
    //     if(block.number > final_commit_block + REVEAL_PERIOD){
    //         find_winner();
    //     }
    // }

    // Reset the contract to the initial state
    function reset() public onlyOwner {  
        toCommitState();
        info.lastPrize = 0;
        delete info.lastWinners;
    }

    // Commit your guess. 
    event SuccesfulCommit(
        bytes32 hashed_commit
    );

    function commit(bytes32 hashedCommit) public payable whenNotPaused {
        
        //require(game_state == GameState.COMMIT_STATE);
        require(state.gameState == GameState.COMMIT_STATE);

        require(msg.value == config.STAKE_SIZE);

        commits[msg.sender] = hashedCommit;
        state.currNumberCommits++;

        // Notify the user that their bet reached us
        emit SuccesfulCommit(hashedCommit);

        // If we received the MAX_PLAYER number of commits, it is time for
        // us to change state.
        if (state.currNumberCommits == config.MAX_PLAYERS) {
            toRevealState();
        }
    }

    event SuccesfulReveal(
        string guess,
        string random
    );

    event DebugEvent(
        string error
    );
    function reveal(string guess, string random) public whenNotPaused {
        
        //require(game_state == GameState.REVEAL_STATE);
        emit DebugEvent("Entered reveal");
        require(state.gameState == GameState.REVEAL_STATE);
        
        // DEBUG: Need to make sure it throws if the guess is not integer
        uint guess_num = stringToUint(guess);
        
        require(guess_num >= config.MIN_GUESS && guess_num <= config.MAX_GUESS);

        // Check that the hashes match
        require(commits[msg.sender] == keccak256(guess, random));

        // When they do, we add the revealed guess to game data
        game_data[msg.sender] = guess_num;
        player_addrs.push(msg.sender);
        state.currNumberReveals++;

        emit SuccesfulReveal(guess, random);

        emit DebugEvent("All good here");

        if(state.currNumberReveals == config.MAX_PLAYERS){
            emit SuccesfulReveal(guess, random);
            find_winner();
        }
    }

    event DebugWinner(address addr, uint n);


    function find_winner() private {

        emit DebugWinner(1, player_addrs.length);
        
        // Calculate the 2/3 average
        uint guess_sum = 0;
        for(uint i = 0; i < player_addrs.length; i++){
            uint tmp = game_data[player_addrs[i]];
            guess_sum += tmp;
        }

        guess_sum = guess_sum * 10000;
        uint average = div(guess_sum, player_addrs.length);
        uint twothirdsavg = div(mul(average, 2), 3);
        twothirdsavg = twothirdsavg / 10000;

        //DEBUG
        average23 = twothirdsavg;

        address[] memory winners = new address[](config.MAX_PLAYERS);
        uint winIndex = 0;
        // We also flush the last list of winner
        delete info.lastWinners;


        // Find the guessers who are the closest to the 2/3 average
        uint min_diff = config.MAX_GUESS;
        uint cur_diff;
        for(i = 0; i < player_addrs.length; i++) {
            
            uint cur_guess = game_data[player_addrs[i]];

            // Find the difference between the guess and the average
            if(twothirdsavg > cur_guess){
                cur_diff = twothirdsavg - cur_guess;
            }
            else{
                cur_diff = cur_guess - twothirdsavg;
            }
            
            // If the difference is less than the smallest difference,
            // we delete all the winners and add the new candidate
            if(cur_diff < min_diff) {
                
                // NOT A BUG!
                //delete winners; // WTF you might ask? There is no necessity to delete elements.
                
                winIndex = 0;
                winners[winIndex] = player_addrs[i];
                winIndex++;

                min_diff = cur_diff;
            // Else, if the difference are the same, we add the candidate to the 
            // list of winners
            } else if(cur_diff == min_diff){
                winners[winIndex] = player_addrs[i];
                winIndex++;
            }
        }

         emit DebugWinner(2, player_addrs.length);

        // winrIndex here has the number of winner in our array
        require(winIndex > 0);

        // Lets pay ourselves some money
        uint gamefee = (address(this).balance/100) * config.GAME_FEE_PERCENT;
        config.FEE_ADDRESS.transfer(gamefee);

        // Split the rest equally among winners
        uint prize = address(this).balance/winIndex;
        for(i = 0; i < winIndex; i++){
            winners[i].transfer(prize); 
            info.lastWinners.push(winners[i]);
            //emit DebugWinner(winners[i], winners.length);
        }
        info.lastPrize = prize;

        // DEBUG: Make sure no ether is lost due to rounding. 

        // RESET STATE
        toCommitState();
    }

    // Call this funtion to get to COMMIT_STATE
    event DebugCommitState(uint last_win_l, uint win_l);

    function toCommitState() internal {

        state.gameState = GameState.COMMIT_STATE;
        state.gameStateDebug = 0;

        delete player_addrs;
        //delete info.lastWinners;
        
        //info.lastWinners = winners;

        //delete winners;
        
        state.currNumberCommits = 0;
        state.currNumberReveals = 0;

        //DebugCommitState(num_last_winners, winners.length);
    }

    // Call this function to get to REVEAL_STATE
    function toRevealState() internal {

        state.gameState = GameState.REVEAL_STATE;
        state.gameStateDebug = 1;
        state.finalCommitBlock = block.number;

    }

}