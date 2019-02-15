package jsonrpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSONRPCUnmarhsal(t *testing.T) {
	type TestCase struct {
		Description string
		Raw         string
		Expected    interface{}
	}

	testCases := []TestCase{
		{
			Description: "Request w/ String ID with No Params",
			Raw:         `{"method":"eth_blockNumber","id":"27a5fbbcaa23c1dcca4deb04f1501efb","jsonrpc":"2.0"}`,
			Expected: &Request{
				JSONRPC: "2.0",
				ID: ID{
					Str:      "27a5fbbcaa23c1dcca4deb04f1501efb",
					IsString: true,
				},
				Method: "eth_blockNumber",
				Params: nil,
			},
		},
		{
			Description: "Request w/ Int ID with Single Object Param",
			Raw:         `{"method":"eth_blockNumber","id": 42,"jsonrpc":"2.0", "params":[{"foo":"bar"}]}`,
			Expected: &Request{
				JSONRPC: "2.0",
				ID: ID{
					Num: 42,
				},
				Method: "eth_blockNumber",
				Params: MustParams(map[string]string{"foo": "bar"}),
			},
		},
		{
			Description: "Result with null response",
			Raw:         `{"jsonrpc":"2.0","id":1,"result":null}`,
			Expected: &Response{
				JSONRPC: "2.0",
				ID:      ID{Num: 1},
				Result:  nil,
			},
		},
		{
			Description: "Result with Invalid request response",
			Raw:         `{"jsonrpc":"2.0","error":{"code":-32600,"message":"Invalid request"},"id":null}`,
			Expected: &Response{
				JSONRPC: "2.0",
				ID: ID{
					Num: 0,
				},
				Result: nil,
				Error: map[string]interface{}{
					// TODO: we should define Error struct so code is an int instead
					"code":    float64(-32600),
					"message": "Invalid request",
				},
			},
		},
		{
			Description: "Result with Block response from Fastly",
			Raw:         `{"jsonrpc":"2.0","id":1,"result":{"difficulty":"0x2","extraData":"0xd783010802846765746887676f312e392e34856c696e75780000000000000000cecbd9b0f87e66435801f5a7bcb5919e82492bd6b0ca5d1e66347b3825eda12e34dca3c287438c6f7d8be8fe4b0e7718a148acc9d17ff8c4318b5bec7d6766d201","gasLimit":"0x744b1e","gasUsed":"0x10b0d2","hash":"0x6d9b5e5696d73a4ccfb54895cf843559b1fdc3870f4a040cb7eb1fe668f39707","logsBloom":"0x406000000000000000000000000000000080200100000000000000000008004040000000080020000080000000002000000000000a00000000000010000401000000200040000000000000080000000000000001000c00000000000000801a000040000040000000005084401000080018000c000000200000000010000000000000000000000000080000000200800000000000000042000000000000000000000400804000008200400140800200082000800000410000020000000000000000000002000100000002040000000400000000082000002000000000000040000001000040000000000020000000040000200000000000008002000000000000","miner":"0x0000000000000000000000000000000000000000","mixHash":"0x0000000000000000000000000000000000000000000000000000000000000000","nonce":"0x0000000000000000","number":"0x26cffa","parentHash":"0xebeabac492512016ef0a23f06dc0b1ea708d8f24d042076177a7c68c2279cc44","receiptsRoot":"0x4fcd6093a07c2195a00fc9a1efd3854e0b3013d47d82cccc8d49e08285b961ed","sha3Uncles":"0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347","size":"0xa6e","stateRoot":"0x34a97432816e49acdee7eabd63b96d9d14ec017fcfd2edf15fdef97d9709a71a","timestamp":"0x5b355e6c","totalDifficulty":"0x492fe6","transactions":[{"blockHash":"0x6d9b5e5696d73a4ccfb54895cf843559b1fdc3870f4a040cb7eb1fe668f39707","blockNumber":"0x26cffa","from":"0x498e1a1cd0fb46d82e77836ef4ab973d36bef799","gas":"0x47b760","gasPrice":"0x174876e800","hash":"0x603d89c07c533cfa078427d6238252d860fa5072e3b6255e00ad6dc88edfb8b6","input":"0x454a2ab30000000000000000000000000000000000000000000000000000000000000001","nonce":"0x19","to":"0xfcc46d6898c6c6abff6d768a1777b0f60df596ef","transactionIndex":"0x0","value":"0x2386f26fc10000","v":"0x1c","r":"0xbb82b621b5359fd0b286e6fec84280f42852424a893afaa48803c990638851d3","s":"0x7afff02a99fdefa4d0b71ed06c0804dd112392dd32d315285f25d834d562fbd9"},{"blockHash":"0x6d9b5e5696d73a4ccfb54895cf843559b1fdc3870f4a040cb7eb1fe668f39707","blockNumber":"0x26cffa","from":"0x88b837d3528fb13192e0b7b1009632a8372d8c8a","gas":"0xf4240","gasPrice":"0x37e11d600","hash":"0xabe5ade676001154cda079e8a9345d18767ec8b50c75983dedd942ae813d6e1d","input":"0xe278fe6f","nonce":"0x30e7","to":"0x6f9b3f0640bf7358c87d1d7f2df1a546df0e8c08","transactionIndex":"0x1","value":"0x0","v":"0x2c","r":"0x46ce26e1a30012bb1d9c899f53aa6f6555aa4e7ba2edf1655f41af7ec78a8d8","s":"0x4af0fa939b80d8d9c1e53922c99f2f1fdded2bb6d13d0fa8cb9d50c5f881cc11"},{"blockHash":"0x6d9b5e5696d73a4ccfb54895cf843559b1fdc3870f4a040cb7eb1fe668f39707","blockNumber":"0x26cffa","from":"0x190957f34c9fbffb6b3439c09826969e9bead8ef","gas":"0x697984","gasPrice":"0x1dcd65000","hash":"0x94bb3655ce8973b1a97a8c20cb9d3e3167ea4d16c1fb7beb0f1501a2ca710e2b","input":"0x940a4e450000000000000000000000000000000000000000000000000000000000000001","nonce":"0x7","to":"0x964a56decbf18d97c9a14de2d40050d77d4ecfeb","transactionIndex":"0x2","value":"0x0","v":"0x1c","r":"0x900ad505375d7435932e8579383779ff24f09ee2d0dc750811d0d280fedea2f","s":"0x67e42d8a01ff8f3e38b749f4bd455ba1601ce1ce7fd7b71977cac33efd05286b"},{"blockHash":"0x6d9b5e5696d73a4ccfb54895cf843559b1fdc3870f4a040cb7eb1fe668f39707","blockNumber":"0x26cffa","from":"0x9a65a9f89ca9ba5db7e519fe32b901599fe3d824","gas":"0x3598d","gasPrice":"0x77359400","hash":"0x65e5e5fe97cb2fbb61a14976252c726ecccc8d241ada80c15334f59b823ef810","input":"0x7e4791fe000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000005b355e4f000000000000000000000000000000000000000000000000000000000000000d47616d655f313132385f45544800000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000d3432352d343330207c20302d3000000000000000000000000000000000000000","nonce":"0x113","to":"0x200e281b50899775ef4c9aa7aea19a94cc5323bb","transactionIndex":"0x3","value":"0x3bcbaa6082ea800","v":"0x2c","r":"0xcdd36f3b859c79ea1538e0d0c0400f0bd79cea5233e1c8e6b3f9c21bfcb8328e","s":"0x3e91ee073f86f03f375c5e59f5a8d74ccb3d6d91042bee5bc6a367ce50adc230"},{"blockHash":"0x6d9b5e5696d73a4ccfb54895cf843559b1fdc3870f4a040cb7eb1fe668f39707","blockNumber":"0x26cffa","from":"0xb94898cf42463cedffbe576677b51db6a0f3557f","gas":"0x3598d","gasPrice":"0x77359400","hash":"0x305a5feff74489ac28dbc951a71c1430f03e32424c9d212c567777b64242cc2d","input":"0x7e4791fe000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000005b355e54000000000000000000000000000000000000000000000000000000000000000d47616d655f313132385f45544800000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000d3436312d343633207c20302d3000000000000000000000000000000000000000","nonce":"0xa7","to":"0x200e281b50899775ef4c9aa7aea19a94cc5323bb","transactionIndex":"0x4","value":"0x1995a1245b41800","v":"0x2b","r":"0x5cf4114835cc9a1ae3ba6ba20deafcf65f3cdbbb37ae3e5cb78e24c864e8622d","s":"0x62b9f193fbba819f36d44f07e263ee55b8f3ae86a50bd0699c1f3f5a8faed5ed"},{"blockHash":"0x6d9b5e5696d73a4ccfb54895cf843559b1fdc3870f4a040cb7eb1fe668f39707","blockNumber":"0x26cffa","from":"0xcb71eb21f53a2f4de0f26dc90518df10be13d1ec","gas":"0x321af","gasPrice":"0x3b9aca00","hash":"0x3f9aca4a9b4565314bdffec8071bc82980642a491b09bc4b5b449201b3a4f057","input":"0xc7a1865b6461736461736461202020202020202020202020202020202020202020202020","nonce":"0xd6","to":"0xb8998227bed70369e1c477f6abcb97f1409a4006","transactionIndex":"0x5","value":"0x16345785d8a0000","v":"0x2c","r":"0x3e6316ec4096f6a9e820cd929bc72fcb5e35132c85c6ab1f5d66afc44ac2f143","s":"0x60ebc69ef70c5b37de2265a9c2cf1d4df8f36e986c41f90d18d379f32972a840"},{"blockHash":"0x6d9b5e5696d73a4ccfb54895cf843559b1fdc3870f4a040cb7eb1fe668f39707","blockNumber":"0x26cffa","from":"0x6f9e990f5b59265e63f12722acb49acc55e2fe56","gas":"0x1d40a","gasPrice":"0x3b9aca00","hash":"0x33865fc3e6cc986731e8cc6c8ef7c1b7d52388af8fc946937c71307f2a382015","input":"0xf7d8c8830000000000000000000000000000000000000000000000000000000000000b040000000000000000000000000000000000000000000000000000000000000f7c","nonce":"0x872","to":"0x5328276603d169165d0f71ca67ccc89c45027df3","transactionIndex":"0x6","value":"0x71afd498d0000","v":"0x2b","r":"0x46c4fef9267f8a48e78c46c5b91cdff6c6442761a63b28bae434545507870d7c","s":"0x7276c04cc4ab00867124b3af5f676b14276b1fcc801fa9b92d5ac43a0bad166d"},{"blockHash":"0x6d9b5e5696d73a4ccfb54895cf843559b1fdc3870f4a040cb7eb1fe668f39707","blockNumber":"0x26cffa","from":"0xa608b142496c257331de4c62bc0948bd514d7c9f","gas":"0x477ba","gasPrice":"0x3b9aca00","hash":"0x19f7df2e775c094db738111e3c8f4b629721ce6b9cb6dd57640165f8a1ad6f01","input":"0xea94496b000000000000000000000000000000000000000000000000000000000008577400000000000000000000000000000000000000000000000000000000000855ac","nonce":"0xae","to":"0xfe2149773b3513703e79ad23d05a778a185016ee","transactionIndex":"0x7","value":"0xaa87bee538000","v":"0x2b","r":"0x621eeeba24db422161e8a65f794d710d425fa72abe493fa5d4c1287869dacec7","s":"0xa0bf0af07bf989e2fca9ff006eea450d2752c403d786bb378e387d870f2bb6"},{"blockHash":"0x6d9b5e5696d73a4ccfb54895cf843559b1fdc3870f4a040cb7eb1fe668f39707","blockNumber":"0x26cffa","from":"0xca39e90cec69838e73cc4f24ec5077dac44b47d6","gas":"0x311e4","gasPrice":"0x3b9aca00","hash":"0x12373ee6bbf5d3e63003f03c2812291835ae7173793271b8efb5049eec36d946","input":"0x3d7d3f5a000000000000000000000000000000000000000000000000000000000008579a00000000000000000000000000000000000000000000000009b6e64a8ec60000000000000000000000000000000000000000000000000000058d15e176280000000000000000000000000000000000000000000000000000000000000002a300","nonce":"0x47f2","to":"0xfe2149773b3513703e79ad23d05a778a185016ee","transactionIndex":"0x8","value":"0x0","v":"0x1c","r":"0xdcee56602d0308ab10bbfea8fc59c07df121087a32f441a86ed219dd5bc157f7","s":"0x4717e57682bf8305b4a7c9d54949c4c8de52ab9ae00af8a68c7ef91c2cbee733"},{"blockHash":"0x6d9b5e5696d73a4ccfb54895cf843559b1fdc3870f4a040cb7eb1fe668f39707","blockNumber":"0x26cffa","from":"0xca39e90cec69838e73cc4f24ec5077dac44b47d6","gas":"0x311e4","gasPrice":"0x3b9aca00","hash":"0x03f4e5505914d08d70f09d477b9d42591d3537b5fd9a067ff25cb657ee9691f1","input":"0x3d7d3f5a000000000000000000000000000000000000000000000000000000000008579b00000000000000000000000000000000000000000000000009b6e64a8ec60000000000000000000000000000000000000000000000000000058d15e176280000000000000000000000000000000000000000000000000000000000000002a300","nonce":"0x47f3","to":"0xfe2149773b3513703e79ad23d05a778a185016ee","transactionIndex":"0x9","value":"0x0","v":"0x1c","r":"0xca4da179b34ab3dfac23443f26f75adf8b3b841c7bdddf542a177b4864d4070","s":"0x5c372293c5667898188a339e85ead39f67842107dd23712fa9bdbe41550be5b6"}],"transactionsRoot":"0xf5e56a5dc26fb78a01ef141aa8e6ee1ae80da6fef36eba097118b04e928d510e","uncles":[]}}`,
			Expected: &Response{
				JSONRPC: "2.0",
				ID: ID{
					Num: 1,
				},
				Result: map[string]interface{}{"size": "0xa6e", "stateRoot": "0x34a97432816e49acdee7eabd63b96d9d14ec017fcfd2edf15fdef97d9709a71a", "timestamp": "0x5b355e6c", "gasLimit": "0x744b1e", "hash": "0x6d9b5e5696d73a4ccfb54895cf843559b1fdc3870f4a040cb7eb1fe668f39707", "transactions": []interface{}{map[string]interface{}{"s": "0x7afff02a99fdefa4d0b71ed06c0804dd112392dd32d315285f25d834d562fbd9", "v": "0x1c", "blockHash": "0x6d9b5e5696d73a4ccfb54895cf843559b1fdc3870f4a040cb7eb1fe668f39707", "blockNumber": "0x26cffa", "hash": "0x603d89c07c533cfa078427d6238252d860fa5072e3b6255e00ad6dc88edfb8b6", "input": "0x454a2ab30000000000000000000000000000000000000000000000000000000000000001", "from": "0x498e1a1cd0fb46d82e77836ef4ab973d36bef799", "gas": "0x47b760", "transactionIndex": "0x0", "gasPrice": "0x174876e800", "nonce": "0x19", "r": "0xbb82b621b5359fd0b286e6fec84280f42852424a893afaa48803c990638851d3", "to": "0xfcc46d6898c6c6abff6d768a1777b0f60df596ef", "value": "0x2386f26fc10000"}, map[string]interface{}{"blockNumber": "0x26cffa", "value": "0x0", "blockHash": "0x6d9b5e5696d73a4ccfb54895cf843559b1fdc3870f4a040cb7eb1fe668f39707", "gas": "0xf4240", "hash": "0xabe5ade676001154cda079e8a9345d18767ec8b50c75983dedd942ae813d6e1d", "to": "0x6f9b3f0640bf7358c87d1d7f2df1a546df0e8c08", "v": "0x2c", "nonce": "0x30e7", "transactionIndex": "0x1", "from": "0x88b837d3528fb13192e0b7b1009632a8372d8c8a", "gasPrice": "0x37e11d600", "input": "0xe278fe6f", "r": "0x46ce26e1a30012bb1d9c899f53aa6f6555aa4e7ba2edf1655f41af7ec78a8d8", "s": "0x4af0fa939b80d8d9c1e53922c99f2f1fdded2bb6d13d0fa8cb9d50c5f881cc11"}, map[string]interface{}{"from": "0x190957f34c9fbffb6b3439c09826969e9bead8ef", "value": "0x0", "to": "0x964a56decbf18d97c9a14de2d40050d77d4ecfeb", "transactionIndex": "0x2", "r": "0x900ad505375d7435932e8579383779ff24f09ee2d0dc750811d0d280fedea2f", "s": "0x67e42d8a01ff8f3e38b749f4bd455ba1601ce1ce7fd7b71977cac33efd05286b", "hash": "0x94bb3655ce8973b1a97a8c20cb9d3e3167ea4d16c1fb7beb0f1501a2ca710e2b", "input": "0x940a4e450000000000000000000000000000000000000000000000000000000000000001", "gas": "0x697984", "gasPrice": "0x1dcd65000", "nonce": "0x7", "v": "0x1c", "blockHash": "0x6d9b5e5696d73a4ccfb54895cf843559b1fdc3870f4a040cb7eb1fe668f39707", "blockNumber": "0x26cffa"}, map[string]interface{}{"blockNumber": "0x26cffa", "gasPrice": "0x77359400", "nonce": "0x113", "to": "0x200e281b50899775ef4c9aa7aea19a94cc5323bb", "blockHash": "0x6d9b5e5696d73a4ccfb54895cf843559b1fdc3870f4a040cb7eb1fe668f39707", "from": "0x9a65a9f89ca9ba5db7e519fe32b901599fe3d824", "gas": "0x3598d", "hash": "0x65e5e5fe97cb2fbb61a14976252c726ecccc8d241ada80c15334f59b823ef810", "r": "0xcdd36f3b859c79ea1538e0d0c0400f0bd79cea5233e1c8e6b3f9c21bfcb8328e", "v": "0x2c", "value": "0x3bcbaa6082ea800", "transactionIndex": "0x3", "input": "0x7e4791fe000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000005b355e4f000000000000000000000000000000000000000000000000000000000000000d47616d655f313132385f45544800000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000d3432352d343330207c20302d3000000000000000000000000000000000000000", "s": "0x3e91ee073f86f03f375c5e59f5a8d74ccb3d6d91042bee5bc6a367ce50adc230"}, map[string]interface{}{"from": "0xb94898cf42463cedffbe576677b51db6a0f3557f", "gas": "0x3598d", "r": "0x5cf4114835cc9a1ae3ba6ba20deafcf65f3cdbbb37ae3e5cb78e24c864e8622d", "gasPrice": "0x77359400", "s": "0x62b9f193fbba819f36d44f07e263ee55b8f3ae86a50bd0699c1f3f5a8faed5ed", "value": "0x1995a1245b41800", "hash": "0x305a5feff74489ac28dbc951a71c1430f03e32424c9d212c567777b64242cc2d", "nonce": "0xa7", "to": "0x200e281b50899775ef4c9aa7aea19a94cc5323bb", "transactionIndex": "0x4", "blockHash": "0x6d9b5e5696d73a4ccfb54895cf843559b1fdc3870f4a040cb7eb1fe668f39707", "blockNumber": "0x26cffa", "input": "0x7e4791fe000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000005b355e54000000000000000000000000000000000000000000000000000000000000000d47616d655f313132385f45544800000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000d3436312d343633207c20302d3000000000000000000000000000000000000000", "v": "0x2b"}, map[string]interface{}{"input": "0xc7a1865b6461736461736461202020202020202020202020202020202020202020202020", "r": "0x3e6316ec4096f6a9e820cd929bc72fcb5e35132c85c6ab1f5d66afc44ac2f143", "transactionIndex": "0x5", "value": "0x16345785d8a0000", "gas": "0x321af", "hash": "0x3f9aca4a9b4565314bdffec8071bc82980642a491b09bc4b5b449201b3a4f057", "blockHash": "0x6d9b5e5696d73a4ccfb54895cf843559b1fdc3870f4a040cb7eb1fe668f39707", "from": "0xcb71eb21f53a2f4de0f26dc90518df10be13d1ec", "gasPrice": "0x3b9aca00", "nonce": "0xd6", "s": "0x60ebc69ef70c5b37de2265a9c2cf1d4df8f36e986c41f90d18d379f32972a840", "v": "0x2c", "blockNumber": "0x26cffa", "to": "0xb8998227bed70369e1c477f6abcb97f1409a4006"}, map[string]interface{}{"gasPrice": "0x3b9aca00", "r": "0x46c4fef9267f8a48e78c46c5b91cdff6c6442761a63b28bae434545507870d7c", "to": "0x5328276603d169165d0f71ca67ccc89c45027df3", "transactionIndex": "0x6", "v": "0x2b", "gas": "0x1d40a", "nonce": "0x872", "blockNumber": "0x26cffa", "from": "0x6f9e990f5b59265e63f12722acb49acc55e2fe56", "input": "0xf7d8c8830000000000000000000000000000000000000000000000000000000000000b040000000000000000000000000000000000000000000000000000000000000f7c", "s": "0x7276c04cc4ab00867124b3af5f676b14276b1fcc801fa9b92d5ac43a0bad166d", "blockHash": "0x6d9b5e5696d73a4ccfb54895cf843559b1fdc3870f4a040cb7eb1fe668f39707", "hash": "0x33865fc3e6cc986731e8cc6c8ef7c1b7d52388af8fc946937c71307f2a382015", "value": "0x71afd498d0000"}, map[string]interface{}{"s": "0xa0bf0af07bf989e2fca9ff006eea450d2752c403d786bb378e387d870f2bb6", "blockNumber": "0x26cffa", "from": "0xa608b142496c257331de4c62bc0948bd514d7c9f", "gas": "0x477ba", "nonce": "0xae", "transactionIndex": "0x7", "blockHash": "0x6d9b5e5696d73a4ccfb54895cf843559b1fdc3870f4a040cb7eb1fe668f39707", "to": "0xfe2149773b3513703e79ad23d05a778a185016ee", "r": "0x621eeeba24db422161e8a65f794d710d425fa72abe493fa5d4c1287869dacec7", "v": "0x2b", "value": "0xaa87bee538000", "gasPrice": "0x3b9aca00", "hash": "0x19f7df2e775c094db738111e3c8f4b629721ce6b9cb6dd57640165f8a1ad6f01", "input": "0xea94496b000000000000000000000000000000000000000000000000000000000008577400000000000000000000000000000000000000000000000000000000000855ac"}, map[string]interface{}{"v": "0x1c", "blockHash": "0x6d9b5e5696d73a4ccfb54895cf843559b1fdc3870f4a040cb7eb1fe668f39707", "gasPrice": "0x3b9aca00", "to": "0xfe2149773b3513703e79ad23d05a778a185016ee", "input": "0x3d7d3f5a000000000000000000000000000000000000000000000000000000000008579a00000000000000000000000000000000000000000000000009b6e64a8ec60000000000000000000000000000000000000000000000000000058d15e176280000000000000000000000000000000000000000000000000000000000000002a300", "nonce": "0x47f2", "blockNumber": "0x26cffa", "gas": "0x311e4", "transactionIndex": "0x8", "s": "0x4717e57682bf8305b4a7c9d54949c4c8de52ab9ae00af8a68c7ef91c2cbee733", "value": "0x0", "from": "0xca39e90cec69838e73cc4f24ec5077dac44b47d6", "hash": "0x12373ee6bbf5d3e63003f03c2812291835ae7173793271b8efb5049eec36d946", "r": "0xdcee56602d0308ab10bbfea8fc59c07df121087a32f441a86ed219dd5bc157f7"}, map[string]interface{}{"s": "0x5c372293c5667898188a339e85ead39f67842107dd23712fa9bdbe41550be5b6", "value": "0x0", "gas": "0x311e4", "hash": "0x03f4e5505914d08d70f09d477b9d42591d3537b5fd9a067ff25cb657ee9691f1", "input": "0x3d7d3f5a000000000000000000000000000000000000000000000000000000000008579b00000000000000000000000000000000000000000000000009b6e64a8ec60000000000000000000000000000000000000000000000000000058d15e176280000000000000000000000000000000000000000000000000000000000000002a300", "to": "0xfe2149773b3513703e79ad23d05a778a185016ee", "v": "0x1c", "blockHash": "0x6d9b5e5696d73a4ccfb54895cf843559b1fdc3870f4a040cb7eb1fe668f39707", "gasPrice": "0x3b9aca00", "nonce": "0x47f3", "transactionIndex": "0x9", "blockNumber": "0x26cffa", "from": "0xca39e90cec69838e73cc4f24ec5077dac44b47d6", "r": "0xca4da179b34ab3dfac23443f26f75adf8b3b841c7bdddf542a177b4864d4070"}}, "difficulty": "0x2", "sha3Uncles": "0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347", "logsBloom": "0x406000000000000000000000000000000080200100000000000000000008004040000000080020000080000000002000000000000a00000000000010000401000000200040000000000000080000000000000001000c00000000000000801a000040000040000000005084401000080018000c000000200000000010000000000000000000000000080000000200800000000000000042000000000000000000000400804000008200400140800200082000800000410000020000000000000000000002000100000002040000000400000000082000002000000000000040000001000040000000000020000000040000200000000000008002000000000000", "mixHash": "0x0000000000000000000000000000000000000000000000000000000000000000", "nonce": "0x0000000000000000", "parentHash": "0xebeabac492512016ef0a23f06dc0b1ea708d8f24d042076177a7c68c2279cc44", "receiptsRoot": "0x4fcd6093a07c2195a00fc9a1efd3854e0b3013d47d82cccc8d49e08285b961ed", "totalDifficulty": "0x492fe6", "extraData": "0xd783010802846765746887676f312e392e34856c696e75780000000000000000cecbd9b0f87e66435801f5a7bcb5919e82492bd6b0ca5d1e66347b3825eda12e34dca3c287438c6f7d8be8fe4b0e7718a148acc9d17ff8c4318b5bec7d6766d201", "gasUsed": "0x10b0d2", "transactionsRoot": "0xf5e56a5dc26fb78a01ef141aa8e6ee1ae80da6fef36eba097118b04e928d510e", "uncles": []interface{}{}, "miner": "0x0000000000000000000000000000000000000000", "number": "0x26cffa"},
			},
		},
		{
			Description: "Notification result from Parity",
			Raw:         `{"jsonrpc":"2.0","method":"parity_subscription","params":{"result":"0x3342d6","subscription":"0x0c2f1dc472de1be0"}}`,
			Expected: &Notification{
				JSONRPC: "2.0",
				Method:  "parity_subscription",
				Params:  NotificationParams(`{"result":"0x3342d6","subscription":"0x0c2f1dc472de1be0"}`),
			},
		},
		{
			Description: "Notification result from Geth",
			Raw:         `{"jsonrpc":"2.0","method":"eth_subscription","params":{"subscription":"0x3eb3487232a1bf601f92757e0a5d0b18","result":{"parentHash":"0x28f2668a84038a5b07d13564b7b11421c7ca74867f80a06c8bf429057c5000dd","truncated":"..."}}}`,
			Expected: &Notification{
				JSONRPC: "2.0",
				Method:  "eth_subscription",
				Params:  NotificationParams(`{"subscription":"0x3eb3487232a1bf601f92757e0a5d0b18","result":{"parentHash":"0x28f2668a84038a5b07d13564b7b11421c7ca74867f80a06c8bf429057c5000dd","truncated":"..."}}`),
			},
		},
	}

	for _, testCase := range testCases {
		msg, err := Unmarshal([]byte(testCase.Raw))
		assert.NoError(t, err, "Should not error in jsonrpc.Unmarshal")

		switch msg := msg.(type) {
		case *Response:
			assert.Equal(t, testCase.Expected.(*Response), msg, "%v", &testCase)
		case *Request:
			assert.Equal(t, testCase.Expected.(*Request), msg, "%v", &testCase)
		case *Notification:
			assert.Equal(t, testCase.Expected.(*Notification), msg, "%v", &testCase)
		}
	}
}
