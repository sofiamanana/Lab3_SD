package main

import (
	"context"
	"fmt"
	"log"

	pb "Lab3_SD/proto"
	"google.golang.org/grpc"
)

func AddCiudad() (planeta string, ciudad string, rebeldes string) {
	log.Printf("¿Cuál es el nombre del planeta?:\n")
	fmt.Scan(&planeta)
	log.Printf("¿Cuál es el nombre de la ciudad?:\n")
	fmt.Scan(&ciudad)
	log.Printf("¿Cuántos rebeldes?:\n")
	fmt.Scan(&rebeldes)
	return
}

func UpdateCiudad() (planeta string, ciudad string, nueva_city string) {
	log.Printf("¿Cuál es el nombre del planeta que contiene la ciudad a cambiar?:\n")
	fmt.Scan(&planeta)
	log.Printf("¿Cuál es el nombre de la ciudad a cambiar?:\n")
	fmt.Scan(&ciudad)
	log.Printf("¿Cuál es el nuevo nombre de la ciudad a cambiar?:\n")
	fmt.Scan(&nueva_city)
	return
}

func UpdateRebeldes() (planeta string, ciudad string, new_rebeldes string) {
	log.Printf("¿Cuál es el nombre del planeta que contiene la cantidad de rebeldes a cambiar?:\n")
	fmt.Scan(&planeta)
	log.Printf("¿En qué ciudad?:\n")
	fmt.Scan(&ciudad)
	log.Printf("¿Cuál es el nuevo número de rebeldes en la ciudad?:\n")
	fmt.Scan(&new_rebeldes)
	return
}

func DeleteCiudad() (planeta string, ciudad string) {
	log.Printf("¿Cuál es el nombre del planeta que contiene la ciudad a eliminar?:\n")
	fmt.Scan(&planeta)
	log.Printf("¿Cuál es el nombre de la ciudad a eliminar?:\n")
	fmt.Scan(&ciudad)
	return
}

