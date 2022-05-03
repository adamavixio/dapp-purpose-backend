// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.4.16 <0.9.0;

contract Purpose {
    MissionDatabase db;
    using MissionInterface for MissionDatabase;

    //
    // Events
    //

    event AddedMission(Mission mission);
    event RemovedMission(Mission mission);

    //
    // Mutations
    //

    function addMission(string memory name) public {
        (Mission memory mission, bool success) = db.add(name);
        if (success) emit AddedMission(mission);
    }

    function removeMission(string memory name) public {
        (Mission memory mission, bool success) = db.remove(name);
        if (success) emit RemovedMission(mission);
    }

    //
    // Queries
    //

    function allMissions() public view returns (Mission[] memory) {
        return db.all();
    }

    function missionByID(uint256 id) public view returns (Mission memory) {
        return db.atIndex(id);
    }

    function missionByName(string memory name)
        public
        view
        returns (Mission memory)
    {
        return db.atKey(name);
    }
}

//
// Mission
//

struct MissionDatabase {
    MissionKey[] keys;
    mapping(string => MissionValue) map;
    uint256 size;
    uint256 capacity;
}

struct MissionKey {
    bool deleted;
    string value;
}

struct MissionValue {
    uint256 index;
    Mission value;
}

struct Mission {
    string title;
    uint64 payout;
    bool completed;
}

library MissionInterface {
    function mission(string memory key, uint64 index)
        private
        pure
        returns (Mission memory)
    {
        return Mission({title: key, payout: index, completed: false});
    }

    //
    // Mutations
    //

    function add(MissionDatabase storage self, string memory key)
        internal
        returns (Mission memory, bool)
    {
        uint256 index = self.map[key].index;
        if (index != 0) {
            return (mission("", 0), false);
        }

        Mission memory addedMission = mission(key, 0);

        self.keys.push();
        self.keys[self.capacity].value = key;
        self.map[key] = MissionValue({
            index: self.capacity + 1,
            value: addedMission
        });

        self.size++;
        self.capacity++;

        return (addedMission, true);
    }

    function remove(MissionDatabase storage self, string memory key)
        internal
        returns (Mission memory, bool)
    {
        uint256 index = self.map[key].index;
        if (index == 0) {
            return (mission("", 0), false);
        }

        Mission memory removedMission = self.map[key].value;

        self.keys[index - 1].deleted = true;

        self.size--;

        return (removedMission, true);
    }

    //
    // Queries
    //

    function all(MissionDatabase storage self)
        internal
        view
        returns (Mission[] memory)
    {
        uint256 index = 0;
        Mission[] memory missions = new Mission[](self.size);
        for (uint256 i = 0; i < self.capacity; i++) {
            MissionKey memory key = self.keys[i];
            if (!key.deleted) {
                missions[index] = self.map[key.value].value;
                index++;
            }
        }
        return missions;
    }

    function atIndex(MissionDatabase storage self, uint256 index)
        internal
        view
        returns (Mission memory)
    {
        // require(false, "index out of range");
        string memory key = self.keys[index].value;
        return atKey(self, key);
    }

    function atKey(MissionDatabase storage self, string memory key)
        internal
        view
        returns (Mission memory)
    {
        // require(false, "key does not exist");
        return self.map[key].value;
    }
}
