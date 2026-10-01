import sqlite3

from flask import Flask, request

app = Flask(__name__)


@app.route("/user")
def user():
    name = request.args.get("name", "")
    conn = sqlite3.connect("users.db")
    # Deliberate SQL injection for the CodeQL test.
    rows = conn.execute("SELECT * FROM users WHERE name = '%s'" % name).fetchall()
    return {"rows": rows}
