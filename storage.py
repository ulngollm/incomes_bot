from income import Income

class Storage:
    def __init__(self) -> None:
        self.incomes = []

    def add_income(self, income: Income) -> None:
        self.incomes.append(income)

    def get_today_sum(self) -> int:
        # todo get today sum
        return sum([x.value for x in self.incomes])
    
    def get_today_list(self) -> list:
        return ["%+d\t%s" % (x.value, x.name) for x in self.incomes]
    
    