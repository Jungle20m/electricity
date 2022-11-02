package httpserver

type HttpServer struct {
	AppContext string
}

func NewServer() *HttpServer {
	return &HttpServer{}
}

func (sv *HttpServer) Serve() {

}

//server := http.Server{
//Addr:              fmt.Sprintf("%s:%s", config.Webserver.Host, config.Webserver.Port),
//Handler:           handler,
//TLSConfig:         nil,
//ReadTimeout:       0,
//ReadHeaderTimeout: 0,
//WriteTimeout:      0,
//IdleTimeout:       0,
//MaxHeaderBytes:    0,
//TLSNextProto:      nil,
//ConnState:         nil,
//ErrorLog:          nil,
//BaseContext:       nil,
//ConnContext:       nil,
//}
