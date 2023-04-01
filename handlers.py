from state import State
state = State()
import input as command_input


async def add(client, message):
    await message.reply(
        'Запишите название и сумму в формате *название. -100*. Валюту указывать не надо'
    )
    state.readInput('add')



async def today_sum(client, message):
    await message.reply(
        "Ваши доходы за сегодня составили 0 руб."
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
    pass

handlers = {
    "add": command_input.add
}