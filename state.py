class State:
    # состояния - считать новую команду или считать текст
    STATE_DEFAULT = 0
    STATE_INPUT = 1
    # дополнительно - стек для хранения команд

    def __init__(self) -> None:
        self.state = self.STATE_DEFAULT
        self.commands = []
        self.parameters = []


    def handleCommand(self):
        self.state = self.STATE_DEFAULT
        if len(self.commands) == 0:
            return None
        return self.commands.pop()


    def readInput(self, command, parameters=None):
        self.state = self.STATE_INPUT
        if parameters:
            self.parameters.append(parameters)
        self.commands.append(command)

    def get_parameters(self):
        if len(self.parameters) > 1:
            return self.parameters.pop()
        return None