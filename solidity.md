# copyright.sol
pragma solidity>=0.4.22 <0.6.0;
import './ERC721.sol';
import './AddressUtils.sol';
import './ERC721TokenReceiver.sol';
import './SafeMath.sol';
contract copyright is ERC721 {

    using AddressUtils for address; 
    using SafeMath for uint256;
    bytes4 constant MAGIC_ON_ERC721_RECEIVED =0x150b7a02;
    
    //owner->tokenId
    mapping(address=>uint256) _ownerTokenCount;
    //tokenId->owne
    mapping(uint256=>address) _tokenOwner;
    //tokenId->Approval
    mapping(uint256=>address) _tokenApprovals;
    mapping(address=>mapping(address=>bool)) _operatorApprovals;
    
    //shuzi zichan dindyi
    struct Asset{
        bytes32 contentHash;//zichan hash
        string copyrightTran;//banquan jiaoyi
        string tran;//jiaoyi
        string name;//mingzi
    }
    
    Asset[] public assets;
    
    constructor()public {}
    
        //can zhuan zhang???
    modifier canTransfer(uint _tokenId){
        address tokenOwner=_tokenOwner[_tokenId];
        require(msg.sender == tokenOwner 
        || msg.sender == _getApproved(_tokenId)
        || _operatorApprovals[tokenOwner][msg.sender]
        );
        _;
    }
    
    //can caozuo???
    modifier canOperate(uint256 _tokenId){
        address tokenOwner=_tokenOwner[_tokenId];
        require(msg.sender == tokenOwner
            ||_operatorApprovals[tokenOwner][msg.sender]
            );
            _;
    }
    
    //
    modifier validToken(uint256 _tokenId){
        require(_tokenOwner[_tokenId]!=address(0));
        _;
    }
    
    //721
    function balanceOf(address _owner) external view returns (uint256){
        require(address(0)!=_owner);
        return  _ownerTokenCount[_owner];
    }
    
    //721  token de zhuren
    function ownerOf(uint256 _tokenId) external view returns (address){
        address owner=_tokenOwner[_tokenId];
        require(address(0)!=owner);
        return owner;
    }
    
    //siyou zhuanzhang fenzhuang
    function _transfer(address _to,uint _tokenId)private{
        address tokenOwner=_tokenOwner[_tokenId];
        clearApproval(_tokenId);//*
        removeToken(tokenOwner,_tokenId);//*
        addToken(_to,_tokenId);//*
        emit Transfer(tokenOwner,_to,_tokenId);
    }  
    
    //siyou qingchu shouquan fenzhuang
    function clearApproval(uint256 _tokenId)private{
        if(_tokenApprovals[_tokenId] != address(0)){
            delete(_tokenApprovals[_tokenId]);
        }
    }
    
    //siyou qingchu fenzhuang
    function removeToken(address _from,uint _tokenId)private{
        require(_tokenOwner[_tokenId] == _from);
        assert(_ownerTokenCount[_from]>0);
        _ownerTokenCount[_from]=_ownerTokenCount[_from]-1;
        delete(_tokenOwner[_tokenId]);
    }

    //siyou tianjia fenzhuang
    function addToken(address _to,uint _tokenId)private{
        require(_tokenOwner[_tokenId] ==address(0));
        _tokenOwner[_tokenId]=_to;
        _ownerTokenCount[_to]=_ownerTokenCount[_to].add(1);
    }
    
    //syou safe tran
    function _safeTransferFrom(address _from, address _to, uint256 _tokenId,bytes memory _data) private canTransfer(_tokenId) validToken(_tokenId){
        address tokenOwner=_tokenOwner[_tokenId];
        require(_to != address(0));
        _transfer(_to,_tokenId);
        //safe or no???
        if(_to.isContract()){
            bytes4 retval=ERC721TokenReceiver(_to).onERC721Received(msg.sender,_from,_tokenId,_data);
            require(retval == MAGIC_ON_ERC721_RECEIVED);
        }
        
    }
    
    //721   safe zhuan zhang
    function safeTransferFrom(address _from, address _to, uint256 _tokenId,bytes calldata _data) external payable{
        _safeTransferFrom(_from,_to,_tokenId,_data);
    }
    //721
    function safeTransferFrom(address _from, address _to, uint256 _tokenId) external payable{
        _safeTransferFrom(_from,_to,_tokenId,"");
    }
    
    //721 trans
    function transferFrom(address _from, address _to, uint256 _tokenId)canTransfer(_tokenId) validToken(_tokenId) external payable{
        address tokenOwner=_tokenOwner[_tokenId];
        require(address(0)!=_to);
        _transfer(_to,_tokenId);
    }
    
    //721   shou quan
    function approve(address _approved, uint256 _tokenId)validToken(_tokenId) external{
        address tokenOwner=_tokenOwner[_tokenId];
        require(_approved != address(0));
        _tokenApprovals[_tokenId]=_approved;
    }
    
    //siyou get
    function _getApproved(uint256 _tokenId)validToken(_tokenId) private view returns (address){
        return _tokenApprovals[_tokenId];
    }
    
    //721   all shou quan 
    function setApprovalForAll(address _operator, bool _approved) external{
        require(_operator != address(0));
        require(_ownerTokenCount[_operator]>0);
        _operatorApprovals[msg.sender][_operator]=_approved;
    }
    //721 get
    function getApproved(uint256 _tokenId) external view returns (address){
        return _getApproved(_tokenId);
    }
    
    
    //721   is all shouquan
    function isApprovedForAll(address _owner, address _operator) external view returns (bool){
        return _operatorApprovals[_owner][_operator];
    }
    
    event newAsset(bytes32 _hash,address _owner,uint256 _tokenId);
    
    //siyou up
    function _newAsset(bytes32  _hash,string memory _data)private returns(uint256){
        Asset memory a=Asset(_hash,"","",_data);
        uint256 tokenId=assets.push(a)-1;
        emit newAsset(_hash,msg.sender,tokenId);
        return tokenId;
    }
    
    //siyou wa kuang
    function mint(bytes32  _hash,string calldata _data)external{
        uint256 tokenId=_newAsset(_hash,_data);
        _ownerTokenCount[msg.sender]=_ownerTokenCount[msg.sender].add(1);
        _tokenOwner[tokenId]=msg.sender;
    }
    
    
    //siyou lianjie
    function strConcat(string memory  _a, string memory  _b) private returns (string memory){
        bytes memory _ba = bytes(_a);
        bytes memory _bb = bytes(_b);
        string memory ret = new string(_ba.length + _bb.length);
        bytes memory bret = bytes(ret);
        uint k = 0;
        for (uint i = 0; i < _ba.length; i++)bret[k++] = _ba[i];
        for (uint i = 0; i < _bb.length; i++) bret[k++] = _bb[i];
        return string(ret);
   }
   
   //siyou banquan jiaoyi jilu
   function copyRecode(string memory _recode,uint256 _tokenId)public payable returns(string memory){
        string memory yuan=assets[_tokenId].copyrightTran;
        string memory xin=strConcat(yuan,_recode);
        assets[_tokenId].copyrightTran=xin;
        return xin;
   }
   
   //siyou  zichan jiaoyi jilu
   function assRecode(string memory _tran,uint256 _tokenId)public payable returns(string memory){
        string memory yuan=assets[_tokenId].tran;
        string memory xin=strConcat(yuan,_tran);
        assets[_tokenId].tran=xin;
        return xin;
   } 
   
   //siyou huoqu
   function getCopyrecode(uint256 _tokenId)public view returns(string memory){
       return assets[_tokenId].copyrightTran;
   }
   function getAssRecode(uint256 _tokenId)public view returns(string memory){
       return assets[_tokenId].tran;
   }
   
}  
# ERC721.sol
pragma solidity>=0.4.22 <0.6.0;
interface ERC721  {
    event Transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId);
    event Approval(address indexed _owner, address indexed _approved, uint256 indexed _tokenId);
    event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved);
    //_owner de token shuliang
    function balanceOf(address _owner) external view returns (uint256);
    //token de zhuren
    function ownerOf(uint256 _tokenId) external view returns (address);
    //safe zhuan zhang
    function safeTransferFrom(address _from, address _to, uint256 _tokenId,bytes calldata _data) external payable;
    function safeTransferFrom(address _from, address _to, uint256 _tokenId) external payable;
    //zhuan zhang
    function transferFrom(address _from, address _to, uint256 _tokenId) external payable;
    //shou quan
    function approve(address _approved, uint256 _tokenId) external;
    //all shou quan 
    function setApprovalForAll(address _operator, bool _approved) external;
    function getApproved(uint256 _tokenId) external view returns (address);
    function isApprovedForAll(address _owner, address _operator) external view returns (bool);
}
# ERC721TokenReceiver.sol
pragma solidity>=0.4.22 <0.6.0;
/** 
* @dev ERC-721 interface for accepting safe transfers. See https://goo.gl/pc9yoS. 
*/ 
interface ERC721TokenReceiver {
  /**   
  * @dev Handle the receipt of a NFT. The ERC721 smart contract calls this function on the   
  * recipient after a `transfer`. This function MAY throw to revert and reject the transfer. Return   
  * of other than the magic value MUST result in the transaction being reverted.  
  * Returns `bytes4(keccak256("onERC721Received(address,address,uint256,bytes)" ))` unless throwing.   
  * @notice The contract address is always the message sender. A wallet/broker/auction application  
  * MUST implement the wallet interface if it will accept safe transfers.  
  * @param _operator The address which called `safeTransferFrom` function.   
  * @param _from The address which previously owned the token.   
  * @param _tokenId The NFT identifier which is being transferred.  
  * @param _data Additional data with no specified format.   
  */  
  function onERC721Received(    address _operator,    address _from,    uint256 _tokenId,    bytes calldata _data  )    external    returns(bytes4);
  }

