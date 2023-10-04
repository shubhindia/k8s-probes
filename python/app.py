import flask
from werkzeug.middleware.dispatcher import DispatcherMiddleware
import socket
from prometheus_client import Counter, make_wsgi_app

app = flask.Flask(__name__)
app.wsgi_app = DispatcherMiddleware(app.wsgi_app, {
    '/metrics': make_wsgi_app()
})


pong_counter = Counter('ping_request_count', 'No of request handled by Ping handler', ['handler'])

@app.route('/healthz')
def healthz():
    return "hello from: " + socket.gethostname() + "\n"

@app.route('/ping')
def ping():
    pong_counter.labels(handler='/ping').inc()
    return "pong from: " + socket.gethostname() + "\n"

if __name__ == '__main__':
    #time.sleep(30)

    app.run(host='0.0.0.0', port=8080)