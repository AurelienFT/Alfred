import os.path
import logging
import sys
from flask import Flask, request
from pymessenger.bot import Bot
ACCESS_TOKEN = 'ACCESS_TOKEN'
VERIFY_TOKEN = 'VERIFY_TOKEN'
import send_message
import get_message
from neural_network.train_neural_network import train
from neural_network.classify import classify
import facebook_routes

class Facebook:
    def __init__(self, bot, access_token, verify_token):
            self.bot = Bot(ACCESS_TOKEN)
            self.acess_token = ACCESS_TOKEN
            self.verify_token = VERIFY_TOKEN

class Alfred:
    def __init__(self, platform):
        if platform == "facebook":
            self.platform = "facebook"
            self.app = Flask(__name__)
            self.facebook = Facebook(Bot(ACCESS_TOKEN), ACCESS_TOKEN, VERIFY_TOKEN)

        return

    def get_message(self):
        return get_message.get_message(self)

    def send_message(self, message):
        return send_message.send_message(self, message)

if not os.path.isfile("synapses.json"):
    logging.info("Training the bot...")
    train(hidden_neurons=20, alpha=0.1, epochs=100000, dropout=False, dropout_percent=0.2)
if len(sys.argv) < 2:
    Alfred = Alfred("cli")
    while True:
        data = raw_input("Ask me anything: ")
        print(classify(data))
else:
    Alfred = Alfred(sys.argv[1])
