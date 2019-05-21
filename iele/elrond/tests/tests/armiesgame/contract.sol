pragma solidity ^0.5.7;

contract ArmiesGame {

    event NewArmyEvent(uint armyId, string name, address owner);
    event SoldierLevelUpEvent(string status, uint soldierId, uint newLevel);
    event SoldierRenameEvent(uint soldierId, string newName);
    event NewSoldierEvent(string status, uint soldierId);
    event BattleEvent(string status, uint32 date, address attacker, address opponent, uint gain);
    event TopUpEvent(string status, uint points);

    string defaultSoldierName = "Soldier";
    uint defaultSoldierDna = 1000000000000000;
    address payable owner;
    
    // coolDown time
    uint cooldownTime = 1 minutes;
    
    // fees
    uint levelUpFee = 100; // 100 x level number
    uint newSoldierFee = 500;
    uint attackFee = 100;

    // points
    uint newSoldierPoints = 100;
    uint levelUpSoldierPoints = 10; // 10 x level number 
    uint attackWinPoints = 50; 

    struct Army {
        string name;
        address owner;
        uint points;
        uint gold;
        uint[] soldiersList;
    }

    struct Soldier {
        string name;
        address owner;
        uint dna;
        uint32 level;
        uint32 readyTime;
    }

    Army[] public armies;
    Soldier[] public soldiers;

    mapping (address => Army) ownerToArmy;

    constructor() public {
        owner = msg.sender;
    }

    modifier aboveLevel(uint _level, uint _soldierId) {
        require(soldiers[_soldierId].level >= _level);
        _;
    }

    modifier onlyOwnerOf(uint _soldierId) {
        require(msg.sender == soldiers[_soldierId].owner);
        _;
    }

    function topUp(uint val) external payable {
        //get eth and gold values based on parameter
        uint topUpFee = 3 ether;
        uint topUpGold = 10000; 
        if(val == 0.1) {
            topUpFee = 0.1 ether;
            topUpGold = 150;    
        } else if(val == 0.5) {
            topUpFee = 0.5 ether;
            topUpGold = 1000;
        } else if(val == 1) {
            topUpFee = 1 ether;
            topUpGold = 2500;
        }
        
        // check to see if the fee was payed
        require(msg.value >= topUpFee);

        // add points
        ownerToArmy[msg.sender].gold += topUpGold;

        // emit topUp event
        emit TopUpEvent("success", topUpGold);
    }

    function transferBalance() external payable{
        owner.transfer(address(this).balance);
    }

    function getBalance() public view returns(uint256) {
        return address(this).balance;
    }

    function addArmy(string memory _name) public returns(uint) {
        // check if tis is the first army of this owner
        require(ownerToArmy[msg.sender].owner == address(0));

        // create new army and add it to the array 
        uint armyId = armies.push(Army(_name, msg.sender, 0, 1000, new uint[](0))) - 1;

        // generat new foldier for this new army and asign it to the army
        uint newSoldierId = _generateNewSoldier(msg.sender);
        armies[armyId].soldiersList.push(newSoldierId);

        // mark owner that he has an army
        ownerToArmy[msg.sender] = armies[armyId];

        emit NewArmyEvent(armyId, _name, msg.sender);
        return armyId;
    }

    function getArmyDetails(address _address) public view returns(string memory, uint, uint, address, uint[] memory) {
        uint[] memory localSoldiersList = new uint[](ownerToArmy[_address].soldiersList.length);

        for(uint i = 0; i < ownerToArmy[_address].soldiersList.length; i++) {
            localSoldiersList[i] = ownerToArmy[_address].soldiersList[i];
        }

        return (ownerToArmy[_address].name, ownerToArmy[_address].points, ownerToArmy[_address].gold, _address, localSoldiersList);
    }

    function _generateNewSoldier(address _owner) internal returns(uint) {
        // create new soldier
        Soldier memory soldier = Soldier(defaultSoldierName, _owner, defaultSoldierDna, 1, uint32(now + cooldownTime));

        // add it to the soldiers array 
        uint soldierId = soldiers.push(soldier) - 1;

        // emit NewArmyEvent(soldierId, _name, msg.sender);
        return soldierId;
    }

    function levelUp(uint _soldierId) external {
        // check to see if the army has enough gold
        if(ownerToArmy[msg.sender].gold < (levelUpFee * soldiers[_soldierId].level)) {
            emit SoldierLevelUpEvent("neg", _soldierId, soldiers[_soldierId].level + 1);
            return;
        }

        // deduct gold from the balance
        ownerToArmy[msg.sender].gold -= levelUpFee * soldiers[_soldierId].level;

        // add points to the army
        ownerToArmy[msg.sender].points += soldiers[_soldierId].level * levelUpSoldierPoints;

        // change the level
        soldiers[_soldierId].level = soldiers[_soldierId].level + 1;

        // add cool down time
        soldiers[_soldierId].readyTime = uint32(now + cooldownTime);

        // emit level up event
        emit SoldierLevelUpEvent("success", _soldierId, soldiers[_soldierId].level);
    }

    function changeName(uint _soldierId, string calldata _newName) external aboveLevel(2, _soldierId) onlyOwnerOf(_soldierId) {
        soldiers[_soldierId].name = _newName;

        // add cool down time
        soldiers[_soldierId].readyTime = uint32(now + cooldownTime);

        // emit rename event
        emit SoldierRenameEvent(_soldierId, _newName);
    }

    function buyNewSoldier() external {
        // check to see if the army has enough gold
        if(ownerToArmy[msg.sender].gold < newSoldierFee) {
            emit NewSoldierEvent("neg", 0);
            return;
        }

        // deduct gold from the balance
        ownerToArmy[msg.sender].gold -= newSoldierFee;

        // check if owner has an army already
        require(ownerToArmy[msg.sender].owner != address(0));

        // generate new soldier
        uint newSoldierId = _generateNewSoldier(msg.sender);

        // add soldier id in hes army soldiers array
        ownerToArmy[msg.sender].soldiersList.push(newSoldierId);

        // add points to the army
        ownerToArmy[msg.sender].points += newSoldierPoints;

        // emit new soldier event
        emit NewSoldierEvent("success", newSoldierId);
    }

    function attack() external {
        // check to see if the army has enough gold
        if(ownerToArmy[msg.sender].gold < attackFee) {
            emit BattleEvent("neg", uint32(now), msg.sender, msg.sender, 0);
            return;
        }

        if(armies.length < 2) {
            emit BattleEvent("canceled", uint32(now), msg.sender, msg.sender, 0);
            return;
        }

        // get attacker army level
        uint attackerLevel = getArmyLevel(msg.sender);
        if(attackerLevel == 0) {
            // unable to attack as, all attacker soldiers ar in coolDown 
            emit BattleEvent("cool_down", uint32(now), msg.sender, msg.sender, 0);
            return;
        }

        uint opponentArmyId = 0;
        string memory status = "canceled";

        // get random opponent army (try max 50 times to find)
        for(uint i = 0; i < 50; i++) {
            opponentArmyId = randMod(armies.length);
            if(address(armies[opponentArmyId].owner) != address(msg.sender)) {
                status = "available_army";
                break;
            }
        }
        
        // if there is no available army, return cancel status
        if(compareStrings(status, "canceled")) {
            emit BattleEvent("canceled", uint32(now), msg.sender, msg.sender, 0);
            return;
        }

        // deduct gold from the balance
        ownerToArmy[msg.sender].gold -= attackFee;

        address opponentAddr = armies[opponentArmyId].owner;
        uint opponentLevel = getArmyLevel(opponentAddr);

        uint winPoints = 0;
        if(opponentLevel > 0) {
            winPoints = randMod(attackerLevel + opponentLevel);
            if(winPoints < attackerLevel) {
                // attacker win
    
                // add points to the attacker army
                ownerToArmy[msg.sender].points += attackWinPoints;
                
                // add cooldownTime for all opponent soldiers
                coolDownArmy(ownerToArmy[opponentAddr].soldiersList);
                
                // get half of the opponents gold
                uint gain = ownerToArmy[opponentAddr].gold / 2;
                ownerToArmy[msg.sender].gold += gain;
                ownerToArmy[opponentAddr].gold -= gain;

                // emit BattleEvent
                emit BattleEvent("win", uint32(now), msg.sender, opponentAddr, gain);
                return;
            }
        }

        // opponent win

        // add points to the opponent army
        ownerToArmy[opponentAddr].points += attackWinPoints;

        // add cooldownTime for all attacker soldiers
        coolDownArmy(ownerToArmy[msg.sender].soldiersList);

        // emit BattleEvent
        emit BattleEvent("loose", uint32(now), msg.sender, opponentAddr, 0);
    }

    function getArmyLevel(address _armyAddress) internal view returns(uint) {
        uint level = 0;
        for(uint i=0; i< ownerToArmy[_armyAddress].soldiersList.length; i++) {
            if(soldiers[ownerToArmy[_armyAddress].soldiersList[i]].readyTime < now)
                level += soldiers[ownerToArmy[_armyAddress].soldiersList[i]].level;
        }
        return level;
    }

    function coolDownArmy(uint[] memory _soldiersList) internal {
        for(uint i=0; i< _soldiersList.length; i++) {
            soldiers[_soldiersList[i]].readyTime = uint32(now + cooldownTime);
        }
    }
    
    function randMod(uint _modulus) internal view returns(uint) {
        return now % _modulus;
//        return (block.timestamp + block.difficulty ) % _modulus;
    }

    function compareStrings (string memory firstString, string memory secondString) internal pure returns (bool) {
        return uint(keccak256(abi.encodePacked(firstString))) == uint(keccak256(abi.encodePacked(secondString)));
    }
}
