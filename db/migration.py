import sqlite3
from dotenv import load_dotenv
import os

load_dotenv()

DB_NAME = os.getenv('DB_NAME')
conn = sqlite3.connect(DB_NAME)

cursor = conn.cursor()

cursor.execute('''
    CREATE TABLE incomes
    (id INTEGER PRIMARY KEY,
    user_id INTEGER,
    desc TEXT DEFAULT '',
    sum INTEGER,
    date INTEGER)
''')


conn.commit()
conn.close()