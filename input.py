from storage import Storage
from income import Income
from pyrogram.types import Message
import re

class Input:
    def __init__(self, storage: Storage) -> None:
        self.storage = storage
    
    
    def add(self, message: Message, date=None):
        [value, name] = self.__normalize_input(message.matches.pop())
        self.storage.add_income(income=Income(name, value, message.from_user.id, date))

    
    def __normalize_input(self, matches):
        return int(matches.group('sum')), matches.group('desc')
    