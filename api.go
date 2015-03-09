package paginaBE

import "github.com/crhym3/go-endpoints/endpoints"



func init() {
  servicioNoticias := &ServicioNoticias{}
  api, err := endpoints.RegisterService(servicioNoticias,
    "noticias", "v1", "noticiasGDG API", true)
  if err != nil {
    panic(err.Error())
  }

  info := api.MethodByName("ListaNoticias").Info()
  info.Name, info.HttpMethod, info.Path, info.Desc =
    "noticias.listar", "GET", "listaNoticias", "Lista las ultimas noticias de GDG UNAM"

info2 := api.MethodByName("InsertaNoticia").Info()
  info2.Name, info2.HttpMethod, info2.Path, info2.Desc =
    "noticias.inserta", "GET", "insertaNoticias", "Inserta una noticia en el datastore de noticias"

ServicioNosotros := &ServicioNosotros{}
  api2, err2 := endpoints.RegisterService(ServicioNosotros,
    "presentacion", "v1", "presentacion GDG API", true)
  if err2 != nil {
    panic(err2.Error())
  }

info3 := api2.MethodByName("ListaNosotros").Info()
  info3.Name, info3.HttpMethod, info3.Path, info3.Desc =
    "presentacion.listar", "GET", "listaPresentacion", "Lista las tarjetas de presentacion de los GDG UNAM"

info4 := api2.MethodByName("InsertaRegistroNosotros").Info()
  info4.Name, info4.HttpMethod, info4.Path, info4.Desc =
    "presentacion.inserta", "GET", "insertaPresentacion", "Inserta una tarjeta de presentacion de algun GDG UNAM Organizer"

ServicioContacto := &ServicioContacto{}
  api3, err3 := endpoints.RegisterService(ServicioContacto,
    "contacto", "v1", "envia correos a los organizadores y al que contacto a  GDG", true)
  if err3 != nil {
    panic(err3.Error())
  }

info5 := api3.MethodByName("AccionContacto").Info()
  info5.Name, info5.HttpMethod, info5.Path, info5.Desc =
    "contacto.enviaCorreos", "GET", "accionContacto", "envia correos a los organizadores y al que contacto a  GDG"


ServicioPolymerWorkshop := &ServicioPolymerWorkshop{}
  api4, err4 := endpoints.RegisterService(ServicioPolymerWorkshop,
    "taller", "v1", "Envia un correo para ilustrar el uso de ajax en el taller", true)
  if err4 != nil {
    panic(err4.Error())
  }

info6 := api4.MethodByName("EnviarCorreoTaller").Info()
  info6.Name, info6.HttpMethod, info6.Path, info6.Desc =
    "taller.enviaCorreos", "GET", "correoTaller", "envia correo para ilustrar uso de ajax en taller"



  ServicioCorreoAmigos := &ServicioCorreoAmigos{}
  api5, err5 := endpoints.RegisterService(ServicioCorreoAmigos,
    "amigos", "v1", "Envia un correo para contactar a los clientes", true)
  if err5 != nil {
    panic(err5.Error())
  }

info7 := api5.MethodByName("EnviarCorreoAmigos").Info()
  info7.Name, info7.HttpMethod, info7.Path, info7.Desc =
    "amigos.enviaCorreos", "GET", "correoAmigos", "envia correo para contactar a los clientes"


  endpoints.HandleHttp()
}