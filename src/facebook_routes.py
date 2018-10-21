from initialize import app, VERIFY_TOKEN
from flask import Flask, request

@app.route("/facebook", methods=['GET', 'POST'])
def receive_message():
    if request.method == 'GET':
        token_sent = request.args.get("hub.verify_token")
        return verify_fb_token(token_sent)

@app.route("/ping", methods=['GET', 'POST'])
def ping():
    return "pong"

def verify_fb_token(token_sent): 
    if token_sent == VERIFY_TOKEN:
        return request.args.get("hub.challenge")
    return 'Invalid verification token'