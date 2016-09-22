this_dir = File.expand_path(File.dirname(__FILE__))
lib_dir = File.join(this_dir, '../lib')
$LOAD_PATH.unshift(lib_dir) unless $LOAD_PATH.include?(lib_dir)

require 'grpc'
require 'gacha_services_pb'

class GachaServer < Gacha::Gacha::Service
  def lottery(req, _unused_call)
     Gacha::Response.new(card: req.cards.sample, ret_code: 1)
  end
end

def main
  s = GRPC::RpcServer.new
  s.add_http2_port('0.0.0.0:50055', :this_port_is_insecure)
  s.handle(GachaServer)
  s.run_till_terminated
end

main
