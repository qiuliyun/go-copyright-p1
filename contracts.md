
### ERC20.sol 

```
pragma solidity^0.4.20;

contract ERC20 {
    function totalSupply() constant returns (uint totalSupply);      
    function balanceOf(address _owner) constant returns (uint balance);
    function transfer(address _to, uint _value) returns (bool success);
    function transferFrom(address _from, address _to, uint _value) returns (bool success);
    function approve(address _spender, uint _value) returns (bool success);                  
    function allowance(address _owner, address _spender) constant returns (uint remaining);
    event Transfer(address indexed _from, address indexed _to, uint _value);                   
    event Approval(address indexed _owner, address indexed _spender, uint _value);
}
```


### pxCoin.sol 

```

pragma solidity^0.4.20;

import './ERC20.sol';

contract pxCoin is ERC20 {
    string public name = "pxcb";
    //token的符号
    string public symbol = "pxc";
    //基金会地址
    address public foundation;
    //issuer 
    address public issuer;
    //总发行量
    uint private _totalSupply  ;
    //账户余额
    mapping(address=>uint) _balance;
    //授权余额
    mapping(address=>mapping(address=>uint)) _allowance;
     
    constructor(uint totalSupply, address _owner) public {
        _totalSupply = totalSupply;
        foundation = _owner;
        _balance[foundation] = totalSupply * 20 / 100;
        issuer = msg.sender;
        _balance[issuer] = totalSupply * 80 / 100;
    }
    //查询总发行量
    function totalSupply() constant returns (uint totalSupply) {
        totalSupply = _totalSupply;
        return;
    }
    //查询余额
    function balanceOf(address _owner) constant returns (uint balance) {
        return _balance[_owner];
    }
    //转账处理
    function transfer(address _to, uint _value) returns (bool success) {
        // 余额充足+不难溢出 
        //assert( _balance[msg.sender] >= _value );
        if( _balance[msg.sender] >= _value &&
            address(0) != _to  &&
            _balance[_to] + _value > 0
        ) {
            // 一个增加，一个减少 
            _balance[msg.sender] -= _value;
            _balance[_to]  += _value;// if _to.val chao ji da > uint256
            emit Transfer(msg.sender, _to, _value);
            return true;
        }
        else {
            return false;
        }
        
        
    }
    //授权使用的转账
    function transferFrom(address _from, address _to, uint _value) returns (bool success) {
        if( address(0) != _to &&
            _balance[_to] + _value > 0 &&
            _allowance[_from][_to] >= _value &&
            _balance[_from] >= _value
        ) {
            _allowance[_from][_to] -= _value;
            _balance[_to] += _value;
            _balance[_from] -= _value;
            // dui zhang对账 
            //require(_balance[_to] + _balance[_from] = _allowance[_from][_to]);
            return true;
        }
        else {
            return false;
        }
    }
    //授权处理
    function approve(address _spender, uint _value) returns (bool success) {
        if( _balance[msg.sender] >= _value &&
            address(0) != _spender 
        ) {
            // 余额必须充足
            _allowance[msg.sender][_spender] = _value; 
            emit Approval(msg.sender, _spender, _value);
            
        }
        else {
            return false;
        }
    }
    // 查询授权的余额
    function allowance(address _owner, address _spender) constant returns (uint remaining) {
        return _allowance[_owner][_spender];
    }
    //返回合约地址
    function getAddr() public view returns (address) {
        return address(this);
    }
    
}
```


### SafeMath.sol 

```
pragma solidity ^0.4.24;

/**
 * @dev Math operations with safety checks that throw on error. This contract is based
 * on the source code at https://goo.gl/iyQsmU.
 */
library SafeMath {

  /**
   * @dev Multiplies two numbers, throws on overflow.
   * @param _a Factor number.
   * @param _b Factor number.
   */
  function mul(
    uint256 _a,
    uint256 _b
  )
    internal
    pure
    returns (uint256)
  {
    if (_a == 0) {
      return 0;
    }
    uint256 c = _a * _b;
    assert(c / _a == _b);
    return c;
  }

  /**
   * @dev Integer division of two numbers, truncating the quotient.
   * @param _a Dividend number.
   * @param _b Divisor number.
   */
  function div(
    uint256 _a,
    uint256 _b
  )
    internal
    pure
    returns (uint256)
  {
    uint256 c = _a / _b;
    // assert(b > 0); // Solidity automatically throws when dividing by 0
    // assert(a == b * c + a % b); // There is no case in which this doesn't hold
    return c;
  }

  /**
   * @dev Substracts two numbers, throws on overflow (i.e. if subtrahend is greater than minuend).
   * @param _a Minuend number.
   * @param _b Subtrahend number.
   */
  function sub(
    uint256 _a,
    uint256 _b
  )
    internal
    pure
    returns (uint256)
  {
    assert(_b <= _a);
    return _a - _b;
  }

  /**
   * @dev Adds two numbers, throws on overflow.
   * @param _a Number.
   * @param _b Number.
   */
  function add(
    uint256 _a,
    uint256 _b
  )
    internal
    pure
    returns (uint256)
  {
    uint256 c = _a + _b;
    assert(c >= _a);
    return c;
  }

}
```

