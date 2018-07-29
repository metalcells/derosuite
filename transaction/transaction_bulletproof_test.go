// Copyright 2017-2018 DERO Project. All rights reserved.
// Use of this source code in any form is governed by RESEARCH license.
// license can be found in the LICENSE file.
// GPG: 0F39 E425 8C65 3947 702A  8234 08B2 0360 A03A 9DE8
//
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY
// EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL
// THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO,
// PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
// INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT,
// STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF
// THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package transaction

//import "fmt"
import "testing"
import "encoding/hex"

import "github.com/deroproject/derosuite/crypto/ringct"

// test transaction from XMR testnet 0a65ce0098f48085fe2dce36773d696f98b334d1ca2a2510617843bd40600a8d
func Test_Transaction_BP_Simple_Deserialisation_(t *testing.T) {

	//bulletproof_compatibility = true
	ringct.XMR_COMPATIBILITY = true

	defer func() {
		//bulletproof_compatibility = false
		ringct.XMR_COMPATIBILITY = false
	}()

	// transaction from height 1302238 txid 0a65ce0098f48085fe2dce36773d696f98b334d1ca2a2510617843bd40600a8d
	// prefix hash 1bbfda600fa6affc80dae05b1124bf05ed0e20890aa42601441dbe0f6fa81f4b
	tx_data_hex := "0200020200078cbd13bdad01abe002f3345bfb018a03d1fd713c68ae49fc76da03402cd810b3f59bcd301ae8c78ae9c4c56096bf48cc020007b4a40bec8f02ef8203a667d0e004d4c901a73c9a919c5f8b92fa6ac89490523c5cf1e799fe68c2e42e10dcfb5d9001923c879102000240d14ba59d0d8e4dce99dfae958cbedcbd51d11051da490f0ab6cba495a7ce5900025a20a9085da6d8dd787d6f07eceb3d86381fb329830dbaf0ef72818ad24d6b0c2c020901a2fcacef820edaca015fed4954732192ca977a775ba4e9f2186e665a760a427742acb7ce0bfd987a5a04d0d9bcbc0383481b4c1026a1222cc480dc09fa92549c6ea194b731fb7aa0b0655a1ce86504ba1b87994c31d7cd4b326cd458a2d6d7c85e91024e554b035954bb3867e7af074803ea766b3a75a1efa9806928ac39443538a7b708915a30d1a2b8f07f3f8902841b5dc4aed19d16899746a5a60e58b2a9ac11f57b3ca8de9fc38a6eef58cc073b94720a8f20b6048e16bcb6622ebde8e89890bb9f51c4ba3a956e44f4ac5943d2ba2274c6ebb7ba262dc06f52a434f14f0719cda0898775f531db70a58d290c619a00ab2e0bd32b2015af2dd369d4851d05bc0c4eac45f55e7c461ddaab79b1031c1795eb735212c99fdedc36b2c899cacc08c2894cc0c89f59978845367e63b78149bb5a0d21df0d1d63c01ad88241be4dba3781319029a8a7677cccd3c65f54e1a7b06927e63f9bcfabd3e76518770ed611e06e481624b307b8c07eb42defdac91e687dbdedfb5e55b7a4e8da526b9b368f0f9beeb7671cac883bb14da20853a8f8df20472a526f972c92e14a56ab04422446d6c017a6fa585737067d3a0f06630aab413ee23dc0f7c78b98cd8ce076b4dce9ebee66009bb027ae0487d9342516de0ef5887425fcf24572fd6e82439989e2c08e0671575915a4416fc78e05f5aba5678c32ad1f887d7fc6113c8e641635be4b2395e5413afe1c2158b6f0b8fe1088e7dc471f28979afeb2931102ffc7fcd383d0b962e7c086c2540a2290df807ea54482ecc533e5defd32b1f3cf21d71a05916680cf46c256cdae790397fd66da43e9d028efbb79fd1728c1861abf8f04abb01123426da9018f0c2f41a8c48506a24612db82e1003ec0462cb112018705b833166c8d1b47642cfbf49f8a9d690803ac948579dc14c8545fb5504c3de6016733ed05030ba85412e0aaf6511fe9894e581f765924984c46b55bbea9cfe3eaddd0ca484b1a925beee3167f881d1d7906bec5dad652e0a42483f12e6f0cc83459e42a4bc3813ccdc84d582d77e3120aa60c0a51d8a2bef571661ce528093d5b9aea32eb4ef7a966e992f29b99415bcc2552b7ffd86abea43d455751851e4d666844897e4ca9e641fed0ea0c81e82e17d937ca32855fa4547278e0c0f98ec40897afed7d9c1fc04ad11fb3ac96c20a085319691738322456034a4804811f874b8403121599a17e921a0db1ef99968d0ac7d1b07f5c4dc43df7878dbd45e8d0bdc4909027c3168a2c9744ac62175f09020fb5888c29d751ad95a801b2a142ce9c6ffd067ecc4f7ba97e2610101da23df54db448eaf9b0f917d912efaa17adf1d36c6f4b64af65fb3cd3275ce3832fb55613f579756219dd1d9f46a3fe799bf64c5369ad88c70f1375c872efc006414e1d1dad995077a03723cb2879decee5c6e9bf2cd0005c9902570317f9f11e786b188de56d01520dc828af0ecb462e54172c1c4b2ff604fb37f17b34178d2c67e00638943cf48794d70de5cc846b84a6c7a213b53d6a1277b897afd2f6b9dd21fb02064ae706f8282974ea88019aa978ea20475401f02492eb435cf7267610ed156d76711471e896f304d41d99ecc74b7d508e913451fa718daa553c5b9e7d6b6fe0ba30b7aad1408374a4bf7e80756a5c8b0b4d770c0e98d51d8a33b17444fb1d3908d63d7aa577f827ab2e7aa95f5755043ae2fd5f20b7d3bafe9b0faf196579e0eba4750c600da2b9e0a276de65fe32e363d378e166f0845647cc402543cd53d2c5441a8018dd5be27658feb54158c83a6b5afe8df43bd95c36b837b20edea2dffd06aa4f59f903f86211167a14661c637e570f587ec839edc5d63c5b46c68992cea76ca73281dce8c21f823fac6d4d3cbf4556a18e7a1b19042facf0e33fd7d16de69b3d6d8396ef755ab106227b4d11779e1f8d644bff0871ac2c7ae94d2c32c3c15d63132dd6b538c711813e15d2554fd76daa9b791707de823954e9980a831caf02e73ed2999eb8382892aff2d0a7395e04e6daa7879ce9b1bbed62d1e2700dc828184250c457d7dab38d47d09d7a71fcc643b1dbe378b49d0a8a97c5e404cf74cd70bf3e22dbbe2b827e9d539ae45b2f8b5c78ef5913653332bb0bb3c1fbf305374b342df67a3df0d0c02a5e933de8f043927afb5b0d3e04b791176da9bf08019cfd35b0e20bd0fd6c4b1e4f18185d481d63dd21f41a9c471a0f3ecbc3f14e0c6da1dfd146ae9e6e7cb1789be5027db64f61e93e5a86b400fe7b5c80a406550b805fe03c1ca6fff4de32b4fe1e192e26281caf9e0924d46e1cdeaa7d3aa6ae047839e8eb93ebb8638f56fa671a070e6a53d99617bd80746798a1e8567fe27c0de1d5ddb4c6fdef7008539507fa836eb0919b41d2fe9482f81b0e5f72cc676f0ed353cd81e067e00e59eb3bc60d7f14f4cebeb4ae00816caacf9a4a27863bd3097bb9834ce07a6658a340d7da72ea70363feb28acffb30d5528ace4d6d6931a0207b99887f7facf89c5b513c909f6f29b50ade2f91b4ce70cb0ec576fa19fcd0edf482643071c1c0a34e612bf77af2fb760604007bfcae1356ffded3dcc3bbb01b966fb8153004ec43f09747d93fb703a1ee80c77467a8aa54cdb2f2412eec80e91b6cbadc89830e3f481202045c33028ebb89c2f529298bc9ba0ddf472105706b6728a44627cef5438242f121625b262ef4c576e82e2013f3dee94712d46ba0ab537ec28472a6d9b4ec8e086a1a90a5263249f2daf77350dd5bfd004b7ae7c0f69a877536f8f7a994a9384578a94b6724a00e713b97ae0b83df80f5074ca9009dfcd7e8faaf4be1100fb2762c865c60809e59f5bc331655e4f84fb73bfabc7051be7f9c92a7b074f98e70f4d02093bd0c1e6a19a4c31ef669a3ea0eace284d08ee04e12a28c6992c9d7c827bf9f7d454ef521baee3dec80d41c7e9b96e3a250b9e9584a323ef2d020d51fd4ff51990d47f10241ef73c5f4c6cbd6cb8bd31da0f5e7a25614d5d00851f549b83dc4706f99e425d3cb9e618f91c17196db28e5109cbab475ca9a2d9ee8da917a6cf60388219b0ac3d2c88066f0a76d7ce054098002c5cdff1cd06e44828f47f7658c8f3d453cc4ddc61b6b01e55f082f9ea711700dac0e08b38b0a836d07a3a6d4ca42c1aa78fff2370c5757705582ac77619290ef0b936c95cd1d9dd9e0cd58b97e8731b11d62858da749c3548255786cdd4fb05600aaf3e4dfc11881f0538a3fc215377d94f3ab41622d5b1055a9fdfddda080007264f973137fe933158999ee7e07dbcbd2f129da1444ab5cbb5dbed9888410922fa5b86fba1b93d71af7b86d1726cc44cc2953d10958b44844d74ee8361f208349b096e8ca266a82fe4d18a4b23ef67a478a24569fdf709ef6fd7cd333b2104e4dbbcf791eed428714899a9d8fc336baf767a48aa8728025da9cb6758ae3603694168809b957cdf9310a25e9a8fdae4d87c1342e30121b7748552312551250b35b0a70ef4a007e25f56819c363ca6f82755b07158390bbc9da395b79b18d608ee0b6d5c4d5ef4bbda3e2012c5bb7bc3fa92e321a7037efe4bfc8ce717e97508844cd41a067712ec82d3f083ce8130a9a7a05cda632d79829f0b781569f4d485d429e7a0458a4c331f81ed3a4aef57a92e133d7003ad903fb79be07ea3fb5eab"
	tx_data_blob, _ := hex.DecodeString(tx_data_hex)

	var tx Transaction

	err := tx.DeserializeHeader(tx_data_blob)

	if err != nil {
		t.Errorf("Deserialize  transaction  %s\n", err)
		return
	}
	/*
		// verify prefix hash
		calculated_prefix_hash := tx.GetPrefixHash()

		if hex.EncodeToString([]byte(calculated_prefix_hash[:])) != "1bbfda600fa6affc80dae05b1124bf05ed0e20890aa42601441dbe0f6fa81f4b" {
			t.Error("Prefix hash mismatched")
		}

		// be9d2cf9b473dbbb2c59ffb07b5d812516f94d64121d87ad61956386a4bc3843
		calculated_hash := tx.GetHash()

		if hex.EncodeToString([]byte(calculated_hash[:])) != "be9d2cf9b473dbbb2c59ffb07b5d812516f94d64121d87ad61956386a4bc3843" {
			t.Error("transaction hash mismatched")
		}
	*/
	// now serialize once again
	//
	serialised := tx.SerializeHeader()
	_ = serialised

	tx.Clear()

	err = tx.DeserializeHeader(serialised)

	//if err != nil {
	//	t.Error("DeserializeHeader  transaction from blockheight 1302238 beb76a82ea17400cd6d7f595f70e1667d2018ed8f5a78d1ce07484222618c3cd\n")
	//	return
	//}

	if hex.EncodeToString(serialised) != hex.EncodeToString(tx.SerializeHeader()) {

		t.Logf("serialized  %s\n", hex.EncodeToString(serialised))
		t.Logf("reserialized %s\n", hex.EncodeToString(tx.SerializeHeader()))
		t.Error("Serialize TX  Failed")
	}

	// check full tx serialisation
	tx.Clear()
	err = tx.DeserializeHeader(tx_data_blob)

	if err != nil {
		t.Errorf("Bulletproof deserialisation failed")

	}

	serialised = tx.Serialize()
	if hex.EncodeToString(serialised) != tx_data_hex {

		t.Logf("serialized  %s\n", hex.EncodeToString(serialised))
		//t.Logf("reserialized %s\n", hex.EncodeToString(tx.Serialize()))
		t.Error("Serialize TX  Failed")
	}

	//t.Logf("serialised bp %s",hex.EncodeToString(serialised) )

	if tx.IsCoinbase() == true {
		t.Errorf("be9d2cf9b473dbbb2c59ffb07b5d812516f94d64121d87ad61956386a4bc3843 is NOT Coinbase\n")
	}

}