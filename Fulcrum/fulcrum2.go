package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	pb "Lab3_SD/proto"
	"google.golang.org/grpc"
)

type Server3 struct {
	pb.UnimplementedInformanteFulcrumServer
}

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

func (ahsoka1 *Server3) AddCity(ctx context.Context, in *pb.Estructura) (*pb.Vector, error) {
	log.Printf("Informante desea crear un planeta de nombre: %s", in.Planeta)
	log.Printf("Con ciudad de nombre: %s", in.Ciudad)
	log.Printf("Con tantos rebeldes: %d", in.Rebeldes)
	//var vector[3]int{0,0,0} ??
	//AgregarCiudad(in.Planeta, in.Ciudad, in.Rebeldes)
	return &pb.Vector{X: 0, Y: 0, Z: 0}, nil
}

func main() {
	//Conexión a Informante Ahsoka
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9070))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	ahsoka1 := grpc.NewServer()

	pb.RegisterInformanteFulcrumServer(ahsoka1, &Server3{})
	if err := ahsoka1.Serve(lis); err != nil {
		log.Fatalf("falló la conexión informante-fulcrum: %s", err)
	}
}
