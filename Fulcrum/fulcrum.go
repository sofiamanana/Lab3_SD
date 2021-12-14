package main

import (
	"bufio"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	//"time"
	pb "Lab3_SD/proto"
	"google.golang.org/grpc"
	"strings"
	//"sync"
)

var Vector = make(map[string][]int32)

type Server2 struct {
	pb.UnimplementedFulcrumServer
}

func (s *Server2) PreguntarInformantes(ctx context.Context, in *pb.PlanetaCiudad) (*pb.Vectorcito, error) {
	split := strings.Split(in.Body, ",")
	planeta := split[0]
	ciudad := split[1]
	var rebeldes string
	var flag int32 = 0
	log.Printf("Broker pregunto por el planeta %s y la ciudad %s", planeta, ciudad)
	//leer archivo
	file, err := os.Open(planeta + ".txt")
	if err != nil {
		flag = 1
		log.Printf("No existe ese planeta\n")
	}
	defer file.Close()
	var texto string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		textarray := strings.Split(scanner.Text(), " ")
		if textarray[1] == ciudad {
			rebeldes = textarray[2]
		} else {
			texto += scanner.Text() + "\n"
		}
	}
	if flag == 1 {
		Vector[planeta] = []int32{0, 0, 0}
		rebeldes = "0"
	}
	return &pb.Vectorcito{X: Vector[planeta][0], Y: Vector[planeta][1], Z: Vector[planeta][2], Body: rebeldes}, nil
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

