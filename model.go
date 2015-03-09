package paginaBE

import "time"
import "appengine/datastore"



// Greeting is a datastore entity that represents a single greeting.
// It also serves as (a part of) a response of GreetingService.
type Noticia struct {
  Key     *datastore.Key `json:"id" datastore:"-"`
  Titulo  string         `json:"titulo"`
  Imagen string         `json:"imagen" datastore:",noindex"`
  Video  string         `json:"video" datastore:",noindex" endpoints:"req"`
  Descripcion string		 `json:"descripcion"`
  Date    time.Time      `json:"date"`
  ImagenWidth string		 `json:"imageWidth"`
  ImagenHeight string		 `json:"imageHeight"`
  VideoWidth string		 `json:"videoWidth"`
  VideoHeight string		 `json:"videoHeight"`
  Autor string `json:"autor"`
  Categoria string `json:"categoria"`

}

// GreetingsList is a response type of GreetingService.List method
type Noticias struct {
  Noticias []*Noticia `json:"noticias"`
}

// Request type for GreetingService.List
type NoticiasAMostrar struct {
  Limite int `json:"limite" endpoints:"d=10"`
}


type Business struct {
  Key     *datastore.Key `json:"id" datastore:"-"`
  Imagen string         `json:"imagen" datastore:",noindex" endpoints:"req"`
  Descripcion string     `json:"descripcion"`
  Date    time.Time      `json:"date"`
  Nombre string     `json:"nombre"`
  LinkedIn string `json:"linkedIn"`
}

// GreetingsList is a response type of GreetingService.List method
type BusinessCards struct {
  BusinessCards []*Business `json:"businessCards"`
}

// Request type for GreetingService.List
type BusinessCardsAmostrar struct {
  Limite int `json:"limite" endpoints:"d=10"`
}


type ContactoMensaje struct{
  Key     *datastore.Key `json:"id" datastore:"-"`
  Nombre string         `json:"nombre" datastore:",noindex" endpoints:"d=guillermo"`
  Correo string     `json:"correo" endpoints:"d=10"`
  Date    time.Time      `json:"date" endpoints:"d=10"`
  Mensaje string     `json:"mensaje" endpoints:"d=10"`
  Asunto string `json:"asunto" endpoints:"d=10"`

}

type ContactoReq struct{

  Nombre string         `json:"nombre"`
  Correo string     `json:"correo" `
  Mensaje string     `json:"mensaje" `
  Asunto string `json:"asunto" `
}

type MailTallerReq struct{

  Nombre string         `json:"nombre"`
  Correo string     `json:"correo" `
 
}



type MailTaller struct{

  Nombre string         `json:"nombre"`
  Key     *datastore.Key `json:"id" datastore:"-"`
  Correo string     `json:"correo" `
  Date    time.Time      `json:"date" endpoints:"d=10"`
}

