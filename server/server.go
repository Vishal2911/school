package server

import (
	"fmt"

	"github.com/vishal2911/school/store/pgress"
)

type Server struct{



}

func (s Server)NewServer() {

	fmt.Println("Creating new server .....")
	pgressStore := pgress.PgressStore{}
	pgressStore.NewStore() 

}