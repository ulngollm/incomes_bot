from state import State
from input import Input
from storage import Storage
state = State()
storage = Storage()


async def add(client, message):
    await message.reply(
        'Запишите название и сумму в формате *название. -100*. Валюту указывать не надо'
    )
    state.readInput('add')



async def today_sum(client, message):
    sum = storage.get_today_sum()
    await message.reply(
        "Ваши доходы за сегодня составили %s руб." % sum
    )

async def week_sum(client, message):
    await message.reply(
        "Ваши доходы за неделю составили 0 руб."
    )

async def month_sum(client, message):
    await message.reply(
        "Ваши доходы за месяц составили 0 руб."
    )

async def read_input(client, message):
    handler =  handlers[state.handleCommand()]
    handler(message.text)
    await message.reply(
        "Ок"
    )

handlers = {
    "add": Input(storage).add
}