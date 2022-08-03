from flask import Flask, url_for, render_template
from flask import request
from flask import jsonify
import json
from flask_cors import CORS
from flask_sqlalchemy import SQLAlchemy
import psycopg2
import requests




app = Flask(__name__)
cors = CORS(app)
app.config ['SQLALCHEMY_DATABASE_URI'] = 'sqlite:///students.sqlite3'
db = SQLAlchemy(app)

class students(db.Model):
    id = db.Column('student_id', db.Integer, primary_key = True)
    name = db.Column(db.String(100))
    city = db.Column(db.String(50))  
    addr = db.Column(db.String(200))
    pin = db.Column(db.String(10))

    def __init__(self, name, city, addr,pin):
        self.name = name
        self.city = city
        self.addr = addr
        self.pin = pin


# if incoming http request has the '/' 
# suffix, flask will route that request through 
#  this resource and return whatever this function 
#  returns.
#  this function can also return rendered html file. 

@app.route('/',methods=['GET'])
def index():
    db.create_all()
    return render_template('index.html')

@app.route('/json',methods=['POST'])
def read_json():
    value = request.get_json()
    print(value)
    requests.post("http://localhost:8080/dbgo", value)
    # here push code to pgadmin, and run over the comments NLP
    # run over rest gdb 

    return jsonify({'value':1})

if __name__ == '__main__':
    app.run(debug=True)