from audioop import add
from socket import *
import time

clientSocket = socket(AF_INET, SOCK_DGRAM)
clientSocket.settimeout(1)
for i in range(1, 11):
    try:
        now = time.time()
        clientSocket.sendto('Ping {} {}'.format(i, now).encode(), ('', 12000,))
        message, address = clientSocket.recvfrom(1024)
        print('#{}: Takes {} seconds'.format(i, time.time()-now))
    except timeout:
        print('#{}: Request timed out'.format(i))
clientSocket.close()
