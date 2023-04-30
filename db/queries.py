import sqlite3

class Repo():
    def __init__(self, db_name) -> None:
        self.db = sqlite3.connect(db_name)


    def add(self, user, desc, value, date):
        # todo дату в unix формат, если выборки будут работать неправильно
        self.db.execute('''
            INSERT into incomes(user_id, desc, sum, date) values(?,?,?,?)
        ''', 
        (user, desc, value, date,))
        self.db.commit()


    def get_today(self, user) -> int:
        return self.db.execute('''
            SELECT sum(sum) from incomes where user_id = ? and date=date('now', 'localtime')
        ''',
        (user,)).fetchone()
    

    def get_daily(self, user, date) -> int:
        return self.db.execute('''
            SELECT sum(sum) from incomes where user_id = ? and date = ?
        ''',
        (user, date,)).fetchone()
    

    def get_today_list(self, user) -> int:
        return self.db.execute('''
            SELECT * from incomes where user_id = ? and date=date('now', 'localtime')
        ''',
        (user,)).fetchall()
    

    def get_week_sum(self, user) -> int:
        return self.db.execute('''
            SELECT sum(sum) from incomes where user_id = ? and date<=date('now', 'localtime') and date >= date('now', '-7 day', 'weekday 1')
        ''',
        (user,)).fetchone()
    

    def get_month_sum(self, user) -> int:
        return self.db.execute('''
            SELECT sum(sum) from incomes where user_id = ? and date <= date('now', 'localtime') and date >= date('now', 'localtime', 'start of month')
        ''',
        (user,)).fetchone()
    