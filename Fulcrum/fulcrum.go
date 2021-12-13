package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"io/ioutil"
	"bufio"
	"time"
	pb "Lab3_SD/proto"
	"google.golang.org/grpc"
	"strings"
	"sync"
)

var Vector = make(map[string][]int)

type Server2 struct {
	pb.UnimplementedFulcrumServer
}

func AgregarCiudad(nombre_planeta string, nombre_ciudad string, nuevo_valor string) {
	content, err := ioutil.ReadFile(nombre_planeta + ".txt")
	if err != nil {
		ioutil.WriteFile(nombre_planeta+".txt", ([]byte(nombre_planeta + " " + nombre_ciudad + " " + nuevo_valor + "\n")), 0644)
	} else {
		content = append(content, ([]byte(nombre_planeta + " " + nombre_ciudad + " " + nuevo_valor + "\n"))...)
		err = ioutil.WriteFile(nombre_planeta+".txt", content, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
	content2, err2 := ioutil.ReadFile("Log" + nombre_planeta + ".txt")
	if err2 != nil {
		ioutil.WriteFile("Log"+nombre_planeta+".txt", ([]byte("AddCitty " + nombre_planeta + " " + nombre_ciudad + " " + nuevo_valor + "\n")), 0644)
	} else {
		content2 = append(content2, ([]byte("AddCitty " + nombre_planeta + " " + nombre_ciudad + " " + nuevo_valor + "\n"))...)
		err = ioutil.WriteFile("Log"+nombre_planeta+".txt", content2, 0644)
		if err2 != nil {
			log.Fatal(err2)
		}
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
	content2, err2 := ioutil.ReadFile("Log" + nombre_planeta + ".txt")
	if err2 != nil {
		ioutil.WriteFile("Log"+nombre_planeta+".txt", ([]byte("UpdateName " + nombre_planeta + " " + nombre_ciudad + " " + nuevo_valor + "\n")), 0644)
	} else {
		content2 = append(content2, ([]byte("UpdateName " + nombre_planeta + " " + nombre_ciudad + " " + nuevo_valor + "\n"))...)
		err = ioutil.WriteFile("Log"+nombre_planeta+".txt", content2, 0644)
		if err2 != nil {
			log.Fatal(err2)
		}
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
	content2, err2 := ioutil.ReadFile("Log" + nombre_planeta + ".txt")
	if err2 != nil {
		ioutil.WriteFile("Log"+nombre_planeta+".txt", ([]byte("UpdateNumber " + nombre_planeta + " " + nombre_ciudad + " " + nuevo_valor + "\n")), 0644)
	} else {
		content2 = append(content2, ([]byte("UpdateNumber " + nombre_planeta + " " + nombre_ciudad + " " + nuevo_valor + "\n"))...)
		err = ioutil.WriteFile("Log"+nombre_planeta+".txt", content2, 0644)
		if err2 != nil {
			log.Fatal(err2)
		}
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
	content2, err2 := ioutil.ReadFile("Log" + nombre_planeta + ".txt")
	if err2 != nil {
		ioutil.WriteFile("Log"+nombre_planeta+".txt", ([]byte("DeleteCity " + nombre_planeta + " " + nombre_ciudad +  "\n")), 0644)
	} else {
		content2 = append(content2, ([]byte("DeleteCity " + nombre_planeta + " " + nombre_ciudad + "\n"))...)
		err = ioutil.WriteFile("Log"+nombre_planeta+".txt", content2, 0644)
		if err2 != nil {
			log.Fatal(err2)
		}
	}
}


func (ahsoka *Server2) AddCity(ctx context.Context, in *pb.Estructura) (*pb.Vector, error) {
	log.Printf("Informante desea crear un planeta de nombre: %s", in.Planeta)
	log.Printf("Con ciudad de nombre: %s", in.Ciudad)
	log.Printf("Con tantos rebeldes: %s", in.Rebeldes)
	//var vector[3]int{0,0,0} ??
	//AgregarCiudad(in.Planeta, in.Ciudad, in.Rebeldes)
	AgregarCiudad(in.Planeta, in.Ciudad, in.Rebeldes)
	Vector[in.Planeta] = []int{0,0,0}
	Vector[in.Planeta][0]++
	return &pb.Vector{X: Vector[in.Planeta][0], Y: Vector[in.Planeta][1], Z: Vector[in.Planeta][2]}, nil
}

func Merge(){
	for range time.Tick(time.Minute * 1) {
			log.Printf("Iniciando merge \n")
			fmt.Println(time.Now())
			//aqui meter el lock y todo lo relacionado al merge
			//var m sync.Mutex
			//m.Lock()
			/*
			var conn *grpc.ClientConn
			conn, err := grpc.Dial("10.6.40.170:9070", grpc.WithInsecure())
			if err != nil {
				log.Fatalf("did not connect: %s", err)
			}
			defer conn.Close()

			c := pb.NewFulcrumClient(conn)
			//recorrer llaves:
			//cada llave es un planeta

			for k, v := range Vector { 
				//fmt.Printf("key[%s] value[%s]\n", k, v)
				response, err := c.Mergecito12(context.Background(), &pb.PlanetaCiudad{Body: k})
				if err != nil {
					log.Fatalf("Error when calling SayHello: %s", err)
				}
				log.Printf("Respuesta del Fulcrum 2: Vector para %s es %d, %d, %d \n",k,response.X,response.Y,response.Z)
				}
			

			//m.Release()*/
			
	}
}


func ConexionServer(){ //Conexión a Informante Ahsoka
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

func main() {
	//Conexión a Informante Ahsoka
	Vector["Chilito"] = []int{0,0,0}

	go ConexionServer()
	go Merge()
	
}
