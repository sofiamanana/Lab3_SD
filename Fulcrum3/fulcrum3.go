package main

import (
	pb "Lab3_SD/proto"
	"bufio"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
)

type Server4 struct {
	pb.UnimplementedFulcrumServer
}

var Vector = make(map[string][]int32)

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
			rebeldes = textarray[2]
		} else {
			texto += scanner.Text() + "\n"
		}
	}
	return &pb.Numero{Num: rebeldes}, nil
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

func IniciarVector(planeta string) {
	file, err := os.Open(planeta + ".txt")
	if err != nil {
		Vector[planeta] = []int32{0, 0, 0}
	}
	defer file.Close()
	Vector[planeta][2]++
}

func (ahsoka *Server2) AddCity(ctx context.Context, in *pb.Estructura) (*pb.Vector, error) {
	log.Printf("Informante desea crear un planeta de nombre: %s", in.Planeta)
	log.Printf("Con ciudad de nombre: %s", in.Ciudad)
	log.Printf("Con tantos rebeldes: %s", in.Rebeldes)
	//var vector[3]int{0,0,0} ??
	//AgregarCiudad(in.Planeta, in.Ciudad, in.Rebeldes)
	IniciarVector(in.Planeta)
	AgregarCiudad(in.Planeta, in.Ciudad, in.Rebeldes)

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

/*
func Merge(){
	for range time.Tick(time.Minute * 1) {
			log.Printf("Iniciando merge \n")
			fmt.Println(time.Now())
			//aqui meter el lock y todo lo relacionado al merge
			//var m sync.Mutex
			//m.Lock()

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


			//m.Release()

	}
}
*/

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
