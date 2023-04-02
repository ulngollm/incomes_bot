from income import Income
from db.queries import Repo

class Storage:
    def __init__(self) -> None:
        self.repo = Repo('app.db')

    def add_income(self, income: Income) -> None:
        self.repo.add(income.owner_id, income.name, income.value, income.date)

    def get_today_sum(self, user_id) -> int:
        return self.repo.get_today(user_id)
    
    def get_today_list(self, user_id) -> list:
        result = self.repo.get_today_list(user_id)
        return ["%+d\t%s" % (x[3], x[2]) for x in result]
    
    