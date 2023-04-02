from datetime import datetime

class Income:
    def __init__(self, name, value, user_id, date=None) -> None:
        self.name = name
        self.value = value
        self.owner_id = user_id
        self.date = datetime.strptime(date, "%d.%m.%Y").date() if date != None else datetime.now().date()