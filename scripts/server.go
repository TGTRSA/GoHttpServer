type Node struct {
	data any
	next *Node
}
// Server object that has an address and listener object (not sure what that is)
type Server struct {
	addr string
	ln net.Listener
}


func main() {
	server , err:= NewServer("192.168.1.67:8080")
	if err != nil {
		logErr(err)
	}
	server.run()
}

func logErr(err error){
	fmt.Println(err)
}

func NewServer(addr string)(*Server, error){
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println(err)
	}
	return &Server{addr:addr, ln: ln}, nil
}

func (s *Server) run(){
	fmt.Printf("Starting server on %s\n", s.addr)
	for {
		conn, connErr := s.ln.Accept()
		if connErr != nil {
			logErr(connErr)
		}
		go func() {
			fmt.Printf("Got connection from %s\n", conn.RemoteAddr())
			handle(conn)
			defer conn.Close()
		}()
	}
}

func handle(conn net.Conn){
	buf := make([]byte, 1024)
	// I assume n is come lenght of some sort ??
	n, err := conn.Read(buf)
	if err != nil{
		logErr(err)
	}
	fmt.Printf("%s\n",string(buf[:n]))
	response := "Hello buddy"
	conn.Write([]byte(response))
}