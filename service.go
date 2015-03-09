package paginaBE

import "net/http"
import "github.com/crhym3/go-endpoints/endpoints"
import "appengine/datastore"
import "time"
import "appengine/mail"
import "fmt"
import "appengine"
//para sendgird
import "appengine/urlfetch"
import "github.com/sendgrid/sendgrid-go"


// GreetingService can sign the guesbook, list all greetings and delete
// a greeting from the guestbook.
type ServicioNoticias struct {
}

// List responds with a list of all greetings ordered by Date field.
// Most recent greets come first.
func (gs *ServicioNoticias) ListaNoticias(
  r *http.Request, req *NoticiasAMostrar, resp *Noticias) error {

  if req.Limite <= 0 {
    req.Limite = 10
  }

  c := endpoints.NewContext(r)
  q := datastore.NewQuery("Noticia").Order("-Date").Limit(req.Limite)
  noticias := make([]*Noticia, 0, req.Limite)
  keys, err := q.GetAll(c, &noticias)
  if err != nil {
    return err
  }

  for i, k := range keys {
    noticias[i].Key = k
  }
  resp.Noticias = noticias
  return nil
}

func (gs *ServicioNoticias) InsertaNoticia(
  r *http.Request, req *Noticia, resp *Noticia) error {

  c := endpoints.NewContext(r)

  noticia := &Noticia{


    Titulo: req.Titulo,//"Android wear",  
    //Imagen: "https://encrypted-tbn2.gstatic.com/images?q=tbn:ANd9GcTjg9p1aOE-6WGkAQxpx7F9ILWIAiHqJXMukFLCtM5GqhXmkZrm",
    Imagen: "",
    Video: req.Video,//"//www.youtube.com/embed/LeLKV4i7agA",
    //VideoHeight: "400vh",
    //VideoWidth: "400vh",
    //ImagenWidth: "0",
    //ImagenHeight: "0",
    Descripcion: req.Descripcion,//"Esta es una introduccion al concepto de wearables que tendra un fuerte empuje por parte de Google en los proximos meses, comienza a programar para relojes inteligentes con GDG UNAM como tutor :)",    
    Autor: req.Autor,// "Guillermo Romero ",
    Categoria : req.Categoria,//"Programming language",
    Date: time.Now(),   
    
  }



  key := datastore.NewIncompleteKey(c, "Noticia", nil)
  _, err := datastore.Put(c, key, noticia)


  if err != nil {
    return err
  }

  resp= noticia

  return nil
}



type ServicioNosotros struct {
}

// List responds with a list of all greetings ordered by Date field.
// Most recent greets come first.
func (gs *ServicioNosotros) ListaNosotros(
  r *http.Request, req *BusinessCardsAmostrar, resp *BusinessCards) error {

  if req.Limite <= 0 {
    req.Limite = 10
  }

  c := endpoints.NewContext(r)
  q := datastore.NewQuery("Business").Order("-Date").Limit(req.Limite)
  listaNosotros := make([]*Business, 0, req.Limite)
  keys, err := q.GetAll(c, &listaNosotros)
  if err != nil {
    return err
  }

  for i, k := range keys {
    listaNosotros[i].Key = k
  }
  resp.BusinessCards = listaNosotros
  return nil
}

func (gs *ServicioNosotros) InsertaRegistroNosotros(
  r *http.Request, req *Business, resp *Business) error {

  c := endpoints.NewContext(r)

  registroNosotros := &Business{

    Imagen: req.Imagen,//"https://lh6.googleusercontent.com/-xPOt15zUMn4/AAAAAAAAAAI/AAAAAAAADqE/2CYtJ7_xB4k/s120-c/photo.jpg",
    Descripcion: req.Descripcion,//"Android programmer, lover of Google (its just a test) Android programmer, lover of Google (its just a test) Android programmer, lover of Google (its just a test) Android programmer, lover of Google (its just a test) Android programmer, lover of Google (its just a test)",    
    Nombre: req.Nombre,//"Salvador Murillo",
    Date: time.Now(),  
    LinkedIn : req.LinkedIn,//"http://lnkd.in/bPu5Nyk",
    
  }



  key := datastore.NewIncompleteKey(c, "Business", nil)
  _, err := datastore.Put(c, key, registroNosotros)


  if err != nil {
    return err
  }

  resp= registroNosotros

  return nil
}






type ServicioContacto struct {
}

