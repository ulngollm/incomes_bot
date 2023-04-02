from storage import Storage
from income import Income

class Input:
    def __init__(self, storage: Storage) -> None:
        self.storage = storage
    
    def add(self, message, date=None):
        # todo parse text
        # todo проверить регуляркой, того ли формата текст
        # todo нормализовать текст
        [value, name] = message.text.split(" ", 1)
        value = int(value)
        self.storage.add_income(income=Income(name, value, message.from_user.id, date))

    