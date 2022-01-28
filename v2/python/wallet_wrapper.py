from google.protobuf.json_format import MessageToDict, ParseDict

import json

import wallet_pb2

def from_json(wallet_json, wallet_pb):
  wallet_dict = json.loads(wallet_json)
  settings = wallet_dict['preferences']['shared']['value']['value']['settings']
  settings['lbryum_servers'] = [
    {'domain': domain, 'port': port}
    for [domain, port]
    in settings['lbryum_servers']
  ]
  return ParseDict(wallet_dict, wallet_pb)

def to_json(wallet_pb, including_default_value_fields):
  wallet_dict = MessageToDict(wallet_pb, including_default_value_fields=including_default_value_fields)
  settings = wallet_dict['preferences']['shared']['value']['value']['settings']
  settings['lbryum_servers'] = [
    [domain_port_pair['domain'], domain_port_pair['port']]
    for domain_port_pair
    in settings['lbryum_servers']
  ]
  return json.dumps(wallet_dict)
