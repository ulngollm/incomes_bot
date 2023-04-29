from dotenv import load_dotenv
import os


load_dotenv()
API_ID = os.getenv('API_ID')
API_HASH = os.getenv('API_HASH')
BOT_API_TOKEN = os.getenv('BOT_API_TOKEN')
DB_NAME = os.getenv('DB_NAME')