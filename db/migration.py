import sqlite3

# todo remove hardcode
conn = sqlite3.connect('app.db')

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