func ActualizarNombre(nombre_planeta string, nombre_ciudad string, nuevo_valor string) (flag int) {
	file, err := os.Open(nombre_planeta + ".txt")
	flag = 0
	if err != nil {
		//log.Fatal(err)
		log.Printf("El planeta no existe en este Fulcrum")
		flag = 1
		return flag
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
	return flag
}

func ActualizarNumero(nombre_planeta string, nombre_ciudad string, nuevo_valor string) (flag int) {
	file, err := os.Open(nombre_planeta + ".txt")
	flag = 0
	if err != nil {
		//log.Fatal(err)
		log.Printf("El planeta no existe en este Fulcrum")
		flag = 1
		return flag
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
	return flag
}

func EliminarCiudad(nombre_planeta string, nombre_ciudad string) (flag int) {
	file, err := os.Open(nombre_planeta + ".txt")
	flag = 0
	if err != nil {
		//log.Fatal(err)
		log.Printf("El planeta no existe en este Fulcrum")
		flag = 1
		return flag
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
		content2 = append(content2, ([]byte("DeleteCity " + nombre_planeta + " " + nombre_ciudad + "\n"))...)
		err = ioutil.WriteFile("Log"+nombre_planeta+".txt", content2, 0644)
		if err2 != nil {
			log.Fatal(err2)
		}
	}
	return flag
}

func LeerArchivo(nombre_planeta string){
	file, err := os.Open(nombre_planeta + ".txt")
	var textarray []string
	if err != nil {
		//log.Fatal(err)
		log.Printf("El planeta no existe en este Fulcrum")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		textarray = strings.Split(scanner.Text(), " ")
	}
	log.Printf("%s\n",textarray)
}

func IniciarVector(planeta string) {
	file, err := os.Open(planeta + ".txt")
	if err != nil {
		Vector[planeta] = []int32{0, 0, 0}
	}
	defer file.Close()
	Vector[planeta][0]++
}

func (ahsoka *Server2) AddCity(ctx context.Context, in *pb.Estructura) (*pb.Vector, error) {
	log.Printf("Informante desea crear un planeta de nombre: %s", in.Planeta)
	log.Printf("Con ciudad de nombre: %s", in.Ciudad)
	log.Printf("Con tantos rebeldes: %s", in.Rebeldes)
	//var vector[3]int{0,0,0} ??
	//AgregarCiudad(in.Planeta, in.Ciudad, in.Rebeldes)
	IniciarVector(in.Planeta)
	AgregarCiudad(in.Planeta, in.Ciudad, in.Rebeldes)
	LeerArchivo(in.Planeta)
	return &pb.Vector{X: Vector[in.Planeta][0], Y: Vector[in.Planeta][1], Z: Vector[in.Planeta][2]}, nil
}

func (ahsoka *Server2) UpdateName(ctx context.Context, in *pb.Estructura) (*pb.Vector, error) {
	log.Printf("Informante desea cambiar el nombre de una ciudad en el planeta: %s", in.Planeta)
	log.Printf("La ciudad a cambiar es : %s", in.Ciudad)
	log.Printf("El nuevo nombre de la ciudad es: %s", in.Rebeldes)
	//var vector[3]int{0,0,0} ??
	//AgregarCiudad(in.Planeta, in.Ciudad, in.Rebeldes)
	flag := ActualizarNombre(in.Planeta, in.Ciudad, in.Rebeldes)
	if flag == 0 {
		IniciarVector(in.Planeta)
	} else {
		Vector[in.Planeta] = []int32{0, 0, 0}
	}
	//Vector[in.Planeta] = []int32{0,0,0}
	//Vector[in.Planeta][0]++
	LeerArchivo(in.Planeta)
	return &pb.Vector{X: Vector[in.Planeta][0], Y: Vector[in.Planeta][1], Z: Vector[in.Planeta][2]}, nil
	//return &pb.Vector{X: 0, Y: 0, Z: 0}, nil
}

func (ahsoka *Server2) UpdateNumber(ctx context.Context, in *pb.Estructura) (*pb.Vector, error) {
	log.Printf("Informante desea cambiar numero del planeta: %s", in.Planeta)
	log.Printf("El numero antiguo de rebeldes es: %s", in.Ciudad)
	log.Printf("El nuevo numero es: %s", in.Rebeldes)
	//var vector[3]int{0,0,0} ??
	//AgregarCiudad(in.Planeta, in.Ciudad, in.Rebeldes)
	flag := ActualizarNumero(in.Planeta, in.Ciudad, in.Rebeldes)
	if flag == 0 {
		IniciarVector(in.Planeta)
	} else {
		Vector[in.Planeta] = []int32{0, 0, 0}
	}
	//Vector[in.Planeta] = []int32{0,0,0}
	//Vector[in.Planeta][0]++
	LeerArchivo(in.Planeta)
	return &pb.Vector{X: Vector[in.Planeta][0], Y: Vector[in.Planeta][1], Z: Vector[in.Planeta][2]}, nil

	//return &pb.Vector{X: 0, Y: 0, Z: 0}, nil
}

func (ahsoka *Server2) DeleteCity(ctx context.Context, in *pb.Estructura3) (*pb.Vector, error) {
	log.Printf("Se desea eliminar una ciudad del planeta: %s", in.Planeta)
	log.Printf("La ciudad a eliminar es: %s", in.Ciudad)
	//var vector[3]int{0,0,0} ??
	//AgregarCiudad(in.Planeta, in.Ciudad, in.Rebeldes)
	flag := EliminarCiudad(in.Planeta, in.Ciudad)
	if flag == 0 {
		IniciarVector(in.Planeta)
	} else {
		Vector[in.Planeta] = []int32{0, 0, 0}
	}

	return &pb.Vector{X: Vector[in.Planeta][0], Y: Vector[in.Planeta][1], Z: Vector[in.Planeta][2]}, nil
	//return &pb.Vector{X: 0, Y: 0, Z: 0}, nil
}


func Merge(){
	log.Printf("Iniciando Merge\n")
	for k, v := range Vector{
		log.Printf("key[%s]\n", k)
		log.Printf("Vector %d, %d, %d\n", v[0],v[1],v[2])
	}
}


func ConexionServer() { //Conexión a Informante Ahsoka
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
	log.Printf("Iniciando Fulcrum 1")
	Vector["Chilito"] = []int32{0, 0, 0}

	go ConexionServer()
	go Merge()

}
