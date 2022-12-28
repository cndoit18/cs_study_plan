#!/usr/bin/env python3

from socket import *
import sys

# if len(sys.argv) <= 1:
#     print('Usage : "python ProxyServer.py server_ip"\n[server_ip : It is the IP Address Of Proxy Server')
#     sys.exit(2)

# Create a server socket, bind it to a port and start listening
tcpSerSock = socket(AF_INET, SOCK_STREAM)
# Fill in start.
tcpSerSock.bind((
    '0.0.0.0',
    8888,
))
tcpSerSock.listen(1)
# Fill in end.
while True:
    print('Ready to serve...')
    tcpCliSock, addr = tcpSerSock.accept()
    print('Received a connection from:', addr)
    message = tcpCliSock.recv(1024)
    print(message)
    # Extract the filename from the given message
    print(message.split()[1])
    filename = message.split()[1].partition(b'/')[2]
    fileExist = False
    try:
        f = open(filename.replace(b'/', b'_'), 'r')
        outputdata = f.readlines()
        fileExist = True
        # ProxyServer finds a cache hit and generates a response message
        # Fill in start.
        for line in outputdata:
            tcpCliSock.send(line.replace('\n', '\r\n').encode())
        # # Fill in end.
        print('Read from cache')
    # Error handling for file not found in cache
    except IOError:
        if not fileExist:
            c = socket(AF_INET, SOCK_STREAM)
            hostn = filename.partition(b'/')[0]
            print('hostn: %s\nfilename: %s' % (hostn, filename))
            try:
                # Connect to the socket to port 80
                # Fill in start.
                c.connect((
                    hostn,
                    80,
                ))
                # Fill in end.
                # Create a temporary file on this socket and ask port 80
                # for the file requested by the client
                c.send(b'GET /' + filename.partition(b'/')[2] +
                       b' HTTP/1.0\r\n')
                c.send(b'Content-Type:text/html\r\n\r\n')
                # Read the response into buffer
                # Fill in start.
                # Fill in end.
                # Create a new file in the cache for the requested file.
                # Also send the response in the buffer to client socket and the corresponding file in the cache
                tmpFile = open("./" + filename.replace(b'/', b'_').decode(),
                               "wb")
                while True:
                    resp = c.recv(1024)
                    if not resp: break
                    tmpFile.write(resp)
                tmpFile.close()
                c.close()
            except:
                print('Illegal request')
        else:
            # HTTP response message for file not found
            # Do stuff here
            print('File Not Found...Stupid Andy')
    # Close the client and the server sockets
    tcpCliSock.close()
tcpSerSock.close()