// List responds with a list of all greetings ordered by Date field.
// Most recent greets come first.
func (gs *ServicioContacto) AccionContacto(
  r *http.Request, req *ContactoReq, resp *Business) error {
  
c := endpoints.NewContext(r)

  mensaje := &ContactoMensaje{

  Nombre : req.Nombre,//"Guillermo Romero ", 
  Correo  : req.Correo,//"memo@hola",
  Mensaje : req.Mensaje,//"hola",
  Asunto :req.Asunto,//"saludo",
  Date: time.Now(),   
    
  }

  key := datastore.NewIncompleteKey(c, "ContactoMensaje", nil)
  _, err := datastore.Put(c, key, mensaje)


if err != nil {
    return err
  }

var textoMensaje = 
`Estimado organizador de GDG UNAM:
`+req.Nombre+` envio un mensaje:
 asunto:  `+req.Asunto+` 
 mensaje: `+req.Mensaje+`
 contacto:`+req.Correo+`
 PD: Este mensaje ya ha sido guardado en el datastore.`



const textoMensaje2 = 

`Gracias por el interés en GDG UNAM! 

Tu mensaje ha sido enviado a los organizadores exitosamente. Te invitamos a que continúes informado y disfrutes del contenido y los eventos que hemos creado para ti.

Recibe un cordial saludo,
GDG UNAM Team
`


  c2 := appengine.NewContext(r)

        addr := "memogrr.dohko@gmail.com"//r.FormValue("email")
        //url := createConfirmationURL(r)
        msg := &mail.Message{
                Sender:  "gdgunam@gmail.com",
                To:      []string{addr,"leomindezji@gmail.com","salvadoredgar99@gmail.com","juanpflores94@gmail.com"},
                Subject: "Contacto GDG UNAM",
                Body:    fmt.Sprintf(textoMensaje),
        }
        if err2 := mail.Send(c2, msg); err2 != nil {
                c2.Errorf("Couldn't send email: %v", err2)
        }



        
        //url := createConfirmationURL(r)
        msg2 := &mail.Message{
                Sender:  "gdgunam@gmail.com",
                To:      []string{req.Correo},
                Subject: "Contacto GDG UNAM",
                Body:    fmt.Sprintf(textoMensaje2),
        }
        if err3 := mail.Send(c2, msg2); err3 != nil {
                c2.Errorf("Couldn't send email: %v", err3)
        }


  return nil
}


type ServicioPolymerWorkshop struct {
}

// List responds with a list of all greetings ordered by Date field.
// Most recent greets come first.
func (gs *ServicioPolymerWorkshop) EnviarCorreoTaller(
  r *http.Request, req *MailTallerReq, resp *Business) error {
  
c := endpoints.NewContext(r)

  mailTaller := &MailTaller{

  Nombre : req.Nombre,//"Guillermo Romero ", 
  Correo  : req.Correo,//"memo@hola",
  Date: time.Now(),   
    
  }

  key := datastore.NewIncompleteKey(c, "MailTaller", nil)
  _, err := datastore.Put(c, key, mailTaller)


if err != nil {
    return err
  }




const textoMensaje2 = 

`Mensaje de taller Polymer
Este correo se envio desde App Engine para ilustrar el uso de ajax con Polymer
Recibe un cordial saludo,
GDG UNAM Team
`


  c2 := appengine.NewContext(r)
        
        //url := createConfirmationURL(r)
        msg2 := &mail.Message{
                Sender:  "gdgunam@gmail.com",
                To:      []string{req.Correo},
                Subject: "Taller Polymer",
                Body:    fmt.Sprintf(textoMensaje2),
        }
        if err3 := mail.Send(c2, msg2); err3 != nil {
                c2.Errorf("Couldn't send email: %v", err3)
        }


  return nil
}




type ServicioCorreoAmigos struct {
}

// List responds with a list of all greetings ordered by Date field.
// Most recent greets come first.

func (gs *ServicioCorreoAmigos) EnviarCorreoAmigos(
  r *http.Request, req *MailTallerReq, resp *Business) error {

/*
c := endpoints.NewContext(r)

  mailTaller := &MailTaller{

  Nombre : req.Nombre,//"Guillermo Romero ", 
  Correo  : req.Correo,//"memo@hola",
  Date: time.Now(),   
    
  }

  key := datastore.NewIncompleteKey(c, "MailTaller", nil)
  _, err := datastore.Put(c, key, mailTaller)


if err != nil {
    return err
  }




const textoMensaje2 = 

`Gracias por el interes en querer coolaborar con Amigos, te contactaremos en breve para verificar tu disponibilidad y otras consideraciones,
 recibe un cordial saludo por parte del equipo de Amigos.

`


  c2 := appengine.NewContext(r)
        
        //url := createConfirmationURL(r)
        msg2 := &mail.Message{
                Sender:  "gdgunam@gmail.com",
                To:      []string{req.Correo},
                Subject: "Amigos contacto",
                Body:    fmt.Sprintf(textoMensaje2),
        }
        if err3 := mail.Send(c2, msg2); err3 != nil {
                c2.Errorf("Couldn't send email: %v", err3)
        }


 


*/


  sg := sendgrid.NewSendGridClient("memocostecho", "D3$t!ny.")
  c := appengine.NewContext(r)
  // set http.Client to use the appengine client
  sg.Client = urlfetch.Client(c) //Just perform this swap, and you are good to go.
  message := sendgrid.NewMail()
  message.AddTo(req.Correo)
  message.SetSubject("Amigos contact")
  message.SetHTML(`Gracias por el interes en querer coolaborar con Amigos, te contactaremos en breve para verificar tu disponibilidad y otras consideraciones,
 recibe un cordial saludo por parte del equipo de Amigos.
`)
  message.SetFrom("gdgunam@gmail.com")
  if r := sg.Send(message); r == nil {
    fmt.Println("Email sent!")

  } else {
    c.Errorf("Unable to send mail %v",r)
    
  }



  return nil

}










