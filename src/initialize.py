import random
from flask import Flask, request
from pymessenger.bot import Bot
app = Flask(__name__)
ACCESS_TOKEN = 'ACCESS_TOKEN'
VERIFY_TOKEN = 'VERIFY_TOKEN'
bot = Bot(ACCESS_TOKEN)

def initialize(alfred, platform):
    if platform == "facebook":
        alfred.app = app
        import facebook_routes
        app.run()
    return alfred