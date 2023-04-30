from state import State
from input import Input
from storage import Storage
from pyrogram.enums import ParseMode
from pyrogram.types import (InlineKeyboardMarkup, InlineKeyboardButton)
import auth
from config import DB_NAME

state = State()
storage = Storage(DB_NAME)
input = Input(storage)


def check_access(func):
    async def wrapper(client, message):
        if not auth.Access.check_access(message.from_user.id):
            await message.reply('Доступ запрещен. Обратитесь к владельцу, чтобы получить доступ.')
            return
        await func(client, message)
    return wrapper
    


async def add(client, message):
    date = message.command[1] if len(message.command) > 1 else None
    await message.reply(
        "Напишите транзакцию в формате `1000 перевод`",
        parse_mode=ParseMode.MARKDOWN
    )
    state.readInput('add', date)


async def today_sum(client, message):    
    sum = storage.get_today_sum(message.from_user.id)
    await message.reply(
        "Ваш итог за сегодня %+d руб." % sum,
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
    sum = storage.get_week_sum(message.from_user.id)
    await message.reply(
        "Ваш итог за неделю %+d руб." % sum,
    )


async def month_sum(client, message):
    sum = storage.get_month_sum(message.from_user.id)
    await message.reply(
        "Ваш итог за месяц %+d руб." % sum,
    )

async def read_input(client, message):
    lastCommand = state.handleCommand()

    # по умолчанию записывает новый расход
    if not lastCommand:
        input.add(message)
        await message.reply(
            "✅"
        )
        return
    
    handler =  handlers[lastCommand]
    handler(message, state.get_parameters())
    await message.reply(
        "✅"
    )

async def button_handler(client, callback_query):
    # todo считать последюю команду, для которой должен быть обработчик
    await callback_query.message.reply(
        '\n'.join(storage.get_today_list(callback_query.from_user.id))
    )


def default_handler(client, message):
    message.reply(
        "Вы хотели записать новую транзакцию? Этот бот понимает формат `[сумма][описание]` \nНапример, `-100 транспорт`",
        parse_mode=ParseMode.MARKDOWN
    )


def help(client, message):
    message.reply(
    """
Чтобы записать транзакцию, напишите сумму и ее описание через пробел, например `0 какая-то транзакция`.
Для суммы можно использовать знаки `-+`, например, 
`+1000 доход` или `-200 расход`. 
Сумма без знака будет интерпретироваться как доход, например, `100 штука`.
Валюту не указывайте.  
Между знаком и числом не должно быть пробелов. Сообщение вида `- 100 расход` бот попросит вас исправить.

По умолчанию бот добавляет транзакции за сегодня. Чтобы добавить за другой день, пишите `/add дд.мм.гггг`, например `/add 20.04.2023`.
    """,
        parse_mode=ParseMode.MARKDOWN
    )

handlers = {
    "add": input.add
}