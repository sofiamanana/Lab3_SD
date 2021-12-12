package main

import (
	"context"
	"fmt"
	//"io/ioutil"
	"log"
	"net"
	//"strings"

	pb "Lab3_SD/proto"
	"google.golang.org/grpc"
)

type Server3 struct {
	pb.UnimplementedFulcrumServer
}

func AddCity(nombre_planeta string, nombre_ciudad string, nuevo_valor string) {
	content, err := ioutil.ReadFile(nombre_planeta + ".txt")
	if err != nil {
		ioutil.WriteFile(nombre_planeta+".txt", ([]byte(nombre_planeta + " " + nombre_ciudad + " " + nuevo_valor + "\n")), 0644)
		return
	}
	content = append(content, ([]byte(nombre_planeta + " " + nombre_ciudad + " " + nuevo_valor + "\n"))...)
	err = ioutil.WriteFile(nombre_planeta+".txt", content, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
func UpdateName(nombre_planeta string, nombre_ciudad string, nuevo_valor string) {
	file, err := os.Open(nombre_planeta + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var texto string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		textarray := strings.Split(scanner.Text(), " ")
		if textarray[1] == nombre_ciudad {
			texto += string(textarray[0] + " " + nuevo_valor + " " + textarray[2] + "\n")
		} else {
			texto += scanner.Text() + "\n"
		}
	}
	log.Println(texto)
	err = ioutil.WriteFile(nombre_planeta+".txt", []byte(texto), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateNumber(nombre_planeta string, nombre_ciudad string, nuevo_valor string) {
	file, err := os.Open(nombre_planeta + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var texto string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		textarray := strings.Split(scanner.Text(), " ")
		if textarray[1] == nombre_ciudad {
			texto += string(textarray[0] + " " + textarray[1] + " " + nuevo_valor + "\n")
		} else {
			texto += scanner.Text() + "\n"
		}
	}
	err = ioutil.WriteFile(nombre_planeta+".txt", []byte(texto), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteCity(nombre_planeta string, nombre_ciudad string) {
	file, err := os.Open(nombre_planeta + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var texto string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		textarray := strings.Split(scanner.Text(), " ")
		if textarray[1] == nombre_ciudad {
			continue
		} else {
			texto += scanner.Text() + "\n"
		}
	}
	err = ioutil.WriteFile(nombre_planeta+".txt", []byte(texto), 0644)
	if err != nil {
		log.Fatal(err)
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

	pb.RegisterFulcrumServer(ahsoka1, &Server3{})
	if err := ahsoka1.Serve(lis); err != nil {
		log.Fatalf("falló la conexión informante-fulcrum: %s", err)
	}
}