func main() {
	//--------------CONEXIONES FULCRUM ------------
	var conn1 *grpc.ClientConn //FULCRUM 1
	conn1, err1 := grpc.Dial("localhost:9060", grpc.WithInsecure())
	if err1 != nil {
		log.Fatalf("did not connect: %s", err1)
	}
	defer conn1.Close()
	fulcrum1 := pb.NewFulcrumClient(conn1)

	var conn2 *grpc.ClientConn //FULCRUM 2
	conn2, err2 := grpc.Dial("10.6.40.170:9070", grpc.WithInsecure())
	if err2 != nil {
		log.Fatalf("did not connect: %s", err2)
	}
	defer conn2.Close()
	fulcrum2 := pb.NewFulcrumClient(conn2)

	var conn3 *grpc.ClientConn //FULCRUM 3
	conn3, err3 := grpc.Dial("10.6.40.171:9040", grpc.WithInsecure())
	if err3 != nil {
		log.Fatalf("did not connect: %s", err3)
	}
	defer conn3.Close()
	fulcrum3 := pb.NewFulcrumClient(conn3)
	// -------- FIN CONEXIONES FULCRUM ----------------

	//Comienza conexion con el broker
	log.Printf("Informante Ahsoka Tano iniciada. \n")
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("10.6.40.172:9050", grpc.WithInsecure()) //Conexión a Broker
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := pb.NewBrokerClient(conn)

	log.Printf("¿Qué acción desea realizar?:\n")
	log.Printf("[1] Añadir una nueva ciudad.\n")
	log.Printf("[2] Actualizar nombre ciudad.\n")
	log.Printf("[3] Actualizar número de rebeldes de ciudad.\n")
	log.Printf("[4] Borrar una ciudad.\n")
	log.Printf("[5] No quiero hacer ni una wea más.\n")
	var opcion int32 = 0
	for ok := true; ok; ok = (opcion != 5) {
		fmt.Scan(&opcion)
		if opcion == 1 { //Añadir ciudad
			response, err := c.QuieroHacer(context.Background(), &pb.Comando{Comando: "AddCity"})
			if err != nil {
				log.Fatalf("Error when calling QuieroHacer: %s", err)
			}
			log.Printf("Respuesta del Broker: %s", response.Valor)
			var planet, city, rebelds string
			planet, city, rebelds = AddCiudad()
			if response.Valor == "10.6.40.169" { //fulcrum1 localhots
				res_fulcrum1, err_f1 := fulcrum1.AddCity(context.Background(), &pb.Estructura{Planeta: planet, Ciudad: city, Rebeldes: rebelds})
				if err_f1 != nil {
					log.Fatalf("Error when calling AddCity: %s", err_f1)
				}
				log.Printf("Respuesta del Fulcrum 1: %d", res_fulcrum1.X)

			} else if response.Valor == "10.6.40.170" { //fulcrum2
				res_fulcrum2, err_f2 := fulcrum2.AddCity(context.Background(), &pb.Estructura{Planeta: planet, Ciudad: city, Rebeldes: rebelds})
				if err_f2 != nil {
					log.Fatalf("Error when calling AddCity: %s", err_f2)
				}
				log.Printf("Respuesta del Fulcrum 2: %d", res_fulcrum2.X)
			} else { //fulcrum 3
				res_fulcrum3, err_f3 := fulcrum3.AddCity(context.Background(), &pb.Estructura{Planeta: planet, Ciudad: city, Rebeldes: rebelds})
				if err_f3 != nil {
					log.Fatalf("Error when calling AddCity: %s", err_f3)
				}
				log.Printf("Respuesta del Fulcrum 3: %d", res_fulcrum3.X)

			}

		} else if opcion == 2 { //Actualizar ciudad

			response2, err2 := c.QuieroHacer(context.Background(), &pb.Comando{Comando: "UpdateName"})
			if err2 != nil {
				log.Fatalf("Error when calling QuieroHacer: %s", err2)
			}
			log.Printf("Respuesta del Broker: %s", response2.Valor)
			var planet, city, new_city string
			planet, city, new_city = UpdateCiudad()

			if response2.Valor == "10.6.40.169" { //fulcrum1 localhots
				res_fulcrum1, err_f1 := fulcrum1.UpdateName(context.Background(), &pb.Estructura{Planeta: planet, Ciudad: city, Rebeldes: new_city})
				if err_f1 != nil {
					log.Fatalf("Error when calling UpdateName: %s", err_f1)
				}
				log.Printf("Respuesta del Fulcrum 1: %d", res_fulcrum1.X)

			} else if response2.Valor == "10.6.40.170" { //fulcrum2
				res_fulcrum2, err_f2 := fulcrum2.UpdateName(context.Background(), &pb.Estructura{Planeta: planet, Ciudad: city, Rebeldes: new_city})
				if err_f2 != nil {
					log.Fatalf("Error when calling UpdateName: %s", err_f2)
				}
				log.Printf("Respuesta del Fulcrum 2: %d", res_fulcrum2.X)
			} else { //fulcrum 3
				res_fulcrum3, err_f3 := fulcrum3.UpdateName(context.Background(), &pb.Estructura{Planeta: planet, Ciudad: city, Rebeldes: new_city})
				if err_f3 != nil {
					log.Fatalf("Error when calling UpdateName: %s", err_f3)
				}
				log.Printf("Respuesta del Fulcrum 3: %d", res_fulcrum3.X)

			}

		} else if opcion == 3 { //Actualizar rebeldes

			response3, err3 := c.QuieroHacer(context.Background(), &pb.Comando{Comando: "UpdateNumber"})
			if err3 != nil {
				log.Fatalf("Error when calling QuieroHacer: %s", err3)
			}
			log.Printf("Respuesta del Broker: %s", response3.Valor)
			var planet, city, new_rebeldes string
			planet, city, new_rebeldes = UpdateRebeldes()

			if response3.Valor == "10.6.40.169" { //fulcrum1 localhots
				res_fulcrum1, err_f1 := fulcrum1.UpdateNumber(context.Background(), &pb.Estructura{Planeta: planet, Ciudad: city, Rebeldes: new_rebeldes})
				if err_f1 != nil {
					log.Fatalf("Error when calling UpdateNumber: %s", err_f1)
				}
				log.Printf("Respuesta del Fulcrum 1: %d", res_fulcrum1.X)

			} else if response3.Valor == "10.6.40.170" { //fulcrum2
				res_fulcrum2, err_f2 := fulcrum2.UpdateNumber(context.Background(), &pb.Estructura{Planeta: planet, Ciudad: city, Rebeldes: new_rebeldes})
				if err_f2 != nil {
					log.Fatalf("Error when calling UpdateNumber: %s", err_f2)
				}
				log.Printf("Respuesta del Fulcrum 2: %d", res_fulcrum2.X)
			} else { //fulcrum 3
				res_fulcrum3, err_f3 := fulcrum3.UpdateNumber(context.Background(), &pb.Estructura{Planeta: planet, Ciudad: city, Rebeldes: new_rebeldes})
				if err_f3 != nil {
					log.Fatalf("Error when calling UpdateNumber: %s", err_f3)
				}
				log.Printf("Respuesta del Fulcrum 3: %d", res_fulcrum3.X)

			}

		} else {

			response4, err4 := c.QuieroHacer(context.Background(), &pb.Comando{Comando: "DeleteCity"})
			if err4 != nil {
				log.Fatalf("Error when calling QuieroHacer: %s", err4)
			}
			log.Printf("Respuesta del Broker: %s", response4.Valor)
			var planet, city string
			planet, city = DeleteCiudad()

			if response4.Valor == "10.6.40.169" { //fulcrum1 localhots
				res_fulcrum1, err_f1 := fulcrum1.DeleteCity(context.Background(), &pb.Estructura3{Planeta: planet, Ciudad: city})
				if err_f1 != nil {
					log.Fatalf("Error when calling DeleteCity: %s", err_f1)
				}
				log.Printf("Respuesta del Fulcrum 1: %d", res_fulcrum1.X)

			} else if response4.Valor == "10.6.40.170" { //fulcrum2
				res_fulcrum2, err_f2 := fulcrum2.DeleteCity(context.Background(), &pb.Estructura3{Planeta: planet, Ciudad: city})
				if err_f2 != nil {
					log.Fatalf("Error when calling DeleteCity: %s", err_f2)
				}
				log.Printf("Respuesta del Fulcrum 2: %d", res_fulcrum2.X)
			} else { //fulcrum 3
				res_fulcrum3, err_f3 := fulcrum3.DeleteCity(context.Background(), &pb.Estructura3{Planeta: planet, Ciudad: city})
				if err_f3 != nil {
					log.Fatalf("Error when calling DeleteCity: %s", err_f3)
				}
				log.Printf("Respuesta del Fulcrum 3: %d", res_fulcrum3.X)

			}

		}
	}

}
