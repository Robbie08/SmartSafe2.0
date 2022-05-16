import requests as req
import random
import keys
from twilio.rest import Client



# This function will send our payload with password
# @params: A password that we will send over to our http server
def sendPayload(passwd):
    payload = {'code': passwd}
    req.post('http://localhost:8080/api/v2/pyClient', params=payload)

# This function will generate a password
# @params length: determines the number of characters in our password
# @return : generated password
def generatePassword(length):
    return ''.join([chr(random.randint(48,122)) for _ in range(length)])
password = generatePassword(8)

#make Client to send message
client = Client(keys.account_sid , keys.auth_token)
#This sends the message to a given phone number
message = client.messages.create(
    body = password,#the password that is generated
    from_= keys.twilio_number,#twilio phone number
    to = keys.target_number #This can be any phone number, must in this format +1##########
)
#print(message.body)
sendPayload(password) # Generate password, package it into a payload and send it via http