import random
from flask import Flask, request
from pymessenger.bot import Bot
app = Flask(__name__)
FACEBOOK_ACCESS_TOKEN = 'EAAMoWPp4rwoBABiii4iDV8uSuZBl5MblHHtDhmbqw7lk93wwpirM53ZBeUWhseUWbxEsCsGGgpg2pfDUXuEhLL6LI3RWMBwm1IlFnP8B1q9ZC0QkEaWoAgCBn10g7no4VBlCyNJcVykHqImvZBwZClRe748ZA7ZCqh20qmkvrgvRwZDZD'
VERIFY_TOKEN = 'income_messenger'
bot = Bot(FACEBOOK_ACCESS_TOKEN)

def initialize(alfred, platform):
    if platform == "facebook":
        alfred.app = app
        import facebook_routes
        app.run()
    return alfred