### AddressUtils.sol


```
pragma solidity ^0.4.24;

/**
 * @dev Utility library of inline functions on addresses.
 */
library AddressUtils {

  /**
   * @dev Returns whether the target address is a contract.
   * @param _addr Address to check.
   */
  function isContract(
    address _addr
  )
    internal
    view
    returns (bool)
  {
    uint256 size;

    /**
     * XXX Currently there is no better way to check if there is a contract in an address than to
     * check the size of the code at that address.
     * See https://ethereum.stackexchange.com/a/14016/36603 for more details about how this works.
     * TODO: Check this again before the Serenity release, because all addresses will be
     * contracts then.
     */
    assembly { size := extcodesize(_addr) } // solium-disable-line security/no-inline-assembly
    return size > 0;
  }

}

```


### ERC721TokenReceiver.sol 

```
pragma solidity ^0.4.24;

/**
 * @dev ERC-721 interface for accepting safe transfers. See https://goo.gl/pc9yoS.
 */
interface ERC721TokenReceiver {

  /**
   * @dev Handle the receipt of a NFT. The ERC721 smart contract calls this function on the
   * recipient after a `transfer`. This function MAY throw to revert and reject the transfer. Return
   * of other than the magic value MUST result in the transaction being reverted.
   * Returns `bytes4(keccak256("onERC721Received(address,address,uint256,bytes)"))` unless throwing.
   * @notice The contract address is always the message sender. A wallet/broker/auction application
   * MUST implement the wallet interface if it will accept safe transfers.
   * @param _operator The address which called `safeTransferFrom` function.
   * @param _from The address which previously owned the token.
   * @param _tokenId The NFT identifier which is being transferred.
   * @param _data Additional data with no specified format.
   */
  function onERC721Received(
    address _operator,
    address _from,
    uint256 _tokenId,
    bytes _data
  )
    external
    returns(bytes4);
    
}

```


### ERC721.sol

```
pragma solidity^0.4.24;

interface ERC721  {
 
    event Transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId);
    event Approval(address indexed _owner, address indexed _approved, uint256 indexed _tokenId);
    event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved);
    function balanceOf(address _owner) external view returns (uint256);
    function ownerOf(uint256 _tokenId) external view returns (address);
    function safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes data) external;
    function safeTransferFrom(address _from, address _to, uint256 _tokenId) external;
    function transferFrom(address _from, address _to, uint256 _tokenId) external;
    function approve(address _approved, uint256 _tokenId) external;
    function setApprovalForAll(address _operator, bool _approved) external;
    function getApproved(uint256 _tokenId) external view returns (address);
    function isApprovedForAll(address _owner, address _operator) external view returns (bool);
}
```


