from dotenv import load_dotenv
import os

load_dotenv()

class Access:
    ALLOWED_USERS = [int(x) for x in os.getenv('ALLOWED_USERS').split(',')]

    @staticmethod
    def check_access(user_id) -> bool:
        return Access.ALLOWED_USERS.count(user_id) > 0