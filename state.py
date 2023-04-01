class State:
    # состояния - считать новую команду или считать текст
    STATE_DEFAULT = 0
    STATE_INPUT = 1
    # дополнительно - стек для хранения команд

    def __init__(self) -> None:
        self.state = self.STATE_DEFAULT
        self.commands = []


    def handleCommand(self):
        self.state = self.STATE_DEFAULT
        return self.commands.pop()


    def readInput(self, command):
        self.state = self.STATE_INPUT
        self.commands.append(command)