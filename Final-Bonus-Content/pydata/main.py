#!python
from bottle import route, run

@route('/')
def hello():
    return "Hello World!"

# Starts a local development server on port 8080
run(host='0.0.0.0', port=8080, debug=True)
