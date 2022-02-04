import wallet_pb2, json
from google.protobuf.json_format import ParseDict, MessageToDict

original = json.load(open('/home/user/projects/wallet-spec/example-wallets/example_wallet'))
back_and_forth = MessageToDict(ParseDict(original, wallet_pb2.Wallet()))

open('/tmp/original', 'w').write(json.dumps(original, sort_keys=True, indent=4))
open('/tmp/back_and_forth', 'w').write(json.dumps(back_and_forth, sort_keys=True, indent=4))

print('meld /tmp/original /tmp/back_and_forth')
