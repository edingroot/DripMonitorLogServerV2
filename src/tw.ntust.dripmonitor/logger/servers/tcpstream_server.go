package servers

import (
	"bufio"
	"log"
	"net"
	"tw.ntust.dripmonitor/logger/helpers"
	"fmt"
)

const logTag = "[TCPStream]"

// TCP server
type TcpServer struct {
	address                  string // Address to open connection - host:port
	onNewClientCallback      func(c *TcpClient)
	onClientConnectionClosed func(c *TcpClient, err error)
	onNewMessage             func(c *TcpClient, message string)
}

// TcpClient holds info about connection
type TcpClient struct {
	conn   net.Conn
	Server *TcpServer
}

func InitializeTCPStream(config *helpers.Configuration, mysqlConn *helpers.MySQLConn) *TcpServer {
	address := fmt.Sprintf("%s:%d", config.StreamListenHost, config.StreamListenPort)

	fmt.Printf("%s Creating tcp server...\n", logTag)
	server := &TcpServer{
		address: address,
	}

	server.OnNewClient(func(c *TcpClient) {
		fmt.Printf("%s New client connected: %s\n", logTag, c.conn.RemoteAddr().String())
	})

	server.OnNewMessage(func(c *TcpClient, message string) {
		fmt.Printf("%s Received message from %s, length=%d\n", logTag, c.conn.RemoteAddr().String(), len(message))
		// TODO
	})

	server.OnClientConnectionClosed(func(c *TcpClient, err error) {
		if err != nil {
			fmt.Printf("%s TCP streaming server %s\n", logTag, err.Error())
		}
		fmt.Printf("%s Client disconnected: %s\n", logTag, c.conn.RemoteAddr().String())
	})

	go server.Listen()

	fmt.Printf("%s TCP streaming server now listening on %s\n", logTag, address)

	return server
}

// Read client data from channel
func (c *TcpClient) listen() {
	reader := bufio.NewReader(c.conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("%s Failed to read from tcp client: %s\n", logTag, err.Error())
			c.conn.Close()
			c.Server.onClientConnectionClosed(c, err)
			return
		}
		c.Server.onNewMessage(c, message)
	}
}

/*
// Send text message to client
func (c *TcpClient) Send(message string) error {
	_, err := c.conn.Write([]byte(message))
	return err
}

// Send bytes to client
func (c *TcpClient) SendBytes(b []byte) error {
	_, err := c.conn.Write(b)
	return err
}
*/

func (c *TcpClient) Conn() net.Conn {
	return c.conn
}

func (c *TcpClient) Close() error {
	return c.conn.Close()
}

// Called right after server starts listening new client
func (s *TcpServer) OnNewClient(callback func(c *TcpClient)) {
	s.onNewClientCallback = callback
}

// Called right after connection closed
func (s *TcpServer) OnClientConnectionClosed(callback func(c *TcpClient, err error)) {
	s.onClientConnectionClosed = callback
}

// Called when Client receives new message
func (s *TcpServer) OnNewMessage(callback func(c *TcpClient, message string)) {
	s.onNewMessage = callback
}

// Start network server
func (s *TcpServer) Listen() {
	listener, err := net.Listen("tcp", s.address)
	if err != nil {
		log.Fatal("Error starting TCP server.")
	}
	defer listener.Close()

	for {
		conn, _ := listener.Accept()
		client := &TcpClient{
			conn:   conn,
			Server: s,
		}
		go client.listen()
		s.onNewClientCallback(client)
	}
}
