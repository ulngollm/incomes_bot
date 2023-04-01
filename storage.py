from income import Income

class Storage:
    def __init__(self) -> None:
        self.incomes = []

    def add_income(self, income: Income):
        self.incomes.append(income)

    def get_today_sum(self):
        # todo get today summ
        return sum([x.value for x in self.incomes])