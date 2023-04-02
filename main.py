from dotenv import load_dotenv
from pyrogram import Client
from pyrogram.handlers import MessageHandler, CallbackQueryHandler
from pyrogram import filters
from handlers import *
import os


load_dotenv()
API_ID = os.getenv('API_ID')
API_HASH = os.getenv('API_HASH')
BOT_API_TOKEN = os.getenv('BOT_API_TOKEN')

app = Client('bot', API_ID, API_HASH, bot_token=BOT_API_TOKEN)



app.add_handler(MessageHandler(add, filters.command(['add'])))
app.add_handler(MessageHandler(today_sum, filters.command(['today'])))
app.add_handler(MessageHandler(week_sum, filters.command(['week'])))
app.add_handler(MessageHandler(month_sum, filters.command(['month'])))

app.add_handler(MessageHandler(read_input, filters.text))
app.add_handler(CallbackQueryHandler(button_handler))

app.run()