# AddressUtils.sol
/** * 
@dev Utility library of inline functions on addresses. 
*/ 
library AddressUtils {
  /**  
  * @dev Returns whether the target address is a contract.   * @param _addr Address to check.  
  */
  function isContract(    address _addr  )    internal    view    returns (bool)  {   
      uint256 size;
    /**     
    * XXX Currently there is no better way to check if there is a contract in an address than to     
    * check the size of the code at that address.     
    * See https://ethereum.stackexchange.com/a/14016/36603 for more details about how this works.     
    * TODO: Check this again before the Serenity release, because all addresses will be     
    * contracts then.     
    */    
    assembly { size := extcodesize(_addr) }
    // solium-disable-line security/no-inline-assembly    
    return size > 0;  }
}
# SafeMath.sol
pragma solidity>=0.4.22 <0.6.0;
/**
/**
* @dev Math operations with safety checks that throw on error. This contract is based 
* on the source code at https://goo.gl/iyQsmU.
*/ 
library SafeMath {
  /**   * @dev Multiplies two numbers, throws on overflow.  
  * @param _a Factor number.   * @param _b Factor number.  
  */ 
  function mul(    uint256 _a,    uint256 _b  )    internal    pure    returns (uint256)  {    
      if (_a == 0) {      return 0;
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
  function div(    uint256 _a,    uint256 _b  )    internal    pure    returns (uint256)  {   
      uint256 c = _a / _b;   
      // assert(b > 0);
      // Solidity automatically throws when dividing by 0    
      // assert(a == b * c + a % b);
      // There is no case in which this doesn't hold    return c;
      }
  /**  
  * @dev Substracts two numbers, throws on overflow (i.e. if subtrahend is greater than minuend).
  * @param _a Minuend number.
  * @param _b Subtrahend number.
  */
  function sub(    uint256 _a,    uint256 _b  )    internal    pure    returns (uint256)  {
      assert(_b <= _a);
      return _a - _b;
      }
  /**  
  * @dev Adds two numbers, throws on overflow. 
  * @param _a Number.   * @param _b Number. 
  */ 
  function add(    uint256 _a,    uint256 _b  )    internal    pure    returns (uint256)  {  
      uint256 c = _a + _b;
      assert(c >= _a);
      return c;
      }
}
