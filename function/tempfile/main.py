import os
import datetime
import re
import json
import string
import random
import requests
import urllib
from datetime import datetime, timezone, timedelta

from bs4 import BeautifulSoup
from flask import Flask, request, render_template, redirect
from google.cloud import firestore
import functions_framework
import jsonpickle

project_id = os.environ['GCP_PROJECT']
bucket = "temp.akiicat.com"

app = Flask(__name__)
db = firestore.Client(project=project_id)

# Register an HTTP function with the Functions Framework
# Your function is passed a single parameter, (request), which is a Flask Request object.
# https://flask.palletsprojects.com/en/1.0.x/api/#flask.Request
@functions_framework.http
def entrypoint(request):
    # Create a new app context for the app
    internal_ctx = app.test_request_context(
        path=request.full_path, method=request.method
    )
    # Copy the request headers to the app context
    internal_ctx.request = request
    # Activate the context
    internal_ctx.push()
    # Dispatch the request to the internal app and get the result
    return_value = app.full_dispatch_request()
    # Offload the context
    internal_ctx.pop()
    # Return the result of the internal app routing and processing
    return return_value

def check_collision(key):
    doc_ref = db.collection(u'tempfile').document(key)
    doc = doc_ref.get()
    if doc.exists:
        doc_ref.delete()
    return doc.exists

def random_key(length):
   letters = "0123456789abcdefghjkmnopqrstuvwxyzABCDEFGHJKMNOPQRSTUVWXYZ"
   return ''.join(random.choice(letters) for i in range(length))

@app.route("/", methods=["GET"])
def tempfile():
    return render_template('index.html'), 200

@app.route("/upload", methods=["POST"])
def upload():
    filesize = request.form.get('filesize', default=0, type=int)
    filename = request.form.get('filename', default="", type=str)
    nonce = request.form.get('nonce', default="", type=str)
    
    if filesize == 0 or not filename or not nonce:
        return "Status Bad Request", 400

    key = random_key(3)
    while check_collision(key):
        key = random_key(3)

    doc_ref = db.collection(u'tempfile').document(key)

    url = "https://storage.googleapis.com/" + bucket + "/" + key + "/" + urllib.parse.quote(filename)

    doc_ref.set({
        u'Timestamp': firestore.SERVER_TIMESTAMP,  # datetime.datetime.now()
        u'Key': key,
        u'Name': filename,
        u'Nonce': nonce,
        u'Url': url,
    })
    
    return {
        "url": url,
        "token": key,
        "name": filename,
        "nonce": nonce,
    }, 200

@app.route("/<string:key>", methods=["GET"])
def download(key):
    if len(key.strip()) != 3:
        return "Not Found", 404

    doc_ref = db.collection(u'tempfile').document(key)

    doc = doc_ref.get()
    if not doc.exists:
        return "Not Found", 404

    obj = doc.to_dict()
    url = "https://storage.googleapis.com/{}/{}/{}".format(bucket, key, urllib.parse.quote(obj['Name']))
    return redirect(url, code=302)

