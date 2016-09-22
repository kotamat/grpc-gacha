this_dir = File.expand_path(File.dirname(__FILE__))
lib_dir = File.join(this_dir, '../lib')
$LOAD_PATH.unshift(lib_dir) unless $LOAD_PATH.include?(lib_dir)

require 'grpc'
require 'gacha_services_pb'

def main
  stub = Gacha::Gacha::Stub.new('localhost:50055', :this_channel_is_insecure)
  cards = [
      Gacha::Card.new(name: "card1"),
      Gacha::Card.new(name: "card2"),
  ]

  begin
  res = stub.lottery(Gacha::Request.new(cards: cards))
  rescue => e
    p e
  end
  if res.ret_code == 1
    p "gained: #{res.card.name}"
  else
    p "error"
  end
end

main
