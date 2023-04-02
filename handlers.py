from state import State
from input import Input
from storage import Storage
from pyrogram.enums import ParseMode
from pyrogram.types import (InlineKeyboardMarkup,InlineKeyboardButton)

state = State()
storage = Storage()


async def add(client, message):
    await message.reply(
        "Запишите название и сумму в формате `-100 название`. Валюту указывать не надо",
        parse_mode=ParseMode.MARKDOWN
    )
    state.readInput('add')



async def today_sum(client, message):
    sum = storage.get_today_sum()
    await message.reply(
        "Ваш итог за сегодня %s руб." % sum,
        reply_markup=InlineKeyboardMarkup([
            [
                InlineKeyboardButton(
                    "Подробнее",
                    callback_data="today.list"
                ),
            ]
        ])
    )

async def week_sum(client, message):
    await message.reply(
        "Ваш итог за неделю 0 руб."
    )

async def month_sum(client, message):
    await message.reply(
        "Ваш итог за месяц 0 руб."
    )

async def read_input(client, message):
    lastCommand = state.handleCommand()

    # по умолчанию записывает новый расход
    if not lastCommand:
        Input(storage).add(message.text)
        await message.reply(
            "Ок"
        )
        return
    
    handler =  handlers[lastCommand]
    handler(message.text)
    await message.reply(
        "Ок"
    )

async def button_handler(client, callback_query):
    # todo считать последюю команду, для которой должен быть обработчик
    await callback_query.message.reply(
        '\n'.join(storage.get_today_list())
    )

handlers = {
    "add": Input(storage).add
}