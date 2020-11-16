from flask import Flask,request,redirect,Response
import requests
import logging

logging.basicConfig(level=logging.DEBUG)
app = Flask(__name__)
SITE_NAME = 'https://google.com/'

@app.route('/')
def index():
    return 'Flask is running!'


@app.route('/<path:path>',methods=['GET','POST','DELETE'])
def proxy(path):
    global SITE_NAME
    if request.method=='GET':
        resp = requests.get(f'{SITE_NAME}{path}', verify=False, headers=request.headers)
        excluded_headers = ['content-encoding', 'content-length', 'transfer-encoding', 'connection']
        headers = [(name, value) for (name, value) in  resp.raw.headers.items() if name.lower() not in excluded_headers]
        response = Response(resp.content, resp.status_code, headers)
        return response
    elif request.method=='POST':
        resp = requests.post(f'{SITE_NAME}{path}',json=request.get_json(), verify=False, headers=request.headers)
        excluded_headers = ['content-encoding', 'content-length', 'transfer-encoding', 'connection']
        headers = [(name, value) for (name, value) in resp.raw.headers.items() if name.lower() not in excluded_headers]
        response = Response(resp.content, resp.status_code, headers)
        return response
    elif request.method=='DELETE':
        resp = requests.delete(f'{SITE_NAME}{path}', verify=False, headers=request.headers).content
        response = Response(resp.content, resp.status_code, headers)
        return response

if __name__ == '__main__':
    app.run(debug = True,port=8080, host='0.0.0.0')
