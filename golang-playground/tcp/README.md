# Building TCP in GO

Typically there are two types of applications that use tcp, a client and a server. A TCP server listens on a well known ip and port combination and accepts connections from a TCP client.

A TCP client initiates a connection request to a TCP server in order to set up a connection with a server. A real TCP server can accept multiple connections on a socket. 

The server should be able to handle the following steps:
1. Create -> Creates a TCP socker
2. Bind -> assigns a local protocol address to a socket. With the Internet protocols, the address is the combination of an IPv4 or IPv6 address (32-bit or 128-bit) address along with a 16 bit TCP port number.
3. Listen -> Puts the socket in a listening mode, waiting for the client to approach the server to make a connection
4. Accept -> Accepts that the connection between server is established and is ready to transfer data.
5. Listen -> see step 3

The client should be able to do the following:
1. 


## Resources
- [TCP implementation in C](https://www.geeksforgeeks.org/tcp-server-client-implementation-in-c/)

- [Simple TCP server in GO](https://medium.com/@iggeehu/learning-go-by-writing-a-simple-tcp-server-d8ed260f67ac)
- [Messing with TCP and systemcalls](https://medium.com/@wu.victor.95/tinkering-with-tcp-and-sockets-70255a707fa0)
- [Wikipedia TCP](https://en.wikipedia.org/wiki/Transmission_Control_Protocol)