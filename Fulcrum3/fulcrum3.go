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

type Server4 struct {
	pb.UnimplementedFulcrumServer
}

func (s *Server4) PreguntarInformantes(ctx context.Context, in *pb.PlanetaCiudad) (*pb.Numero, error) {
	split := strings.Split(in.Body, ",")
	planeta := split[0]
	ciudad := split[1]
	var rebeldes string
	log.Printf("Broker pregunto por el planeta %s y la ciudad %s", planeta, ciudad)
	//leer archivo
	file, err := os.Open(planeta + ".txt")
	if err != nil {
		log.Fatal(err)
		log.Printf("No existe ese planeta\n")
	}
	defer file.Close()
	var texto string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		textarray := strings.Split(scanner.Text(), " ")
		if textarray[1] == ciudad {
			rebeldes := textarray[2]
		} else {
			texto += scanner.Text() + "\n"
		}
	}
	return &pb.Numero{Num: int(rebeldes)}, nil
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
		ioutil.WriteFile("Log"+nombre_planeta+".txt", ([]byte("UpdateNumber " + nombre_planeta + " " + nombre_ciudad +  "\n")), 0644)
	} else {
		content2 = append(content2, ([]byte("UpdateNumber " + nombre_planeta + " " + nombre_ciudad +"\n"))...)
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
		ioutil.WriteFile("Log"+nombre_planeta+".txt", ([]byte("DeleteCity " + nombre_planeta + " " + nombre_ciudad + "\n")), 0644)
	} else {
		content2 = append(content2, ([]byte("DeleteCity " + nombre_planeta + " " + nombre_ciudad +  "\n"))...)
		err = ioutil.WriteFile("Log"+nombre_planeta+".txt", content2, 0644)
		if err2 != nil {
			log.Fatal(err2)
		}
	}
}


func (ahsoka1 *Server4) AddCity(ctx context.Context, in *pb.Estructura) (*pb.Vector, error) {
	log.Printf("Informante desea crear un planeta de nombre: %s", in.Planeta)
	log.Printf("Con ciudad de nombre: %s", in.Ciudad)
	log.Printf("Con tantos rebeldes: %s", in.Rebeldes)
	//var vector[3]int{0,0,0} ??
	//AgregarCiudad(in.Planeta, in.Ciudad, in.Rebeldes)
	return &pb.Vector{X: 0, Y: 0, Z: 0}, nil
}



func main() {
	//Conexión a Leia
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9040))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	leia := grpc.NewServer()

	pb.RegisterFulcrumServer(leia, &Server4{})
	if err := leia.Serve(lis); err != nil {
		log.Fatalf("fallo la conexion informante-fulcrum: %s", err)
	}

}
