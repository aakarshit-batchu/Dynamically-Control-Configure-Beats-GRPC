/******** AUTHOR: NAGA SAI AAKARSHIT BATCHU ********/

var PROTO_PATH = './iot/iot.proto';

var grpc = require('grpc');
var fs = require('fs');
var flags = require('flags');

var iot = grpc.load(PROTO_PATH).iot;

function main() {

  flags.defineString('beat','beat','Beat-Type');
  flags.defineString('action','nothing','Action to be Performed');
  flags.defineString('cert_file','','Public Key');
  flags.defineString('config_path','','Configuration Path');
  flags.defineString('addr','localhost:7771','Host Address and Port Where the beatgrpc Service is Running');
  flags.parse()

  beattype = flags.get('beat')
  actiontobe = flags.get('action')
  certFile = flags.get('cert_file')
  configPath = flags.get('config_path')
  beatgrpcserviceaddress = flags.get('addr')
  var root_certs = fs.readFileSync(certFile);
  var ssl_creds = grpc.credentials.createSsl(root_certs);
  var options = {
  'grpc.ssl_target_name_override' : 'beat',	//The Common Name or Server Host Override Name you give while generating a self-signed public key.
  'grpc.default_authority': 'beat'		//The Common Name or Server Host Override Name you give while generating a self-signed public key.
  };
  var client = new iot.IOT(beatgrpcserviceaddress,ssl_creds,options)
  var yml = fs.readFileSync(configPath)
  client.beat({beat: beattype , action: actiontobe , data: yml }, function(err, response) {
    console.log(response);
  });
}

main();

/******** AUTHOR: NAGA SAI AAKARSHIT BATCHU ********/
