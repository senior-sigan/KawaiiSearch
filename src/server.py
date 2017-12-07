# -*- coding: utf-8 -*-

from flask import Flask, render_template

from imlucky import imlucky

app = Flask(__name__, template_folder='../templates', static_folder='../data')


@app.route('/', methods=['GET'])
def index():
    return render_template('index.html')


@app.route('/', methods=['POST'])
def imlucky_action():
    files = imlucky()
    return render_template('index.html', files=files)


if __name__ == "__main__":
    app.run(host='0.0.0.0', port=5000)
