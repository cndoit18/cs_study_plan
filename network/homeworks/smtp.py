import base64
from socket import *
import os

msg = "\r\n I love computer networks!"
endmsg = "\r\n.\r\n"
# Choose a mail server (e.g. Google mail server) and call it mailserver 
mailserver = ('smtp.qq.com', 25)
# Create socket called clientSocket and establish a TCP connection with mailserver
#Fill in start
clientSocket = socket(AF_INET, SOCK_STREAM)
clientSocket.connect(mailserver)
#Fill in end
recv = clientSocket.recv(1024)
print(recv)
if recv[:3].decode() != '220':
    print('220 reply not received from server.')

# Send HELO command and print server response.
clientSocket.send('HELO Alice\r\n'.encode())
recv1 = clientSocket.recv(1024)
print(recv1)
if recv1[:3].decode() != '250':
    print('250 reply not received from server.')

clientSocket.send('AUTH LOGIN\r\n'.encode())
recv1 = clientSocket.recv(1024)
print(recv1)
if recv1[:3].decode() != '334':
    print('334 reply not received from server.')

clientSocket.send('{}\r\n'.format(base64.b64encode('1047439649@qq.com'.encode()).decode()).encode())
recv1 = clientSocket.recv(1024)
print(recv1)
if recv1[:3].decode() != '334':
    print('334 reply not received from server.')
clientSocket.send('{}\r\n'.format(base64.b64encode(os.getenv('EMIAL_PWD').encode()).decode()).encode())
recv1 = clientSocket.recv(1024)
print(recv1)
if recv1[:3].decode() != '235':
    print('235 reply not received from server.')
# Send MAIL FROM command and print server response.
# Fill in start
clientSocket.send('MAIL FROM:<1047439649@qq.com>\r\n'.encode())
recv1 = clientSocket.recv(1024)
print(recv1)
if recv1[:3].decode() != '250':
    print('250 reply not received from server.')
# Fill in end

# Send RCPT TO command and print server response.
# Fill in start
clientSocket.send('RCPT TO:<cndoit18@outlook.com>\r\n'.encode())
recv1 = clientSocket.recv(1024)
print(recv1)
if recv1[:3].decode() != '250':
    print('250 reply not received from server.')
# Fill in end

# Send DATA command and print server response.
# Fill in start
clientSocket.send('DATA\r\n'.encode())
recv1 = clientSocket.recv(1024)
print(recv1)
if recv1[:3].decode() != '354':
    print('354 reply not received from server.')
# Fill in end

# Send message data.
# Fill in start
clientSocket.send(msg.encode())
# Fill in end 

# Message ends with a single period.
# Fill in start
clientSocket.send(endmsg.encode())
recv1 = clientSocket.recv(1024)
print(recv1)
if recv1[:3].decode() != '250':
    print('250 reply not received from server.')
# Fill in end

# Send QUIT command and get server response.
# Fill in start
clientSocket.send('QUIT\r\n'.encode())
recv1 = clientSocket.recv(1024)
print(recv1)
if recv1[:3].decode() != '221':
    print('221 reply not received from server.')
# Fill in end

clientSocket.close()