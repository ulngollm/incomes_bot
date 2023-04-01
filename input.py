from storage import Storage
from income import Income

class Input:
    def __init__(self, storage) -> None:
        self.storage = storage
    
    def add(self, text):
        # todo parse text
        # todo проверить регуляркой, того ли формата текст
        # todo нормализовать текст
        [name, value] = text.split(". ")
        value = int(value)
        self.storage.add_income(income=Income(name, value))

    