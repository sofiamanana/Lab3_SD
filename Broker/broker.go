package main

import (
	pb "Lab3_SD/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
	"strings"
	"time"
)

type Server struct {
	pb.UnimplementedBrokerServer
}

func IrFulcrum1(planetaciudad string) (retorno string, X int32, Y int32, Z int32) {
	var conn1 *grpc.ClientConn //FULCRUM 1
	conn1, err1 := grpc.Dial("10.6.40.169:9060", grpc.WithInsecure())
	if err1 != nil {
		log.Fatalf("did not connect: %s", err1)
	}
	defer conn1.Close()
	fulcrum1 := pb.NewFulcrumClient(conn1)

	response, err := fulcrum1.PreguntarInformantes(context.Background(), &pb.PlanetaCiudad{Body: planetaciudad})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	retorno = response.Body
	X = response.X
	Y = response.Y
	Z = response.Z
	return retorno, X, Y, Z
}

func IrFulcrum2(planetaciudad string) (retorno string, X int32, Y int32, Z int32) {
	var conn2 *grpc.ClientConn //FULCRUM 2
	conn2, err2 := grpc.Dial("10.6.40.170:9070", grpc.WithInsecure())
	if err2 != nil {
		log.Fatalf("did not connect: %s", err2)
	}
	defer conn2.Close()
	fulcrum2 := pb.NewFulcrumClient(conn2)

	response, err := fulcrum2.PreguntarInformantes(context.Background(), &pb.PlanetaCiudad{Body: planetaciudad})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	retorno = response.Body
	X = response.X
	Y = response.Y
	Z = response.Z
	return retorno, X, Y, Z
}

func IrFulcrum3(planetaciudad string) (retorno string, X int32, Y int32, Z int32) {
	var conn3 *grpc.ClientConn //FULCRUM 3
	conn3, err3 := grpc.Dial("10.6.40.171:9040", grpc.WithInsecure())
	if err3 != nil {
		log.Fatalf("did not connect: %s", err3)
	}
	defer conn3.Close()
	fulcrum3 := pb.NewFulcrumClient(conn3)

	response, err := fulcrum3.PreguntarInformantes(context.Background(), &pb.PlanetaCiudad{Body: planetaciudad})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	retorno = response.Body
	X = response.X
	Y = response.Y
	Z = response.Z
	return retorno, X, Y, Z
}

func (s *Server) GetNumberRebels(ctx context.Context, in *pb.PlanetaCiudad) (*pb.Vect, error) {
	split := strings.Split(in.Body, ",")
	planeta := split[0]
	ciudad := split[1]
	var ip string
	var numero string
	var ful int32 = 0
	var X, Y, Z int32
	log.Printf("Leia pregunto por el planeta %s y la ciudad %s", planeta, ciudad)

	//tirar al azar server:
	ful = rand.Int31n(3) + 1 //Se escoge un numero al azar entre 1 y 3 (corresponden a los 3 fulcrum)
	if ful == 1 {
		ip = "10.6.40.169" //ip del dist29 10.6.40.169
		numero, X, Y, Z = IrFulcrum1(in.Body)
	} else if ful == 2 { //ip del dist30 10.6.40.170
		ip = "10.6.40.170"
		numero, X, Y, Z = IrFulcrum2(in.Body)

	} else { //ip del dist31 10.6.40.171
		ip = "10.6.40.171"
		numero, X, Y, Z = IrFulcrum3(in.Body)
	}

	return &pb.Vect{X: X, Y: Y, Z: Z, Body: numero, Ip: ip}, nil
}

func (inf *Server) QuieroHacer(ctx context.Context, in *pb.Comando) (*pb.Redirigido, error) {
	var fulcrum int32 = 0
	var ip string
	log.Printf("El informante desea hacer: %s", in.Comando)
	fulcrum = rand.Int31n(3) + 1 //Se escoge un numero al azar entre 1 y 3 (corresponden a los 3 fulcrum)
	if fulcrum == 1 {            //ip del dist29 10.6.40.169
		ip = "10.6.40.169"
		return &pb.Redirigido{Valor: ip}, nil
	} else if fulcrum == 2 { //ip del dist30 10.6.40.170
		ip = "10.6.40.170"
		return &pb.Redirigido{Valor: ip}, nil
	} else { //ip del dist31 10.6.40.171
		ip = "10.6.40.171"
		return &pb.Redirigido{Valor: ip}, nil
	}
}

func ServerInformante() { //Conexión para conectar este broker (servidor) al informante (cliente)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9050))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	inf := grpc.NewServer()

	pb.RegisterBrokerServer(inf, &Server{})
	if err := inf.Serve(lis); err != nil {
		log.Fatalf("falló la conexión informante-broker: %s", err)
	}
}

func main() {
	log.Printf("Servidor Broker iniciado. \n")
	rand.Seed(time.Now().UnixNano()) //para seleccionar el server fulcrum al azar
	ServerInformante()
}
