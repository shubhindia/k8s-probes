import flask
import time
import socket

app = flask.Flask(__name__)


@app.route('/healthz')
def healthz():
    return "hello from: " + socket.gethostname() + "\n"


if __name__ == '__main__':
    #time.sleep(30)
    app.run(host='0.0.0.0', port=8080)