package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"io/ioutil"
	"bufio"
	pb "Lab3_SD/proto"
	"google.golang.org/grpc"
	"strings"
)

var Vector = make(map[string][]int32)

type Server2 struct {
	pb.UnimplementedFulcrumServer
}

func AgregarCiudad(nombre_planeta string, nombre_ciudad string, nuevo_valor string) {
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


func (ahsoka *Server2) AddCity(ctx context.Context, in *pb.Estructura) (*pb.Vector, error) {
	log.Printf("Informante desea crear un planeta de nombre: %s", in.Planeta)
	log.Printf("Con ciudad de nombre: %s", in.Ciudad)
	log.Printf("Con tantos rebeldes: %s", in.Rebeldes)
	//var vector[3]int{0,0,0} ??
	//AgregarCiudad(in.Planeta, in.Ciudad, in.Rebeldes)
	AgregarCiudad(in.Planeta, in.Ciudad, in.Rebeldes)
	Vector[in.Planeta] = []int32{0,0,0}
	Vector[in.Planeta][0]++
	return &pb.Vector{X: Vector[in.Planeta][0], Y: Vector[in.Planeta][1], Z: Vector[in.Planeta][2]}, nil
}


func main() {
	//Conexión a Informante Ahsoka
	Vector["Chilito"] = []int32{0,0,0}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9060))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	ahsoka := grpc.NewServer()

	pb.RegisterFulcrumServer(ahsoka, &Server2{})
	if err := ahsoka.Serve(lis); err != nil {
		log.Fatalf("falló la conexión informante-fulcrum: %s", err)
	}
}
