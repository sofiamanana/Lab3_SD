package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	pb "Lab3_SD/proto"
)
func AddCity( nombre_planeta string,nombre_ciudad string,nuevo_valor=0 int){
	file, err := os.Open( nombre_planeta +".txt")
	if err != nil {
		file, err := os.Create(nombre_planeta +".txt")
		if err != nil {
			log.Fatal(err)
		  }
	}
	defer file.Close()
	_, err2 := file.WriteString(nombre_planeta+" "+nombre_ciudad+" "+nuevo_valor+"/n")
	if err2 != nil {
		log.Fatal(err2)
	}
	//fmt.Println("listo")
}
func UpdateName( nombre_planeta string,nombre_ciudad string,nuevo_valor string){
file, err := os.Open( nombre_planeta+"txt")
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


func main(){
	
}