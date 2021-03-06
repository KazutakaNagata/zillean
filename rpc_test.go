package zillean

import (
	"math/big"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewRPC(t *testing.T) {
	Convey("returns a new rpc", t, func() {
		So(NewRPC(localNet), ShouldHaveSameTypeAs, &RPC{})
	})
}

func TestRPC_GetBalance(t *testing.T) {
	Convey("returns the balance and nonce of a given address", t, func() {
		result, err := NewRPC(testNet).GetBalance("5568CF7C38334A4E960BC99D8F22C1E90645E5F2")
		So(err, ShouldBeNil)
		So(result.Balance, ShouldEqual, "1000000000000")
		So(result.Nonce, ShouldEqual, 0)
	})
}

func TestRPC_GetDsBlock(t *testing.T) {
	Convey("returns details of a Directory Service block by block number", t, func() {
		result, err := NewRPC(testNet).GetDsBlock("1")
		So(err, ShouldBeNil)
		So(result.Header.BlockNum, ShouldEqual, "1")
		So(result.Header.Difficulty, ShouldEqual, 3)
		So(result.Header.DifficultyDS, ShouldEqual, 5)
		So(result.Header.GasPrice, ShouldEqual, "1000000000")
		So(result.Header.LeaderPubKey, ShouldEqual, "0x02002F3550709EDA6A480F4E63FD7EAAB151149A7DA5726E37B317AC39D78549CF")
		So(result.Header.PoWWinners, ShouldResemble, []string{"0x038F3F572D3C873EF1F06AFC9247D584163CDEE3AE6AF943CCB8F7293FAF2FE500", "0x039045C744B71E64A5098DCC2E3614989D461BA5C195AF71C15B56F38BFE1E6809", "0x039E5811FE01268DDDAF5281A72915B1C564F8BC6424202488B2E249EFB33D6ED9", "0x03B03E2854A5E302B25BCA1F1DFE03DEA537B4EBEDA5F65E72505EA51E9CB52287", "0x03B5F0F75B69A3C6E4CD31F68352A924F97AB983694D7461AE441EBA397F6A3DD5", "0x03D87B6FD7B3D8A80285B8DE5D8137E6A1C4EE3670F8E3A036EF150A57D86A8045", "0x03E36CA078F3E6E28F156BD623D6407E9E746E062B3685DC7D10BA34E1FA168A5F", "0x03E87174AE9993E8ABC1D823EE8D3AD3B78FB598FB778EB4B22DB92DAE1C15E8EE", "0x03F2374DE8FCF30EA993592B93E976E5171073F1240C8B3F3687711428028DC3F0", "0x03F6D3A5767221E970B40A38840C99649ADC398F83EE4848C4561291E3AE633C92"})
		So(result.Header.Prevhash, ShouldEqual, "ba127538d2c63eec121629011ae8173210589689dca54d1e11904dd82c68e9da")
		So(result.Header.Timestamp, ShouldEqual, "1545390003677310")
		So(result.Signature, ShouldEqual, "09A271C06660111DD2CA581D018C987155E58E1B4811B0E2BA079CE9060B2401F2B10C6E5D2A3AA3DC4C61A11763FE29E0A6AC673A24D4D323457023D8F1E252")
	})
}

func TestRPC_GetTxBlock(t *testing.T) {
	Convey("returns details of a Transaction block by block number.", t, func() {
		result, err := NewRPC(testNet).GetTxBlock("100")
		So(err, ShouldBeNil)
		So(result.Body.HeaderSign, ShouldEqual, "A45D6BF82D00C13360C87625EECC8D9D7CD4DF62D9C4CB037D2F0AFB725CDCA9F5965BC4AB1E55D220153C0D7336B8E9D794F6478A9F115F9DE046A0A14C3BCA")
		So(len(result.Body.MicroBlockInfos), ShouldEqual, 4)
		So(result.Header.BlockNum, ShouldEqual, "100")
		So(result.Header.DsBlockNum, ShouldEqual, "2")
		So(result.Header.GasLimit, ShouldEqual, "2000000")
		So(result.Header.GasUsed, ShouldEqual, "0")
		So(result.Header.MbInfoHash, ShouldEqual, "2e138097da1707573580ac2538e7681582adee762a8ee367efe04b657a653c3a")
		So(result.Header.MinerPubKey, ShouldEqual, "0x02129C96C0ABDA482E959AE0C7EF0FB19D056233DCE96E6E8E7DE49165383BE2CC")
		So(result.Header.NumMicroBlocks, ShouldEqual, 4)
		So(result.Header.NumTxns, ShouldEqual, 0)
		So(result.Header.PrevBlockHash, ShouldEqual, "69a621c3fdc64945d2eed7a38a0da9b1ebe5de1f64cc4f48fc1219665fdeef84")
		So(result.Header.Rewards, ShouldEqual, "0")
		So(result.Header.StateDeltaHash, ShouldEqual, "0000000000000000000000000000000000000000000000000000000000000000")
		So(result.Header.StateRootHash, ShouldEqual, "1e0fea2b34b7a8ebe6d1d9a31d6c3c4910529a95d702481a13cd9e788fc26182")
		So(result.Header.Timestamp, ShouldEqual, "1545393785852975")
		So(result.Header.Type, ShouldEqual, 1)
		So(result.Header.Version, ShouldEqual, 0)
	})
}

func TestRPC_GetLatestDsBlock(t *testing.T) {
	Convey("returns details of the most recent Directory Service block", t, func() {
		result, err := NewRPC(testNet).GetLatestDsBlock()
		So(err, ShouldBeNil)
		So(result.Header.BlockNum, ShouldNotBeBlank)
		So(result.Header.Difficulty, ShouldBeGreaterThan, 0)
		So(result.Header.LeaderPubKey, ShouldNotBeBlank)
		So(result.Header.Prevhash, ShouldNotBeBlank)
		So(result.Header.Timestamp, ShouldNotBeBlank)
		So(result.Signature, ShouldNotBeBlank)
	})
}

func TestRPC_GetLatestTxBlock(t *testing.T) {
	Convey("returns details of the most recent Transaction block", t, func() {
		result, err := NewRPC(testNet).GetLatestTxBlock()
		So(err, ShouldBeNil)
		So(result.Body.HeaderSign, ShouldNotBeBlank)
		So(result.Body.MicroBlockInfos, ShouldNotBeNil)
		So(result.Header.BlockNum, ShouldNotBeBlank)
		So(result.Header.DsBlockNum, ShouldNotBeBlank)
		So(result.Header.GasLimit, ShouldNotBeBlank)
		So(result.Header.GasUsed, ShouldNotBeBlank)
		So(result.Header.MbInfoHash, ShouldNotBeBlank)
		So(result.Header.MinerPubKey, ShouldNotBeBlank)
		So(result.Header.NumMicroBlocks, ShouldBeGreaterThan, 0)
		So(result.Header.NumTxns, ShouldHaveSameTypeAs, int64(0))
		So(result.Header.PrevBlockHash, ShouldNotBeBlank)
		So(result.Header.Rewards, ShouldNotBeBlank)
		So(result.Header.StateDeltaHash, ShouldNotBeBlank)
		So(result.Header.StateRootHash, ShouldNotBeBlank)
		So(result.Header.Timestamp, ShouldNotBeBlank)
		So(result.Header.Type, ShouldHaveSameTypeAs, int64(1))
		So(result.Header.Version, ShouldHaveSameTypeAs, int64(0))
	})
}

func TestRPC_GetTransaction(t *testing.T) {
	Convey("returns details of a Transaction by its hash", t, func() {
		result, err := NewRPC(testNet).GetTransaction("5a09b08a28ca4b6ef1935bb3d00307530ad64e4150fba573cdb9f0dea847d1c7")
		So(err, ShouldBeNil)
		So(result.ID, ShouldEqual, "5a09b08a28ca4b6ef1935bb3d00307530ad64e4150fba573cdb9f0dea847d1c7")
		So(result.Amount, ShouldEqual, "1000000000000")
		So(result.GasLimit, ShouldEqual, "1")
		So(result.GasPrice, ShouldEqual, "1000000000")
		So(result.Nonce, ShouldEqual, "1")
		So(result.Receipt.CumulativeGas, ShouldEqual, "1")
		So(result.Receipt.Success, ShouldBeTrue)
		So(result.SenderPubKey, ShouldEqual, "0x03B994288C5B6CB99A694BEF9F7E01B1773DDB8F662ADE12D98B880487232D3565")
		So(result.Signature, ShouldEqual, "0x2576BC15EC10C5BB5A673EE4A5B72F8D3C824EC0C96AE3B69EE098290BA6044EE605BD1D9EDF42271AC16639703B42685AE276EC4703F3B2F86BF47DE3BA658D")
		So(result.ToAddr, ShouldEqual, "5568cf7c38334a4e960bc99d8f22c1e90645e5f2")
		So(result.Version, ShouldEqual, "0")
	})
}

func TestRPC_CreateTransaction(t *testing.T) {
	Convey("returns a hash of created Transaction", t, func() {
		zillean := NewZillean(testNet)
		privateKey := "AAFD338492962FAD674EE3BD6EBC57C8373B2C9BADBAC8806D890F1FE8C571DF"
		publicKey, _ := zillean.GetPublicKeyFromPrivateKey(privateKey)
		rawTx := RawTransaction{
			Version:  0,
			Nonce:    1,
			To:       "5568CF7C38334A4E960BC99D8F22C1E90645E5F2",
			Amount:   "1000000000000",
			PubKey:   publicKey,
			GasPrice: big.NewInt(1000000000),
			GasLimit: 1,
		}
		k, _ := GenerateDRN(EncodeTransaction(rawTx))
		signature, _ := zillean.SignTransaction(k, rawTx, privateKey)
		result, err := zillean.RPC.CreateTransaction(rawTx, signature)
		So(err, ShouldBeNil)
		So(result, ShouldEqual, "5a09b08a28ca4b6ef1935bb3d00307530ad64e4150fba573cdb9f0dea847d1c7")
	})
}

func TestRPC_GetSmartContracts(t *testing.T) {
	Convey("returns the list of smart contracts created by an address", t, func() {
		result, err := NewRPC(testNet).GetSmartContracts("f3d2005b55102d6588dd9771e9356f1908c9d97f")
		So(err, ShouldBeNil)
		So(len(result), ShouldEqual, 1)
		So(result[0].Address, ShouldEqual, "83536f90ed096b5d14ba2c296a32f37849dd3221")
		So(result[0].State[0].Type, ShouldEqual, "Uint128")
		So(result[0].State[0].Value, ShouldEqual, "0")
		So(result[0].State[0].Vname, ShouldEqual, "_balance")
	})
}

func TestRPC_GetSmartContractState(t *testing.T) {
	Convey("returns the state variables (mutable) of a smart contract address", t, func() {
		result, err := NewRPC(testNet).GetSmartContractState("83536f90ed096b5d14ba2c296a32f37849dd3221")
		So(err, ShouldBeNil)
		So(len(result), ShouldEqual, 1)
		So(result[0].Type, ShouldEqual, "Uint128")
		So(result[0].Value, ShouldEqual, "0")
		So(result[0].Vname, ShouldEqual, "_balance")
	})
}

func TestRPC_GetSmartContractCode(t *testing.T) {
	Convey("returns the Scilla code of a smart contract address", t, func() {
		result, err := NewRPC(testNet).GetSmartContractCode("83536f90ed096b5d14ba2c296a32f37849dd3221")
		So(err, ShouldBeNil)
		So(result, ShouldEqual, "scilla_version 0\n\n    (* HelloWorld contract *)\n    \n    import ListUtils\n    \n    (***************************************************)\n    (*               Associated library                *)\n    (***************************************************)\n    library HelloWorld\n    \n    let one_msg = \n      fun (msg : Message) => \n      let nil_msg = Nil {Message} in\n      Cons {Message} msg nil_msg\n    \n    let not_owner_code = Int32 1\n    let set_hello_code = Int32 2\n    \n    (***************************************************)\n    (*             The contract definition             *)\n    (***************************************************)\n    \n    contract HelloWorld\n    (owner: ByStr20)\n    \n    field welcome_msg : String = \"\"\n    \n    transition setHello (msg : String)\n      is_owner = builtin eq owner _sender;\n      match is_owner with\n      | False =>\n        msg = {_tag : \"Main\"; _recipient : _sender; _amount : Uint128 0; code : not_owner_code};\n        msgs = one_msg msg;\n        send msgs\n      | True =>\n        welcome_msg := msg;\n        msg = {_tag : \"Main\"; _recipient : _sender; _amount : Uint128 0; code : set_hello_code};\n        msgs = one_msg msg;\n        send msgs\n      end\n    end\n    \n    \n    transition getHello ()\n        r <- welcome_msg;\n        e = {_eventname: \"getHello()\"; msg: r};\n        event e\n    end")
	})
}

func TestRPC_GetSmartContractInit(t *testing.T) {
	Convey("returns the initialization parameters (immutable) of a given smart contract address", t, func() {
		result, err := NewRPC(testNet).GetSmartContractInit("83536f90ed096b5d14ba2c296a32f37849dd3221")
		So(err, ShouldBeNil)
		So(len(result), ShouldEqual, 3)
		So(result[0].Type, ShouldEqual, "Uint32")
		So(result[0].Value, ShouldEqual, "0")
		So(result[0].Vname, ShouldEqual, "_scilla_version")
		So(result[1].Type, ShouldEqual, "ByStr20")
		So(result[1].Value, ShouldEqual, "0xf3d2005b55102d6588dd9771e9356f1908c9d97f")
		So(result[1].Vname, ShouldEqual, "owner")
		So(result[2].Type, ShouldEqual, "BNum")
		So(result[2].Value, ShouldEqual, "9247")
		So(result[2].Vname, ShouldEqual, "_creation_block")
	})
}

func TestRPC_GetBlockchainInfo(t *testing.T) {
	Convey("returns statistics about the specified zilliqa node", t, func() {
		result, err := NewRPC(testNet).GetBlockchainInfo()
		So(err, ShouldBeNil)
		So(result.CurrentDSEpoch, ShouldNotBeBlank)
		So(result.CurrentMiniEpoch, ShouldNotBeBlank)
		So(result.DSBlockRate, ShouldBeGreaterThan, 0)
		So(result.NumDSBlocks, ShouldNotBeBlank)
		So(result.NumPeers, ShouldBeGreaterThan, 0)
		So(result.NumTransactions, ShouldNotBeBlank)
		So(result.NumTxBlocks, ShouldNotBeBlank)
		So(result.NumTxnsDSEpoch, ShouldNotBeBlank)
		So(result.NumTxnsTxEpoch, ShouldHaveSameTypeAs, int64(0))
		So(result.CurrentDSEpoch, ShouldNotBeBlank)
		So(result.ShardingStructure.NumPeers, ShouldHaveSameTypeAs, []int64{})
		So(len(result.ShardingStructure.NumPeers), ShouldBeGreaterThan, 0)
		So(result.TransactionRate, ShouldHaveSameTypeAs, int64(0))
		So(result.TxBlockRate, ShouldBeGreaterThan, 0)
	})
}

func TestRPC_GetNetworkID(t *testing.T) {
	Convey("returns the network ID of the specified zilliqa node", t, func() {
		result, err := NewRPC(testNet).GetNetworkID()
		So(err, ShouldBeNil)
		So(result, ShouldEqual, "TestNet")
	})
}

func TestRPC_GetRecentTransactions(t *testing.T) {
	Convey("returns  the most recent transactions (upto 100) accepted by the specified zilliqa node.", t, func() {
		result, err := NewRPC(testNet).GetRecentTransactions()
		So(err, ShouldBeNil)
		So(len(result.TxnHashes), ShouldEqual, result.Number)
	})
}

func TestRPC_GetDSBlockListing(t *testing.T) {
	Convey("returns a paginated list of Directory Service blocks", t, func() {
		result, err := NewRPC(testNet).DSBlockListing(1)
		So(err, ShouldBeNil)
		So(len(result.Data), ShouldEqual, 10)
		So(result.MaxPages, ShouldBeGreaterThan, 0)
	})
}

func TestRPC_GetTxBlockListing(t *testing.T) {
	Convey("returns a paginated list of Transaction blocks", t, func() {
		result, err := NewRPC(testNet).TxBlockListing(1)
		So(err, ShouldBeNil)
		So(len(result.Data), ShouldEqual, 10)
		So(result.MaxPages, ShouldBeGreaterThan, 0)
	})
}
