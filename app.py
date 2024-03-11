from flask import Flask, render_template
from flask_socketio import SocketIO,send
import time

datch = {
    "temp":11,
    "high":22,
}

print(datch['high'])

app = Flask(__name__)

socket_ = SocketIO(app)

@app.route('/')
def index():
    return render_template('index.html')


@socket_.on('message')
def handle_message(data):
    time.sleep(1)
    print('received message: ' + data)
    mas = datch
    send(mas )

"""
ggdgdgdg
dgdgdgdgg 
"""


if __name__ == '__main__':
    socket_.run(app, debug=True)