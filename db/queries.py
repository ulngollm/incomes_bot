import sqlite3

class Repo():
    def __init__(self, db_name) -> None:
        self.db = sqlite3.connect(db_name)


    def add(self, user, desc, value,):
        self.db.execute('''
            INSERT into incomes(user_id, desc, sum, date) values(?,?,?,date('now'))
        ''', 
        (user, desc, value,))
        self.db.commit()


    def get_today(self, user) -> int:
        return self.db.execute('''
            SELECT sum(sum) from incomes where user_id = ? and date=date('now')
        ''',
        (user,)).fetchone()
    
    def get_today_list(self, user) -> int:
        return self.db.execute('''
            SELECT * from incomes where user_id = ? and date=date('now')
        ''',
        (user,)).fetchall()
    