/******** AUTHOR: NAGA SAI AAKARSHIT BATCHU ********/

package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "./iot"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
)

var (
	tls       = flag.Bool("tls", true, "Connection uses TLS if true, else plain TCP")
	certFile  = flag.String("cert_file", "", "The TLS cert file")
	keyFile   = flag.String("key_file", "", "The TLS key file")
	port      = flag.Int("port", 7771, "The server port")
	fbolddata []byte
	mbolddata []byte
)

type server struct{}

type beatError struct {
	er string
}

func (e *beatError) Error() string {
	return fmt.Sprintf("%s", e.er)
}

func UpperCase(str string) string {
	var upstr string = strings.ToUpper(str)
	return upstr
}

func LowerCase(str string) string {
	var lowstr string = strings.ToLower(str)
	return lowstr
}

func Trim(str, cutset string) string {
	var trimstr string = strings.Trim(str, cutset)
	return trimstr
}

func GenerateCode(str1, str2, str3 string) string {
	var codeout string = str1 + str2 + str3
	return codeout
}

func Execute(command, execbeat, execaction string) ([]byte, error) {
	out, err := exec.Command("sudo", command, execbeat, execaction).Output()
	return out, err
}

func (s *server) Beat(ctx context.Context, in *pb.Config) (*pb.Response, error) {

	var (
		response      string
		beattype      string = UpperCase(in.Beat)
		executebeat   string = LowerCase(beattype)
		action        string = UpperCase(in.Action)
		executeaction string = LowerCase(action)
		data          []byte = in.Data
		filepath      string
		beatcode      string
		actioncode    string
		errorcode     string
		code          string
	)

	const (
		sudo         string      = "sudo"
		commands     string      = "service"
		actionerrmsg string      = "failed to perform the action on beat"
		writeerrmsg  string      = "failed to write data to the file"
		permission   os.FileMode = 0644
		zero         string      = "0"
		one          string      = "1"
		two          string      = "2"
		three        string      = "3"
		four         string      = "4"
		five         string      = "5"
		six          string      = "6"
		seven        string      = "7"
		eight        string      = "8"
		nine         string      = "9"
	)

	if beattype == "FILEBEAT" {
		filepath = "/etc/filebeat/filebeat.yml"
		beatcode = one
	} else if beattype == "METRICBEAT" {
		filepath = "/etc/metricbeat/metricbeat.yml"
		beatcode = two
	}

	if (beattype == "FILEBEAT") || (beattype == "METRICBEAT") {

		if action == "START" {
			actioncode = one
			writeerr := ioutil.WriteFile(filepath, data, permission)
			if writeerr != nil {
				errorcode = one
				code = GenerateCode(beatcode, actioncode, errorcode)
				log.Printf("%s: %v", "Failed to write Data to File", writeerr)
				return &pb.Response{Name: executebeat, Action: executeaction, Code: code, Result: "failed", Message: writeerrmsg}, nil
			} else {
				startout, starterr := Execute(commands, executebeat, executeaction)
				_ = startout
				response = "Started " + executebeat
				if starterr != nil {
					errorcode = two
					code = GenerateCode(beatcode, actioncode, errorcode)
					log.Printf("Failed to %s %s Code: %s Error: %v", executeaction, executebeat, code, starterr)
					return &pb.Response{Name: executebeat, Action: executeaction, Code: code, Result: "failed", Message: actionerrmsg}, nil
				} else {
					errorcode = zero
					code = GenerateCode(beatcode, actioncode, errorcode)
					log.Printf("%s, Code: %s", response, code)
					if beattype == "FILEBEAT" {
						fbolddata = data
					} else if beattype == "METRICBEAT" {
						mbolddata = data
					}
				}
			}

		} else if action == "STOP" {
			actioncode = two
			writeerr := ioutil.WriteFile(filepath, []byte(""), permission)
			if writeerr != nil {
				errorcode = one
				code = GenerateCode(beatcode, actioncode, errorcode)
				log.Printf("%s: %v", "Failed to reset the File", writeerr)
				return &pb.Response{Name: executebeat, Action: executeaction, Code: code, Result: "failed", Message: writeerrmsg}, nil
			} else {
				stopout, stoperr := Execute(commands, executebeat, executeaction)
				_ = stopout
				response = "Stopped " + executebeat
				if stoperr != nil {
					errorcode = two
					code = GenerateCode(beatcode, actioncode, errorcode)
					log.Printf("Failed to %s %s Code: %s Error: %v", executeaction, executebeat, code, stoperr)
					return &pb.Response{Name: executebeat, Action: executeaction, Code: code, Result: "failed", Message: actionerrmsg}, nil
				} else {
					errorcode = zero
					code = GenerateCode(beatcode, actioncode, errorcode)
					log.Printf("%s, Code: %s", response, code)
				}
			}

		} else if action == "STATUS" {
			actioncode = three
			statusout, statuserr := Execute(commands, executebeat, executeaction)
			response = Trim(string(statusout), " * ")
			response = strings.TrimSpace(response)
			errorcode = zero
			code = GenerateCode(beatcode, actioncode, errorcode)
			log.Printf("%s, Code: %s", response, code)
			_ = statuserr

		} else if action == "PAUSE" {
			actioncode = four
			pauseout, pauseerr := Execute(commands, executebeat, "stop")
			_ = pauseout
			response = "Paused " + executebeat
			if pauseerr != nil {
				errorcode = one
				code = GenerateCode(beatcode, actioncode, errorcode)
				log.Printf("Failed to %s %s Code: %s Error: %v", executeaction, executebeat, code, pauseerr)
				return &pb.Response{Name: executebeat, Action: executeaction, Code: code, Result: "failed", Message: actionerrmsg}, nil
			} else {
				errorcode = zero
				code = GenerateCode(beatcode, actioncode, errorcode)
				log.Printf("%s, Code: %s", response, code)
			}

		} else if action == "RESUME" {
			actioncode = five
			resumeout, resumeerr := Execute(commands, executebeat, "start")
			_ = resumeout
			response = "Resumed" + executebeat
			if resumeerr != nil {
				errorcode = one
				code = GenerateCode(beatcode, actioncode, errorcode)
				log.Printf("Failed to %s %s Code: %s Error: %v", executeaction, executebeat, code, resumeerr)
				return &pb.Response{Name: executebeat, Action: executeaction, Code: code, Result: "failed", Message: actionerrmsg}, nil
			} else {
				errorcode = zero
				code = GenerateCode(beatcode, actioncode, errorcode)
				log.Printf("%s, Code: %s", response, code)
			}

		} else if action == "RESTART" {
			actioncode = six
			if beattype == "FILEBEAT" {
				data = fbolddata
			} else if beattype == "METRICBEAT" {
				data = mbolddata
			}

			writeerr := ioutil.WriteFile(filepath, data, permission)
			if writeerr != nil {
				errorcode = one
				code = GenerateCode(beatcode, actioncode, errorcode)
				log.Printf("%s: %v", "Failed to write Data to File", writeerr)
				return &pb.Response{Name: executebeat, Action: executeaction, Code: code, Result: "failed", Message: writeerrmsg}, nil
			} else {
				startout, starterr := Execute(commands, executebeat, "start")
				_ = startout
				response = "Restarted " + executebeat
				if starterr != nil {
					errorcode = two
					code = GenerateCode(beatcode, actioncode, errorcode)
					log.Printf("Failed to %s %s Code: %s Error: %v", executeaction, executebeat, code, starterr)
					return &pb.Response{Name: executebeat, Action: executeaction, Code: code, Result: "failed", Message: actionerrmsg}, nil
				} else {
					errorcode = zero
					code = GenerateCode(beatcode, actioncode, errorcode)
					log.Printf("%s, Code: %s", response, code)
				}
			}

		} else {
			actioncode = nine
			errorcode = zero
			code = GenerateCode(beatcode, actioncode, errorcode)
			response = "Action cannot be performed: " + action
			return &pb.Response{Name: executebeat, Action: executeaction, Code: code, Result: "failed", Message: response}, nil
		}
	} else {
		beatcode = nine
		actioncode = zero
		errorcode = zero
		code = GenerateCode(beatcode, actioncode, errorcode)
		return &pb.Response{Name: executebeat, Action: executeaction, Code: code, Result: "failed", Message: "Cannot Work with that Beat"}, nil
	}
	return &pb.Response{Name: executebeat, Action: executeaction, Code: code, Result: "success", Message: response}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	if *tls {
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			grpclog.Fatalf("Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterIOTServer(grpcServer, &server{})
	grpcServer.Serve(lis)
}

/******** AUTHOR: NAGA SAI AAKARSHIT BATCHU ********/
