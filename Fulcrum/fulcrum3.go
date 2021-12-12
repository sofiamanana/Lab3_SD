package main

import (
	"context"
	"fmt"
	"log"
	"net"
	//"os"

	pb "Lab3_SD/proto"
	"google.golang.org/grpc"
)

type Server4 struct {
	pb.UnimplementedBrokerServer
}

func (s *Server) PreguntarInformantes(ctx context.Context, in *pb.PlanetaCiudad) (*pb.Numero, error) {
	split := strings.Split(in.Body, ",")
	planeta := split[0]
	ciudad := split[1]
	log.Printf("Broker pregunto por el planeta %s y la ciudad %s", planeta, ciudad)
	return &pb.Numero{Num: 5}, nil
}

/*
func AgregarCiudad(nombre_planeta string, nombre_ciudad string, nuevo_valor int) {
	file, err := os.Open(nombre_planeta + ".txt")
	if err != nil {
		file, err := os.Create(nombre_planeta + ".txt")
		if err != nil {
			log.Fatal(err)
		}
	}
	defer file.Close()
	_, err2 := file.WriteString(nombre_planeta + " " + nombre_ciudad + " " + nuevo_valor + "/n")
	if err2 != nil {
		log.Fatal(err2)
	}
	//fmt.Println("listo")
}
func UpdateName(nombre_planeta string, nombre_ciudad string, nuevo_valor string) {
	file, err := os.Open(nombre_planeta + "txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		fmt.Println(scanner.Bytes())
	}
}
*/

func (ahsoka2 *Server4) AddCity(ctx context.Context, in *pb.Estructura) (*pb.Vector, error) {
	log.Printf("Informante desea crear un planeta de nombre: %s", in.Planeta)
	log.Printf("Con ciudad de nombre: %s", in.Ciudad)
	log.Printf("Con tantos rebeldes: %d", in.Rebeldes)
	//var vector[3]int{0,0,0} ??
	//AgregarCiudad(in.Planeta, in.Ciudad, in.Rebeldes)
	return &pb.Vector{X: "0", Y: "0", Z: "0"}, nil
}

func main() {
	//Conexión a Informante Ahsoka
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9040))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	ahsoka2 := grpc.NewServer()

	pb.RegisterBrokerServer(ahsoka2, &Server4{})
	if err := ahsoka2.Serve(lis); err != nil {
		log.Fatalf("fallo la conexion informante-fulcrum: %s", err)
	}
}
