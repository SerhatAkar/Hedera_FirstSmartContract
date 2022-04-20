// SPDX-License-Identifier: GPL-3.0

pragma solidity ^0.8.13;

contract LookupContract {

    mapping(string => uint32) public myDirectory;

    constructor(string memory _name, uint32 _mobileNumber) public {
        myDirectory[_name] = _mobileNumber;
    }

    function setMobileNumber(string memory _name, uint32 _mobileNumber) public {
         myDirectory[_name] = _mobileNumber;
    }
    
    function getMobileNumber(string memory _name) public view returns(uint32) {
        return myDirectory[_name];
    }
}