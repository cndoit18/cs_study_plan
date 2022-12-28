#!/usr/bin/env python3
from socket import *

serverName = '0.0.0.0'
serverPort = 8080

serverSocket = socket(AF_INET, SOCK_STREAM)
serverSocket.bind((
    serverName,
    serverPort,
))
serverSocket.listen(1)

print('The server is ready to receive')
while True:
    connectionSocket, addr = serverSocket.accept()
    try:
        message = connectionSocket.recv(1024)
        filename = message.split()[1]
        f = open(filename[1:])
        body = '\n'.join(f.readlines())
        connectionSocket.send('HTTP/1.1 200 OK\r\n'.encode())
        connectionSocket.send('Content-Length: {}\r\n\r\n'.format(
            len(body)).encode())
        connectionSocket.send(body.encode())
    except IOError:
        body = 'resource not found\n'
        connectionSocket.send('HTTP/1.1 404 Not Found\r\n'.encode())
        connectionSocket.send('Content-Length: {}\r\n\r\n'.format(
            len(body)).encode())
        connectionSocket.send(body.encode())
    finally:
        connectionSocket.close()
