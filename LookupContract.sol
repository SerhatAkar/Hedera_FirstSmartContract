pragma solidity ^0.8.13;

contract LookupContract {

    mapping(string => uint) public myDirectory;

    constructor(string memory _name, uint _mobileNumber) public {
        myDirectory[_name] = _mobileNumber;
    }

    function setMobileNumber(string memory _name, uint _mobileNumber) public {
         myDirectory[_name] = _mobileNumber;
    }
    
    function getMobileNumber(string memory _name) public view returns(uint) {
        return myDirectory[_name];
    }
}