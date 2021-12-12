package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	//"strings"

	pb "Lab3_SD/proto"
	"google.golang.org/grpc"
)

type Server3 struct {
	pb.UnimplementedBrokerServer
}

/*
func AgregarCiudad(nombre_planeta string, nombre_ciudad string, nuevo_valor int) {
	content,err:=ioutill.ReadFile(filename:nombre_planeta+".txt")
	if err!=nil{
		err2:=ioutill.WriteFile(nombre_planeta+".txt",nombre_planeta+" "+nombre_ciudad+" "+nuevo_valor,0644)
	}
	content=append(content,([]byte(nombre_planeta)+[]byte(" ")+[]byte(nombre_ciudad)+[]byte(" ")+[]byte(nuevo_valor))...,0644)
	err=ioutill.WriteFile(nombre_planeta+".txt",content,0644)
	if err!=nil{
		log.Fatal(err)
	}
}
*/
/*
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

func (ahsoka1 *Server3) AddCity(ctx context.Context, in *pb.Estructura) (*pb.Vector, error) {
	log.Printf("Informante desea crear un planeta de nombre: %s", in.Planeta)
	log.Printf("Con ciudad de nombre: %s", in.Ciudad)
	log.Printf("Con tantos rebeldes: %d", in.Rebeldes)
	//var vector[3]int{0,0,0} ??
	//AgregarCiudad(in.Planeta, in.Ciudad, in.Rebeldes)
	return &pb.Vector{X: "0", Y: "0", Z: "0"}, nil
}

func main() {
	//Conexión a Informante Ahsoka
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9070))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	ahsoka1 := grpc.NewServer()

	pb.RegisterBrokerServer(ahsoka1, &Server3{})
	if err := ahsoka1.Serve(lis); err != nil {
		log.Fatalf("falló la conexión informante-fulcrum: %s", err)
	}
}