### pxAsset.sol 
```
pragma solidity^0.4.24;

import './ERC721.sol';
import './SafeMath.sol';
import './AddressUtils.sol';
import './ERC721TokenReceiver.sol';
import './pxCoin.sol';

contract pxAsset is ERC721 {
    
    using SafeMath for uint256;
    using AddressUtils for address;
    address public foundation;
    mapping(address=>uint) _ownerTokenCount;
    mapping(uint=>address) _tokenOwner;
    mapping(uint=>address) _tokenApprovals;
    mapping(address=>mapping(address=>bool)) _operatorApprovals;
    pxCoin pxcoin;
    
    //a tokenId is a Asset,a tokenId must belong to a address
    struct Asset {
        string contentHash;
        uint price;
        uint weight;
        string metaData;
        uint voteCount;
    }
    
    //all assets here,get a new tokenId when push a new asset 
    Asset[] public assets;
    
    /**
   * @dev Magic value of a smart contract that can recieve NFT.
   * Equal to: bytes4(keccak256("onERC721Received(address,address,uint256,bytes)")).
   */
    bytes4 constant MAGIC_ON_ERC721_RECEIVED = 0x150b7a02;
    
    constructor() public {
        foundation = msg.sender;
        pxcoin = new pxCoin(1000000000, msg.sender);
    }
    
    modifier onlyOwner() {
        require( msg.sender == foundation );
        _;
    }
    modifier canOperate(uint _tokenId) {
        address tokenOwner = _tokenOwner[_tokenId];
        require( tokenOwner == msg.sender ||
                 _operatorApprovals[tokenOwner][msg.sender] );
        _;
    }
    
    
    
    modifier canTransfer(uint _tokenId) {
        address tokenOwner = _tokenOwner[_tokenId];
        require( tokenOwner == msg.sender ||
                 _operatorApprovals[tokenOwner][msg.sender] ||
                 _getApproved(_tokenId) == msg.sender );
        _;
    }
    
    modifier validToken( uint _tokenId ) {
        require( _tokenOwner[_tokenId] != address(0) );
        _;
    }
    function _getApproved(uint256 _tokenId) private validToken(_tokenId) view returns (address) {
        return _tokenApprovals[_tokenId];
    }
    function balanceOf(address _owner) external view returns (uint256) {
        require( _owner != address(0) );
        return _ownerTokenCount[_owner];
    }
    function ownerOf(uint256 _tokenId) external view returns (address) {
        address owner = _tokenOwner[_tokenId];
        require( address(0) != owner );
        return owner;
    }
    function removeToken(address _from, uint256 _tokenId) internal {
        require( _tokenOwner[_tokenId] == _from );
        assert( _ownerTokenCount[_from] > 0 );
        _ownerTokenCount[_from] = _ownerTokenCount[_from] - 1;
        delete _tokenOwner[_tokenId];
    }
    function addToken(address _to, uint256 _tokenId) internal {
        require( _tokenOwner[_tokenId] == address(0) );
        _tokenOwner[_tokenId] = _to;
        _ownerTokenCount[_to] = _ownerTokenCount[_to].add(1);
    }
    
    function clearApproval(uint256 _tokenId) private {
        if( _tokenApprovals[_tokenId] != address(0) ) {
            delete _tokenApprovals[_tokenId];
        }
    }
    
    function _transfer(address _to,uint256 _tokenId) private {
        address _from = _tokenOwner[_tokenId];
        clearApproval(_tokenId);
        
        removeToken(_from, _tokenId);
        addToken(_to, _tokenId);
        
        emit Transfer(_from, _to, _tokenId);
    }
    event logSafeTransfer(bool isCont, bytes4 retval);
    function _safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes _data)  internal canTransfer(_tokenId) validToken(_tokenId) {
        address tokenOwner = _tokenOwner[_tokenId];
        require(tokenOwner == _from);
        require(_to != address(0));

        _transfer(_to, _tokenId);
        if (_to.isContract()) 
        {
          bytes4 retval = ERC721TokenReceiver(_to).onERC721Received(msg.sender, _from, _tokenId, _data);
          require(retval == MAGIC_ON_ERC721_RECEIVED);
        }
        //logSafeTransfer(_to.isContract(), retval);
    }
    function safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes _data) external {
        _safeTransferFrom(_from, _to, _tokenId, _data);
    }
    function safeTransferFrom(address _from, address _to, uint256 _tokenId) external {
        _safeTransferFrom(_from, _to, _tokenId, "");
    }
    function transferFrom(address _from, address _to, uint256 _tokenId) external canTransfer(_tokenId) validToken(_tokenId)  {
        address tokenOwner = _tokenOwner[_tokenId];
        require(tokenOwner == _from);
        require(_to != address(0));
        _transfer(_to, _tokenId);
    }
    function approve(address _approved, uint256 _tokenId) external canOperate(_tokenId) validToken(_tokenId) {
        address tokenOwner = _tokenOwner[_tokenId];
        require( tokenOwner != _approved );
        
        _tokenApprovals[_tokenId] = _approved;
        emit Approval(tokenOwner, _approved, _tokenId);
        
    }
    function setApprovalForAll(address _operator, bool _approved) external {
        require( _operator != address(0) );
        _operatorApprovals[msg.sender][_operator] = _approved;
        emit ApprovalForAll(msg.sender, _operator, _approved);
    }
    function getApproved(uint256 _tokenId) external view returns (address) {
        return _getApproved(_tokenId);
    }
    function isApprovedForAll(address _owner, address _operator) external view returns (bool) {
        require( address(0) != _owner );
        require( address(0) != _operator );
        return _operatorApprovals[_owner][_operator];
    }
    
    function _newAsset(string _contentHash, uint _price, uint _weight, string _data) internal returns (uint) {
        Asset memory a = Asset(_contentHash, _price, _weight, _data, 0);
        uint tokenId = assets.push(a) - 1;
        return tokenId;
    }
    
    function mint(string _contentHash, uint _price, string _data) external {
        uint tokenId = _newAsset(_contentHash, _price, 100, _data);
        //add mapping 
        _ownerTokenCount[msg.sender] = _ownerTokenCount[msg.sender].add(1);
        _tokenOwner[tokenId] = msg.sender;
        //pxc transfer, transfer()' msg.sender is local contract address 
        pxcoin.transfer(msg.sender, 100);
    }
    
    function splitAsset(uint _tokenId, uint _weight, address _buyer) onlyOwner validToken(_tokenId) external returns(uint tokenId) {
        
        assert( _weight < 100 );//if _weight == 100, should be call transferFrom
        require( address(0) != _buyer );
        Asset  a = assets[_tokenId];
        require( a.weight > _weight );//if _weight == weight, should be call transferFrom
        //a.weight = _weight;
        tokenId = assets.push(a) - 1;
        a = assets[tokenId];
        a.weight = _weight;
        addToken(_buyer, tokenId);
        a = assets[_tokenId];
        a.weight = a.weight.sub(_weight);
        return tokenId;
    }
    
    function getPXCAddr() view returns (address) {
        return address(pxcoin);
    }
    
    function getPXCBalance(address _owner) view returns( uint ) {
        return pxcoin.balanceOf(_owner);
    } 
    
}
```
